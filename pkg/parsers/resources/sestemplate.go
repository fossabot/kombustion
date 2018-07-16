package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// SESTemplate Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-template.html
type SESTemplate struct {
	Type       string                `yaml:"Type"`
	Properties SESTemplateProperties `yaml:"Properties"`
	Condition  interface{}           `yaml:"Condition,omitempty"`
	Metadata   interface{}           `yaml:"Metadata,omitempty"`
	DependsOn  interface{}           `yaml:"DependsOn,omitempty"`
}

// SESTemplate Properties
type SESTemplateProperties struct {
	Template *properties.TemplateTemplate `yaml:"Template,omitempty"`
}

// NewSESTemplate constructor creates a new SESTemplate
func NewSESTemplate(properties SESTemplateProperties, deps ...interface{}) SESTemplate {
	return SESTemplate{
		Type:       "AWS::SES::Template",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseSESTemplate parses SESTemplate
func ParseSESTemplate(name string, data string) (cf types.TemplateObject, err error) {
	var resource SESTemplate
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: SESTemplate - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseSESTemplate validator
func (resource SESTemplate) Validate() []error {
	return resource.Properties.Validate()
}

// ParseSESTemplateProperties validator
func (resource SESTemplateProperties) Validate() []error {
	errs := []error{}
	return errs
}
