package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// DAXCluster Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dax-cluster.html
type DAXCluster struct {
	Type       string               `yaml:"Type"`
	Properties DAXClusterProperties `yaml:"Properties"`
	Condition  interface{}          `yaml:"Condition,omitempty"`
	Metadata   interface{}          `yaml:"Metadata,omitempty"`
	DependsOn  interface{}          `yaml:"DependsOn,omitempty"`
}

// DAXCluster Properties
type DAXClusterProperties struct {
	ClusterName                interface{} `yaml:"ClusterName,omitempty"`
	Description                interface{} `yaml:"Description,omitempty"`
	IAMRoleARN                 interface{} `yaml:"IAMRoleARN"`
	NodeType                   interface{} `yaml:"NodeType"`
	NotificationTopicARN       interface{} `yaml:"NotificationTopicARN,omitempty"`
	ParameterGroupName         interface{} `yaml:"ParameterGroupName,omitempty"`
	PreferredMaintenanceWindow interface{} `yaml:"PreferredMaintenanceWindow,omitempty"`
	ReplicationFactor          interface{} `yaml:"ReplicationFactor"`
	SubnetGroupName            interface{} `yaml:"SubnetGroupName,omitempty"`
	Tags                       interface{} `yaml:"Tags,omitempty"`
	AvailabilityZones          interface{} `yaml:"AvailabilityZones,omitempty"`
	SecurityGroupIds           interface{} `yaml:"SecurityGroupIds,omitempty"`
}

// NewDAXCluster constructor creates a new DAXCluster
func NewDAXCluster(properties DAXClusterProperties, deps ...interface{}) DAXCluster {
	return DAXCluster{
		Type:       "AWS::DAX::Cluster",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseDAXCluster parses DAXCluster
func ParseDAXCluster(name string, data string) (cf types.TemplateObject, err error) {
	var resource DAXCluster
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: DAXCluster - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseDAXCluster validator
func (resource DAXCluster) Validate() []error {
	return resource.Properties.Validate()
}

// ParseDAXClusterProperties validator
func (resource DAXClusterProperties) Validate() []error {
	errs := []error{}
	if resource.IAMRoleARN == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'IAMRoleARN'"))
	}
	if resource.NodeType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'NodeType'"))
	}
	if resource.ReplicationFactor == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ReplicationFactor'"))
	}
	return errs
}