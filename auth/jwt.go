package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

var (
	TokenExpired     error  = errors.New("token失效")
	TokenNotValidYet error  = errors.New("token未激活")
	TokenMalformed   error  = errors.New("非token")
	TokenInvalid     error  = errors.New("无法解析token")
	SignKey          string = "peak_exchange"
)

type Claims struct {
	Mobile   string `json:"mobile"`
	Id       int    `json:"id"`
	LoginPwd string `json:"login_pwd"`
	jwt.StandardClaims
}

type JWT struct {
	SigningKey []byte
}

// 获取签名
func GetSignKey() string {
	return SignKey
}

// 设置签名
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

// 生成JWT实例
func NewJwt() *JWT {
	return &JWT{[]byte(GetSignKey())}
}

// 创建token
func (j *JWT) CreateToken(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 解析token
func (j *JWT) ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

// 刷新token
//func (j *JWT)RefreshToken(tokenStr string)(string,error)  {
//
//}
