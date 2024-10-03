package refreshTokenRepository

import (
	"auth/pkg/domain/refreshToken"
	"auth/pkg/domain/response"
	"auth/pkg/useCases/Helpers/databaseHelper"

	"gorm.io/gorm"
)

type RefreshTokenRepository struct {
}

type Repository interface {
	CreateRefreshToken(refreshToken *refreshToken.RefreshToken) (*refreshToken.RefreshToken, response.Status)
	DeleteRefreshToken(userId int, token string) response.Status
	GetRefreshToken(userId int, token string) (*refreshToken.RefreshToken, response.Status)
}

func (rr *RefreshTokenRepository) CreateRefreshToken(refreshToken *refreshToken.RefreshToken) (*refreshToken.RefreshToken, response.Status) {
	db := databaseHelper.Db
	result := db.Omit("id").Omit("created_at").Create(refreshToken)
	if result.Error != nil {
		if result.Error.Error() == gorm.ErrDuplicatedKey.Error() {
			return refreshToken, response.DBItemAlreadyExists
		}
		return refreshToken, response.DBExecutionError
	}

	return refreshToken, response.SuccessfulCreation
}

func (rr *RefreshTokenRepository) DeleteRefreshToken(userId int, token string) response.Status {
	db := databaseHelper.Db
	result := db.Where("user_id = ? AND token = ?", userId, token).Delete(&refreshToken.RefreshToken{})
	if result.Error != nil {
		return response.DBExecutionError
	}
	if result.RowsAffected == 0 {
		return response.NotFound
	}
	return response.SuccessfulDeletion
}

func (rr *RefreshTokenRepository) GetRefreshToken(userId int, token string) (*refreshToken.RefreshToken, response.Status) {
	db := databaseHelper.Db
	var refresh refreshToken.RefreshToken
	result := db.Where("user_id = ? AND token = ?", userId, token).First(&refresh)
	if result.Error != nil {
		if result.Error.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, response.NotFound
		}
		return nil, response.DBExecutionError
	}
	return &refresh, response.SuccessfulSearch
}
