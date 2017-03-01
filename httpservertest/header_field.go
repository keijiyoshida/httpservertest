package httpservertest

// HeaderField represents a field of an expected Header.
type HeaderField struct {
	Name  string `yaml:"Name"`
	Value string `yaml:"Value"`
}
