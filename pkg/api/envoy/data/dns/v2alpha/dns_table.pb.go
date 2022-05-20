// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: envoy/data/dns/v2alpha/dns_table.proto

package v2alpha

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	matcher "github.com/datawire/ambassador/v2/pkg/api/envoy/type/matcher"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	duration "github.com/golang/protobuf/ptypes/duration"
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

// This message contains the configuration for the DNS Filter if populated
// from the control plane
type DnsTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Control how many times envoy makes an attempt to forward a query to
	// an external server
	ExternalRetryCount uint32 `protobuf:"varint,1,opt,name=external_retry_count,json=externalRetryCount,proto3" json:"external_retry_count,omitempty"`
	// Fully qualified domain names for which Envoy will respond to queries
	VirtualDomains []*DnsTable_DnsVirtualDomain `protobuf:"bytes,2,rep,name=virtual_domains,json=virtualDomains,proto3" json:"virtual_domains,omitempty"`
	// This field serves to help Envoy determine whether it can authoritatively
	// answer a query for a name matching a suffix in this list. If the query
	// name does not match a suffix in this list, Envoy will forward
	// the query to an upstream DNS server
	KnownSuffixes []*matcher.StringMatcher `protobuf:"bytes,3,rep,name=known_suffixes,json=knownSuffixes,proto3" json:"known_suffixes,omitempty"`
}

func (x *DnsTable) Reset() {
	*x = DnsTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_data_dns_v2alpha_dns_table_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsTable) ProtoMessage() {}

func (x *DnsTable) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_data_dns_v2alpha_dns_table_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsTable.ProtoReflect.Descriptor instead.
func (*DnsTable) Descriptor() ([]byte, []int) {
	return file_envoy_data_dns_v2alpha_dns_table_proto_rawDescGZIP(), []int{0}
}

func (x *DnsTable) GetExternalRetryCount() uint32 {
	if x != nil {
		return x.ExternalRetryCount
	}
	return 0
}

func (x *DnsTable) GetVirtualDomains() []*DnsTable_DnsVirtualDomain {
	if x != nil {
		return x.VirtualDomains
	}
	return nil
}

func (x *DnsTable) GetKnownSuffixes() []*matcher.StringMatcher {
	if x != nil {
		return x.KnownSuffixes
	}
	return nil
}

// This message contains a list of IP addresses returned for a query for a known name
type DnsTable_AddressList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This field contains a well formed IP address that is returned
	// in the answer for a name query. The address field can be an
	// IPv4 or IPv6 address. Address family detection is done automatically
	// when Envoy parses the string. Since this field is repeated,
	// Envoy will return one randomly chosen entry from this list in the
	// DNS response. The random index will vary per query so that we prevent
	// clients pinning on a single address for a configured domain
	Address []string `protobuf:"bytes,1,rep,name=address,proto3" json:"address,omitempty"`
}

func (x *DnsTable_AddressList) Reset() {
	*x = DnsTable_AddressList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_data_dns_v2alpha_dns_table_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsTable_AddressList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsTable_AddressList) ProtoMessage() {}

func (x *DnsTable_AddressList) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_data_dns_v2alpha_dns_table_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsTable_AddressList.ProtoReflect.Descriptor instead.
func (*DnsTable_AddressList) Descriptor() ([]byte, []int) {
	return file_envoy_data_dns_v2alpha_dns_table_proto_rawDescGZIP(), []int{0, 0}
}

func (x *DnsTable_AddressList) GetAddress() []string {
	if x != nil {
		return x.Address
	}
	return nil
}

// This message type is extensible and can contain a list of addresses
// or dictate some other method for resolving the addresses for an
// endpoint
type DnsTable_DnsEndpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to EndpointConfig:
	//	*DnsTable_DnsEndpoint_AddressList
	EndpointConfig isDnsTable_DnsEndpoint_EndpointConfig `protobuf_oneof:"endpoint_config"`
}

func (x *DnsTable_DnsEndpoint) Reset() {
	*x = DnsTable_DnsEndpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_data_dns_v2alpha_dns_table_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsTable_DnsEndpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsTable_DnsEndpoint) ProtoMessage() {}

func (x *DnsTable_DnsEndpoint) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_data_dns_v2alpha_dns_table_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsTable_DnsEndpoint.ProtoReflect.Descriptor instead.
func (*DnsTable_DnsEndpoint) Descriptor() ([]byte, []int) {
	return file_envoy_data_dns_v2alpha_dns_table_proto_rawDescGZIP(), []int{0, 1}
}

