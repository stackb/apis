// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.14.0
// source: build/stack/protobuf/package/v1alpha1/package.proto

package v1alpha1

import (
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProtoRepository struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host       string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Owner      string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Name       string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Repository string `protobuf:"bytes,4,opt,name=repository,proto3" json:"repository,omitempty"`
}

func (x *ProtoRepository) Reset() {
	*x = ProtoRepository{}
	if protoimpl.UnsafeEnabled {
		mi := &file_build_stack_protobuf_package_v1alpha1_package_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoRepository) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoRepository) ProtoMessage() {}

func (x *ProtoRepository) ProtoReflect() protoreflect.Message {
	mi := &file_build_stack_protobuf_package_v1alpha1_package_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoRepository.ProtoReflect.Descriptor instead.
func (*ProtoRepository) Descriptor() ([]byte, []int) {
	return file_build_stack_protobuf_package_v1alpha1_package_proto_rawDescGZIP(), []int{0}
}

func (x *ProtoRepository) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *ProtoRepository) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *ProtoRepository) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProtoRepository) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

type ProtoSourceLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Repository *ProtoRepository `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	Prefix     string           `protobuf:"bytes,2,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Commit     string           `protobuf:"bytes,3,opt,name=commit,proto3" json:"commit,omitempty"`
	RefName    string           `protobuf:"bytes,5,opt,name=ref_name,json=refName,proto3" json:"ref_name,omitempty"`
	RefType    string           `protobuf:"bytes,6,opt,name=ref_type,json=refType,proto3" json:"ref_type,omitempty"`
}

func (x *ProtoSourceLocation) Reset() {
	*x = ProtoSourceLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_build_stack_protobuf_package_v1alpha1_package_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoSourceLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoSourceLocation) ProtoMessage() {}

func (x *ProtoSourceLocation) ProtoReflect() protoreflect.Message {
	mi := &file_build_stack_protobuf_package_v1alpha1_package_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoSourceLocation.ProtoReflect.Descriptor instead.
func (*ProtoSourceLocation) Descriptor() ([]byte, []int) {
	return file_build_stack_protobuf_package_v1alpha1_package_proto_rawDescGZIP(), []int{1}
}

func (x *ProtoSourceLocation) GetRepository() *ProtoRepository {
	if x != nil {
		return x.Repository
	}
	return nil
}

func (x *ProtoSourceLocation) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *ProtoSourceLocation) GetCommit() string {
	if x != nil {
		return x.Commit
	}
	return ""
}

func (x *ProtoSourceLocation) GetRefName() string {
	if x != nil {
		return x.RefName
	}
	return ""
}

func (x *ProtoSourceLocation) GetRefType() string {
	if x != nil {
		return x.RefType
	}
	return ""
}

