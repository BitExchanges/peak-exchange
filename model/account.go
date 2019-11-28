package model

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
	Type       int     `json:"type"`                               //类型
}
