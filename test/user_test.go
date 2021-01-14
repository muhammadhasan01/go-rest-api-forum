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
		Input:              "mhasan01",
		ExpectedStatusCode: 200,
	})
	// A false test case
	testCase = append(testCase, interfaces.TestStruct{
		Input:              "NotMhasan01",
		ExpectedStatusCode: 400,
	})

	// Check every testcase
	for _, tc := range testCase {
		fmt.Println(tc.Input)
		// Make a request to the /user/{username}
		req, err := http.NewRequest("GET", fmt.Sprintf("/user/%v", tc.Input), bytes.NewBufferString(tc.Input))
		if err != nil {
			t.Errorf("Error trying to get new request get to /user/%v: %v", tc.Input, err)
		}
		req = mux.SetURLVars(req, map[string]string{"username": tc.Input})
		// Create http test
		r := httptest.NewRecorder()
		h := http.HandlerFunc(user.GetUserHandler)
		h.ServeHTTP(r, req)
		// Assert the result with the expected result
		assert.Equal(t, r.Code, tc.ExpectedStatusCode)
		fmt.Println()
	}
}
