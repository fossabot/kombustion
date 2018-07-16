package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// GlueDevEndpoint Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-devendpoint.html
type GlueDevEndpoint struct {
	Type       string                    `yaml:"Type"`
	Properties GlueDevEndpointProperties `yaml:"Properties"`
	Condition  interface{}               `yaml:"Condition,omitempty"`
	Metadata   interface{}               `yaml:"Metadata,omitempty"`
	DependsOn  interface{}               `yaml:"DependsOn,omitempty"`
}

// GlueDevEndpoint Properties
type GlueDevEndpointProperties struct {
	EndpointName          interface{} `yaml:"EndpointName,omitempty"`
	ExtraJarsS3Path       interface{} `yaml:"ExtraJarsS3Path,omitempty"`
	ExtraPythonLibsS3Path interface{} `yaml:"ExtraPythonLibsS3Path,omitempty"`
	NumberOfNodes         interface{} `yaml:"NumberOfNodes,omitempty"`
	PublicKey             interface{} `yaml:"PublicKey"`
	RoleArn               interface{} `yaml:"RoleArn"`
	SubnetId              interface{} `yaml:"SubnetId,omitempty"`
	SecurityGroupIds      interface{} `yaml:"SecurityGroupIds,omitempty"`
}

// NewGlueDevEndpoint constructor creates a new GlueDevEndpoint
func NewGlueDevEndpoint(properties GlueDevEndpointProperties, deps ...interface{}) GlueDevEndpoint {
	return GlueDevEndpoint{
		Type:       "AWS::Glue::DevEndpoint",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseGlueDevEndpoint parses GlueDevEndpoint
func ParseGlueDevEndpoint(name string, data string) (cf types.TemplateObject, err error) {
	var resource GlueDevEndpoint
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: GlueDevEndpoint - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseGlueDevEndpoint validator
func (resource GlueDevEndpoint) Validate() []error {
	return resource.Properties.Validate()
}

// ParseGlueDevEndpointProperties validator
func (resource GlueDevEndpointProperties) Validate() []error {
	errs := []error{}
	if resource.PublicKey == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'PublicKey'"))
	}
	if resource.RoleArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RoleArn'"))
	}
	return errs
}
