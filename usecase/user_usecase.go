package usecase

import (
	"github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/models"
	"github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/repository"
)

type UserUsecase interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(userID string) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	UpdateUser(userID string, updatedUser *models.User) (*models.User, error)
	DeleteUser(userID string) error
}

type userUsecase struct {
	userRepo repository.UserRepository // ポインタ型ではなくインターフェースそのものを保持するよう修正
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

// ユーザー全取得
func (uu *userUsecase) GetAllUsers() ([]models.User, error) {
	return uu.userRepo.GetAllUsers()
}

// ユーザー個別取得
func (uu *userUsecase) GetUserByID(userID string) (*models.User, error) {
	return uu.userRepo.GetUserByID(userID)
}

// ユーザー作成
func (uu *userUsecase) CreateUser(user *models.User) (*models.User, error) {
	return uu.userRepo.CreateUser(user)
}

// ユーザー情報編集更新
func (uu *userUsecase) UpdateUser(userID string, updatedUser *models.User) (*models.User, error) {
	return uu.userRepo.UpdateUser(userID, updatedUser)
}

// ユーザー削除
func (uu *userUsecase) DeleteUser(userID string) error {
	return uu.userRepo.DeleteUser(userID)
}
