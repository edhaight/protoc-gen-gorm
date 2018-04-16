// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example/feature_demo/test.proto

/*
Package example is a generated protocol buffer package.

It is generated from these files:
	example/feature_demo/test.proto
	example/feature_demo/test2.proto

It has these top-level messages:
	TestTypes
	TypeWithId
	MultitenantTypeWithId
	MultitenantTypeWithoutId
	ApiOnlyType
	TypeBecomesEmpty
	IntPoint
	CreateIntPointRequest
	CreateIntPointResponse
	ReadIntPointRequest
	ReadIntPointResponse
	UpdateIntPointRequest
	UpdateIntPointResponse
	DeleteIntPointRequest
	ListIntPointResponse
	Something
*/
package example

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/infobloxopen/protoc-gen-gorm/options"
import gormable_types "github.com/infobloxopen/protoc-gen-gorm/types"
import google_protobuf1 "github.com/golang/protobuf/ptypes/wrappers"
import google_protobuf2 "github.com/golang/protobuf/ptypes/empty"
import google_protobuf3 "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TestTypesStatus int32

const (
	TestTypes_UNKNOWN TestTypesStatus = 0
	TestTypes_GOOD    TestTypesStatus = 1
	TestTypes_BAD     TestTypesStatus = 2
)

var TestTypesStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "GOOD",
	2: "BAD",
}
var TestTypesStatus_value = map[string]int32{
	"UNKNOWN": 0,
	"GOOD":    1,
	"BAD":     2,
}

func (x TestTypesStatus) String() string {
	return proto.EnumName(TestTypesStatus_name, int32(x))
}
func (TestTypesStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// test_types is a message that serves as an example
type TestTypes struct {
	ApiOnlyString  string                        `protobuf:"bytes,1,opt,name=api_only_string,json=apiOnlyString" json:"api_only_string,omitempty"`
	Numbers        []int32                       `protobuf:"varint,2,rep,packed,name=numbers" json:"numbers,omitempty"`
	OptionalString *google_protobuf1.StringValue `protobuf:"bytes,3,opt,name=optional_string,json=optionalString" json:"optional_string,omitempty"`
	BecomesInt     TestTypesStatus               `protobuf:"varint,4,opt,name=becomes_int,json=becomesInt,enum=example.TestTypesStatus" json:"becomes_int,omitempty"`
	Nothingness    *google_protobuf2.Empty       `protobuf:"bytes,5,opt,name=nothingness" json:"nothingness,omitempty"`
	Uuid           *gormable_types.UUIDValue     `protobuf:"bytes,6,opt,name=uuid" json:"uuid,omitempty"`
	CreatedAt      *google_protobuf3.Timestamp   `protobuf:"bytes,7,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
}

func (m *TestTypes) Reset()                    { *m = TestTypes{} }
func (m *TestTypes) String() string            { return proto.CompactTextString(m) }
func (*TestTypes) ProtoMessage()               {}
func (*TestTypes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TestTypes) GetApiOnlyString() string {
	if m != nil {
		return m.ApiOnlyString
	}
	return ""
}

func (m *TestTypes) GetNumbers() []int32 {
	if m != nil {
		return m.Numbers
	}
	return nil
}

func (m *TestTypes) GetOptionalString() *google_protobuf1.StringValue {
	if m != nil {
		return m.OptionalString
	}
	return nil
}

func (m *TestTypes) GetBecomesInt() TestTypesStatus {
	if m != nil {
		return m.BecomesInt
	}
	return TestTypes_UNKNOWN
}

func (m *TestTypes) GetNothingness() *google_protobuf2.Empty {
	if m != nil {
		return m.Nothingness
	}
	return nil
}

func (m *TestTypes) GetUuid() *gormable_types.UUIDValue {
	if m != nil {
		return m.Uuid
	}
	return nil
}

func (m *TestTypes) GetCreatedAt() *google_protobuf3.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type TypeWithId struct {
	Ip            string       `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	Things        []*TestTypes `protobuf:"bytes,3,rep,name=things" json:"things,omitempty"`
	ANestedObject *TestTypes   `protobuf:"bytes,4,opt,name=a_nested_object,json=aNestedObject" json:"a_nested_object,omitempty"`
}

func (m *TypeWithId) Reset()                    { *m = TypeWithId{} }
func (m *TypeWithId) String() string            { return proto.CompactTextString(m) }
func (*TypeWithId) ProtoMessage()               {}
func (*TypeWithId) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TypeWithId) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *TypeWithId) GetThings() []*TestTypes {
	if m != nil {
		return m.Things
	}
	return nil
}

