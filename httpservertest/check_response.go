package httpservertest

import (
	"fmt"
	"net/http"
	"time"
)

// checkResponse checks if the response is valid or not.
func checkResponse(expectedRes *ExpectedResponse, res *http.Response, elapsed time.Duration) error {
	if expectedRes.StatusCode != nil {
		if res.StatusCode != *expectedRes.StatusCode {
			return newFailedErr(
				fmt.Sprintf("StatusCode => %d, expected %d", res.StatusCode, *expectedRes.StatusCode),
				elapsed,
			)
		}
	}

	return nil
}
