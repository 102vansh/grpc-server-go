// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: proto/greet.proto

package proto

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

type Noparams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Noparams) Reset() {
	*x = Noparams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_greet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Noparams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Noparams) ProtoMessage() {}

func (x *Noparams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_greet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Noparams.ProtoReflect.Descriptor instead.
func (*Noparams) Descriptor() ([]byte, []int) {
	return file_proto_greet_proto_rawDescGZIP(), []int{0}
}

type Hellorequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Hellorequest) Reset() {
	*x = Hellorequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_greet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hellorequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hellorequest) ProtoMessage() {}

func (x *Hellorequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_greet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hellorequest.ProtoReflect.Descriptor instead.
func (*Hellorequest) Descriptor() ([]byte, []int) {
	return file_proto_greet_proto_rawDescGZIP(), []int{1}
}

func (x *Hellorequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Helloresponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Helloresponse) Reset() {
	*x = Helloresponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_greet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Helloresponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Helloresponse) ProtoMessage() {}

func (x *Helloresponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_greet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Helloresponse.ProtoReflect.Descriptor instead.
func (*Helloresponse) Descriptor() ([]byte, []int) {
	return file_proto_greet_proto_rawDescGZIP(), []int{2}
}

func (x *Helloresponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Messagelist struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []string `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *Messagelist) Reset() {
	*x = Messagelist{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_greet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Messagelist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Messagelist) ProtoMessage() {}

func (x *Messagelist) ProtoReflect() protoreflect.Message {
	mi := &file_proto_greet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Messagelist.ProtoReflect.Descriptor instead.
func (*Messagelist) Descriptor() ([]byte, []int) {
	return file_proto_greet_proto_rawDescGZIP(), []int{3}
}

func (x *Messagelist) GetMessages() []string {
	if x != nil {
		return x.Messages
	}
	return nil
}

type Namelist struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Names []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
}

func (x *Namelist) Reset() {
	*x = Namelist{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_greet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Namelist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Namelist) ProtoMessage() {}

func (x *Namelist) ProtoReflect() protoreflect.Message {
	mi := &file_proto_greet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Namelist.ProtoReflect.Descriptor instead.
func (*Namelist) Descriptor() ([]byte, []int) {
	return file_proto_greet_proto_rawDescGZIP(), []int{4}
}

func (x *Namelist) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

var File_proto_greet_proto protoreflect.FileDescriptor

var file_proto_greet_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x67, 0x72, 0x65, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x22, 0x0a, 0x0a, 0x08, 0x4e, 0x6f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x22,
	0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x29, 0x0a, 0x0d, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x29, 0x0a,
	0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x20, 0x0a, 0x08, 0x4e, 0x61, 0x6d, 0x65,
	0x6c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x32, 0xdd, 0x02, 0x0a, 0x0d, 0x47,
	0x72, 0x65, 0x65, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x08,
	0x53, 0x61, 0x79, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x17, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4e, 0x6f, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x1a, 0x1c, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x52, 0x0a, 0x17, 0x53, 0x61, 0x79, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x17, 0x2e, 0x67, 0x72, 0x65,
	0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x6c,
	0x69, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x30, 0x01, 0x12, 0x54, 0x0a, 0x17, 0x53, 0x61, 0x79, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x1b,
	0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x72,
	0x65, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x28, 0x01, 0x12, 0x5f, 0x0a, 0x1e, 0x53, 0x61, 0x79,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x62, 0x69, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x1b, 0x2e, 0x67, 0x72,
	0x65, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_greet_proto_rawDescOnce sync.Once
	file_proto_greet_proto_rawDescData = file_proto_greet_proto_rawDesc
)

func file_proto_greet_proto_rawDescGZIP() []byte {
	file_proto_greet_proto_rawDescOnce.Do(func() {
		file_proto_greet_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_greet_proto_rawDescData)
	})
	return file_proto_greet_proto_rawDescData
}

var file_proto_greet_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_greet_proto_goTypes = []interface{}{
	(*Noparams)(nil),      // 0: greet_service.Noparams
	(*Hellorequest)(nil),  // 1: greet_service.Hellorequest
	(*Helloresponse)(nil), // 2: greet_service.Helloresponse
	(*Messagelist)(nil),   // 3: greet_service.messagelist
	(*Namelist)(nil),      // 4: greet_service.Namelist
}
var file_proto_greet_proto_depIdxs = []int32{
	0, // 0: greet_service.GreetsService.Sayhello:input_type -> greet_service.Noparams
	4, // 1: greet_service.GreetsService.Sayhelloserverstreaming:input_type -> greet_service.Namelist
	1, // 2: greet_service.GreetsService.Sayhelloclientstreaming:input_type -> greet_service.Hellorequest
	1, // 3: greet_service.GreetsService.Sayhellobidirectionalstreaming:input_type -> greet_service.Hellorequest
	2, // 4: greet_service.GreetsService.Sayhello:output_type -> greet_service.Helloresponse
	2, // 5: greet_service.GreetsService.Sayhelloserverstreaming:output_type -> greet_service.Helloresponse
	3, // 6: greet_service.GreetsService.Sayhelloclientstreaming:output_type -> greet_service.messagelist
	2, // 7: greet_service.GreetsService.Sayhellobidirectionalstreaming:output_type -> greet_service.Helloresponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_greet_proto_init() }
func file_proto_greet_proto_init() {
	if File_proto_greet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_greet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Noparams); i {
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
		file_proto_greet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hellorequest); i {
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
		file_proto_greet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Helloresponse); i {
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
		file_proto_greet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Messagelist); i {
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
		file_proto_greet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Namelist); i {
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
			RawDescriptor: file_proto_greet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_greet_proto_goTypes,
		DependencyIndexes: file_proto_greet_proto_depIdxs,
		MessageInfos:      file_proto_greet_proto_msgTypes,
	}.Build()
	File_proto_greet_proto = out.File
	file_proto_greet_proto_rawDesc = nil
	file_proto_greet_proto_goTypes = nil
	file_proto_greet_proto_depIdxs = nil
}
