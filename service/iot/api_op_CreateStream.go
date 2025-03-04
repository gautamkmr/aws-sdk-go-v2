// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateStreamInput struct {
	_ struct{} `type:"structure"`

	// A description of the stream.
	Description *string `locationName:"description" type:"string"`

	// The files to stream.
	//
	// Files is a required field
	Files []StreamFile `locationName:"files" min:"1" type:"list" required:"true"`

	// An IAM role that allows the IoT service principal assumes to access your
	// S3 files.
	//
	// RoleArn is a required field
	RoleArn *string `locationName:"roleArn" min:"20" type:"string" required:"true"`

	// The stream ID.
	//
	// StreamId is a required field
	StreamId *string `location:"uri" locationName:"streamId" min:"1" type:"string" required:"true"`

	// Metadata which can be used to manage streams.
	Tags []Tag `locationName:"tags" type:"list"`
}

// String returns the string representation
func (s CreateStreamInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateStreamInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateStreamInput"}

	if s.Files == nil {
		invalidParams.Add(aws.NewErrParamRequired("Files"))
	}
	if s.Files != nil && len(s.Files) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Files", 1))
	}

	if s.RoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleArn"))
	}
	if s.RoleArn != nil && len(*s.RoleArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleArn", 20))
	}

	if s.StreamId == nil {
		invalidParams.Add(aws.NewErrParamRequired("StreamId"))
	}
	if s.StreamId != nil && len(*s.StreamId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StreamId", 1))
	}
	if s.Files != nil {
		for i, v := range s.Files {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Files", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateStreamInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Files != nil {
		v := s.Files

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "files", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.RoleArn != nil {
		v := *s.RoleArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "roleArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "tags", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.StreamId != nil {
		v := *s.StreamId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "streamId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateStreamOutput struct {
	_ struct{} `type:"structure"`

	// A description of the stream.
	Description *string `locationName:"description" type:"string"`

	// The stream ARN.
	StreamArn *string `locationName:"streamArn" type:"string"`

	// The stream ID.
	StreamId *string `locationName:"streamId" min:"1" type:"string"`

	// The version of the stream.
	StreamVersion *int64 `locationName:"streamVersion" type:"integer"`
}

// String returns the string representation
func (s CreateStreamOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateStreamOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StreamArn != nil {
		v := *s.StreamArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "streamArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StreamId != nil {
		v := *s.StreamId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "streamId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StreamVersion != nil {
		v := *s.StreamVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "streamVersion", protocol.Int64Value(v), metadata)
	}
	return nil
}

const opCreateStream = "CreateStream"

// CreateStreamRequest returns a request value for making API operation for
// AWS IoT.
//
// Creates a stream for delivering one or more large files in chunks over MQTT.
// A stream transports data bytes in chunks or blocks packaged as MQTT messages
// from a source like S3. You can have one or more files associated with a stream.
// The total size of a file associated with the stream cannot exceed more than
// 2 MB. The stream will be created with version 0. If a stream is created with
// the same streamID as a stream that existed and was deleted within last 90
// days, we will resurrect that old stream by incrementing the version by 1.
//
//    // Example sending a request using CreateStreamRequest.
//    req := client.CreateStreamRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateStreamRequest(input *CreateStreamInput) CreateStreamRequest {
	op := &aws.Operation{
		Name:       opCreateStream,
		HTTPMethod: "POST",
		HTTPPath:   "/streams/{streamId}",
	}

	if input == nil {
		input = &CreateStreamInput{}
	}

	req := c.newRequest(op, input, &CreateStreamOutput{})
	return CreateStreamRequest{Request: req, Input: input, Copy: c.CreateStreamRequest}
}

// CreateStreamRequest is the request type for the
// CreateStream API operation.
type CreateStreamRequest struct {
	*aws.Request
	Input *CreateStreamInput
	Copy  func(*CreateStreamInput) CreateStreamRequest
}

// Send marshals and sends the CreateStream API request.
func (r CreateStreamRequest) Send(ctx context.Context) (*CreateStreamResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateStreamResponse{
		CreateStreamOutput: r.Request.Data.(*CreateStreamOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateStreamResponse is the response type for the
// CreateStream API operation.
type CreateStreamResponse struct {
	*CreateStreamOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateStream request.
func (r *CreateStreamResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
