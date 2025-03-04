// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListSecurityProfilesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return at one time.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The token for the next set of results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListSecurityProfilesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListSecurityProfilesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListSecurityProfilesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListSecurityProfilesInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListSecurityProfilesOutput struct {
	_ struct{} `type:"structure"`

	// A token that can be used to retrieve the next set of results, or null if
	// there are no additional results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// A list of security profile identifiers (names and ARNs).
	SecurityProfileIdentifiers []SecurityProfileIdentifier `locationName:"securityProfileIdentifiers" type:"list"`
}

// String returns the string representation
func (s ListSecurityProfilesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListSecurityProfilesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SecurityProfileIdentifiers != nil {
		v := s.SecurityProfileIdentifiers

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "securityProfileIdentifiers", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListSecurityProfiles = "ListSecurityProfiles"

// ListSecurityProfilesRequest returns a request value for making API operation for
// AWS IoT.
//
// Lists the Device Defender security profiles you have created. You can use
// filters to list only those security profiles associated with a thing group
// or only those associated with your account.
//
//    // Example sending a request using ListSecurityProfilesRequest.
//    req := client.ListSecurityProfilesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListSecurityProfilesRequest(input *ListSecurityProfilesInput) ListSecurityProfilesRequest {
	op := &aws.Operation{
		Name:       opListSecurityProfiles,
		HTTPMethod: "GET",
		HTTPPath:   "/security-profiles",
	}

	if input == nil {
		input = &ListSecurityProfilesInput{}
	}

	req := c.newRequest(op, input, &ListSecurityProfilesOutput{})
	return ListSecurityProfilesRequest{Request: req, Input: input, Copy: c.ListSecurityProfilesRequest}
}

// ListSecurityProfilesRequest is the request type for the
// ListSecurityProfiles API operation.
type ListSecurityProfilesRequest struct {
	*aws.Request
	Input *ListSecurityProfilesInput
	Copy  func(*ListSecurityProfilesInput) ListSecurityProfilesRequest
}

// Send marshals and sends the ListSecurityProfiles API request.
func (r ListSecurityProfilesRequest) Send(ctx context.Context) (*ListSecurityProfilesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListSecurityProfilesResponse{
		ListSecurityProfilesOutput: r.Request.Data.(*ListSecurityProfilesOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListSecurityProfilesResponse is the response type for the
// ListSecurityProfiles API operation.
type ListSecurityProfilesResponse struct {
	*ListSecurityProfilesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListSecurityProfiles request.
func (r *ListSecurityProfilesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
