package routes

import (
	"auth/pkg/auth"
	"auth/pkg/domain/login"
	"auth/pkg/domain/response"
	"auth/pkg/domain/user"
	"auth/pkg/useCases/Helpers/responseHelper"
	"auth/pkg/useCases/services"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

type UserRouter struct {
	Handler *services.UserService
}

func (ur UserRouter) Login(w http.ResponseWriter, r *http.Request) {
	var credentials login.Login
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		responseHelper.WriteResponse(w, response.BadRequest, nil)
		return
	}

	response, status := ur.Handler.Login(credentials)
	responseHelper.WriteResponse(w, status, response)
}

func (ur UserRouter) RefreshToken(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		responseHelper.WriteResponse(w, response.Unauthorized, nil)
		return
	}

	response, status := ur.Handler.RefreshToken(token)
	responseHelper.WriteResponse(w, status, response)
}

func (ur UserRouter) Logout(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		responseHelper.WriteResponse(w, response.Unauthorized, nil)
		return
	}

	status := ur.Handler.Logout(token)
	responseHelper.WriteResponse(w, status, nil)
}

func (ur UserRouter) Register(w http.ResponseWriter, r *http.Request) {
	var user user.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		responseHelper.WriteResponse(w, response.BadRequest, nil)
		return
	}

	response, status := ur.Handler.Register(user)
	responseHelper.WriteResponse(w, status, response)
}

func (ur UserRouter) ProtectedRoute(w http.ResponseWriter, r *http.Request) {
	responseHelper.WriteResponse(w, response.SuccessfulSearch, nil)
}

func (ur UserRouter) Routes() http.Handler {
	r := chi.NewRouter()

	r.Post("/login", ur.Login)
	r.Post("/register", ur.Register)
	r.Post("/refresh", ur.RefreshToken)
	r.Post("/logout", ur.Logout)

	r.With(auth.JWTMiddleware).Get("/protected", ur.ProtectedRoute)
	return r
}
