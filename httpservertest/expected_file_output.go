package httpservertest

// ExpectedFileOutput represents an expected file output of a test case.
type ExpectedFileOutput struct {
	// Path is a file path which the expected output is written to.
	Path string `yaml:"Path"`
	// NunLines is the number of lines which tail reads from the bottom of the file.
	// Default value is 1.
	NumLines *int `yaml:"NumLines"`
	// Pattern is a textual regular expression of the expected output.
	Pattern *string `yaml:"Pattern"`
}
