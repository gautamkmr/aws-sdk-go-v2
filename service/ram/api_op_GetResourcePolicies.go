// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ram

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ram-2018-01-04/GetResourcePoliciesRequest
type GetResourcePoliciesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int64 `locationName:"maxResults" min:"1" type:"integer"`

	// The token for the next page of results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The principal.
	Principal *string `locationName:"principal" type:"string"`

	// The Amazon Resource Names (ARN) of the resources.
	//
	// ResourceArns is a required field
	ResourceArns []string `locationName:"resourceArns" type:"list" required:"true"`
}

// String returns the string representation
func (s GetResourcePoliciesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetResourcePoliciesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetResourcePoliciesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.ResourceArns == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceArns"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetResourcePoliciesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Principal != nil {
		v := *s.Principal

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "principal", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceArns != nil {
		v := s.ResourceArns

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "resourceArns", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ram-2018-01-04/GetResourcePoliciesResponse
type GetResourcePoliciesOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// A key policy document, in JSON format.
	Policies []string `locationName:"policies" type:"list"`
}

// String returns the string representation
func (s GetResourcePoliciesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetResourcePoliciesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Policies != nil {
		v := s.Policies

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "policies", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

const opGetResourcePolicies = "GetResourcePolicies"

// GetResourcePoliciesRequest returns a request value for making API operation for
// AWS Resource Access Manager.
//
// Gets the policies for the specifies resources.
//
//    // Example sending a request using GetResourcePoliciesRequest.
//    req := client.GetResourcePoliciesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ram-2018-01-04/GetResourcePolicies
func (c *Client) GetResourcePoliciesRequest(input *GetResourcePoliciesInput) GetResourcePoliciesRequest {
	op := &aws.Operation{
		Name:       opGetResourcePolicies,
		HTTPMethod: "POST",
		HTTPPath:   "/getresourcepolicies",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetResourcePoliciesInput{}
	}

	req := c.newRequest(op, input, &GetResourcePoliciesOutput{})
	return GetResourcePoliciesRequest{Request: req, Input: input, Copy: c.GetResourcePoliciesRequest}
}

// GetResourcePoliciesRequest is the request type for the
// GetResourcePolicies API operation.
type GetResourcePoliciesRequest struct {
	*aws.Request
	Input *GetResourcePoliciesInput
	Copy  func(*GetResourcePoliciesInput) GetResourcePoliciesRequest
}

// Send marshals and sends the GetResourcePolicies API request.
func (r GetResourcePoliciesRequest) Send(ctx context.Context) (*GetResourcePoliciesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetResourcePoliciesResponse{
		GetResourcePoliciesOutput: r.Request.Data.(*GetResourcePoliciesOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetResourcePoliciesRequestPaginator returns a paginator for GetResourcePolicies.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetResourcePoliciesRequest(input)
//   p := ram.NewGetResourcePoliciesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetResourcePoliciesPaginator(req GetResourcePoliciesRequest) GetResourcePoliciesPaginator {
	return GetResourcePoliciesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetResourcePoliciesInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// GetResourcePoliciesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetResourcePoliciesPaginator struct {
	aws.Pager
}

func (p *GetResourcePoliciesPaginator) CurrentPage() *GetResourcePoliciesOutput {
	return p.Pager.CurrentPage().(*GetResourcePoliciesOutput)
}

// GetResourcePoliciesResponse is the response type for the
// GetResourcePolicies API operation.
type GetResourcePoliciesResponse struct {
	*GetResourcePoliciesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetResourcePolicies request.
func (r *GetResourcePoliciesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