func (m *TypeWithId) GetANestedObject() *TestTypes {
	if m != nil {
		return m.ANestedObject
	}
	return nil
}

type MultitenantTypeWithId struct {
	Id        uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	SomeField string `protobuf:"bytes,2,opt,name=some_field,json=someField" json:"some_field,omitempty"`
}

func (m *MultitenantTypeWithId) Reset()                    { *m = MultitenantTypeWithId{} }
func (m *MultitenantTypeWithId) String() string            { return proto.CompactTextString(m) }
func (*MultitenantTypeWithId) ProtoMessage()               {}
func (*MultitenantTypeWithId) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MultitenantTypeWithId) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MultitenantTypeWithId) GetSomeField() string {
	if m != nil {
		return m.SomeField
	}
	return ""
}

type MultitenantTypeWithoutId struct {
	SomeField string `protobuf:"bytes,1,opt,name=some_field,json=someField" json:"some_field,omitempty"`
}

func (m *MultitenantTypeWithoutId) Reset()                    { *m = MultitenantTypeWithoutId{} }
func (m *MultitenantTypeWithoutId) String() string            { return proto.CompactTextString(m) }
func (*MultitenantTypeWithoutId) ProtoMessage()               {}
func (*MultitenantTypeWithoutId) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MultitenantTypeWithoutId) GetSomeField() string {
	if m != nil {
		return m.SomeField
	}
	return ""
}

type ApiOnlyType struct {
	Contents string `protobuf:"bytes,1,opt,name=contents" json:"contents,omitempty"`
}

func (m *ApiOnlyType) Reset()                    { *m = ApiOnlyType{} }
func (m *ApiOnlyType) String() string            { return proto.CompactTextString(m) }
func (*ApiOnlyType) ProtoMessage()               {}
func (*ApiOnlyType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ApiOnlyType) GetContents() string {
	if m != nil {
		return m.Contents
	}
	return ""
}

type TypeBecomesEmpty struct {
	AThing *ApiOnlyType `protobuf:"bytes,1,opt,name=a_thing,json=aThing" json:"a_thing,omitempty"`
}

func (m *TypeBecomesEmpty) Reset()                    { *m = TypeBecomesEmpty{} }
func (m *TypeBecomesEmpty) String() string            { return proto.CompactTextString(m) }
func (*TypeBecomesEmpty) ProtoMessage()               {}
func (*TypeBecomesEmpty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *TypeBecomesEmpty) GetAThing() *ApiOnlyType {
	if m != nil {
		return m.AThing
	}
	return nil
}

func init() {
	proto.RegisterType((*TestTypes)(nil), "example.test_types")
	proto.RegisterType((*TypeWithId)(nil), "example.type_with_id")
	proto.RegisterType((*MultitenantTypeWithId)(nil), "example.multitenant_type_with_id")
	proto.RegisterType((*MultitenantTypeWithoutId)(nil), "example.multitenant_type_without_id")
	proto.RegisterType((*ApiOnlyType)(nil), "example.api_only_type")
	proto.RegisterType((*TypeBecomesEmpty)(nil), "example.type_becomes_empty")
	proto.RegisterEnum("example.TestTypesStatus", TestTypesStatus_name, TestTypesStatus_value)
}

