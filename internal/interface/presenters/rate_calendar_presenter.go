package presenters

import (
	"net/http"

	"github.com/Cirrus-Ltd/moves-clone-calendar/internal/usecase"
	"github.com/labstack/echo/v4"
)

type IPresenter interface {
	usecase.ISaveRateOutput
	PresentBadRequest(c echo.Context, message string) error
	PresentInternalServerError(c echo.Context, err error) error
}

type RateCalendarPresenter struct{}

func NewRateCalendarPresenter() *RateCalendarPresenter {
	return &RateCalendarPresenter{}
}

func (p *RateCalendarPresenter) SaveRateOutputPresenter(c echo.Context, output usecase.SaveRateOutputData) error {
	response := make([]map[string]interface{}, len(output.RateCalendars))
	for i, rateCalendar := range output.RateCalendars {
		response[i] = map[string]interface{}{
			"id":   rateCalendar.ID().Format("2006-01-02"),
			"rate": rateCalendar.Rate(),
		}
	}
	return c.JSON(http.StatusOK, response)
}

func (p *RateCalendarPresenter) PresentBadRequest(c echo.Context, message string) error {
	response := struct {
		Error string `json:"error"`
	}{
		Error: message,
	}

	return c.JSON(http.StatusBadRequest, response)
}

func (p *RateCalendarPresenter) PresentInternalServerError(c echo.Context, err error) error {
	response := struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	}

	return c.JSON(http.StatusInternalServerError, response)
}
