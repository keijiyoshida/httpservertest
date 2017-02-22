package httpservertest

import (
	"bytes"
	"os/exec"
	"strconv"
	"strings"
)

// tail executes the tail command and returns the result.
func tail(path string, n int) (string, error) {
	bf := new(bytes.Buffer)

	cmd := exec.Command("tail", "-n", strconv.Itoa(n), path)
	cmd.Stdout = bf

	if err := cmd.Run(); err != nil {
		return "", err
	}

	return strings.TrimSpace(bf.String()), nil
}
