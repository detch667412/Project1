package apitest

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

type ApiResponse struct {
	Origin string `json:"origin"`
	URL    string `json:"url"`
}

func FetchData(apiURL string) (ApiResponse, error) {
	// Create a new fasthttp request
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	// Set the request URL
	req.SetRequestURI(apiURL)

	// Create a new fasthttp response
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	// Make the HTTP GET request
	if err := fasthttp.Do(req, resp); err != nil {
		return ApiResponse{}, fmt.Errorf("error making the request: %s", err)
	}

	// Check the response status code
	if resp.StatusCode() != fasthttp.StatusOK {
		return ApiResponse{}, fmt.Errorf("non-OK status code: %d", resp.StatusCode())
	}

	// Get the response body as a string
	responseBody := resp.Body()

	fmt.Println(string(responseBody))
	// Parse the JSON response
	var apiResponse ApiResponse
	err := json.Unmarshal(responseBody, &apiResponse)
	if err != nil {
		return ApiResponse{}, fmt.Errorf("error parsing JSON: %s", err)
	}

	return apiResponse, nil
}
