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

// KinesisStream Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-stream.html
type KinesisStream struct {
	Type       string                  `yaml:"Type"`
	Properties KinesisStreamProperties `yaml:"Properties"`
	Condition  interface{}             `yaml:"Condition,omitempty"`
	Metadata   interface{}             `yaml:"Metadata,omitempty"`
	DependsOn  interface{}             `yaml:"DependsOn,omitempty"`
}

// KinesisStream Properties
type KinesisStreamProperties struct {
	Name                 interface{}                        `yaml:"Name,omitempty"`
	RetentionPeriodHours interface{}                        `yaml:"RetentionPeriodHours,omitempty"`
	ShardCount           interface{}                        `yaml:"ShardCount"`
	StreamEncryption     *properties.StreamStreamEncryption `yaml:"StreamEncryption,omitempty"`
	Tags                 interface{}                        `yaml:"Tags,omitempty"`
}

// NewKinesisStream constructor creates a new KinesisStream
func NewKinesisStream(properties KinesisStreamProperties, deps ...interface{}) KinesisStream {
	return KinesisStream{
		Type:       "AWS::Kinesis::Stream",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseKinesisStream parses KinesisStream
func ParseKinesisStream(name string, data string) (cf types.TemplateObject, err error) {
	var resource KinesisStream
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: KinesisStream - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseKinesisStream validator
func (resource KinesisStream) Validate() []error {
	return resource.Properties.Validate()
}

// ParseKinesisStreamProperties validator
func (resource KinesisStreamProperties) Validate() []error {
	errs := []error{}
	if resource.ShardCount == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ShardCount'"))
	}
	return errs
}
