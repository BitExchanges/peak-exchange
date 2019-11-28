package model

import "time"

var (
	UNKNOWN    = 0
	FIX        = 1
	STRIKE_FEE = 100
)

type Account struct {
	CommonModel
	UserId     int     `json:"user_id"`                            //用户ID
	CurrencyId int     `json:"currency_id"`                        //币种ID
	Currency   string  `json:"currency"`                           //币种
	Balance    float64 `json:"balance" gorm:"type:decimal(24,12)"` //当前资金
	Locked     float64 `json:"locked" gorm:"type:decimal(24,12)"`  //冻结资金
	Type       int     `json:"type"`                               //类型 0真实资金 1虚拟资金
}

// 创建虚拟账户
func CreateVirtualAccount(userId int) *Account {
	return &Account{
		CommonModel: CommonModel{CreateAt: time.Now(), UpdateAt: time.Now()},
		UserId:      userId,
		CurrencyId:  0,
		Currency:    "USD",
		Balance:     10000,
		Locked:      0,
		Type:        1,
	}
}

// 创建真实账户
func CreateRealAccount(userId int) *Account {
	return &Account{
		CommonModel: CommonModel{CreateAt: time.Now(), UpdateAt: time.Now()},
		UserId:      userId,
		CurrencyId:  0,
		Currency:    "USD",
		Balance:     0,
		Locked:      0,
		Type:        0,
	}
}
