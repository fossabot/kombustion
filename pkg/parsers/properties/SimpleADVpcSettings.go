package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// SimpleADVpcSettings Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-directoryservice-simplead-vpcsettings.html
type SimpleADVpcSettings struct {
	VpcId     interface{} `yaml:"VpcId"`
	SubnetIds interface{} `yaml:"SubnetIds"`
}

// SimpleADVpcSettings validation
func (resource SimpleADVpcSettings) Validate() []error {
	errs := []error{}

	if resource.VpcId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'VpcId'"))
	}
	if resource.SubnetIds == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SubnetIds'"))
	}
	return errs
}
