package refreshToken

type RefreshToken struct {
	Id        int    `json:"id"`
	Token     string `json:"token"`
	UserId    int    `json:"userId"`
	CreatedAt string `json:"createdAt"`
}
