package resources


// AWSKinesisFirehoseDeliveryStream_S3DestinationConfiguration AWS CloudFormation Resource (AWS::KinesisFirehose::DeliveryStream.S3DestinationConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration.html
type AWSKinesisFirehoseDeliveryStream_S3DestinationConfiguration struct {
    
    // BucketARN AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-bucketarn
    BucketARN string `json:"BucketARN,omitempty"`
    
    // BufferingHints AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-bufferinghints
    BufferingHints *AWSKinesisFirehoseDeliveryStream_BufferingHints `json:"BufferingHints,omitempty"`
    
    // CloudWatchLoggingOptions AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-cloudwatchloggingoptions
    CloudWatchLoggingOptions *AWSKinesisFirehoseDeliveryStream_CloudWatchLoggingOptions `json:"CloudWatchLoggingOptions,omitempty"`
    
    // CompressionFormat AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-compressionformat
    CompressionFormat string `json:"CompressionFormat,omitempty"`
    
    // EncryptionConfiguration AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-encryptionconfiguration
    EncryptionConfiguration *AWSKinesisFirehoseDeliveryStream_EncryptionConfiguration `json:"EncryptionConfiguration,omitempty"`
    
    // Prefix AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-prefix
    Prefix string `json:"Prefix,omitempty"`
    
    // RoleARN AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-rolearn
    RoleARN string `json:"RoleARN,omitempty"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisFirehoseDeliveryStream_S3DestinationConfiguration) AWSCloudFormationType() string {
    return "AWS::KinesisFirehose::DeliveryStream.S3DestinationConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisFirehoseDeliveryStream_S3DestinationConfiguration) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}
