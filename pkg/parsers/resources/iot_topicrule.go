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

type IoTTopicRule struct {
	Type       string                 `yaml:"Type"`
	Properties IoTTopicRuleProperties `yaml:"Properties"`
	Condition  interface{}            `yaml:"Condition,omitempty"`
	Metadata   interface{}            `yaml:"Metadata,omitempty"`
	DependsOn  interface{}            `yaml:"DependsOn,omitempty"`
}

type IoTTopicRuleProperties struct {
	RuleName         interface{}                            `yaml:"RuleName,omitempty"`
	TopicRulePayload *properties.TopicRule_TopicRulePayload `yaml:"TopicRulePayload"`
}

func NewIoTTopicRule(properties IoTTopicRuleProperties, deps ...interface{}) IoTTopicRule {
	return IoTTopicRule{
		Type:       "AWS::IoT::TopicRule",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseIoTTopicRule(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource IoTTopicRule
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: IoTTopicRule - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

func (resource IoTTopicRule) Validate() []error {
	return resource.Properties.Validate()
}

func (resource IoTTopicRuleProperties) Validate() []error {
	errs := []error{}
	if resource.TopicRulePayload == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TopicRulePayload'"))
	} else {
		errs = append(errs, resource.TopicRulePayload.Validate()...)
	}
	return errs
}