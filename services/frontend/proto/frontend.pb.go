// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: frontend.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Not a replacement for proper monitoring - just a simple load
// monitor for the dashboard. For proper monitoring use
// graphana/prometheus.
type Metrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProcessCpuNanoSecondsTotal int64 `protobuf:"varint,1,opt,name=process_cpu_nano_seconds_total,json=processCpuNanoSecondsTotal,proto3" json:"process_cpu_nano_seconds_total,omitempty"`
	// Percent of one core used by this frontend process.
	CpuLoadPercent                int64 `protobuf:"varint,4,opt,name=cpu_load_percent,json=cpuLoadPercent,proto3" json:"cpu_load_percent,omitempty"`
	ClientCommsCurrentConnections int64 `protobuf:"varint,2,opt,name=client_comms_current_connections,json=clientCommsCurrentConnections,proto3" json:"client_comms_current_connections,omitempty"`
	ProcessResidentMemoryBytes    int64 `protobuf:"varint,3,opt,name=process_resident_memory_bytes,json=processResidentMemoryBytes,proto3" json:"process_resident_memory_bytes,omitempty"`
}

func (x *Metrics) Reset() {
	*x = Metrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_frontend_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metrics) ProtoMessage() {}

func (x *Metrics) ProtoReflect() protoreflect.Message {
	mi := &file_frontend_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metrics.ProtoReflect.Descriptor instead.
func (*Metrics) Descriptor() ([]byte, []int) {
	return file_frontend_proto_rawDescGZIP(), []int{0}
}

func (x *Metrics) GetProcessCpuNanoSecondsTotal() int64 {
	if x != nil {
		return x.ProcessCpuNanoSecondsTotal
	}
	return 0
}

func (x *Metrics) GetCpuLoadPercent() int64 {
	if x != nil {
		return x.CpuLoadPercent
	}
	return 0
}

func (x *Metrics) GetClientCommsCurrentConnections() int64 {
	if x != nil {
		return x.ClientCommsCurrentConnections
	}
	return 0
}

func (x *Metrics) GetProcessResidentMemoryBytes() int64 {
	if x != nil {
		return x.ProcessResidentMemoryBytes
	}
	return 0
}

// Describe the frontend state.
type FrontendState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique name of the frontend.
	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version string `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	// Updated periodically to indicate the frontend is alive (in
	// nanoseconds).
	Heartbeat int64 `protobuf:"varint,2,opt,name=heartbeat,proto3" json:"heartbeat,omitempty"`
	// The URL over which the frontend is accessible to clients.
	Url     string   `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Metrics *Metrics `protobuf:"bytes,4,opt,name=metrics,proto3" json:"metrics,omitempty"`
	// If set other frontends will not redirect to this frontend
	// automatically. This makes sense if you have an external load
	// balancer that will redirect traffic to this node.
	DoNotRedirect bool `protobuf:"varint,6,opt,name=do_not_redirect,json=doNotRedirect,proto3" json:"do_not_redirect,omitempty"`
}

func (x *FrontendState) Reset() {
	*x = FrontendState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_frontend_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FrontendState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrontendState) ProtoMessage() {}

func (x *FrontendState) ProtoReflect() protoreflect.Message {
	mi := &file_frontend_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrontendState.ProtoReflect.Descriptor instead.
func (*FrontendState) Descriptor() ([]byte, []int) {
	return file_frontend_proto_rawDescGZIP(), []int{1}
}

func (x *FrontendState) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FrontendState) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *FrontendState) GetHeartbeat() int64 {
	if x != nil {
		return x.Heartbeat
	}
	return 0
}

func (x *FrontendState) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *FrontendState) GetMetrics() *Metrics {
	if x != nil {
		return x.Metrics
	}
	return nil
}

func (x *FrontendState) GetDoNotRedirect() bool {
	if x != nil {
		return x.DoNotRedirect
	}
	return false
}

var File_frontend_proto protoreflect.FileDescriptor

var file_frontend_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x02, 0x0a, 0x07, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x12, 0x42, 0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x63,
	0x70, 0x75, 0x5f, 0x6e, 0x61, 0x6e, 0x6f, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x5f,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1a, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x43, 0x70, 0x75, 0x4e, 0x61, 0x6e, 0x6f, 0x53, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x73, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x70, 0x75, 0x5f, 0x6c,
	0x6f, 0x61, 0x64, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0e, 0x63, 0x70, 0x75, 0x4c, 0x6f, 0x61, 0x64, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x12, 0x47, 0x0a, 0x20, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x6d,
	0x73, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1d, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x73, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x41, 0x0a, 0x1d, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x6d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x1a, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x42, 0x79, 0x74, 0x65, 0x73, 0x22, 0xbf, 0x01,
	0x0a, 0x0d, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a,
	0x09, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x28, 0x0a,
	0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x07,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x64, 0x6f, 0x5f, 0x6e, 0x6f,
	0x74, 0x5f, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0d, 0x64, 0x6f, 0x4e, 0x6f, 0x74, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x42,
	0x3f, 0x5a, 0x3d, 0x77, 0x77, 0x77, 0x2e, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x76, 0x65, 0x6c, 0x6f,
	0x63, 0x69, 0x72, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_frontend_proto_rawDescOnce sync.Once
	file_frontend_proto_rawDescData = file_frontend_proto_rawDesc
)

func file_frontend_proto_rawDescGZIP() []byte {
	file_frontend_proto_rawDescOnce.Do(func() {
		file_frontend_proto_rawDescData = protoimpl.X.CompressGZIP(file_frontend_proto_rawDescData)
	})
	return file_frontend_proto_rawDescData
}

var file_frontend_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_frontend_proto_goTypes = []interface{}{
	(*Metrics)(nil),       // 0: proto.Metrics
	(*FrontendState)(nil), // 1: proto.FrontendState
}
var file_frontend_proto_depIdxs = []int32{
	0, // 0: proto.FrontendState.metrics:type_name -> proto.Metrics
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_frontend_proto_init() }
func file_frontend_proto_init() {
	if File_frontend_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_frontend_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metrics); i {
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
		file_frontend_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FrontendState); i {
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
			RawDescriptor: file_frontend_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_frontend_proto_goTypes,
		DependencyIndexes: file_frontend_proto_depIdxs,
		MessageInfos:      file_frontend_proto_msgTypes,
	}.Build()
	File_frontend_proto = out.File
	file_frontend_proto_rawDesc = nil
	file_frontend_proto_goTypes = nil
	file_frontend_proto_depIdxs = nil
}
