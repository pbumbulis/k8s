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
// source: k8s.io/kube-aggregator/pkg/apis/apiregistration/v1beta1/generated.proto

package v1beta1

import (
	"github.com/ericchiang/k8s/apis/meta/v1"
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

// APIService represents a server for a particular GroupVersion.
// Name must be "version.group".
type APIService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *v1.ObjectMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// Spec contains information for locating and communicating with a server
	Spec *APIServiceSpec `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	// Status contains derived information about an API server
	Status *APIServiceStatus `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
}

func (x *APIService) Reset() {
	*x = APIService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APIService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APIService) ProtoMessage() {}

func (x *APIService) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APIService.ProtoReflect.Descriptor instead.
func (*APIService) Descriptor() ([]byte, []int) {
	return file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_rawDescGZIP(), []int{0}
}

func (x *APIService) GetMetadata() *v1.ObjectMeta {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *APIService) GetSpec() *APIServiceSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *APIService) GetStatus() *APIServiceStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// APIServiceCondition describes the state of an APIService at a particular point
type APIServiceCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type is the type of the condition.
	Type *string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	// Status is the status of the condition.
	// Can be True, False, Unknown.
	Status *string `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
	// Last time the condition transitioned from one status to another.
	// +optional
	LastTransitionTime *v1.Time `protobuf:"bytes,3,opt,name=lastTransitionTime" json:"lastTransitionTime,omitempty"`
	// Unique, one-word, CamelCase reason for the condition's last transition.
	// +optional
	Reason *string `protobuf:"bytes,4,opt,name=reason" json:"reason,omitempty"`
	// Human-readable message indicating details about last transition.
	// +optional
	Message *string `protobuf:"bytes,5,opt,name=message" json:"message,omitempty"`
}

func (x *APIServiceCondition) Reset() {
	*x = APIServiceCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APIServiceCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APIServiceCondition) ProtoMessage() {}

func (x *APIServiceCondition) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APIServiceCondition.ProtoReflect.Descriptor instead.
func (*APIServiceCondition) Descriptor() ([]byte, []int) {
	return file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_rawDescGZIP(), []int{1}
}

func (x *APIServiceCondition) GetType() string {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return ""
}

func (x *APIServiceCondition) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

func (x *APIServiceCondition) GetLastTransitionTime() *v1.Time {
	if x != nil {
		return x.LastTransitionTime
	}
	return nil
}

func (x *APIServiceCondition) GetReason() string {
	if x != nil && x.Reason != nil {
		return *x.Reason
	}
	return ""
}

func (x *APIServiceCondition) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

// APIServiceList is a list of APIService objects.
type APIServiceList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *v1.ListMeta  `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	Items    []*APIService `protobuf:"bytes,2,rep,name=items" json:"items,omitempty"`
}

func (x *APIServiceList) Reset() {
	*x = APIServiceList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APIServiceList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APIServiceList) ProtoMessage() {}

func (x *APIServiceList) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APIServiceList.ProtoReflect.Descriptor instead.
func (*APIServiceList) Descriptor() ([]byte, []int) {
	return file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_rawDescGZIP(), []int{2}
}

func (x *APIServiceList) GetMetadata() *v1.ListMeta {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *APIServiceList) GetItems() []*APIService {
	if x != nil {
		return x.Items
	}
	return nil
}

