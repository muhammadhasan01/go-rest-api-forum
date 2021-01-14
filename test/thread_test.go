package test

import (
	"backend-forum/interfaces"
	"backend-forum/thread"
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	assert "gopkg.in/go-playground/assert.v1"
)

// TestGetThread is to test getting a thread by a threadID
// it tests the endpoint GET /thread/{threadID}
func TestGetThread(t *testing.T) {
	// Create a slice of testcase struct
	testCase := make([]interfaces.TestStruct, 0)
	// A correct test case
	testCase = append(testCase, interfaces.TestStruct{
		ThreadID:           "2",
		ExpectedStatusCode: 200,
	})
	// A false test case [id not found]
	testCase = append(testCase, interfaces.TestStruct{
		ThreadID:           "0",
		ExpectedStatusCode: 400,
	})
	// A false test case [id not an integer]
	testCase = append(testCase, interfaces.TestStruct{
		ThreadID:           "notAnInteger",
		ExpectedStatusCode: 400,
	})

	// Check every testcase
	for _, tc := range testCase {
		fmt.Println(tc.ThreadID)
		// Make a request to the /user/{username}
		req, err := http.NewRequest("GET", fmt.Sprintf("/thread/%v", tc.ThreadID), bytes.NewBufferString(""))
		if err != nil {
			t.Errorf("Error trying to get new request get to /thread/%v: %v", tc.ThreadID, err)
		}
		// Sets the path threadID
		req = mux.SetURLVars(req, map[string]string{"threadID": tc.ThreadID})
		// Create http test
		r := httptest.NewRecorder()
		h := http.HandlerFunc(thread.GetThreadHandler)
		h.ServeHTTP(r, req)
		// Assert the result with the expected result
		assert.Equal(t, r.Code, tc.ExpectedStatusCode)
		fmt.Println()
	}
}

// TestAddThread is to test getting a thread by a threadID
// it tests the endpoint POST /thread/add
func TestAddThread(t *testing.T) {
	// Create a slice of testcase struct
	testCase := make([]interfaces.TestStruct, 0)
	// A correct test case
	testCase = append(testCase, interfaces.TestStruct{
		Input:              `{"name":"threadName", "description":"descriptionThread"}`,
		ExpectedStatusCode: 200,
	})

	// Check every testcase
	for _, tc := range testCase {
		fmt.Println(tc.Input)
		// Make a request to the /thread/add
		req, err := http.NewRequest("GET", "/thread/add", bytes.NewBufferString(tc.Input))
		if err != nil {
			t.Errorf("Error trying to get new request POST to /thread/add")
		}
		// Sets the token for the header
		req.Header.Set("Token", Token)
		// Create http test
		r := httptest.NewRecorder()
		h := http.HandlerFunc(thread.AddThreadHandler)
		h.ServeHTTP(r, req)
		// Assert the result with the expected result
		assert.Equal(t, r.Code, tc.ExpectedStatusCode)
		fmt.Println()
	}
}

// TestUpdateThread is to test updating a thread by a threadID
// it tests the endpoint PUT /thread/{threadID}
func TestUpdateThread(t *testing.T) {
	// Create a slice of testcase struct
	testCase := make([]interfaces.TestStruct, 0)
	// A correct test case
	testCase = append(testCase, interfaces.TestStruct{
		ThreadID:           "11",
		Input:              `{"name":"threadName", "description":"descriptionThread"}`,
		ExpectedStatusCode: 200,
	})
	// A false test case [id not found]
	testCase = append(testCase, interfaces.TestStruct{
		ThreadID:           "0",
		Input:              `{"name":"threadName", "description":"descriptionThread"}`,
		ExpectedStatusCode: 400,
	})
	// A false test case [id not an integer]
	testCase = append(testCase, interfaces.TestStruct{
		ThreadID:           "notAnInteger",
		Input:              `{"name":"threadName", "description":"descriptionThread"}`,
		ExpectedStatusCode: 400,
	})
	// A false test case [cannot change other person thread]
	testCase = append(testCase, interfaces.TestStruct{
		ThreadID:           "2",
		Input:              `{"name":"threadName", "description":"descriptionThread"}`,
		ExpectedStatusCode: 400,
	})

	// Check every testcase
	for _, tc := range testCase {
		fmt.Println(tc.ThreadID, tc.Input)
		// Make a request to the /user/{username}
		req, err := http.NewRequest("GET", fmt.Sprintf("/thread/%v", tc.ThreadID), bytes.NewBufferString(tc.Input))
		if err != nil {
			t.Errorf("Error trying to get new request get to /thread/%v: %v", tc.ThreadID, err)
		}
		// Sets the path threadID
		req = mux.SetURLVars(req, map[string]string{"threadID": tc.ThreadID})
		// Sets the token for the header
		req.Header.Set("Token", Token)
		// Create http test
		r := httptest.NewRecorder()
		h := http.HandlerFunc(thread.UpdateThreadHandler)
		h.ServeHTTP(r, req)
		// Assert the result with the expected result
		assert.Equal(t, r.Code, tc.ExpectedStatusCode)
		fmt.Println()
	}
}

// TestDeleteThread is to test deleting a thread by a threadID
// it tests the endpoint DELETE /thread/{threadID}
func TestDeleteThread(t *testing.T) {
	// Create a slice of testcase struct
	testCase := make([]interfaces.TestStruct, 0)
	// A correct test case
	testCase = append(testCase, interfaces.TestStruct{
		ThreadID:           "11",
		ExpectedStatusCode: 200,
	})
	// A false test case [id not found]
	testCase = append(testCase, interfaces.TestStruct{
		ThreadID:           "0",
		ExpectedStatusCode: 400,
	})
	// A false test case [id not an integer]
	testCase = append(testCase, interfaces.TestStruct{
		ThreadID:           "notAnInteger",
		ExpectedStatusCode: 400,
	})
	// A false test case [cannot delete other person thread]
	testCase = append(testCase, interfaces.TestStruct{
		ThreadID:           "2",
		ExpectedStatusCode: 400,
	})

	// Check every testcase
	for _, tc := range testCase {
		fmt.Println(tc.ThreadID)
		// Make a request to the /user/{username}
		req, err := http.NewRequest("DELETE", fmt.Sprintf("/thread/%v", tc.ThreadID), bytes.NewBufferString(""))
		if err != nil {
			t.Errorf("Error trying to get new request get to /thread/%v: %v", tc.ThreadID, err)
		}
		// Sets the path threadID
		req = mux.SetURLVars(req, map[string]string{"threadID": tc.ThreadID})
		// Sets the token for the header
		req.Header.Set("Token", Token)
		// Create http test
		r := httptest.NewRecorder()
		h := http.HandlerFunc(thread.DeleteThreadHandler)
		h.ServeHTTP(r, req)
		// Assert the result with the expected result
		assert.Equal(t, r.Code, tc.ExpectedStatusCode)
		fmt.Println()
	}
}
