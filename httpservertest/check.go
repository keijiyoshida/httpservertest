package httpservertest

import (
	"net/http"
	"time"
)

func check(testCase TestCase, res *http.Response, elapsed time.Duration) error {
	if testCase.ExpectedResponse != nil {
		if err := checkResponse(testCase.ExpectedResponse, res, elapsed); err != nil {
			return err
		}
	}

	return nil
}
