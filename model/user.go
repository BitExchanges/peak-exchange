package model

import "time"

type User struct {
	CommonModel
	Id           int       `json:"-"`                                    //ID
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
}
