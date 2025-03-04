// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package guardduty

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/ListFiltersRequest
type ListFiltersInput struct {
	_ struct{} `type:"structure"`

	// DetectorId is a required field
	DetectorId *string `location:"uri" locationName:"detectorId" type:"string" required:"true"`

	// You can use this parameter to indicate the maximum number of items that you
	// want in the response.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListFiltersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListFiltersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListFiltersInput"}

	if s.DetectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorId"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListFiltersInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.DetectorId != nil {
		v := *s.DetectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "detectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
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

// ListFilters response object.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/ListFiltersResponse
type ListFiltersOutput struct {
	_ struct{} `type:"structure"`

	// A list of filter names
	FilterNames []string `locationName:"filterNames" type:"list"`

	// You can use this parameter when paginating results. Set the value of this
	// parameter to null on your first call to the list action. For subsequent calls
	// to the action fill nextToken in the request with the value of NextToken from
	// the previous response to continue listing data.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListFiltersOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListFiltersOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.FilterNames != nil {
		v := s.FilterNames

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "filterNames", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListFilters = "ListFilters"

// ListFiltersRequest returns a request value for making API operation for
// Amazon GuardDuty.
//
// Returns a paginated list of the current filters.
//
//    // Example sending a request using ListFiltersRequest.
//    req := client.ListFiltersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/ListFilters
func (c *Client) ListFiltersRequest(input *ListFiltersInput) ListFiltersRequest {
	op := &aws.Operation{
		Name:       opListFilters,
		HTTPMethod: "GET",
		HTTPPath:   "/detector/{detectorId}/filter",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListFiltersInput{}
	}

	req := c.newRequest(op, input, &ListFiltersOutput{})
	return ListFiltersRequest{Request: req, Input: input, Copy: c.ListFiltersRequest}
}

// ListFiltersRequest is the request type for the
// ListFilters API operation.
type ListFiltersRequest struct {
	*aws.Request
	Input *ListFiltersInput
	Copy  func(*ListFiltersInput) ListFiltersRequest
}

// Send marshals and sends the ListFilters API request.
func (r ListFiltersRequest) Send(ctx context.Context) (*ListFiltersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListFiltersResponse{
		ListFiltersOutput: r.Request.Data.(*ListFiltersOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListFiltersRequestPaginator returns a paginator for ListFilters.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListFiltersRequest(input)
//   p := guardduty.NewListFiltersRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListFiltersPaginator(req ListFiltersRequest) ListFiltersPaginator {
	return ListFiltersPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListFiltersInput
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

// ListFiltersPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListFiltersPaginator struct {
	aws.Pager
}

func (p *ListFiltersPaginator) CurrentPage() *ListFiltersOutput {
	return p.Pager.CurrentPage().(*ListFiltersOutput)
}

// ListFiltersResponse is the response type for the
// ListFilters API operation.
type ListFiltersResponse struct {
	*ListFiltersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListFilters request.
func (r *ListFiltersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
