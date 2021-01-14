package test

import (
	"backend-forum/auth"
	"backend-forum/interfaces"
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	assert "gopkg.in/go-playground/assert.v1"
)

// TestRegister is used to test registering a user
// it tests the endpoint /auth/register
func TestRegister(t *testing.T) {
	// Create a slice of testcase struct
	testCase := make([]interfaces.TestBody, 0)
	// A correct test case
	testCase = append(testCase, interfaces.TestBody{
		InputBodyJSON:      `{"username":"NewTester", "email":"NewTester@gmail.com", "password":"tester"}`,
		ExpectedStatusCode: 200,
	})
	// A false test case [username has already been taken]
	testCase = append(testCase, interfaces.TestBody{
		InputBodyJSON:      `{"username":"tester", "email":"NewTester@gmail.com", "password":"tester"}`,
		ExpectedStatusCode: 400,
	})
	// A false test case [email has already been taken]
	testCase = append(testCase, interfaces.TestBody{
		InputBodyJSON:      `{"username":"VeryNewTester", "email":"tester@gmail.com", "password":"tester"}`,
		ExpectedStatusCode: 400,
	})

	// Check every testcase
	for _, tc := range testCase {
		fmt.Println(tc.InputBodyJSON)
		// Make a request to the /auth/register
		req, err := http.NewRequest("POST", "/auth/register", bytes.NewBufferString(tc.InputBodyJSON))
		if err != nil {
			t.Errorf("Error trying to get new request post to /auth/register: %v", err)
		}
		// Create http test
		r := httptest.NewRecorder()
		h := http.HandlerFunc(auth.RegisterHandler)
		h.ServeHTTP(r, req)
		// Assert the result with the expected result
		assert.Equal(t, r.Code, tc.ExpectedStatusCode)
		fmt.Println()
	}
}

// TestLogin is used to to test the login
// endpoint (/auth/login)
func TestLogin(t *testing.T) {
	// Create a slice of testcase struct
	testCase := make([]interfaces.TestBody, 0)
	// A correct test case
	testCase = append(testCase, interfaces.TestBody{
		InputBodyJSON:      `{"username":"tester", "password":"tester"}`,
		ExpectedStatusCode: 202,
	})
	// A false test case [username not found]
	testCase = append(testCase, interfaces.TestBody{
		InputBodyJSON:      `{"username":"notTester", "password":"tester"}`,
		ExpectedStatusCode: 400,
	})
	// A false test case [password is wrong]
	testCase = append(testCase, interfaces.TestBody{
		InputBodyJSON:      `{"username":"tester", "password":"notPasswordTester"}`,
		ExpectedStatusCode: 400,
	})
	// A false test case [user cannot login twice]
	testCase = append(testCase, interfaces.TestBody{
		InputBodyJSON:      `{"username":"tester", "password":"tester"}`,
		ExpectedStatusCode: 400,
	})

	// Check every testcase
	for _, tc := range testCase {
		fmt.Println(tc.InputBodyJSON)
		// Make a request to the /auth/login
		req, err := http.NewRequest("POST", "/auth/login", bytes.NewBufferString(tc.InputBodyJSON))
		if err != nil {
			t.Errorf("Error trying to get new request post to /auth/login: %v", err)
		}
		// Create http test
		r := httptest.NewRecorder()
		h := http.HandlerFunc(auth.LoginHandler)
		h.ServeHTTP(r, req)
		// Assert the result with the expected result
		assert.Equal(t, r.Code, tc.ExpectedStatusCode)
		fmt.Println()
	}
}

// TestLogout is used to to test the logout
// endpoint (/auth/logout)
func TestLogout(t *testing.T) {
	// Create a slice of testcase struct
	testCase := make([]interfaces.TestBody, 0)
	// A correct test case
	testCase = append(testCase, interfaces.TestBody{
		InputBodyJSON:      `{"username":"tester", "password":"tester"}`,
		ExpectedStatusCode: 200,
	})

	// Check every testcase
	for _, tc := range testCase {
		fmt.Println(tc.InputBodyJSON)
		// Make a request to the /auth/login
		req, err := http.NewRequest("POST", "/auth/logout", bytes.NewBufferString(tc.InputBodyJSON))
		if err != nil {
			t.Errorf("Error trying to get new request post to /auth/logout: %v", err)
		}
		// Sets the token for the header
		req.Header.Set("Token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyMDUsInVzZXJuYW1lIjoidGVzdGVyIn0.Qf1IxctVIPxT2EL2XwoMn1FKShT0-yq3Vxj6sj8wXTI")
		// Create http test
		r := httptest.NewRecorder()
		h := http.HandlerFunc(auth.LogoutHandler)
		h.ServeHTTP(r, req)
		// Assert the result with the expected result
		assert.Equal(t, r.Code, tc.ExpectedStatusCode)
		fmt.Println()
	}
}
