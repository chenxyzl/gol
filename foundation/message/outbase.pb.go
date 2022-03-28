// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: outbase.proto

package message

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sn   uint32 `protobuf:"varint,1,opt,name=sn,proto3" json:"sn,omitempty"`    //流水号
	Cmd  uint32 `protobuf:"varint,2,opt,name=cmd,proto3" json:"cmd,omitempty"`  //rpc的id
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"` //数据 //第一个消息必须是登录。gate会保存里面的uid
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_outbase_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_outbase_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_outbase_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetSn() uint32 {
	if x != nil {
		return x.Sn
	}
	return 0
}

func (x *Request) GetCmd() uint32 {
	if x != nil {
		return x.Cmd
	}
	return 0
}

func (x *Request) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sn   uint32 `protobuf:"varint,1,opt,name=sn,proto3" json:"sn,omitempty"`                       //流水号
	Cmd  uint32 `protobuf:"varint,2,opt,name=cmd,proto3" json:"cmd,omitempty"`                     //rpc的id
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`                    //数据
	Code Code   `protobuf:"varint,4,opt,name=code,proto3,enum=message.Code" json:"code,omitempty"` //错误码
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_outbase_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_outbase_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_outbase_proto_rawDescGZIP(), []int{1}
}

func (x *Reply) GetSn() uint32 {
	if x != nil {
		return x.Sn
	}
	return 0
}

func (x *Reply) GetCmd() uint32 {
	if x != nil {
		return x.Cmd
	}
	return 0
}

func (x *Reply) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Reply) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_OK
}

type Notify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sn   uint32 `protobuf:"varint,1,opt,name=sn,proto3" json:"sn,omitempty"`    //流水号
	Cmd  uint32 `protobuf:"varint,2,opt,name=cmd,proto3" json:"cmd,omitempty"`  //rpc的id
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"` //数据
}

func (x *Notify) Reset() {
	*x = Notify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_outbase_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notify) ProtoMessage() {}

func (x *Notify) ProtoReflect() protoreflect.Message {
	mi := &file_outbase_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notify.ProtoReflect.Descriptor instead.
func (*Notify) Descriptor() ([]byte, []int) {
	return file_outbase_proto_rawDescGZIP(), []int{2}
}

func (x *Notify) GetSn() uint32 {
	if x != nil {
		return x.Sn
	}
	return 0
}

func (x *Notify) GetCmd() uint32 {
	if x != nil {
		return x.Cmd
	}
	return 0
}

func (x *Notify) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_outbase_proto protoreflect.FileDescriptor

var file_outbase_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6f, 0x75, 0x74, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0a, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x73, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x73, 0x6e, 0x12,
	0x10, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x63, 0x6d,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x60, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x73, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x73, 0x6e, 0x12, 0x10,
	0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x63, 0x6d, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x21, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x3e, 0x0a, 0x06, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x73, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x73,
	0x6e, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03,
	0x63, 0x6d, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0a, 0x5a, 0x08, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_outbase_proto_rawDescOnce sync.Once
	file_outbase_proto_rawDescData = file_outbase_proto_rawDesc
)

func file_outbase_proto_rawDescGZIP() []byte {
	file_outbase_proto_rawDescOnce.Do(func() {
		file_outbase_proto_rawDescData = protoimpl.X.CompressGZIP(file_outbase_proto_rawDescData)
	})
	return file_outbase_proto_rawDescData
}

var file_outbase_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_outbase_proto_goTypes = []interface{}{
	(*Request)(nil), // 0: message.Request
	(*Reply)(nil),   // 1: message.Reply
	(*Notify)(nil),  // 2: message.Notify
	(Code)(0),       // 3: message.Code
}
var file_outbase_proto_depIdxs = []int32{
	3, // 0: message.Reply.code:type_name -> message.Code
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_outbase_proto_init() }
func file_outbase_proto_init() {
	if File_outbase_proto != nil {
		return
	}
	file_code_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_outbase_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_outbase_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
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
		file_outbase_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notify); i {
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
			RawDescriptor: file_outbase_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_outbase_proto_goTypes,
		DependencyIndexes: file_outbase_proto_depIdxs,
		MessageInfos:      file_outbase_proto_msgTypes,
	}.Build()
	File_outbase_proto = out.File
	file_outbase_proto_rawDesc = nil
	file_outbase_proto_goTypes = nil
	file_outbase_proto_depIdxs = nil
}
