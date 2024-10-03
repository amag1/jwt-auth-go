package routes

import (
	"auth/internal/data/infrastructure/refreshTokenRepository"
	"auth/internal/data/infrastructure/userRepository"
	"auth/pkg/useCases/Helpers/encoder"
	"auth/pkg/useCases/Helpers/jwtHelper"
	"auth/pkg/useCases/services"
	"net/http"

	"github.com/go-chi/chi"
)

var (
	INTERNAL_SERVER_ERROR = []byte("500: Internal Server Error")
	ERR_ALREADY_COMMITTED = "already been committed"
)

func New() http.Handler {
	r := chi.NewRouter()

	ur := UserRouter{
		Handler: &services.UserService{
			UserRepository:         &userRepository.UserRepository{},
			RefreshTokenRepository: &refreshTokenRepository.RefreshTokenRepository{},
			Encoder:                &encoder.EncoderImpl{},
			JwtHelper:              &jwtHelper.JwtHelperImpl{},
		},
	}

	r.Mount("/user", ur.Routes())

	return r
}
