// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisvideoarchivedmedia

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-video-archived-media-2017-09-30/GetMediaForFragmentListInput
type GetMediaForFragmentListInput struct {
	_ struct{} `type:"structure"`

	// A list of the numbers of fragments for which to retrieve media. You retrieve
	// these values with ListFragments.
	//
	// Fragments is a required field
	Fragments []string `min:"1" type:"list" required:"true"`

	// The name of the stream from which to retrieve fragment media.
	//
	// StreamName is a required field
	StreamName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetMediaForFragmentListInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetMediaForFragmentListInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetMediaForFragmentListInput"}

	if s.Fragments == nil {
		invalidParams.Add(aws.NewErrParamRequired("Fragments"))
	}
	if s.Fragments != nil && len(s.Fragments) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Fragments", 1))
	}

	if s.StreamName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StreamName"))
	}
	if s.StreamName != nil && len(*s.StreamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StreamName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetMediaForFragmentListInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Fragments != nil {
		v := s.Fragments

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Fragments", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.StreamName != nil {
		v := *s.StreamName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "StreamName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-video-archived-media-2017-09-30/GetMediaForFragmentListOutput
type GetMediaForFragmentListOutput struct {
	_ struct{} `type:"structure" payload:"Payload"`

	// The content type of the requested media.
	ContentType *string `location:"header" locationName:"Content-Type" min:"1" type:"string"`

	// The payload that Kinesis Video Streams returns is a sequence of chunks from
	// the specified stream. For information about the chunks, see PutMedia (http://docs.aws.amazon.com/kinesisvideostreams/latest/dg/API_dataplane_PutMedia.html).
	// The chunks that Kinesis Video Streams returns in the GetMediaForFragmentList
	// call also include the following additional Matroska (MKV) tags:
	//
	//    * AWS_KINESISVIDEO_FRAGMENT_NUMBER - Fragment number returned in the chunk.
	//
	//    * AWS_KINESISVIDEO_SERVER_SIDE_TIMESTAMP - Server-side timestamp of the
	//    fragment.
	//
	//    * AWS_KINESISVIDEO_PRODUCER_SIDE_TIMESTAMP - Producer-side timestamp of
	//    the fragment.
	//
	// The following tags will be included if an exception occurs:
	//
	//    * AWS_KINESISVIDEO_FRAGMENT_NUMBER - The number of the fragment that threw
	//    the exception
	//
	//    * AWS_KINESISVIDEO_EXCEPTION_ERROR_CODE - The integer code of the exception
	//
	//    * AWS_KINESISVIDEO_EXCEPTION_MESSAGE - A text description of the exception
	Payload io.ReadCloser `type:"blob"`
}

// String returns the string representation
func (s GetMediaForFragmentListOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetMediaForFragmentListOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ContentType != nil {
		v := *s.ContentType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	// Skipping Payload Output type's body not valid.
	return nil
}

const opGetMediaForFragmentList = "GetMediaForFragmentList"

// GetMediaForFragmentListRequest returns a request value for making API operation for
// Amazon Kinesis Video Streams Archived Media.
//
// Gets media for a list of fragments (specified by fragment number) from the
// archived data in an Amazon Kinesis video stream.
//
// You must first call the GetDataEndpoint API to get an endpoint. Then send
// the GetMediaForFragmentList requests to this endpoint using the --endpoint-url
// parameter (https://docs.aws.amazon.com/cli/latest/reference/).
//
// The following limits apply when using the GetMediaForFragmentList API:
//
//    * A client can call GetMediaForFragmentList up to five times per second
//    per stream.
//
//    * Kinesis Video Streams sends media data at a rate of up to 25 megabytes
//    per second (or 200 megabits per second) during a GetMediaForFragmentList
//    session.
//
//    // Example sending a request using GetMediaForFragmentListRequest.
//    req := client.GetMediaForFragmentListRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-video-archived-media-2017-09-30/GetMediaForFragmentList
func (c *Client) GetMediaForFragmentListRequest(input *GetMediaForFragmentListInput) GetMediaForFragmentListRequest {
	op := &aws.Operation{
		Name:       opGetMediaForFragmentList,
		HTTPMethod: "POST",
		HTTPPath:   "/getMediaForFragmentList",
	}

	if input == nil {
		input = &GetMediaForFragmentListInput{}
	}

	req := c.newRequest(op, input, &GetMediaForFragmentListOutput{})
	return GetMediaForFragmentListRequest{Request: req, Input: input, Copy: c.GetMediaForFragmentListRequest}
}

// GetMediaForFragmentListRequest is the request type for the
// GetMediaForFragmentList API operation.
type GetMediaForFragmentListRequest struct {
	*aws.Request
	Input *GetMediaForFragmentListInput
	Copy  func(*GetMediaForFragmentListInput) GetMediaForFragmentListRequest
}

// Send marshals and sends the GetMediaForFragmentList API request.
func (r GetMediaForFragmentListRequest) Send(ctx context.Context) (*GetMediaForFragmentListResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetMediaForFragmentListResponse{
		GetMediaForFragmentListOutput: r.Request.Data.(*GetMediaForFragmentListOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetMediaForFragmentListResponse is the response type for the
// GetMediaForFragmentList API operation.
type GetMediaForFragmentListResponse struct {
	*GetMediaForFragmentListOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetMediaForFragmentList request.
func (r *GetMediaForFragmentListResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
