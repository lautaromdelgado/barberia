package models

import "time"

type UserTokens struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	User_ID    uint      `json:"user_id" gorm:"not null;index"`
	User       User      `json:"user" gorm:"foreignKey:User_ID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Token      string    `json:"token"`
	Expires_At time.Time `json:"expires_at"`
	Created_At time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func (UserTokens) TableName() string {
	return "user_tokens"
}
