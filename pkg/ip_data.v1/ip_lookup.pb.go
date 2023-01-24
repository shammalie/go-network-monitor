// https://dev.bitolog.com/grpc-long-lived-streaming/

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: proto/ip_lookup/v1/ip_lookup.proto

package ip_data_v1

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

type BearerStatusRequest_BearerStatus int32

const (
	BearerStatusRequest_AVAILABLE    BearerStatusRequest_BearerStatus = 0
	BearerStatusRequest_UNAVAILABLE  BearerStatusRequest_BearerStatus = 1
	BearerStatusRequest_RATE_LIMITED BearerStatusRequest_BearerStatus = 2
)

// Enum value maps for BearerStatusRequest_BearerStatus.
var (
	BearerStatusRequest_BearerStatus_name = map[int32]string{
		0: "AVAILABLE",
		1: "UNAVAILABLE",
		2: "RATE_LIMITED",
	}
	BearerStatusRequest_BearerStatus_value = map[string]int32{
		"AVAILABLE":    0,
		"UNAVAILABLE":  1,
		"RATE_LIMITED": 2,
	}
)

func (x BearerStatusRequest_BearerStatus) Enum() *BearerStatusRequest_BearerStatus {
	p := new(BearerStatusRequest_BearerStatus)
	*p = x
	return p
}

func (x BearerStatusRequest_BearerStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BearerStatusRequest_BearerStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_ip_lookup_v1_ip_lookup_proto_enumTypes[0].Descriptor()
}

func (BearerStatusRequest_BearerStatus) Type() protoreflect.EnumType {
	return &file_proto_ip_lookup_v1_ip_lookup_proto_enumTypes[0]
}

func (x BearerStatusRequest_BearerStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BearerStatusRequest_BearerStatus.Descriptor instead.
func (BearerStatusRequest_BearerStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_ip_lookup_v1_ip_lookup_proto_rawDescGZIP(), []int{2, 0}
}

type BearerIdentity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Ip       string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	Hostname string `protobuf:"bytes,3,opt,name=hostname,proto3" json:"hostname,omitempty"`
	ApiKey   string `protobuf:"bytes,4,opt,name=apiKey,proto3" json:"apiKey,omitempty"`
}

func (x *BearerIdentity) Reset() {
	*x = BearerIdentity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ip_lookup_v1_ip_lookup_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BearerIdentity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BearerIdentity) ProtoMessage() {}

func (x *BearerIdentity) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ip_lookup_v1_ip_lookup_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BearerIdentity.ProtoReflect.Descriptor instead.
func (*BearerIdentity) Descriptor() ([]byte, []int) {
	return file_proto_ip_lookup_v1_ip_lookup_proto_rawDescGZIP(), []int{0}
}

func (x *BearerIdentity) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BearerIdentity) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *BearerIdentity) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *BearerIdentity) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

type BearerStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BearerStatusResponse) Reset() {
	*x = BearerStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ip_lookup_v1_ip_lookup_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BearerStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BearerStatusResponse) ProtoMessage() {}

func (x *BearerStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ip_lookup_v1_ip_lookup_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BearerStatusResponse.ProtoReflect.Descriptor instead.
func (*BearerStatusResponse) Descriptor() ([]byte, []int) {
	return file_proto_ip_lookup_v1_ip_lookup_proto_rawDescGZIP(), []int{1}
}

type BearerStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BearerIdentity *BearerIdentity                  `protobuf:"bytes,1,opt,name=bearerIdentity,proto3" json:"bearerIdentity,omitempty"`
	BearerStatus   BearerStatusRequest_BearerStatus `protobuf:"varint,2,opt,name=bearerStatus,proto3,enum=ip_data.v1.BearerStatusRequest_BearerStatus" json:"bearerStatus,omitempty"`
}

func (x *BearerStatusRequest) Reset() {
	*x = BearerStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ip_lookup_v1_ip_lookup_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BearerStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BearerStatusRequest) ProtoMessage() {}

