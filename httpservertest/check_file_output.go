package httpservertest

import (
	"fmt"
	"regexp"
	"time"
)

//  Default number of lines which tail reads from the bottom of the file.
const defaultNumLines = 1

// checkFileOutput checks if the file outputs are valid or not.
func checkFileOutput(expectedFileOutput ExpectedFileOutput, elapsed time.Duration) error {
	numLines := defaultNumLines
	if expectedFileOutput.NumLines != nil {
		numLines = *expectedFileOutput.NumLines
	}

	fileOutput, err := tail(expectedFileOutput.Path, numLines)
	if err != nil {
		return err
	}

	if expectedFileOutput.Pattern != nil {
		matchResult, err := regexp.MatchString(*expectedFileOutput.Pattern, fileOutput)
		if err != nil {
			return err
		}

		if !matchResult {
			return newFailedErr(
				fmt.Sprintf("FileOutput does not match Pattern. Path: %s, FileOutput: %q, Pattern: %s", expectedFileOutput.Path, fileOutput, *expectedFileOutput.Pattern),
				elapsed,
			)
		}
	}

	return nil
}
