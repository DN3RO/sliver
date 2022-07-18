// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: commonpb/common.proto

package commonpb

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonpb_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_commonpb_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_commonpb_common_proto_rawDescGZIP(), []int{0}
}

// Request - Common fields used in all gRPC requests
type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Async     bool   `protobuf:"varint,1,opt,name=Async,proto3" json:"Async,omitempty"`
	Timeout   int64  `protobuf:"varint,2,opt,name=Timeout,proto3" json:"Timeout,omitempty"`
	BeaconID  string `protobuf:"bytes,8,opt,name=BeaconID,proto3" json:"BeaconID,omitempty"`
	SessionID string `protobuf:"bytes,9,opt,name=SessionID,proto3" json:"SessionID,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonpb_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_commonpb_common_proto_msgTypes[1]
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
	return file_commonpb_common_proto_rawDescGZIP(), []int{1}
}

func (x *Request) GetAsync() bool {
	if x != nil {
		return x.Async
	}
	return false
}

func (x *Request) GetTimeout() int64 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *Request) GetBeaconID() string {
	if x != nil {
		return x.BeaconID
	}
	return ""
}

func (x *Request) GetSessionID() string {
	if x != nil {
		return x.SessionID
	}
	return ""
}

// Response - Common fields used in all gRPC responses. Note that the Err field
//            only used when the implant needs to return an error to the server.
//            Client<->Server comms should use normal gRPC error handling.
type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Err      string `protobuf:"bytes,1,opt,name=Err,proto3" json:"Err,omitempty"`
	Async    bool   `protobuf:"varint,2,opt,name=Async,proto3" json:"Async,omitempty"`
	BeaconID string `protobuf:"bytes,8,opt,name=BeaconID,proto3" json:"BeaconID,omitempty"`
	TaskID   string `protobuf:"bytes,9,opt,name=TaskID,proto3" json:"TaskID,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonpb_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_commonpb_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_commonpb_common_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

func (x *Response) GetAsync() bool {
	if x != nil {
		return x.Async
	}
	return false
}

func (x *Response) GetBeaconID() string {
	if x != nil {
		return x.BeaconID
	}
	return ""
}

func (x *Response) GetTaskID() string {
	if x != nil {
		return x.TaskID
	}
	return ""
}

// File - A basic file data type
type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonpb_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_commonpb_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_commonpb_common_proto_rawDescGZIP(), []int{3}
}

func (x *File) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *File) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// Process - A basic process data type
type Process struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid          int32    `protobuf:"varint,1,opt,name=Pid,proto3" json:"Pid,omitempty"`
	Ppid         int32    `protobuf:"varint,2,opt,name=Ppid,proto3" json:"Ppid,omitempty"`
	Executable   string   `protobuf:"bytes,3,opt,name=Executable,proto3" json:"Executable,omitempty"`
	Owner        string   `protobuf:"bytes,4,opt,name=Owner,proto3" json:"Owner,omitempty"`
	Architecture string   `protobuf:"bytes,7,opt,name=Architecture,proto3" json:"Architecture,omitempty"`
	SessionID    int32    `protobuf:"varint,5,opt,name=SessionID,proto3" json:"SessionID,omitempty"`
	CmdLine      []string `protobuf:"bytes,6,rep,name=CmdLine,proto3" json:"CmdLine,omitempty"`
}

func (x *Process) Reset() {
	*x = Process{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonpb_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Process) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Process) ProtoMessage() {}

func (x *Process) ProtoReflect() protoreflect.Message {
	mi := &file_commonpb_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Process.ProtoReflect.Descriptor instead.
func (*Process) Descriptor() ([]byte, []int) {
	return file_commonpb_common_proto_rawDescGZIP(), []int{4}
}

func (x *Process) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *Process) GetPpid() int32 {
	if x != nil {
		return x.Ppid
	}
	return 0
}

func (x *Process) GetExecutable() string {
	if x != nil {
		return x.Executable
	}
	return ""
}

func (x *Process) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *Process) GetArchitecture() string {
	if x != nil {
		return x.Architecture
	}
	return ""
}

func (x *Process) GetSessionID() int32 {
	if x != nil {
		return x.SessionID
	}
	return 0
}

func (x *Process) GetCmdLine() []string {
	if x != nil {
		return x.CmdLine
	}
	return nil
}

// EnvVar - Environment variable K/V
type EnvVar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *EnvVar) Reset() {
	*x = EnvVar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonpb_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnvVar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvVar) ProtoMessage() {}

func (x *EnvVar) ProtoReflect() protoreflect.Message {
	mi := &file_commonpb_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnvVar.ProtoReflect.Descriptor instead.
func (*EnvVar) Descriptor() ([]byte, []int) {
	return file_commonpb_common_proto_rawDescGZIP(), []int{5}
}

func (x *EnvVar) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *EnvVar) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_commonpb_common_proto protoreflect.FileDescriptor

var file_commonpb_common_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70,
	0x62, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x73, 0x0a, 0x07, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x73, 0x79, 0x6e, 0x63, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x41, 0x73, 0x79, 0x6e, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x54,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x54, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x49,
	0x44, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x49,
	0x44, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x22,
	0x66, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x45,
	0x72, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x45, 0x72, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x41, 0x73, 0x79, 0x6e, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x41, 0x73,
	0x79, 0x6e, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x49, 0x44, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x49, 0x44, 0x12,
	0x16, 0x0a, 0x06, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x22, 0x2e, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0xc1, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x50, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x50, 0x70, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x4f, 0x77, 0x6e,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12,
	0x22, 0x0a, 0x0c, 0x41, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x41, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6d, 0x64, 0x4c, 0x69, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x43, 0x6d, 0x64, 0x4c, 0x69, 0x6e, 0x65, 0x22, 0x30, 0x0a, 0x06, 0x45,
	0x6e, 0x76, 0x56, 0x61, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x2f, 0x5a,
	0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x69, 0x73, 0x68,
	0x6f, 0x70, 0x66, 0x6f, 0x78, 0x2f, 0x73, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commonpb_common_proto_rawDescOnce sync.Once
	file_commonpb_common_proto_rawDescData = file_commonpb_common_proto_rawDesc
)

func file_commonpb_common_proto_rawDescGZIP() []byte {
	file_commonpb_common_proto_rawDescOnce.Do(func() {
		file_commonpb_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonpb_common_proto_rawDescData)
	})
	return file_commonpb_common_proto_rawDescData
}

var file_commonpb_common_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_commonpb_common_proto_goTypes = []interface{}{
	(*Empty)(nil),    // 0: commonpb.Empty
	(*Request)(nil),  // 1: commonpb.Request
	(*Response)(nil), // 2: commonpb.Response
	(*File)(nil),     // 3: commonpb.File
	(*Process)(nil),  // 4: commonpb.Process
	(*EnvVar)(nil),   // 5: commonpb.EnvVar
}
var file_commonpb_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_commonpb_common_proto_init() }
func file_commonpb_common_proto_init() {
	if File_commonpb_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_commonpb_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_commonpb_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_commonpb_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_commonpb_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
		file_commonpb_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Process); i {
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
		file_commonpb_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnvVar); i {
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
			RawDescriptor: file_commonpb_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_commonpb_common_proto_goTypes,
		DependencyIndexes: file_commonpb_common_proto_depIdxs,
		MessageInfos:      file_commonpb_common_proto_msgTypes,
	}.Build()
	File_commonpb_common_proto = out.File
	file_commonpb_common_proto_rawDesc = nil
	file_commonpb_common_proto_goTypes = nil
	file_commonpb_common_proto_depIdxs = nil
}
