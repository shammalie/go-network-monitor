// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: proto/network_capture/v1/pcap_event.proto

package network_capture_v1

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

type NetworkCaptureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetworkLayer     *NetworkLayer     `protobuf:"bytes,1,opt,name=network_layer,json=networkLayer,proto3" json:"network_layer,omitempty"`
	TransportLayer   *TransportLayer   `protobuf:"bytes,2,opt,name=transport_layer,json=transportLayer,proto3" json:"transport_layer,omitempty"`
	ApplicationLayer *ApplicationLayer `protobuf:"bytes,3,opt,name=application_layer,json=applicationLayer,proto3" json:"application_layer,omitempty"`
	Metadata         *Metadata         `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *NetworkCaptureRequest) Reset() {
	*x = NetworkCaptureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_network_capture_v1_pcap_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkCaptureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkCaptureRequest) ProtoMessage() {}

func (x *NetworkCaptureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_network_capture_v1_pcap_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkCaptureRequest.ProtoReflect.Descriptor instead.
func (*NetworkCaptureRequest) Descriptor() ([]byte, []int) {
	return file_proto_network_capture_v1_pcap_event_proto_rawDescGZIP(), []int{0}
}

func (x *NetworkCaptureRequest) GetNetworkLayer() *NetworkLayer {
	if x != nil {
		return x.NetworkLayer
	}
	return nil
}

func (x *NetworkCaptureRequest) GetTransportLayer() *TransportLayer {
	if x != nil {
		return x.TransportLayer
	}
	return nil
}

func (x *NetworkCaptureRequest) GetApplicationLayer() *ApplicationLayer {
	if x != nil {
		return x.ApplicationLayer
	}
	return nil
}

func (x *NetworkCaptureRequest) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type Protocol struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Protocol) Reset() {
	*x = Protocol{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_network_capture_v1_pcap_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Protocol) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Protocol) ProtoMessage() {}

func (x *Protocol) ProtoReflect() protoreflect.Message {
	mi := &file_proto_network_capture_v1_pcap_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Protocol.ProtoReflect.Descriptor instead.
func (*Protocol) Descriptor() ([]byte, []int) {
	return file_proto_network_capture_v1_pcap_event_proto_rawDescGZIP(), []int{1}
}

func (x *Protocol) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type NetworkLayer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SrcIp    string    `protobuf:"bytes,1,opt,name=src_ip,json=srcIp,proto3" json:"src_ip,omitempty"`
	DstIp    string    `protobuf:"bytes,2,opt,name=dst_ip,json=dstIp,proto3" json:"dst_ip,omitempty"`
	Protocol *Protocol `protobuf:"bytes,3,opt,name=protocol,proto3" json:"protocol,omitempty"`
}

func (x *NetworkLayer) Reset() {
	*x = NetworkLayer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_network_capture_v1_pcap_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkLayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkLayer) ProtoMessage() {}

func (x *NetworkLayer) ProtoReflect() protoreflect.Message {
	mi := &file_proto_network_capture_v1_pcap_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkLayer.ProtoReflect.Descriptor instead.
func (*NetworkLayer) Descriptor() ([]byte, []int) {
	return file_proto_network_capture_v1_pcap_event_proto_rawDescGZIP(), []int{2}
}

func (x *NetworkLayer) GetSrcIp() string {
	if x != nil {
		return x.SrcIp
	}
	return ""
}

func (x *NetworkLayer) GetDstIp() string {
	if x != nil {
		return x.DstIp
	}
	return ""
}

func (x *NetworkLayer) GetProtocol() *Protocol {
	if x != nil {
		return x.Protocol
	}
	return nil
}

type TransportLayer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SrcPort  string    `protobuf:"bytes,1,opt,name=src_port,json=srcPort,proto3" json:"src_port,omitempty"`
	DstPort  string    `protobuf:"bytes,2,opt,name=dst_port,json=dstPort,proto3" json:"dst_port,omitempty"`
	Protocol *Protocol `protobuf:"bytes,3,opt,name=protocol,proto3" json:"protocol,omitempty"`
}

func (x *TransportLayer) Reset() {
	*x = TransportLayer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_network_capture_v1_pcap_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransportLayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransportLayer) ProtoMessage() {}

func (x *TransportLayer) ProtoReflect() protoreflect.Message {
	mi := &file_proto_network_capture_v1_pcap_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransportLayer.ProtoReflect.Descriptor instead.
func (*TransportLayer) Descriptor() ([]byte, []int) {
	return file_proto_network_capture_v1_pcap_event_proto_rawDescGZIP(), []int{3}
}

func (x *TransportLayer) GetSrcPort() string {
	if x != nil {
		return x.SrcPort
	}
	return ""
}

func (x *TransportLayer) GetDstPort() string {
	if x != nil {
		return x.DstPort
	}
	return ""
}

func (x *TransportLayer) GetProtocol() *Protocol {
	if x != nil {
		return x.Protocol
	}
	return nil
}

type ApplicationLayer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Protocol *Protocol `protobuf:"bytes,1,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Payload  []byte    `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *ApplicationLayer) Reset() {
	*x = ApplicationLayer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_network_capture_v1_pcap_event_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationLayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationLayer) ProtoMessage() {}

