// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        v3.19.4
// source: sim.proto

package sim

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

// The braodcast message
type BroadcastTxn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *BroadcastTxn) Reset() {
	*x = BroadcastTxn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sim_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastTxn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastTxn) ProtoMessage() {}

func (x *BroadcastTxn) ProtoReflect() protoreflect.Message {
	mi := &file_sim_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastTxn.ProtoReflect.Descriptor instead.
func (*BroadcastTxn) Descriptor() ([]byte, []int) {
	return file_sim_proto_rawDescGZIP(), []int{0}
}

func (x *BroadcastTxn) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type AcknowledgementTxn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *AcknowledgementTxn) Reset() {
	*x = AcknowledgementTxn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sim_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcknowledgementTxn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcknowledgementTxn) ProtoMessage() {}

func (x *AcknowledgementTxn) ProtoReflect() protoreflect.Message {
	mi := &file_sim_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcknowledgementTxn.ProtoReflect.Descriptor instead.
func (*AcknowledgementTxn) Descriptor() ([]byte, []int) {
	return file_sim_proto_rawDescGZIP(), []int{1}
}

func (x *AcknowledgementTxn) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_sim_proto protoreflect.FileDescriptor

var file_sim_proto_rawDesc = []byte{
	0x0a, 0x09, 0x73, 0x69, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x73, 0x69, 0x6d,
	0x22, 0x28, 0x0a, 0x0c, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x54, 0x78, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x2e, 0x0a, 0x12, 0x41, 0x63,
	0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x78, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32, 0x47, 0x0a, 0x0a, 0x53, 0x69,
	0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x09, 0x42, 0x72, 0x6f, 0x61,
	0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x11, 0x2e, 0x73, 0x69, 0x6d, 0x2e, 0x42, 0x72, 0x6f, 0x61,
	0x64, 0x63, 0x61, 0x73, 0x74, 0x54, 0x78, 0x6e, 0x1a, 0x17, 0x2e, 0x73, 0x69, 0x6d, 0x2e, 0x41,
	0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x78,
	0x6e, 0x22, 0x00, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_sim_proto_rawDescOnce sync.Once
	file_sim_proto_rawDescData = file_sim_proto_rawDesc
)

func file_sim_proto_rawDescGZIP() []byte {
	file_sim_proto_rawDescOnce.Do(func() {
		file_sim_proto_rawDescData = protoimpl.X.CompressGZIP(file_sim_proto_rawDescData)
	})
	return file_sim_proto_rawDescData
}

var file_sim_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sim_proto_goTypes = []interface{}{
	(*BroadcastTxn)(nil),       // 0: sim.BroadcastTxn
	(*AcknowledgementTxn)(nil), // 1: sim.AcknowledgementTxn
}
var file_sim_proto_depIdxs = []int32{
	0, // 0: sim.Simulation.Broadcast:input_type -> sim.BroadcastTxn
	1, // 1: sim.Simulation.Broadcast:output_type -> sim.AcknowledgementTxn
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sim_proto_init() }
func file_sim_proto_init() {
	if File_sim_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sim_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastTxn); i {
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
		file_sim_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcknowledgementTxn); i {
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
			RawDescriptor: file_sim_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sim_proto_goTypes,
		DependencyIndexes: file_sim_proto_depIdxs,
		MessageInfos:      file_sim_proto_msgTypes,
	}.Build()
	File_sim_proto = out.File
	file_sim_proto_rawDesc = nil
	file_sim_proto_goTypes = nil
	file_sim_proto_depIdxs = nil
}
