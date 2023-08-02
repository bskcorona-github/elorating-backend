package repository

import (
	"github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/database"
	"github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/models"
)

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(userID string) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	UpdateUser(userID string, updatedUser *models.User) (*models.User, error)
	DeleteUser(userID string) error
}

type userRepository struct {
	db database.UserDatabase // ポインタ型ではなくインターフェースそのものを保持するよう修正
}

func NewUserRepository(db database.UserDatabase) UserRepository {
	return &userRepository{
		db: db,
	}
}

// ユーザー全取得
func (ur *userRepository) GetAllUsers() ([]models.User, error) {
	return ur.db.GetAllUsers()
}

// ユーザー個別取得
func (ur *userRepository) GetUserByID(userID string) (*models.User, error) {
	return ur.db.GetUserByID(userID)
}

// ユーザー作成
func (ur *userRepository) CreateUser(user *models.User) (*models.User, error) {
	return ur.db.CreateUser(user)
}

// ユーザー情報編集更新
func (ur *userRepository) UpdateUser(userID string, updatedUser *models.User) (*models.User, error) {
	return ur.db.UpdateUser(userID, updatedUser)
}

// ユーザー削除
func (ur *userRepository) DeleteUser(userID string) error {
	return ur.db.DeleteUser(userID)
}
