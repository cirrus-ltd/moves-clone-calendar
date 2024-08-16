package usecase

import "github.com/Cirrus-Ltd/moves-clone-calendar/internal/domain"

// <I> Data Access Interface
type IRateCalendarRepository interface {
	Save(rateCalendars []domain.RateCalendar) error
	// FindById()
}
