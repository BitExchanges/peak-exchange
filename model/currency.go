package model

import "peak-exchange/utils"

type Currency struct {
	CommonModel
	Id          int    `json:"id" gorm:"primary_key"` //主键
	Key         string `json:"key"`
	Code        string `json:"code"`
	Symbol      string `json:"symbol"`
	Coin        string `json:"coin"`
	Precision   int    `json:"precision"`    //精度
	Erc20       bool   `json:"erc20"`        //erc20
	Erc23       bool   `json:"erc23"`        //erc23
	Visible     bool   `json:"visible"`      //可见
	Tradeable   bool   `json:"trade_able"`   //允许交易
	Depositable bool   `json:"deposit_able"` //
}

var AllCurrencies []Currency

// 初始化所有可用交易币种
func InitAllCurrency(db *utils.GormDB) {
	db.Where("visible=?", true).Find(&AllCurrencies)
}

// 是否基于以太开发
func (currency Currency) IsEthereum() (result bool) {
	if currency.Code == "eth" || currency.Erc20 || currency.Erc23 {
		result = true
	}
	return
}
