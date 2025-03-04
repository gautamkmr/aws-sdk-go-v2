// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateJobInput struct {
	_ struct{} `type:"structure"`

	// Allows you to create criteria to abort a job.
	AbortConfig *AbortConfig `locationName:"abortConfig" type:"structure"`

	// A short text description of the job.
	Description *string `locationName:"description" type:"string"`

	// The job document.
	//
	// If the job document resides in an S3 bucket, you must use a placeholder link
	// when specifying the document.
	//
	// The placeholder link is of the following form:
	//
	// ${aws:iot:s3-presigned-url:https://s3.amazonaws.com/bucket/key}
	//
	// where bucket is your bucket name and key is the object in the bucket to which
	// you are linking.
	Document *string `locationName:"document" type:"string"`

	// An S3 link to the job document.
	DocumentSource *string `locationName:"documentSource" min:"1" type:"string"`

	// Allows you to create a staged rollout of the job.
	JobExecutionsRolloutConfig *JobExecutionsRolloutConfig `locationName:"jobExecutionsRolloutConfig" type:"structure"`

	// A job identifier which must be unique for your AWS account. We recommend
	// using a UUID. Alpha-numeric characters, "-" and "_" are valid for use here.
	//
	// JobId is a required field
	JobId *string `location:"uri" locationName:"jobId" min:"1" type:"string" required:"true"`

	// Configuration information for pre-signed S3 URLs.
	PresignedUrlConfig *PresignedUrlConfig `locationName:"presignedUrlConfig" type:"structure"`

	// Metadata which can be used to manage the job.
	Tags []Tag `locationName:"tags" type:"list"`

	// Specifies whether the job will continue to run (CONTINUOUS), or will be complete
	// after all those things specified as targets have completed the job (SNAPSHOT).
	// If continuous, the job may also be run on a thing when a change is detected
	// in a target. For example, a job will run on a thing when the thing is added
	// to a target group, even after the job was completed by all things originally
	// in the group.
	TargetSelection TargetSelection `locationName:"targetSelection" type:"string" enum:"true"`

	// A list of things and thing groups to which the job should be sent.
	//
	// Targets is a required field
	Targets []string `locationName:"targets" min:"1" type:"list" required:"true"`

	// Specifies the amount of time each device has to finish its execution of the
	// job. The timer is started when the job execution status is set to IN_PROGRESS.
	// If the job execution status is not set to another terminal state before the
	// time expires, it will be automatically set to TIMED_OUT.
	TimeoutConfig *TimeoutConfig `locationName:"timeoutConfig" type:"structure"`
}

// String returns the string representation
func (s CreateJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateJobInput"}
	if s.DocumentSource != nil && len(*s.DocumentSource) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DocumentSource", 1))
	}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}
	if s.JobId != nil && len(*s.JobId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobId", 1))
	}

	if s.Targets == nil {
		invalidParams.Add(aws.NewErrParamRequired("Targets"))
	}
	if s.Targets != nil && len(s.Targets) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Targets", 1))
	}
	if s.AbortConfig != nil {
		if err := s.AbortConfig.Validate(); err != nil {
			invalidParams.AddNested("AbortConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.JobExecutionsRolloutConfig != nil {
		if err := s.JobExecutionsRolloutConfig.Validate(); err != nil {
			invalidParams.AddNested("JobExecutionsRolloutConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.PresignedUrlConfig != nil {
		if err := s.PresignedUrlConfig.Validate(); err != nil {
			invalidParams.AddNested("PresignedUrlConfig", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateJobInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.AbortConfig != nil {
		v := s.AbortConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "abortConfig", v, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Document != nil {
		v := *s.Document

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "document", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DocumentSource != nil {
		v := *s.DocumentSource

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "documentSource", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.JobExecutionsRolloutConfig != nil {
		v := s.JobExecutionsRolloutConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "jobExecutionsRolloutConfig", v, metadata)
	}
	if s.PresignedUrlConfig != nil {
		v := s.PresignedUrlConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "presignedUrlConfig", v, metadata)
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
	if len(s.TargetSelection) > 0 {
		v := s.TargetSelection

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "targetSelection", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Targets != nil {
		v := s.Targets

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "targets", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.TimeoutConfig != nil {
		v := s.TimeoutConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "timeoutConfig", v, metadata)
	}
	if s.JobId != nil {
		v := *s.JobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "jobId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateJobOutput struct {
	_ struct{} `type:"structure"`

	// The job description.
	Description *string `locationName:"description" type:"string"`

	// The job ARN.
	JobArn *string `locationName:"jobArn" type:"string"`

	// The unique identifier you assigned to this job.
	JobId *string `locationName:"jobId" min:"1" type:"string"`
}

// String returns the string representation
func (s CreateJobOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateJobOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.JobArn != nil {
		v := *s.JobArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "jobArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.JobId != nil {
		v := *s.JobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "jobId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateJob = "CreateJob"

// CreateJobRequest returns a request value for making API operation for
// AWS IoT.
//
// Creates a job.
//
//    // Example sending a request using CreateJobRequest.
//    req := client.CreateJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateJobRequest(input *CreateJobInput) CreateJobRequest {
	op := &aws.Operation{
		Name:       opCreateJob,
		HTTPMethod: "PUT",
		HTTPPath:   "/jobs/{jobId}",
	}

	if input == nil {
		input = &CreateJobInput{}
	}

	req := c.newRequest(op, input, &CreateJobOutput{})
	return CreateJobRequest{Request: req, Input: input, Copy: c.CreateJobRequest}
}

// CreateJobRequest is the request type for the
// CreateJob API operation.
type CreateJobRequest struct {
	*aws.Request
	Input *CreateJobInput
	Copy  func(*CreateJobInput) CreateJobRequest
}

// Send marshals and sends the CreateJob API request.
func (r CreateJobRequest) Send(ctx context.Context) (*CreateJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateJobResponse{
		CreateJobOutput: r.Request.Data.(*CreateJobOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateJobResponse is the response type for the
// CreateJob API operation.
type CreateJobResponse struct {
	*CreateJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateJob request.
func (r *CreateJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