func (x *ApplicationLayer) ProtoReflect() protoreflect.Message {
	mi := &file_proto_network_capture_v1_pcap_event_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationLayer.ProtoReflect.Descriptor instead.
func (*ApplicationLayer) Descriptor() ([]byte, []int) {
	return file_proto_network_capture_v1_pcap_event_proto_rawDescGZIP(), []int{4}
}

func (x *ApplicationLayer) GetProtocol() *Protocol {
	if x != nil {
		return x.Protocol
	}
	return nil
}

func (x *ApplicationLayer) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp            int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	CaptureLength        int64 `protobuf:"varint,2,opt,name=capture_length,json=captureLength,proto3" json:"capture_length,omitempty"`
	OriginalPacketLength int64 `protobuf:"varint,3,opt,name=original_packet_length,json=originalPacketLength,proto3" json:"original_packet_length,omitempty"`
	Truncated            bool  `protobuf:"varint,4,opt,name=truncated,proto3" json:"truncated,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_network_capture_v1_pcap_event_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_proto_network_capture_v1_pcap_event_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_proto_network_capture_v1_pcap_event_proto_rawDescGZIP(), []int{5}
}

func (x *Metadata) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Metadata) GetCaptureLength() int64 {
	if x != nil {
		return x.CaptureLength
	}
	return 0
}

func (x *Metadata) GetOriginalPacketLength() int64 {
	if x != nil {
		return x.OriginalPacketLength
	}
	return 0
}

func (x *Metadata) GetTruncated() bool {
	if x != nil {
		return x.Truncated
	}
	return false
}

type NetworkCaptureResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip        string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Action    string `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	Timestamp int64  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *NetworkCaptureResponse) Reset() {
	*x = NetworkCaptureResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_network_capture_v1_pcap_event_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkCaptureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkCaptureResponse) ProtoMessage() {}

func (x *NetworkCaptureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_network_capture_v1_pcap_event_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkCaptureResponse.ProtoReflect.Descriptor instead.
func (*NetworkCaptureResponse) Descriptor() ([]byte, []int) {
	return file_proto_network_capture_v1_pcap_event_proto_rawDescGZIP(), []int{6}
}

func (x *NetworkCaptureResponse) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *NetworkCaptureResponse) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *NetworkCaptureResponse) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_proto_network_capture_v1_pcap_event_proto protoreflect.FileDescriptor

var file_proto_network_capture_v1_pcap_event_proto_rawDesc = []byte{
	0x0a, 0x29, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f,
	0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x63, 0x61, 0x70, 0x5f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x22,
	0xb8, 0x02, 0x0a, 0x15, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x61, 0x70, 0x74, 0x75,
	0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x0d, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x63, 0x61, 0x70, 0x74, 0x75,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4c, 0x61, 0x79,
	0x65, 0x72, 0x52, 0x0c, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4c, 0x61, 0x79, 0x65, 0x72,
	0x12, 0x4b, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x5f, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x0e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x51, 0x0a,
	0x11, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x5f, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x10,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x61, 0x79, 0x65, 0x72,
	0x12, 0x38, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x63, 0x61, 0x70,
	0x74, 0x75, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x1e, 0x0a, 0x08, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x76, 0x0a, 0x0c, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x72,
	0x63, 0x5f, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x72, 0x63, 0x49,
	0x70, 0x12, 0x15, 0x0a, 0x06, 0x64, 0x73, 0x74, 0x5f, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x64, 0x73, 0x74, 0x49, 0x70, 0x12, 0x38, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x22, 0x80, 0x01, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x4c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x72, 0x63, 0x5f, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x72, 0x63, 0x50, 0x6f, 0x72, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x64, 0x73, 0x74, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x64, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x38, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x08, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x66, 0x0a, 0x10, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0xa3, 0x01,
	0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x61, 0x70, 0x74,
	0x75, 0x72, 0x65, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12,
	0x34, 0x0a, 0x16, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x63, 0x6b,
	0x65, 0x74, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x14, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x4c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x72, 0x75, 0x6e, 0x63, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x74, 0x72, 0x75, 0x6e, 0x63, 0x61,
	0x74, 0x65, 0x64, 0x22, 0x5e, 0x0a, 0x16, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x61,
	0x70, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x32, 0x86, 0x01, 0x0a, 0x15, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43,
	0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6d, 0x0a,
	0x0e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x29, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x61, 0x70, 0x74,
	0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x14, 0x5a, 0x12,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_network_capture_v1_pcap_event_proto_rawDescOnce sync.Once
	file_proto_network_capture_v1_pcap_event_proto_rawDescData = file_proto_network_capture_v1_pcap_event_proto_rawDesc
)

