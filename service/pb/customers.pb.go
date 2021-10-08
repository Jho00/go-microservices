// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: customers.proto

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

type GetCustomersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCustomersRequest) Reset() {
	*x = GetCustomersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomersRequest) ProtoMessage() {}

func (x *GetCustomersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomersRequest.ProtoReflect.Descriptor instead.
func (*GetCustomersRequest) Descriptor() ([]byte, []int) {
	return file_customers_proto_rawDescGZIP(), []int{0}
}

type GetCustomersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Customers []*Customer `protobuf:"bytes,1,rep,name=customers,proto3" json:"customers,omitempty"`
}

func (x *GetCustomersResponse) Reset() {
	*x = GetCustomersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomersResponse) ProtoMessage() {}

func (x *GetCustomersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomersResponse.ProtoReflect.Descriptor instead.
func (*GetCustomersResponse) Descriptor() ([]byte, []int) {
	return file_customers_proto_rawDescGZIP(), []int{1}
}

func (x *GetCustomersResponse) GetCustomers() []*Customer {
	if x != nil {
		return x.Customers
	}
	return nil
}

type Customer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Customerid   int32  `protobuf:"varint,1,opt,name=Customerid,proto3" json:"Customerid,omitempty"`
	Customername string `protobuf:"bytes,2,opt,name=Customername,proto3" json:"Customername,omitempty"`
}

func (x *Customer) Reset() {
	*x = Customer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Customer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Customer) ProtoMessage() {}

func (x *Customer) ProtoReflect() protoreflect.Message {
	mi := &file_customers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Customer.ProtoReflect.Descriptor instead.
func (*Customer) Descriptor() ([]byte, []int) {
	return file_customers_proto_rawDescGZIP(), []int{2}
}

func (x *Customer) GetCustomerid() int32 {
	if x != nil {
		return x.Customerid
	}
	return 0
}

func (x *Customer) GetCustomername() string {
	if x != nil {
		return x.Customername
	}
	return ""
}

var File_customers_proto protoreflect.FileDescriptor

var file_customers_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x15, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x42, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x09, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x09, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73,
	0x22, 0x4e, 0x0a, 0x08, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x0c,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x32, 0x56, 0x0a, 0x0f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x73, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3d, 0x0a, 0x13, 0x69, 0x6f, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x42,
	0x0f, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x13, 0x67, 0x6f, 0x2d, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_customers_proto_rawDescOnce sync.Once
	file_customers_proto_rawDescData = file_customers_proto_rawDesc
)

func file_customers_proto_rawDescGZIP() []byte {
	file_customers_proto_rawDescOnce.Do(func() {
		file_customers_proto_rawDescData = protoimpl.X.CompressGZIP(file_customers_proto_rawDescData)
	})
	return file_customers_proto_rawDescData
}

var file_customers_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_customers_proto_goTypes = []interface{}{
	(*GetCustomersRequest)(nil),  // 0: pb.GetCustomersRequest
	(*GetCustomersResponse)(nil), // 1: pb.GetCustomersResponse
	(*Customer)(nil),             // 2: pb.Customer
}
var file_customers_proto_depIdxs = []int32{
	2, // 0: pb.GetCustomersResponse.customers:type_name -> pb.Customer
	0, // 1: pb.CustomerService.GetCustomers:input_type -> pb.GetCustomersRequest
	1, // 2: pb.CustomerService.GetCustomers:output_type -> pb.GetCustomersResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_customers_proto_init() }
func file_customers_proto_init() {
	if File_customers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_customers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomersRequest); i {
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
		file_customers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomersResponse); i {
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
		file_customers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Customer); i {
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
			RawDescriptor: file_customers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_customers_proto_goTypes,
		DependencyIndexes: file_customers_proto_depIdxs,
		MessageInfos:      file_customers_proto_msgTypes,
	}.Build()
	File_customers_proto = out.File
	file_customers_proto_rawDesc = nil
	file_customers_proto_goTypes = nil
	file_customers_proto_depIdxs = nil
}