package authHelper

import (
	"errors"
	"net/http"
)

func GetEmailFromContext(r *http.Request) (string, error) {
	if r.Context().Value("email") == nil {
		return "", errors.New("email not found in context")
	}

	email := r.Context().Value("email").(string)
	return email, nil
}
