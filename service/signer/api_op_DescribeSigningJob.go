// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package signer

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/signer-2017-08-25/DescribeSigningJobRequest
type DescribeSigningJobInput struct {
	_ struct{} `type:"structure"`

	// The ID of the signing job on input.
	//
	// JobId is a required field
	JobId *string `location:"uri" locationName:"jobId" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeSigningJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeSigningJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeSigningJobInput"}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeSigningJobInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.JobId != nil {
		v := *s.JobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "jobId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/signer-2017-08-25/DescribeSigningJobResponse
type DescribeSigningJobOutput struct {
	_ struct{} `type:"structure"`

	// Date and time that the signing job was completed.
	CompletedAt *time.Time `locationName:"completedAt" type:"timestamp" timestampFormat:"unix"`

	// Date and time that the signing job was created.
	CreatedAt *time.Time `locationName:"createdAt" type:"timestamp" timestampFormat:"unix"`

	// The ID of the signing job on output.
	JobId *string `locationName:"jobId" type:"string"`

	// A list of any overrides that were applied to the signing operation.
	Overrides *SigningPlatformOverrides `locationName:"overrides" type:"structure"`

	// The microcontroller platform to which your signed code image will be distributed.
	PlatformId *string `locationName:"platformId" type:"string"`

	// The name of the profile that initiated the signing operation.
	ProfileName *string `locationName:"profileName" min:"2" type:"string"`

	// The IAM principal that requested the signing job.
	RequestedBy *string `locationName:"requestedBy" type:"string"`

	// Name of the S3 bucket where the signed code image is saved by AWS Signer.
	SignedObject *SignedObject `locationName:"signedObject" type:"structure"`

	// Amazon Resource Name (ARN) of your code signing certificate.
	SigningMaterial *SigningMaterial `locationName:"signingMaterial" type:"structure"`

	// Map of user-assigned key-value pairs used during signing. These values contain
	// any information that you specified for use in your signing job.
	SigningParameters map[string]string `locationName:"signingParameters" type:"map"`

	// The object that contains the name of your S3 bucket or your raw code.
	Source *Source `locationName:"source" type:"structure"`

	// Status of the signing job.
	Status SigningStatus `locationName:"status" type:"string" enum:"true"`

	// String value that contains the status reason.
	StatusReason *string `locationName:"statusReason" type:"string"`
}

// String returns the string representation
func (s DescribeSigningJobOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeSigningJobOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CompletedAt != nil {
		v := *s.CompletedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "completedAt", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.CreatedAt != nil {
		v := *s.CreatedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "createdAt", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.JobId != nil {
		v := *s.JobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "jobId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Overrides != nil {
		v := s.Overrides

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "overrides", v, metadata)
	}
	if s.PlatformId != nil {
		v := *s.PlatformId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "platformId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ProfileName != nil {
		v := *s.ProfileName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "profileName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestedBy != nil {
		v := *s.RequestedBy

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "requestedBy", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SignedObject != nil {
		v := s.SignedObject

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "signedObject", v, metadata)
	}
	if s.SigningMaterial != nil {
		v := s.SigningMaterial

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "signingMaterial", v, metadata)
	}
	if s.SigningParameters != nil {
		v := s.SigningParameters

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "signingParameters", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.Source != nil {
		v := s.Source

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "source", v, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.StatusReason != nil {
		v := *s.StatusReason

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "statusReason", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDescribeSigningJob = "DescribeSigningJob"

// DescribeSigningJobRequest returns a request value for making API operation for
// AWS Signer.
//
// Returns information about a specific code signing job. You specify the job
// by using the jobId value that is returned by the StartSigningJob operation.
//
//    // Example sending a request using DescribeSigningJobRequest.
//    req := client.DescribeSigningJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/signer-2017-08-25/DescribeSigningJob
func (c *Client) DescribeSigningJobRequest(input *DescribeSigningJobInput) DescribeSigningJobRequest {
	op := &aws.Operation{
		Name:       opDescribeSigningJob,
		HTTPMethod: "GET",
		HTTPPath:   "/signing-jobs/{jobId}",
	}

	if input == nil {
		input = &DescribeSigningJobInput{}
	}

	req := c.newRequest(op, input, &DescribeSigningJobOutput{})
	return DescribeSigningJobRequest{Request: req, Input: input, Copy: c.DescribeSigningJobRequest}
}

// DescribeSigningJobRequest is the request type for the
// DescribeSigningJob API operation.
type DescribeSigningJobRequest struct {
	*aws.Request
	Input *DescribeSigningJobInput
	Copy  func(*DescribeSigningJobInput) DescribeSigningJobRequest
}

// Send marshals and sends the DescribeSigningJob API request.
func (r DescribeSigningJobRequest) Send(ctx context.Context) (*DescribeSigningJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeSigningJobResponse{
		DescribeSigningJobOutput: r.Request.Data.(*DescribeSigningJobOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeSigningJobResponse is the response type for the
// DescribeSigningJob API operation.
type DescribeSigningJobResponse struct {
	*DescribeSigningJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeSigningJob request.
func (r *DescribeSigningJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
