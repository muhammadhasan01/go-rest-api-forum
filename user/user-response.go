package user

// LoginBody is the type that needs
// to be in a body request when user wants to login
type LoginBody struct {
	Username string `json:"username" example:"tester"`
	Password string `json:"password" example:"some_password"`
}

// RegisterBody is the type that needs
// to be in a body request when user wants to register
type RegisterBody struct {
	Username string `json:"username" example:"tester"`
	Email    string `json:"email" example:"tester@gmail.com"`
	Password string `json:"password" example:"some_password"`
}

// UserResponse is the type that will
// be responsed to a user GET request
type UserResponse struct {
	UserID   uint   `json:"user_id" example:"205"`
	Username string `json:"username" example:"tester"`
	Email    string `json:"email" example:"tester@gmail.com"`
}

// RegisterResponse is the type that will
// be responded to a register request
type RegisterResponse struct {
	Message  string `json:"message" example:"user registered successfully!"`
	Username string `json:"username" example:"tester"`
	Email    string `json:"email" example:"tester@gmail.com"`
}

// LogoutResponse is the type that will
// be responded to a logout request
type LogoutResponse struct {
	Message  string `json:"message" example:"you have been logged out successfully"`
	Username string `json:"username" example:"tester"`
}

// ErrorResponse is the response
// to give when an error occurs
type ErrorResponse struct {
	Msg string `json:"error" example:"something wrong happened"`
}
