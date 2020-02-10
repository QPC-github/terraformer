// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Gets information about a specified traffic policy instance.
type GetTrafficPolicyInstanceInput struct {
	_ struct{} `type:"structure"`

	// The ID of the traffic policy instance that you want to get information about.
	//
	// Id is a required field
	Id *string `location:"uri" locationName:"Id" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetTrafficPolicyInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetTrafficPolicyInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetTrafficPolicyInstanceInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetTrafficPolicyInstanceInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Id", protocol.StringValue(v), metadata)
	}
	return nil
}

// A complex type that contains information about the resource record sets that
// Amazon Route 53 created based on a specified traffic policy.
type GetTrafficPolicyInstanceOutput struct {
	_ struct{} `type:"structure"`

	// A complex type that contains settings for the traffic policy instance.
	//
	// TrafficPolicyInstance is a required field
	TrafficPolicyInstance *TrafficPolicyInstance `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetTrafficPolicyInstanceOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetTrafficPolicyInstanceOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.TrafficPolicyInstance != nil {
		v := s.TrafficPolicyInstance

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "TrafficPolicyInstance", v, metadata)
	}
	return nil
}

const opGetTrafficPolicyInstance = "GetTrafficPolicyInstance"

// GetTrafficPolicyInstanceRequest returns a request value for making API operation for
// Amazon Route 53.
//
// Gets information about a specified traffic policy instance.
//
// After you submit a CreateTrafficPolicyInstance or an UpdateTrafficPolicyInstance
// request, there's a brief delay while Amazon Route 53 creates the resource
// record sets that are specified in the traffic policy definition. For more
// information, see the State response element.
//
// In the Route 53 console, traffic policy instances are known as policy records.
//
//    // Example sending a request using GetTrafficPolicyInstanceRequest.
//    req := client.GetTrafficPolicyInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/GetTrafficPolicyInstance
func (c *Client) GetTrafficPolicyInstanceRequest(input *GetTrafficPolicyInstanceInput) GetTrafficPolicyInstanceRequest {
	op := &aws.Operation{
		Name:       opGetTrafficPolicyInstance,
		HTTPMethod: "GET",
		HTTPPath:   "/2013-04-01/trafficpolicyinstance/{Id}",
	}

	if input == nil {
		input = &GetTrafficPolicyInstanceInput{}
	}

	req := c.newRequest(op, input, &GetTrafficPolicyInstanceOutput{})
	return GetTrafficPolicyInstanceRequest{Request: req, Input: input, Copy: c.GetTrafficPolicyInstanceRequest}
}

// GetTrafficPolicyInstanceRequest is the request type for the
// GetTrafficPolicyInstance API operation.
type GetTrafficPolicyInstanceRequest struct {
	*aws.Request
	Input *GetTrafficPolicyInstanceInput
	Copy  func(*GetTrafficPolicyInstanceInput) GetTrafficPolicyInstanceRequest
}

// Send marshals and sends the GetTrafficPolicyInstance API request.
func (r GetTrafficPolicyInstanceRequest) Send(ctx context.Context) (*GetTrafficPolicyInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetTrafficPolicyInstanceResponse{
		GetTrafficPolicyInstanceOutput: r.Request.Data.(*GetTrafficPolicyInstanceOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetTrafficPolicyInstanceResponse is the response type for the
// GetTrafficPolicyInstance API operation.
type GetTrafficPolicyInstanceResponse struct {
	*GetTrafficPolicyInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetTrafficPolicyInstance request.
func (r *GetTrafficPolicyInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}