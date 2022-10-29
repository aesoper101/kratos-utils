// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: response.proto

package httppb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OperationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *OperationResponse) Reset() {
	*x = OperationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationResponse) ProtoMessage() {}

func (x *OperationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationResponse.ProtoReflect.Descriptor instead.
func (*OperationResponse) Descriptor() ([]byte, []int) {
	return file_response_proto_rawDescGZIP(), []int{0}
}

func (x *OperationResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type FuzzySearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keywords string   `protobuf:"bytes,1,opt,name=Keywords,proto3" json:"Keywords,omitempty"`
	Field    []string `protobuf:"bytes,2,rep,name=field,proto3" json:"field,omitempty"`
}

func (x *FuzzySearchRequest) Reset() {
	*x = FuzzySearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_response_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FuzzySearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FuzzySearchRequest) ProtoMessage() {}

func (x *FuzzySearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_response_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FuzzySearchRequest.ProtoReflect.Descriptor instead.
func (*FuzzySearchRequest) Descriptor() ([]byte, []int) {
	return file_response_proto_rawDescGZIP(), []int{1}
}

func (x *FuzzySearchRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *FuzzySearchRequest) GetField() []string {
	if x != nil {
		return x.Field
	}
	return nil
}

type AdvancedFilterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field    string     `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Operator string     `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
	Value    *anypb.Any `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *AdvancedFilterRequest) Reset() {
	*x = AdvancedFilterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_response_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdvancedFilterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdvancedFilterRequest) ProtoMessage() {}

func (x *AdvancedFilterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_response_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdvancedFilterRequest.ProtoReflect.Descriptor instead.
func (*AdvancedFilterRequest) Descriptor() ([]byte, []int) {
	return file_response_proto_rawDescGZIP(), []int{2}
}

func (x *AdvancedFilterRequest) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *AdvancedFilterRequest) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

func (x *AdvancedFilterRequest) GetValue() *anypb.Any {
	if x != nil {
		return x.Value
	}
	return nil
}

type SortRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Order string `protobuf:"bytes,2,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *SortRequest) Reset() {
	*x = SortRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_response_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SortRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SortRequest) ProtoMessage() {}

func (x *SortRequest) ProtoReflect() protoreflect.Message {
	mi := &file_response_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SortRequest.ProtoReflect.Descriptor instead.
func (*SortRequest) Descriptor() ([]byte, []int) {
	return file_response_proto_rawDescGZIP(), []int{3}
}

func (x *SortRequest) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *SortRequest) GetOrder() string {
	if x != nil {
		return x.Order
	}
	return ""
}

var File_response_proto protoreflect.FileDescriptor

var file_response_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x11, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x46, 0x0a, 0x12, 0x46, 0x75, 0x7a, 0x7a, 0x79,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x22,
	0x75, 0x0a, 0x15, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x39, 0x0a, 0x0b, 0x53, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x65, 0x73, 0x6f, 0x70, 0x65, 0x72, 0x31, 0x30, 0x31, 0x2f, 0x6b, 0x72, 0x61, 0x74, 0x6f,
	0x73, 0x2d, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x70, 0x62, 0x3b, 0x68, 0x74,
	0x74, 0x70, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_response_proto_rawDescOnce sync.Once
	file_response_proto_rawDescData = file_response_proto_rawDesc
)

func file_response_proto_rawDescGZIP() []byte {
	file_response_proto_rawDescOnce.Do(func() {
		file_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_response_proto_rawDescData)
	})
	return file_response_proto_rawDescData
}

var file_response_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_response_proto_goTypes = []interface{}{
	(*OperationResponse)(nil),     // 0: kratos.http.OperationResponse
	(*FuzzySearchRequest)(nil),    // 1: kratos.http.FuzzySearchRequest
	(*AdvancedFilterRequest)(nil), // 2: kratos.http.AdvancedFilterRequest
	(*SortRequest)(nil),           // 3: kratos.http.SortRequest
	(*anypb.Any)(nil),             // 4: google.protobuf.Any
}
var file_response_proto_depIdxs = []int32{
	4, // 0: kratos.http.AdvancedFilterRequest.value:type_name -> google.protobuf.Any
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_response_proto_init() }
func file_response_proto_init() {
	if File_response_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationResponse); i {
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
		file_response_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FuzzySearchRequest); i {
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
		file_response_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdvancedFilterRequest); i {
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
		file_response_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SortRequest); i {
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
			RawDescriptor: file_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_response_proto_goTypes,
		DependencyIndexes: file_response_proto_depIdxs,
		MessageInfos:      file_response_proto_msgTypes,
	}.Build()
	File_response_proto = out.File
	file_response_proto_rawDesc = nil
	file_response_proto_goTypes = nil
	file_response_proto_depIdxs = nil
}
