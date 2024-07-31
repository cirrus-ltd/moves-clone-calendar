package domain

import (
	"context"
)

type RateCalendarRepository interface {
	Save(ctx context.Context, id string)
	FindByID(ctx context.Context, rateCalendar *RateCalendar) error
}
