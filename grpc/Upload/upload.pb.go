// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.1
// source: upload.proto

package Upload

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_upload_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[0]
	if x != nil {
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
	return file_upload_proto_rawDescGZIP(), []int{0}
}

type UploadRequestBody struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Client_IP     string                 `protobuf:"bytes,1,opt,name=client_IP,json=clientIP,proto3" json:"client_IP,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadRequestBody) Reset() {
	*x = UploadRequestBody{}
	mi := &file_upload_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadRequestBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadRequestBody) ProtoMessage() {}

func (x *UploadRequestBody) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadRequestBody.ProtoReflect.Descriptor instead.
func (*UploadRequestBody) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{1}
}

func (x *UploadRequestBody) GetClient_IP() string {
	if x != nil {
		return x.Client_IP
	}
	return ""
}

type UploadResponseBody struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SelectedPort  string                 `protobuf:"bytes,1,opt,name=selected_port,json=selectedPort,proto3" json:"selected_port,omitempty"`
	DataNode_IP   string                 `protobuf:"bytes,2,opt,name=data_node_IP,json=dataNodeIP,proto3" json:"data_node_IP,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadResponseBody) Reset() {
	*x = UploadResponseBody{}
	mi := &file_upload_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadResponseBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadResponseBody) ProtoMessage() {}

func (x *UploadResponseBody) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadResponseBody.ProtoReflect.Descriptor instead.
func (*UploadResponseBody) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{2}
}

func (x *UploadResponseBody) GetSelectedPort() string {
	if x != nil {
		return x.SelectedPort
	}
	return ""
}

func (x *UploadResponseBody) GetDataNode_IP() string {
	if x != nil {
		return x.DataNode_IP
	}
	return ""
}

type UploadFileRequestBody struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileData      []byte                 `protobuf:"bytes,1,opt,name=file_data,json=fileData,proto3" json:"file_data,omitempty"`
	FileName      string                 `protobuf:"bytes,2,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadFileRequestBody) Reset() {
	*x = UploadFileRequestBody{}
	mi := &file_upload_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadFileRequestBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileRequestBody) ProtoMessage() {}

func (x *UploadFileRequestBody) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileRequestBody.ProtoReflect.Descriptor instead.
func (*UploadFileRequestBody) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{3}
}

func (x *UploadFileRequestBody) GetFileData() []byte {
	if x != nil {
		return x.FileData
	}
	return nil
}

func (x *UploadFileRequestBody) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

type NodeMasterAckRequestBody struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        bool                   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NodeMasterAckRequestBody) Reset() {
	*x = NodeMasterAckRequestBody{}
	mi := &file_upload_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NodeMasterAckRequestBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeMasterAckRequestBody) ProtoMessage() {}

func (x *NodeMasterAckRequestBody) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeMasterAckRequestBody.ProtoReflect.Descriptor instead.
func (*NodeMasterAckRequestBody) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{4}
}

func (x *NodeMasterAckRequestBody) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type MasterClientAckRequestBody struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        bool                   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MasterClientAckRequestBody) Reset() {
	*x = MasterClientAckRequestBody{}
	mi := &file_upload_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MasterClientAckRequestBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MasterClientAckRequestBody) ProtoMessage() {}

func (x *MasterClientAckRequestBody) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MasterClientAckRequestBody.ProtoReflect.Descriptor instead.
func (*MasterClientAckRequestBody) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{5}
}

func (x *MasterClientAckRequestBody) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type DownloadPortsRequestBody struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileName      string                 `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DownloadPortsRequestBody) Reset() {
	*x = DownloadPortsRequestBody{}
	mi := &file_upload_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadPortsRequestBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadPortsRequestBody) ProtoMessage() {}

func (x *DownloadPortsRequestBody) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadPortsRequestBody.ProtoReflect.Descriptor instead.
func (*DownloadPortsRequestBody) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{6}
}

func (x *DownloadPortsRequestBody) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

type DownloadPortsResponseBody struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Addresses     []string               `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
	FileSize      int64                  `protobuf:"varint,2,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DownloadPortsResponseBody) Reset() {
	*x = DownloadPortsResponseBody{}
	mi := &file_upload_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadPortsResponseBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadPortsResponseBody) ProtoMessage() {}

