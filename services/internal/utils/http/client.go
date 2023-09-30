package request

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type HTTP interface {
	GET(url string) (int, []byte, error)
	POST(url string, body string) (int, []byte, error)
}

type HTTPClient struct {
    httpClient *http.Client
    baseURL    string
}

type MockHTTPClient struct {
	Responses map[string]MockResponse
}

type MockResponse struct {
	Response []byte
	Code     int
	Err      error
}

func (c HTTPClient) NewHTTPClient(baseURL string) *HTTPClient {
    return &HTTPClient{
        httpClient: &http.Client{},
        baseURL: baseURL,
    }
}

func (c HTTPClient) GET(path string) (int, []byte, error) {
	// Create a request
    url := c.baseURL + path
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, nil, err
	}

	// Send the request and get the response
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return resp.StatusCode, nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, err
	}

	return resp.StatusCode, respBody, err
}
func (c HTTPClient) POST(path string, body string) (int, []byte, error) {
	// Create a request
    url := c.baseURL + path
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(body))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		return 0, nil, err
	}

	// Send the request and get the response
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return resp.StatusCode, nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, err
	}

	return resp.StatusCode, respBody, err
}

func (c MockHTTPClient ) NewMockHTTPClient() *MockHTTPClient {
	return &MockHTTPClient{
		Responses: make(map[string]MockResponse),
	}
}
func (c MockHTTPClient) POST(url string, body string) (int, []byte, error) {
	if response, ok := c.Responses[url]; ok {
		if response.Response != nil {
			return response.Code, response.Response, response.Err
		}
		// If the response is nil, return an empty array
		return response.Code, []byte{}, response.Err
	}
	return 0, nil, errors.New("URL not found in mock responses")
}

func (c MockHTTPClient) GET(url string) (int, []byte, error) {
	if response, ok := c.Responses[url]; ok {
		if response.Response != nil {
			return response.Code, response.Response, response.Err
		}
		// If the response is nil, return an empty array
		return response.Code, []byte{}, response.Err
	}
	return 0, nil, errors.New("URL not found in mock responses")
}