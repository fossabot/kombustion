package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// WAFRegionalIPSet Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-ipset.html
type WAFRegionalIPSet struct {
	Type       string                     `yaml:"Type"`
	Properties WAFRegionalIPSetProperties `yaml:"Properties"`
	Condition  interface{}                `yaml:"Condition,omitempty"`
	Metadata   interface{}                `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                `yaml:"DependsOn,omitempty"`
}

// WAFRegionalIPSet Properties
type WAFRegionalIPSetProperties struct {
	Name             interface{} `yaml:"Name"`
	IPSetDescriptors interface{} `yaml:"IPSetDescriptors,omitempty"`
}

// NewWAFRegionalIPSet constructor creates a new WAFRegionalIPSet
func NewWAFRegionalIPSet(properties WAFRegionalIPSetProperties, deps ...interface{}) WAFRegionalIPSet {
	return WAFRegionalIPSet{
		Type:       "AWS::WAFRegional::IPSet",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseWAFRegionalIPSet parses WAFRegionalIPSet
func ParseWAFRegionalIPSet(name string, data string) (cf types.TemplateObject, err error) {
	var resource WAFRegionalIPSet
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: WAFRegionalIPSet - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseWAFRegionalIPSet validator
func (resource WAFRegionalIPSet) Validate() []error {
	return resource.Properties.Validate()
}

// ParseWAFRegionalIPSetProperties validator
func (resource WAFRegionalIPSetProperties) Validate() []error {
	errs := []error{}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	return errs
}
