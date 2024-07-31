package models

import "time"

type RateCalendar struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Version   string    `json:"version"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	RateDate  time.Time `json:"rate_date"`
	Rate      int       `json:"rate"`
}

func (RateCalendar) TableName() string {
	return "rate_calendar"
}
