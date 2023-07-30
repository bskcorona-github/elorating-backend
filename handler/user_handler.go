package handlers

import (
	"net/http"

	"github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/models"
	"github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/usecase"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase // ポインタ型ではなくインターフェースそのものを保持するよう修正
}

func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

// ユーザー全取得
func (uh *UserHandler) GetAllUsers(c echo.Context) error {
	users, err := uh.userUsecase.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}

// ユーザー個別取得
func (uh *UserHandler) GetUserByID(c echo.Context) error {
	userID := c.Param("id")
	user, err := uh.userUsecase.GetUserByID(userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, "User not found")
	}
	return c.JSON(http.StatusOK, user)
}

// ユーザー作成
func (uh *UserHandler) CreateUser(c echo.Context) error {
	var newUser models.User
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}
	createdUser, err := uh.userUsecase.CreateUser(&newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create user")
	}
	return c.JSON(http.StatusCreated, createdUser)
}

// ユーザー情報編集更新
func (uh *UserHandler) UpdateUser(c echo.Context) error {
	userID := c.Param("id")
	var updatedUser *models.User
	if err := c.Bind(&updatedUser); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}
	updatedUser, err := uh.userUsecase.UpdateUser(userID, updatedUser)
	if err != nil {
		return c.JSON(http.StatusNotFound, "User not found")
	}
	return c.JSON(http.StatusOK, updatedUser)
}

// ユーザー削除
func (uh *UserHandler) DeleteUser(c echo.Context) error {
	userID := c.Param("id")
	err := uh.userUsecase.DeleteUser(userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, "User not found")
	}
	return c.NoContent(http.StatusNoContent)
}
