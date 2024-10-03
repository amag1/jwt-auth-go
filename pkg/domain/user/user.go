package user

import (
	"auth/pkg/domain/refreshToken"
)

type User struct {
	Id            int                         `json:"id"`
	Email         string                      `json:"email"`
	Password      string                      `json:"password"`
	Name          string                      `json:"name"`
	RefreshTokens []refreshToken.RefreshToken `json:"-"`
}
