// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotanalytics

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/CreateChannelRequest
type CreateChannelInput struct {
	_ struct{} `type:"structure"`

	// The name of the channel.
	//
	// ChannelName is a required field
	ChannelName *string `locationName:"channelName" min:"1" type:"string" required:"true"`

	// How long, in days, message data is kept for the channel.
	RetentionPeriod *RetentionPeriod `locationName:"retentionPeriod" type:"structure"`

	// Metadata which can be used to manage the channel.
	Tags []Tag `locationName:"tags" min:"1" type:"list"`
}

// String returns the string representation
func (s CreateChannelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateChannelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateChannelInput"}

	if s.ChannelName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChannelName"))
	}
	if s.ChannelName != nil && len(*s.ChannelName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChannelName", 1))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}
	if s.RetentionPeriod != nil {
		if err := s.RetentionPeriod.Validate(); err != nil {
			invalidParams.AddNested("RetentionPeriod", err.(aws.ErrInvalidParams))
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateChannelInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.ChannelName != nil {
		v := *s.ChannelName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "channelName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RetentionPeriod != nil {
		v := s.RetentionPeriod

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "retentionPeriod", v, metadata)
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
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/CreateChannelResponse
type CreateChannelOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the channel.
	ChannelArn *string `locationName:"channelArn" type:"string"`

	// The name of the channel.
	ChannelName *string `locationName:"channelName" min:"1" type:"string"`

	// How long, in days, message data is kept for the channel.
	RetentionPeriod *RetentionPeriod `locationName:"retentionPeriod" type:"structure"`
}

// String returns the string representation
func (s CreateChannelOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateChannelOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ChannelArn != nil {
		v := *s.ChannelArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "channelArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ChannelName != nil {
		v := *s.ChannelName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "channelName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RetentionPeriod != nil {
		v := s.RetentionPeriod

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "retentionPeriod", v, metadata)
	}
	return nil
}

const opCreateChannel = "CreateChannel"

// CreateChannelRequest returns a request value for making API operation for
// AWS IoT Analytics.
//
// Creates a channel. A channel collects data from an MQTT topic and archives
// the raw, unprocessed messages before publishing the data to a pipeline.
//
//    // Example sending a request using CreateChannelRequest.
//    req := client.CreateChannelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/CreateChannel
func (c *Client) CreateChannelRequest(input *CreateChannelInput) CreateChannelRequest {
	op := &aws.Operation{
		Name:       opCreateChannel,
		HTTPMethod: "POST",
		HTTPPath:   "/channels",
	}

	if input == nil {
		input = &CreateChannelInput{}
	}

	req := c.newRequest(op, input, &CreateChannelOutput{})
	return CreateChannelRequest{Request: req, Input: input, Copy: c.CreateChannelRequest}
}

// CreateChannelRequest is the request type for the
// CreateChannel API operation.
type CreateChannelRequest struct {
	*aws.Request
	Input *CreateChannelInput
	Copy  func(*CreateChannelInput) CreateChannelRequest
}

// Send marshals and sends the CreateChannel API request.
func (r CreateChannelRequest) Send(ctx context.Context) (*CreateChannelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateChannelResponse{
		CreateChannelOutput: r.Request.Data.(*CreateChannelOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateChannelResponse is the response type for the
// CreateChannel API operation.
type CreateChannelResponse struct {
	*CreateChannelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateChannel request.
func (r *CreateChannelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
