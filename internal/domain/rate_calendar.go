package domain

import (
	"errors"
	"time"
)

type RateCalendar struct {
	id   time.Time
	rate int
}

func NewRateCalendar(idStr string, rate int) (*RateCalendar, error) {
	id, err := time.Parse("2006-01-02", idStr)
	if err != nil {
		return nil, errors.New("id must be in the format YYYY-MM-DD")
	}

	// バリデーション: rateが負の値でないことを確認
	if rate < 0 {
		return nil, errors.New("rate cannot be negative")
	}

	return &RateCalendar{
		id:   id,
		rate: rate,
	}, nil
}

func (r *RateCalendar) ID() time.Time {
	return r.id
}

func (r *RateCalendar) Rate() int {
	return r.rate
}
