// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        (unknown)
// source: health.proto

package v1health

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

type ConnectionStatus int32

const (
	ConnectionStatus_Disconnected ConnectionStatus = 0 // internet is disconneced
	ConnectionStatus_Active       ConnectionStatus = 1 // internet is connected
)

// Enum value maps for ConnectionStatus.
var (
	ConnectionStatus_name = map[int32]string{
		0: "Disconnected",
		1: "Active",
	}
	ConnectionStatus_value = map[string]int32{
		"Disconnected": 0,
		"Active":       1,
	}
)

func (x ConnectionStatus) Enum() *ConnectionStatus {
	p := new(ConnectionStatus)
	*p = x
	return p
}

func (x ConnectionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConnectionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_health_proto_enumTypes[0].Descriptor()
}

func (ConnectionStatus) Type() protoreflect.EnumType {
	return &file_health_proto_enumTypes[0]
}

func (x ConnectionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConnectionStatus.Descriptor instead.
func (ConnectionStatus) EnumDescriptor() ([]byte, []int) {
	return file_health_proto_rawDescGZIP(), []int{0}
}

type ServiceStatus int32

const (
	ServiceStatus_Stopped  ServiceStatus = 0 // service stopped
	ServiceStatus_Running  ServiceStatus = 1 // service running
	ServiceStatus_Degraded ServiceStatus = 2 // service heath is degraded
)

// Enum value maps for ServiceStatus.
var (
	ServiceStatus_name = map[int32]string{
		0: "Stopped",
		1: "Running",
		2: "Degraded",
	}
	ServiceStatus_value = map[string]int32{
		"Stopped":  0,
		"Running":  1,
		"Degraded": 2,
	}
)

func (x ServiceStatus) Enum() *ServiceStatus {
	p := new(ServiceStatus)
	*p = x
	return p
}

func (x ServiceStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServiceStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_health_proto_enumTypes[1].Descriptor()
}

func (ServiceStatus) Type() protoreflect.EnumType {
	return &file_health_proto_enumTypes[1]
}

func (x ServiceStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServiceStatus.Descriptor instead.
func (ServiceStatus) EnumDescriptor() ([]byte, []int) {
	return file_health_proto_rawDescGZIP(), []int{1}
}

type Health struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimestampUtc        string                `protobuf:"bytes,1,opt,name=timestamp_utc,json=timestampUtc,proto3" json:"timestamp_utc,omitempty"`                                    // current timestamp in utc
	ServiceName         string                `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`                                       // service name
	ServiceProvider     string                `protobuf:"bytes,3,opt,name=service_provider,json=serviceProvider,proto3" json:"service_provider,omitempty"`                           // service provider name
	ServiceVersion      string                `protobuf:"bytes,4,opt,name=service_version,json=serviceVersion,proto3" json:"service_version,omitempty"`                              // service version
	ServiceStatus       ServiceStatus         `protobuf:"varint,5,opt,name=service_status,json=serviceStatus,proto3,enum=v1.v1health.ServiceStatus" json:"service_status,omitempty"` // service status
	ServiceStartTimeUtc string                `protobuf:"bytes,6,opt,name=service_start_time_utc,json=serviceStartTimeUtc,proto3" json:"service_start_time_utc,omitempty"`           // service start time in utc
	Uptime              float64               `protobuf:"fixed64,7,opt,name=uptime,proto3" json:"uptime,omitempty"`                                                                  // service uptime in utc
	InboundConnections  []*InboundConnection  `protobuf:"bytes,8,rep,name=inbound_connections,json=inboundConnections,proto3" json:"inbound_connections,omitempty"`                  // inbound connections list
	OutboundConnections []*OutboundConnection `protobuf:"bytes,9,rep,name=outbound_connections,json=outboundConnections,proto3" json:"outbound_connections,omitempty"`               // outbound connections list
}

func (x *Health) Reset() {
	*x = Health{}
	if protoimpl.UnsafeEnabled {
		mi := &file_health_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Health) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Health) ProtoMessage() {}

func (x *Health) ProtoReflect() protoreflect.Message {
	mi := &file_health_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Health.ProtoReflect.Descriptor instead.
func (*Health) Descriptor() ([]byte, []int) {
	return file_health_proto_rawDescGZIP(), []int{0}
}

func (x *Health) GetTimestampUtc() string {
	if x != nil {
		return x.TimestampUtc
	}
	return ""
}

func (x *Health) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *Health) GetServiceProvider() string {
	if x != nil {
		return x.ServiceProvider
	}
	return ""
}

func (x *Health) GetServiceVersion() string {
	if x != nil {
		return x.ServiceVersion
	}
	return ""
}

func (x *Health) GetServiceStatus() ServiceStatus {
	if x != nil {
		return x.ServiceStatus
	}
	return ServiceStatus_Stopped
}

func (x *Health) GetServiceStartTimeUtc() string {
	if x != nil {
		return x.ServiceStartTimeUtc
	}
	return ""
}

func (x *Health) GetUptime() float64 {
	if x != nil {
		return x.Uptime
	}
	return 0
}

func (x *Health) GetInboundConnections() []*InboundConnection {
	if x != nil {
		return x.InboundConnections
	}
	return nil
}

func (x *Health) GetOutboundConnections() []*OutboundConnection {
	if x != nil {
		return x.OutboundConnections
	}
	return nil
}

type InboundConnection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApplicationName  string           `protobuf:"bytes,1,opt,name=application_name,json=applicationName,proto3" json:"application_name,omitempty"`                                       // name of the application
	ConnectionStatus ConnectionStatus `protobuf:"varint,2,opt,name=connection_status,json=connectionStatus,proto3,enum=v1.v1health.ConnectionStatus" json:"connection_status,omitempty"` // connectins status of the application
	TimestampUtc     string           `protobuf:"bytes,3,opt,name=timestamp_utc,json=timestampUtc,proto3" json:"timestamp_utc,omitempty"`                                                // current timestamp in utc
	Hostname         string           `protobuf:"bytes,4,opt,name=hostname,proto3" json:"hostname,omitempty"`                                                                            // hostname
	Address          string           `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`                                                                              // ip address of the application
	Os               string           `protobuf:"bytes,6,opt,name=os,proto3" json:"os,omitempty"`                                                                                        // OS
}

