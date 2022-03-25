// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: cl.proto

//登录服～http协议

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

//角色简单信息
type SimpleRole struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid  uint64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Tid  uint32 `protobuf:"varint,3,opt,name=tid,proto3" json:"tid,omitempty"` //模版id
	Exp  uint64 `protobuf:"varint,4,opt,name=exp,proto3" json:"exp,omitempty"` //经验
}

func (x *SimpleRole) Reset() {
	*x = SimpleRole{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cl_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleRole) ProtoMessage() {}

func (x *SimpleRole) ProtoReflect() protoreflect.Message {
	mi := &file_cl_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleRole.ProtoReflect.Descriptor instead.
func (*SimpleRole) Descriptor() ([]byte, []int) {
	return file_cl_proto_rawDescGZIP(), []int{0}
}

func (x *SimpleRole) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *SimpleRole) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SimpleRole) GetTid() uint32 {
	if x != nil {
		return x.Tid
	}
	return 0
}

func (x *SimpleRole) GetExp() uint64 {
	if x != nil {
		return x.Exp
	}
	return 0
}

//登录验证+返回角色列表 CL_Login LC_Login
type CL_Login struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *CL_Login) Reset() {
	*x = CL_Login{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cl_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CL_Login) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CL_Login) ProtoMessage() {}

func (x *CL_Login) ProtoReflect() protoreflect.Message {
	mi := &file_cl_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CL_Login.ProtoReflect.Descriptor instead.
func (*CL_Login) Descriptor() ([]byte, []int) {
	return file_cl_proto_rawDescGZIP(), []int{1}
}

func (x *CL_Login) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type LC_Login struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleList []*SimpleRole `protobuf:"bytes,1,rep,name=roleList,proto3" json:"roleList,omitempty"` //角色列表
}

func (x *LC_Login) Reset() {
	*x = LC_Login{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cl_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LC_Login) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LC_Login) ProtoMessage() {}

func (x *LC_Login) ProtoReflect() protoreflect.Message {
	mi := &file_cl_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LC_Login.ProtoReflect.Descriptor instead.
func (*LC_Login) Descriptor() ([]byte, []int) {
	return file_cl_proto_rawDescGZIP(), []int{2}
}

func (x *LC_Login) GetRoleList() []*SimpleRole {
	if x != nil {
		return x.RoleList
	}
	return nil
}

var File_cl_proto protoreflect.FileDescriptor

var file_cl_proto_rawDesc = []byte{
	0x0a, 0x08, 0x63, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x56, 0x0a, 0x0a, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x6f, 0x6c,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03,
	0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x74, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x70,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x65, 0x78, 0x70, 0x22, 0x20, 0x0a, 0x08, 0x43,
	0x4c, 0x5f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3b, 0x0a,
	0x08, 0x4c, 0x43, 0x5f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x2f, 0x0a, 0x08, 0x72, 0x6f, 0x6c,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x6f, 0x6c, 0x65,
	0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x0a, 0x5a, 0x08, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cl_proto_rawDescOnce sync.Once
	file_cl_proto_rawDescData = file_cl_proto_rawDesc
)

func file_cl_proto_rawDescGZIP() []byte {
	file_cl_proto_rawDescOnce.Do(func() {
		file_cl_proto_rawDescData = protoimpl.X.CompressGZIP(file_cl_proto_rawDescData)
	})
	return file_cl_proto_rawDescData
}

var file_cl_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cl_proto_goTypes = []interface{}{
	(*SimpleRole)(nil), // 0: message.SimpleRole
	(*CL_Login)(nil),   // 1: message.CL_Login
	(*LC_Login)(nil),   // 2: message.LC_Login
}
var file_cl_proto_depIdxs = []int32{
	0, // 0: message.LC_Login.roleList:type_name -> message.SimpleRole
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cl_proto_init() }
func file_cl_proto_init() {
	if File_cl_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cl_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleRole); i {
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
		file_cl_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CL_Login); i {
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
		file_cl_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LC_Login); i {
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
			RawDescriptor: file_cl_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cl_proto_goTypes,
		DependencyIndexes: file_cl_proto_depIdxs,
		MessageInfos:      file_cl_proto_msgTypes,
	}.Build()
	File_cl_proto = out.File
	file_cl_proto_rawDesc = nil
	file_cl_proto_goTypes = nil
	file_cl_proto_depIdxs = nil
}
