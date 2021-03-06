package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// BucketRedirectAllRequestsTo Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-redirectallrequeststo.html
type BucketRedirectAllRequestsTo struct {
	HostName interface{} `yaml:"HostName"`
	Protocol interface{} `yaml:"Protocol,omitempty"`
}

// BucketRedirectAllRequestsTo validation
func (resource BucketRedirectAllRequestsTo) Validate() []error {
	errs := []error{}

	if resource.HostName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'HostName'"))
	}
	return errs
}
