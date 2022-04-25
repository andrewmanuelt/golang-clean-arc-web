package web

type LoginRequest struct {
	Email string
}

type LoginResponse struct {
	Username string
	Email    string
	Password string
}

type RegisterRequest struct {
	Username string
	Email    string
	Password string
}
