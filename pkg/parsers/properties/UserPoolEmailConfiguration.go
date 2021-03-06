package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// UserPoolEmailConfiguration Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-emailconfiguration.html
type UserPoolEmailConfiguration struct {
	ReplyToEmailAddress interface{} `yaml:"ReplyToEmailAddress,omitempty"`
	SourceArn           interface{} `yaml:"SourceArn,omitempty"`
}

// UserPoolEmailConfiguration validation
func (resource UserPoolEmailConfiguration) Validate() []error {
	errs := []error{}

	return errs
}
