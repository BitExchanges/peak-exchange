package model

import "time"

type CommonModel struct {
	Id       int       `json:"id" gorm:"primary_key"`
	CreateAt time.Time `json:"-"`
	UpdateAt time.Time `json:"-"`
}

type EmailMessage struct {
	Message string // 消息
	Type    int    // 类型 0-异地登录通知  1-充值通知
	Ip      string //ip地址
	Date    string //时间
	Head    string //头部
	Title   string //邮件标题
}
