package usecase

import (
	"context"

	"github.com/Cirrus-Ltd/moves-clone-calendar/internal/domain"
)

// Input Boundary <I>
type ISaveRateInteractor interface {
	Execute(input SaveRateInputData) (SaveRateOutputData, error)
}

// Output Boundary <I>
type ISaveRateOutput interface {
	SaveRateOutputPresenter(ctx context.Context, output SaveRateOutputData) error
}

// Input Data <DS>
type SaveRateInputData struct {
	DateRate map[string]int
}

// Output Data <DS>
type SaveRateOutputData struct {
	RateCalendars []domain.RateCalendar
}

// Data Access
type SaveRateInteractor struct {
	rateRepository IRateCalendarRepository
}

func NewSaveRateInteractor(
	rateRepository IRateCalendarRepository,
) ISaveRateInteractor {
	return &SaveRateInteractor{
		rateRepository: rateRepository,
	}
}

func (i *SaveRateInteractor) Execute(input SaveRateInputData) (SaveRateOutputData, error) {
	var rateCalendars []domain.RateCalendar
	for dateStr, rate := range input.DateRate {
		rateCalendar, err := domain.NewRateCalendar(dateStr, rate)
		if err != nil {
			return SaveRateOutputData{}, err
		}
		rateCalendars = append(rateCalendars, *rateCalendar)
	}

	err := i.rateRepository.Save(rateCalendars)
	if err != nil {
		return SaveRateOutputData{}, err
	}
	outputData := SaveRateOutputData{RateCalendars: rateCalendars}

	return outputData, nil
}
