package httpservertest

// Expires represents an expected expires value.
type Expires struct {
	// AfterOrEqualToRequestTimePlus is a check condition of Expires.
	AfterOrEqualToRequestTimePlus *int64 `yaml:"AfterOrEqualToRequestTimePlus"`
}
