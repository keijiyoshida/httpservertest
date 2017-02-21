package httpservertest

// ExpectedResponse represents an expected response of a test case.
type ExpectedResponse struct {
	// StatusCode is an expected status code of the response.
	StatusCode *int `yaml:"StatusCode"`
}
