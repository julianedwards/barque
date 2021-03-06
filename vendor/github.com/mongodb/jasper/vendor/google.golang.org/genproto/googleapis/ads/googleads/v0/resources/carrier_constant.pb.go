// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/resources/carrier_constant.proto

package resources // import "google.golang.org/genproto/googleapis/ads/googleads/v0/resources"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A carrier criterion that can be used in campaign targeting.
type CarrierConstant struct {
	// The resource name of the carrier criterion.
	// Carrier criterion resource names have the form:
	//
	// `carrierConstants/{criterion_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the carrier criterion.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// The full name of the carrier in English.
	Name *wrappers.StringValue `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// The country code of the country where the carrier is located, e.g., "AR",
	// "FR", etc.
	CountryCode          *wrappers.StringValue `protobuf:"bytes,4,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CarrierConstant) Reset()         { *m = CarrierConstant{} }
func (m *CarrierConstant) String() string { return proto.CompactTextString(m) }
func (*CarrierConstant) ProtoMessage()    {}
func (*CarrierConstant) Descriptor() ([]byte, []int) {
	return fileDescriptor_carrier_constant_f3e35162f1ca864e, []int{0}
}
func (m *CarrierConstant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CarrierConstant.Unmarshal(m, b)
}
func (m *CarrierConstant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CarrierConstant.Marshal(b, m, deterministic)
}
func (dst *CarrierConstant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CarrierConstant.Merge(dst, src)
}
func (m *CarrierConstant) XXX_Size() int {
	return xxx_messageInfo_CarrierConstant.Size(m)
}
func (m *CarrierConstant) XXX_DiscardUnknown() {
	xxx_messageInfo_CarrierConstant.DiscardUnknown(m)
}

var xxx_messageInfo_CarrierConstant proto.InternalMessageInfo

func (m *CarrierConstant) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CarrierConstant) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *CarrierConstant) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *CarrierConstant) GetCountryCode() *wrappers.StringValue {
	if m != nil {
		return m.CountryCode
	}
	return nil
}

func init() {
	proto.RegisterType((*CarrierConstant)(nil), "google.ads.googleads.v0.resources.CarrierConstant")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/resources/carrier_constant.proto", fileDescriptor_carrier_constant_f3e35162f1ca864e)
}

var fileDescriptor_carrier_constant_f3e35162f1ca864e = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4a, 0xfb, 0x30,
	0x1c, 0xc7, 0x69, 0x36, 0xfe, 0xf0, 0xcf, 0x26, 0x42, 0xf1, 0x30, 0x54, 0x64, 0x53, 0x06, 0x03,
	0x21, 0x2d, 0x2a, 0x22, 0xf1, 0x20, 0xdd, 0x0e, 0x43, 0x0f, 0x32, 0x26, 0xf4, 0x20, 0x85, 0x91,
	0x35, 0x31, 0x14, 0xb6, 0xa4, 0x24, 0xed, 0xc4, 0xa3, 0xaf, 0xe2, 0xd1, 0x47, 0xf1, 0x01, 0x7c,
	0x08, 0x9f, 0x42, 0xda, 0x34, 0x39, 0x28, 0xa8, 0xb7, 0x2f, 0xed, 0xe7, 0xfb, 0xf9, 0x25, 0xf9,
	0xc1, 0x0b, 0x2e, 0x25, 0x5f, 0xb1, 0x80, 0x50, 0x1d, 0x98, 0x58, 0xa5, 0x4d, 0x18, 0x28, 0xa6,
	0x65, 0xa9, 0x52, 0xa6, 0x83, 0x94, 0x28, 0x95, 0x31, 0xb5, 0x48, 0xa5, 0xd0, 0x05, 0x11, 0x05,
	0xca, 0x95, 0x2c, 0xa4, 0x3f, 0x30, 0x38, 0x22, 0x54, 0x23, 0xd7, 0x44, 0x9b, 0x10, 0xb9, 0xe6,
	0xee, 0x41, 0x23, 0xaf, 0x0b, 0xcb, 0xf2, 0x21, 0x78, 0x54, 0x24, 0xcf, 0x99, 0xd2, 0x46, 0x71,
	0xf8, 0xee, 0xc1, 0xed, 0x89, 0xb1, 0x4f, 0x1a, 0xb9, 0x7f, 0x04, 0xb7, 0xac, 0x60, 0x21, 0xc8,
	0x9a, 0xf5, 0xbc, 0xbe, 0x37, 0xfa, 0x3f, 0xef, 0xda, 0x8f, 0xb7, 0x64, 0xcd, 0xfc, 0x63, 0x08,
	0x32, 0xda, 0x03, 0x7d, 0x6f, 0xd4, 0x39, 0xd9, 0x6b, 0xa6, 0x23, 0x3b, 0x05, 0x5d, 0x8b, 0xe2,
	0xfc, 0x2c, 0x26, 0xab, 0x92, 0xcd, 0x41, 0x46, 0xfd, 0x10, 0xb6, 0x6b, 0x51, 0xab, 0xc6, 0xf7,
	0xbf, 0xe1, 0x77, 0x85, 0xca, 0x04, 0x37, 0x7c, 0x4d, 0xfa, 0x57, 0xb0, 0x9b, 0xca, 0x52, 0x14,
	0xea, 0x69, 0x91, 0x4a, 0xca, 0x7a, 0xed, 0x3f, 0x34, 0x3b, 0x4d, 0x63, 0x22, 0x29, 0x1b, 0x3f,
	0x03, 0x38, 0x4c, 0xe5, 0x1a, 0xfd, 0xfa, 0x44, 0xe3, 0x9d, 0x2f, 0xf7, 0x9f, 0x55, 0xee, 0x99,
	0x77, 0x7f, 0xd3, 0x54, 0xb9, 0x5c, 0x11, 0xc1, 0x91, 0x54, 0x3c, 0xe0, 0x4c, 0xd4, 0x93, 0xed,
	0x9e, 0xf2, 0x4c, 0xff, 0xb0, 0xb6, 0x4b, 0x97, 0x5e, 0x40, 0x6b, 0x1a, 0x45, 0xaf, 0x60, 0x30,
	0x35, 0xca, 0x88, 0x6a, 0x64, 0x62, 0x95, 0xe2, 0x10, 0xcd, 0x2d, 0xf9, 0x66, 0x99, 0x24, 0xa2,
	0x3a, 0x71, 0x4c, 0x12, 0x87, 0x89, 0x63, 0x3e, 0xc0, 0xd0, 0xfc, 0xc0, 0x38, 0xa2, 0x1a, 0x63,
	0x47, 0x61, 0x1c, 0x87, 0x18, 0x3b, 0x6e, 0xf9, 0xaf, 0x3e, 0xec, 0xe9, 0x67, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xf9, 0x7f, 0x40, 0xbc, 0x62, 0x02, 0x00, 0x00,
}
