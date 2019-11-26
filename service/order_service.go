package service

import (
	"fmt"
	. "peak-exchange/model"
	. "peak-exchange/utils"
	"strconv"
	"time"
)

// 创建订单
func SaveOrder(order Order) {

	orderNo := strconv.FormatInt(GenerateSnowflakeId(), 10)
	order.OrderNo = orderNo
	order.CreateAt = time.Now()
	order.UpdateAt = time.Now()

	DB := MainDbBegin()
	defer DB.DbRollback()
	result := DB.Create(&order)
	fmt.Println(result.RowsAffected)

}