func (x *InboundConnection) Reset() {
	*x = InboundConnection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_health_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InboundConnection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InboundConnection) ProtoMessage() {}

func (x *InboundConnection) ProtoReflect() protoreflect.Message {
	mi := &file_health_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InboundConnection.ProtoReflect.Descriptor instead.
func (*InboundConnection) Descriptor() ([]byte, []int) {
	return file_health_proto_rawDescGZIP(), []int{1}
}

func (x *InboundConnection) GetApplicationName() string {
	if x != nil {
		return x.ApplicationName
	}
	return ""
}

func (x *InboundConnection) GetConnectionStatus() ConnectionStatus {
	if x != nil {
		return x.ConnectionStatus
	}
	return ConnectionStatus_Disconnected
}

func (x *InboundConnection) GetTimestampUtc() string {
	if x != nil {
		return x.TimestampUtc
	}
	return ""
}

func (x *InboundConnection) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *InboundConnection) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *InboundConnection) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

type OutboundConnection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApplicationName  string           `protobuf:"bytes,1,opt,name=application_name,json=applicationName,proto3" json:"application_name,omitempty"`                                       // applcition name
	TimestampUtc     string           `protobuf:"bytes,2,opt,name=timestamp_utc,json=timestampUtc,proto3" json:"timestamp_utc,omitempty"`                                                // current timestamp in utc
	Urls             []string         `protobuf:"bytes,3,rep,name=urls,proto3" json:"urls,omitempty"`                                                                                    // connection urls
	ConnectionStatus ConnectionStatus `protobuf:"varint,4,opt,name=connection_status,json=connectionStatus,proto3,enum=v1.v1health.ConnectionStatus" json:"connection_status,omitempty"` // connection status of application
}

func (x *OutboundConnection) Reset() {
	*x = OutboundConnection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_health_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutboundConnection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutboundConnection) ProtoMessage() {}

