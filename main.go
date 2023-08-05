package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/database"
	handlers "github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/handler"
	"github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/repository"
	"github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/routes"
	"github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/usecase"
)

func main() {
	// 設定情報の読み込み
	//cfg := config.NewConfig()

	// データベース接続の設定
	dbURI := "host=localhost port=5432 user=postgres dbname=postgres password=tkz2001r sslmode=disable"
	db, err := database.NewDatabase(dbURI)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// PlayerRepositoryの作成
	playerRepository := repository.NewPlayerRepository(db.GetDB())

	// ELORepositoryの作成
	eloRepository := repository.NewELORepository(db.GetDB())

	// PlayerUsecaseの作成
	playerUsecase := usecase.NewPlayerUsecase(playerRepository)

	// ELOUsecaseの作成
	eloUsecase := usecase.NewELOUsecase(*playerRepository, eloRepository)

	// プレイヤーハンドラの作成
	playerHandler := handlers.NewPlayerHandler(playerUsecase)

	// ELOハンドラの作成
	eloHandler := handlers.NewELOHandler(*eloUsecase)

	// Echoインスタンスの作成
	e := echo.New()

	// Middlewareの設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティングの設定
	routes.RegisterPlayerRoutes(e, playerHandler)
	routes.RegisterELORoutes(e, eloHandler)

	// サーバー起動
	e.Logger.Fatal(e.Start(":8080"))
}
