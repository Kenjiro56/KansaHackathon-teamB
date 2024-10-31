package models

import "time"

type Todo struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `json:"user_id"`
	User        User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	ObjID       uint      `json:"obj_id"`
	Obj         Obj       `gorm:"foreignKey:ObjID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"obj"`
	Status      bool      `gorm:"not null;default:false" json:"status"` //true = DONE
	Progress    int       `gorm:"not null;default:0" json:"progress"`
	MaxProgress int       `gorm:"not null;default:1" json:"max_progress"`
	DeleteFlag  bool      `gorm:"not null;default:false" json:"delete_flag"` //true = 削除済み
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
