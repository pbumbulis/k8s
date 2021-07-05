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
// source: k8s.io/apimachinery/pkg/apis/meta/v1beta1/generated.proto

package v1beta1

import (
	"github.com/ericchiang/k8s/apis/meta/v1"
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

// PartialObjectMetadataList contains a list of objects containing only their metadata.
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type PartialObjectMetadataList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Standard list metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// +optional
	Metadata *v1.ListMeta `protobuf:"bytes,2,opt,name=metadata" json:"metadata,omitempty"`
	// items contains each of the included items.
	Items []*v1.PartialObjectMetadata `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (x *PartialObjectMetadataList) Reset() {
	*x = PartialObjectMetadataList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_apimachinery_pkg_apis_meta_v1beta1_generated_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartialObjectMetadataList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartialObjectMetadataList) ProtoMessage() {}

func (x *PartialObjectMetadataList) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_apimachinery_pkg_apis_meta_v1beta1_generated_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartialObjectMetadataList.ProtoReflect.Descriptor instead.
func (*PartialObjectMetadataList) Descriptor() ([]byte, []int) {
	return file_k8s_io_apimachinery_pkg_apis_meta_v1beta1_generated_proto_rawDescGZIP(), []int{0}
}

func (x *PartialObjectMetadataList) GetMetadata() *v1.ListMeta {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *PartialObjectMetadataList) GetItems() []*v1.PartialObjectMetadata {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_k8s_io_apimachinery_pkg_apis_meta_v1beta1_generated_proto protoreflect.FileDescriptor

var file_k8s_io_apimachinery_pkg_apis_meta_v1beta1_generated_proto_rawDesc = []byte{
	0x0a, 0x39, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d,
	0x65, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x29, 0x6b, 0x38, 0x73,
	0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79,
	0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x34, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x6b, 0x38,
	0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72,
	0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x01, 0x0a, 0x19, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x4a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70,
	0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x51,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e,
	0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x72, 0x79, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x42, 0x2b, 0x5a, 0x29, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x6d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
}

var (
	file_k8s_io_apimachinery_pkg_apis_meta_v1beta1_generated_proto_rawDescOnce sync.Once
	file_k8s_io_apimachinery_pkg_apis_meta_v1beta1_generated_proto_rawDescData = file_k8s_io_apimachinery_pkg_apis_meta_v1beta1_generated_proto_rawDesc
)

func file_k8s_io_apimachinery_pkg_apis_meta_v1beta1_generated_proto_rawDescGZIP() []byte {
	file_k8s_io_apimachinery_pkg_apis_meta_v1beta1_generated_proto_rawDescOnce.Do(func() {
		file_k8s_io_apimachinery_pkg_apis_meta_v1beta1_generated_proto_rawDescData = protoimpl.X.CompressGZIP(file_k8s_io_apimachinery_pkg_apis_meta_v1beta1_generated_proto_rawDescData)
	})
	return file_k8s_io_apimachinery_pkg_apis_meta_v1beta1_generated_proto_rawDescData
}

var file_k8s_io_apimachinery_pkg_apis_meta_v1beta1_generated_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_k8s_io_apimachinery_pkg_apis_meta_v1beta1_generated_proto_goTypes = []interface{}{
	(*PartialObjectMetadataList)(nil), // 0: k8s.io.apimachinery.pkg.apis.meta.v1beta1.PartialObjectMetadataList
	(*v1.ListMeta)(nil),               // 1: k8s.io.apimachinery.pkg.apis.meta.v1.ListMeta
	(*v1.PartialObjectMetadata)(nil),  // 2: k8s.io.apimachinery.pkg.apis.meta.v1.PartialObjectMetadata
}
var file_k8s_io_apimachinery_pkg_apis_meta_v1beta1_generated_proto_depIdxs = []int32{
	1, // 0: k8s.io.apimachinery.pkg.apis.meta.v1beta1.PartialObjectMetadataList.metadata:type_name -> k8s.io.apimachinery.pkg.apis.meta.v1.ListMeta
	2, // 1: k8s.io.apimachinery.pkg.apis.meta.v1beta1.PartialObjectMetadataList.items:type_name -> k8s.io.apimachinery.pkg.apis.meta.v1.PartialObjectMetadata
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_k8s_io_apimachinery_pkg_apis_meta_v1beta1_generated_proto_init() }
func file_k8s_io_apimachinery_pkg_apis_meta_v1beta1_generated_proto_init() {
	if File_k8s_io_apimachinery_pkg_apis_meta_v1beta1_generated_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_k8s_io_apimachinery_pkg_apis_meta_v1beta1_generated_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartialObjectMetadataList); i {
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
			RawDescriptor: file_k8s_io_apimachinery_pkg_apis_meta_v1beta1_generated_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_k8s_io_apimachinery_pkg_apis_meta_v1beta1_generated_proto_goTypes,
		DependencyIndexes: file_k8s_io_apimachinery_pkg_apis_meta_v1beta1_generated_proto_depIdxs,
		MessageInfos:      file_k8s_io_apimachinery_pkg_apis_meta_v1beta1_generated_proto_msgTypes,
	}.Build()
	File_k8s_io_apimachinery_pkg_apis_meta_v1beta1_generated_proto = out.File
	file_k8s_io_apimachinery_pkg_apis_meta_v1beta1_generated_proto_rawDesc = nil
	file_k8s_io_apimachinery_pkg_apis_meta_v1beta1_generated_proto_goTypes = nil
	file_k8s_io_apimachinery_pkg_apis_meta_v1beta1_generated_proto_depIdxs = nil
}
