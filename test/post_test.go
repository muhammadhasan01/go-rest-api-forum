package test

import (
	"backend-forum/interfaces"
	"backend-forum/post"
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	assert "gopkg.in/go-playground/assert.v1"
)

// TestGetPost is to test getting a post by a threadID and a postID
// it tests the endpoint GET /thread/{threadID}/post/{postID}
func TestGetPost(t *testing.T) {
	// Create a slice of testcase struct
	testCase := make([]interfaces.TestStruct, 0)
	// A correct test case
	testCase = append(testCase, interfaces.TestStruct{
		ThreadID:           "7",
		PostID:             "7",
		ExpectedStatusCode: 200,
	})
	// A false test case [thread ID not found]
	testCase = append(testCase, interfaces.TestStruct{
		ThreadID:           "0",
		PostID:             "8",
		ExpectedStatusCode: 400,
	})
	// A false test case [post ID not found]
	testCase = append(testCase, interfaces.TestStruct{
		ThreadID:           "7",
		PostID:             "0",
		ExpectedStatusCode: 400,
	})
	// A false test case [thread ID not an integer]
	testCase = append(testCase, interfaces.TestStruct{
		ThreadID:           "notInteger",
		PostID:             "7",
		ExpectedStatusCode: 400,
	})
	// A false test case [post ID not an integer]
	testCase = append(testCase, interfaces.TestStruct{
		ThreadID:           "7",
		PostID:             "notInteger",
		ExpectedStatusCode: 400,
	})

	// Check every testcase
	for _, tc := range testCase {
		fmt.Println(tc.ThreadID, tc.PostID)
		// Make a request to the /thread/{threadID}/post/{postID}
		req, err := http.NewRequest("GET", fmt.Sprintf("/thread/%v/post/%v", tc.ThreadID, tc.PostID), bytes.NewBufferString(""))
		if err != nil {
			t.Errorf("Error trying to get new request get to /thread/%v/post/%v: %v", tc.ThreadID, tc.PostID, err)
		}
		// Sets the path threadID and postID
		req = mux.SetURLVars(req, map[string]string{"threadID": tc.ThreadID, "postID": tc.PostID})
		// Create http test
		r := httptest.NewRecorder()
		h := http.HandlerFunc(post.GetPostHandler)
		h.ServeHTTP(r, req)
		// Assert the result with the expected result
		assert.Equal(t, r.Code, tc.ExpectedStatusCode)
		fmt.Println()
	}
}
