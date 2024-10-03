package auth

import (
	constants "auth"
	"auth/pkg/domain/response"
	"auth/pkg/useCases/Helpers/jwtHelper"
	"auth/pkg/useCases/Helpers/responseHelper"
	"context"
	"net/http"
	"strings"
)

var jwtHelperInstance jwtHelper.JwtHelperImpl

// JWTMiddleware is a middleware function to validate JWT tokens
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the Authorization header from the request
		authHeader := r.Header.Get("Authorization")

		// Check if the Authorization header is present
		if authHeader == "" {
			response := response.Status{Text: "Missing Authorization header", Code: 401}
			responseHelper.WriteResponse(w, response, nil)
			return
		}

		// Ensure it's a Bearer token
		if !strings.HasPrefix(authHeader, "Bearer ") {
			response := response.Status{Text: "Invalid Authorization header", Code: 401}
			responseHelper.WriteResponse(w, response, nil)
			return
		}

		// Extract the token from the Authorization header
		tokenSlice := strings.Split(authHeader, " ")
		if len(tokenSlice) != 2 {
			response := response.Status{Text: "Invalid Authorization header", Code: 401}
			responseHelper.WriteResponse(w, response, nil)
			return
		}

		tokenString := tokenSlice[1]
		email, err := jwtHelperInstance.ValidateToken(tokenString, constants.AccessTokenSecret)
		if err != nil {
			response := response.Status{Text: "Invalid token", Code: 401}
			responseHelper.WriteResponse(w, response, nil)
			return
		}

		// Add the email to the request context and call the next handler
		ctx := context.WithValue(r.Context(), "email", email)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
