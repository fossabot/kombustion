package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// EC2SecurityGroupIngress Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-ingress.html
type EC2SecurityGroupIngress struct {
	Type       string                            `yaml:"Type"`
	Properties EC2SecurityGroupIngressProperties `yaml:"Properties"`
	Condition  interface{}                       `yaml:"Condition,omitempty"`
	Metadata   interface{}                       `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                       `yaml:"DependsOn,omitempty"`
}

// EC2SecurityGroupIngress Properties
type EC2SecurityGroupIngressProperties struct {
	CidrIp                     interface{} `yaml:"CidrIp,omitempty"`
	CidrIpv6                   interface{} `yaml:"CidrIpv6,omitempty"`
	Description                interface{} `yaml:"Description,omitempty"`
	FromPort                   interface{} `yaml:"FromPort,omitempty"`
	GroupId                    interface{} `yaml:"GroupId,omitempty"`
	GroupName                  interface{} `yaml:"GroupName,omitempty"`
	IpProtocol                 interface{} `yaml:"IpProtocol"`
	SourceSecurityGroupId      interface{} `yaml:"SourceSecurityGroupId,omitempty"`
	SourceSecurityGroupName    interface{} `yaml:"SourceSecurityGroupName,omitempty"`
	SourceSecurityGroupOwnerId interface{} `yaml:"SourceSecurityGroupOwnerId,omitempty"`
	ToPort                     interface{} `yaml:"ToPort,omitempty"`
}

// NewEC2SecurityGroupIngress constructor creates a new EC2SecurityGroupIngress
func NewEC2SecurityGroupIngress(properties EC2SecurityGroupIngressProperties, deps ...interface{}) EC2SecurityGroupIngress {
	return EC2SecurityGroupIngress{
		Type:       "AWS::EC2::SecurityGroupIngress",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseEC2SecurityGroupIngress parses EC2SecurityGroupIngress
func ParseEC2SecurityGroupIngress(name string, data string) (cf types.TemplateObject, err error) {
	var resource EC2SecurityGroupIngress
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2SecurityGroupIngress - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseEC2SecurityGroupIngress validator
func (resource EC2SecurityGroupIngress) Validate() []error {
	return resource.Properties.Validate()
}

// ParseEC2SecurityGroupIngressProperties validator
func (resource EC2SecurityGroupIngressProperties) Validate() []error {
	errs := []error{}
	if resource.IpProtocol == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'IpProtocol'"))
	}
	return errs
}
