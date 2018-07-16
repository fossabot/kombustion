package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// StepFunctionsActivity Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-activity.html
type StepFunctionsActivity struct {
	Type       string                          `yaml:"Type"`
	Properties StepFunctionsActivityProperties `yaml:"Properties"`
	Condition  interface{}                     `yaml:"Condition,omitempty"`
	Metadata   interface{}                     `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                     `yaml:"DependsOn,omitempty"`
}

// StepFunctionsActivity Properties
type StepFunctionsActivityProperties struct {
	Name interface{} `yaml:"Name"`
}

// NewStepFunctionsActivity constructor creates a new StepFunctionsActivity
func NewStepFunctionsActivity(properties StepFunctionsActivityProperties, deps ...interface{}) StepFunctionsActivity {
	return StepFunctionsActivity{
		Type:       "AWS::StepFunctions::Activity",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseStepFunctionsActivity parses StepFunctionsActivity
func ParseStepFunctionsActivity(name string, data string) (cf types.TemplateObject, err error) {
	var resource StepFunctionsActivity
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: StepFunctionsActivity - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseStepFunctionsActivity validator
func (resource StepFunctionsActivity) Validate() []error {
	return resource.Properties.Validate()
}

// ParseStepFunctionsActivityProperties validator
func (resource StepFunctionsActivityProperties) Validate() []error {
	errs := []error{}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	return errs
}
