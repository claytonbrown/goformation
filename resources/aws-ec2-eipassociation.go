package resources


import (
	"encoding/json"
	"fmt"
	"errors"
)

// AWSEC2EIPAssociation AWS CloudFormation Resource (AWS::EC2::EIPAssociation)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip-association.html
type AWSEC2EIPAssociation struct {
    
    // AllocationId AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip-association.html#cfn-ec2-eipassociation-allocationid
    AllocationId string `json:"AllocationId,omitempty"`
    
    // EIP AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip-association.html#cfn-ec2-eipassociation-eip
    EIP string `json:"EIP,omitempty"`
    
    // InstanceId AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip-association.html#cfn-ec2-eipassociation-instanceid
    InstanceId string `json:"InstanceId,omitempty"`
    
    // NetworkInterfaceId AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip-association.html#cfn-ec2-eipassociation-networkinterfaceid
    NetworkInterfaceId string `json:"NetworkInterfaceId,omitempty"`
    
    // PrivateIpAddress AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip-association.html#cfn-ec2-eipassociation-PrivateIpAddress
    PrivateIpAddress string `json:"PrivateIpAddress,omitempty"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2EIPAssociation) AWSCloudFormationType() string {
    return "AWS::EC2::EIPAssociation"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2EIPAssociation) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into 
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSEC2EIPAssociation) MarshalJSON() ([]byte, error) {
	type Properties AWSEC2EIPAssociation
	return json.Marshal(&struct{
		Type string
		Properties Properties
	}{
		Type: r.AWSCloudFormationType(),
		Properties: (Properties)(*r),
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *AWSEC2EIPAssociation) UnmarshalJSON(b []byte) error {
	type Properties AWSEC2EIPAssociation
	res := &struct {
		Type string
		Properties *Properties
	}{}
	if err := json.Unmarshal(b, &res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}
	*r = AWSEC2EIPAssociation(*res.Properties)
	return nil	
}

// GetAllAWSEC2EIPAssociationResources retrieves all AWSEC2EIPAssociation items from an AWS CloudFormation template
func (t *CloudFormationTemplate) GetAllAWSEC2EIPAssociationResources () map[string]AWSEC2EIPAssociation {
    results := map[string]AWSEC2EIPAssociation{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSEC2EIPAssociation:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::EC2::EIPAssociation" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSEC2EIPAssociation
						if err := json.Unmarshal(b, &result); err == nil {
							results[name] = result
						}
					}
				}
			}
		}
	}
	return results
}

// GetAWSEC2EIPAssociationWithName retrieves all AWSEC2EIPAssociation items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *CloudFormationTemplate) GetAWSEC2EIPAssociationWithName (name string) (AWSEC2EIPAssociation, error) {
	if untyped, ok := t.Resources[name]; ok {		
		switch resource := untyped.(type) {
		case AWSEC2EIPAssociation:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::EC2::EIPAssociation" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSEC2EIPAssociation
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}	
		}
	}
    return AWSEC2EIPAssociation{}, errors.New("resource not found")
}
