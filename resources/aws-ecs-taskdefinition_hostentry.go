package resources


// AWSECSTaskDefinition_HostEntry AWS CloudFormation Resource (AWS::ECS::TaskDefinition.HostEntry)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-hostentry.html
type AWSECSTaskDefinition_HostEntry struct {
    
    // Hostname AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-hostentry.html#cfn-ecs-taskdefinition-containerdefinition-hostentry-hostname
    Hostname string `json:"Hostname,omitempty"`
    
    // IpAddress AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-hostentry.html#cfn-ecs-taskdefinition-containerdefinition-hostentry-ipaddress
    IpAddress string `json:"IpAddress,omitempty"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSECSTaskDefinition_HostEntry) AWSCloudFormationType() string {
    return "AWS::ECS::TaskDefinition.HostEntry"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSECSTaskDefinition_HostEntry) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}
