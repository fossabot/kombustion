package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// VPNConnectionVpnTunnelOptionsSpecification Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html
type VPNConnectionVpnTunnelOptionsSpecification struct {
	PreSharedKey     interface{} `yaml:"PreSharedKey,omitempty"`
	TunnelInsideCidr interface{} `yaml:"TunnelInsideCidr,omitempty"`
}

// VPNConnectionVpnTunnelOptionsSpecification validation
func (resource VPNConnectionVpnTunnelOptionsSpecification) Validate() []error {
	errs := []error{}

	return errs
}
