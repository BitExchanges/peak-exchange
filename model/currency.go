package model

import "peak-exchange/utils"

type Currency struct {
	CommonModel
	Key         string `json:"key"`
	Code        string `json:"code"`
	Symbol      string `json:"symbol"`
	Coin        string `json:"coin"`
	Precision   int    `json:"precision"`
	Erc20       bool   `json:"erc20"`
	Erc23       bool   `json:"erc23"`
	Visible     bool   `json:"visible"`
	Tradable    bool   `json:"tradable"`
	Depositable bool   `json:"depositable"`
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