func (m *DnsTable_DnsEndpoint) GetEndpointConfig() isDnsTable_DnsEndpoint_EndpointConfig {
	if m != nil {
		return m.EndpointConfig
	}
	return nil
}

func (x *DnsTable_DnsEndpoint) GetAddressList() *DnsTable_AddressList {
	if x, ok := x.GetEndpointConfig().(*DnsTable_DnsEndpoint_AddressList); ok {
		return x.AddressList
	}
	return nil
}

type isDnsTable_DnsEndpoint_EndpointConfig interface {
	isDnsTable_DnsEndpoint_EndpointConfig()
}

type DnsTable_DnsEndpoint_AddressList struct {
	AddressList *DnsTable_AddressList `protobuf:"bytes,1,opt,name=address_list,json=addressList,proto3,oneof"`
}

func (*DnsTable_DnsEndpoint_AddressList) isDnsTable_DnsEndpoint_EndpointConfig() {}

type DnsTable_DnsVirtualDomain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The domain name for which Envoy will respond to query requests
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The configuration containing the method to determine the address
	// of this endpoint
	Endpoint *DnsTable_DnsEndpoint `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	// Sets the TTL in dns answers from Envoy returned to the client
	AnswerTtl *duration.Duration `protobuf:"bytes,3,opt,name=answer_ttl,json=answerTtl,proto3" json:"answer_ttl,omitempty"`
}

func (x *DnsTable_DnsVirtualDomain) Reset() {
	*x = DnsTable_DnsVirtualDomain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_data_dns_v2alpha_dns_table_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsTable_DnsVirtualDomain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsTable_DnsVirtualDomain) ProtoMessage() {}

func (x *DnsTable_DnsVirtualDomain) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_data_dns_v2alpha_dns_table_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsTable_DnsVirtualDomain.ProtoReflect.Descriptor instead.
func (*DnsTable_DnsVirtualDomain) Descriptor() ([]byte, []int) {
	return file_envoy_data_dns_v2alpha_dns_table_proto_rawDescGZIP(), []int{0, 2}
}

func (x *DnsTable_DnsVirtualDomain) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DnsTable_DnsVirtualDomain) GetEndpoint() *DnsTable_DnsEndpoint {
	if x != nil {
		return x.Endpoint
	}
	return nil
}

func (x *DnsTable_DnsVirtualDomain) GetAnswerTtl() *duration.Duration {
	if x != nil {
		return x.AnswerTtl
	}
	return nil
}

var File_envoy_data_dns_v2alpha_dns_table_proto protoreflect.FileDescriptor

var file_envoy_data_dns_v2alpha_dns_table_proto_rawDesc = []byte{
	0x0a, 0x26, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x64, 0x6e, 0x73,
	0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x64, 0x6e, 0x73, 0x5f, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe2, 0x04, 0x0a, 0x08, 0x44, 0x6e,
	0x73, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65,
	0x74, 0x72, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x64, 0x0a, 0x0f, 0x76, 0x69, 0x72, 0x74,
	0x75, 0x61, 0x6c, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x31, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x64,
	0x6e, 0x73, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x44, 0x6e, 0x73, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x2e, 0x44, 0x6e, 0x73, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x0e,
	0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x12, 0x48,
	0x0a, 0x0e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x0d, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x53, 0x75, 0x66, 0x66, 0x69, 0x78, 0x65, 0x73, 0x1a, 0x37, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x0e, 0xfa, 0x42, 0x0b, 0x92, 0x01, 0x08,
	0x08, 0x01, 0x22, 0x04, 0x72, 0x02, 0x10, 0x03, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x1a, 0x78, 0x0a, 0x0b, 0x44, 0x6e, 0x73, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x12, 0x51, 0x0a, 0x0c, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x44, 0x6e, 0x73, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x16, 0x0a, 0x0f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x1a, 0xc0, 0x01, 0x0a, 0x10,
	0x44, 0x6e, 0x73, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x12, 0x1e, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a,
	0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x02, 0xc0, 0x01, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x48, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x64, 0x6e, 0x73, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x44, 0x6e, 0x73, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x2e, 0x44, 0x6e, 0x73, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x42, 0x0a, 0x0a, 0x61, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x5f, 0x74, 0x74, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05, 0xaa, 0x01,
	0x02, 0x2a, 0x00, 0x52, 0x09, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x54, 0x74, 0x6c, 0x42, 0x86,
	0x01, 0x0a, 0x24, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x64, 0x6e, 0x73, 0x2e,
	0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x0d, 0x44, 0x6e, 0x73, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f,
	0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65,
	0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x64, 0x6e, 0x73, 0x2f,
	0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x08, 0x01, 0xba,
	0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_data_dns_v2alpha_dns_table_proto_rawDescOnce sync.Once
	file_envoy_data_dns_v2alpha_dns_table_proto_rawDescData = file_envoy_data_dns_v2alpha_dns_table_proto_rawDesc
)

func file_envoy_data_dns_v2alpha_dns_table_proto_rawDescGZIP() []byte {
	file_envoy_data_dns_v2alpha_dns_table_proto_rawDescOnce.Do(func() {
		file_envoy_data_dns_v2alpha_dns_table_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_data_dns_v2alpha_dns_table_proto_rawDescData)
	})
	return file_envoy_data_dns_v2alpha_dns_table_proto_rawDescData
}

var file_envoy_data_dns_v2alpha_dns_table_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_envoy_data_dns_v2alpha_dns_table_proto_goTypes = []interface{}{
	(*DnsTable)(nil),                  // 0: envoy.data.dns.v2alpha.DnsTable
	(*DnsTable_AddressList)(nil),      // 1: envoy.data.dns.v2alpha.DnsTable.AddressList
	(*DnsTable_DnsEndpoint)(nil),      // 2: envoy.data.dns.v2alpha.DnsTable.DnsEndpoint
	(*DnsTable_DnsVirtualDomain)(nil), // 3: envoy.data.dns.v2alpha.DnsTable.DnsVirtualDomain
	(*matcher.StringMatcher)(nil),     // 4: envoy.type.matcher.StringMatcher
	(*duration.Duration)(nil),         // 5: google.protobuf.Duration
}
var file_envoy_data_dns_v2alpha_dns_table_proto_depIdxs = []int32{
	3, // 0: envoy.data.dns.v2alpha.DnsTable.virtual_domains:type_name -> envoy.data.dns.v2alpha.DnsTable.DnsVirtualDomain
	4, // 1: envoy.data.dns.v2alpha.DnsTable.known_suffixes:type_name -> envoy.type.matcher.StringMatcher
	1, // 2: envoy.data.dns.v2alpha.DnsTable.DnsEndpoint.address_list:type_name -> envoy.data.dns.v2alpha.DnsTable.AddressList
	2, // 3: envoy.data.dns.v2alpha.DnsTable.DnsVirtualDomain.endpoint:type_name -> envoy.data.dns.v2alpha.DnsTable.DnsEndpoint
	5, // 4: envoy.data.dns.v2alpha.DnsTable.DnsVirtualDomain.answer_ttl:type_name -> google.protobuf.Duration
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_envoy_data_dns_v2alpha_dns_table_proto_init() }
func file_envoy_data_dns_v2alpha_dns_table_proto_init() {
	if File_envoy_data_dns_v2alpha_dns_table_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_data_dns_v2alpha_dns_table_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsTable); i {
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
		file_envoy_data_dns_v2alpha_dns_table_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsTable_AddressList); i {
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
		file_envoy_data_dns_v2alpha_dns_table_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsTable_DnsEndpoint); i {
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
		file_envoy_data_dns_v2alpha_dns_table_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsTable_DnsVirtualDomain); i {
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
	file_envoy_data_dns_v2alpha_dns_table_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*DnsTable_DnsEndpoint_AddressList)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_data_dns_v2alpha_dns_table_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_data_dns_v2alpha_dns_table_proto_goTypes,
		DependencyIndexes: file_envoy_data_dns_v2alpha_dns_table_proto_depIdxs,
		MessageInfos:      file_envoy_data_dns_v2alpha_dns_table_proto_msgTypes,
	}.Build()
	File_envoy_data_dns_v2alpha_dns_table_proto = out.File
	file_envoy_data_dns_v2alpha_dns_table_proto_rawDesc = nil
	file_envoy_data_dns_v2alpha_dns_table_proto_goTypes = nil
	file_envoy_data_dns_v2alpha_dns_table_proto_depIdxs = nil
}
