package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// DMSEventSubscription Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-eventsubscription.html
type DMSEventSubscription struct {
	Type       string                         `yaml:"Type"`
	Properties DMSEventSubscriptionProperties `yaml:"Properties"`
	Condition  interface{}                    `yaml:"Condition,omitempty"`
	Metadata   interface{}                    `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                    `yaml:"DependsOn,omitempty"`
}

// DMSEventSubscription Properties
type DMSEventSubscriptionProperties struct {
	Enabled          interface{} `yaml:"Enabled,omitempty"`
	SnsTopicArn      interface{} `yaml:"SnsTopicArn"`
	SourceType       interface{} `yaml:"SourceType,omitempty"`
	SubscriptionName interface{} `yaml:"SubscriptionName,omitempty"`
	EventCategories  interface{} `yaml:"EventCategories,omitempty"`
	SourceIds        interface{} `yaml:"SourceIds,omitempty"`
	Tags             interface{} `yaml:"Tags,omitempty"`
}

// NewDMSEventSubscription constructor creates a new DMSEventSubscription
func NewDMSEventSubscription(properties DMSEventSubscriptionProperties, deps ...interface{}) DMSEventSubscription {
	return DMSEventSubscription{
		Type:       "AWS::DMS::EventSubscription",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseDMSEventSubscription parses DMSEventSubscription
func ParseDMSEventSubscription(name string, data string) (cf types.TemplateObject, err error) {
	var resource DMSEventSubscription
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: DMSEventSubscription - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseDMSEventSubscription validator
func (resource DMSEventSubscription) Validate() []error {
	return resource.Properties.Validate()
}

// ParseDMSEventSubscriptionProperties validator
func (resource DMSEventSubscriptionProperties) Validate() []error {
	errs := []error{}
	if resource.SnsTopicArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SnsTopicArn'"))
	}
	return errs
}
