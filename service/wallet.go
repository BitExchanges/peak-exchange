package service

import (
	"errors"
	. "peak-exchange/model"
	"peak-exchange/utils"
)

func SelectAddressByPage(limit, page int) ([]Wallet, int, error) {
	var wallets []Wallet
	var count int
	db := utils.MainDbBegin()
	defer db.DbCommit()
	db.Table("wallet").Count(&count)
	if db.Offset(limit * (page - 1)).Limit(limit).Find(&wallets).RecordNotFound() {
		return nil, 0, errors.New("暂无数据")
	}
	return wallets, count, nil
}

//批量插入
func BatchSaveAddress(wallet Wallet) {
	db := utils.MainDbBegin()
	defer db.DbCommit()
	db.Create(&wallet)
}

// 根据用户ID 删除钱包地址
func DeleteWallet(userId, id int) {
	db := utils.MainDbBegin()
	defer db.CommonDB()
	db.Delete(&Wallet{}, "id=? AND user_id=?", id, userId)
}