// APIServiceSpec contains information for locating and communicating with a server.
// Only https is supported, though you are able to disable certificate verification.
type APIServiceSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Service is a reference to the service for this API server.  It must communicate
	// on port 443.
	// If the Service is nil, that means the handling for the API groupversion is handled locally on this server.
	// The call will simply delegate to the normal handler chain to be fulfilled.
	// +optional
	Service *ServiceReference `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
	// Group is the API group name this server hosts
	Group *string `protobuf:"bytes,2,opt,name=group" json:"group,omitempty"`
	// Version is the API version this server hosts.  For example, "v1"
	Version *string `protobuf:"bytes,3,opt,name=version" json:"version,omitempty"`
	// InsecureSkipTLSVerify disables TLS certificate verification when communicating with this server.
	// This is strongly discouraged.  You should use the CABundle instead.
	InsecureSkipTLSVerify *bool `protobuf:"varint,4,opt,name=insecureSkipTLSVerify" json:"insecureSkipTLSVerify,omitempty"`
	// CABundle is a PEM encoded CA bundle which will be used to validate an API server's serving certificate.
	// If unspecified, system trust roots on the apiserver are used.
	// +listType=atomic
	// +optional
	CaBundle []byte `protobuf:"bytes,5,opt,name=caBundle" json:"caBundle,omitempty"`
	// GroupPriorityMininum is the priority this group should have at least. Higher priority means that the group is preferred by clients over lower priority ones.
	// Note that other versions of this group might specify even higher GroupPriorityMininum values such that the whole group gets a higher priority.
	// The primary sort is based on GroupPriorityMinimum, ordered highest number to lowest (20 before 10).
	// The secondary sort is based on the alphabetical comparison of the name of the object.  (v1.bar before v1.foo)
	// We'd recommend something like: *.k8s.io (except extensions) at 18000 and
	// PaaSes (OpenShift, Deis) are recommended to be in the 2000s
	GroupPriorityMinimum *int32 `protobuf:"varint,7,opt,name=groupPriorityMinimum" json:"groupPriorityMinimum,omitempty"`
	// VersionPriority controls the ordering of this API version inside of its group.  Must be greater than zero.
	// The primary sort is based on VersionPriority, ordered highest to lowest (20 before 10).
	// Since it's inside of a group, the number can be small, probably in the 10s.
	// In case of equal version priorities, the version string will be used to compute the order inside a group.
	// If the version string is "kube-like", it will sort above non "kube-like" version strings, which are ordered
	// lexicographically. "Kube-like" versions start with a "v", then are followed by a number (the major version),
	// then optionally the string "alpha" or "beta" and another number (the minor version). These are sorted first
	// by GA > beta > alpha (where GA is a version with no suffix such as beta or alpha), and then by comparing major
	// version, then minor version. An example sorted list of versions:
	// v10, v2, v1, v11beta2, v10beta3, v3beta1, v12alpha1, v11alpha2, foo1, foo10.
	VersionPriority *int32 `protobuf:"varint,8,opt,name=versionPriority" json:"versionPriority,omitempty"`
}

func (x *APIServiceSpec) Reset() {
	*x = APIServiceSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APIServiceSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APIServiceSpec) ProtoMessage() {}

func (x *APIServiceSpec) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APIServiceSpec.ProtoReflect.Descriptor instead.
func (*APIServiceSpec) Descriptor() ([]byte, []int) {
	return file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_rawDescGZIP(), []int{3}
}

func (x *APIServiceSpec) GetService() *ServiceReference {
	if x != nil {
		return x.Service
	}
	return nil
}

func (x *APIServiceSpec) GetGroup() string {
	if x != nil && x.Group != nil {
		return *x.Group
	}
	return ""
}

func (x *APIServiceSpec) GetVersion() string {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return ""
}

func (x *APIServiceSpec) GetInsecureSkipTLSVerify() bool {
	if x != nil && x.InsecureSkipTLSVerify != nil {
		return *x.InsecureSkipTLSVerify
	}
	return false
}

func (x *APIServiceSpec) GetCaBundle() []byte {
	if x != nil {
		return x.CaBundle
	}
	return nil
}

func (x *APIServiceSpec) GetGroupPriorityMinimum() int32 {
	if x != nil && x.GroupPriorityMinimum != nil {
		return *x.GroupPriorityMinimum
	}
	return 0
}

func (x *APIServiceSpec) GetVersionPriority() int32 {
	if x != nil && x.VersionPriority != nil {
		return *x.VersionPriority
	}
	return 0
}

// APIServiceStatus contains derived information about an API server
type APIServiceStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Current service state of apiService.
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []*APIServiceCondition `protobuf:"bytes,1,rep,name=conditions" json:"conditions,omitempty"`
}

func (x *APIServiceStatus) Reset() {
	*x = APIServiceStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APIServiceStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APIServiceStatus) ProtoMessage() {}

func (x *APIServiceStatus) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APIServiceStatus.ProtoReflect.Descriptor instead.
func (*APIServiceStatus) Descriptor() ([]byte, []int) {
	return file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_rawDescGZIP(), []int{4}
}

func (x *APIServiceStatus) GetConditions() []*APIServiceCondition {
	if x != nil {
		return x.Conditions
	}
	return nil
}

// ServiceReference holds a reference to Service.legacy.k8s.io
type ServiceReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Namespace is the namespace of the service
	Namespace *string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	// Name is the name of the service
	Name *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// If specified, the port on the service that hosting webhook.
	// Default to 443 for backward compatibility.
	// `port` should be a valid port number (1-65535, inclusive).
	// +optional
	Port *int32 `protobuf:"varint,3,opt,name=port" json:"port,omitempty"`
}

func (x *ServiceReference) Reset() {
	*x = ServiceReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceReference) ProtoMessage() {}

func (x *ServiceReference) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceReference.ProtoReflect.Descriptor instead.
func (*ServiceReference) Descriptor() ([]byte, []int) {
	return file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_rawDescGZIP(), []int{5}
}

func (x *ServiceReference) GetNamespace() string {
	if x != nil && x.Namespace != nil {
		return *x.Namespace
	}
	return ""
}

func (x *ServiceReference) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *ServiceReference) GetPort() int32 {
	if x != nil && x.Port != nil {
		return *x.Port
	}
	return 0
}

var File_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto protoreflect.FileDescriptor

var file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_rawDesc = []byte{
	0x0a, 0x47, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x2d, 0x61, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x37, 0x6b, 0x38, 0x73, 0x2e, 0x69,
	0x6f, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x5f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x1a, 0x34, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x6d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f,
	0x2f, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x6b, 0x38, 0x73, 0x2e, 0x69,
	0x6f, 0x2f, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x9a, 0x02, 0x0a, 0x0a, 0x41, 0x50, 0x49, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x4c, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x6d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x5b,
	0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x6b,
	0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x5f, 0x61, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x41, 0x50, 0x49, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x61, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x49, 0x2e, 0x6b, 0x38,
	0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x5f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70,
	0x69, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x41, 0x50, 0x49, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xcf,
	0x01, 0x0a, 0x13, 0x41, 0x50, 0x49, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x5a, 0x0a, 0x12, 0x6c, 0x61, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a,
	0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6d, 0x65,
	0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x12, 0x6c, 0x61, 0x73, 0x74,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0xb7, 0x01, 0x0a, 0x0e, 0x41, 0x50, 0x49, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61,
	0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x6b, 0x67, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x59, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x43,
	0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x5f, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x41, 0x50, 0x49, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xd5, 0x02, 0x0a, 0x0e, 0x41,
	0x50, 0x49, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x63, 0x0a,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x49,
	0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x5f, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x15, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x53, 0x6b,
	0x69, 0x70, 0x54, 0x4c, 0x53, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x15, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x53, 0x6b, 0x69, 0x70, 0x54,
	0x4c, 0x53, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x42, 0x75,
	0x6e, 0x64, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x63, 0x61, 0x42, 0x75,
	0x6e, 0x64, 0x6c, 0x65, 0x12, 0x32, 0x0a, 0x14, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x72, 0x69,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x4d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x14, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x4d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x12, 0x28, 0x0a, 0x0f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x22, 0x80, 0x01, 0x0a, 0x10, 0x41, 0x50, 0x49, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x6c, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4c, 0x2e, 0x6b, 0x38,
	0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x5f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70,
	0x69, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x41, 0x50, 0x49, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x58, 0x0a, 0x10, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x42,
	0x39, 0x5a, 0x37, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x2d, 0x61,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
}

var (
	file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_rawDescOnce sync.Once
	file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_rawDescData = file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_rawDesc
)

func file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_rawDescGZIP() []byte {
	file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_rawDescOnce.Do(func() {
		file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_rawDescData = protoimpl.X.CompressGZIP(file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_rawDescData)
	})
	return file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_rawDescData
}

var file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_goTypes = []interface{}{
	(*APIService)(nil),          // 0: k8s.io.kube_aggregator.pkg.apis.apiregistration.v1beta1.APIService
	(*APIServiceCondition)(nil), // 1: k8s.io.kube_aggregator.pkg.apis.apiregistration.v1beta1.APIServiceCondition
	(*APIServiceList)(nil),      // 2: k8s.io.kube_aggregator.pkg.apis.apiregistration.v1beta1.APIServiceList
	(*APIServiceSpec)(nil),      // 3: k8s.io.kube_aggregator.pkg.apis.apiregistration.v1beta1.APIServiceSpec
	(*APIServiceStatus)(nil),    // 4: k8s.io.kube_aggregator.pkg.apis.apiregistration.v1beta1.APIServiceStatus
	(*ServiceReference)(nil),    // 5: k8s.io.kube_aggregator.pkg.apis.apiregistration.v1beta1.ServiceReference
	(*v1.ObjectMeta)(nil),       // 6: k8s.io.apimachinery.pkg.apis.meta.v1.ObjectMeta
	(*v1.Time)(nil),             // 7: k8s.io.apimachinery.pkg.apis.meta.v1.Time
	(*v1.ListMeta)(nil),         // 8: k8s.io.apimachinery.pkg.apis.meta.v1.ListMeta
}
var file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_depIdxs = []int32{
	6, // 0: k8s.io.kube_aggregator.pkg.apis.apiregistration.v1beta1.APIService.metadata:type_name -> k8s.io.apimachinery.pkg.apis.meta.v1.ObjectMeta
	3, // 1: k8s.io.kube_aggregator.pkg.apis.apiregistration.v1beta1.APIService.spec:type_name -> k8s.io.kube_aggregator.pkg.apis.apiregistration.v1beta1.APIServiceSpec
	4, // 2: k8s.io.kube_aggregator.pkg.apis.apiregistration.v1beta1.APIService.status:type_name -> k8s.io.kube_aggregator.pkg.apis.apiregistration.v1beta1.APIServiceStatus
	7, // 3: k8s.io.kube_aggregator.pkg.apis.apiregistration.v1beta1.APIServiceCondition.lastTransitionTime:type_name -> k8s.io.apimachinery.pkg.apis.meta.v1.Time
	8, // 4: k8s.io.kube_aggregator.pkg.apis.apiregistration.v1beta1.APIServiceList.metadata:type_name -> k8s.io.apimachinery.pkg.apis.meta.v1.ListMeta
	0, // 5: k8s.io.kube_aggregator.pkg.apis.apiregistration.v1beta1.APIServiceList.items:type_name -> k8s.io.kube_aggregator.pkg.apis.apiregistration.v1beta1.APIService
	5, // 6: k8s.io.kube_aggregator.pkg.apis.apiregistration.v1beta1.APIServiceSpec.service:type_name -> k8s.io.kube_aggregator.pkg.apis.apiregistration.v1beta1.ServiceReference
	1, // 7: k8s.io.kube_aggregator.pkg.apis.apiregistration.v1beta1.APIServiceStatus.conditions:type_name -> k8s.io.kube_aggregator.pkg.apis.apiregistration.v1beta1.APIServiceCondition
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_init() }
func file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_init() {
	if File_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APIService); i {
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
		file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APIServiceCondition); i {
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
		file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APIServiceList); i {
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
		file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APIServiceSpec); i {
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
		file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APIServiceStatus); i {
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
		file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceReference); i {
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
			RawDescriptor: file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_goTypes,
		DependencyIndexes: file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_depIdxs,
		MessageInfos:      file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_msgTypes,
	}.Build()
	File_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto = out.File
	file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_rawDesc = nil
	file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_goTypes = nil
	file_k8s_io_kube_aggregator_pkg_apis_apiregistration_v1beta1_generated_proto_depIdxs = nil
}
