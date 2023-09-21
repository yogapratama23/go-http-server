package auth

type SignupInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SigninInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignoutInput struct {
	ID    int    `json:"id"`
	Token string `json:"token"`
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

type ContextKey string

type UserInfo struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}
