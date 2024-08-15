package presenters

import (
	"context"
	"errors"
	"net/http"

	"github.com/Cirrus-Ltd/moves-clone-calendar/internal/usecase"
	"github.com/labstack/echo/v4"
)

type RateCalendarPresenter struct {
	saveRateOutput usecase.ISaveRateOutput
}

func NewRateCalendarPresenter(
	saveRateOutput usecase.ISaveRateOutput,
) *RateCalendarPresenter {
	return &RateCalendarPresenter{
		saveRateOutput: saveRateOutput,
	}
}

func (p *RateCalendarPresenter) SaveRateOutputPresenter(ctx context.Context, output usecase.SaveRateOutputData) error {
	c, ok := ctx.(echo.Context)
	if !ok {
		return errors.New("invalid context type")
	}
	response := make([]map[string]interface{}, len(output.RateCalendars))
	for i, rateCalendar := range output.RateCalendars {
		response[i] = map[string]interface{}{
			"id":   rateCalendar.ID().Format("2006-01-02"),
			"rate": rateCalendar.Rate(),
		}
	}
	return c.JSON(http.StatusOK, response)
}
