package httpservertest

import (
	"bytes"
	"html/template"

	yaml "gopkg.in/yaml.v2"
)

// newTestCases parses the test cases file,
// creates test cases and returns them.
func newTestCases(path string, parameters interface{}) ([]TestCase, error) {
	t, err := template.ParseFiles(path)
	if err != nil {
		return nil, err
	}

	b := new(bytes.Buffer)

	if err := t.Execute(b, TestCasesTemplateData{parameters}); err != nil {
		return nil, err
	}

	var testCases []TestCase

	if err := yaml.Unmarshal(b.Bytes(), &testCases); err != nil {
		return nil, err
	}

	return testCases, nil
}
