package test

import (
	"backend-forum/auth"
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	assert "gopkg.in/go-playground/assert.v1"
)

// TestLogin is used to to test the login
// endpoint (/auth/login)
// ! Make sure the user has never logged in yet
func TestLogin(t *testing.T) {
	// Create defined testLogin struct
	type testLogin struct {
		inputBodyJSON      string
		expectedStatusCode int
	}
	// Create a slice of testcase struct
	testCase := make([]testLogin, 0)
	// A correct test case
	testCase = append(testCase, testLogin{
		inputBodyJSON:      `{"username":"tester", "password":"tester"}`,
		expectedStatusCode: 202,
	})
	// A false username test case
	testCase = append(testCase, testLogin{
		inputBodyJSON:      `{"username":"notTester", "password":"tester"}`,
		expectedStatusCode: 400,
	})
	// A false password test case
	testCase = append(testCase, testLogin{
		inputBodyJSON:      `{"username":"tester", "password":"notPasswordTester"}`,
		expectedStatusCode: 400,
	})
	// User cannot login twice
	testCase = append(testCase, testLogin{
		inputBodyJSON:      `{"username":"tester", "password":"tester"}`,
		expectedStatusCode: 400,
	})

	// Check every testcase
	for _, tc := range testCase {
		fmt.Println(tc.inputBodyJSON)
		// Make a request to the /auth/login
		req, err := http.NewRequest("POST", "/auth/login", bytes.NewBufferString(tc.inputBodyJSON))
		if err != nil {
			t.Errorf("Error trying to get new request post to /auth/login: %v", err)
		}
		// Create a http test
		r := httptest.NewRecorder()
		h := http.HandlerFunc(auth.LoginHandler)
		h.ServeHTTP(r, req)
		// Assert the result with the expected result
		assert.Equal(t, r.Code, tc.expectedStatusCode)
	}
}
