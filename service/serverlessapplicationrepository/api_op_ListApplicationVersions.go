// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package serverlessapplicationrepository

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/serverlessrepo-2017-09-08/ListApplicationVersionsRequest
type ListApplicationVersionsInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"applicationId" type:"string" required:"true"`

	MaxItems *int64 `location:"querystring" locationName:"maxItems" min:"1" type:"integer"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListApplicationVersionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListApplicationVersionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListApplicationVersionsInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}
	if s.MaxItems != nil && *s.MaxItems < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxItems", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListApplicationVersionsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "applicationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxItems != nil {
		v := *s.MaxItems

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxItems", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/serverlessrepo-2017-09-08/ListApplicationVersionsResponse
type ListApplicationVersionsOutput struct {
	_ struct{} `type:"structure"`

	NextToken *string `locationName:"nextToken" type:"string"`

	Versions []VersionSummary `locationName:"versions" type:"list"`
}

// String returns the string representation
func (s ListApplicationVersionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListApplicationVersionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Versions != nil {
		v := s.Versions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "versions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListApplicationVersions = "ListApplicationVersions"

// ListApplicationVersionsRequest returns a request value for making API operation for
// AWSServerlessApplicationRepository.
//
// Lists versions for the specified application.
//
//    // Example sending a request using ListApplicationVersionsRequest.
//    req := client.ListApplicationVersionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/serverlessrepo-2017-09-08/ListApplicationVersions
func (c *Client) ListApplicationVersionsRequest(input *ListApplicationVersionsInput) ListApplicationVersionsRequest {
	op := &aws.Operation{
		Name:       opListApplicationVersions,
		HTTPMethod: "GET",
		HTTPPath:   "/applications/{applicationId}/versions",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxItems",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListApplicationVersionsInput{}
	}

	req := c.newRequest(op, input, &ListApplicationVersionsOutput{})
	return ListApplicationVersionsRequest{Request: req, Input: input, Copy: c.ListApplicationVersionsRequest}
}

// ListApplicationVersionsRequest is the request type for the
// ListApplicationVersions API operation.
type ListApplicationVersionsRequest struct {
	*aws.Request
	Input *ListApplicationVersionsInput
	Copy  func(*ListApplicationVersionsInput) ListApplicationVersionsRequest
}

// Send marshals and sends the ListApplicationVersions API request.
func (r ListApplicationVersionsRequest) Send(ctx context.Context) (*ListApplicationVersionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListApplicationVersionsResponse{
		ListApplicationVersionsOutput: r.Request.Data.(*ListApplicationVersionsOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListApplicationVersionsRequestPaginator returns a paginator for ListApplicationVersions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListApplicationVersionsRequest(input)
//   p := serverlessapplicationrepository.NewListApplicationVersionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListApplicationVersionsPaginator(req ListApplicationVersionsRequest) ListApplicationVersionsPaginator {
	return ListApplicationVersionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListApplicationVersionsInput
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

// ListApplicationVersionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListApplicationVersionsPaginator struct {
	aws.Pager
}

func (p *ListApplicationVersionsPaginator) CurrentPage() *ListApplicationVersionsOutput {
	return p.Pager.CurrentPage().(*ListApplicationVersionsOutput)
}

// ListApplicationVersionsResponse is the response type for the
// ListApplicationVersions API operation.
type ListApplicationVersionsResponse struct {
	*ListApplicationVersionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListApplicationVersions request.
func (r *ListApplicationVersionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
