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
// source: k8s.io/api/node/v1beta1/generated.proto

package v1beta1

import (
	v11 "github.com/ericchiang/k8s/apis/core/v1"
	"github.com/ericchiang/k8s/apis/meta/v1"
	"github.com/ericchiang/k8s/apis/resource"
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

// Overhead structure represents the resource overhead associated with running a pod.
type Overhead struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// PodFixed represents the fixed resource overhead associated with running a pod.
	// +optional
	PodFixed map[string]*resource.Quantity `protobuf:"bytes,1,rep,name=podFixed" json:"podFixed,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (x *Overhead) Reset() {
	*x = Overhead{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_api_node_v1beta1_generated_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Overhead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Overhead) ProtoMessage() {}

func (x *Overhead) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_api_node_v1beta1_generated_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Overhead.ProtoReflect.Descriptor instead.
func (*Overhead) Descriptor() ([]byte, []int) {
	return file_k8s_io_api_node_v1beta1_generated_proto_rawDescGZIP(), []int{0}
}

func (x *Overhead) GetPodFixed() map[string]*resource.Quantity {
	if x != nil {
		return x.PodFixed
	}
	return nil
}

// RuntimeClass defines a class of container runtime supported in the cluster.
// The RuntimeClass is used to determine which container runtime is used to run
// all containers in a pod. RuntimeClasses are (currently) manually defined by a
// user or cluster provisioner, and referenced in the PodSpec. The Kubelet is
// responsible for resolving the RuntimeClassName reference before running the
// pod.  For more details, see
// https://git.k8s.io/enhancements/keps/sig-node/runtime-class.md
type RuntimeClass struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	Metadata *v1.ObjectMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// Handler specifies the underlying runtime and configuration that the CRI
	// implementation will use to handle pods of this class. The possible values
	// are specific to the node & CRI configuration.  It is assumed that all
	// handlers are available on every node, and handlers of the same name are
	// equivalent on every node.
	// For example, a handler called "runc" might specify that the runc OCI
	// runtime (using native Linux containers) will be used to run the containers
	// in a pod.
	// The Handler must be lowercase, conform to the DNS Label (RFC 1123) requirements,
	// and is immutable.
	Handler *string `protobuf:"bytes,2,opt,name=handler" json:"handler,omitempty"`
	// Overhead represents the resource overhead associated with running a pod for a
	// given RuntimeClass. For more details, see
	// https://git.k8s.io/enhancements/keps/sig-node/20190226-pod-overhead.md
	// This field is alpha-level as of Kubernetes v1.15, and is only honored by servers that enable the PodOverhead feature.
	// +optional
	Overhead *Overhead `protobuf:"bytes,3,opt,name=overhead" json:"overhead,omitempty"`
	// Scheduling holds the scheduling constraints to ensure that pods running
	// with this RuntimeClass are scheduled to nodes that support it.
	// If scheduling is nil, this RuntimeClass is assumed to be supported by all
	// nodes.
	// +optional
	Scheduling *Scheduling `protobuf:"bytes,4,opt,name=scheduling" json:"scheduling,omitempty"`
}

func (x *RuntimeClass) Reset() {
	*x = RuntimeClass{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_api_node_v1beta1_generated_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RuntimeClass) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuntimeClass) ProtoMessage() {}

func (x *RuntimeClass) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_api_node_v1beta1_generated_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuntimeClass.ProtoReflect.Descriptor instead.
func (*RuntimeClass) Descriptor() ([]byte, []int) {
	return file_k8s_io_api_node_v1beta1_generated_proto_rawDescGZIP(), []int{1}
}

func (x *RuntimeClass) GetMetadata() *v1.ObjectMeta {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *RuntimeClass) GetHandler() string {
	if x != nil && x.Handler != nil {
		return *x.Handler
	}
	return ""
}

func (x *RuntimeClass) GetOverhead() *Overhead {
	if x != nil {
		return x.Overhead
	}
	return nil
}

func (x *RuntimeClass) GetScheduling() *Scheduling {
	if x != nil {
		return x.Scheduling
	}
	return nil
}

// RuntimeClassList is a list of RuntimeClass objects.
type RuntimeClassList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Standard list metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	Metadata *v1.ListMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// Items is a list of schema objects.
	Items []*RuntimeClass `protobuf:"bytes,2,rep,name=items" json:"items,omitempty"`
}

func (x *RuntimeClassList) Reset() {
	*x = RuntimeClassList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_api_node_v1beta1_generated_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RuntimeClassList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuntimeClassList) ProtoMessage() {}

func (x *RuntimeClassList) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_api_node_v1beta1_generated_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuntimeClassList.ProtoReflect.Descriptor instead.
func (*RuntimeClassList) Descriptor() ([]byte, []int) {
	return file_k8s_io_api_node_v1beta1_generated_proto_rawDescGZIP(), []int{2}
}

func (x *RuntimeClassList) GetMetadata() *v1.ListMeta {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *RuntimeClassList) GetItems() []*RuntimeClass {
	if x != nil {
		return x.Items
	}
	return nil
}

// Scheduling specifies the scheduling constraints for nodes supporting a
// RuntimeClass.
type Scheduling struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// nodeSelector lists labels that must be present on nodes that support this
	// RuntimeClass. Pods using this RuntimeClass can only be scheduled to a
	// node matched by this selector. The RuntimeClass nodeSelector is merged
	// with a pod's existing nodeSelector. Any conflicts will cause the pod to
	// be rejected in admission.
	// +optional
	NodeSelector map[string]string `protobuf:"bytes,1,rep,name=nodeSelector" json:"nodeSelector,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// tolerations are appended (excluding duplicates) to pods running with this
	// RuntimeClass during admission, effectively unioning the set of nodes
	// tolerated by the pod and the RuntimeClass.
	// +optional
	// +listType=atomic
	Tolerations []*v11.Toleration `protobuf:"bytes,2,rep,name=tolerations" json:"tolerations,omitempty"`
}

