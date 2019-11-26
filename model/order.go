package model

import (
	//. "github.com/shopspring/decimal"
	"time"
)

type Order struct {
	CommonModel
	Id         int       `json:"-" gorm:"primary_key "`
	UserId     int       `json:"user_id"`                                //用户ID
	OrderNo    string    `json:"order_no"`                               //订单号
	Symbol     string    `json:"symbol"`                                 //交易对
	State      int       `json:"-"`                                      //订单状态
	Deposit    float64   `gorm:"type:decimal(24,12)" json:"deposit"`     //保证金
	Type       string    `json:"type"`                                   //订单类型
	Price      float64   `gorm:"type:decimal(24,12)" json:"price"`       //委托价格
	Amount     float64   `gorm:"type:decimal(24,12)" json:"amount"`      //委托量
	Volume     float64   `gorm:"type:decimal(24,12)" json:"volume"`      //成交量
	AvgPrice   float64   `gorm:"type:decimal(24,12)" json:"avg_price"`   //成交均价
	HoldAmount float64   `gorm:"type:decimal(24,12)" json:"hold_amount"` //持仓量
	Received   float64   `gorm:"type:decimal(24,12)" json:"received"`    //已收到
	DealAt     time.Time `gorm:"default:null"`                           //成交时间
	Sn         string    `json:"sign"`                                   //订单签名

}
