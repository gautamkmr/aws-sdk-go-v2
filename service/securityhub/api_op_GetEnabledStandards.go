// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package securityhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/GetEnabledStandardsRequest
type GetEnabledStandardsInput struct {
	_ struct{} `type:"structure"`

	// Indicates the maximum number of items that you want in the response.
	MaxResults *int64 `min:"1" type:"integer"`

	// Paginates results. Set the value of this parameter to NULL on your first
	// call to the GetEnabledStandards operation. For subsequent calls to the operation,
	// fill nextToken in the request with the value of nextToken from the previous
	// response to continue listing data.
	NextToken *string `type:"string"`

	// The list of standards subscription ARNS that you want to list and describe.
	StandardsSubscriptionArns []string `min:"1" type:"list"`
}

// String returns the string representation
func (s GetEnabledStandardsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetEnabledStandardsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetEnabledStandardsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.StandardsSubscriptionArns != nil && len(s.StandardsSubscriptionArns) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StandardsSubscriptionArns", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetEnabledStandardsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MaxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StandardsSubscriptionArns != nil {
		v := s.StandardsSubscriptionArns

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "StandardsSubscriptionArns", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/GetEnabledStandardsResponse
type GetEnabledStandardsOutput struct {
	_ struct{} `type:"structure"`

	// The token that is required for pagination.
	NextToken *string `type:"string"`

	// The standards subscription details returned by the operation.
	StandardsSubscriptions []StandardsSubscription `type:"list"`
}

// String returns the string representation
func (s GetEnabledStandardsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetEnabledStandardsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StandardsSubscriptions != nil {
		v := s.StandardsSubscriptions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "StandardsSubscriptions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opGetEnabledStandards = "GetEnabledStandards"

// GetEnabledStandardsRequest returns a request value for making API operation for
// AWS SecurityHub.
//
// Lists and describes enabled standards.
//
//    // Example sending a request using GetEnabledStandardsRequest.
//    req := client.GetEnabledStandardsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/GetEnabledStandards
func (c *Client) GetEnabledStandardsRequest(input *GetEnabledStandardsInput) GetEnabledStandardsRequest {
	op := &aws.Operation{
		Name:       opGetEnabledStandards,
		HTTPMethod: "POST",
		HTTPPath:   "/standards/get",
	}

	if input == nil {
		input = &GetEnabledStandardsInput{}
	}

	req := c.newRequest(op, input, &GetEnabledStandardsOutput{})
	return GetEnabledStandardsRequest{Request: req, Input: input, Copy: c.GetEnabledStandardsRequest}
}

// GetEnabledStandardsRequest is the request type for the
// GetEnabledStandards API operation.
type GetEnabledStandardsRequest struct {
	*aws.Request
	Input *GetEnabledStandardsInput
	Copy  func(*GetEnabledStandardsInput) GetEnabledStandardsRequest
}

// Send marshals and sends the GetEnabledStandards API request.
func (r GetEnabledStandardsRequest) Send(ctx context.Context) (*GetEnabledStandardsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetEnabledStandardsResponse{
		GetEnabledStandardsOutput: r.Request.Data.(*GetEnabledStandardsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetEnabledStandardsResponse is the response type for the
// GetEnabledStandards API operation.
type GetEnabledStandardsResponse struct {
	*GetEnabledStandardsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetEnabledStandards request.
func (r *GetEnabledStandardsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
