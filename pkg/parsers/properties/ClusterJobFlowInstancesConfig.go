package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ClusterJobFlowInstancesConfig Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html
type ClusterJobFlowInstancesConfig struct {
	Ec2KeyName                     interface{}                 `yaml:"Ec2KeyName,omitempty"`
	Ec2SubnetId                    interface{}                 `yaml:"Ec2SubnetId,omitempty"`
	EmrManagedMasterSecurityGroup  interface{}                 `yaml:"EmrManagedMasterSecurityGroup,omitempty"`
	EmrManagedSlaveSecurityGroup   interface{}                 `yaml:"EmrManagedSlaveSecurityGroup,omitempty"`
	HadoopVersion                  interface{}                 `yaml:"HadoopVersion,omitempty"`
	ServiceAccessSecurityGroup     interface{}                 `yaml:"ServiceAccessSecurityGroup,omitempty"`
	TerminationProtected           interface{}                 `yaml:"TerminationProtected,omitempty"`
	Placement                      *ClusterPlacementType       `yaml:"Placement,omitempty"`
	AdditionalSlaveSecurityGroups  interface{}                 `yaml:"AdditionalSlaveSecurityGroups,omitempty"`
	AdditionalMasterSecurityGroups interface{}                 `yaml:"AdditionalMasterSecurityGroups,omitempty"`
	CoreInstanceGroup              *ClusterInstanceGroupConfig `yaml:"CoreInstanceGroup,omitempty"`
	MasterInstanceGroup            *ClusterInstanceGroupConfig `yaml:"MasterInstanceGroup,omitempty"`
	CoreInstanceFleet              *ClusterInstanceFleetConfig `yaml:"CoreInstanceFleet,omitempty"`
	MasterInstanceFleet            *ClusterInstanceFleetConfig `yaml:"MasterInstanceFleet,omitempty"`
}

// ClusterJobFlowInstancesConfig validation
func (resource ClusterJobFlowInstancesConfig) Validate() []error {
	errs := []error{}

	return errs
}
