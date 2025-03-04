// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package amplify

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Request structure for update branch request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/UpdateBranchRequest
type UpdateBranchInput struct {
	_ struct{} `type:"structure"`

	// Unique Id for an Amplify App.
	//
	// AppId is a required field
	AppId *string `location:"uri" locationName:"appId" min:"1" type:"string" required:"true"`

	// Basic Authorization credentials for the branch.
	BasicAuthCredentials *string `locationName:"basicAuthCredentials" type:"string"`

	// Name for the branch.
	//
	// BranchName is a required field
	BranchName *string `location:"uri" locationName:"branchName" min:"1" type:"string" required:"true"`

	// BuildSpec for the branch.
	BuildSpec *string `locationName:"buildSpec" min:"1" type:"string"`

	// Description for the branch.
	Description *string `locationName:"description" type:"string"`

	// Enables auto building for the branch.
	EnableAutoBuild *bool `locationName:"enableAutoBuild" type:"boolean"`

	// Enables Basic Auth for the branch.
	EnableBasicAuth *bool `locationName:"enableBasicAuth" type:"boolean"`

	// Enables notifications for the branch.
	EnableNotification *bool `locationName:"enableNotification" type:"boolean"`

	// Environment Variables for the branch.
	EnvironmentVariables map[string]string `locationName:"environmentVariables" type:"map"`

	// Framework for the branch.
	Framework *string `locationName:"framework" type:"string"`

	// Stage for the branch.
	Stage Stage `locationName:"stage" type:"string" enum:"true"`

	// The content TTL for the website in seconds.
	Ttl *string `locationName:"ttl" type:"string"`
}

// String returns the string representation
func (s UpdateBranchInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateBranchInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateBranchInput"}

	if s.AppId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AppId"))
	}
	if s.AppId != nil && len(*s.AppId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AppId", 1))
	}

	if s.BranchName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BranchName"))
	}
	if s.BranchName != nil && len(*s.BranchName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("BranchName", 1))
	}
	if s.BuildSpec != nil && len(*s.BuildSpec) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("BuildSpec", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateBranchInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.BasicAuthCredentials != nil {
		v := *s.BasicAuthCredentials

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "basicAuthCredentials", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BuildSpec != nil {
		v := *s.BuildSpec

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "buildSpec", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EnableAutoBuild != nil {
		v := *s.EnableAutoBuild

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "enableAutoBuild", protocol.BoolValue(v), metadata)
	}
	if s.EnableBasicAuth != nil {
		v := *s.EnableBasicAuth

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "enableBasicAuth", protocol.BoolValue(v), metadata)
	}
	if s.EnableNotification != nil {
		v := *s.EnableNotification

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "enableNotification", protocol.BoolValue(v), metadata)
	}
	if s.EnvironmentVariables != nil {
		v := s.EnvironmentVariables

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "environmentVariables", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.Framework != nil {
		v := *s.Framework

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "framework", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Stage) > 0 {
		v := s.Stage

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "stage", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Ttl != nil {
		v := *s.Ttl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ttl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AppId != nil {
		v := *s.AppId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "appId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BranchName != nil {
		v := *s.BranchName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "branchName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Result structure for update branch request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/UpdateBranchResult
type UpdateBranchOutput struct {
	_ struct{} `type:"structure"`

	// Branch structure for an Amplify App.
	//
	// Branch is a required field
	Branch *Branch `locationName:"branch" type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateBranchOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateBranchOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Branch != nil {
		v := s.Branch

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "branch", v, metadata)
	}
	return nil
}

const opUpdateBranch = "UpdateBranch"

// UpdateBranchRequest returns a request value for making API operation for
// AWS Amplify.
//
// Updates a branch for an Amplify App.
//
//    // Example sending a request using UpdateBranchRequest.
//    req := client.UpdateBranchRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/UpdateBranch
func (c *Client) UpdateBranchRequest(input *UpdateBranchInput) UpdateBranchRequest {
	op := &aws.Operation{
		Name:       opUpdateBranch,
		HTTPMethod: "POST",
		HTTPPath:   "/apps/{appId}/branches/{branchName}",
	}

	if input == nil {
		input = &UpdateBranchInput{}
	}

	req := c.newRequest(op, input, &UpdateBranchOutput{})
	return UpdateBranchRequest{Request: req, Input: input, Copy: c.UpdateBranchRequest}
}

// UpdateBranchRequest is the request type for the
// UpdateBranch API operation.
type UpdateBranchRequest struct {
	*aws.Request
	Input *UpdateBranchInput
	Copy  func(*UpdateBranchInput) UpdateBranchRequest
}

// Send marshals and sends the UpdateBranch API request.
func (r UpdateBranchRequest) Send(ctx context.Context) (*UpdateBranchResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateBranchResponse{
		UpdateBranchOutput: r.Request.Data.(*UpdateBranchOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateBranchResponse is the response type for the
// UpdateBranch API operation.
type UpdateBranchResponse struct {
	*UpdateBranchOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateBranch request.
func (r *UpdateBranchResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
