package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// CodeDeployApplication Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-application.html
type CodeDeployApplication struct {
	Type       string                          `yaml:"Type"`
	Properties CodeDeployApplicationProperties `yaml:"Properties"`
	Condition  interface{}                     `yaml:"Condition,omitempty"`
	Metadata   interface{}                     `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                     `yaml:"DependsOn,omitempty"`
}

// CodeDeployApplication Properties
type CodeDeployApplicationProperties struct {
	ApplicationName interface{} `yaml:"ApplicationName,omitempty"`
	ComputePlatform interface{} `yaml:"ComputePlatform,omitempty"`
}

// NewCodeDeployApplication constructor creates a new CodeDeployApplication
func NewCodeDeployApplication(properties CodeDeployApplicationProperties, deps ...interface{}) CodeDeployApplication {
	return CodeDeployApplication{
		Type:       "AWS::CodeDeploy::Application",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseCodeDeployApplication parses CodeDeployApplication
func ParseCodeDeployApplication(name string, data string) (cf types.TemplateObject, err error) {
	var resource CodeDeployApplication
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: CodeDeployApplication - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseCodeDeployApplication validator
func (resource CodeDeployApplication) Validate() []error {
	return resource.Properties.Validate()
}

// ParseCodeDeployApplicationProperties validator
func (resource CodeDeployApplicationProperties) Validate() []error {
	errs := []error{}
	return errs
}
