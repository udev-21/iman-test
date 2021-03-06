// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: pkg/post-crud/pb/post-crud.proto

package pb

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

type Response int32

const (
	Response_statusOK  Response = 0
	Response_statusErr Response = 1
)

// Enum value maps for Response.
var (
	Response_name = map[int32]string{
		0: "statusOK",
		1: "statusErr",
	}
	Response_value = map[string]int32{
		"statusOK":  0,
		"statusErr": 1,
	}
)

func (x Response) Enum() *Response {
	p := new(Response)
	*p = x
	return p
}

func (x Response) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Response) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_post_crud_pb_post_crud_proto_enumTypes[0].Descriptor()
}

func (Response) Type() protoreflect.EnumType {
	return &file_pkg_post_crud_pb_post_crud_proto_enumTypes[0]
}

func (x Response) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Response.Descriptor instead.
func (Response) EnumDescriptor() ([]byte, []int) {
	return file_pkg_post_crud_pb_post_crud_proto_rawDescGZIP(), []int{0}
}

type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID     int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	UserID int64  `protobuf:"varint,2,opt,name=userID,proto3" json:"userID,omitempty"`
	Title  string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Body   string `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_post_crud_pb_post_crud_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_post_crud_pb_post_crud_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_pkg_post_crud_pb_post_crud_proto_rawDescGZIP(), []int{0}
}

func (x *Post) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Post) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *Post) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Post) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type CreatePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post *Post `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
}

func (x *CreatePostRequest) Reset() {
	*x = CreatePostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_post_crud_pb_post_crud_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePostRequest) ProtoMessage() {}

func (x *CreatePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_post_crud_pb_post_crud_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePostRequest.ProtoReflect.Descriptor instead.
func (*CreatePostRequest) Descriptor() ([]byte, []int) {
	return file_pkg_post_crud_pb_post_crud_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePostRequest) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type CreatePostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   Response `protobuf:"varint,1,opt,name=status,proto3,enum=postcrud.Response" json:"status,omitempty"`
	ErrorMsg *string  `protobuf:"bytes,2,opt,name=errorMsg,proto3,oneof" json:"errorMsg,omitempty"`
	ID       int64    `protobuf:"varint,3,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *CreatePostResponse) Reset() {
	*x = CreatePostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_post_crud_pb_post_crud_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePostResponse) ProtoMessage() {}

func (x *CreatePostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_post_crud_pb_post_crud_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePostResponse.ProtoReflect.Descriptor instead.
func (*CreatePostResponse) Descriptor() ([]byte, []int) {
	return file_pkg_post_crud_pb_post_crud_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePostResponse) GetStatus() Response {
	if x != nil {
		return x.Status
	}
	return Response_statusOK
}

func (x *CreatePostResponse) GetErrorMsg() string {
	if x != nil && x.ErrorMsg != nil {
		return *x.ErrorMsg
	}
	return ""
}

func (x *CreatePostResponse) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type UpdatePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post *Post `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
}

func (x *UpdatePostRequest) Reset() {
	*x = UpdatePostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_post_crud_pb_post_crud_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePostRequest) ProtoMessage() {}

func (x *UpdatePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_post_crud_pb_post_crud_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePostRequest.ProtoReflect.Descriptor instead.
func (*UpdatePostRequest) Descriptor() ([]byte, []int) {
	return file_pkg_post_crud_pb_post_crud_proto_rawDescGZIP(), []int{3}
}

func (x *UpdatePostRequest) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type DeletePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DeletePostRequest) Reset() {
	*x = DeletePostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_post_crud_pb_post_crud_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePostRequest) ProtoMessage() {}

func (x *DeletePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_post_crud_pb_post_crud_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePostRequest.ProtoReflect.Descriptor instead.
func (*DeletePostRequest) Descriptor() ([]byte, []int) {
	return file_pkg_post_crud_pb_post_crud_proto_rawDescGZIP(), []int{4}
}

func (x *DeletePostRequest) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type ReadPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *ReadPostRequest) Reset() {
	*x = ReadPostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_post_crud_pb_post_crud_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadPostRequest) ProtoMessage() {}

