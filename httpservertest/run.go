package httpservertest

import "fmt"

// Run runs HTTP server testing.
func Run(conf Config) error {
	if conf.Version {
		fmt.Printf("HTTPServerTest %s\n", version)
		return nil
	}

	params, err := newParameters(conf.ParametersFilePath)
	if err != nil {
		return err
	}

	testCases, err := newTestCases(conf.TestCasesFilePath, params)
	if err != nil {
		return err
	}

	client := newHTTPClient(conf.InsecureSkipVerify)

	for _, testCase := range testCases {
		if err := execute(conf.BaseURL, testCase, client); err != nil {
			return err
		}
	}

	return nil
}