type ProtoAsset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File      *descriptorpb.FileDescriptorProto `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	Hash      string                            `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Sha256    string                            `protobuf:"bytes,3,opt,name=sha256,proto3" json:"sha256,omitempty"`
	Size      uint64                            `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	AssetUrl  string                            `protobuf:"bytes,5,opt,name=asset_url,json=assetUrl,proto3" json:"asset_url,omitempty"`
	SourceUrl string                            `protobuf:"bytes,6,opt,name=source_url,json=sourceUrl,proto3" json:"source_url,omitempty"`
}

func (x *ProtoAsset) Reset() {
	*x = ProtoAsset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_build_stack_protobuf_package_v1alpha1_package_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoAsset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoAsset) ProtoMessage() {}

func (x *ProtoAsset) ProtoReflect() protoreflect.Message {
	mi := &file_build_stack_protobuf_package_v1alpha1_package_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoAsset.ProtoReflect.Descriptor instead.
func (*ProtoAsset) Descriptor() ([]byte, []int) {
	return file_build_stack_protobuf_package_v1alpha1_package_proto_rawDescGZIP(), []int{2}
}

func (x *ProtoAsset) GetFile() *descriptorpb.FileDescriptorProto {
	if x != nil {
		return x.File
	}
	return nil
}

func (x *ProtoAsset) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *ProtoAsset) GetSha256() string {
	if x != nil {
		return x.Sha256
	}
	return ""
}

func (x *ProtoAsset) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ProtoAsset) GetAssetUrl() string {
	if x != nil {
		return x.AssetUrl
	}
	return ""
}

func (x *ProtoAsset) GetSourceUrl() string {
	if x != nil {
		return x.SourceUrl
	}
	return ""
}

type ProtoCompiler struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Url     string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *ProtoCompiler) Reset() {
	*x = ProtoCompiler{}
	if protoimpl.UnsafeEnabled {
		mi := &file_build_stack_protobuf_package_v1alpha1_package_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoCompiler) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoCompiler) ProtoMessage() {}

func (x *ProtoCompiler) ProtoReflect() protoreflect.Message {
	mi := &file_build_stack_protobuf_package_v1alpha1_package_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoCompiler.ProtoReflect.Descriptor instead.
func (*ProtoCompiler) Descriptor() ([]byte, []int) {
	return file_build_stack_protobuf_package_v1alpha1_package_proto_rawDescGZIP(), []int{3}
}

func (x *ProtoCompiler) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProtoCompiler) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ProtoCompiler) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type ProtoPackage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location  *ProtoSourceLocation   `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	Compiler  *ProtoCompiler         `protobuf:"bytes,2,opt,name=compiler,proto3" json:"compiler,omitempty"`
	Assets    []*ProtoAsset          `protobuf:"bytes,4,rep,name=assets,proto3" json:"assets,omitempty"`
	Sha256    string                 `protobuf:"bytes,5,opt,name=sha256,proto3" json:"sha256,omitempty"`
	Url       string                 `protobuf:"bytes,6,opt,name=url,proto3" json:"url,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Name      string                 `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ProtoPackage) Reset() {
	*x = ProtoPackage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_build_stack_protobuf_package_v1alpha1_package_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoPackage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoPackage) ProtoMessage() {}