func (x *ReadPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_post_crud_pb_post_crud_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadPostRequest.ProtoReflect.Descriptor instead.
func (*ReadPostRequest) Descriptor() ([]byte, []int) {
	return file_pkg_post_crud_pb_post_crud_proto_rawDescGZIP(), []int{5}
}

func (x *ReadPostRequest) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type RequestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   Response `protobuf:"varint,1,opt,name=status,proto3,enum=postcrud.Response" json:"status,omitempty"`
	ErrorMsg *string  `protobuf:"bytes,2,opt,name=errorMsg,proto3,oneof" json:"errorMsg,omitempty"`
}

func (x *RequestResponse) Reset() {
	*x = RequestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_post_crud_pb_post_crud_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestResponse) ProtoMessage() {}

func (x *RequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_post_crud_pb_post_crud_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestResponse.ProtoReflect.Descriptor instead.
func (*RequestResponse) Descriptor() ([]byte, []int) {
	return file_pkg_post_crud_pb_post_crud_proto_rawDescGZIP(), []int{6}
}

func (x *RequestResponse) GetStatus() Response {
	if x != nil {
		return x.Status
	}
	return Response_statusOK
}

func (x *RequestResponse) GetErrorMsg() string {
	if x != nil && x.ErrorMsg != nil {
		return *x.ErrorMsg
	}
	return ""
}

type ReadPostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   Response `protobuf:"varint,1,opt,name=status,proto3,enum=postcrud.Response" json:"status,omitempty"`
	ErrorMsg *string  `protobuf:"bytes,2,opt,name=errorMsg,proto3,oneof" json:"errorMsg,omitempty"`
	Post     *Post    `protobuf:"bytes,3,opt,name=post,proto3" json:"post,omitempty"`
}

func (x *ReadPostResponse) Reset() {
	*x = ReadPostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_post_crud_pb_post_crud_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadPostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadPostResponse) ProtoMessage() {}

func (x *ReadPostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_post_crud_pb_post_crud_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadPostResponse.ProtoReflect.Descriptor instead.
func (*ReadPostResponse) Descriptor() ([]byte, []int) {
	return file_pkg_post_crud_pb_post_crud_proto_rawDescGZIP(), []int{7}
}

func (x *ReadPostResponse) GetStatus() Response {
	if x != nil {
		return x.Status
	}
	return Response_statusOK
}

func (x *ReadPostResponse) GetErrorMsg() string {
	if x != nil && x.ErrorMsg != nil {
		return *x.ErrorMsg
	}
	return ""
}

func (x *ReadPostResponse) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type ListPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID *int64 `protobuf:"varint,1,opt,name=userID,proto3,oneof" json:"userID,omitempty"`
	Page   *int64 `protobuf:"varint,2,opt,name=page,proto3,oneof" json:"page,omitempty"`
}

func (x *ListPostRequest) Reset() {
	*x = ListPostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_post_crud_pb_post_crud_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPostRequest) ProtoMessage() {}

func (x *ListPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_post_crud_pb_post_crud_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPostRequest.ProtoReflect.Descriptor instead.
func (*ListPostRequest) Descriptor() ([]byte, []int) {
	return file_pkg_post_crud_pb_post_crud_proto_rawDescGZIP(), []int{8}
}

func (x *ListPostRequest) GetUserID() int64 {
	if x != nil && x.UserID != nil {
		return *x.UserID
	}
	return 0
}

func (x *ListPostRequest) GetPage() int64 {
	if x != nil && x.Page != nil {
		return *x.Page
	}
	return 0
}

type ListPostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   Response `protobuf:"varint,1,opt,name=status,proto3,enum=postcrud.Response" json:"status,omitempty"`
	ErrorMsg *string  `protobuf:"bytes,2,opt,name=errorMsg,proto3,oneof" json:"errorMsg,omitempty"`
	Posts    []*Post  `protobuf:"bytes,3,rep,name=posts,proto3" json:"posts,omitempty"`
}

func (x *ListPostResponse) Reset() {
	*x = ListPostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_post_crud_pb_post_crud_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPostResponse) ProtoMessage() {}

