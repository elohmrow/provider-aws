/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ApplicationObservation struct {

	// The ARN of the Kinesis Analytics Appliation.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The CloudWatch log stream options to monitor application errors.
	// See CloudWatch Logging Options below for more details.
	// +kubebuilder:validation:Optional
	CloudwatchLoggingOptions []CloudwatchLoggingOptionsObservation `json:"cloudwatchLoggingOptions,omitempty" tf:"cloudwatch_logging_options,omitempty"`

	// The Timestamp when the application version was created.
	CreateTimestamp *string `json:"createTimestamp,omitempty" tf:"create_timestamp,omitempty"`

	// The ARN of the Kinesis Analytics Application.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Input configuration of the application. See Inputs below for more details.
	// +kubebuilder:validation:Optional
	Inputs []InputsObservation `json:"inputs,omitempty" tf:"inputs,omitempty"`

	// The Timestamp when the application was last updated.
	LastUpdateTimestamp *string `json:"lastUpdateTimestamp,omitempty" tf:"last_update_timestamp,omitempty"`

	// Output destination configuration of the application. See Outputs below for more details.
	// +kubebuilder:validation:Optional
	Outputs []OutputsObservation `json:"outputs,omitempty" tf:"outputs,omitempty"`

	// An S3 Reference Data Source for the application.
	// See Reference Data Sources below for more details.
	// +kubebuilder:validation:Optional
	ReferenceDataSources []ReferenceDataSourcesObservation `json:"referenceDataSources,omitempty" tf:"reference_data_sources,omitempty"`

	// The Status of the application.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The Version of the application.
	Version *float64 `json:"version,omitempty" tf:"version,omitempty"`
}

type ApplicationParameters struct {

	// The CloudWatch log stream options to monitor application errors.
	// See CloudWatch Logging Options below for more details.
	// +kubebuilder:validation:Optional
	CloudwatchLoggingOptions []CloudwatchLoggingOptionsParameters `json:"cloudwatchLoggingOptions,omitempty" tf:"cloudwatch_logging_options,omitempty"`

	// SQL Code to transform input data, and generate output.
	// +kubebuilder:validation:Optional
	Code *string `json:"code,omitempty" tf:"code,omitempty"`

	// Description of the application.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Input configuration of the application. See Inputs below for more details.
	// +kubebuilder:validation:Optional
	Inputs []InputsParameters `json:"inputs,omitempty" tf:"inputs,omitempty"`

	// Output destination configuration of the application. See Outputs below for more details.
	// +kubebuilder:validation:Optional
	Outputs []OutputsParameters `json:"outputs,omitempty" tf:"outputs,omitempty"`

	// An S3 Reference Data Source for the application.
	// See Reference Data Sources below for more details.
	// +kubebuilder:validation:Optional
	ReferenceDataSources []ReferenceDataSourcesParameters `json:"referenceDataSources,omitempty" tf:"reference_data_sources,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Whether to start or stop the Kinesis Analytics Application. To start an application, an input with a defined starting_position must be configured.
	// To modify an application's starting position, first stop the application by setting start_application = false, then update starting_position and set start_application = true.
	// +kubebuilder:validation:Optional
	StartApplication *bool `json:"startApplication,omitempty" tf:"start_application,omitempty"`

	// Key-value map of tags for the Kinesis Analytics Application. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type CloudwatchLoggingOptionsObservation struct {

	// The ARN of the Kinesis Analytics Application.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type CloudwatchLoggingOptionsParameters struct {

	// The ARN of the CloudWatch Log Stream.
	// +kubebuilder:validation:Required
	LogStreamArn *string `json:"logStreamArn" tf:"log_stream_arn,omitempty"`

	// The ARN of the IAM Role used to send application messages.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`
}

type CsvObservation struct {
}

type CsvParameters struct {

	// The Column Delimiter.
	// +kubebuilder:validation:Required
	RecordColumnDelimiter *string `json:"recordColumnDelimiter" tf:"record_column_delimiter,omitempty"`

	// The Row Delimiter.
	// +kubebuilder:validation:Required
	RecordRowDelimiter *string `json:"recordRowDelimiter" tf:"record_row_delimiter,omitempty"`
}

