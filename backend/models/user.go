package models

type User struct {
	ID       uint   `gorm:"primarykey;autoIncrement" json:"id"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	Name     string `json:"name" gorm:"not null"`
	LastName string `json:"last_name" gorm:"not null"`
}
