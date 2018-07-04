package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// DeploymentGroupDeployment Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment.html
type DeploymentGroupDeployment struct {
	Description                   interface{}                      `yaml:"Description,omitempty"`
	IgnoreApplicationStopFailures interface{}                      `yaml:"IgnoreApplicationStopFailures,omitempty"`
	Revision                      *DeploymentGroupRevisionLocation `yaml:"Revision"`
}

// DeploymentGroupDeployment validation
func (resource DeploymentGroupDeployment) Validate() []error {
	errs := []error{}

	if resource.Revision == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Revision'"))
	} else {
		errs = append(errs, resource.Revision.Validate()...)
	}
	return errs
}