func (x *ListPostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_post_crud_pb_post_crud_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPostResponse.ProtoReflect.Descriptor instead.
func (*ListPostResponse) Descriptor() ([]byte, []int) {
	return file_pkg_post_crud_pb_post_crud_proto_rawDescGZIP(), []int{9}
}

func (x *ListPostResponse) GetStatus() Response {
	if x != nil {
		return x.Status
	}
	return Response_statusOK
}

func (x *ListPostResponse) GetErrorMsg() string {
	if x != nil && x.ErrorMsg != nil {
		return *x.ErrorMsg
	}
	return ""
}

func (x *ListPostResponse) GetPosts() []*Post {
	if x != nil {
		return x.Posts
	}
	return nil
}

var File_pkg_post_crud_pb_post_crud_proto protoreflect.FileDescriptor

var file_pkg_post_crud_pb_post_crud_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2d, 0x63, 0x72, 0x75, 0x64, 0x2f,
	0x70, 0x62, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2d, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x63, 0x72, 0x75, 0x64, 0x22, 0x58, 0x0a, 0x04,
	0x50, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x37, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x70,
	0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x6f, 0x73, 0x74,
	0x63, 0x72, 0x75, 0x64, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x22,
	0x7e, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x63, 0x72, 0x75, 0x64,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1f, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x88,
	0x01, 0x01, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x49, 0x44, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x22,
	0x37, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x22, 0x21, 0x0a,
	0x0f, 0x52, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44,
	0x22, 0x6b, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1f, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x88, 0x01, 0x01,
	0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x22, 0x90, 0x01,
	0x0a, 0x10, 0x52, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f,
	0x0a, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x88, 0x01, 0x01, 0x12,
	0x22, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x70, 0x6f, 0x73, 0x74, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x70,
	0x6f, 0x73, 0x74, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67,
	0x22, 0x5b, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x88, 0x01, 0x01,
	0x12, 0x17, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x22, 0x92, 0x01,
	0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f,
	0x0a, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x88, 0x01, 0x01, 0x12,
	0x24, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x70, 0x6f, 0x73, 0x74, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x05,
	0x70, 0x6f, 0x73, 0x74, 0x73, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d,
	0x73, 0x67, 0x2a, 0x27, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0c,
	0x0a, 0x08, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x72, 0x72, 0x10, 0x01, 0x32, 0xe1, 0x02, 0x0a, 0x0f,
	0x50, 0x6f, 0x73, 0x74, 0x43, 0x52, 0x55, 0x44, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x45, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x70, 0x6f, 0x73, 0x74,
	0x63, 0x72, 0x75, 0x64, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x63, 0x72, 0x75,
	0x64, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x1b, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x70, 0x6f, 0x73, 0x74, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x63, 0x72, 0x75, 0x64, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f,
	0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x19, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x63, 0x72, 0x75,
	0x64, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x52, 0x65, 0x61,
	0x64, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3e, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x63, 0x72,
	0x75, 0x64, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x14, 0x5a, 0x12, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2d, 0x63, 0x72,
	0x75, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_post_crud_pb_post_crud_proto_rawDescOnce sync.Once
	file_pkg_post_crud_pb_post_crud_proto_rawDescData = file_pkg_post_crud_pb_post_crud_proto_rawDesc
)

func file_pkg_post_crud_pb_post_crud_proto_rawDescGZIP() []byte {
	file_pkg_post_crud_pb_post_crud_proto_rawDescOnce.Do(func() {
		file_pkg_post_crud_pb_post_crud_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_post_crud_pb_post_crud_proto_rawDescData)
	})
	return file_pkg_post_crud_pb_post_crud_proto_rawDescData
}

