package httpservertest

import (
	"crypto/tls"
	"net/http"
)

// httpClient is an HTTP client.
type httpClient interface {
	Do(r *http.Request) (*http.Response, error)
}

// newHTTPClient creates an HTTP client and returns it.
func newHTTPClient(insecureSkipVerify bool) httpClient {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: insecureSkipVerify,
			},
		},
		CheckRedirect: doNotFollowRedirect,
	}
}

// doNotFollowRedirect returns http.ErrUseLastResponse.
// This function is supposed to be used for http.Client.CheckRedirect
// to prevent the HTTP client from following a redirect.
func doNotFollowRedirect(_ *http.Request, _ []*http.Request) error {
	return http.ErrUseLastResponse
}
