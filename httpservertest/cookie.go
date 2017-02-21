package httpservertest

// Cookie represents a Cookie of an expected response.
type Cookie struct {
	// Name is an expected name.
	Name *string `yaml:"Name"`
	// Value is an expected value.
	Value *string `yaml:"Value"`
	// Path is an expected path.
	Path *string `yaml:"Path"`
	// Domain is an expected domain.
	Domain *string `yaml:"Domain"`
	// Secure is an expected secure value.
	Secure *bool `yaml:"Secure"`
	// Expires is an expected expires value.
	Expires *Expires `yaml:"Expires"`
}
