package model

import "time"

type Token struct {
	CommonModel
	Token        string    `json:"token" gorm:"type:varchar(64)"`
	UserId       int       `json:"user_id"`
	Mobile       string    `json:"mobile"`
	ExpireAt     time.Time `json:"expire_at"`
	LastVerifyAt time.Time `json:"last_verify_at"`
}
