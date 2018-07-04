package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ApplicationReferenceDataSourceReferenceDataSource Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-referencedatasource.html
type ApplicationReferenceDataSourceReferenceDataSource struct {
	TableName             interface{}                                          `yaml:"TableName,omitempty"`
	S3ReferenceDataSource *ApplicationReferenceDataSourceS3ReferenceDataSource `yaml:"S3ReferenceDataSource,omitempty"`
	ReferenceSchema       *ApplicationReferenceDataSourceReferenceSchema       `yaml:"ReferenceSchema"`
}

// ApplicationReferenceDataSourceReferenceDataSource validation
func (resource ApplicationReferenceDataSourceReferenceDataSource) Validate() []error {
	errs := []error{}

	if resource.ReferenceSchema == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ReferenceSchema'"))
	} else {
		errs = append(errs, resource.ReferenceSchema.Validate()...)
	}
	return errs
}