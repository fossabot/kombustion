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

// ElasticLoadBalancingLoadBalancer Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html
type ElasticLoadBalancingLoadBalancer struct {
	Type       string                                     `yaml:"Type"`
	Properties ElasticLoadBalancingLoadBalancerProperties `yaml:"Properties"`
	Condition  interface{}                                `yaml:"Condition,omitempty"`
	Metadata   interface{}                                `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                                `yaml:"DependsOn,omitempty"`
}

// ElasticLoadBalancingLoadBalancer Properties
type ElasticLoadBalancingLoadBalancerProperties struct {
	CrossZone                 interface{}                                      `yaml:"CrossZone,omitempty"`
	LoadBalancerName          interface{}                                      `yaml:"LoadBalancerName,omitempty"`
	Scheme                    interface{}                                      `yaml:"Scheme,omitempty"`
	Tags                      interface{}                                      `yaml:"Tags,omitempty"`
	AppCookieStickinessPolicy interface{}                                      `yaml:"AppCookieStickinessPolicy,omitempty"`
	AvailabilityZones         interface{}                                      `yaml:"AvailabilityZones,omitempty"`
	Subnets                   interface{}                                      `yaml:"Subnets,omitempty"`
	SecurityGroups            interface{}                                      `yaml:"SecurityGroups,omitempty"`
	Policies                  interface{}                                      `yaml:"Policies,omitempty"`
	Instances                 interface{}                                      `yaml:"Instances,omitempty"`
	LBCookieStickinessPolicy  interface{}                                      `yaml:"LBCookieStickinessPolicy,omitempty"`
	Listeners                 interface{}                                      `yaml:"Listeners"`
	HealthCheck               *properties.LoadBalancerHealthCheck              `yaml:"HealthCheck,omitempty"`
	ConnectionSettings        *properties.LoadBalancerConnectionSettings       `yaml:"ConnectionSettings,omitempty"`
	ConnectionDrainingPolicy  *properties.LoadBalancerConnectionDrainingPolicy `yaml:"ConnectionDrainingPolicy,omitempty"`
	AccessLoggingPolicy       *properties.LoadBalancerAccessLoggingPolicy      `yaml:"AccessLoggingPolicy,omitempty"`
}

// NewElasticLoadBalancingLoadBalancer constructor creates a new ElasticLoadBalancingLoadBalancer
func NewElasticLoadBalancingLoadBalancer(properties ElasticLoadBalancingLoadBalancerProperties, deps ...interface{}) ElasticLoadBalancingLoadBalancer {
	return ElasticLoadBalancingLoadBalancer{
		Type:       "AWS::ElasticLoadBalancing::LoadBalancer",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseElasticLoadBalancingLoadBalancer parses ElasticLoadBalancingLoadBalancer
func ParseElasticLoadBalancingLoadBalancer(name string, data string) (cf types.TemplateObject, err error) {
	var resource ElasticLoadBalancingLoadBalancer
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ElasticLoadBalancingLoadBalancer - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseElasticLoadBalancingLoadBalancer validator
func (resource ElasticLoadBalancingLoadBalancer) Validate() []error {
	return resource.Properties.Validate()
}

// ParseElasticLoadBalancingLoadBalancerProperties validator
func (resource ElasticLoadBalancingLoadBalancerProperties) Validate() []error {
	errs := []error{}
	if resource.Listeners == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Listeners'"))
	}
	return errs
}