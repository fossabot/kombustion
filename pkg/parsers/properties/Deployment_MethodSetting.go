package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

type Deployment_MethodSetting struct {
	CacheDataEncrypted   interface{} `yaml:"CacheDataEncrypted,omitempty"`
	CacheTtlInSeconds    interface{} `yaml:"CacheTtlInSeconds,omitempty"`
	CachingEnabled       interface{} `yaml:"CachingEnabled,omitempty"`
	DataTraceEnabled     interface{} `yaml:"DataTraceEnabled,omitempty"`
	HttpMethod           interface{} `yaml:"HttpMethod,omitempty"`
	LoggingLevel         interface{} `yaml:"LoggingLevel,omitempty"`
	MetricsEnabled       interface{} `yaml:"MetricsEnabled,omitempty"`
	ResourcePath         interface{} `yaml:"ResourcePath,omitempty"`
	ThrottlingBurstLimit interface{} `yaml:"ThrottlingBurstLimit,omitempty"`
	ThrottlingRateLimit  interface{} `yaml:"ThrottlingRateLimit,omitempty"`
}

func (resource Deployment_MethodSetting) Validate() []error {
	errs := []error{}

	return errs
}