package models

import "time"

type Obj struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	UserID     uint      `json:"user_id"`
	User       User      `gorm:"foreignKey:UserID" json:"user"`
	ObjTitle   string    `gorm:"size:255;not null" json:"obj_title"`
	DeleteFlag bool      `gorm:"not null;default:false" json:"delete_flag"` //true = 削除済み
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
