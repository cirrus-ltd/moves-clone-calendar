package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/Cirrus-Ltd/moves-clone-calendar/internal/interface/presenters"
	"github.com/Cirrus-Ltd/moves-clone-calendar/internal/usecase"
	"github.com/labstack/echo/v4"
)

type RateCalendarController struct {
	saveRateInteractor    usecase.ISaveRateInteractor
	rateCalendarPresenter presenters.RateCalendarPresenter
	// FindRateInteractor usecase.IFindRateInteractor
}

func NewRateCalendarCalendar(
	saveRateInteractor usecase.ISaveRateInteractor,
	rateCalendarPresenter presenters.RateCalendarPresenter,
	// findRateInteractor usecase.IFindRateInteractor,
) *RateCalendarController {
	return &RateCalendarController{
		saveRateInteractor:    saveRateInteractor,
		rateCalendarPresenter: rateCalendarPresenter,
		// findRateInteractor: findRateInteractor,
	}
}

func (rc *RateCalendarController) RateRegister(c echo.Context) error {
	var input usecase.SaveRateInputData
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	// 日付の形式をチェック
	for dateStr := range input.DateRate {
		if _, err := time.Parse("2006-01-02", dateStr); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid date format"})
		}
	}

	output, err := rc.saveRateInteractor.Execute(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	ctx := context.Background()
	return rc.rateCalendarPresenter.SaveRateOutputPresenter(ctx, output)
}
