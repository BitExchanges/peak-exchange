package model

import "time"

// 当前期权订单
type OptionOrder struct {
	Id        int       `json:"id"`         //主键
	UserId    int       `json:"user_id"`    //用户ID
	OrderNo   string    `json:"order_no"`   //订单号
	Symbol    string    `json:"symbol"`     //标的物
	Direction int       `json:"direction"`  //方向 0-涨  1-跌
	OpenPrice float64   `json:"open_price"` //开仓价格
	Amount    int       `json:"amount"`     //开仓数量  USD
	Profit    float64   `json:"profit"`     //利润
	CloseAt   time.Time `json:"close_at"`   //结算时间
	CommonModel
}

// 历史期权订单
type OptionOrderHistory struct {
	OptionOrder
}
