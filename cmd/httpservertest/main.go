package main

import (
	"flag"
	"log"

	"github.com/keijiyoshida/httpservertest/httpservertest"
)

func main() {
	if err := httpservertest.Run(newConfig()); err != nil {
		log.Fatal(err)
	}
}

// newConfig parses the command-line flags,
// creates httpservertest.Config and returns it.
func newConfig() httpservertest.Config {
	baseURL := flag.String("u", "", "endpoint URL without a path (e.g. http://yourdomain.com)")
	testCasesFilePath := flag.String("t", "", "test cases file path")
	parametersFilePath := flag.String("p", "", "parameters file path (optional)")

	flag.Parse()

	return httpservertest.Config{
		BaseURL:            *baseURL,
		TestCasesFilePath:  *testCasesFilePath,
		ParametersFilePath: *parametersFilePath,
	}
}
