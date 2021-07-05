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
// source: k8s.io/api/admission/v1beta1/generated.proto

package v1beta1

import (
	v11 "github.com/ericchiang/k8s/apis/authentication/v1"
	"github.com/ericchiang/k8s/apis/meta/v1"
	"github.com/ericchiang/k8s/runtime"
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

// AdmissionRequest describes the admission.Attributes for the admission request.
type AdmissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// UID is an identifier for the individual request/response. It allows us to distinguish instances of requests which are
	// otherwise identical (parallel requests, requests when earlier requests did not modify etc)
	// The UID is meant to track the round trip (request/response) between the KAS and the WebHook, not the user request.
	// It is suitable for correlating log entries between the webhook and apiserver, for either auditing or debugging.
	Uid *string `protobuf:"bytes,1,opt,name=uid" json:"uid,omitempty"`
	// Kind is the fully-qualified type of object being submitted (for example, v1.Pod or autoscaling.v1.Scale)
	Kind *v1.GroupVersionKind `protobuf:"bytes,2,opt,name=kind" json:"kind,omitempty"`
	// Resource is the fully-qualified resource being requested (for example, v1.pods)
	Resource *v1.GroupVersionResource `protobuf:"bytes,3,opt,name=resource" json:"resource,omitempty"`
	// SubResource is the subresource being requested, if any (for example, "status" or "scale")
	// +optional
	SubResource *string `protobuf:"bytes,4,opt,name=subResource" json:"subResource,omitempty"`
	// RequestKind is the fully-qualified type of the original API request (for example, v1.Pod or autoscaling.v1.Scale).
	// If this is specified and differs from the value in "kind", an equivalent match and conversion was performed.
	//
	// For example, if deployments can be modified via apps/v1 and apps/v1beta1, and a webhook registered a rule of
	// `apiGroups:["apps"], apiVersions:["v1"], resources: ["deployments"]` and `matchPolicy: Equivalent`,
	// an API request to apps/v1beta1 deployments would be converted and sent to the webhook
	// with `kind: {group:"apps", version:"v1", kind:"Deployment"}` (matching the rule the webhook registered for),
	// and `requestKind: {group:"apps", version:"v1beta1", kind:"Deployment"}` (indicating the kind of the original API request).
	//
	// See documentation for the "matchPolicy" field in the webhook configuration type for more details.
	// +optional
	RequestKind *v1.GroupVersionKind `protobuf:"bytes,13,opt,name=requestKind" json:"requestKind,omitempty"`
	// RequestResource is the fully-qualified resource of the original API request (for example, v1.pods).
	// If this is specified and differs from the value in "resource", an equivalent match and conversion was performed.
	//
	// For example, if deployments can be modified via apps/v1 and apps/v1beta1, and a webhook registered a rule of
	// `apiGroups:["apps"], apiVersions:["v1"], resources: ["deployments"]` and `matchPolicy: Equivalent`,
	// an API request to apps/v1beta1 deployments would be converted and sent to the webhook
	// with `resource: {group:"apps", version:"v1", resource:"deployments"}` (matching the resource the webhook registered for),
	// and `requestResource: {group:"apps", version:"v1beta1", resource:"deployments"}` (indicating the resource of the original API request).
	//
	// See documentation for the "matchPolicy" field in the webhook configuration type.
	// +optional
	RequestResource *v1.GroupVersionResource `protobuf:"bytes,14,opt,name=requestResource" json:"requestResource,omitempty"`
	// RequestSubResource is the name of the subresource of the original API request, if any (for example, "status" or "scale")
	// If this is specified and differs from the value in "subResource", an equivalent match and conversion was performed.
	// See documentation for the "matchPolicy" field in the webhook configuration type.
	// +optional
	RequestSubResource *string `protobuf:"bytes,15,opt,name=requestSubResource" json:"requestSubResource,omitempty"`
	// Name is the name of the object as presented in the request.  On a CREATE operation, the client may omit name and
	// rely on the server to generate the name.  If that is the case, this field will contain an empty string.
	// +optional
	Name *string `protobuf:"bytes,5,opt,name=name" json:"name,omitempty"`
	// Namespace is the namespace associated with the request (if any).
	// +optional
	Namespace *string `protobuf:"bytes,6,opt,name=namespace" json:"namespace,omitempty"`
	// Operation is the operation being performed. This may be different than the operation
	// requested. e.g. a patch can result in either a CREATE or UPDATE Operation.
	Operation *string `protobuf:"bytes,7,opt,name=operation" json:"operation,omitempty"`
	// UserInfo is information about the requesting user
	UserInfo *v11.UserInfo `protobuf:"bytes,8,opt,name=userInfo" json:"userInfo,omitempty"`
	// Object is the object from the incoming request.
	// +optional
	Object *runtime.RawExtension `protobuf:"bytes,9,opt,name=object" json:"object,omitempty"`
	// OldObject is the existing object. Only populated for DELETE and UPDATE requests.
	// +optional
	OldObject *runtime.RawExtension `protobuf:"bytes,10,opt,name=oldObject" json:"oldObject,omitempty"`
	// DryRun indicates that modifications will definitely not be persisted for this request.
	// Defaults to false.
	// +optional
	DryRun *bool `protobuf:"varint,11,opt,name=dryRun" json:"dryRun,omitempty"`
	// Options is the operation option structure of the operation being performed.
	// e.g. `meta.k8s.io/v1.DeleteOptions` or `meta.k8s.io/v1.CreateOptions`. This may be
	// different than the options the caller provided. e.g. for a patch request the performed
	// Operation might be a CREATE, in which case the Options will a
	// `meta.k8s.io/v1.CreateOptions` even though the caller provided `meta.k8s.io/v1.PatchOptions`.
	// +optional
	Options *runtime.RawExtension `protobuf:"bytes,12,opt,name=options" json:"options,omitempty"`
}

