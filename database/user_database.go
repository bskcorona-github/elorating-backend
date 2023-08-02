package database

import (
	"github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/models"
	"gorm.io/gorm"
)

type UserDatabase struct {
	db *gorm.DB
}

func NewUserDatabase(db *gorm.DB) *UserDatabase {
	return &UserDatabase{
		db: db,
	}
}

// ユーザー全取得
func (udb *UserDatabase) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := udb.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// ユーザー個別取得
func (udb *UserDatabase) GetUserByID(userID string) (*models.User, error) {
	var user models.User
	if err := udb.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// ユーザー作成
func (udb *UserDatabase) CreateUser(user *models.User) (*models.User, error) {
	if err := udb.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// ユーザー情報編集更新
func (udb *UserDatabase) UpdateUser(userID string, updatedUser *models.User) (*models.User, error) {
	var user models.User
	if err := udb.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}

	// Update user fields here...
	// For example:
	// user.Name = updatedUser.Name
	// user.Age = updatedUser.Age
	// ...

	if err := udb.db.Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// ユーザー削除
func (udb *UserDatabase) DeleteUser(userID string) error {
	var user models.User
	if err := udb.db.Where("id = ?", userID).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
