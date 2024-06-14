// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: string.proto

package pb

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

type StringRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input string `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
}

func (x *StringRequest) Reset() {
	*x = StringRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_string_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringRequest) ProtoMessage() {}

func (x *StringRequest) ProtoReflect() protoreflect.Message {
	mi := &file_string_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringRequest.ProtoReflect.Descriptor instead.
func (*StringRequest) Descriptor() ([]byte, []int) {
	return file_string_proto_rawDescGZIP(), []int{0}
}

func (x *StringRequest) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

type StringResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Output string `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *StringResponse) Reset() {
	*x = StringResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_string_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringResponse) ProtoMessage() {}

func (x *StringResponse) ProtoReflect() protoreflect.Message {
	mi := &file_string_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringResponse.ProtoReflect.Descriptor instead.
func (*StringResponse) Descriptor() ([]byte, []int) {
	return file_string_proto_rawDescGZIP(), []int{1}
}

func (x *StringResponse) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

type LengthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Length int32 `protobuf:"varint,1,opt,name=length,proto3" json:"length,omitempty"`
}

func (x *LengthResponse) Reset() {
	*x = LengthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_string_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LengthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LengthResponse) ProtoMessage() {}

func (x *LengthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_string_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LengthResponse.ProtoReflect.Descriptor instead.
func (*LengthResponse) Descriptor() ([]byte, []int) {
	return file_string_proto_rawDescGZIP(), []int{2}
}

func (x *LengthResponse) GetLength() int32 {
	if x != nil {
		return x.Length
	}
	return 0
}

var File_string_proto protoreflect.FileDescriptor

var file_string_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25,
	0x0a, 0x0d, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x28, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22,
	0x28, 0x0a, 0x0e, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x32, 0xc6, 0x01, 0x0a, 0x0d, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x54,
	0x6f, 0x55, 0x70, 0x70, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x07, 0x54, 0x6f, 0x4c,
	0x6f, 0x77, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x07, 0x52, 0x65, 0x76, 0x65, 0x72,
	0x73, 0x65, 0x12, 0x0e, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x06, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12,
	0x0e, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0f, 0x2e, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_string_proto_rawDescOnce sync.Once
	file_string_proto_rawDescData = file_string_proto_rawDesc
)

func file_string_proto_rawDescGZIP() []byte {
	file_string_proto_rawDescOnce.Do(func() {
		file_string_proto_rawDescData = protoimpl.X.CompressGZIP(file_string_proto_rawDescData)
	})
	return file_string_proto_rawDescData
}

var file_string_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_string_proto_goTypes = []any{
	(*StringRequest)(nil),  // 0: StringRequest
	(*StringResponse)(nil), // 1: StringResponse
	(*LengthResponse)(nil), // 2: LengthResponse
}
var file_string_proto_depIdxs = []int32{
	0, // 0: StringService.ToUpper:input_type -> StringRequest
	0, // 1: StringService.ToLower:input_type -> StringRequest
	0, // 2: StringService.Reverse:input_type -> StringRequest
	0, // 3: StringService.Length:input_type -> StringRequest
	1, // 4: StringService.ToUpper:output_type -> StringResponse
	1, // 5: StringService.ToLower:output_type -> StringResponse
	1, // 6: StringService.Reverse:output_type -> StringResponse
	2, // 7: StringService.Length:output_type -> LengthResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_string_proto_init() }
func file_string_proto_init() {
	if File_string_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_string_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*StringRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_string_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*StringResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_string_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*LengthResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_string_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_string_proto_goTypes,
		DependencyIndexes: file_string_proto_depIdxs,
		MessageInfos:      file_string_proto_msgTypes,
	}.Build()
	File_string_proto = out.File
	file_string_proto_rawDesc = nil
	file_string_proto_goTypes = nil
	file_string_proto_depIdxs = nil
}
