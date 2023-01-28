package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
	"time"
)

type Event struct {
	ID          uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Title       string `gorm:"size:255;" json:"title"`
	Description string `gorm:"size:255;" json:"description"`
	//Creator     User      `gorm:"foreignKey:CreatorID;references:ID" json:"creator"`
	CreatorID    uint32        `json:"creator_id"`
	Location     Location      `gorm:"embedded" json:"location"`
	Time         time.Time     `gorm:"default:current_timestamp" json:"time"`
	Limit        int           `gorm:"default:10" json:"limit"`
	Participants pq.Int32Array `gorm:"type:integer[]" json:"participants"`
	Status       string        `gorm:"size:255;default:'active'" json:"status"`
}

type Location struct {
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
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
