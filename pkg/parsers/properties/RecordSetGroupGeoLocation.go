package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// RecordSetGroupGeoLocation Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset-geolocation.html
type RecordSetGroupGeoLocation struct {
	ContinentCode   interface{} `yaml:"ContinentCode,omitempty"`
	CountryCode     interface{} `yaml:"CountryCode,omitempty"`
	SubdivisionCode interface{} `yaml:"SubdivisionCode,omitempty"`
}

// RecordSetGroupGeoLocation validation
func (resource RecordSetGroupGeoLocation) Validate() []error {
	errs := []error{}

	return errs
}
