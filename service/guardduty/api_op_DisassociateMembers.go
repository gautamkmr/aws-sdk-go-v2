// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package guardduty

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// DisassociateMembers request body.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/DisassociateMembersRequest
type DisassociateMembersInput struct {
	_ struct{} `type:"structure"`

	// A list of account IDs of the GuardDuty member accounts that you want to disassociate
	// from master.
	//
	// AccountIds is a required field
	AccountIds []string `locationName:"accountIds" type:"list" required:"true"`

	// DetectorId is a required field
	DetectorId *string `location:"uri" locationName:"detectorId" type:"string" required:"true"`
}

// String returns the string representation
func (s DisassociateMembersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisassociateMembersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisassociateMembersInput"}

	if s.AccountIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountIds"))
	}

	if s.DetectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DisassociateMembersInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.AccountIds != nil {
		v := s.AccountIds

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "accountIds", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.DetectorId != nil {
		v := *s.DetectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "detectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// DisassociateMembers response object.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/DisassociateMembersResponse
type DisassociateMembersOutput struct {
	_ struct{} `type:"structure"`

	// A list of objects containing the unprocessed account and a result string
	// explaining why it was unprocessed.
	UnprocessedAccounts []UnprocessedAccount `locationName:"unprocessedAccounts" type:"list"`
}

// String returns the string representation
func (s DisassociateMembersOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DisassociateMembersOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.UnprocessedAccounts != nil {
		v := s.UnprocessedAccounts

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "unprocessedAccounts", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opDisassociateMembers = "DisassociateMembers"

// DisassociateMembersRequest returns a request value for making API operation for
// Amazon GuardDuty.
//
// Disassociates GuardDuty member accounts (to the current GuardDuty master
// account) specified by the account IDs.
//
//    // Example sending a request using DisassociateMembersRequest.
//    req := client.DisassociateMembersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/DisassociateMembers
func (c *Client) DisassociateMembersRequest(input *DisassociateMembersInput) DisassociateMembersRequest {
	op := &aws.Operation{
		Name:       opDisassociateMembers,
		HTTPMethod: "POST",
		HTTPPath:   "/detector/{detectorId}/member/disassociate",
	}

	if input == nil {
		input = &DisassociateMembersInput{}
	}

	req := c.newRequest(op, input, &DisassociateMembersOutput{})
	return DisassociateMembersRequest{Request: req, Input: input, Copy: c.DisassociateMembersRequest}
}

// DisassociateMembersRequest is the request type for the
// DisassociateMembers API operation.
type DisassociateMembersRequest struct {
	*aws.Request
	Input *DisassociateMembersInput
	Copy  func(*DisassociateMembersInput) DisassociateMembersRequest
}

// Send marshals and sends the DisassociateMembers API request.
func (r DisassociateMembersRequest) Send(ctx context.Context) (*DisassociateMembersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociateMembersResponse{
		DisassociateMembersOutput: r.Request.Data.(*DisassociateMembersOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociateMembersResponse is the response type for the
// DisassociateMembers API operation.
type DisassociateMembersResponse struct {
	*DisassociateMembersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociateMembers request.
func (r *DisassociateMembersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
