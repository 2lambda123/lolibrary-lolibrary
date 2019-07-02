// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/lolibrary/lolibrary/service.feature/proto/feature.proto

package featureproto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/lolibrary/lolibrary/cmd/protoc-gen-router/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Feature struct {
	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Slug string `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug,omitempty"`
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// timestamps
	CreatedAt            string   `protobuf:"bytes,100,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,101,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Feature) Reset()         { *m = Feature{} }
func (m *Feature) String() string { return proto.CompactTextString(m) }
func (*Feature) ProtoMessage()    {}
func (*Feature) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9eff94eab42a591, []int{0}
}

func (m *Feature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Feature.Unmarshal(m, b)
}
func (m *Feature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Feature.Marshal(b, m, deterministic)
}
func (m *Feature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Feature.Merge(m, src)
}
func (m *Feature) XXX_Size() int {
	return xxx_messageInfo_Feature.Size(m)
}
func (m *Feature) XXX_DiscardUnknown() {
	xxx_messageInfo_Feature.DiscardUnknown(m)
}

var xxx_messageInfo_Feature proto.InternalMessageInfo

func (m *Feature) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Feature) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *Feature) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Feature) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Feature) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type GETReadFeatureRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Slug                 string   `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GETReadFeatureRequest) Reset()         { *m = GETReadFeatureRequest{} }
func (m *GETReadFeatureRequest) String() string { return proto.CompactTextString(m) }
func (*GETReadFeatureRequest) ProtoMessage()    {}
func (*GETReadFeatureRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9eff94eab42a591, []int{1}
}