type InputsObservation struct {

	// The ARN of the Kinesis Analytics Application.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Schema format of the data in the streaming source. See Source Schema below for more details.
	// +kubebuilder:validation:Required
	Schema []SchemaObservation `json:"schema,omitempty" tf:"schema,omitempty"`

	StreamNames []*string `json:"streamNames,omitempty" tf:"stream_names,omitempty"`
}

type InputsParameters struct {

	// The Kinesis Firehose configuration for the streaming source. Conflicts with kinesis_stream.
	// See Kinesis Firehose below for more details.
	// +kubebuilder:validation:Optional
	KinesisFirehose []KinesisFirehoseParameters `json:"kinesisFirehose,omitempty" tf:"kinesis_firehose,omitempty"`

	// The Kinesis Stream configuration for the streaming source. Conflicts with kinesis_firehose.
	// See Kinesis Stream below for more details.
	// +kubebuilder:validation:Optional
	KinesisStream []KinesisStreamParameters `json:"kinesisStream,omitempty" tf:"kinesis_stream,omitempty"`

	// The Name Prefix to use when creating an in-application stream.
	// +kubebuilder:validation:Required
	NamePrefix *string `json:"namePrefix" tf:"name_prefix,omitempty"`

	// The number of Parallel in-application streams to create.
	// See Parallelism below for more details.
	// +kubebuilder:validation:Optional
	Parallelism []ParallelismParameters `json:"parallelism,omitempty" tf:"parallelism,omitempty"`

	// The Processing Configuration to transform records as they are received from the stream.
	// See Processing Configuration below for more details.
	// +kubebuilder:validation:Optional
	ProcessingConfiguration []ProcessingConfigurationParameters `json:"processingConfiguration,omitempty" tf:"processing_configuration,omitempty"`

	// The Schema format of the data in the streaming source. See Source Schema below for more details.
	// +kubebuilder:validation:Required
	Schema []SchemaParameters `json:"schema" tf:"schema,omitempty"`

	// The point at which the application starts processing records from the streaming source.
	// See Starting Position Configuration below for more details.
	// +kubebuilder:validation:Optional
	StartingPositionConfiguration []StartingPositionConfigurationParameters `json:"startingPositionConfiguration,omitempty" tf:"starting_position_configuration,omitempty"`
}

type JSONObservation struct {
}

type JSONParameters struct {

	// Path to the top-level parent that contains the records.
	// +kubebuilder:validation:Required
	RecordRowPath *string `json:"recordRowPath" tf:"record_row_path,omitempty"`
}

type KinesisFirehoseObservation struct {
}

type KinesisFirehoseParameters struct {

	// The ARN of the Lambda function.
	// +kubebuilder:validation:Required
	ResourceArn *string `json:"resourceArn" tf:"resource_arn,omitempty"`

	// The IAM Role ARN to read the data.
	// +kubebuilder:validation:Required
	RoleArn *string `json:"roleArn" tf:"role_arn,omitempty"`
}

type KinesisStreamObservation struct {
}

type KinesisStreamParameters struct {

	// The ARN of the Lambda function.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kinesis/v1beta1.Stream
	// +kubebuilder:validation:Optional
	ResourceArn *string `json:"resourceArn,omitempty" tf:"resource_arn,omitempty"`

	// Reference to a Stream in kinesis to populate resourceArn.
	// +kubebuilder:validation:Optional
	ResourceArnRef *v1.Reference `json:"resourceArnRef,omitempty" tf:"-"`

	// Selector for a Stream in kinesis to populate resourceArn.
	// +kubebuilder:validation:Optional
	ResourceArnSelector *v1.Selector `json:"resourceArnSelector,omitempty" tf:"-"`

	// The IAM Role ARN to read the data.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`
}

type LambdaObservation struct {
}

type LambdaParameters struct {

	// The ARN of the Lambda function.
	// +kubebuilder:validation:Required
	ResourceArn *string `json:"resourceArn" tf:"resource_arn,omitempty"`

	// The IAM Role ARN to read the data.
	// +kubebuilder:validation:Required
	RoleArn *string `json:"roleArn" tf:"role_arn,omitempty"`
}

