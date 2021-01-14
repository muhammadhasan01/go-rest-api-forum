package auth

// LoginBody is the type that needs
// to be in a body request when user wants to login
type LoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// RegisterBody is the type that needs
// to be in a body request when user wants to register
type RegisterBody struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse is the type that will
// be responsed to a login request
type LoginResponse struct {
	Message string `json:"message" example:"you have been logged in successfully!"`
	Token   string `json:"token" example="someJwtToken"`
}

// ErrorResponse is the response
// to give when an error occurs
type ErrorResponse struct {
	Msg string `json:"error" example:"username has already been taken" description:"error message"`
}
