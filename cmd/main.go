package main

import (
	"log"

	handlers "github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/internal/api/handler"
	"github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/internal/api/routes"
	"github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/internal/infrastructures/database"
	"github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/internal/repository"
	"github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/internal/usecase"
	"github.com/labstack/echo/v4"
)

func main() {
	// データベース接続の初期化
	db, err := database.NewDatabaseConnection()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// データベースマイグレーション
	err = database.RunMigrations(db)
	if err != nil {
		log.Fatal("Database migration failed:", err)
	}

	// Repository層の初期化
	playerRepository := repository.NewPlayerRepository(db)

	// Usecase層の初期化
	playerUsecase := usecase.NewPlayerUsecase(playerRepository)

	// Handler層の初期化
	playerHandler := handlers.NewPlayerHandler(playerUsecase)

	// Echoフレームワークのインスタンス化
	e := echo.New()

	// ルーティングの設定
	registerRoutes(e, playerHandler)

	// Webサーバーの起動
	e.Start(":8080")
}

func registerRoutes(e *echo.Echo, playerHandler *handlers.PlayerHandler) {
	// Player関連のルーティングを登録
	routes.RegisterPlayerRoutes(e, *playerHandler)
}