type MappingParametersCsvObservation struct {
}

type MappingParametersCsvParameters struct {

	// The Column Delimiter.
	// +kubebuilder:validation:Required
	RecordColumnDelimiter *string `json:"recordColumnDelimiter" tf:"record_column_delimiter,omitempty"`

	// The Row Delimiter.
	// +kubebuilder:validation:Required
	RecordRowDelimiter *string `json:"recordRowDelimiter" tf:"record_row_delimiter,omitempty"`
}

type MappingParametersJSONObservation struct {
}

type MappingParametersJSONParameters struct {

	// Path to the top-level parent that contains the records.
	// +kubebuilder:validation:Required
	RecordRowPath *string `json:"recordRowPath" tf:"record_row_path,omitempty"`
}

type MappingParametersObservation struct {
}

type MappingParametersParameters struct {

	// Mapping information when the record format uses delimiters.
	// See CSV Mapping Parameters below for more details.
	// +kubebuilder:validation:Optional
	Csv []CsvParameters `json:"csv,omitempty" tf:"csv,omitempty"`

	// Mapping information when JSON is the record format on the streaming source.
	// See JSON Mapping Parameters below for more details.
	// +kubebuilder:validation:Optional
	JSON []JSONParameters `json:"json,omitempty" tf:"json,omitempty"`
}

type OutputsKinesisFirehoseObservation struct {
}

type OutputsKinesisFirehoseParameters struct {

	// The ARN of the Lambda function.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/firehose/v1beta1.DeliveryStream
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",false)
	// +kubebuilder:validation:Optional
	ResourceArn *string `json:"resourceArn,omitempty" tf:"resource_arn,omitempty"`

	// Reference to a DeliveryStream in firehose to populate resourceArn.
	// +kubebuilder:validation:Optional
	ResourceArnRef *v1.Reference `json:"resourceArnRef,omitempty" tf:"-"`

	// Selector for a DeliveryStream in firehose to populate resourceArn.
	// +kubebuilder:validation:Optional
	ResourceArnSelector *v1.Selector `json:"resourceArnSelector,omitempty" tf:"-"`

	// The IAM Role ARN to read the data.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`
}

type OutputsKinesisStreamObservation struct {
}

type OutputsKinesisStreamParameters struct {

	// The ARN of the Lambda function.
	// +kubebuilder:validation:Required
	ResourceArn *string `json:"resourceArn" tf:"resource_arn,omitempty"`

	// The IAM Role ARN to read the data.
	// +kubebuilder:validation:Required
	RoleArn *string `json:"roleArn" tf:"role_arn,omitempty"`
}

type OutputsLambdaObservation struct {
}

type OutputsLambdaParameters struct {

	// The ARN of the Lambda function.
	// +kubebuilder:validation:Required
	ResourceArn *string `json:"resourceArn" tf:"resource_arn,omitempty"`

	// The IAM Role ARN to read the data.
	// +kubebuilder:validation:Required
	RoleArn *string `json:"roleArn" tf:"role_arn,omitempty"`
}

type OutputsObservation struct {

	// The ARN of the Kinesis Analytics Application.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type OutputsParameters struct {

	// The Kinesis Firehose configuration for the destination stream. Conflicts with kinesis_stream.
	// See Kinesis Firehose below for more details.
	// +kubebuilder:validation:Optional
	KinesisFirehose []OutputsKinesisFirehoseParameters `json:"kinesisFirehose,omitempty" tf:"kinesis_firehose,omitempty"`

	// The Kinesis Stream configuration for the destination stream. Conflicts with kinesis_firehose.
	// See Kinesis Stream below for more details.
	// +kubebuilder:validation:Optional
	KinesisStream []OutputsKinesisStreamParameters `json:"kinesisStream,omitempty" tf:"kinesis_stream,omitempty"`

	// The Lambda function destination. See Lambda below for more details.
	// +kubebuilder:validation:Optional
	Lambda []OutputsLambdaParameters `json:"lambda,omitempty" tf:"lambda,omitempty"`

	// The Name of the in-application stream.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The Schema format of the data written to the destination. See Destination Schema below for more details.
	// +kubebuilder:validation:Required
	Schema []OutputsSchemaParameters `json:"schema" tf:"schema,omitempty"`
}

