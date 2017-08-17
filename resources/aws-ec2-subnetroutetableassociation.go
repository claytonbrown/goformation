package resources


import (
	"encoding/json"
	"fmt"
	"errors"
)

// AWSEC2SubnetRouteTableAssociation AWS CloudFormation Resource (AWS::EC2::SubnetRouteTableAssociation)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet-route-table-assoc.html
type AWSEC2SubnetRouteTableAssociation struct {
    
    // RouteTableId AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet-route-table-assoc.html#cfn-ec2-subnetroutetableassociation-routetableid
    RouteTableId string `json:"RouteTableId,omitempty"`
    
    // SubnetId AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet-route-table-assoc.html#cfn-ec2-subnetroutetableassociation-subnetid
    SubnetId string `json:"SubnetId,omitempty"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2SubnetRouteTableAssociation) AWSCloudFormationType() string {
    return "AWS::EC2::SubnetRouteTableAssociation"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2SubnetRouteTableAssociation) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into 
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSEC2SubnetRouteTableAssociation) MarshalJSON() ([]byte, error) {
	type Properties AWSEC2SubnetRouteTableAssociation
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
func (r *AWSEC2SubnetRouteTableAssociation) UnmarshalJSON(b []byte) error {
	type Properties AWSEC2SubnetRouteTableAssociation
	res := &struct {
		Type string
		Properties *Properties
	}{}
	if err := json.Unmarshal(b, &res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}
	*r = AWSEC2SubnetRouteTableAssociation(*res.Properties)
	return nil	
}

// GetAllAWSEC2SubnetRouteTableAssociationResources retrieves all AWSEC2SubnetRouteTableAssociation items from an AWS CloudFormation template
func (t *CloudFormationTemplate) GetAllAWSEC2SubnetRouteTableAssociationResources () map[string]AWSEC2SubnetRouteTableAssociation {
    results := map[string]AWSEC2SubnetRouteTableAssociation{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSEC2SubnetRouteTableAssociation:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::EC2::SubnetRouteTableAssociation" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSEC2SubnetRouteTableAssociation
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

// GetAWSEC2SubnetRouteTableAssociationWithName retrieves all AWSEC2SubnetRouteTableAssociation items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *CloudFormationTemplate) GetAWSEC2SubnetRouteTableAssociationWithName (name string) (AWSEC2SubnetRouteTableAssociation, error) {
	if untyped, ok := t.Resources[name]; ok {		
		switch resource := untyped.(type) {
		case AWSEC2SubnetRouteTableAssociation:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::EC2::SubnetRouteTableAssociation" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSEC2SubnetRouteTableAssociation
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}	
		}
	}
    return AWSEC2SubnetRouteTableAssociation{}, errors.New("resource not found")
}
