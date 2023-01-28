package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID      uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Bonus   uint32 `gorm:"default:0" json:"bonus"`
	Name    string `gorm:"size:255;" json:"name"`
	Surname string `gorm:"size:100;" json:"surname"`
	Phone   string `gorm:"size:100;" json:"phone"`
	Email   string `gorm:"size:100;not null;" json:"email"`
}

func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

type Participate struct {
	UserID  int `json:"user_id"`
	EventID int `json:"event_id"`
}
