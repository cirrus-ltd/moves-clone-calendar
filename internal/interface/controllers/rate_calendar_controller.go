package controllers

import (
	"log"
	"time"

	"github.com/Cirrus-Ltd/moves-clone-calendar/internal/interface/presenters"
	"github.com/Cirrus-Ltd/moves-clone-calendar/internal/usecase"
	"github.com/labstack/echo/v4"
)

type RateCalendarController struct {
	saveRateInteractor usecase.ISaveRateInteractor
	presenter          presenters.IPresenter
}

func NewRateCalendarController(
	saveRateInteractor usecase.ISaveRateInteractor,
	presenter presenters.IPresenter,
) *RateCalendarController {
	return &RateCalendarController{
		saveRateInteractor: saveRateInteractor,
		presenter:          presenter,
	}
}

func (rc *RateCalendarController) RateRegister(c echo.Context) error {
	var input usecase.SaveRateInputData
	if err := c.Bind(&input); err != nil {
		return rc.presenter.PresentBadRequest(c, "Invalid request")
	}
	// 日付の形式をチェック
	for dateStr := range input.DateRate {
		if _, err := time.Parse("2006-01-02", dateStr); err != nil {
			return rc.presenter.PresentBadRequest(c, "Invalid date format")
		}
	}
	log.Printf("Received input: %+v", input.DateRate)
	output, err := rc.saveRateInteractor.Execute(input)
	if err != nil {
		log.Printf("Error executing interactor: %v", err)
		return rc.presenter.PresentInternalServerError(c, err)
	}

	return rc.presenter.SaveRateOutputPresenter(c, output)
}
