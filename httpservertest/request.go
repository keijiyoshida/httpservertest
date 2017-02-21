package httpservertest

// Request represents a request of a test case.
type Request struct {
	// Method is a method of the request.
	Method string `yaml:"Method"`
	// Path is a path of the request.
	Path string `yaml:"Path"`
}
