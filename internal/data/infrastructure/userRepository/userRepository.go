package userRepository

import (
	"auth/pkg/domain/response"
	"auth/pkg/domain/user"
	"auth/pkg/useCases/Helpers/databaseHelper"

	"gorm.io/gorm"
)

type UserRepository struct {
}

type Repository interface {
	CreateUser(user *user.User) (*user.User, response.Status)
	UpdateUser(user user.User) response.Status
	UpdateUserPassword(user user.User) response.Status
	DeleteUserById(userId int) response.Status
	GetUserByEmail(email string) (user.User, response.Status)
}

func (ur *UserRepository) CreateUser(user *user.User) (*user.User, response.Status) {
	db := databaseHelper.Db
	result := db.Omit("id").Create(user)
	if result.Error != nil {
		if result.Error.Error() == gorm.ErrDuplicatedKey.Error() {
			return user, response.DBItemAlreadyExists
		}
		return user, response.DBExecutionError
	}

	return user, response.SuccessfulCreation
}

func (ur *UserRepository) GetUserById(userId uint) (user.User, response.Status) {
	var user user.User
	db := databaseHelper.Db
	result := db.Omit("Password").Preload("RefreshTokens").First(&user, userId)

	if err := result.Error; err != nil {
		if result.Error.Error() == gorm.ErrRecordNotFound.Error() {
			return user, response.NotFound
		}
		return user, response.DBExecutionError
	}

	return user, response.SuccessfulSearch
}

func (ur *UserRepository) UpdateUser(user user.User) response.Status {
	db := databaseHelper.Db
	result := db.Omit("password").Save(&user)
	if err := result.Error; err != nil {
		if result.Error.Error() == gorm.ErrRecordNotFound.Error() {
			return response.NotFound
		}
		return response.DBExecutionError
	}
	return response.SuccessfulUpdate
}

func (ur *UserRepository) UpdateUserPassword(user user.User) response.Status {
	db := databaseHelper.Db
	result := db.Select("password").Save(&user)
	if err := result.Error; err != nil {
		if result.Error.Error() == gorm.ErrRecordNotFound.Error() {
			return response.NotFound
		}
		return response.DBExecutionError
	}
	return response.SuccessfulUpdate
}

func (ur *UserRepository) DeleteUserById(userId int) response.Status {
	db := databaseHelper.Db
	result := db.Delete(&user.User{}, userId)
	if err := result.Error; err != nil {
		if result.Error.Error() == gorm.ErrRecordNotFound.Error() {
			return response.FailedSearch
		}
		return response.DBExecutionError
	}
	return response.SuccessfulDeletion
}

func (ur *UserRepository) GetUserByEmail(email string) (user.User, response.Status) {
	var user user.User
	db := databaseHelper.Db
	result := db.Preload("RefreshTokens").Where("email = ?", email).First(&user)
	if err := result.Error; err != nil {
		if result.Error.Error() == gorm.ErrRecordNotFound.Error() {
			return user, response.NotFound
		}
		return user, response.DBExecutionError
	}
	return user, response.SuccessfulSearch
}
