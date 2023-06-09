package schema

type LoginReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenReq struct {
	UserID       int
	RefreshToken string
}

type RefreshTokenResp struct {
	AccessToken string `json:"access_token"`
}
