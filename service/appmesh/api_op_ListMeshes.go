// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appmesh

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/ListMeshesInput
type ListMeshesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results returned by ListMeshes in paginated output.
	// When you use this parameter, ListMeshes returns only limit results in a single
	// page along with a nextToken response element. You can see the remaining results
	// of the initial request by sending another ListMeshes request with the returned
	// nextToken value. This value can be between 1 and 100. If you don't use this
	// parameter, ListMeshes returns up to 100 results and a nextToken value if
	// applicable.
	Limit *int64 `location:"querystring" locationName:"limit" min:"1" type:"integer"`

	// The nextToken value returned from a previous paginated ListMeshes request
	// where limit was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// nextToken value.
	//
	// This token should be treated as an opaque identifier that is used only to
	// retrieve the next items in a list and not for other programmatic purposes.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListMeshesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListMeshesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListMeshesInput"}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListMeshesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.Limit != nil {
		v := *s.Limit

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "limit", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/ListMeshesOutput
type ListMeshesOutput struct {
	_ struct{} `type:"structure"`

	// The list of existing service meshes.
	//
	// Meshes is a required field
	Meshes []MeshRef `locationName:"meshes" type:"list" required:"true"`

	// The nextToken value to include in a future ListMeshes request. When the results
	// of a ListMeshes request exceed limit, you can use this value to retrieve
	// the next page of results. This value is null when there are no more results
	// to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListMeshesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListMeshesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Meshes != nil {
		v := s.Meshes

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "meshes", metadata)
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

const opListMeshes = "ListMeshes"

// ListMeshesRequest returns a request value for making API operation for
// AWS App Mesh.
//
// Returns a list of existing service meshes.
//
//    // Example sending a request using ListMeshesRequest.
//    req := client.ListMeshesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/ListMeshes
func (c *Client) ListMeshesRequest(input *ListMeshesInput) ListMeshesRequest {
	op := &aws.Operation{
		Name:       opListMeshes,
		HTTPMethod: "GET",
		HTTPPath:   "/v20190125/meshes",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListMeshesInput{}
	}

	req := c.newRequest(op, input, &ListMeshesOutput{})
	return ListMeshesRequest{Request: req, Input: input, Copy: c.ListMeshesRequest}
}

// ListMeshesRequest is the request type for the
// ListMeshes API operation.
type ListMeshesRequest struct {
	*aws.Request
	Input *ListMeshesInput
	Copy  func(*ListMeshesInput) ListMeshesRequest
}

// Send marshals and sends the ListMeshes API request.
func (r ListMeshesRequest) Send(ctx context.Context) (*ListMeshesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListMeshesResponse{
		ListMeshesOutput: r.Request.Data.(*ListMeshesOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListMeshesRequestPaginator returns a paginator for ListMeshes.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListMeshesRequest(input)
//   p := appmesh.NewListMeshesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListMeshesPaginator(req ListMeshesRequest) ListMeshesPaginator {
	return ListMeshesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListMeshesInput
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

// ListMeshesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListMeshesPaginator struct {
	aws.Pager
}

func (p *ListMeshesPaginator) CurrentPage() *ListMeshesOutput {
	return p.Pager.CurrentPage().(*ListMeshesOutput)
}

// ListMeshesResponse is the response type for the
// ListMeshes API operation.
type ListMeshesResponse struct {
	*ListMeshesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListMeshes request.
func (r *ListMeshesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
