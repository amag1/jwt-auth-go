package services

import (
	constants "auth"
	"auth/internal/data/infrastructure/refreshTokenRepository"
	"auth/internal/data/infrastructure/userRepository"
	"auth/pkg/domain/login"
	"auth/pkg/domain/refreshToken"
	"auth/pkg/domain/response"
	"auth/pkg/domain/user"
	"auth/pkg/useCases/Helpers/encoder"
	"auth/pkg/useCases/Helpers/jwtHelper"
)

type UserService struct {
	UserRepository         userRepository.Repository
	RefreshTokenRepository refreshTokenRepository.Repository
	Encoder                encoder.Encoder
	JwtHelper              jwtHelper.JwtHelper
}

func (us UserService) Login(credentials login.Login) (login.LoginResponse, response.Status) {
	user, status := us.UserRepository.GetUserByEmail(credentials.Email)
	if status != response.SuccessfulSearch {
		return login.LoginResponse{}, status
	}

	if err := us.Encoder.ComparePasswords(user.Password, credentials.Password); err != nil {
		return login.LoginResponse{}, response.Unauthorized
	}

	refresh, err := us.JwtHelper.GenerateToken(user.Email, constants.RefreshTokenSecret, constants.RefreshTokenExpireTime)
	if err != nil {
		return login.LoginResponse{}, response.InternalServerError
	}

	_, status = us.RefreshTokenRepository.CreateRefreshToken(&refreshToken.RefreshToken{
		UserId: user.Id,
		Token:  refresh,
	})

	if status != response.SuccessfulCreation {
		return login.LoginResponse{}, status
	}

	access, err := us.JwtHelper.GenerateToken(user.Email, constants.AccessTokenSecret, constants.AccessTokenExpireTime)
	if err != nil {
		return login.LoginResponse{}, response.InternalServerError
	}

	user.Password = ""
	return login.LoginResponse{
		AccessToken:  access,
		RefreshToken: refresh,
		User:         user,
	}, response.SuccessfulCreation
}

func (us UserService) RefreshToken(token string) (login.LoginResponse, response.Status) {
	email, err := us.JwtHelper.ValidateToken(token, constants.RefreshTokenSecret)
	if err != nil {
		return login.LoginResponse{}, response.Unauthorized
	}

	user, status := us.UserRepository.GetUserByEmail(email)
	if status != response.SuccessfulSearch {
		return login.LoginResponse{}, response.Unauthorized
	}
	user.Password = ""

	// Check if the refresh token exists
	_, status = us.RefreshTokenRepository.GetRefreshToken(user.Id, token)
	if status != response.SuccessfulSearch {
		return login.LoginResponse{}, response.Unauthorized
	}

	accessToken, err := us.JwtHelper.GenerateToken(user.Email, constants.AccessTokenSecret, constants.AccessTokenExpireTime)
	if err != nil {
		return login.LoginResponse{}, response.InternalServerError
	}

	return login.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: token,
		User:         user,
	}, response.SuccessfulCreation
}

func (us UserService) Logout(token string) response.Status {
	// Validate the refresh token
	email, err := us.JwtHelper.ValidateToken(token, constants.RefreshTokenSecret)
	if err != nil {
		return response.Unauthorized
	}

	// Get the user by email
	user, status := us.UserRepository.GetUserByEmail(email)
	if status != response.SuccessfulSearch {
		return status
	}

	// Invalidate the refresh token by removing it from the repository
	status = us.RefreshTokenRepository.DeleteRefreshToken(user.Id, token)
	if status != response.SuccessfulDeletion {
		return response.InternalServerError
	}

	return response.SuccessfulDeletion
}

func (us UserService) Register(newUser user.User) (user.User, response.Status) {
	// Check if user already exists
	_, status := us.UserRepository.GetUserByEmail(newUser.Email)
	if status == response.SuccessfulSearch {
		return user.User{}, response.Conflict // User already exists
	}

	// Hash the password before storing
	hashedPassword, err := us.Encoder.HashAndSalt(newUser.Password)
	if err != nil {
		return user.User{}, response.InternalServerError
	}
	newUser.Password = string(hashedPassword)

	// Save the new user in the repository
	newUser.RefreshTokens = nil
	createdUser, status := us.UserRepository.CreateUser(&newUser)

	if status != response.SuccessfulCreation {
		return user.User{}, status
	}

	newUser.Password = "" // Do not return the password
	return *createdUser, response.SuccessfulCreation
}