func (m *GETReadFeatureRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GETReadFeatureRequest.Unmarshal(m, b)
}
func (m *GETReadFeatureRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GETReadFeatureRequest.Marshal(b, m, deterministic)
}
func (m *GETReadFeatureRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GETReadFeatureRequest.Merge(m, src)
}
func (m *GETReadFeatureRequest) XXX_Size() int {
	return xxx_messageInfo_GETReadFeatureRequest.Size(m)
}
func (m *GETReadFeatureRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GETReadFeatureRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GETReadFeatureRequest proto.InternalMessageInfo

func (m *GETReadFeatureRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GETReadFeatureRequest) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

type GETReadFeatureResponse struct {
	Feature                *Feature   `protobuf:"bytes,1,opt,name=feature,proto3" json:"feature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GETReadFeatureResponse) Reset()         { *m = GETReadFeatureResponse{} }
func (m *GETReadFeatureResponse) String() string { return proto.CompactTextString(m) }
func (*GETReadFeatureResponse) ProtoMessage()    {}
func (*GETReadFeatureResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9eff94eab42a591, []int{2}
}

func (m *GETReadFeatureResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GETReadFeatureResponse.Unmarshal(m, b)
}
func (m *GETReadFeatureResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GETReadFeatureResponse.Marshal(b, m, deterministic)
}
func (m *GETReadFeatureResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GETReadFeatureResponse.Merge(m, src)
}
func (m *GETReadFeatureResponse) XXX_Size() int {
	return xxx_messageInfo_GETReadFeatureResponse.Size(m)
}
func (m *GETReadFeatureResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GETReadFeatureResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GETReadFeatureResponse proto.InternalMessageInfo

func (m *GETReadFeatureResponse) GetFeature() *Feature {
	if m != nil {
		return m.Feature
	}
	return nil
}

type GETListFeaturesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GETListFeaturesRequest) Reset()         { *m = GETListFeaturesRequest{} }
func (m *GETListFeaturesRequest) String() string { return proto.CompactTextString(m) }
func (*GETListFeaturesRequest) ProtoMessage()    {}
func (*GETListFeaturesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9eff94eab42a591, []int{3}
}

func (m *GETListFeaturesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GETListFeaturesRequest.Unmarshal(m, b)
}
func (m *GETListFeaturesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GETListFeaturesRequest.Marshal(b, m, deterministic)
}
func (m *GETListFeaturesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GETListFeaturesRequest.Merge(m, src)
}
func (m *GETListFeaturesRequest) XXX_Size() int {
	return xxx_messageInfo_GETListFeaturesRequest.Size(m)
}
func (m *GETListFeaturesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GETListFeaturesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GETListFeaturesRequest proto.InternalMessageInfo

type GETListFeaturesResponse struct {
	Features               []*Feature `protobuf:"bytes,1,rep,name=features,proto3" json:"features,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GETListFeaturesResponse) Reset()         { *m = GETListFeaturesResponse{} }
func (m *GETListFeaturesResponse) String() string { return proto.CompactTextString(m) }
func (*GETListFeaturesResponse) ProtoMessage()    {}
func (*GETListFeaturesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9eff94eab42a591, []int{4}
}

func (m *GETListFeaturesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GETListFeaturesResponse.Unmarshal(m, b)
}
func (m *GETListFeaturesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GETListFeaturesResponse.Marshal(b, m, deterministic)
}
func (m *GETListFeaturesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GETListFeaturesResponse.Merge(m, src)
}
func (m *GETListFeaturesResponse) XXX_Size() int {
	return xxx_messageInfo_GETListFeaturesResponse.Size(m)
}
func (m *GETListFeaturesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GETListFeaturesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GETListFeaturesResponse proto.InternalMessageInfo

func (m *GETListFeaturesResponse) GetFeatures() []*Feature {
	if m != nil {
		return m.Features
	}
	return nil
}

type POSTCreateFeatureRequest struct {
	Slug                 string   `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *POSTCreateFeatureRequest) Reset()         { *m = POSTCreateFeatureRequest{} }
func (m *POSTCreateFeatureRequest) String() string { return proto.CompactTextString(m) }
func (*POSTCreateFeatureRequest) ProtoMessage()    {}
func (*POSTCreateFeatureRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9eff94eab42a591, []int{5}
}

func (m *POSTCreateFeatureRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_POSTCreateFeatureRequest.Unmarshal(m, b)
}
func (m *POSTCreateFeatureRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_POSTCreateFeatureRequest.Marshal(b, m, deterministic)
}
func (m *POSTCreateFeatureRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_POSTCreateFeatureRequest.Merge(m, src)
}
func (m *POSTCreateFeatureRequest) XXX_Size() int {
	return xxx_messageInfo_POSTCreateFeatureRequest.Size(m)
}
func (m *POSTCreateFeatureRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_POSTCreateFeatureRequest.DiscardUnknown(m)
}

var xxx_messageInfo_POSTCreateFeatureRequest proto.InternalMessageInfo

func (m *POSTCreateFeatureRequest) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *POSTCreateFeatureRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type POSTCreateFeatureResponse struct {
	Feature                *Feature   `protobuf:"bytes,1,opt,name=feature,proto3" json:"feature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *POSTCreateFeatureResponse) Reset()         { *m = POSTCreateFeatureResponse{} }
func (m *POSTCreateFeatureResponse) String() string { return proto.CompactTextString(m) }
func (*POSTCreateFeatureResponse) ProtoMessage()    {}
func (*POSTCreateFeatureResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9eff94eab42a591, []int{6}
}

func (m *POSTCreateFeatureResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_POSTCreateFeatureResponse.Unmarshal(m, b)
}
func (m *POSTCreateFeatureResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_POSTCreateFeatureResponse.Marshal(b, m, deterministic)
}
func (m *POSTCreateFeatureResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_POSTCreateFeatureResponse.Merge(m, src)
}
func (m *POSTCreateFeatureResponse) XXX_Size() int {
	return xxx_messageInfo_POSTCreateFeatureResponse.Size(m)
}
func (m *POSTCreateFeatureResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_POSTCreateFeatureResponse.DiscardUnknown(m)
}

var xxx_messageInfo_POSTCreateFeatureResponse proto.InternalMessageInfo

func (m *POSTCreateFeatureResponse) GetFeature() *Feature {
	if m != nil {
		return m.Feature
	}
	return nil
}

type PUTUpdateFeatureRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PUTUpdateFeatureRequest) Reset()         { *m = PUTUpdateFeatureRequest{} }
func (m *PUTUpdateFeatureRequest) String() string { return proto.CompactTextString(m) }
func (*PUTUpdateFeatureRequest) ProtoMessage()    {}
func (*PUTUpdateFeatureRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9eff94eab42a591, []int{7}
}

func (m *PUTUpdateFeatureRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PUTUpdateFeatureRequest.Unmarshal(m, b)
}
func (m *PUTUpdateFeatureRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PUTUpdateFeatureRequest.Marshal(b, m, deterministic)
}
func (m *PUTUpdateFeatureRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PUTUpdateFeatureRequest.Merge(m, src)
}
func (m *PUTUpdateFeatureRequest) XXX_Size() int {
	return xxx_messageInfo_PUTUpdateFeatureRequest.Size(m)
}
func (m *PUTUpdateFeatureRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PUTUpdateFeatureRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PUTUpdateFeatureRequest proto.InternalMessageInfo

func (m *PUTUpdateFeatureRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PUTUpdateFeatureRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PUTUpdateFeatureResponse struct {
	Feature                *Feature   `protobuf:"bytes,1,opt,name=feature,proto3" json:"feature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PUTUpdateFeatureResponse) Reset()         { *m = PUTUpdateFeatureResponse{} }
func (m *PUTUpdateFeatureResponse) String() string { return proto.CompactTextString(m) }
func (*PUTUpdateFeatureResponse) ProtoMessage()    {}
func (*PUTUpdateFeatureResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9eff94eab42a591, []int{8}
}

func (m *PUTUpdateFeatureResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PUTUpdateFeatureResponse.Unmarshal(m, b)
}
func (m *PUTUpdateFeatureResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PUTUpdateFeatureResponse.Marshal(b, m, deterministic)
}
func (m *PUTUpdateFeatureResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PUTUpdateFeatureResponse.Merge(m, src)
}
func (m *PUTUpdateFeatureResponse) XXX_Size() int {
	return xxx_messageInfo_PUTUpdateFeatureResponse.Size(m)
}
func (m *PUTUpdateFeatureResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PUTUpdateFeatureResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PUTUpdateFeatureResponse proto.InternalMessageInfo

func (m *PUTUpdateFeatureResponse) GetFeature() *Feature {
	if m != nil {
		return m.Feature
	}
	return nil
}

type DELETERemoveFeatureRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DELETERemoveFeatureRequest) Reset()         { *m = DELETERemoveFeatureRequest{} }
func (m *DELETERemoveFeatureRequest) String() string { return proto.CompactTextString(m) }
func (*DELETERemoveFeatureRequest) ProtoMessage()    {}
func (*DELETERemoveFeatureRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9eff94eab42a591, []int{9}
}

func (m *DELETERemoveFeatureRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DELETERemoveFeatureRequest.Unmarshal(m, b)
}
func (m *DELETERemoveFeatureRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DELETERemoveFeatureRequest.Marshal(b, m, deterministic)
}
func (m *DELETERemoveFeatureRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DELETERemoveFeatureRequest.Merge(m, src)
}
func (m *DELETERemoveFeatureRequest) XXX_Size() int {
	return xxx_messageInfo_DELETERemoveFeatureRequest.Size(m)
}
func (m *DELETERemoveFeatureRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DELETERemoveFeatureRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DELETERemoveFeatureRequest proto.InternalMessageInfo

func (m *DELETERemoveFeatureRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DELETERemoveFeatureResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DELETERemoveFeatureResponse) Reset()         { *m = DELETERemoveFeatureResponse{} }
func (m *DELETERemoveFeatureResponse) String() string { return proto.CompactTextString(m) }
func (*DELETERemoveFeatureResponse) ProtoMessage()    {}
func (*DELETERemoveFeatureResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9eff94eab42a591, []int{10}
}

func (m *DELETERemoveFeatureResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DELETERemoveFeatureResponse.Unmarshal(m, b)
}
func (m *DELETERemoveFeatureResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DELETERemoveFeatureResponse.Marshal(b, m, deterministic)
}
func (m *DELETERemoveFeatureResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DELETERemoveFeatureResponse.Merge(m, src)
}
func (m *DELETERemoveFeatureResponse) XXX_Size() int {
	return xxx_messageInfo_DELETERemoveFeatureResponse.Size(m)
}
func (m *DELETERemoveFeatureResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DELETERemoveFeatureResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DELETERemoveFeatureResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Feature)(nil), "featureproto.Feature")
	proto.RegisterType((*GETReadFeatureRequest)(nil), "featureproto.GETReadFeatureRequest")
	proto.RegisterType((*GETReadFeatureResponse)(nil), "featureproto.GETReadFeatureResponse")
	proto.RegisterType((*GETListFeaturesRequest)(nil), "featureproto.GETListFeaturesRequest")
	proto.RegisterType((*GETListFeaturesResponse)(nil), "featureproto.GETListFeaturesResponse")
	proto.RegisterType((*POSTCreateFeatureRequest)(nil), "featureproto.POSTCreateFeatureRequest")
	proto.RegisterType((*POSTCreateFeatureResponse)(nil), "featureproto.POSTCreateFeatureResponse")
	proto.RegisterType((*PUTUpdateFeatureRequest)(nil), "featureproto.PUTUpdateFeatureRequest")
	proto.RegisterType((*PUTUpdateFeatureResponse)(nil), "featureproto.PUTUpdateFeatureResponse")
	proto.RegisterType((*DELETERemoveFeatureRequest)(nil), "featureproto.DELETERemoveFeatureRequest")
	proto.RegisterType((*DELETERemoveFeatureResponse)(nil), "featureproto.DELETERemoveFeatureResponse")
}

func init() {
	proto.RegisterFile("github.com/lolibrary/lolibrary/service.feature/proto/feature.proto", fileDescriptor_d9eff94eab42a591)
}

var fileDescriptor_d9eff94eab42a591 = []byte{
	// 504 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0x55, 0xd2, 0xa4, 0x91, 0xc7, 0xea, 0xef, 0x27, 0x16, 0x1a, 0x8c, 0x51, 0x44, 0xba, 0x80,
	0xf8, 0x23, 0xc5, 0x96, 0xca, 0x09, 0x21, 0x01, 0x69, 0x1b, 0xf5, 0x52, 0x89, 0x2a, 0xa4, 0x27,
	0x24, 0x90, 0xe3, 0x1d, 0x05, 0x4b, 0x4e, 0x36, 0x78, 0xd7, 0x95, 0x38, 0x70, 0xc9, 0xf7, 0xe0,
	0xda, 0x6f, 0xd0, 0xef, 0x87, 0x3c, 0xeb, 0x10, 0xc7, 0x75, 0x13, 0x7a, 0x1b, 0xcf, 0xbc, 0x79,
	0x6f, 0xf6, 0xed, 0xac, 0xe1, 0xfd, 0x24, 0xd2, 0xdf, 0xd3, 0xb1, 0x17, 0xca, 0xa9, 0x1f, 0xcb,
	0x38, 0x1a, 0x27, 0x41, 0xf2, 0xb3, 0x10, 0x29, 0x4c, 0x2e, 0xa3, 0x10, 0xbd, 0x50, 0xc6, 0x32,
	0xf1, 0xe7, 0x89, 0xd4, 0xd2, 0xa7, 0xd8, 0xa3, 0x98, 0x01, 0x7d, 0x50, 0xec, 0x9e, 0x6e, 0xe1,
	0x0a, 0xa7, 0xc2, 0x30, 0x84, 0xbd, 0x09, 0xce, 0x7a, 0x89, 0x4c, 0x35, 0x2e, 0x39, 0xcd, 0x87,
	0x21, 0xe5, 0xbf, 0xa0, 0x79, 0x9c, 0xd1, 0xb2, 0xff, 0xa0, 0x1e, 0x09, 0xa7, 0xd6, 0xad, 0xbd,
	0xb4, 0x86, 0xf5, 0x48, 0x30, 0x06, 0x0d, 0x15, 0xa7, 0x13, 0xa7, 0x4e, 0x19, 0x8a, 0xb3, 0xdc,
	0x2c, 0x98, 0xa2, 0xd3, 0x30, 0xb9, 0x2c, 0x66, 0x1d, 0x80, 0x30, 0xc1, 0x40, 0xa3, 0xf8, 0x16,
	0x68, 0x47, 0x50, 0xc5, 0xca, 0x33, 0x7d, 0x9d, 0x95, 0xd3, 0xb9, 0x58, 0x96, 0xd1, 0x94, 0xf3,
	0x4c, 0x5f, 0xf3, 0xb7, 0x70, 0xff, 0x74, 0x30, 0x1a, 0x62, 0x20, 0x68, 0x8a, 0x21, 0xfe, 0x48,
	0x51, 0xe9, 0x7f, 0x19, 0x86, 0x7f, 0x80, 0x07, 0xeb, 0xad, 0x6a, 0x2e, 0x67, 0x0a, 0xd9, 0x0b,
	0x68, 0x92, 0x51, 0xd4, 0x6e, 0x1f, 0xde, 0xf3, 0x56, 0xb6, 0x79, 0x06, 0x69, 0xea, 0xbc, 0x4d,
	0x04, 0x67, 0x91, 0xd2, 0x94, 0x56, 0xb9, 0x38, 0x3f, 0x82, 0xfd, 0x52, 0x3e, 0x67, 0x7e, 0x05,
	0xbb, 0xd4, 0xa9, 0x9c, 0x5a, 0x77, 0xa7, 0x9a, 0x3a, 0x07, 0xf0, 0x8f, 0xd0, 0x3e, 0xff, 0xf4,
	0x79, 0x74, 0x4c, 0x3e, 0xac, 0x1d, 0x6d, 0x79, 0x94, 0x5a, 0x85, 0xaf, 0x3b, 0x2b, 0x5f, 0xf9,
	0x11, 0x3c, 0xbc, 0xc1, 0x70, 0xd7, 0x13, 0xbe, 0x83, 0xfd, 0xf3, 0x8b, 0xd1, 0x05, 0xb9, 0xbd,
	0xcd, 0x5f, 0x1a, 0xa0, 0x5e, 0x18, 0xa0, 0x0f, 0xed, 0x72, 0xf3, 0x5d, 0xf5, 0x5f, 0x83, 0x73,
	0x32, 0x38, 0x1b, 0x8c, 0x06, 0x43, 0x9c, 0xca, 0xcb, 0x8d, 0x23, 0xf0, 0xc7, 0xf0, 0xa8, 0x02,
	0x6b, 0x14, 0x0f, 0x7f, 0x37, 0x72, 0x49, 0xf6, 0x05, 0xac, 0xbf, 0x57, 0xce, 0x9e, 0x14, 0x95,
	0x2b, 0xf6, 0xc8, 0xed, 0xde, 0x0e, 0x30, 0xcc, 0xdc, 0x5e, 0x5c, 0x3b, 0x2d, 0x68, 0xfa, 0x09,
	0x06, 0x82, 0x85, 0x60, 0x17, 0xce, 0xcb, 0x0e, 0x8a, 0xdd, 0x95, 0x46, 0xba, 0x7c, 0x13, 0x24,
	0x97, 0xd8, 0x5b, 0x5c, 0x3b, 0x16, 0xb4, 0x7c, 0xb3, 0xf5, 0x2c, 0x02, 0xfb, 0x04, 0x63, 0x5c,
	0x8a, 0x3c, 0x2b, 0x32, 0xdc, 0xe6, 0x96, 0xfb, 0x7c, 0x0b, 0xaa, 0x24, 0x25, 0x48, 0x80, 0x21,
	0xd8, 0x85, 0xfd, 0x61, 0xeb, 0xc3, 0x56, 0xae, 0xa7, 0xfb, 0x74, 0x23, 0xa6, 0x24, 0x63, 0x9e,
	0x39, 0xfb, 0x0a, 0xb0, 0x7a, 0x2d, 0xac, 0xec, 0xf9, 0x8d, 0x07, 0xe6, 0x1e, 0x6c, 0x40, 0x94,
	0xae, 0x25, 0x8e, 0xb2, 0x5b, 0x5c, 0x5c, 0x75, 0xfe, 0x87, 0xbd, 0xb5, 0xff, 0xe3, 0xe2, 0xaa,
	0x63, 0xb1, 0x96, 0xea, 0xd1, 0xc7, 0x78, 0x97, 0xb8, 0xde, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff,
	0x9d, 0xde, 0x58, 0xfb, 0x63, 0x05, 0x00, 0x00,
}
