package httpservertest

// ExpectedResponse represents an expected response of a test case.
type ExpectedResponse struct {
	// StatusCode is an expected status code of the response.
	StatusCode *int `yaml:"StatusCode"`
	// ContentType is an expected content type of the response.
	ContentType *string `yaml:"ContentType"`
	// ContentLength is an expected content length of the response.
	ContentLength *int64 `yaml:"ContentLength"`
	// Close is an expected connection status of the response.
	Close *bool `yaml:"Close"`
	// Location is an expected location of the response.
	Location *string `yaml:"Location"`
	// Cookies is expected Cooikes of the response.
	Cookies *[]Cookie `yaml:"Cookies"`
	// Header is an expected Header of the response.
	Header *[]HeaderField `yaml:"Header"`
	// Body is an expected body of the response.
	Body *string `yaml:"Body"`
}