var file_pkg_post_crud_pb_post_crud_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_post_crud_pb_post_crud_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_pkg_post_crud_pb_post_crud_proto_goTypes = []interface{}{
	(Response)(0),              // 0: postcrud.Response
	(*Post)(nil),               // 1: postcrud.Post
	(*CreatePostRequest)(nil),  // 2: postcrud.CreatePostRequest
	(*CreatePostResponse)(nil), // 3: postcrud.CreatePostResponse
	(*UpdatePostRequest)(nil),  // 4: postcrud.UpdatePostRequest
	(*DeletePostRequest)(nil),  // 5: postcrud.DeletePostRequest
	(*ReadPostRequest)(nil),    // 6: postcrud.ReadPostRequest
	(*RequestResponse)(nil),    // 7: postcrud.RequestResponse
	(*ReadPostResponse)(nil),   // 8: postcrud.ReadPostResponse
	(*ListPostRequest)(nil),    // 9: postcrud.ListPostRequest
	(*ListPostResponse)(nil),   // 10: postcrud.ListPostResponse
}
var file_pkg_post_crud_pb_post_crud_proto_depIdxs = []int32{
	1,  // 0: postcrud.CreatePostRequest.post:type_name -> postcrud.Post
	0,  // 1: postcrud.CreatePostResponse.status:type_name -> postcrud.Response
	1,  // 2: postcrud.UpdatePostRequest.post:type_name -> postcrud.Post
	0,  // 3: postcrud.RequestResponse.status:type_name -> postcrud.Response
	0,  // 4: postcrud.ReadPostResponse.status:type_name -> postcrud.Response
	1,  // 5: postcrud.ReadPostResponse.post:type_name -> postcrud.Post
	0,  // 6: postcrud.ListPostResponse.status:type_name -> postcrud.Response
	1,  // 7: postcrud.ListPostResponse.posts:type_name -> postcrud.Post
	2,  // 8: postcrud.PostCRUDService.Create:input_type -> postcrud.CreatePostRequest
	4,  // 9: postcrud.PostCRUDService.Update:input_type -> postcrud.UpdatePostRequest
	5,  // 10: postcrud.PostCRUDService.Delete:input_type -> postcrud.DeletePostRequest
	6,  // 11: postcrud.PostCRUDService.Read:input_type -> postcrud.ReadPostRequest
	9,  // 12: postcrud.PostCRUDService.List:input_type -> postcrud.ListPostRequest
	3,  // 13: postcrud.PostCRUDService.Create:output_type -> postcrud.CreatePostResponse
	7,  // 14: postcrud.PostCRUDService.Update:output_type -> postcrud.RequestResponse
	7,  // 15: postcrud.PostCRUDService.Delete:output_type -> postcrud.RequestResponse
	8,  // 16: postcrud.PostCRUDService.Read:output_type -> postcrud.ReadPostResponse
	7,  // 17: postcrud.PostCRUDService.List:output_type -> postcrud.RequestResponse
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_pkg_post_crud_pb_post_crud_proto_init() }
func file_pkg_post_crud_pb_post_crud_proto_init() {
	if File_pkg_post_crud_pb_post_crud_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_post_crud_pb_post_crud_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Post); i {
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
		file_pkg_post_crud_pb_post_crud_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePostRequest); i {
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
		file_pkg_post_crud_pb_post_crud_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePostResponse); i {
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
		file_pkg_post_crud_pb_post_crud_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePostRequest); i {
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
		file_pkg_post_crud_pb_post_crud_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePostRequest); i {
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
		file_pkg_post_crud_pb_post_crud_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadPostRequest); i {
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
		file_pkg_post_crud_pb_post_crud_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestResponse); i {
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
		file_pkg_post_crud_pb_post_crud_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadPostResponse); i {
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
		file_pkg_post_crud_pb_post_crud_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPostRequest); i {
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
		file_pkg_post_crud_pb_post_crud_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPostResponse); i {
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
	file_pkg_post_crud_pb_post_crud_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_pkg_post_crud_pb_post_crud_proto_msgTypes[6].OneofWrappers = []interface{}{}
	file_pkg_post_crud_pb_post_crud_proto_msgTypes[7].OneofWrappers = []interface{}{}
	file_pkg_post_crud_pb_post_crud_proto_msgTypes[8].OneofWrappers = []interface{}{}
	file_pkg_post_crud_pb_post_crud_proto_msgTypes[9].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_post_crud_pb_post_crud_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_post_crud_pb_post_crud_proto_goTypes,
		DependencyIndexes: file_pkg_post_crud_pb_post_crud_proto_depIdxs,
		EnumInfos:         file_pkg_post_crud_pb_post_crud_proto_enumTypes,
		MessageInfos:      file_pkg_post_crud_pb_post_crud_proto_msgTypes,
	}.Build()
	File_pkg_post_crud_pb_post_crud_proto = out.File
	file_pkg_post_crud_pb_post_crud_proto_rawDesc = nil
	file_pkg_post_crud_pb_post_crud_proto_goTypes = nil
	file_pkg_post_crud_pb_post_crud_proto_depIdxs = nil
}
