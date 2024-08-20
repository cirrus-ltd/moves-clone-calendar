package database

import (
	"github.com/Cirrus-Ltd/moves-clone-calendar/internal/domain"
	"github.com/Cirrus-Ltd/moves-clone-calendar/internal/infrastructure/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RateCalendarRepository struct {
	db *gorm.DB
}

func NewRateCalendarRepository(db *gorm.DB) *RateCalendarRepository {
	return &RateCalendarRepository{db: db}
}

func (r *RateCalendarRepository) Save(rateCalendars []domain.RateCalendar) error {
	var modelRateCalendars []model.RateCalendar
	for _, rateCalendar := range rateCalendars {
		modelRateCalendars = append(modelRateCalendars, model.RateCalendar{
			ID:      rateCalendar.ID(),
			Rate:    rateCalendar.Rate(),
			Version: 1,
		})
	}
	// 複数のレコードを一度に登録し、主キーが重複している場合は上書き
	if err := r.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&modelRateCalendars).Error; err != nil {
		return err
	}
	return nil
}
