package httpservertest

// Config is configuration for the HTTP server testing process.
type Config struct {
	// BaseURL is the endpoint URL without a path. (e.g. http://yourdomain.com)
	BaseURL string
	// TestCasesFilePath is the test cases file path.
	TestCasesFilePath string
	// ParametersFilePath is the parameters file path. This field is optional.
	ParametersFilePath string
}
