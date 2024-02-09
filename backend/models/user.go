package models

type User struct {
	ID       uint   `gorm:"primarykey;autoIncrement" json:"id"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}