type OutputsSchemaObservation struct {
}

type OutputsSchemaParameters struct {

	// The Format Type of the records on the output stream. Can be CSV or JSON.
	// +kubebuilder:validation:Required
	RecordFormatType *string `json:"recordFormatType" tf:"record_format_type,omitempty"`
}

type ParallelismObservation struct {
}

type ParallelismParameters struct {

	// The Count of streams.
	// +kubebuilder:validation:Optional
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`
}

type ProcessingConfigurationObservation struct {
}

type ProcessingConfigurationParameters struct {

	// The Lambda function configuration. See Lambda below for more details.
	// +kubebuilder:validation:Required
	Lambda []LambdaParameters `json:"lambda" tf:"lambda,omitempty"`
}

type RecordColumnsObservation struct {
}

type RecordColumnsParameters struct {

	// The Mapping reference to the data element.
	// +kubebuilder:validation:Optional
	Mapping *string `json:"mapping,omitempty" tf:"mapping,omitempty"`

	// Name of the column.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The SQL Type of the column.
	// +kubebuilder:validation:Required
	SQLType *string `json:"sqlType" tf:"sql_type,omitempty"`
}

type RecordFormatMappingParametersObservation struct {
}

type RecordFormatMappingParametersParameters struct {

	// Mapping information when the record format uses delimiters.
	// See CSV Mapping Parameters below for more details.
	// +kubebuilder:validation:Optional
	Csv []MappingParametersCsvParameters `json:"csv,omitempty" tf:"csv,omitempty"`

	// Mapping information when JSON is the record format on the streaming source.
	// See JSON Mapping Parameters below for more details.
	// +kubebuilder:validation:Optional
	JSON []MappingParametersJSONParameters `json:"json,omitempty" tf:"json,omitempty"`
}

type RecordFormatObservation struct {

	// The Format Type of the records on the output stream. Can be CSV or JSON.
	RecordFormatType *string `json:"recordFormatType,omitempty" tf:"record_format_type,omitempty"`
}

type RecordFormatParameters struct {

	// The Mapping Information for the record format.
	// See Mapping Parameters below for more details.
	// +kubebuilder:validation:Optional
	MappingParameters []MappingParametersParameters `json:"mappingParameters,omitempty" tf:"mapping_parameters,omitempty"`
}

type ReferenceDataSourcesObservation struct {

	// The ARN of the Kinesis Analytics Application.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Schema format of the data in the streaming source. See Source Schema below for more details.
	// +kubebuilder:validation:Required
	Schema []ReferenceDataSourcesSchemaObservation `json:"schema,omitempty" tf:"schema,omitempty"`
}

type ReferenceDataSourcesParameters struct {

	// The S3 configuration for the reference data source. See S3 Reference below for more details.
	// +kubebuilder:validation:Required
	S3 []S3Parameters `json:"s3" tf:"s3,omitempty"`

	// The Schema format of the data in the streaming source. See Source Schema below for more details.
	// +kubebuilder:validation:Required
	Schema []ReferenceDataSourcesSchemaParameters `json:"schema" tf:"schema,omitempty"`

	// The in-application Table Name.
	// +kubebuilder:validation:Required
	TableName *string `json:"tableName" tf:"table_name,omitempty"`
}

type ReferenceDataSourcesSchemaObservation struct {

	// The Record Format and mapping information to schematize a record.
	// See Record Format below for more details.
	// +kubebuilder:validation:Required
	RecordFormat []SchemaRecordFormatObservation `json:"recordFormat,omitempty" tf:"record_format,omitempty"`
}

type ReferenceDataSourcesSchemaParameters struct {

	// The Record Column mapping for the streaming source data element.
	// See Record Columns below for more details.
	// +kubebuilder:validation:Required
	RecordColumns []SchemaRecordColumnsParameters `json:"recordColumns" tf:"record_columns,omitempty"`

	// The Encoding of the record in the streaming source.
	// +kubebuilder:validation:Optional
	RecordEncoding *string `json:"recordEncoding,omitempty" tf:"record_encoding,omitempty"`

	// The Record Format and mapping information to schematize a record.
	// See Record Format below for more details.
	// +kubebuilder:validation:Required
	RecordFormat []SchemaRecordFormatParameters `json:"recordFormat" tf:"record_format,omitempty"`
}

type S3Observation struct {
}

type S3Parameters struct {

	// The S3 Bucket ARN.
	// +kubebuilder:validation:Required
	BucketArn *string `json:"bucketArn" tf:"bucket_arn,omitempty"`

	// The File Key name containing reference data.
	// +kubebuilder:validation:Required
	FileKey *string `json:"fileKey" tf:"file_key,omitempty"`

	// The IAM Role ARN to read the data.
	// +kubebuilder:validation:Required
	RoleArn *string `json:"roleArn" tf:"role_arn,omitempty"`
}

type SchemaObservation struct {

	// The Record Format and mapping information to schematize a record.
	// See Record Format below for more details.
	// +kubebuilder:validation:Required
	RecordFormat []RecordFormatObservation `json:"recordFormat,omitempty" tf:"record_format,omitempty"`
}

type SchemaParameters struct {

	// The Record Column mapping for the streaming source data element.
	// See Record Columns below for more details.
	// +kubebuilder:validation:Required
	RecordColumns []RecordColumnsParameters `json:"recordColumns" tf:"record_columns,omitempty"`

	// The Encoding of the record in the streaming source.
	// +kubebuilder:validation:Optional
	RecordEncoding *string `json:"recordEncoding,omitempty" tf:"record_encoding,omitempty"`

	// The Record Format and mapping information to schematize a record.
	// See Record Format below for more details.
	// +kubebuilder:validation:Required
	RecordFormat []RecordFormatParameters `json:"recordFormat" tf:"record_format,omitempty"`
}

type SchemaRecordColumnsObservation struct {
}

type SchemaRecordColumnsParameters struct {

	// The Mapping reference to the data element.
	// +kubebuilder:validation:Optional
	Mapping *string `json:"mapping,omitempty" tf:"mapping,omitempty"`

	// Name of the column.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The SQL Type of the column.
	// +kubebuilder:validation:Required
	SQLType *string `json:"sqlType" tf:"sql_type,omitempty"`
}

type SchemaRecordFormatObservation struct {

	// The Format Type of the records on the output stream. Can be CSV or JSON.
	RecordFormatType *string `json:"recordFormatType,omitempty" tf:"record_format_type,omitempty"`
}

type SchemaRecordFormatParameters struct {

	// The Mapping Information for the record format.
	// See Mapping Parameters below for more details.
	// +kubebuilder:validation:Optional
	MappingParameters []RecordFormatMappingParametersParameters `json:"mappingParameters,omitempty" tf:"mapping_parameters,omitempty"`
}

type StartingPositionConfigurationObservation struct {
}

type StartingPositionConfigurationParameters struct {

	// The starting position on the stream. Valid values: LAST_STOPPED_POINT, NOW, TRIM_HORIZON.
	// +kubebuilder:validation:Optional
	StartingPosition *string `json:"startingPosition,omitempty" tf:"starting_position,omitempty"`
}

// ApplicationSpec defines the desired state of Application
type ApplicationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApplicationParameters `json:"forProvider"`
}

// ApplicationStatus defines the observed state of Application.
type ApplicationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApplicationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Application is the Schema for the Applications API. Provides a AWS Kinesis Analytics Application
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Application struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApplicationSpec   `json:"spec"`
	Status            ApplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationList contains a list of Applications
type ApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Application `json:"items"`
}

// Repository type metadata.
var (
	Application_Kind             = "Application"
	Application_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Application_Kind}.String()
	Application_KindAPIVersion   = Application_Kind + "." + CRDGroupVersion.String()
	Application_GroupVersionKind = CRDGroupVersion.WithKind(Application_Kind)
)

func init() {
	SchemeBuilder.Register(&Application{}, &ApplicationList{})
}
