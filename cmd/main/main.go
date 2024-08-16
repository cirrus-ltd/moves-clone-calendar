package main

import (
	"log"

	"github.com/Cirrus-Ltd/moves-clone-calendar/internal/config"
	"github.com/Cirrus-Ltd/moves-clone-calendar/internal/infrastructure/database"
	"github.com/Cirrus-Ltd/moves-clone-calendar/internal/infrastructure/router"
	"github.com/Cirrus-Ltd/moves-clone-calendar/internal/interface/controllers"
	"github.com/Cirrus-Ltd/moves-clone-calendar/internal/interface/presenters"
	"github.com/Cirrus-Ltd/moves-clone-calendar/internal/usecase"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := database.NewDB(*cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: ")
	}

	// リポジトリの初期化
	rateCalendarRepository := database.NewRateCalendarRepository(db)

	// プレゼンターの初期化
	rateCalendarPresenter := presenters.NewRateCalendarPresenter()

	// インタラクターの初期化
	saveRateInteractor := usecase.NewSaveRateInteractor(rateCalendarRepository)

	// コントローラーの初期化
	rateCalendarController := controllers.NewRateCalendarController(saveRateInteractor, rateCalendarPresenter)

	// ルーターの初期化
	router.InitRouter(rateCalendarController)
}
