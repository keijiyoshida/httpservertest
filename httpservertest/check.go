package httpservertest

import (
	"net/http"
	"time"
)

func check(testCase TestCase, res *http.Response, resBody string, requestTime time.Time, elapsed time.Duration) error {
	if testCase.ExpectedResponse != nil {
		if err := checkResponse(testCase.ExpectedResponse, res, resBody, requestTime, elapsed); err != nil {
			return err
		}
	}

	if testCase.ExpectedFileOutputs != nil {
		for _, expectedFileOutput := range *testCase.ExpectedFileOutputs {
			if err := checkFileOutput(expectedFileOutput, elapsed); err != nil {
				return err
			}
		}
	}
	return nil
}
