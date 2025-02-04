//
//Copyright The Kubernetes Authors.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// This file was autogenerated by go-to-protobuf. Do not edit it manually!

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: k8s.io/api/discovery/v1alpha1/generated.proto

package v1alpha1

import (
	"github.com/ericchiang/k8s/apis/core/v1"
	v11 "github.com/ericchiang/k8s/apis/meta/v1"
	_ "github.com/ericchiang/k8s/runtime"
	_ "github.com/ericchiang/k8s/runtime/schema"
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

// Endpoint represents a single logical "backend" implementing a service.
type Endpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// addresses of this endpoint. The contents of this field are interpreted
	// according to the corresponding EndpointSlice addressType field. Consumers
	// must handle different types of addresses in the context of their own
	// capabilities. This must contain at least one address but no more than
	// 100.
	// +listType=set
	Addresses []string `protobuf:"bytes,1,rep,name=addresses" json:"addresses,omitempty"`
	// conditions contains information about the current status of the endpoint.
	Conditions *EndpointConditions `protobuf:"bytes,2,opt,name=conditions" json:"conditions,omitempty"`
	// hostname of this endpoint. This field may be used by consumers of
	// endpoints to distinguish endpoints from each other (e.g. in DNS names).
	// Multiple endpoints which use the same hostname should be considered
	// fungible (e.g. multiple A values in DNS). Must be lowercase and pass
	// DNS label (RFC 1123) validation.
	// +optional
	Hostname *string `protobuf:"bytes,3,opt,name=hostname" json:"hostname,omitempty"`
	// targetRef is a reference to a Kubernetes object that represents this
	// endpoint.
	// +optional
	TargetRef *v1.ObjectReference `protobuf:"bytes,4,opt,name=targetRef" json:"targetRef,omitempty"`
	// topology contains arbitrary topology information associated with the
	// endpoint. These key/value pairs must conform with the label format.
	// https://kubernetes.io/docs/concepts/overview/working-with-objects/labels
	// Topology may include a maximum of 16 key/value pairs. This includes, but
	// is not limited to the following well known keys:
	// * kubernetes.io/hostname: the value indicates the hostname of the node
	//   where the endpoint is located. This should match the corresponding
	//   node label.
	// * topology.kubernetes.io/zone: the value indicates the zone where the
	//   endpoint is located. This should match the corresponding node label.
	// * topology.kubernetes.io/region: the value indicates the region where the
	//   endpoint is located. This should match the corresponding node label.
	// This field is deprecated and will be removed in future api versions.
	// +optional
	Topology map[string]string `protobuf:"bytes,5,rep,name=topology" json:"topology,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// nodeName represents the name of the Node hosting this endpoint. This can
	// be used to determine endpoints local to a Node. This field can be enabled
	// with the EndpointSliceNodeName feature gate.
	// +optional
	NodeName *string `protobuf:"bytes,6,opt,name=nodeName" json:"nodeName,omitempty"`
}

func (x *Endpoint) Reset() {
	*x = Endpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_api_discovery_v1alpha1_generated_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Endpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Endpoint) ProtoMessage() {}

func (x *Endpoint) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_api_discovery_v1alpha1_generated_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Endpoint.ProtoReflect.Descriptor instead.
func (*Endpoint) Descriptor() ([]byte, []int) {
	return file_k8s_io_api_discovery_v1alpha1_generated_proto_rawDescGZIP(), []int{0}
}

func (x *Endpoint) GetAddresses() []string {
	if x != nil {
		return x.Addresses
	}
	return nil
}

func (x *Endpoint) GetConditions() *EndpointConditions {
	if x != nil {
		return x.Conditions
	}
	return nil
}

func (x *Endpoint) GetHostname() string {
	if x != nil && x.Hostname != nil {
		return *x.Hostname
	}
	return ""
}

func (x *Endpoint) GetTargetRef() *v1.ObjectReference {
	if x != nil {
		return x.TargetRef
	}
	return nil
}

func (x *Endpoint) GetTopology() map[string]string {
	if x != nil {
		return x.Topology
	}
	return nil
}

func (x *Endpoint) GetNodeName() string {
	if x != nil && x.NodeName != nil {
		return *x.NodeName
	}
	return ""
}

// EndpointConditions represents the current condition of an endpoint.
type EndpointConditions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ready indicates that this endpoint is prepared to receive traffic,
	// according to whatever system is managing the endpoint. A nil value
	// indicates an unknown state. In most cases consumers should interpret this
	// unknown state as ready. For compatibility reasons, ready should never be
	// "true" for terminating endpoints.
	// +optional
	Ready *bool `protobuf:"varint,1,opt,name=ready" json:"ready,omitempty"`
	// serving is identical to ready except that it is set regardless of the
	// terminating state of endpoints. This condition should be set to true for
	// a ready endpoint that is terminating. If nil, consumers should defer to
	// the ready condition. This field can be enabled with the
	// EndpointSliceTerminatingCondition feature gate.
	// +optional
	Serving *bool `protobuf:"varint,2,opt,name=serving" json:"serving,omitempty"`
	// terminating indicates that this endpoint is terminating. A nil value
	// indicates an unknown state. Consumers should interpret this unknown state
	// to mean that the endpoint is not terminating. This field can be enabled
	// with the EndpointSliceTerminatingCondition feature gate.
	// +optional
	Terminating *bool `protobuf:"varint,3,opt,name=terminating" json:"terminating,omitempty"`
}

func (x *EndpointConditions) Reset() {
	*x = EndpointConditions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_api_discovery_v1alpha1_generated_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointConditions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointConditions) ProtoMessage() {}

func (x *EndpointConditions) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_api_discovery_v1alpha1_generated_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointConditions.ProtoReflect.Descriptor instead.
func (*EndpointConditions) Descriptor() ([]byte, []int) {
	return file_k8s_io_api_discovery_v1alpha1_generated_proto_rawDescGZIP(), []int{1}
}

func (x *EndpointConditions) GetReady() bool {
	if x != nil && x.Ready != nil {
		return *x.Ready
	}
	return false
}

func (x *EndpointConditions) GetServing() bool {
	if x != nil && x.Serving != nil {
		return *x.Serving
	}
	return false
}

func (x *EndpointConditions) GetTerminating() bool {
	if x != nil && x.Terminating != nil {
		return *x.Terminating
	}
	return false
}

// EndpointPort represents a Port used by an EndpointSlice
type EndpointPort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of this port. All ports in an EndpointSlice must have a unique
	// name. If the EndpointSlice is dervied from a Kubernetes service, this
	// corresponds to the Service.ports[].name.
	// Name must either be an empty string or pass DNS_LABEL validation:
	// * must be no more than 63 characters long.
	// * must consist of lower case alphanumeric characters or '-'.
	// * must start and end with an alphanumeric character.
	// Default is empty string.
	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The IP protocol for this port.
	// Must be UDP, TCP, or SCTP.
	// Default is TCP.
	Protocol *string `protobuf:"bytes,2,opt,name=protocol" json:"protocol,omitempty"`
	// The port number of the endpoint.
	// If this is not specified, ports are not restricted and must be
	// interpreted in the context of the specific consumer.
	Port *int32 `protobuf:"varint,3,opt,name=port" json:"port,omitempty"`
	// The application protocol for this port.
	// This field follows standard Kubernetes label syntax.
	// Un-prefixed names are reserved for IANA standard service names (as per
	// RFC-6335 and http://www.iana.org/assignments/service-names).
	// Non-standard protocols should use prefixed names.
	// Default is empty string.
	AppProtocol *string `protobuf:"bytes,4,opt,name=appProtocol" json:"appProtocol,omitempty"`
}

func (x *EndpointPort) Reset() {
	*x = EndpointPort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_api_discovery_v1alpha1_generated_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointPort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointPort) ProtoMessage() {}

func (x *EndpointPort) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_api_discovery_v1alpha1_generated_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointPort.ProtoReflect.Descriptor instead.
func (*EndpointPort) Descriptor() ([]byte, []int) {
	return file_k8s_io_api_discovery_v1alpha1_generated_proto_rawDescGZIP(), []int{2}
}

func (x *EndpointPort) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *EndpointPort) GetProtocol() string {
	if x != nil && x.Protocol != nil {
		return *x.Protocol
	}
	return ""
}

func (x *EndpointPort) GetPort() int32 {
	if x != nil && x.Port != nil {
		return *x.Port
	}
	return 0
}

func (x *EndpointPort) GetAppProtocol() string {
	if x != nil && x.AppProtocol != nil {
		return *x.AppProtocol
	}
	return ""
}

// EndpointSlice represents a subset of the endpoints that implement a service.
// For a given service there may be multiple EndpointSlice objects, selected by
// labels, which must be joined to produce the full set of endpoints.
type EndpointSlice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Standard object's metadata.
	// +optional
	Metadata *v11.ObjectMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// addressType specifies the type of address carried by this EndpointSlice.
	// All addresses in this slice must be the same type. This field is
	// immutable after creation. The following address types are currently
	// supported:
	// * IPv4: Represents an IPv4 Address.
	// * IPv6: Represents an IPv6 Address.
	// * FQDN: Represents a Fully Qualified Domain Name.
	AddressType *string `protobuf:"bytes,4,opt,name=addressType" json:"addressType,omitempty"`
	// endpoints is a list of unique endpoints in this slice. Each slice may
	// include a maximum of 1000 endpoints.
	// +listType=atomic
	Endpoints []*Endpoint `protobuf:"bytes,2,rep,name=endpoints" json:"endpoints,omitempty"`
	// ports specifies the list of network ports exposed by each endpoint in
	// this slice. Each port must have a unique name. When ports is empty, it
	// indicates that there are no defined ports. When a port is defined with a
	// nil port value, it indicates "all ports". Each slice may include a
	// maximum of 100 ports.
	// +optional
	// +listType=atomic
	Ports []*EndpointPort `protobuf:"bytes,3,rep,name=ports" json:"ports,omitempty"`
}

func (x *EndpointSlice) Reset() {
	*x = EndpointSlice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_api_discovery_v1alpha1_generated_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointSlice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointSlice) ProtoMessage() {}

func (x *EndpointSlice) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_api_discovery_v1alpha1_generated_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointSlice.ProtoReflect.Descriptor instead.
func (*EndpointSlice) Descriptor() ([]byte, []int) {
	return file_k8s_io_api_discovery_v1alpha1_generated_proto_rawDescGZIP(), []int{3}
}

func (x *EndpointSlice) GetMetadata() *v11.ObjectMeta {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *EndpointSlice) GetAddressType() string {
	if x != nil && x.AddressType != nil {
		return *x.AddressType
	}
	return ""
}

func (x *EndpointSlice) GetEndpoints() []*Endpoint {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

func (x *EndpointSlice) GetPorts() []*EndpointPort {
	if x != nil {
		return x.Ports
	}
	return nil
}

// EndpointSliceList represents a list of endpoint slices
type EndpointSliceList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Standard list metadata.
	// +optional
	Metadata *v11.ListMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// List of endpoint slices
	Items []*EndpointSlice `protobuf:"bytes,2,rep,name=items" json:"items,omitempty"`
}

func (x *EndpointSliceList) Reset() {
	*x = EndpointSliceList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_api_discovery_v1alpha1_generated_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointSliceList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointSliceList) ProtoMessage() {}

func (x *EndpointSliceList) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_api_discovery_v1alpha1_generated_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointSliceList.ProtoReflect.Descriptor instead.
func (*EndpointSliceList) Descriptor() ([]byte, []int) {
	return file_k8s_io_api_discovery_v1alpha1_generated_proto_rawDescGZIP(), []int{4}
}

func (x *EndpointSliceList) GetMetadata() *v11.ListMeta {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *EndpointSliceList) GetItems() []*EndpointSlice {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_k8s_io_api_discovery_v1alpha1_generated_proto protoreflect.FileDescriptor

var file_k8s_io_api_discovery_v1alpha1_generated_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1d, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x22,
	0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x34, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x6d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f,
	0x2f, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x6b, 0x38, 0x73, 0x2e, 0x69,
	0x6f, 0x2f, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x86, 0x03, 0x0a, 0x08, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x51, 0x0a, 0x0a,
	0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x31, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x09, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x52, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x66, 0x12, 0x51,
	0x0a, 0x08, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x35, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f,
	0x67, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67,
	0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x3b, 0x0a,
	0x0d, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x66, 0x0a, 0x12, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x61, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x72, 0x65, 0x61, 0x64, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67,
	0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x22, 0x74, 0x0a, 0x0c, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x50, 0x6f,
	0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x70, 0x70,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x89, 0x02, 0x0a, 0x0d, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6b,
	0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x72, 0x79, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x45, 0x0a, 0x09, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x12, 0x41, 0x0a, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x05, 0x70,
	0x6f, 0x72, 0x74, 0x73, 0x22, 0xa3, 0x01, 0x0a, 0x11, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6b,
	0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x72, 0x79, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x42, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x6c,
	0x69, 0x63, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x1f, 0x5a, 0x1d, 0x6b, 0x38,
	0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
}

var (
	file_k8s_io_api_discovery_v1alpha1_generated_proto_rawDescOnce sync.Once
	file_k8s_io_api_discovery_v1alpha1_generated_proto_rawDescData = file_k8s_io_api_discovery_v1alpha1_generated_proto_rawDesc
)

func file_k8s_io_api_discovery_v1alpha1_generated_proto_rawDescGZIP() []byte {
	file_k8s_io_api_discovery_v1alpha1_generated_proto_rawDescOnce.Do(func() {
		file_k8s_io_api_discovery_v1alpha1_generated_proto_rawDescData = protoimpl.X.CompressGZIP(file_k8s_io_api_discovery_v1alpha1_generated_proto_rawDescData)
	})
	return file_k8s_io_api_discovery_v1alpha1_generated_proto_rawDescData
}

var file_k8s_io_api_discovery_v1alpha1_generated_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_k8s_io_api_discovery_v1alpha1_generated_proto_goTypes = []interface{}{
	(*Endpoint)(nil),           // 0: k8s.io.api.discovery.v1alpha1.Endpoint
	(*EndpointConditions)(nil), // 1: k8s.io.api.discovery.v1alpha1.EndpointConditions
	(*EndpointPort)(nil),       // 2: k8s.io.api.discovery.v1alpha1.EndpointPort
	(*EndpointSlice)(nil),      // 3: k8s.io.api.discovery.v1alpha1.EndpointSlice
	(*EndpointSliceList)(nil),  // 4: k8s.io.api.discovery.v1alpha1.EndpointSliceList
	nil,                        // 5: k8s.io.api.discovery.v1alpha1.Endpoint.TopologyEntry
	(*v1.ObjectReference)(nil), // 6: k8s.io.api.core.v1.ObjectReference
	(*v11.ObjectMeta)(nil),     // 7: k8s.io.apimachinery.pkg.apis.meta.v1.ObjectMeta
	(*v11.ListMeta)(nil),       // 8: k8s.io.apimachinery.pkg.apis.meta.v1.ListMeta
}
var file_k8s_io_api_discovery_v1alpha1_generated_proto_depIdxs = []int32{
	1, // 0: k8s.io.api.discovery.v1alpha1.Endpoint.conditions:type_name -> k8s.io.api.discovery.v1alpha1.EndpointConditions
	6, // 1: k8s.io.api.discovery.v1alpha1.Endpoint.targetRef:type_name -> k8s.io.api.core.v1.ObjectReference
	5, // 2: k8s.io.api.discovery.v1alpha1.Endpoint.topology:type_name -> k8s.io.api.discovery.v1alpha1.Endpoint.TopologyEntry
	7, // 3: k8s.io.api.discovery.v1alpha1.EndpointSlice.metadata:type_name -> k8s.io.apimachinery.pkg.apis.meta.v1.ObjectMeta
	0, // 4: k8s.io.api.discovery.v1alpha1.EndpointSlice.endpoints:type_name -> k8s.io.api.discovery.v1alpha1.Endpoint
	2, // 5: k8s.io.api.discovery.v1alpha1.EndpointSlice.ports:type_name -> k8s.io.api.discovery.v1alpha1.EndpointPort
	8, // 6: k8s.io.api.discovery.v1alpha1.EndpointSliceList.metadata:type_name -> k8s.io.apimachinery.pkg.apis.meta.v1.ListMeta
	3, // 7: k8s.io.api.discovery.v1alpha1.EndpointSliceList.items:type_name -> k8s.io.api.discovery.v1alpha1.EndpointSlice
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_k8s_io_api_discovery_v1alpha1_generated_proto_init() }
func file_k8s_io_api_discovery_v1alpha1_generated_proto_init() {
	if File_k8s_io_api_discovery_v1alpha1_generated_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_k8s_io_api_discovery_v1alpha1_generated_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Endpoint); i {
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
		file_k8s_io_api_discovery_v1alpha1_generated_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointConditions); i {
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
		file_k8s_io_api_discovery_v1alpha1_generated_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointPort); i {
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
		file_k8s_io_api_discovery_v1alpha1_generated_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointSlice); i {
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
		file_k8s_io_api_discovery_v1alpha1_generated_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointSliceList); i {
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
			RawDescriptor: file_k8s_io_api_discovery_v1alpha1_generated_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_k8s_io_api_discovery_v1alpha1_generated_proto_goTypes,
		DependencyIndexes: file_k8s_io_api_discovery_v1alpha1_generated_proto_depIdxs,
		MessageInfos:      file_k8s_io_api_discovery_v1alpha1_generated_proto_msgTypes,
	}.Build()
	File_k8s_io_api_discovery_v1alpha1_generated_proto = out.File
	file_k8s_io_api_discovery_v1alpha1_generated_proto_rawDesc = nil
	file_k8s_io_api_discovery_v1alpha1_generated_proto_goTypes = nil
	file_k8s_io_api_discovery_v1alpha1_generated_proto_depIdxs = nil
}
