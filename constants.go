package constants

import (
	"os"
	"time"
)

var AccessTokenSecret = []byte(os.Getenv("JWT_SESSION_SECRET"))
var RefreshTokenSecret = []byte(os.Getenv("JWT_REFRESH_SECRET"))

// 5 minutes
const AccessTokenExpireTime = time.Minute * 5

// 7 days
const RefreshTokenExpireTime = time.Hour * 24 * 7
