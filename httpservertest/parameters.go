package httpservertest

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// newParameters parses the parameters file,
// creates parameters and returns them.
func newParameters(path string) (interface{}, error) {
	if path == "" {
		return nil, nil
	}

	in, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var parameters interface{}

	if err := yaml.Unmarshal(in, &parameters); err != nil {
		return nil, err
	}

	return parameters, nil
}
