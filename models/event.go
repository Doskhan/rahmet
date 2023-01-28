package models

import (
	"database/sql/driver"
	"encoding/json"
	"gorm.io/gorm"
)

type Event struct {
	ID          uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Title       string `gorm:"size:255;" json:"title"`
	Description string `gorm:"size:255;" json:"description"`
	Category    string `gorm:"size:255;" json:"category"`
	//Creator      User          `gorm:"foreignKey:CreatorID;references:ID" json:"creator"`
	CreatorID    uint32    `json:"creator_id"`
	Location     string    `gorm:"size:255" json:"location"`
	Time         int       `json:"time"`
	Limit        int       `gorm:"default:10" json:"limit"`
	Bonus        int       `gorm:"default:0" json:"bonus"`
	Participants UserArray `gorm:"column:participants;embedded" json:"participants"`
	Status       string    `gorm:"size:255;default:'active'" json:"status"`
}

type UserArray []User

func (sla *UserArray) Scan(src interface{}) error {
	return json.Unmarshal([]byte(src.(string)), &sla)
}

func (sla UserArray) Value() (driver.Value, error) {
	val, err := json.Marshal(sla)
	return string(val), err
}

func (e *Event) FindAllEvents(db *gorm.DB) (*[]Event, error) {
	var err error
	events := []Event{}
	err = db.Debug().Model(&Event{}).Limit(100).Find(&events).Error
	if err != nil {
		return &[]Event{}, err
	}
	return &events, err
}
