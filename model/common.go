package model

import "time"

type CommonModel struct {
	Id       int       `json:"id" gorm:"primary_key"`
	CreateAt time.Time `json:"-"`
	UpdateAt time.Time `json:"-"`
}
