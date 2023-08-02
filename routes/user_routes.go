package routes

import (
	handlers "github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/handler"
	"github.com/labstack/echo/v4"
)

func SetupUserRoutes(e *echo.Echo, userHandler *handlers.UserHandler) {
	e.GET("/api/users", userHandler.GetAllUsers)       //ユーザー全取得
	e.GET("/api/users/:id", userHandler.GetUserByID)   //ユーザー個別取得
	e.POST("/api/users", userHandler.CreateUser)       //ユーザー作成
	e.PUT("/api/users/:id", userHandler.UpdateUser)    //ユーザー情報編集更新
	e.DELETE("/api/users/:id", userHandler.DeleteUser) //ユーザー削除
}
