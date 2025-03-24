// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.21.12
// source: kvstore.proto

package kvstore_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetIn struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetIn) Reset() {
	*x = GetIn{}
	mi := &file_kvstore_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIn) ProtoMessage() {}

func (x *GetIn) ProtoReflect() protoreflect.Message {
	mi := &file_kvstore_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIn.ProtoReflect.Descriptor instead.
func (*GetIn) Descriptor() ([]byte, []int) {
	return file_kvstore_proto_rawDescGZIP(), []int{0}
}

func (x *GetIn) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetOut struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         string                 `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOut) Reset() {
	*x = GetOut{}
	mi := &file_kvstore_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOut) ProtoMessage() {}

func (x *GetOut) ProtoReflect() protoreflect.Message {
	mi := &file_kvstore_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOut.ProtoReflect.Descriptor instead.
func (*GetOut) Descriptor() ([]byte, []int) {
	return file_kvstore_proto_rawDescGZIP(), []int{1}
}

func (x *GetOut) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type PutIn struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value         string                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PutIn) Reset() {
	*x = PutIn{}
	mi := &file_kvstore_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PutIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutIn) ProtoMessage() {}

func (x *PutIn) ProtoReflect() protoreflect.Message {
	mi := &file_kvstore_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutIn.ProtoReflect.Descriptor instead.
func (*PutIn) Descriptor() ([]byte, []int) {
	return file_kvstore_proto_rawDescGZIP(), []int{2}
}

func (x *PutIn) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *PutIn) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type DeleteIn struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteIn) Reset() {
	*x = DeleteIn{}
	mi := &file_kvstore_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteIn) ProtoMessage() {}

func (x *DeleteIn) ProtoReflect() protoreflect.Message {
	mi := &file_kvstore_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteIn.ProtoReflect.Descriptor instead.
func (*DeleteIn) Descriptor() ([]byte, []int) {
	return file_kvstore_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteIn) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

var File_kvstore_proto protoreflect.FileDescriptor

var file_kvstore_proto_rawDesc = string([]byte{
	0x0a, 0x0d, 0x6b, 0x76, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6b, 0x76, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x19, 0x0a, 0x05, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x22, 0x1e, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x4f, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x2f, 0x0a, 0x05, 0x50, 0x75, 0x74, 0x49, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x1c, 0x0a, 0x08, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x32,
	0x95, 0x01, 0x0a, 0x07, 0x4b, 0x56, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x26, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x0e, 0x2e, 0x6b, 0x76, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x49, 0x6e, 0x1a, 0x0f, 0x2e, 0x6b, 0x76, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x4f, 0x75, 0x74, 0x12, 0x2d, 0x0a, 0x03, 0x50, 0x75, 0x74, 0x12, 0x0e, 0x2e, 0x6b, 0x76, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x75, 0x74, 0x49, 0x6e, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x33, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x6b,
	0x76, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x17, 0x5a, 0x15, 0x6b, 0x76, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x6b, 0x76, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_kvstore_proto_rawDescOnce sync.Once
	file_kvstore_proto_rawDescData []byte
)

func file_kvstore_proto_rawDescGZIP() []byte {
	file_kvstore_proto_rawDescOnce.Do(func() {
		file_kvstore_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_kvstore_proto_rawDesc), len(file_kvstore_proto_rawDesc)))
	})
	return file_kvstore_proto_rawDescData
}

var file_kvstore_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_kvstore_proto_goTypes = []any{
	(*GetIn)(nil),         // 0: kvstore.GetIn
	(*GetOut)(nil),        // 1: kvstore.GetOut
	(*PutIn)(nil),         // 2: kvstore.PutIn
	(*DeleteIn)(nil),      // 3: kvstore.DeleteIn
	(*emptypb.Empty)(nil), // 4: google.protobuf.Empty
}
var file_kvstore_proto_depIdxs = []int32{
	0, // 0: kvstore.KVStore.Get:input_type -> kvstore.GetIn
	2, // 1: kvstore.KVStore.Put:input_type -> kvstore.PutIn
	3, // 2: kvstore.KVStore.Delete:input_type -> kvstore.DeleteIn
	1, // 3: kvstore.KVStore.Get:output_type -> kvstore.GetOut
	4, // 4: kvstore.KVStore.Put:output_type -> google.protobuf.Empty
	4, // 5: kvstore.KVStore.Delete:output_type -> google.protobuf.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_kvstore_proto_init() }
func file_kvstore_proto_init() {
	if File_kvstore_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_kvstore_proto_rawDesc), len(file_kvstore_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kvstore_proto_goTypes,
		DependencyIndexes: file_kvstore_proto_depIdxs,
		MessageInfos:      file_kvstore_proto_msgTypes,
	}.Build()
	File_kvstore_proto = out.File
	file_kvstore_proto_goTypes = nil
	file_kvstore_proto_depIdxs = nil
}
