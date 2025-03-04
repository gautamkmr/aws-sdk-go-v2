// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package medialive

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/medialive-2017-10-14/ListChannelsRequest
type ListChannelsInput struct {
	_ struct{} `type:"structure"`

	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListChannelsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListChannelsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListChannelsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListChannelsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/medialive-2017-10-14/ListChannelsResponse
type ListChannelsOutput struct {
	_ struct{} `type:"structure"`

	Channels []ChannelSummary `locationName:"channels" type:"list"`

	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListChannelsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListChannelsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Channels != nil {
		v := s.Channels

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "channels", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
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

const opListChannels = "ListChannels"

// ListChannelsRequest returns a request value for making API operation for
// AWS Elemental MediaLive.
//
// Produces list of channels that have been created
//
//    // Example sending a request using ListChannelsRequest.
//    req := client.ListChannelsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/medialive-2017-10-14/ListChannels
func (c *Client) ListChannelsRequest(input *ListChannelsInput) ListChannelsRequest {
	op := &aws.Operation{
		Name:       opListChannels,
		HTTPMethod: "GET",
		HTTPPath:   "/prod/channels",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListChannelsInput{}
	}

	req := c.newRequest(op, input, &ListChannelsOutput{})
	return ListChannelsRequest{Request: req, Input: input, Copy: c.ListChannelsRequest}
}

// ListChannelsRequest is the request type for the
// ListChannels API operation.
type ListChannelsRequest struct {
	*aws.Request
	Input *ListChannelsInput
	Copy  func(*ListChannelsInput) ListChannelsRequest
}

// Send marshals and sends the ListChannels API request.
func (r ListChannelsRequest) Send(ctx context.Context) (*ListChannelsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListChannelsResponse{
		ListChannelsOutput: r.Request.Data.(*ListChannelsOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListChannelsRequestPaginator returns a paginator for ListChannels.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListChannelsRequest(input)
//   p := medialive.NewListChannelsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListChannelsPaginator(req ListChannelsRequest) ListChannelsPaginator {
	return ListChannelsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListChannelsInput
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

// ListChannelsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListChannelsPaginator struct {
	aws.Pager
}

func (p *ListChannelsPaginator) CurrentPage() *ListChannelsOutput {
	return p.Pager.CurrentPage().(*ListChannelsOutput)
}

// ListChannelsResponse is the response type for the
// ListChannels API operation.
type ListChannelsResponse struct {
	*ListChannelsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListChannels request.
func (r *ListChannelsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
