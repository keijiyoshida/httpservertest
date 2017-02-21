package httpservertest

import (
	"net/http"
	"time"
)

func check(testCase TestCase, res *http.Response, requestTime time.Time, elapsed time.Duration) error {
	if testCase.ExpectedResponse != nil {
		if err := checkResponse(testCase.ExpectedResponse, res, requestTime, elapsed); err != nil {
			return err
		}
	}

	return nil
}
