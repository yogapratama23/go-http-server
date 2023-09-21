package auth

type SignupInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SigninInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token  string `json:"token"`
	UserId int    `json:"user_id"`
}
