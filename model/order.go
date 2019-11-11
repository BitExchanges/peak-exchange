package model

import (
	. "github.com/shopspring/decimal"
)

type Order struct {
	CommonModel
	UserId     int     `json:"user_id"`     //用户ID
	OrderNo    string  `json:"order_no"`    //订单号
	Symbol     string  `json:"symbol"`      //交易对
	State      int     `json:"-"`           //订单状态
	Deposit    Decimal `json:"deposit"`     //保证金
	Type       string  `json:"type"`        //订单类型
	Price      Decimal `json:"price"`       //委托价格
	Amount     Decimal `json:"amount"`      //委托量
	Volume     Decimal `json:"volume"`      //成交量
	AvgPrice   Decimal `json:"avg_price"`   //成交均价
	HoldAmount Decimal `json:"hold_amount"` //持仓量
	Received   Decimal `json:"received"`    //已收到
	Sn         string  `json:"sign"`        //订单签名

}
