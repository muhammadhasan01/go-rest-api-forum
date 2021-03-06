// interfaces is a package for global defined struct
package interfaces

import "github.com/jinzhu/gorm"

// User struct defines user schema in the database
type User struct {
	gorm.Model

	Username string
	Email    string
	Password string
	Role     string
}

// Auth struct is used for the 'whitelist' database
// it is used to handle logout
type Auth struct {
	gorm.Model

	UserID uint
	Token  string
}

// Thread struct defines thread schema in the database
type Thread struct {
	gorm.Model

	UserID      uint
	Username    string
	Name        string
	Description string
}

// Post struct defines post schema in the database
type Post struct {
	gorm.Model

	UserID      uint
	Username    string
	ThreadID    uint
	Title       string
	Description string
}

// TestStruct struct is used for testing http request
// with the data given in the body
type TestStruct struct {
	Input              string
	ExpectedStatusCode int
	Username           string
	ThreadID           string
	PostID             string
}
