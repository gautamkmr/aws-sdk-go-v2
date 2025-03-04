// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package securityhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/GetInsightsRequest
type GetInsightsInput struct {
	_ struct{} `type:"structure"`

	// The ARNS of the insights that you want to describe.
	InsightArns []string `type:"list"`

	// Indicates the maximum number of items that you want in the response.
	MaxResults *int64 `min:"1" type:"integer"`

	// Paginates results. Set the value of this parameter to NULL on your first
	// call to the GetInsights operation. For subsequent calls to the operation,
	// fill nextToken in the request with the value of nextToken from the previous
	// response to continue listing data.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetInsightsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetInsightsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetInsightsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetInsightsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.InsightArns != nil {
		v := s.InsightArns

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "InsightArns", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
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
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/GetInsightsResponse
type GetInsightsOutput struct {
	_ struct{} `type:"structure"`

	// The insights returned by the operation.
	//
	// Insights is a required field
	Insights []Insight `type:"list" required:"true"`

	// The token that is required for pagination.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetInsightsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetInsightsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Insights != nil {
		v := s.Insights

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Insights", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetInsights = "GetInsights"

// GetInsightsRequest returns a request value for making API operation for
// AWS SecurityHub.
//
// Lists and describes insights that are specified by insight ARNs.
//
//    // Example sending a request using GetInsightsRequest.
//    req := client.GetInsightsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/GetInsights
func (c *Client) GetInsightsRequest(input *GetInsightsInput) GetInsightsRequest {
	op := &aws.Operation{
		Name:       opGetInsights,
		HTTPMethod: "POST",
		HTTPPath:   "/insights/get",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetInsightsInput{}
	}

	req := c.newRequest(op, input, &GetInsightsOutput{})
	return GetInsightsRequest{Request: req, Input: input, Copy: c.GetInsightsRequest}
}

// GetInsightsRequest is the request type for the
// GetInsights API operation.
type GetInsightsRequest struct {
	*aws.Request
	Input *GetInsightsInput
	Copy  func(*GetInsightsInput) GetInsightsRequest
}

// Send marshals and sends the GetInsights API request.
func (r GetInsightsRequest) Send(ctx context.Context) (*GetInsightsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetInsightsResponse{
		GetInsightsOutput: r.Request.Data.(*GetInsightsOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetInsightsRequestPaginator returns a paginator for GetInsights.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetInsightsRequest(input)
//   p := securityhub.NewGetInsightsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetInsightsPaginator(req GetInsightsRequest) GetInsightsPaginator {
	return GetInsightsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetInsightsInput
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

// GetInsightsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetInsightsPaginator struct {
	aws.Pager
}

func (p *GetInsightsPaginator) CurrentPage() *GetInsightsOutput {
	return p.Pager.CurrentPage().(*GetInsightsOutput)
}

// GetInsightsResponse is the response type for the
// GetInsights API operation.
type GetInsightsResponse struct {
	*GetInsightsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetInsights request.
func (r *GetInsightsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