func init() { proto.RegisterFile("example/feature_demo/test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 690 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x5d, 0x6b, 0xe3, 0x38,
	0x14, 0x5d, 0x3b, 0x69, 0x92, 0xde, 0x6c, 0x93, 0xa0, 0x2e, 0x8b, 0x93, 0xee, 0x6e, 0xb3, 0x66,
	0x17, 0xb2, 0x94, 0xda, 0x90, 0xbe, 0x6c, 0xd3, 0xa7, 0x96, 0xb4, 0x43, 0x19, 0x26, 0x01, 0xb7,
	0x9d, 0x81, 0x79, 0x31, 0x72, 0xac, 0xb8, 0x9a, 0xb1, 0x25, 0x61, 0xc9, 0xb4, 0xf9, 0x69, 0x09,
	0xcc, 0x0f, 0x98, 0x7f, 0x35, 0x58, 0xb6, 0x3b, 0x33, 0xfd, 0x80, 0xbe, 0x18, 0xa4, 0x7b, 0xce,
	0xb9, 0xe7, 0x1e, 0x59, 0x82, 0x7d, 0x72, 0x8f, 0x13, 0x11, 0x13, 0x77, 0x49, 0xb0, 0xca, 0x52,
	0xe2, 0x87, 0x24, 0xe1, 0xae, 0x22, 0x52, 0x39, 0x22, 0xe5, 0x8a, 0xa3, 0x66, 0x09, 0x18, 0x4c,
	0x22, 0xaa, 0x6e, 0xb3, 0xc0, 0x59, 0xf0, 0xc4, 0xa5, 0x6c, 0xc9, 0x83, 0x98, 0xdf, 0x73, 0x41,
	0x98, 0xab, 0x71, 0x8b, 0xc3, 0x88, 0xb0, 0xc3, 0x88, 0xa7, 0x89, 0xcb, 0x85, 0xa2, 0x9c, 0x49,
	0x37, 0x5f, 0x14, 0x22, 0x83, 0xe3, 0xd7, 0x72, 0xd5, 0x4a, 0x10, 0x59, 0x7c, 0x4b, 0xea, 0x5f,
	0x11, 0xe7, 0x51, 0x4c, 0x0a, 0x64, 0x90, 0x2d, 0xdd, 0xbb, 0x14, 0x0b, 0x41, 0xd2, 0xaa, 0xbe,
	0xf7, 0xb8, 0x4e, 0x12, 0xa1, 0x56, 0x65, 0x71, 0xff, 0x71, 0x51, 0xd1, 0x84, 0x48, 0x85, 0x13,
	0x51, 0x00, 0xec, 0x2f, 0x35, 0x80, 0x7c, 0x58, 0x5f, 0xb7, 0x44, 0x0e, 0x74, 0xb1, 0xa0, 0x3e,
	0x67, 0xf1, 0xca, 0x97, 0x2a, 0xa5, 0x2c, 0xb2, 0x8c, 0xa1, 0x31, 0xda, 0x3e, 0x6b, 0x6c, 0xd6,
	0x7d, 0xb3, 0x67, 0x78, 0x3b, 0x58, 0xd0, 0x39, 0x8b, 0x57, 0x57, 0xba, 0x88, 0x2c, 0x68, 0xb2,
	0x2c, 0x09, 0x48, 0x2a, 0x2d, 0x73, 0x58, 0x1b, 0x6d, 0x79, 0xd5, 0x12, 0x9d, 0x43, 0xb7, 0xc8,
	0x01, 0xc7, 0x95, 0x52, 0x6d, 0x68, 0x8c, 0xda, 0xe3, 0x3f, 0x9c, 0xc2, 0x93, 0x53, 0x79, 0x72,
	0x0a, 0xad, 0xf7, 0x38, 0xce, 0x88, 0xd7, 0xa9, 0x48, 0x65, 0x83, 0x13, 0x68, 0x07, 0x64, 0xc1,
	0x13, 0x22, 0x7d, 0xca, 0x94, 0x55, 0x1f, 0x1a, 0xa3, 0xce, 0x78, 0xe0, 0x94, 0x67, 0xe2, 0x7c,
	0xb7, 0xee, 0x48, 0x85, 0x55, 0x26, 0x3d, 0x28, 0xe1, 0x97, 0x4c, 0xa1, 0xff, 0xa1, 0xcd, 0xb8,
	0xba, 0xa5, 0x2c, 0x62, 0x44, 0x4a, 0x6b, 0x4b, 0xf7, 0xff, 0xfd, 0x49, 0xff, 0xf3, 0x3c, 0x30,
	0xef, 0x47, 0x28, 0x3a, 0x84, 0x7a, 0x96, 0xd1, 0xd0, 0x6a, 0x68, 0x4a, 0xdf, 0xc9, 0xcf, 0x06,
	0x07, 0x31, 0x29, 0x7b, 0xdd, 0xdc, 0x5c, 0x4e, 0x0b, 0xbf, 0x1a, 0x86, 0x8e, 0x01, 0x16, 0x29,
	0xc1, 0x8a, 0x84, 0x3e, 0x56, 0x56, 0x53, 0x93, 0x06, 0x4f, 0xfa, 0x5c, 0x57, 0xd9, 0x7b, 0xdb,
	0x25, 0xfa, 0x54, 0xd9, 0x23, 0x68, 0x14, 0xce, 0x51, 0x1b, 0x9a, 0x37, 0xb3, 0xb7, 0xb3, 0xf9,
	0x87, 0x59, 0xef, 0x17, 0xd4, 0x82, 0xfa, 0x9b, 0xf9, 0x7c, 0xda, 0x33, 0x50, 0x13, 0x6a, 0x67,
	0xa7, 0xd3, 0x9e, 0x39, 0xd9, 0xdd, 0xac, 0xfb, 0xdd, 0x96, 0x31, 0x68, 0xcb, 0x84, 0xa7, 0x11,
	0x96, 0x01, 0x4f, 0x43, 0xfb, 0xab, 0x01, 0xbf, 0xe6, 0x9e, 0xfc, 0x3b, 0xaa, 0x6e, 0x7d, 0x1a,
	0xa2, 0x7f, 0xc0, 0xa4, 0xa2, 0x3c, 0xb4, 0xdf, 0x36, 0xeb, 0x7e, 0x0f, 0x3a, 0xb9, 0xf9, 0x89,
	0x4d, 0x85, 0x8f, 0xc3, 0x30, 0xb5, 0x3d, 0x93, 0x0a, 0x74, 0x00, 0x0d, 0x3d, 0xac, 0xb4, 0x6a,
	0xc3, 0xda, 0xa8, 0x3d, 0xde, 0x7d, 0x26, 0x51, 0xaf, 0x84, 0xa0, 0x13, 0xe8, 0x62, 0x9f, 0x11,
	0x99, 0x8f, 0xc7, 0x83, 0x4f, 0x64, 0x51, 0x9c, 0xc3, 0x0b, 0xac, 0x1d, 0x3c, 0xd3, 0xd0, 0xb9,
	0x46, 0x4e, 0xfe, 0xdb, 0xac, 0xfb, 0xff, 0xb6, 0x0c, 0xf4, 0x37, 0x6c, 0x51, 0xa6, 0x8e, 0xc6,
	0x48, 0x27, 0x36, 0x40, 0x85, 0x29, 0x91, 0xd2, 0x04, 0xa7, 0x2b, 0xff, 0x33, 0x59, 0xd9, 0xf6,
	0x15, 0x58, 0x49, 0x16, 0x2b, 0xaa, 0x08, 0xc3, 0xac, 0x90, 0x7b, 0x18, 0xab, 0x03, 0x26, 0x0d,
	0xf5, 0x58, 0x75, 0xcf, 0xa4, 0x21, 0xfa, 0x13, 0x40, 0xf2, 0x84, 0xf8, 0x4b, 0x4a, 0xe2, 0xd0,
	0x32, 0xf3, 0x71, 0xbd, 0xed, 0x7c, 0xe7, 0x22, 0xdf, 0x98, 0xb4, 0x36, 0xeb, 0x7e, 0xbd, 0x65,
	0x0c, 0x0d, 0xfb, 0x02, 0xf6, 0x9e, 0x15, 0xe5, 0x99, 0xf2, 0x9f, 0xe8, 0x18, 0x2f, 0xeb, 0x1c,
	0xc0, 0xce, 0xc3, 0xcd, 0xc8, 0x45, 0xd0, 0x00, 0x5a, 0x0b, 0xce, 0x14, 0x61, 0x4a, 0x96, 0xbc,
	0x87, 0xb5, 0xfd, 0x0e, 0x90, 0x6e, 0x54, 0xfd, 0xba, 0xfa, 0x4a, 0x22, 0x17, 0x9a, 0xd8, 0xd7,
	0x99, 0x6a, 0x42, 0xfe, 0x2b, 0x56, 0xf9, 0xfd, 0x24, 0xed, 0x35, 0xf0, 0x75, 0x8e, 0x9a, 0xe8,
	0x4b, 0xd7, 0x32, 0xce, 0x2e, 0x3e, 0x4e, 0x5f, 0xfb, 0x7e, 0x3c, 0xf7, 0x9a, 0x9d, 0x94, 0x9b,
	0x41, 0x43, 0xa3, 0x8f, 0xbe, 0x05, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x69, 0x31, 0xf4, 0xf4, 0x04,
	0x00, 0x00,
}
