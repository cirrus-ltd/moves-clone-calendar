package domain

import "time"

type RateCalendar struct {
	id   time.Time
	rate int
}

func (r *RateCalendar) ID() time.Time {
	return r.id
}

func (r *RateCalendar) Rate() int {
	return r.rate
}
