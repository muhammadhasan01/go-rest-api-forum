package thread

// UpdateBody is the type that needs
// to be in a body request when user wants to update the password
type UpdateBody struct {
	Password string `json:"password" example:"some_new_password"`
}

// ThreadResponse is the type that will
// be responsed to a thread GET request
type ThreadResponse struct {
	ID          uint   `json:"thread_id" example:"45"`
	Username    string `json:"username" example:"tester"`
	Name        string `json:"name" example:"Name of a thread"`
	Description string `json:"description" example:"Description of the thread"`
}

// LogoutResponse is the type that will
// be responded to a logout request
type LogoutResponse struct {
	Message  string `json:"message" example:"you have been logged out successfully"`
	Username string `json:"username" example:"tester"`
}

// UpdateResponse is the type that will
// be responded to an update request
type UpdateResponse struct {
	Message  string `json:"message" example:"user has been updated with the new password"`
	Username string `json:"username" example:"tester"`
}

// DeleteResponse is the type that will
// be responded to an delete request
type DeleteResponse struct {
	Message  string `json:"message" example:"user has been deleted successfully!"`
	Username string `json:"username" example:"tester"`
}

// ErrorResponse is the response
// to give when an error occurs
type ErrorResponse struct {
	Msg string `json:"error" example:"something wrong happened"`
}