func (x *AdmissionRequest) Reset() {
	*x = AdmissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_api_admission_v1beta1_generated_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdmissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdmissionRequest) ProtoMessage() {}

func (x *AdmissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_api_admission_v1beta1_generated_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdmissionRequest.ProtoReflect.Descriptor instead.
func (*AdmissionRequest) Descriptor() ([]byte, []int) {
	return file_k8s_io_api_admission_v1beta1_generated_proto_rawDescGZIP(), []int{0}
}

func (x *AdmissionRequest) GetUid() string {
	if x != nil && x.Uid != nil {
		return *x.Uid
	}
	return ""
}

func (x *AdmissionRequest) GetKind() *v1.GroupVersionKind {
	if x != nil {
		return x.Kind
	}
	return nil
}

func (x *AdmissionRequest) GetResource() *v1.GroupVersionResource {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *AdmissionRequest) GetSubResource() string {
	if x != nil && x.SubResource != nil {
		return *x.SubResource
	}
	return ""
}

func (x *AdmissionRequest) GetRequestKind() *v1.GroupVersionKind {
	if x != nil {
		return x.RequestKind
	}
	return nil
}

func (x *AdmissionRequest) GetRequestResource() *v1.GroupVersionResource {
	if x != nil {
		return x.RequestResource
	}
	return nil
}

func (x *AdmissionRequest) GetRequestSubResource() string {
	if x != nil && x.RequestSubResource != nil {
		return *x.RequestSubResource
	}
	return ""
}

func (x *AdmissionRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *AdmissionRequest) GetNamespace() string {
	if x != nil && x.Namespace != nil {
		return *x.Namespace
	}
	return ""
}

func (x *AdmissionRequest) GetOperation() string {
	if x != nil && x.Operation != nil {
		return *x.Operation
	}
	return ""
}

func (x *AdmissionRequest) GetUserInfo() *v11.UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

func (x *AdmissionRequest) GetObject() *runtime.RawExtension {
	if x != nil {
		return x.Object
	}
	return nil
}

func (x *AdmissionRequest) GetOldObject() *runtime.RawExtension {
	if x != nil {
		return x.OldObject
	}
	return nil
}

func (x *AdmissionRequest) GetDryRun() bool {
	if x != nil && x.DryRun != nil {
		return *x.DryRun
	}
	return false
}

func (x *AdmissionRequest) GetOptions() *runtime.RawExtension {
	if x != nil {
		return x.Options
	}
	return nil
}

// AdmissionResponse describes an admission response.
type AdmissionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// UID is an identifier for the individual request/response.
	// This should be copied over from the corresponding AdmissionRequest.
	Uid *string `protobuf:"bytes,1,opt,name=uid" json:"uid,omitempty"`
	// Allowed indicates whether or not the admission request was permitted.
	Allowed *bool `protobuf:"varint,2,opt,name=allowed" json:"allowed,omitempty"`
	// Result contains extra details into why an admission request was denied.
	// This field IS NOT consulted in any way if "Allowed" is "true".
	// +optional
	Status *v1.Status `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
	// The patch body. Currently we only support "JSONPatch" which implements RFC 6902.
	// +optional
	Patch []byte `protobuf:"bytes,4,opt,name=patch" json:"patch,omitempty"`
	// The type of Patch. Currently we only allow "JSONPatch".
	// +optional
	PatchType *string `protobuf:"bytes,5,opt,name=patchType" json:"patchType,omitempty"`
	// AuditAnnotations is an unstructured key value map set by remote admission controller (e.g. error=image-blacklisted).
	// MutatingAdmissionWebhook and ValidatingAdmissionWebhook admission controller will prefix the keys with
	// admission webhook name (e.g. imagepolicy.example.com/error=image-blacklisted). AuditAnnotations will be provided by
	// the admission webhook to add additional context to the audit log for this request.
	// +optional
	AuditAnnotations map[string]string `protobuf:"bytes,6,rep,name=auditAnnotations" json:"auditAnnotations,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// warnings is a list of warning messages to return to the requesting API client.
	// Warning messages describe a problem the client making the API request should correct or be aware of.
	// Limit warnings to 120 characters if possible.
	// Warnings over 256 characters and large numbers of warnings may be truncated.
	// +optional
	Warnings []string `protobuf:"bytes,7,rep,name=warnings" json:"warnings,omitempty"`
}

