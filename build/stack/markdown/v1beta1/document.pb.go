// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.14.0
// source: build/stack/markdown/v1beta1/document.proto

package v1beta1

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

type Document struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Title    string     `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Headings []*Heading `protobuf:"bytes,3,rep,name=headings,proto3" json:"headings,omitempty"`
	Blocks   []*Block   `protobuf:"bytes,4,rep,name=blocks,proto3" json:"blocks,omitempty"`
	Html     string     `protobuf:"bytes,5,opt,name=html,proto3" json:"html,omitempty"`
}

func (x *Document) Reset() {
	*x = Document{}
	if protoimpl.UnsafeEnabled {
		mi := &file_build_stack_markdown_v1beta1_document_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Document) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Document) ProtoMessage() {}

func (x *Document) ProtoReflect() protoreflect.Message {
	mi := &file_build_stack_markdown_v1beta1_document_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Document.ProtoReflect.Descriptor instead.
func (*Document) Descriptor() ([]byte, []int) {
	return file_build_stack_markdown_v1beta1_document_proto_rawDescGZIP(), []int{0}
}

func (x *Document) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Document) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Document) GetHeadings() []*Heading {
	if x != nil {
		return x.Headings
	}
	return nil
}

func (x *Document) GetBlocks() []*Block {
	if x != nil {
		return x.Blocks
	}
	return nil
}

func (x *Document) GetHtml() string {
	if x != nil {
		return x.Html
	}
	return ""
}

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Content:
	//	*Block_Unknown_
	//	*Block_Para_
	Content isBlock_Content `protobuf_oneof:"content"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_build_stack_markdown_v1beta1_document_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_build_stack_markdown_v1beta1_document_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_build_stack_markdown_v1beta1_document_proto_rawDescGZIP(), []int{1}
}

func (m *Block) GetContent() isBlock_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *Block) GetUnknown() *Block_Unknown {
	if x, ok := x.GetContent().(*Block_Unknown_); ok {
		return x.Unknown
	}
	return nil
}

func (x *Block) GetPara() *Block_Para {
	if x, ok := x.GetContent().(*Block_Para_); ok {
		return x.Para
	}
	return nil
}

type isBlock_Content interface {
	isBlock_Content()
}

type Block_Unknown_ struct {
	Unknown *Block_Unknown `protobuf:"bytes,1,opt,name=unknown,proto3,oneof"`
}

type Block_Para_ struct {
	Para *Block_Para `protobuf:"bytes,2,opt,name=para,proto3,oneof"`
}

func (*Block_Unknown_) isBlock_Content() {}

func (*Block_Para_) isBlock_Content() {}

type Heading struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text  string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Level int32  `protobuf:"varint,3,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *Heading) Reset() {
	*x = Heading{}
	if protoimpl.UnsafeEnabled {
		mi := &file_build_stack_markdown_v1beta1_document_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Heading) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Heading) ProtoMessage() {}

func (x *Heading) ProtoReflect() protoreflect.Message {
	mi := &file_build_stack_markdown_v1beta1_document_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Heading.ProtoReflect.Descriptor instead.
func (*Heading) Descriptor() ([]byte, []int) {
	return file_build_stack_markdown_v1beta1_document_proto_rawDescGZIP(), []int{2}
}

func (x *Heading) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Heading) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Heading) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

type RenderDocumentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source      []byte `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	LineNumbers bool   `protobuf:"varint,2,opt,name=line_numbers,json=lineNumbers,proto3" json:"line_numbers,omitempty"`
}

func (x *RenderDocumentRequest) Reset() {
	*x = RenderDocumentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_build_stack_markdown_v1beta1_document_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenderDocumentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenderDocumentRequest) ProtoMessage() {}

func (x *RenderDocumentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_build_stack_markdown_v1beta1_document_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenderDocumentRequest.ProtoReflect.Descriptor instead.
func (*RenderDocumentRequest) Descriptor() ([]byte, []int) {
	return file_build_stack_markdown_v1beta1_document_proto_rawDescGZIP(), []int{3}
}

func (x *RenderDocumentRequest) GetSource() []byte {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *RenderDocumentRequest) GetLineNumbers() bool {
	if x != nil {
		return x.LineNumbers
	}
	return false
}

type Block_Unknown struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Block_Unknown) Reset() {
	*x = Block_Unknown{}
	if protoimpl.UnsafeEnabled {
		mi := &file_build_stack_markdown_v1beta1_document_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block_Unknown) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block_Unknown) ProtoMessage() {}

func (x *Block_Unknown) ProtoReflect() protoreflect.Message {
	mi := &file_build_stack_markdown_v1beta1_document_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block_Unknown.ProtoReflect.Descriptor instead.
func (*Block_Unknown) Descriptor() ([]byte, []int) {
	return file_build_stack_markdown_v1beta1_document_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Block_Unknown) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type Block_Para struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Block_Para) Reset() {
	*x = Block_Para{}
	if protoimpl.UnsafeEnabled {
		mi := &file_build_stack_markdown_v1beta1_document_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block_Para) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block_Para) ProtoMessage() {}

