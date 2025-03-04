// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot1clickdevicesservice

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

var _ aws.Config
var _ = awsutil.Prettify

// Please also see https://docs.aws.amazon.com/goto/WebAPI/devices-2018-05-14/Attributes
type Attributes struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s Attributes) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Attributes) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/devices-2018-05-14/Device
type Device struct {
	_ struct{} `type:"structure"`

	// The user specified attributes associated with the device for an event.
	Attributes *Attributes `locationName:"attributes" type:"structure"`

	// The unique identifier of the device.
	DeviceId *string `locationName:"deviceId" type:"string"`

	// The device type, such as "button".
	Type *string `locationName:"type" type:"string"`
}

// String returns the string representation
func (s Device) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Device) MarshalFields(e protocol.FieldEncoder) error {
	if s.Attributes != nil {
		v := s.Attributes

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "attributes", v, metadata)
	}
	if s.DeviceId != nil {
		v := *s.DeviceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "deviceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Type != nil {
		v := *s.Type

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "type", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/devices-2018-05-14/DeviceDescription
type DeviceDescription struct {
	_ struct{} `type:"structure"`

	// The ARN of the device.
	Arn *string `locationName:"arn" type:"string"`

	// An array of zero or more elements of DeviceAttribute objects providing user
	// specified device attributes.
	Attributes map[string]string `locationName:"attributes" type:"map"`

	// The unique identifier of the device.
	DeviceId *string `locationName:"deviceId" type:"string"`

	// A Boolean value indicating whether or not the device is enabled.
	Enabled *bool `locationName:"enabled" type:"boolean"`

	// A value between 0 and 1 inclusive, representing the fraction of life remaining
	// for the device.
	RemainingLife *float64 `locationName:"remainingLife" type:"double"`

	Tags map[string]string `locationName:"tags" type:"map"`

	// The type of the device, such as "button".
	Type *string `locationName:"type" type:"string"`
}

// String returns the string representation
func (s DeviceDescription) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeviceDescription) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Attributes != nil {
		v := s.Attributes

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "attributes", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.DeviceId != nil {
		v := *s.DeviceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "deviceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Enabled != nil {
		v := *s.Enabled

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "enabled", protocol.BoolValue(v), metadata)
	}
	if s.RemainingLife != nil {
		v := *s.RemainingLife

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "remainingLife", protocol.Float64Value(v), metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.Type != nil {
		v := *s.Type

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "type", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/devices-2018-05-14/DeviceEvent
type DeviceEvent struct {
	_ struct{} `type:"structure"`

	// An object representing the device associated with the event.
	Device *Device `locationName:"device" type:"structure"`

	// A serialized JSON object representing the device-type specific event.
	StdEvent *string `locationName:"stdEvent" type:"string"`
}

// String returns the string representation
func (s DeviceEvent) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeviceEvent) MarshalFields(e protocol.FieldEncoder) error {
	if s.Device != nil {
		v := s.Device

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "device", v, metadata)
	}
	if s.StdEvent != nil {
		v := *s.StdEvent

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "stdEvent", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/devices-2018-05-14/DeviceMethod
type DeviceMethod struct {
	_ struct{} `type:"structure"`

	// The type of the device, such as "button".
	DeviceType *string `locationName:"deviceType" type:"string"`

	// The name of the method applicable to the deviceType.
	MethodName *string `locationName:"methodName" type:"string"`
}

// String returns the string representation
func (s DeviceMethod) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeviceMethod) MarshalFields(e protocol.FieldEncoder) error {
	if s.DeviceType != nil {
		v := *s.DeviceType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "deviceType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MethodName != nil {
		v := *s.MethodName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "methodName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}