func (x *AdmissionResponse) Reset() {
	*x = AdmissionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_api_admission_v1beta1_generated_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdmissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdmissionResponse) ProtoMessage() {}

func (x *AdmissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_api_admission_v1beta1_generated_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdmissionResponse.ProtoReflect.Descriptor instead.
func (*AdmissionResponse) Descriptor() ([]byte, []int) {
	return file_k8s_io_api_admission_v1beta1_generated_proto_rawDescGZIP(), []int{1}
}

func (x *AdmissionResponse) GetUid() string {
	if x != nil && x.Uid != nil {
		return *x.Uid
	}
	return ""
}

func (x *AdmissionResponse) GetAllowed() bool {
	if x != nil && x.Allowed != nil {
		return *x.Allowed
	}
	return false
}

func (x *AdmissionResponse) GetStatus() *v1.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *AdmissionResponse) GetPatch() []byte {
	if x != nil {
		return x.Patch
	}
	return nil
}

func (x *AdmissionResponse) GetPatchType() string {
	if x != nil && x.PatchType != nil {
		return *x.PatchType
	}
	return ""
}

func (x *AdmissionResponse) GetAuditAnnotations() map[string]string {
	if x != nil {
		return x.AuditAnnotations
	}
	return nil
}

func (x *AdmissionResponse) GetWarnings() []string {
	if x != nil {
		return x.Warnings
	}
	return nil
}

// AdmissionReview describes an admission review request/response.
type AdmissionReview struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Request describes the attributes for the admission request.
	// +optional
	Request *AdmissionRequest `protobuf:"bytes,1,opt,name=request" json:"request,omitempty"`
	// Response describes the attributes for the admission response.
	// +optional
	Response *AdmissionResponse `protobuf:"bytes,2,opt,name=response" json:"response,omitempty"`
}

func (x *AdmissionReview) Reset() {
	*x = AdmissionReview{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_api_admission_v1beta1_generated_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdmissionReview) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdmissionReview) ProtoMessage() {}

