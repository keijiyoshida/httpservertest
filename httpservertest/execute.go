package httpservertest

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// execute executes a test case.
func execute(baseURL string, testCase TestCase, client httpClient) error {
	log.Printf("testCase: %s\n", testCase.Name)

	req, err := http.NewRequest(testCase.Request.Method, baseURL+testCase.Request.Path, nil)
	if err != nil {
		return err
	}

	requestTime := time.Now()

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	responseTime := time.Now()
	elapsed := responseTime.Sub(requestTime)

	resBody, err := readResponseBody(res)
	if err != nil {
		return err
	}

	if err := check(testCase, res, resBody, requestTime, elapsed); err != nil {
		return err
	}

	log.Printf("result: ok\trequestTime: %s\tresponseTime: %s\telapsed: %s\n", requestTime, responseTime, elapsed)

	return nil
}

// readResponseBody reads the response body.
func readResponseBody(res *http.Response) (string, error) {
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
