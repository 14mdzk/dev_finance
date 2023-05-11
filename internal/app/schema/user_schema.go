package schema

type UserResp struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	FullName     string `json:"full_name"`
	RegisteredAt string `json:"registered_at"`
}

type UserChangePasswordReq struct {
	Password             string `json:"password" validate:"required,min=8"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required,min=8"`
}

type UserRegisterReq struct {
	Username             string `json:"username" validate:"required"`
	FullName             string `json:"full_name" validate:"required"`
	Password             string `json:"password" validate:"required"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required"`
}
