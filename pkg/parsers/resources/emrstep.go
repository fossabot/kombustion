package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// EMRStep Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-step.html
type EMRStep struct {
	Type       string            `yaml:"Type"`
	Properties EMRStepProperties `yaml:"Properties"`
	Condition  interface{}       `yaml:"Condition,omitempty"`
	Metadata   interface{}       `yaml:"Metadata,omitempty"`
	DependsOn  interface{}       `yaml:"DependsOn,omitempty"`
}

// EMRStep Properties
type EMRStepProperties struct {
	ActionOnFailure interface{}                         `yaml:"ActionOnFailure"`
	JobFlowId       interface{}                         `yaml:"JobFlowId"`
	Name            interface{}                         `yaml:"Name"`
	HadoopJarStep   *properties.StepHadoopJarStepConfig `yaml:"HadoopJarStep"`
}

// NewEMRStep constructor creates a new EMRStep
func NewEMRStep(properties EMRStepProperties, deps ...interface{}) EMRStep {
	return EMRStep{
		Type:       "AWS::EMR::Step",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseEMRStep parses EMRStep
func ParseEMRStep(name string, data string) (cf types.TemplateObject, err error) {
	var resource EMRStep
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EMRStep - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseEMRStep validator
func (resource EMRStep) Validate() []error {
	return resource.Properties.Validate()
}

// ParseEMRStepProperties validator
func (resource EMRStepProperties) Validate() []error {
	errs := []error{}
	if resource.ActionOnFailure == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ActionOnFailure'"))
	}
	if resource.JobFlowId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'JobFlowId'"))
	}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.HadoopJarStep == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'HadoopJarStep'"))
	} else {
		errs = append(errs, resource.HadoopJarStep.Validate()...)
	}
	return errs
}
