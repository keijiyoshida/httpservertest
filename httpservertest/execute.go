package httpservertest

import (
	"log"
	"net/http"
	"time"
)

// execute executes a test case.
func execute(baseURL string, testCase TestCase, client httpClient) error {
	log.Printf("testCase: %s\n", testCase.Name)

	startTime := time.Now()

	req, err := http.NewRequest(testCase.Request.Method, baseURL+testCase.Request.Path, nil)
	if err != nil {
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	elapsed := time.Now().Sub(startTime)

	if err := check(testCase, res, elapsed); err != nil {
		return err
	}

	log.Printf("result: ok\telapsed: %s\n", elapsed)

	return nil
}
