// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.29.3
// source: reqresp.proto

package main

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TardisGetFeatureRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FeatureName   string                 `protobuf:"bytes,1,opt,name=feature_name,json=featureName,proto3" json:"feature_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TardisGetFeatureRequest) Reset() {
	*x = TardisGetFeatureRequest{}
	mi := &file_reqresp_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TardisGetFeatureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TardisGetFeatureRequest) ProtoMessage() {}

func (x *TardisGetFeatureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reqresp_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TardisGetFeatureRequest.ProtoReflect.Descriptor instead.
func (*TardisGetFeatureRequest) Descriptor() ([]byte, []int) {
	return file_reqresp_proto_rawDescGZIP(), []int{0}
}

func (x *TardisGetFeatureRequest) GetFeatureName() string {
	if x != nil {
		return x.FeatureName
	}
	return ""
}

type TardisGetFeaturesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Features      string                 `protobuf:"bytes,2,opt,name=features,proto3" json:"features,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TardisGetFeaturesResponse) Reset() {
	*x = TardisGetFeaturesResponse{}
	mi := &file_reqresp_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TardisGetFeaturesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TardisGetFeaturesResponse) ProtoMessage() {}

func (x *TardisGetFeaturesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reqresp_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TardisGetFeaturesResponse.ProtoReflect.Descriptor instead.
func (*TardisGetFeaturesResponse) Descriptor() ([]byte, []int) {
	return file_reqresp_proto_rawDescGZIP(), []int{1}
}

func (x *TardisGetFeaturesResponse) GetFeatures() string {
	if x != nil {
		return x.Features
	}
	return ""
}

type TardisGetCountersRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Counter       string                 `protobuf:"bytes,1,opt,name=counter,proto3" json:"counter,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TardisGetCountersRequest) Reset() {
	*x = TardisGetCountersRequest{}
	mi := &file_reqresp_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TardisGetCountersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TardisGetCountersRequest) ProtoMessage() {}

func (x *TardisGetCountersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reqresp_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TardisGetCountersRequest.ProtoReflect.Descriptor instead.
func (*TardisGetCountersRequest) Descriptor() ([]byte, []int) {
	return file_reqresp_proto_rawDescGZIP(), []int{2}
}

func (x *TardisGetCountersRequest) GetCounter() string {
	if x != nil {
		return x.Counter
	}
	return ""
}

type TardisGetCountersResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Counters      int32                  `protobuf:"varint,2,opt,name=counters,proto3" json:"counters,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TardisGetCountersResponse) Reset() {
	*x = TardisGetCountersResponse{}
	mi := &file_reqresp_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TardisGetCountersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TardisGetCountersResponse) ProtoMessage() {}

func (x *TardisGetCountersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reqresp_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TardisGetCountersResponse.ProtoReflect.Descriptor instead.
func (*TardisGetCountersResponse) Descriptor() ([]byte, []int) {
	return file_reqresp_proto_rawDescGZIP(), []int{3}
}

func (x *TardisGetCountersResponse) GetCounters() int32 {
	if x != nil {
		return x.Counters
	}
	return 0
}

var File_reqresp_proto protoreflect.FileDescriptor

var file_reqresp_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x65, 0x71, 0x72, 0x65, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x61, 0x61, 0x72, 0x73, 0x68, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x3c, 0x0a, 0x17,
	0x54, 0x61, 0x72, 0x64, 0x69, 0x73, 0x47, 0x65, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x37, 0x0a, 0x19, 0x54, 0x61,
	0x72, 0x64, 0x69, 0x73, 0x47, 0x65, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x73, 0x22, 0x34, 0x0a, 0x18, 0x54, 0x61, 0x72, 0x64, 0x69, 0x73, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x22, 0x37, 0x0a, 0x19, 0x54, 0x61, 0x72,
	0x64, 0x69, 0x73, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x73, 0x32, 0xd1, 0x01, 0x0a, 0x14, 0x54, 0x61, 0x72, 0x64, 0x69, 0x73, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5b, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x24, 0x2e, 0x61, 0x61, 0x72,
	0x73, 0x68, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x61, 0x72, 0x64, 0x69, 0x73, 0x47,
	0x65, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x61, 0x61, 0x72, 0x73, 0x68, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x54,
	0x61, 0x72, 0x64, 0x69, 0x73, 0x47, 0x65, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x12, 0x25, 0x2e, 0x61, 0x61, 0x72, 0x73, 0x68, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x61, 0x72, 0x64, 0x69, 0x73, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26,
	0x2e, 0x61, 0x61, 0x72, 0x73, 0x68, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x61, 0x72,
	0x64, 0x69, 0x73, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x61, 0x2f, 0x62, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_reqresp_proto_rawDescOnce sync.Once
	file_reqresp_proto_rawDescData = file_reqresp_proto_rawDesc
)

func file_reqresp_proto_rawDescGZIP() []byte {
	file_reqresp_proto_rawDescOnce.Do(func() {
		file_reqresp_proto_rawDescData = protoimpl.X.CompressGZIP(file_reqresp_proto_rawDescData)
	})
	return file_reqresp_proto_rawDescData
}

var file_reqresp_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_reqresp_proto_goTypes = []any{
	(*TardisGetFeatureRequest)(nil),   // 0: aarshserver.TardisGetFeatureRequest
	(*TardisGetFeaturesResponse)(nil), // 1: aarshserver.TardisGetFeaturesResponse
	(*TardisGetCountersRequest)(nil),  // 2: aarshserver.TardisGetCountersRequest
	(*TardisGetCountersResponse)(nil), // 3: aarshserver.TardisGetCountersResponse
}
var file_reqresp_proto_depIdxs = []int32{
	0, // 0: aarshserver.TardisFeatureService.GetFeatures:input_type -> aarshserver.TardisGetFeatureRequest
	2, // 1: aarshserver.TardisFeatureService.GetCounters:input_type -> aarshserver.TardisGetCountersRequest
	1, // 2: aarshserver.TardisFeatureService.GetFeatures:output_type -> aarshserver.TardisGetFeaturesResponse
	3, // 3: aarshserver.TardisFeatureService.GetCounters:output_type -> aarshserver.TardisGetCountersResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_reqresp_proto_init() }
func file_reqresp_proto_init() {
	if File_reqresp_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_reqresp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_reqresp_proto_goTypes,
		DependencyIndexes: file_reqresp_proto_depIdxs,
		MessageInfos:      file_reqresp_proto_msgTypes,
	}.Build()
	File_reqresp_proto = out.File
	file_reqresp_proto_rawDesc = nil
	file_reqresp_proto_goTypes = nil
	file_reqresp_proto_depIdxs = nil
}
