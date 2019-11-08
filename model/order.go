package model

type Order struct {
	Id      int    `json:"id"`
	OrderNo string `json:"order_no" binding:"required"`
}
