package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// InstanceGroupConfigScalingAction Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-scalingaction.html
type InstanceGroupConfigScalingAction struct {
	Market                           interface{}                                          `yaml:"Market,omitempty"`
	SimpleScalingPolicyConfiguration *InstanceGroupConfigSimpleScalingPolicyConfiguration `yaml:"SimpleScalingPolicyConfiguration"`
}

// InstanceGroupConfigScalingAction validation
func (resource InstanceGroupConfigScalingAction) Validate() []error {
	errs := []error{}

	if resource.SimpleScalingPolicyConfiguration == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SimpleScalingPolicyConfiguration'"))
	} else {
		errs = append(errs, resource.SimpleScalingPolicyConfiguration.Validate()...)
	}
	return errs
}
