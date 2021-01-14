package post

// PostBody is the type that needs
// to be in a body request when user wants to add a post
type PostBody struct {
	Title       string `json:"title" example:"Title of a post"`
	Description string `json:"description" example:"Description of the post"`
}

// PostResponse is the type that will
// be responsed to a post GET request
type PostResponse struct {
	ID          uint   `json:"id" example:"45"`
	Username    string `json:"username" example:"tester"`
	Title       string `json:"title" example:"Title of a post"`
	Description string `json:"description" example:"Description of the post"`
}

// AddPostResponse is the type that will
// be responded to a add post request
type AddPostResponse struct {
	Message  string `json:"message" example:"post has been added successfully!"`
	ID       uint   `json:"id" example:"45"`
	Title    string `json:"title" example:"Title of a post"`
	Username string `json:"username" example:"tester"`
}

// UpdatePostResponse is the type that will
// be responded to an update post request
type UpdatePostResponse struct {
	Message     string `json:"message" example:"post has been updated successfully!"`
	Username    string `json:"username" example:"tester"`
	Title       string `json:"title" example:"new post name"`
	Description string `json:"description" example:"new description post"`
}

// DeletePostResponse is the type that will
// be responded to an delete delete request
type DeletePostResponse struct {
	Message  string `json:"message" example:"post has been deleted successfully!"`
	ID       uint   `json:"id" example:"45"`
	Username string `json:"username" example:"tester"`
}

// ErrorResponse is the response
// to give when an error occurs
type ErrorResponse struct {
	Msg string `json:"error" example:"something wrong happened"`
}