func (x *OutboundConnection) ProtoReflect() protoreflect.Message {
	mi := &file_health_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutboundConnection.ProtoReflect.Descriptor instead.
func (*OutboundConnection) Descriptor() ([]byte, []int) {
	return file_health_proto_rawDescGZIP(), []int{2}
}

func (x *OutboundConnection) GetApplicationName() string {
	if x != nil {
		return x.ApplicationName
	}
	return ""
}

func (x *OutboundConnection) GetTimestampUtc() string {
	if x != nil {
		return x.TimestampUtc
	}
	return ""
}

func (x *OutboundConnection) GetUrls() []string {
	if x != nil {
		return x.Urls
	}
	return nil
}

func (x *OutboundConnection) GetConnectionStatus() ConnectionStatus {
	if x != nil {
		return x.ConnectionStatus
	}
	return ConnectionStatus_Disconnected
}

var File_health_proto protoreflect.FileDescriptor

var file_health_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x76, 0x31, 0x2e, 0x76, 0x31, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x22, 0xd9, 0x03, 0x0a, 0x06,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x5f, 0x75, 0x74, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x55, 0x74, 0x63, 0x12, 0x21, 0x0a, 0x0c, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x29,
	0x0a, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x41, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x76, 0x31, 0x2e,
	0x76, 0x31, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x33, 0x0a, 0x16, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x74, 0x63, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x74, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x70,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x75, 0x70, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x4f, 0x0a, 0x13, 0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x76, 0x31, 0x2e, 0x76, 0x31, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x49, 0x6e,
	0x62, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x12, 0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x52, 0x0a, 0x14, 0x6f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x5f,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x76, 0x31, 0x2e, 0x76, 0x31, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e,
	0x4f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x13, 0x6f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xf5, 0x01, 0x0a, 0x11, 0x49, 0x6e, 0x62, 0x6f,
	0x75, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a,
	0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x4a, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x76, 0x31, 0x2e, 0x76, 0x31, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x5f, 0x75, 0x74, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x55, 0x74, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x0e, 0x0a, 0x02, 0x6f, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x73, 0x22,
	0xc4, 0x01, 0x0a, 0x12, 0x4f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x75,
	0x74, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x55, 0x74, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x12, 0x4a, 0x0a, 0x11, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x76, 0x31, 0x2e, 0x76, 0x31, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x30, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x10, 0x01, 0x2a, 0x37, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x74, 0x6f,
	0x70, 0x70, 0x65, 0x64, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e,
	0x67, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x65, 0x67, 0x72, 0x61, 0x64, 0x65, 0x64, 0x10,
	0x02, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6a, 0x6f, 0x62, 0x62, 0x6f, 0x78, 0x2d, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x72, 0x65, 0x63, 0x72,
	0x75, 0x69, 0x74, 0x65, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2f, 0x76, 0x31,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_health_proto_rawDescOnce sync.Once
	file_health_proto_rawDescData = file_health_proto_rawDesc
)

func file_health_proto_rawDescGZIP() []byte {
	file_health_proto_rawDescOnce.Do(func() {
		file_health_proto_rawDescData = protoimpl.X.CompressGZIP(file_health_proto_rawDescData)
	})
	return file_health_proto_rawDescData
}

var file_health_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_health_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_health_proto_goTypes = []interface{}{
	(ConnectionStatus)(0),      // 0: v1.v1health.ConnectionStatus
	(ServiceStatus)(0),         // 1: v1.v1health.ServiceStatus
	(*Health)(nil),             // 2: v1.v1health.Health
	(*InboundConnection)(nil),  // 3: v1.v1health.InboundConnection
	(*OutboundConnection)(nil), // 4: v1.v1health.OutboundConnection
}
var file_health_proto_depIdxs = []int32{
	1, // 0: v1.v1health.Health.service_status:type_name -> v1.v1health.ServiceStatus
	3, // 1: v1.v1health.Health.inbound_connections:type_name -> v1.v1health.InboundConnection
	4, // 2: v1.v1health.Health.outbound_connections:type_name -> v1.v1health.OutboundConnection
	0, // 3: v1.v1health.InboundConnection.connection_status:type_name -> v1.v1health.ConnectionStatus
	0, // 4: v1.v1health.OutboundConnection.connection_status:type_name -> v1.v1health.ConnectionStatus
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_health_proto_init() }
func file_health_proto_init() {
	if File_health_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_health_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Health); i {
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
		file_health_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InboundConnection); i {
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
		file_health_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutboundConnection); i {
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
			RawDescriptor: file_health_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_health_proto_goTypes,
		DependencyIndexes: file_health_proto_depIdxs,
		EnumInfos:         file_health_proto_enumTypes,
		MessageInfos:      file_health_proto_msgTypes,
	}.Build()
	File_health_proto = out.File
	file_health_proto_rawDesc = nil
	file_health_proto_goTypes = nil
	file_health_proto_depIdxs = nil
}
