package model

type User struct {
	//CommonModel
	Id           int    `json:"id"`            //ID
	UUID         string `json:"uuid"`          //UUID 短号
	NickName     string `json:"nick_name"`     //昵称
	Name         string `json:"name"`          //姓名
	Avatar       string `json:"avatar"`        //头像
	TradePwd     string `json:"trade_pwd"`     //交易密码
	LoginPwd     string `json:"login_pwd"`     //登录密码
	Mobile       string `json:"mobile"`        //手机号
	Email        string `json:"email"`         //邮箱
	Level        string `json:"level"`         //用户等级
	KycLevel     string `json:"kyc_level"`     //认证等级
	IdentityCard string `json:"identity_card"` //身份证
	CardType     int    `json:"card_type"`     //证件类型
}