func (x *AdmissionReview) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_api_admission_v1beta1_generated_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdmissionReview.ProtoReflect.Descriptor instead.
func (*AdmissionReview) Descriptor() ([]byte, []int) {
	return file_k8s_io_api_admission_v1beta1_generated_proto_rawDescGZIP(), []int{2}
}

func (x *AdmissionReview) GetRequest() *AdmissionRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *AdmissionReview) GetResponse() *AdmissionResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

var File_k8s_io_api_admission_v1beta1_generated_proto protoreflect.FileDescriptor

var file_k8s_io_api_admission_v1beta1_generated_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c,
	0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x2c, 0x6b, 0x38,
	0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x6b, 0x38, 0x73, 0x2e,
	0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x76, 0x31,
	0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2f, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x36, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe3, 0x06, 0x0a, 0x10, 0x41, 0x64,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64,
	0x12, 0x4a, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36,
	0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6d, 0x65,
	0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x56, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a,
	0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6d, 0x65,
	0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x75, 0x62, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x4b, 0x69, 0x6e, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x6b, 0x38,
	0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72,
	0x79, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4b,
	0x69, 0x6e, 0x64, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4b, 0x69, 0x6e, 0x64,
	0x12, 0x64, 0x0a, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x6b, 0x38, 0x73, 0x2e,
	0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2e,
	0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x53, 0x75, 0x62, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x12, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x75, 0x62, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x42, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69,
	0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x45, 0x0a, 0x06, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6b, 0x38, 0x73,
	0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79,
	0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x52, 0x61, 0x77,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x4b, 0x0a, 0x09, 0x6f, 0x6c, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70,
	0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72,
	0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x52, 0x61, 0x77, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6f, 0x6c, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x72, 0x79, 0x52, 0x75, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x64, 0x72, 0x79, 0x52, 0x75, 0x6e, 0x12, 0x47, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f,
	0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x6b,
	0x67, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x52, 0x61, 0x77, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x8d, 0x03, 0x0a, 0x11, 0x41, 0x64, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x64, 0x12, 0x44, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2c, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x74, 0x63, 0x68,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x70, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1c, 0x0a,
	0x09, 0x70, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x71, 0x0a, 0x10, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x41, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x10, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x08, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x1a, 0x43, 0x0a, 0x15, 0x41, 0x75,
	0x64, 0x69, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0xa8, 0x01, 0x0a, 0x0f, 0x41, 0x64, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x12, 0x48, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4b, 0x0a,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2f, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x41,
	0x64, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1e, 0x5a, 0x1c, 0x6b, 0x38,
	0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
}

var (
	file_k8s_io_api_admission_v1beta1_generated_proto_rawDescOnce sync.Once
	file_k8s_io_api_admission_v1beta1_generated_proto_rawDescData = file_k8s_io_api_admission_v1beta1_generated_proto_rawDesc
)

func file_k8s_io_api_admission_v1beta1_generated_proto_rawDescGZIP() []byte {
	file_k8s_io_api_admission_v1beta1_generated_proto_rawDescOnce.Do(func() {
		file_k8s_io_api_admission_v1beta1_generated_proto_rawDescData = protoimpl.X.CompressGZIP(file_k8s_io_api_admission_v1beta1_generated_proto_rawDescData)
	})
	return file_k8s_io_api_admission_v1beta1_generated_proto_rawDescData
}

