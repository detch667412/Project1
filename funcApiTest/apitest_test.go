package apitest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchData(t *testing.T) {
	t.Run("Successful request", func(t *testing.T) {
		// Create a test server
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Respond with a successful JSON response
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `{"args":{},"origin":"xxx.xxx.xxx","url":"https://www.testfunction.com"}`)
		}))
		defer ts.Close()

		// Set the API URL to the test server URL
		apiURL := ts.URL

		// Call the fetchData function with the test server URL
		response, err := FetchData(apiURL)
		fmt.Println()
		// Assert the expected values in the response
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.NotEmpty(t, response.Origin)
		assert.NotEmpty(t, response.URL)
	})

	t.Run("Error status code", func(t *testing.T) {
		// Create a test server that returns an error status code
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer ts.Close()

		// Set the API URL to the test server URL
		apiURL := ts.URL

		// Call the fetchData function with the test server URL
		_, err := FetchData(apiURL)

		// Assert that an error occurred
		assert.Error(t, err)
	})
}