func (x *BearerStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ip_lookup_v1_ip_lookup_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BearerStatusRequest.ProtoReflect.Descriptor instead.
func (*BearerStatusRequest) Descriptor() ([]byte, []int) {
	return file_proto_ip_lookup_v1_ip_lookup_proto_rawDescGZIP(), []int{2}
}

func (x *BearerStatusRequest) GetBearerIdentity() *BearerIdentity {
	if x != nil {
		return x.BearerIdentity
	}
	return nil
}

func (x *BearerStatusRequest) GetBearerStatus() BearerStatusRequest_BearerStatus {
	if x != nil {
		return x.BearerStatus
	}
	return BearerStatusRequest_AVAILABLE
}

type LookupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip       string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	BearerId string `protobuf:"bytes,2,opt,name=bearerId,proto3" json:"bearerId,omitempty"`
}

func (x *LookupRequest) Reset() {
	*x = LookupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ip_lookup_v1_ip_lookup_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LookupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupRequest) ProtoMessage() {}

func (x *LookupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ip_lookup_v1_ip_lookup_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupRequest.ProtoReflect.Descriptor instead.
func (*LookupRequest) Descriptor() ([]byte, []int) {
	return file_proto_ip_lookup_v1_ip_lookup_proto_rawDescGZIP(), []int{3}
}

func (x *LookupRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *LookupRequest) GetBearerId() string {
	if x != nil {
		return x.BearerId
	}
	return ""
}

type LookupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip                 string  `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	City               string  `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	Region             string  `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
	RegionCode         string  `protobuf:"bytes,4,opt,name=regionCode,proto3" json:"regionCode,omitempty"`
	CountryCode        string  `protobuf:"bytes,5,opt,name=countryCode,proto3" json:"countryCode,omitempty"`
	CountryCodeIso3    string  `protobuf:"bytes,6,opt,name=countryCodeIso3,proto3" json:"countryCodeIso3,omitempty"`
	CountryName        string  `protobuf:"bytes,7,opt,name=countryName,proto3" json:"countryName,omitempty"`
	CountryCapital     string  `protobuf:"bytes,8,opt,name=countryCapital,proto3" json:"countryCapital,omitempty"`
	CountryTld         string  `protobuf:"bytes,9,opt,name=countryTld,proto3" json:"countryTld,omitempty"`
	ContinentCode      string  `protobuf:"bytes,10,opt,name=continentCode,proto3" json:"continentCode,omitempty"`
	InEu               bool    `protobuf:"varint,11,opt,name=inEu,proto3" json:"inEu,omitempty"`
	Postal             string  `protobuf:"bytes,12,opt,name=postal,proto3" json:"postal,omitempty"`
	Latitude           float32 `protobuf:"fixed32,13,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude          float32 `protobuf:"fixed32,14,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Timezone           string  `protobuf:"bytes,15,opt,name=timezone,proto3" json:"timezone,omitempty"`
	UtcOffset          string  `protobuf:"bytes,16,opt,name=utcOffset,proto3" json:"utcOffset,omitempty"`
	CountryCallingCode string  `protobuf:"bytes,17,opt,name=countryCallingCode,proto3" json:"countryCallingCode,omitempty"`
	Currency           string  `protobuf:"bytes,18,opt,name=currency,proto3" json:"currency,omitempty"`
	CurrencyName       string  `protobuf:"bytes,19,opt,name=currencyName,proto3" json:"currencyName,omitempty"`
	Languages          string  `protobuf:"bytes,20,opt,name=languages,proto3" json:"languages,omitempty"`
	Asn                string  `protobuf:"bytes,21,opt,name=asn,proto3" json:"asn,omitempty"`
	Org                string  `protobuf:"bytes,22,opt,name=org,proto3" json:"org,omitempty"`
	Error              string  `protobuf:"bytes,23,opt,name=error,proto3" json:"error,omitempty"`
	Reason             string  `protobuf:"bytes,24,opt,name=reason,proto3" json:"reason,omitempty"`
	Message            string  `protobuf:"bytes,25,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *LookupResponse) Reset() {
	*x = LookupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ip_lookup_v1_ip_lookup_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LookupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupResponse) ProtoMessage() {}

func (x *LookupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ip_lookup_v1_ip_lookup_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupResponse.ProtoReflect.Descriptor instead.
func (*LookupResponse) Descriptor() ([]byte, []int) {
	return file_proto_ip_lookup_v1_ip_lookup_proto_rawDescGZIP(), []int{4}
}

func (x *LookupResponse) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *LookupResponse) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *LookupResponse) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *LookupResponse) GetRegionCode() string {
	if x != nil {
		return x.RegionCode
	}
	return ""
}

func (x *LookupResponse) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *LookupResponse) GetCountryCodeIso3() string {
	if x != nil {
		return x.CountryCodeIso3
	}
	return ""
}

func (x *LookupResponse) GetCountryName() string {
	if x != nil {
		return x.CountryName
	}
	return ""
}

func (x *LookupResponse) GetCountryCapital() string {
	if x != nil {
		return x.CountryCapital
	}
	return ""
}

func (x *LookupResponse) GetCountryTld() string {
	if x != nil {
		return x.CountryTld
	}
	return ""
}

func (x *LookupResponse) GetContinentCode() string {
	if x != nil {
		return x.ContinentCode
	}
	return ""
}

func (x *LookupResponse) GetInEu() bool {
	if x != nil {
		return x.InEu
	}
	return false
}

func (x *LookupResponse) GetPostal() string {
	if x != nil {
		return x.Postal
	}
	return ""
}

func (x *LookupResponse) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *LookupResponse) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *LookupResponse) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *LookupResponse) GetUtcOffset() string {
	if x != nil {
		return x.UtcOffset
	}
	return ""
}

func (x *LookupResponse) GetCountryCallingCode() string {
	if x != nil {
		return x.CountryCallingCode
	}
	return ""
}

func (x *LookupResponse) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *LookupResponse) GetCurrencyName() string {
	if x != nil {
		return x.CurrencyName
	}
	return ""
}

func (x *LookupResponse) GetLanguages() string {
	if x != nil {
		return x.Languages
	}
	return ""
}

func (x *LookupResponse) GetAsn() string {
	if x != nil {
		return x.Asn
	}
	return ""
}

func (x *LookupResponse) GetOrg() string {
	if x != nil {
		return x.Org
	}
	return ""
}

func (x *LookupResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *LookupResponse) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *LookupResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_ip_lookup_v1_ip_lookup_proto protoreflect.FileDescriptor

var file_proto_ip_lookup_v1_ip_lookup_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x70, 0x5f, 0x6c, 0x6f, 0x6f, 0x6b, 0x75,
	0x70, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x70, 0x5f, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x69, 0x70, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31,
	0x22, 0x64, 0x0a, 0x0e, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x22, 0x16, 0x0a, 0x14, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xed,
	0x01, 0x0a, 0x13, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x0e, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x69, 0x70, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x65, 0x61, 0x72,
	0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x0e, 0x62, 0x65, 0x61, 0x72,
	0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x50, 0x0a, 0x0c, 0x62, 0x65,
	0x61, 0x72, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x2c, 0x2e, 0x69, 0x70, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x65,
	0x61, 0x72, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0c,
	0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x40, 0x0a, 0x0c,
	0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0d, 0x0a, 0x09,
	0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x55,
	0x4e, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c,
	0x52, 0x41, 0x54, 0x45, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x45, 0x44, 0x10, 0x02, 0x22, 0x3b,
	0x0a, 0x0d, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12,
	0x1a, 0x0a, 0x08, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x49, 0x64, 0x22, 0xe2, 0x05, 0x0a, 0x0e,
	0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69,
	0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x28, 0x0a, 0x0f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x49, 0x73, 0x6f, 0x33, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f,
	0x64, 0x65, 0x49, 0x73, 0x6f, 0x33, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x43, 0x61, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x61, 0x70, 0x69, 0x74, 0x61, 0x6c,
	0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x54, 0x6c, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x54, 0x6c, 0x64,
	0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x45, 0x75, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x6e, 0x45, 0x75, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f,
	0x73, 0x74, 0x61, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74,
	0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x74, 0x63, 0x4f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x74, 0x63,
	0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x2e, 0x0a, 0x12, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x43, 0x61, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x12, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x61, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x73, 0x6e, 0x18, 0x15, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x61, 0x73, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x72, 0x67, 0x18, 0x16, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x72, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x32, 0xb0, 0x01, 0x0a, 0x0d, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x5c, 0x0a, 0x15, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x42,
	0x65, 0x61, 0x72, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x2e, 0x69, 0x70,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x69,
	0x70, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x41, 0x0a, 0x06, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x12, 0x19, 0x2e, 0x69, 0x70, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x69, 0x70, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x69, 0x70, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_ip_lookup_v1_ip_lookup_proto_rawDescOnce sync.Once
	file_proto_ip_lookup_v1_ip_lookup_proto_rawDescData = file_proto_ip_lookup_v1_ip_lookup_proto_rawDesc
)

func file_proto_ip_lookup_v1_ip_lookup_proto_rawDescGZIP() []byte {
	file_proto_ip_lookup_v1_ip_lookup_proto_rawDescOnce.Do(func() {
		file_proto_ip_lookup_v1_ip_lookup_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ip_lookup_v1_ip_lookup_proto_rawDescData)
	})
	return file_proto_ip_lookup_v1_ip_lookup_proto_rawDescData
}

var file_proto_ip_lookup_v1_ip_lookup_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_ip_lookup_v1_ip_lookup_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_ip_lookup_v1_ip_lookup_proto_goTypes = []interface{}{
	(BearerStatusRequest_BearerStatus)(0), // 0: ip_data.v1.BearerStatusRequest.BearerStatus
	(*BearerIdentity)(nil),                // 1: ip_data.v1.BearerIdentity
	(*BearerStatusResponse)(nil),          // 2: ip_data.v1.BearerStatusResponse
	(*BearerStatusRequest)(nil),           // 3: ip_data.v1.BearerStatusRequest
	(*LookupRequest)(nil),                 // 4: ip_data.v1.LookupRequest
	(*LookupResponse)(nil),                // 5: ip_data.v1.LookupResponse
}
var file_proto_ip_lookup_v1_ip_lookup_proto_depIdxs = []int32{
	1, // 0: ip_data.v1.BearerStatusRequest.bearerIdentity:type_name -> ip_data.v1.BearerIdentity
	0, // 1: ip_data.v1.BearerStatusRequest.bearerStatus:type_name -> ip_data.v1.BearerStatusRequest.BearerStatus
	3, // 2: ip_data.v1.LookupService.BroadcastBearerStatus:input_type -> ip_data.v1.BearerStatusRequest
	4, // 3: ip_data.v1.LookupService.Lookup:input_type -> ip_data.v1.LookupRequest
	2, // 4: ip_data.v1.LookupService.BroadcastBearerStatus:output_type -> ip_data.v1.BearerStatusResponse
	5, // 5: ip_data.v1.LookupService.Lookup:output_type -> ip_data.v1.LookupResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_ip_lookup_v1_ip_lookup_proto_init() }
func file_proto_ip_lookup_v1_ip_lookup_proto_init() {
	if File_proto_ip_lookup_v1_ip_lookup_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_ip_lookup_v1_ip_lookup_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BearerIdentity); i {
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
		file_proto_ip_lookup_v1_ip_lookup_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BearerStatusResponse); i {
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
		file_proto_ip_lookup_v1_ip_lookup_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BearerStatusRequest); i {
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
		file_proto_ip_lookup_v1_ip_lookup_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LookupRequest); i {
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
		file_proto_ip_lookup_v1_ip_lookup_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LookupResponse); i {
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
			RawDescriptor: file_proto_ip_lookup_v1_ip_lookup_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_ip_lookup_v1_ip_lookup_proto_goTypes,
		DependencyIndexes: file_proto_ip_lookup_v1_ip_lookup_proto_depIdxs,
		EnumInfos:         file_proto_ip_lookup_v1_ip_lookup_proto_enumTypes,
		MessageInfos:      file_proto_ip_lookup_v1_ip_lookup_proto_msgTypes,
	}.Build()
	File_proto_ip_lookup_v1_ip_lookup_proto = out.File
	file_proto_ip_lookup_v1_ip_lookup_proto_rawDesc = nil
	file_proto_ip_lookup_v1_ip_lookup_proto_goTypes = nil
	file_proto_ip_lookup_v1_ip_lookup_proto_depIdxs = nil
}