func (x *Scheduling) Reset() {
	*x = Scheduling{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_api_node_v1beta1_generated_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scheduling) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scheduling) ProtoMessage() {}

func (x *Scheduling) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_api_node_v1beta1_generated_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scheduling.ProtoReflect.Descriptor instead.
func (*Scheduling) Descriptor() ([]byte, []int) {
	return file_k8s_io_api_node_v1beta1_generated_proto_rawDescGZIP(), []int{3}
}

func (x *Scheduling) GetNodeSelector() map[string]string {
	if x != nil {
		return x.NodeSelector
	}
	return nil
}

func (x *Scheduling) GetTolerations() []*v11.Toleration {
	if x != nil {
		return x.Tolerations
	}
	return nil
}

var File_k8s_io_api_node_v1beta1_generated_proto protoreflect.FileDescriptor

var file_k8s_io_api_node_v1beta1_generated_proto_rawDesc = []byte{
	0x0a, 0x27, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x6f, 0x64,
	0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x6b, 0x38, 0x73, 0x2e, 0x69,
	0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x1a, 0x22, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x6b, 0x38,
	0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72,
	0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f,
	0x76, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2f, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x6d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x6d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x75, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc4, 0x01, 0x0a, 0x08,
	0x4f, 0x76, 0x65, 0x72, 0x68, 0x65, 0x61, 0x64, 0x12, 0x4b, 0x0a, 0x08, 0x70, 0x6f, 0x64, 0x46,
	0x69, 0x78, 0x65, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x6b, 0x38, 0x73,
	0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x4f, 0x76, 0x65, 0x72, 0x68, 0x65, 0x61, 0x64, 0x2e, 0x50, 0x6f,
	0x64, 0x46, 0x69, 0x78, 0x65, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x70, 0x6f, 0x64,
	0x46, 0x69, 0x78, 0x65, 0x64, 0x1a, 0x6b, 0x0a, 0x0d, 0x50, 0x6f, 0x64, 0x46, 0x69, 0x78, 0x65,
	0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x44, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f,
	0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x6b,
	0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x51,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0xfa, 0x01, 0x0a, 0x0c, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x12, 0x4c, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61,
	0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x6b, 0x67, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x08, 0x6f,
	0x76, 0x65, 0x72, 0x68, 0x65, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4f, 0x76, 0x65, 0x72, 0x68, 0x65, 0x61, 0x64,
	0x52, 0x08, 0x6f, 0x76, 0x65, 0x72, 0x68, 0x65, 0x61, 0x64, 0x12, 0x43, 0x0a, 0x0a, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x64, 0x65,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x69, 0x6e, 0x67, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x22,
	0x9b, 0x01, 0x0a, 0x10, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e,
	0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x6b, 0x67,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x3b, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x64,
	0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xea, 0x01,
	0x0a, 0x0a, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x12, 0x59, 0x0a, 0x0c,
	0x6e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x35, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x40, 0x0a, 0x0b, 0x74, 0x6f, 0x6c, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6b,
	0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x6f, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x6f,
	0x6c, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x3f, 0x0a, 0x11, 0x4e, 0x6f, 0x64,
	0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x19, 0x5a, 0x17, 0x6b, 0x38,
	0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31,
}

var (
	file_k8s_io_api_node_v1beta1_generated_proto_rawDescOnce sync.Once
	file_k8s_io_api_node_v1beta1_generated_proto_rawDescData = file_k8s_io_api_node_v1beta1_generated_proto_rawDesc
)

