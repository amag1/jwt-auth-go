package login

import "auth/pkg/domain/user"

type LoginResponse struct {
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refreshToken"`
	User         user.User `json:"user"`
}
