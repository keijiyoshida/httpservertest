package main

import (
	"flag"
	"log"

	"github.com/keijiyoshida/httpservertest/httpservertest"
)

func main() {
	if err := httpservertest.Run(newConfig()); err != nil {
		switch err.(type) {
		case *httpservertest.FailedErr:
			log.Fatal(err)
		default:
			log.Fatal("[error] " + err.Error())
		}
	}
}

// newConfig parses the command-line flags,
// creates httpservertest.Config and returns it.
func newConfig() httpservertest.Config {
	baseURL := flag.String("u", "", "endpoint URL without a path (e.g. http://yourdomain.com)")
	testCasesFilePath := flag.String("t", "", "test cases file path")
	parametersFilePath := flag.String("p", "", "parameters file path (optional)")
	insecureSkipVerify := flag.Bool("insecure-skip-verify", false, "controls whether a client verifies the server's certificate chain and host name")
	version := flag.Bool("v", false, "show the version of HTTPServerTest")

	flag.Parse()

	return httpservertest.Config{
		BaseURL:            *baseURL,
		TestCasesFilePath:  *testCasesFilePath,
		ParametersFilePath: *parametersFilePath,
		InsecureSkipVerify: *insecureSkipVerify,
		Version:            *version,
	}
}
