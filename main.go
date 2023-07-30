package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/config"
	"github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/database"
	"github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/routes"
)

func main() {
	// 設定情報の読み込み
	cfg := config.NewConfig()

	// データベース接続の設定
	db, err := database.NewDatabaseConnection(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// Echoインスタンスの作成
	e := echo.New()

	// Middlewareの設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティングの設定
	routes.Setup(e, db)

	// サーバー起動
	e.Logger.Fatal(e.Start(":8080"))
}
