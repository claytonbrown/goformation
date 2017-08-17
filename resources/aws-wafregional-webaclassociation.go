package resources


import (
	"encoding/json"
	"fmt"
	"errors"
)

// AWSWAFRegionalWebACLAssociation AWS CloudFormation Resource (AWS::WAFRegional::WebACLAssociation)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-webaclassociation.html
type AWSWAFRegionalWebACLAssociation struct {
    
    // ResourceArn AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-webaclassociation.html#cfn-wafregional-webaclassociation-resourcearn
    ResourceArn string `json:"ResourceArn,omitempty"`
    
    // WebACLId AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-webaclassociation.html#cfn-wafregional-webaclassociation-webaclid
    WebACLId string `json:"WebACLId,omitempty"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFRegionalWebACLAssociation) AWSCloudFormationType() string {
    return "AWS::WAFRegional::WebACLAssociation"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFRegionalWebACLAssociation) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into 
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSWAFRegionalWebACLAssociation) MarshalJSON() ([]byte, error) {
	type Properties AWSWAFRegionalWebACLAssociation
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
func (r *AWSWAFRegionalWebACLAssociation) UnmarshalJSON(b []byte) error {
	type Properties AWSWAFRegionalWebACLAssociation
	res := &struct {
		Type string
		Properties *Properties
	}{}
	if err := json.Unmarshal(b, &res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}
	*r = AWSWAFRegionalWebACLAssociation(*res.Properties)
	return nil	
}

// GetAllAWSWAFRegionalWebACLAssociationResources retrieves all AWSWAFRegionalWebACLAssociation items from an AWS CloudFormation template
func (t *CloudFormationTemplate) GetAllAWSWAFRegionalWebACLAssociationResources () map[string]AWSWAFRegionalWebACLAssociation {
    results := map[string]AWSWAFRegionalWebACLAssociation{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSWAFRegionalWebACLAssociation:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::WAFRegional::WebACLAssociation" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSWAFRegionalWebACLAssociation
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

// GetAWSWAFRegionalWebACLAssociationWithName retrieves all AWSWAFRegionalWebACLAssociation items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *CloudFormationTemplate) GetAWSWAFRegionalWebACLAssociationWithName (name string) (AWSWAFRegionalWebACLAssociation, error) {
	if untyped, ok := t.Resources[name]; ok {		
		switch resource := untyped.(type) {
		case AWSWAFRegionalWebACLAssociation:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::WAFRegional::WebACLAssociation" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSWAFRegionalWebACLAssociation
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}	
		}
	}
    return AWSWAFRegionalWebACLAssociation{}, errors.New("resource not found")
}
