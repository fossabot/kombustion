package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// DMSReplicationTask Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html
type DMSReplicationTask struct {
	Type       string                       `yaml:"Type"`
	Properties DMSReplicationTaskProperties `yaml:"Properties"`
	Condition  interface{}                  `yaml:"Condition,omitempty"`
	Metadata   interface{}                  `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                  `yaml:"DependsOn,omitempty"`
}

// DMSReplicationTask Properties
type DMSReplicationTaskProperties struct {
	CdcStartTime              interface{} `yaml:"CdcStartTime,omitempty"`
	MigrationType             interface{} `yaml:"MigrationType"`
	ReplicationInstanceArn    interface{} `yaml:"ReplicationInstanceArn"`
	ReplicationTaskIdentifier interface{} `yaml:"ReplicationTaskIdentifier,omitempty"`
	ReplicationTaskSettings   interface{} `yaml:"ReplicationTaskSettings,omitempty"`
	SourceEndpointArn         interface{} `yaml:"SourceEndpointArn"`
	TableMappings             interface{} `yaml:"TableMappings"`
	TargetEndpointArn         interface{} `yaml:"TargetEndpointArn"`
	Tags                      interface{} `yaml:"Tags,omitempty"`
}

// NewDMSReplicationTask constructor creates a new DMSReplicationTask
func NewDMSReplicationTask(properties DMSReplicationTaskProperties, deps ...interface{}) DMSReplicationTask {
	return DMSReplicationTask{
		Type:       "AWS::DMS::ReplicationTask",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseDMSReplicationTask parses DMSReplicationTask
func ParseDMSReplicationTask(name string, data string) (cf types.TemplateObject, err error) {
	var resource DMSReplicationTask
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: DMSReplicationTask - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseDMSReplicationTask validator
func (resource DMSReplicationTask) Validate() []error {
	return resource.Properties.Validate()
}

// ParseDMSReplicationTaskProperties validator
func (resource DMSReplicationTaskProperties) Validate() []error {
	errs := []error{}
	if resource.MigrationType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MigrationType'"))
	}
	if resource.ReplicationInstanceArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ReplicationInstanceArn'"))
	}
	if resource.SourceEndpointArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SourceEndpointArn'"))
	}
	if resource.TableMappings == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TableMappings'"))
	}
	if resource.TargetEndpointArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TargetEndpointArn'"))
	}
	return errs
}