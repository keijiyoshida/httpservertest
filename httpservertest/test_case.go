package httpservertest

// TestCase represents a test case.
type TestCase struct {
	// Name is a name of the test case.
	Name string `yaml:"Name"`
	// Request is a request of the test case.
	Request Request `yaml:"Request"`
	// ExpectedResponse is an expected response of the test case.
	ExpectedResponse *ExpectedResponse `yaml:"ExpectedResponse"`
	// ExpectedFileOutputs is expected file outputs of the test case.
	ExpectedFileOutputs *[]ExpectedFileOutput `yaml:"ExpectedFileOutputs"`
}
