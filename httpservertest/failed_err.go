package httpservertest

import (
	"fmt"
	"time"
)

// failedErr represents a test failed error.
type FailedErr struct {
	reason  string
	elapsed time.Duration
}

func (err *FailedErr) Error() string {
	return fmt.Sprintf("result: failed\telapsed: %s\treason: %s", err.elapsed, err.reason)
}

// newFailedErr creates a test failed error and returns it.
func newFailedErr(reason string, elapsed time.Duration) *FailedErr {
	return &FailedErr{
		reason:  reason,
		elapsed: elapsed,
	}
}
