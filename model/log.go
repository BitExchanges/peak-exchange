package model

import "time"

type Log struct {
	Id      int       `json:"id"`       //主键
	UserId  int       `json:"user_id"`  //用户ID
	Event   string    `json:"event"`    //事件   登录/退出/注册/交易/充值
	CreatAt time.Time `json:"creat_at"` //创建时间
}