func (x *Block_Para) ProtoReflect() protoreflect.Message {
	mi := &file_build_stack_markdown_v1beta1_document_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block_Para.ProtoReflect.Descriptor instead.
func (*Block_Para) Descriptor() ([]byte, []int) {
	return file_build_stack_markdown_v1beta1_document_proto_rawDescGZIP(), []int{1, 1}
}

func (x *Block_Para) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_build_stack_markdown_v1beta1_document_proto protoreflect.FileDescriptor

var file_build_stack_markdown_v1beta1_document_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x6d, 0x61,
	0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x64,
	0x6f, 0x77, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x22, 0xc8, 0x01, 0x0a, 0x08,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x41, 0x0a, 0x08, 0x68, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x73, 0x74, 0x61,
	0x63, 0x6b, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x68, 0x65, 0x61,
	0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x3b, 0x0a, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x74, 0x6d, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x68, 0x74, 0x6d, 0x6c, 0x22, 0xd6, 0x01, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x12, 0x47, 0x0a, 0x07, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2b, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e,
	0x6d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x48, 0x00,
	0x52, 0x07, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x12, 0x3e, 0x0a, 0x04, 0x70, 0x61, 0x72,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x50, 0x61, 0x72,
	0x61, 0x48, 0x00, 0x52, 0x04, 0x70, 0x61, 0x72, 0x61, 0x1a, 0x1d, 0x0a, 0x07, 0x55, 0x6e, 0x6b,
	0x6e, 0x6f, 0x77, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x1a, 0x1a, 0x0a, 0x04, 0x50, 0x61, 0x72, 0x61,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22,
	0x47, 0x0a, 0x07, 0x48, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x52, 0x0a, 0x15, 0x52, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x69, 0x6e,
	0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x32, 0x7c, 0x0a, 0x09,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x6f, 0x0a, 0x0e, 0x52, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x33, 0x2e, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x64, 0x6f,
	0x77, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6d,
	0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x62, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_build_stack_markdown_v1beta1_document_proto_rawDescOnce sync.Once
	file_build_stack_markdown_v1beta1_document_proto_rawDescData = file_build_stack_markdown_v1beta1_document_proto_rawDesc
)

func file_build_stack_markdown_v1beta1_document_proto_rawDescGZIP() []byte {
	file_build_stack_markdown_v1beta1_document_proto_rawDescOnce.Do(func() {
		file_build_stack_markdown_v1beta1_document_proto_rawDescData = protoimpl.X.CompressGZIP(file_build_stack_markdown_v1beta1_document_proto_rawDescData)
	})
	return file_build_stack_markdown_v1beta1_document_proto_rawDescData
}

var file_build_stack_markdown_v1beta1_document_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_build_stack_markdown_v1beta1_document_proto_goTypes = []interface{}{
	(*Document)(nil),              // 0: build.stack.markdown.v1beta1.Document
	(*Block)(nil),                 // 1: build.stack.markdown.v1beta1.Block
	(*Heading)(nil),               // 2: build.stack.markdown.v1beta1.Heading
	(*RenderDocumentRequest)(nil), // 3: build.stack.markdown.v1beta1.RenderDocumentRequest
	(*Block_Unknown)(nil),         // 4: build.stack.markdown.v1beta1.Block.Unknown
	(*Block_Para)(nil),            // 5: build.stack.markdown.v1beta1.Block.Para
}
var file_build_stack_markdown_v1beta1_document_proto_depIdxs = []int32{
	2, // 0: build.stack.markdown.v1beta1.Document.headings:type_name -> build.stack.markdown.v1beta1.Heading
	1, // 1: build.stack.markdown.v1beta1.Document.blocks:type_name -> build.stack.markdown.v1beta1.Block
	4, // 2: build.stack.markdown.v1beta1.Block.unknown:type_name -> build.stack.markdown.v1beta1.Block.Unknown
	5, // 3: build.stack.markdown.v1beta1.Block.para:type_name -> build.stack.markdown.v1beta1.Block.Para
	3, // 4: build.stack.markdown.v1beta1.Documents.RenderDocument:input_type -> build.stack.markdown.v1beta1.RenderDocumentRequest
	0, // 5: build.stack.markdown.v1beta1.Documents.RenderDocument:output_type -> build.stack.markdown.v1beta1.Document
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_build_stack_markdown_v1beta1_document_proto_init() }
func file_build_stack_markdown_v1beta1_document_proto_init() {
	if File_build_stack_markdown_v1beta1_document_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_build_stack_markdown_v1beta1_document_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Document); i {
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
		file_build_stack_markdown_v1beta1_document_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
		file_build_stack_markdown_v1beta1_document_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Heading); i {
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
		file_build_stack_markdown_v1beta1_document_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenderDocumentRequest); i {
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
		file_build_stack_markdown_v1beta1_document_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block_Unknown); i {
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
		file_build_stack_markdown_v1beta1_document_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block_Para); i {
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
	file_build_stack_markdown_v1beta1_document_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Block_Unknown_)(nil),
		(*Block_Para_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_build_stack_markdown_v1beta1_document_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_build_stack_markdown_v1beta1_document_proto_goTypes,
		DependencyIndexes: file_build_stack_markdown_v1beta1_document_proto_depIdxs,
		MessageInfos:      file_build_stack_markdown_v1beta1_document_proto_msgTypes,
	}.Build()
	File_build_stack_markdown_v1beta1_document_proto = out.File
	file_build_stack_markdown_v1beta1_document_proto_rawDesc = nil
	file_build_stack_markdown_v1beta1_document_proto_goTypes = nil
	file_build_stack_markdown_v1beta1_document_proto_depIdxs = nil
}
