package service

import (
	. "peak-exchange/model"
	"peak-exchange/utils"
)

// 根据账户类型查询
// typ      资金类型 0-虚拟账户 1-真实账户
// userId   用户ID
// currency 资金名称
func SelectAccountBalance(typ, userId int, currency string) (account Account) {
	db := utils.MainDbBegin()
	defer db.DbCommit()
	db.Where("type=? AND user_id=? AND currency=?", typ, userId, currency).Find(&account)
	return account
}

// 查询用户当前账户列表
func SelectUserAccountList(userId int) (accountList []Account) {
	db := utils.MainDbBegin()
	defer db.DbCommit()
	db.Where("user_id=?", userId).Find(&accountList)
	return accountList
}

// 创建账户
func SaveAccount(account Account) {
	db := utils.MainDbBegin()
	defer db.Commit()
	db.Create(&account)
}
