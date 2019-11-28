package service

import (
	. "peak-exchange/model"
	"peak-exchange/utils"
)

// 根据账户类型查询
// typ      资金类型 0-虚拟账户 1-真实账户
// userId   用户ID
// currency 资金名称
func SelectUserAccount(typ, userId int, currency string) (account Account) {
	db := utils.MainDbBegin()
	defer db.DbCommit()
	db.Where("type=? AND user_id=? AND currency=?", typ, userId, currency).Find(&account)
	return account
}

// 创建账户
func SaveAccount(account Account) {

}