func (x *ProtoPackage) ProtoReflect() protoreflect.Message {
	mi := &file_build_stack_protobuf_package_v1alpha1_package_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoPackage.ProtoReflect.Descriptor instead.
func (*ProtoPackage) Descriptor() ([]byte, []int) {
	return file_build_stack_protobuf_package_v1alpha1_package_proto_rawDescGZIP(), []int{4}
}

func (x *ProtoPackage) GetLocation() *ProtoSourceLocation {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *ProtoPackage) GetCompiler() *ProtoCompiler {
	if x != nil {
		return x.Compiler
	}
	return nil
}

func (x *ProtoPackage) GetAssets() []*ProtoAsset {
	if x != nil {
		return x.Assets
	}
	return nil
}

func (x *ProtoPackage) GetSha256() string {
	if x != nil {
		return x.Sha256
	}
	return ""
}

func (x *ProtoPackage) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ProtoPackage) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ProtoPackage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateProtoPackageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Root     *ProtoSourceLocation `protobuf:"bytes,1,opt,name=root,proto3" json:"root,omitempty"`
	Compiler *ProtoCompiler       `protobuf:"bytes,2,opt,name=compiler,proto3" json:"compiler,omitempty"`
	Prefix   string               `protobuf:"bytes,3,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Assets   []*ProtoAsset        `protobuf:"bytes,4,rep,name=assets,proto3" json:"assets,omitempty"`
}

func (x *CreateProtoPackageRequest) Reset() {
	*x = CreateProtoPackageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_build_stack_protobuf_package_v1alpha1_package_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProtoPackageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProtoPackageRequest) ProtoMessage() {}

func (x *CreateProtoPackageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_build_stack_protobuf_package_v1alpha1_package_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProtoPackageRequest.ProtoReflect.Descriptor instead.
func (*CreateProtoPackageRequest) Descriptor() ([]byte, []int) {
	return file_build_stack_protobuf_package_v1alpha1_package_proto_rawDescGZIP(), []int{5}
}

func (x *CreateProtoPackageRequest) GetRoot() *ProtoSourceLocation {
	if x != nil {
		return x.Root
	}
	return nil
}

func (x *CreateProtoPackageRequest) GetCompiler() *ProtoCompiler {
	if x != nil {
		return x.Compiler
	}
	return nil
}

func (x *CreateProtoPackageRequest) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *CreateProtoPackageRequest) GetAssets() []*ProtoAsset {
	if x != nil {
		return x.Assets
	}
	return nil
}

type GetProtoPackageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetProtoPackageRequest) Reset() {
	*x = GetProtoPackageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_build_stack_protobuf_package_v1alpha1_package_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProtoPackageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProtoPackageRequest) ProtoMessage() {}

func (x *GetProtoPackageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_build_stack_protobuf_package_v1alpha1_package_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProtoPackageRequest.ProtoReflect.Descriptor instead.
func (*GetProtoPackageRequest) Descriptor() ([]byte, []int) {
	return file_build_stack_protobuf_package_v1alpha1_package_proto_rawDescGZIP(), []int{6}
}

func (x *GetProtoPackageRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_build_stack_protobuf_package_v1alpha1_package_proto protoreflect.FileDescriptor

var file_build_stack_protobuf_package_v1alpha1_package_proto_rawDesc = []byte{
	0x0a, 0x33, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x73, 0x74, 0x61,
	0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e,
	0x69, 0x6e, 0x67, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6f, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x6f, 0x72, 0x79, 0x22, 0xd3, 0x01, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x56, 0x0a,
	0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x36, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x16, 0x0a,
	0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x65, 0x66, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x66, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x72, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x66, 0x54, 0x79, 0x70, 0x65, 0x22, 0xc2, 0x01, 0x0a, 0x0a,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x38, 0x0a, 0x04, 0x66, 0x69,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x04,
	0x66, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x61, 0x32,
	0x35, 0x36, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x73, 0x73, 0x65, 0x74, 0x55, 0x72,
	0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x72, 0x6c,
	0x22, 0x4f, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x22, 0xfc, 0x02, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x12, 0x56, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x73, 0x74, 0x61,
	0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x50, 0x0a, 0x08, 0x63, 0x6f,
	0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c,
	0x65, 0x72, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x12, 0x49, 0x0a, 0x06,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52,
	0x06, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x61, 0x32, 0x35,
	0x36, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0xa0, 0x02, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4e,
	0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x12, 0x50,
	0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x34, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x43, 0x6f,
	0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x49, 0x0a, 0x06, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x06, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x73, 0x22, 0x28, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0x95, 0x02,
	0x0a, 0x08, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x12, 0x7c, 0x0a, 0x12, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x12, 0x40, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6c, 0x6f, 0x6e, 0x67,
	0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x03, 0x90, 0x02, 0x02, 0x28, 0x01, 0x12, 0x8a, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x3d, 0x2e, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x22, 0x03, 0x90, 0x02, 0x01, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_build_stack_protobuf_package_v1alpha1_package_proto_rawDescOnce sync.Once
	file_build_stack_protobuf_package_v1alpha1_package_proto_rawDescData = file_build_stack_protobuf_package_v1alpha1_package_proto_rawDesc
)

func file_build_stack_protobuf_package_v1alpha1_package_proto_rawDescGZIP() []byte {
	file_build_stack_protobuf_package_v1alpha1_package_proto_rawDescOnce.Do(func() {
		file_build_stack_protobuf_package_v1alpha1_package_proto_rawDescData = protoimpl.X.CompressGZIP(file_build_stack_protobuf_package_v1alpha1_package_proto_rawDescData)
	})
	return file_build_stack_protobuf_package_v1alpha1_package_proto_rawDescData
}

var file_build_stack_protobuf_package_v1alpha1_package_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_build_stack_protobuf_package_v1alpha1_package_proto_goTypes = []interface{}{
	(*ProtoRepository)(nil),                  // 0: build.stack.protobuf.package.v1alpha1.ProtoRepository
	(*ProtoSourceLocation)(nil),              // 1: build.stack.protobuf.package.v1alpha1.ProtoSourceLocation
	(*ProtoAsset)(nil),                       // 2: build.stack.protobuf.package.v1alpha1.ProtoAsset
	(*ProtoCompiler)(nil),                    // 3: build.stack.protobuf.package.v1alpha1.ProtoCompiler
	(*ProtoPackage)(nil),                     // 4: build.stack.protobuf.package.v1alpha1.ProtoPackage
	(*CreateProtoPackageRequest)(nil),        // 5: build.stack.protobuf.package.v1alpha1.CreateProtoPackageRequest
	(*GetProtoPackageRequest)(nil),           // 6: build.stack.protobuf.package.v1alpha1.GetProtoPackageRequest
	(*descriptorpb.FileDescriptorProto)(nil), // 7: google.protobuf.FileDescriptorProto
	(*timestamppb.Timestamp)(nil),            // 8: google.protobuf.Timestamp
	(*longrunningpb.Operation)(nil),          // 9: google.longrunning.Operation
}
var file_build_stack_protobuf_package_v1alpha1_package_proto_depIdxs = []int32{
	0,  // 0: build.stack.protobuf.package.v1alpha1.ProtoSourceLocation.repository:type_name -> build.stack.protobuf.package.v1alpha1.ProtoRepository
	7,  // 1: build.stack.protobuf.package.v1alpha1.ProtoAsset.file:type_name -> google.protobuf.FileDescriptorProto
	1,  // 2: build.stack.protobuf.package.v1alpha1.ProtoPackage.location:type_name -> build.stack.protobuf.package.v1alpha1.ProtoSourceLocation
	3,  // 3: build.stack.protobuf.package.v1alpha1.ProtoPackage.compiler:type_name -> build.stack.protobuf.package.v1alpha1.ProtoCompiler
	2,  // 4: build.stack.protobuf.package.v1alpha1.ProtoPackage.assets:type_name -> build.stack.protobuf.package.v1alpha1.ProtoAsset
	8,  // 5: build.stack.protobuf.package.v1alpha1.ProtoPackage.created_at:type_name -> google.protobuf.Timestamp
	1,  // 6: build.stack.protobuf.package.v1alpha1.CreateProtoPackageRequest.root:type_name -> build.stack.protobuf.package.v1alpha1.ProtoSourceLocation
	3,  // 7: build.stack.protobuf.package.v1alpha1.CreateProtoPackageRequest.compiler:type_name -> build.stack.protobuf.package.v1alpha1.ProtoCompiler
	2,  // 8: build.stack.protobuf.package.v1alpha1.CreateProtoPackageRequest.assets:type_name -> build.stack.protobuf.package.v1alpha1.ProtoAsset
	5,  // 9: build.stack.protobuf.package.v1alpha1.Packages.CreateProtoPackage:input_type -> build.stack.protobuf.package.v1alpha1.CreateProtoPackageRequest
	6,  // 10: build.stack.protobuf.package.v1alpha1.Packages.GetProtoPackage:input_type -> build.stack.protobuf.package.v1alpha1.GetProtoPackageRequest
	9,  // 11: build.stack.protobuf.package.v1alpha1.Packages.CreateProtoPackage:output_type -> google.longrunning.Operation
	4,  // 12: build.stack.protobuf.package.v1alpha1.Packages.GetProtoPackage:output_type -> build.stack.protobuf.package.v1alpha1.ProtoPackage
	11, // [11:13] is the sub-list for method output_type
	9,  // [9:11] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_build_stack_protobuf_package_v1alpha1_package_proto_init() }
func file_build_stack_protobuf_package_v1alpha1_package_proto_init() {
	if File_build_stack_protobuf_package_v1alpha1_package_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_build_stack_protobuf_package_v1alpha1_package_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoRepository); i {
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
		file_build_stack_protobuf_package_v1alpha1_package_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoSourceLocation); i {
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
		file_build_stack_protobuf_package_v1alpha1_package_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoAsset); i {
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
		file_build_stack_protobuf_package_v1alpha1_package_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoCompiler); i {
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
		file_build_stack_protobuf_package_v1alpha1_package_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoPackage); i {
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
		file_build_stack_protobuf_package_v1alpha1_package_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProtoPackageRequest); i {
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
		file_build_stack_protobuf_package_v1alpha1_package_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProtoPackageRequest); i {
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
			RawDescriptor: file_build_stack_protobuf_package_v1alpha1_package_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_build_stack_protobuf_package_v1alpha1_package_proto_goTypes,
		DependencyIndexes: file_build_stack_protobuf_package_v1alpha1_package_proto_depIdxs,
		MessageInfos:      file_build_stack_protobuf_package_v1alpha1_package_proto_msgTypes,
	}.Build()
	File_build_stack_protobuf_package_v1alpha1_package_proto = out.File
	file_build_stack_protobuf_package_v1alpha1_package_proto_rawDesc = nil
	file_build_stack_protobuf_package_v1alpha1_package_proto_goTypes = nil
	file_build_stack_protobuf_package_v1alpha1_package_proto_depIdxs = nil
}