func file_proto_network_capture_v1_pcap_event_proto_rawDescGZIP() []byte {
	file_proto_network_capture_v1_pcap_event_proto_rawDescOnce.Do(func() {
		file_proto_network_capture_v1_pcap_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_network_capture_v1_pcap_event_proto_rawDescData)
	})
	return file_proto_network_capture_v1_pcap_event_proto_rawDescData
}

var file_proto_network_capture_v1_pcap_event_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_network_capture_v1_pcap_event_proto_goTypes = []interface{}{
	(*NetworkCaptureRequest)(nil),  // 0: network_capture.v1.NetworkCaptureRequest
	(*Protocol)(nil),               // 1: network_capture.v1.Protocol
	(*NetworkLayer)(nil),           // 2: network_capture.v1.NetworkLayer
	(*TransportLayer)(nil),         // 3: network_capture.v1.TransportLayer
	(*ApplicationLayer)(nil),       // 4: network_capture.v1.ApplicationLayer
	(*Metadata)(nil),               // 5: network_capture.v1.Metadata
	(*NetworkCaptureResponse)(nil), // 6: network_capture.v1.NetworkCaptureResponse
}
var file_proto_network_capture_v1_pcap_event_proto_depIdxs = []int32{
	2, // 0: network_capture.v1.NetworkCaptureRequest.network_layer:type_name -> network_capture.v1.NetworkLayer
	3, // 1: network_capture.v1.NetworkCaptureRequest.transport_layer:type_name -> network_capture.v1.TransportLayer
	4, // 2: network_capture.v1.NetworkCaptureRequest.application_layer:type_name -> network_capture.v1.ApplicationLayer
	5, // 3: network_capture.v1.NetworkCaptureRequest.metadata:type_name -> network_capture.v1.Metadata
	1, // 4: network_capture.v1.NetworkLayer.protocol:type_name -> network_capture.v1.Protocol
	1, // 5: network_capture.v1.TransportLayer.protocol:type_name -> network_capture.v1.Protocol
	1, // 6: network_capture.v1.ApplicationLayer.protocol:type_name -> network_capture.v1.Protocol
	0, // 7: network_capture.v1.NetworkCaptureService.NetworkCapture:input_type -> network_capture.v1.NetworkCaptureRequest
	6, // 8: network_capture.v1.NetworkCaptureService.NetworkCapture:output_type -> network_capture.v1.NetworkCaptureResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_proto_network_capture_v1_pcap_event_proto_init() }
func file_proto_network_capture_v1_pcap_event_proto_init() {
	if File_proto_network_capture_v1_pcap_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_network_capture_v1_pcap_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkCaptureRequest); i {
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
		file_proto_network_capture_v1_pcap_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Protocol); i {
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
		file_proto_network_capture_v1_pcap_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkLayer); i {
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
		file_proto_network_capture_v1_pcap_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransportLayer); i {
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
		file_proto_network_capture_v1_pcap_event_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationLayer); i {
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
		file_proto_network_capture_v1_pcap_event_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
		file_proto_network_capture_v1_pcap_event_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkCaptureResponse); i {
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
			RawDescriptor: file_proto_network_capture_v1_pcap_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_network_capture_v1_pcap_event_proto_goTypes,
		DependencyIndexes: file_proto_network_capture_v1_pcap_event_proto_depIdxs,
		MessageInfos:      file_proto_network_capture_v1_pcap_event_proto_msgTypes,
	}.Build()
	File_proto_network_capture_v1_pcap_event_proto = out.File
	file_proto_network_capture_v1_pcap_event_proto_rawDesc = nil
	file_proto_network_capture_v1_pcap_event_proto_goTypes = nil
	file_proto_network_capture_v1_pcap_event_proto_depIdxs = nil
}