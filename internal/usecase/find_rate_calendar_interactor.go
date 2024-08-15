package usecase

// import (
// 	"context"

// 	"github.com/Cirrus-Ltd/moves-clone-calendar/internal/domain"
// )

// type IFindRateInteractor interface {
// 	Execute(ctx context.Context, input *input.FindRateInput) (*output.FindRateOutput, error)
// }

// type FindRateInteractor struct {
// 	rateRepository Repository.IRateRepository
// }

// func NewFindRateInteractor(
// 	rateRepository repository.IRateRepository,
// ) *IFindRateInteractor {
// 	return &FindRateInteractor{
// 		rateRepository: rateRepository,
// 	}
// }

// // DTO input
// type FindRateInput struct {
// 	StartDate string
// 	EndDate   string
// }

// // DTO output
// type FindRateOutput struct {
// 	Rate []int `json:"rate"`
// }

// type FindRatePresenter struct {
// 	Output([]domain.RateCalendar) []FindRateOutput.Rate
// }

// func (i *FindRateInteractor) Execute(input FindRateInput) (*output.FindRateOutput, error) {
// 	rateCalendar, err := i.rateRepository.FindById(input)
// 	if err != nil {
// 		return i.presenter.Output([]domain.RateCalendar{}), err
// 	}
// 	return i.presenter.Output([]rateCalendar), nil
// }
