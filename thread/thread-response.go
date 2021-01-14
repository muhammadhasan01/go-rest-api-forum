package thread

// ThreadBody is the type that needs
// to be in a body request when user wants to add a thread
type ThreadBody struct {
	Name        string `json:"name" example:"Name of a thread"`
	Description string `json:"description" example:"Description of the thread"`
}

// ThreadResponse is the type that will
// be responsed to a thread GET request
type ThreadResponse struct {
	ID          uint   `json:"id" example:"45"`
	Username    string `json:"username" example:"tester"`
	Name        string `json:"name" example:"Name of a thread"`
	Description string `json:"description" example:"Description of the thread"`
}

// AddThreadResponse is the type that will
// be responded to a add thread request
type AddThreadResponse struct {
	Message  string `json:"message" example:"thread has been added successfully!"`
	ID       uint   `json:"id" example:"45"`
	Name     string `json:"name" example:"Name of a thread"`
	Username string `json:"username" example:"tester"`
}

// UpdateThreadResponse is the type that will
// be responded to an update thread request
type UpdateThreadResponse struct {
	Message     string `json:"message" example:"thread has been updated successfully!"`
	Username    string `json:"username" example:"tester"`
	Name        string `json:"name" example:"new thread name"`
	Description string `json:"description" example:"new description thread"`
}

// DeleteThreadResponse is the type that will
// be responded to an delete delete request
type DeleteThreadResponse struct {
	Message  string `json:"message" example:"thread has been deleted successfully!"`
	ID       uint   `json:"id" example:"45"`
	Username string `json:"username" example:"tester"`
}

// ErrorResponse is the response
// to give when an error occurs
type ErrorResponse struct {
	Msg string `json:"error" example:"something wrong happened"`
}
