package model

import (
	"peak-exchange/utils"
	"time"
)

// 用户
type User struct {
	CommonModel
	ID           int       `json:"id"`                                   //ID
	UUID         string    `json:"uuid"`                                 //UUID 短号
	NickName     string    `json:"nick_name"`                            //昵称
	Name         string    `json:"name"`                                 //姓名
	Avatar       string    `json:"avatar"`                               //头像
	TradePwd     string    `json:"trade_pwd"`                            //交易密码
	LoginPwd     string    `valid:"string,min=6,max=12"json:"login_pwd"` //登录密码
	Mobile       string    `valid:"string,min=11,max=11" json:"mobile"`  //手机号
	Email        string    `json:"email"`                                //邮箱
	Level        string    `json:"level"`                                //用户等级
	KycLevel     string    `json:"kyc_level"`                            //认证等级
	IdentityCard string    `json:"identity_card"`                        //身份证
	CardType     int       `json:"card_type"`                            //证件类型
	LastLoginAt  time.Time `json:"last_login_at"`                        //最后登录时间
	LastLoginIp  string    `json:"last_login_ip"`                        //最后登录IP
	Token        string    `json:"token" gorm:"-"`                       //token标识
	Country      int       `json:"country"`                              //国家标识
}

// 常用地管理
type AuthLoginAddress struct {
	Id        int       `json:"id"`         //主键
	UserId    int       `json:"user_id"`    //用户ID
	Address   string    `json:"address"`    //地址名称
	IpAddress string    `json:"ip_address"` //IP
	State     int       `json:"state"`      //状态  0-未确认  1-已确认
	LoginType string    `json:"login_type"` //登录设备类型
	LoginAt   time.Time `json:"login_at"`   //登录时间
	CommonModel
}

// 钱包
type Wallet struct {
	Id         int    `json:"id"`          //主键
	UserId     int    `json:"user_id"`     //用户ID
	Address    string `json:"address"`     //钱包地址
	Currency   string `json:"currency"`    //币种
	Type       string `json:"type"`        //类型 erc20 | omini |trc20
	PrivateKey string `json:"private_key"` //私钥
	CommonModel
}

// 生成新钱包
func NewWallet(userId int, privateKey, address string) Wallet {
	return Wallet{
		Id:          0,
		UserId:      userId,
		Address:     address,
		Currency:    "usd",
		Type:        "erc20",
		PrivateKey:  privateKey,
		CommonModel: CommonModel{CreateAt: time.Now(), UpdateAt: time.Now()},
	}
}

func (user *User) SendEmail(typ int, ip string) {
	utils.SendEmail("769558579@qq.com", user.Email, "异地登录通知", ip)
}
