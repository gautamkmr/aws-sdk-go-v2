// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateDynamicThingGroupInput struct {
	_ struct{} `type:"structure"`

	// The dynamic thing group index name.
	//
	// Currently one index is supported: "AWS_Things".
	IndexName *string `locationName:"indexName" min:"1" type:"string"`

	// The dynamic thing group search query string.
	//
	// See Query Syntax (https://docs.aws.amazon.com/iot/latest/developerguide/query-syntax.html)
	// for information about query string syntax.
	//
	// QueryString is a required field
	QueryString *string `locationName:"queryString" min:"1" type:"string" required:"true"`

	// The dynamic thing group query version.
	//
	// Currently one query version is supported: "2017-09-30". If not specified,
	// the query version defaults to this value.
	QueryVersion *string `locationName:"queryVersion" type:"string"`

	// Metadata which can be used to manage the dynamic thing group.
	Tags []Tag `locationName:"tags" type:"list"`

	// The dynamic thing group name to create.
	//
	// ThingGroupName is a required field
	ThingGroupName *string `location:"uri" locationName:"thingGroupName" min:"1" type:"string" required:"true"`

	// The dynamic thing group properties.
	ThingGroupProperties *ThingGroupProperties `locationName:"thingGroupProperties" type:"structure"`
}

// String returns the string representation
func (s CreateDynamicThingGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDynamicThingGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDynamicThingGroupInput"}
	if s.IndexName != nil && len(*s.IndexName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IndexName", 1))
	}

	if s.QueryString == nil {
		invalidParams.Add(aws.NewErrParamRequired("QueryString"))
	}
	if s.QueryString != nil && len(*s.QueryString) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("QueryString", 1))
	}

	if s.ThingGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ThingGroupName"))
	}
	if s.ThingGroupName != nil && len(*s.ThingGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ThingGroupName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDynamicThingGroupInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.IndexName != nil {
		v := *s.IndexName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "indexName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.QueryString != nil {
		v := *s.QueryString

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "queryString", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.QueryVersion != nil {
		v := *s.QueryVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "queryVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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
	if s.ThingGroupProperties != nil {
		v := s.ThingGroupProperties

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "thingGroupProperties", v, metadata)
	}
	if s.ThingGroupName != nil {
		v := *s.ThingGroupName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "thingGroupName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateDynamicThingGroupOutput struct {
	_ struct{} `type:"structure"`

	// The dynamic thing group index name.
	IndexName *string `locationName:"indexName" min:"1" type:"string"`

	// The dynamic thing group search query string.
	QueryString *string `locationName:"queryString" min:"1" type:"string"`

	// The dynamic thing group query version.
	QueryVersion *string `locationName:"queryVersion" type:"string"`

	// The dynamic thing group ARN.
	ThingGroupArn *string `locationName:"thingGroupArn" type:"string"`

	// The dynamic thing group ID.
	ThingGroupId *string `locationName:"thingGroupId" min:"1" type:"string"`

	// The dynamic thing group name.
	ThingGroupName *string `locationName:"thingGroupName" min:"1" type:"string"`
}

// String returns the string representation
func (s CreateDynamicThingGroupOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDynamicThingGroupOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.IndexName != nil {
		v := *s.IndexName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "indexName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.QueryString != nil {
		v := *s.QueryString

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "queryString", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.QueryVersion != nil {
		v := *s.QueryVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "queryVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ThingGroupArn != nil {
		v := *s.ThingGroupArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "thingGroupArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ThingGroupId != nil {
		v := *s.ThingGroupId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "thingGroupId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ThingGroupName != nil {
		v := *s.ThingGroupName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "thingGroupName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateDynamicThingGroup = "CreateDynamicThingGroup"

// CreateDynamicThingGroupRequest returns a request value for making API operation for
// AWS IoT.
//
// Creates a dynamic thing group.
//
//    // Example sending a request using CreateDynamicThingGroupRequest.
//    req := client.CreateDynamicThingGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateDynamicThingGroupRequest(input *CreateDynamicThingGroupInput) CreateDynamicThingGroupRequest {
	op := &aws.Operation{
		Name:       opCreateDynamicThingGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/dynamic-thing-groups/{thingGroupName}",
	}

	if input == nil {
		input = &CreateDynamicThingGroupInput{}
	}

	req := c.newRequest(op, input, &CreateDynamicThingGroupOutput{})
	return CreateDynamicThingGroupRequest{Request: req, Input: input, Copy: c.CreateDynamicThingGroupRequest}
}

// CreateDynamicThingGroupRequest is the request type for the
// CreateDynamicThingGroup API operation.
type CreateDynamicThingGroupRequest struct {
	*aws.Request
	Input *CreateDynamicThingGroupInput
	Copy  func(*CreateDynamicThingGroupInput) CreateDynamicThingGroupRequest
}

// Send marshals and sends the CreateDynamicThingGroup API request.
func (r CreateDynamicThingGroupRequest) Send(ctx context.Context) (*CreateDynamicThingGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDynamicThingGroupResponse{
		CreateDynamicThingGroupOutput: r.Request.Data.(*CreateDynamicThingGroupOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDynamicThingGroupResponse is the response type for the
// CreateDynamicThingGroup API operation.
type CreateDynamicThingGroupResponse struct {
	*CreateDynamicThingGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDynamicThingGroup request.
func (r *CreateDynamicThingGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
