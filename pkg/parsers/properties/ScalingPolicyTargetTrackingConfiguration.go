package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ScalingPolicyTargetTrackingConfiguration Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-targettrackingconfiguration.html
type ScalingPolicyTargetTrackingConfiguration struct {
	DisableScaleIn                interface{}                                 `yaml:"DisableScaleIn,omitempty"`
	TargetValue                   interface{}                                 `yaml:"TargetValue"`
	PredefinedMetricSpecification *ScalingPolicyPredefinedMetricSpecification `yaml:"PredefinedMetricSpecification,omitempty"`
	CustomizedMetricSpecification *ScalingPolicyCustomizedMetricSpecification `yaml:"CustomizedMetricSpecification,omitempty"`
}

// ScalingPolicyTargetTrackingConfiguration validation
func (resource ScalingPolicyTargetTrackingConfiguration) Validate() []error {
	errs := []error{}

	if resource.TargetValue == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TargetValue'"))
	}
	return errs
}