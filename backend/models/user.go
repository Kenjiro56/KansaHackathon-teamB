package models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	UserName string `gorm:"size:255;not null" json:"user_name"`
	Email    string `gorm:"size:255;unique;not null" json:"email"`
	Password string `gorm:"size:255;not null" json:"password"`
}
