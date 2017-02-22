package httpservertest

import (
	"fmt"
	"net/http"
	"time"
)

// checkResponse checks if the response is valid or not.
func checkResponse(expectedRes *ExpectedResponse, res *http.Response, resBody string, requestTime time.Time, elapsed time.Duration) error {
	if expectedRes.StatusCode != nil {
		if res.StatusCode != *expectedRes.StatusCode {
			return newFailedErr(
				fmt.Sprintf("StatusCode => %d, expected %d", res.StatusCode, *expectedRes.StatusCode),
				elapsed,
			)
		}
	}

	if expectedRes.ContentType != nil {
		if contentType := res.Header.Get("Content-Type"); contentType != *expectedRes.ContentType {
			return newFailedErr(
				fmt.Sprintf("ContentType => %s, expected %s", contentType, *expectedRes.ContentType),
				elapsed,
			)
		}
	}

	if expectedRes.ContentLength != nil {
		if res.ContentLength != *expectedRes.ContentLength {
			return newFailedErr(
				fmt.Sprintf("ContentLength => %d, expected %d", res.ContentLength, *expectedRes.ContentLength),
				elapsed,
			)
		}
	}

	if expectedRes.Close != nil {
		if res.Close != *expectedRes.Close {
			return newFailedErr(
				fmt.Sprintf("Close => %t, expected %t", res.Close, *expectedRes.Close),
				elapsed,
			)
		}
	}

	if expectedRes.Location != nil {
		url, err := res.Location()
		if err != nil {
			return err
		}

		if location := url.String(); location != *expectedRes.Location {
			return newFailedErr(
				fmt.Sprintf("Location => %s, expected %s", location, *expectedRes.Location),
				elapsed,
			)
		}
	}

	if expectedRes.Cookies != nil {
		cookies := res.Cookies()

		if len(cookies) != len(*expectedRes.Cookies) {
			return newFailedErr(
				fmt.Sprintf("len(Cookies) => %d, expected %d", len(cookies), len(*expectedRes.Cookies)),
				elapsed,
			)
		}

		for i, cookie := range cookies {
			expectedCookie := (*expectedRes.Cookies)[i]

			if expectedCookie.Name != nil {
				if cookie.Name != *expectedCookie.Name {
					return newFailedErr(
						fmt.Sprintf("Cookie.Name => %s, expected %s", cookie.Name, *expectedCookie.Name),
						elapsed,
					)
				}
			}

			if expectedCookie.Value != nil {
				if cookie.Value != *expectedCookie.Value {
					return newFailedErr(
						fmt.Sprintf("Cookie.Value => %s, expected %s", cookie.Value, *expectedCookie.Value),
						elapsed,
					)
				}
			}

			if expectedCookie.Path != nil {
				if cookie.Path != *expectedCookie.Path {
					return newFailedErr(
						fmt.Sprintf("Cookie.Path => %s, expected %s", cookie.Path, *expectedCookie.Path),
						elapsed,
					)
				}
			}

			if expectedCookie.Domain != nil {
				if cookie.Domain != *expectedCookie.Domain {
					return newFailedErr(
						fmt.Sprintf("Cookie.Domain => %s, expected %s", cookie.Domain, *expectedCookie.Domain),
						elapsed,
					)
				}
			}

			if expectedCookie.Secure != nil {
				if cookie.Secure != *expectedCookie.Secure {
					return newFailedErr(
						fmt.Sprintf("Cookie.Secure => %t, expected %t", cookie.Secure, *expectedCookie.Secure),
						elapsed,
					)
				}
			}

			if expectedCookie.Expires != nil {
				if expectedCookie.Expires.AfterOrEqualToRequestTimePlus != nil {
					minTime := requestTime.Add(time.Duration(*expectedCookie.Expires.AfterOrEqualToRequestTimePlus) * time.Second).Add(time.Duration(-requestTime.Nanosecond())).UTC()
					if cookie.Expires.Before(minTime) {
						return newFailedErr(
							fmt.Sprintf("Cookie.Expires (%s) is before %s", cookie.Expires, minTime),
							elapsed,
						)
					}
				}
			}
		}
	}

	if expectedRes.Body != nil {
		if resBody != *expectedRes.Body {
			return newFailedErr(
				fmt.Sprintf("Body => %s, expected %s", resBody, *expectedRes.Body),
				elapsed,
			)
		}
	}

	return nil
}
