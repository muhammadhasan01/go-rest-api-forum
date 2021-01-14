package test

import (
	"backend-forum/interfaces"
	"backend-forum/user"
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	assert "gopkg.in/go-playground/assert.v1"
)

// TestGetUser is to test getting a user by a username
// it tests the endpoint /user/{username}
func TestGetUser(t *testing.T) {
	// Create a slice of testcase struct
	testCase := make([]interfaces.TestStruct, 0)
	// A correct test case
	testCase = append(testCase, interfaces.TestStruct{
		Username:           "mhasan01",
		ExpectedStatusCode: 200,
	})
	// A false test case
	testCase = append(testCase, interfaces.TestStruct{
		Username:           "NotMhasan01",
		ExpectedStatusCode: 400,
	})

	// Check every testcase
	for _, tc := range testCase {
		fmt.Println(tc.Username)
		// Make a request to the /user/{username}
		req, err := http.NewRequest("GET", fmt.Sprintf("/user/%v", tc.Username), bytes.NewBufferString(""))
		if err != nil {
			t.Errorf("Error trying to get new request get to /user/%v: %v", tc.Username, err)
		}
		// Sets the path username
		req = mux.SetURLVars(req, map[string]string{"username": tc.Username})
		// Create http test
		r := httptest.NewRecorder()
		h := http.HandlerFunc(user.GetUserHandler)
		h.ServeHTTP(r, req)
		// Assert the result with the expected result
		assert.Equal(t, r.Code, tc.ExpectedStatusCode)
		fmt.Println()
	}
}

// TestUpdateUser is to test updating a user by its username
// it tests the endpoint /user/{username}
func TestUpdateUser(t *testing.T) {
	// Create a slice of testcase struct
	testCase := make([]interfaces.TestStruct, 0)
	// A correct test case
	testCase = append(testCase, interfaces.TestStruct{
		Username:           "tester",
		Input:              `"password":"tester"`,
		ExpectedStatusCode: 200,
	})
	// A false test case [cannot change other person user]
	testCase = append(testCase, interfaces.TestStruct{
		Username:           "mhasan01",
		Input:              `"password":"tester"`,
		ExpectedStatusCode: 400,
	})

	// Check every testcase
	for _, tc := range testCase {
		fmt.Println(tc.Username)
		// Make a request to the /user/{username}
		req, err := http.NewRequest("PUT", fmt.Sprintf("/user/%v", tc.Username), bytes.NewBufferString(tc.Input))
		if err != nil {
			t.Errorf("Error trying to get new request PUT to /user/%v: %v", tc.Username, err)
		}
		// Sets the path username
		req = mux.SetURLVars(req, map[string]string{"username": tc.Username})
		// Sets the token for the header
		req.Header.Set("Token", Token)
		// Create http test
		r := httptest.NewRecorder()
		h := http.HandlerFunc(user.UpdateUserHandler)
		h.ServeHTTP(r, req)
		// Assert the result with the expected result
		assert.Equal(t, r.Code, tc.ExpectedStatusCode)
		fmt.Println()
	}
}
