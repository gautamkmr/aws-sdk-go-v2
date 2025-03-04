// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glacier

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Provides options to retrieve the vault list owned by the calling user's account.
// The list provides metadata information for each vault.
type ListVaultsInput struct {
	_ struct{} `type:"structure"`

	// The AccountId value is the AWS account ID. This value must match the AWS
	// account ID associated with the credentials used to sign the request. You
	// can either specify an AWS account ID or optionally a single '-' (hyphen),
	// in which case Amazon Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you specify your account ID, do
	// not include any hyphens ('-') in the ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The maximum number of vaults to be returned. The default limit is 10. The
	// number of vaults returned might be fewer than the specified limit, but the
	// number of returned vaults never exceeds the limit.
	Limit *string `location:"querystring" locationName:"limit" type:"string"`

	// A string used for pagination. The marker specifies the vault ARN after which
	// the listing of vaults should begin.
	Marker *string `location:"querystring" locationName:"marker" type:"string"`
}

// String returns the string representation
func (s ListVaultsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListVaultsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListVaultsInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListVaultsInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Limit != nil {
		v := *s.Limit

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "limit", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Marker != nil {
		v := *s.Marker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "marker", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Contains the Amazon Glacier response to your request.
type ListVaultsOutput struct {
	_ struct{} `type:"structure"`

	// The vault ARN at which to continue pagination of the results. You use the
	// marker in another List Vaults request to obtain more vaults in the list.
	Marker *string `type:"string"`

	// List of vaults.
	VaultList []DescribeVaultOutput `type:"list"`
}

// String returns the string representation
func (s ListVaultsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListVaultsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Marker != nil {
		v := *s.Marker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Marker", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VaultList != nil {
		v := s.VaultList

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "VaultList", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListVaults = "ListVaults"

// ListVaultsRequest returns a request value for making API operation for
// Amazon Glacier.
//
// This operation lists all vaults owned by the calling user's account. The
// list returned in the response is ASCII-sorted by vault name.
//
// By default, this operation returns up to 10 items. If there are more vaults
// to list, the response marker field contains the vault Amazon Resource Name
// (ARN) at which to continue the list with a new List Vaults request; otherwise,
// the marker field is null. To return a list of vaults that begins at a specific
// vault, set the marker request parameter to the vault ARN you obtained from
// a previous List Vaults request. You can also limit the number of vaults returned
// in the response by specifying the limit parameter in the request.
//
// An AWS account has full permission to perform all operations (actions). However,
// AWS Identity and Access Management (IAM) users don't have any permissions
// by default. You must grant them explicit permission to perform specific actions.
// For more information, see Access Control Using AWS Identity and Access Management
// (IAM) (http://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html).
//
// For conceptual information and underlying REST API, see Retrieving Vault
// Metadata in Amazon Glacier (http://docs.aws.amazon.com/amazonglacier/latest/dev/retrieving-vault-info.html)
// and List Vaults (http://docs.aws.amazon.com/amazonglacier/latest/dev/api-vaults-get.html)
// in the Amazon Glacier Developer Guide.
//
//    // Example sending a request using ListVaultsRequest.
//    req := client.ListVaultsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListVaultsRequest(input *ListVaultsInput) ListVaultsRequest {
	op := &aws.Operation{
		Name:       opListVaults,
		HTTPMethod: "GET",
		HTTPPath:   "/{accountId}/vaults",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListVaultsInput{}
	}

	req := c.newRequest(op, input, &ListVaultsOutput{})
	return ListVaultsRequest{Request: req, Input: input, Copy: c.ListVaultsRequest}
}

// ListVaultsRequest is the request type for the
// ListVaults API operation.
type ListVaultsRequest struct {
	*aws.Request
	Input *ListVaultsInput
	Copy  func(*ListVaultsInput) ListVaultsRequest
}

// Send marshals and sends the ListVaults API request.
func (r ListVaultsRequest) Send(ctx context.Context) (*ListVaultsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListVaultsResponse{
		ListVaultsOutput: r.Request.Data.(*ListVaultsOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListVaultsRequestPaginator returns a paginator for ListVaults.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListVaultsRequest(input)
//   p := glacier.NewListVaultsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListVaultsPaginator(req ListVaultsRequest) ListVaultsPaginator {
	return ListVaultsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListVaultsInput
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

// ListVaultsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListVaultsPaginator struct {
	aws.Pager
}

func (p *ListVaultsPaginator) CurrentPage() *ListVaultsOutput {
	return p.Pager.CurrentPage().(*ListVaultsOutput)
}

// ListVaultsResponse is the response type for the
// ListVaults API operation.
type ListVaultsResponse struct {
	*ListVaultsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListVaults request.
func (r *ListVaultsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
