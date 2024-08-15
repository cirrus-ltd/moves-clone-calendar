package model

import "time"

type RateCalendar struct {
	ID        time.Time `gorm:"primaryKey;type:date" json:"id"`
	Version   string    `json:"version"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Rate      int       `json:"rate"`
}

func (RateCalendar) TableName() string {
	return "rate_calendar"
}