func (x *DownloadPortsResponseBody) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadPortsResponseBody.ProtoReflect.Descriptor instead.
func (*DownloadPortsResponseBody) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{7}
}

func (x *DownloadPortsResponseBody) GetAddresses() []string {
	if x != nil {
		return x.Addresses
	}
	return nil
}

func (x *DownloadPortsResponseBody) GetFileSize() int64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

type DownloadFileResponseBody struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileData      []byte                 `protobuf:"bytes,1,opt,name=file_data,json=fileData,proto3" json:"file_data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DownloadFileResponseBody) Reset() {
	*x = DownloadFileResponseBody{}
	mi := &file_upload_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadFileResponseBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadFileResponseBody) ProtoMessage() {}

func (x *DownloadFileResponseBody) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadFileResponseBody.ProtoReflect.Descriptor instead.
func (*DownloadFileResponseBody) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{8}
}

func (x *DownloadFileResponseBody) GetFileData() []byte {
	if x != nil {
		return x.FileData
	}
	return nil
}

type DownloadFileRequestBody struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileName      string                 `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	Start         int64                  `protobuf:"varint,2,opt,name=start,proto3" json:"start,omitempty"`
	End           int64                  `protobuf:"varint,3,opt,name=end,proto3" json:"end,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DownloadFileRequestBody) Reset() {
	*x = DownloadFileRequestBody{}
	mi := &file_upload_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadFileRequestBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadFileRequestBody) ProtoMessage() {}

func (x *DownloadFileRequestBody) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadFileRequestBody.ProtoReflect.Descriptor instead.
func (*DownloadFileRequestBody) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{9}
}

func (x *DownloadFileRequestBody) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *DownloadFileRequestBody) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *DownloadFileRequestBody) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

var File_upload_proto protoreflect.FileDescriptor

const file_upload_proto_rawDesc = "" +
	"\n" +
	"\fupload.proto\x12\x06Upload\"\a\n" +
	"\x05Empty\"0\n" +
	"\x11UploadRequestBody\x12\x1b\n" +
	"\tclient_IP\x18\x01 \x01(\tR\bclientIP\"[\n" +
	"\x12UploadResponseBody\x12#\n" +
	"\rselected_port\x18\x01 \x01(\tR\fselectedPort\x12 \n" +
	"\fdata_node_IP\x18\x02 \x01(\tR\n" +
	"dataNodeIP\"Q\n" +
	"\x15UploadFileRequestBody\x12\x1b\n" +
	"\tfile_data\x18\x01 \x01(\fR\bfileData\x12\x1b\n" +
	"\tfile_name\x18\x02 \x01(\tR\bfileName\"2\n" +
	"\x18NodeMasterAckRequestBody\x12\x16\n" +
	"\x06status\x18\x01 \x01(\bR\x06status\"4\n" +
	"\x1aMasterClientAckRequestBody\x12\x16\n" +
	"\x06status\x18\x01 \x01(\bR\x06status\"7\n" +
	"\x18DownloadPortsRequestBody\x12\x1b\n" +
	"\tfile_name\x18\x01 \x01(\tR\bfileName\"V\n" +
	"\x19DownloadPortsResponseBody\x12\x1c\n" +
	"\taddresses\x18\x01 \x03(\tR\taddresses\x12\x1b\n" +
	"\tfile_size\x18\x02 \x01(\x03R\bfileSize\"7\n" +
	"\x18DownloadFileResponseBody\x12\x1b\n" +
	"\tfile_data\x18\x01 \x01(\fR\bfileData\"^\n" +
	"\x17DownloadFileRequestBody\x12\x1b\n" +
	"\tfile_name\x18\x01 \x01(\tR\bfileName\x12\x14\n" +
	"\x05start\x18\x02 \x01(\x03R\x05start\x12\x10\n" +
	"\x03end\x18\x03 \x01(\x03R\x03end2\xdd\x03\n" +
	"\x03DFS\x12F\n" +
	"\rUploadRequest\x12\x19.Upload.UploadRequestBody\x1a\x1a.Upload.UploadResponseBody\x12A\n" +
	"\x11UploadFileRequest\x12\x1d.Upload.UploadFileRequestBody\x1a\r.Upload.Empty\x12G\n" +
	"\x14NodeMasterAckRequest\x12 .Upload.NodeMasterAckRequestBody\x1a\r.Upload.Empty\x12K\n" +
	"\x16MasterClientAckRequest\x12\".Upload.MasterClientAckRequestBody\x1a\r.Upload.Empty\x12[\n" +
	"\x14DownloadPortsRequest\x12 .Upload.DownloadPortsRequestBody\x1a!.Upload.DownloadPortsResponseBody\x12X\n" +
	"\x13DownloadFileRequest\x12\x1f.Upload.DownloadFileRequestBody\x1a .Upload.DownloadFileResponseBodyB?Z=github.com/RawanMostafa08/Distributed-File-System/grpc/Uploadb\x06proto3"

var (
	file_upload_proto_rawDescOnce sync.Once
	file_upload_proto_rawDescData []byte
)

func file_upload_proto_rawDescGZIP() []byte {
	file_upload_proto_rawDescOnce.Do(func() {
		file_upload_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_upload_proto_rawDesc), len(file_upload_proto_rawDesc)))
	})
	return file_upload_proto_rawDescData
}

var file_upload_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_upload_proto_goTypes = []any{
	(*Empty)(nil),                      // 0: Upload.Empty
	(*UploadRequestBody)(nil),          // 1: Upload.UploadRequestBody
	(*UploadResponseBody)(nil),         // 2: Upload.UploadResponseBody
	(*UploadFileRequestBody)(nil),      // 3: Upload.UploadFileRequestBody
	(*NodeMasterAckRequestBody)(nil),   // 4: Upload.NodeMasterAckRequestBody
	(*MasterClientAckRequestBody)(nil), // 5: Upload.MasterClientAckRequestBody
	(*DownloadPortsRequestBody)(nil),   // 6: Upload.DownloadPortsRequestBody
	(*DownloadPortsResponseBody)(nil),  // 7: Upload.DownloadPortsResponseBody
	(*DownloadFileResponseBody)(nil),   // 8: Upload.DownloadFileResponseBody
	(*DownloadFileRequestBody)(nil),    // 9: Upload.DownloadFileRequestBody
}
var file_upload_proto_depIdxs = []int32{
	1, // 0: Upload.DFS.UploadRequest:input_type -> Upload.UploadRequestBody
	3, // 1: Upload.DFS.UploadFileRequest:input_type -> Upload.UploadFileRequestBody
	4, // 2: Upload.DFS.NodeMasterAckRequest:input_type -> Upload.NodeMasterAckRequestBody
	5, // 3: Upload.DFS.MasterClientAckRequest:input_type -> Upload.MasterClientAckRequestBody
	6, // 4: Upload.DFS.DownloadPortsRequest:input_type -> Upload.DownloadPortsRequestBody
	9, // 5: Upload.DFS.DownloadFileRequest:input_type -> Upload.DownloadFileRequestBody
	2, // 6: Upload.DFS.UploadRequest:output_type -> Upload.UploadResponseBody
	0, // 7: Upload.DFS.UploadFileRequest:output_type -> Upload.Empty
	0, // 8: Upload.DFS.NodeMasterAckRequest:output_type -> Upload.Empty
	0, // 9: Upload.DFS.MasterClientAckRequest:output_type -> Upload.Empty
	7, // 10: Upload.DFS.DownloadPortsRequest:output_type -> Upload.DownloadPortsResponseBody
	8, // 11: Upload.DFS.DownloadFileRequest:output_type -> Upload.DownloadFileResponseBody
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_upload_proto_init() }
func file_upload_proto_init() {
	if File_upload_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_upload_proto_rawDesc), len(file_upload_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_upload_proto_goTypes,
		DependencyIndexes: file_upload_proto_depIdxs,
		MessageInfos:      file_upload_proto_msgTypes,
	}.Build()
	File_upload_proto = out.File
	file_upload_proto_goTypes = nil
	file_upload_proto_depIdxs = nil
}