var file_k8s_io_api_admission_v1beta1_generated_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_k8s_io_api_admission_v1beta1_generated_proto_goTypes = []interface{}{
	(*AdmissionRequest)(nil),        // 0: k8s.io.api.admission.v1beta1.AdmissionRequest
	(*AdmissionResponse)(nil),       // 1: k8s.io.api.admission.v1beta1.AdmissionResponse
	(*AdmissionReview)(nil),         // 2: k8s.io.api.admission.v1beta1.AdmissionReview
	nil,                             // 3: k8s.io.api.admission.v1beta1.AdmissionResponse.AuditAnnotationsEntry
	(*v1.GroupVersionKind)(nil),     // 4: k8s.io.apimachinery.pkg.apis.meta.v1.GroupVersionKind
	(*v1.GroupVersionResource)(nil), // 5: k8s.io.apimachinery.pkg.apis.meta.v1.GroupVersionResource
	(*v11.UserInfo)(nil),            // 6: k8s.io.api.authentication.v1.UserInfo
	(*runtime.RawExtension)(nil),    // 7: k8s.io.apimachinery.pkg.runtime.RawExtension
	(*v1.Status)(nil),               // 8: k8s.io.apimachinery.pkg.apis.meta.v1.Status
}
var file_k8s_io_api_admission_v1beta1_generated_proto_depIdxs = []int32{
	4,  // 0: k8s.io.api.admission.v1beta1.AdmissionRequest.kind:type_name -> k8s.io.apimachinery.pkg.apis.meta.v1.GroupVersionKind
	5,  // 1: k8s.io.api.admission.v1beta1.AdmissionRequest.resource:type_name -> k8s.io.apimachinery.pkg.apis.meta.v1.GroupVersionResource
	4,  // 2: k8s.io.api.admission.v1beta1.AdmissionRequest.requestKind:type_name -> k8s.io.apimachinery.pkg.apis.meta.v1.GroupVersionKind
	5,  // 3: k8s.io.api.admission.v1beta1.AdmissionRequest.requestResource:type_name -> k8s.io.apimachinery.pkg.apis.meta.v1.GroupVersionResource
	6,  // 4: k8s.io.api.admission.v1beta1.AdmissionRequest.userInfo:type_name -> k8s.io.api.authentication.v1.UserInfo
	7,  // 5: k8s.io.api.admission.v1beta1.AdmissionRequest.object:type_name -> k8s.io.apimachinery.pkg.runtime.RawExtension
	7,  // 6: k8s.io.api.admission.v1beta1.AdmissionRequest.oldObject:type_name -> k8s.io.apimachinery.pkg.runtime.RawExtension
	7,  // 7: k8s.io.api.admission.v1beta1.AdmissionRequest.options:type_name -> k8s.io.apimachinery.pkg.runtime.RawExtension
	8,  // 8: k8s.io.api.admission.v1beta1.AdmissionResponse.status:type_name -> k8s.io.apimachinery.pkg.apis.meta.v1.Status
	3,  // 9: k8s.io.api.admission.v1beta1.AdmissionResponse.auditAnnotations:type_name -> k8s.io.api.admission.v1beta1.AdmissionResponse.AuditAnnotationsEntry
	0,  // 10: k8s.io.api.admission.v1beta1.AdmissionReview.request:type_name -> k8s.io.api.admission.v1beta1.AdmissionRequest
	1,  // 11: k8s.io.api.admission.v1beta1.AdmissionReview.response:type_name -> k8s.io.api.admission.v1beta1.AdmissionResponse
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_k8s_io_api_admission_v1beta1_generated_proto_init() }
func file_k8s_io_api_admission_v1beta1_generated_proto_init() {
	if File_k8s_io_api_admission_v1beta1_generated_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_k8s_io_api_admission_v1beta1_generated_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdmissionRequest); i {
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
		file_k8s_io_api_admission_v1beta1_generated_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdmissionResponse); i {
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
		file_k8s_io_api_admission_v1beta1_generated_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdmissionReview); i {
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
			RawDescriptor: file_k8s_io_api_admission_v1beta1_generated_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_k8s_io_api_admission_v1beta1_generated_proto_goTypes,
		DependencyIndexes: file_k8s_io_api_admission_v1beta1_generated_proto_depIdxs,
		MessageInfos:      file_k8s_io_api_admission_v1beta1_generated_proto_msgTypes,
	}.Build()
	File_k8s_io_api_admission_v1beta1_generated_proto = out.File
	file_k8s_io_api_admission_v1beta1_generated_proto_rawDesc = nil
	file_k8s_io_api_admission_v1beta1_generated_proto_goTypes = nil
	file_k8s_io_api_admission_v1beta1_generated_proto_depIdxs = nil
}