func file_k8s_io_api_node_v1beta1_generated_proto_rawDescGZIP() []byte {
	file_k8s_io_api_node_v1beta1_generated_proto_rawDescOnce.Do(func() {
		file_k8s_io_api_node_v1beta1_generated_proto_rawDescData = protoimpl.X.CompressGZIP(file_k8s_io_api_node_v1beta1_generated_proto_rawDescData)
	})
	return file_k8s_io_api_node_v1beta1_generated_proto_rawDescData
}

var file_k8s_io_api_node_v1beta1_generated_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_k8s_io_api_node_v1beta1_generated_proto_goTypes = []interface{}{
	(*Overhead)(nil),          // 0: k8s.io.api.node.v1beta1.Overhead
	(*RuntimeClass)(nil),      // 1: k8s.io.api.node.v1beta1.RuntimeClass
	(*RuntimeClassList)(nil),  // 2: k8s.io.api.node.v1beta1.RuntimeClassList
	(*Scheduling)(nil),        // 3: k8s.io.api.node.v1beta1.Scheduling
	nil,                       // 4: k8s.io.api.node.v1beta1.Overhead.PodFixedEntry
	nil,                       // 5: k8s.io.api.node.v1beta1.Scheduling.NodeSelectorEntry
	(*v1.ObjectMeta)(nil),     // 6: k8s.io.apimachinery.pkg.apis.meta.v1.ObjectMeta
	(*v1.ListMeta)(nil),       // 7: k8s.io.apimachinery.pkg.apis.meta.v1.ListMeta
	(*v11.Toleration)(nil),    // 8: k8s.io.api.core.v1.Toleration
	(*resource.Quantity)(nil), // 9: k8s.io.apimachinery.pkg.api.resource.Quantity
}
var file_k8s_io_api_node_v1beta1_generated_proto_depIdxs = []int32{
	4, // 0: k8s.io.api.node.v1beta1.Overhead.podFixed:type_name -> k8s.io.api.node.v1beta1.Overhead.PodFixedEntry
	6, // 1: k8s.io.api.node.v1beta1.RuntimeClass.metadata:type_name -> k8s.io.apimachinery.pkg.apis.meta.v1.ObjectMeta
	0, // 2: k8s.io.api.node.v1beta1.RuntimeClass.overhead:type_name -> k8s.io.api.node.v1beta1.Overhead
	3, // 3: k8s.io.api.node.v1beta1.RuntimeClass.scheduling:type_name -> k8s.io.api.node.v1beta1.Scheduling
	7, // 4: k8s.io.api.node.v1beta1.RuntimeClassList.metadata:type_name -> k8s.io.apimachinery.pkg.apis.meta.v1.ListMeta
	1, // 5: k8s.io.api.node.v1beta1.RuntimeClassList.items:type_name -> k8s.io.api.node.v1beta1.RuntimeClass
	5, // 6: k8s.io.api.node.v1beta1.Scheduling.nodeSelector:type_name -> k8s.io.api.node.v1beta1.Scheduling.NodeSelectorEntry
	8, // 7: k8s.io.api.node.v1beta1.Scheduling.tolerations:type_name -> k8s.io.api.core.v1.Toleration
	9, // 8: k8s.io.api.node.v1beta1.Overhead.PodFixedEntry.value:type_name -> k8s.io.apimachinery.pkg.api.resource.Quantity
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_k8s_io_api_node_v1beta1_generated_proto_init() }
func file_k8s_io_api_node_v1beta1_generated_proto_init() {
	if File_k8s_io_api_node_v1beta1_generated_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_k8s_io_api_node_v1beta1_generated_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Overhead); i {
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
		file_k8s_io_api_node_v1beta1_generated_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RuntimeClass); i {
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
		file_k8s_io_api_node_v1beta1_generated_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RuntimeClassList); i {
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
		file_k8s_io_api_node_v1beta1_generated_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scheduling); i {
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
			RawDescriptor: file_k8s_io_api_node_v1beta1_generated_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_k8s_io_api_node_v1beta1_generated_proto_goTypes,
		DependencyIndexes: file_k8s_io_api_node_v1beta1_generated_proto_depIdxs,
		MessageInfos:      file_k8s_io_api_node_v1beta1_generated_proto_msgTypes,
	}.Build()
	File_k8s_io_api_node_v1beta1_generated_proto = out.File
	file_k8s_io_api_node_v1beta1_generated_proto_rawDesc = nil
	file_k8s_io_api_node_v1beta1_generated_proto_goTypes = nil
	file_k8s_io_api_node_v1beta1_generated_proto_depIdxs = nil
}
