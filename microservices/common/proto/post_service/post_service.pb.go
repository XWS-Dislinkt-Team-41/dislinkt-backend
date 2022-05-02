// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: post_service/post_service.proto

package post

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_post_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_post_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_post_service_post_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post *Post `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_post_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_post_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_post_service_post_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetResponse) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type GetAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllRequest) Reset() {
	*x = GetAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_post_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRequest) ProtoMessage() {}

func (x *GetAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_post_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllRequest.ProtoReflect.Descriptor instead.
func (*GetAllRequest) Descriptor() ([]byte, []int) {
	return file_post_service_post_service_proto_rawDescGZIP(), []int{2}
}

type GetAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts []*Post `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
}

func (x *GetAllResponse) Reset() {
	*x = GetAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_post_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllResponse) ProtoMessage() {}

func (x *GetAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_post_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllResponse.ProtoReflect.Descriptor instead.
func (*GetAllResponse) Descriptor() ([]byte, []int) {
	return file_post_service_post_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllResponse) GetPosts() []*Post {
	if x != nil {
		return x.Posts
	}
	return nil
}

type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Text     string     `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Link     string     `protobuf:"bytes,3,opt,name=link,proto3" json:"link,omitempty"`
	Image    string     `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	OwnerId  string     `protobuf:"bytes,5,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	Likes    int64      `protobuf:"varint,6,opt,name=likes,proto3" json:"likes,omitempty"`
	Dislikes int64      `protobuf:"varint,7,opt,name=dislikes,proto3" json:"dislikes,omitempty"`
	Comments []*Comment `protobuf:"bytes,8,rep,name=comments,proto3" json:"comments,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_post_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_post_service_proto_msgTypes[4]
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
	return file_post_service_post_service_proto_rawDescGZIP(), []int{4}
}

func (x *Post) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Post) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Post) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *Post) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Post) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *Post) GetLikes() int64 {
	if x != nil {
		return x.Likes
	}
	return 0
}

func (x *Post) GetDislikes() int64 {
	if x != nil {
		return x.Dislikes
	}
	return 0
}

func (x *Post) GetComments() []*Comment {
	if x != nil {
		return x.Comments
	}
	return nil
}

type GetPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PostId string `protobuf:"bytes,2,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
}

func (x *GetPostRequest) Reset() {
	*x = GetPostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_post_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostRequest) ProtoMessage() {}

func (x *GetPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_post_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostRequest.ProtoReflect.Descriptor instead.
func (*GetPostRequest) Descriptor() ([]byte, []int) {
	return file_post_service_post_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetPostRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetPostRequest) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

type GetPostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post *Post `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
}

func (x *GetPostResponse) Reset() {
	*x = GetPostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_post_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostResponse) ProtoMessage() {}

func (x *GetPostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_post_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostResponse.ProtoReflect.Descriptor instead.
func (*GetPostResponse) Descriptor() ([]byte, []int) {
	return file_post_service_post_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetPostResponse) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type NewPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Post *Post  `protobuf:"bytes,2,opt,name=post,proto3" json:"post,omitempty"`
}

func (x *NewPostRequest) Reset() {
	*x = NewPostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_post_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewPostRequest) ProtoMessage() {}

func (x *NewPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_post_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewPostRequest.ProtoReflect.Descriptor instead.
func (*NewPostRequest) Descriptor() ([]byte, []int) {
	return file_post_service_post_service_proto_rawDescGZIP(), []int{7}
}

func (x *NewPostRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NewPostRequest) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type NewPostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post *Post `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
}

func (x *NewPostResponse) Reset() {
	*x = NewPostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_post_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewPostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewPostResponse) ProtoMessage() {}

func (x *NewPostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_post_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewPostResponse.ProtoReflect.Descriptor instead.
func (*NewPostResponse) Descriptor() ([]byte, []int) {
	return file_post_service_post_service_proto_rawDescGZIP(), []int{8}
}

func (x *NewPostResponse) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Text string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_post_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_post_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_post_service_post_service_proto_rawDescGZIP(), []int{9}
}

func (x *Comment) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Comment) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type CommentOnPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PostId  string   `protobuf:"bytes,2,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	Comment *Comment `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *CommentOnPostRequest) Reset() {
	*x = CommentOnPostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_post_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentOnPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentOnPostRequest) ProtoMessage() {}

func (x *CommentOnPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_post_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentOnPostRequest.ProtoReflect.Descriptor instead.
func (*CommentOnPostRequest) Descriptor() ([]byte, []int) {
	return file_post_service_post_service_proto_rawDescGZIP(), []int{10}
}

func (x *CommentOnPostRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CommentOnPostRequest) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

func (x *CommentOnPostRequest) GetComment() *Comment {
	if x != nil {
		return x.Comment
	}
	return nil
}

type CommentOnPostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comment *Comment `protobuf:"bytes,1,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *CommentOnPostResponse) Reset() {
	*x = CommentOnPostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_post_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentOnPostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentOnPostResponse) ProtoMessage() {}

func (x *CommentOnPostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_post_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentOnPostResponse.ProtoReflect.Descriptor instead.
func (*CommentOnPostResponse) Descriptor() ([]byte, []int) {
	return file_post_service_post_service_proto_rawDescGZIP(), []int{11}
}

func (x *CommentOnPostResponse) GetComment() *Comment {
	if x != nil {
		return x.Comment
	}
	return nil
}

var File_post_service_post_service_proto protoreflect.FileDescriptor

var file_post_service_post_service_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70,
	0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x2d, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x70, 0x6f,
	0x73, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x32, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x22, 0xcc, 0x01, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6b,
	0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x64, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x29, 0x0a, 0x08, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x70, 0x6f, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x39, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49,
	0x64, 0x22, 0x31, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x04,
	0x70, 0x6f, 0x73, 0x74, 0x22, 0x40, 0x0a, 0x0e, 0x4e, 0x65, 0x77, 0x50, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x22, 0x31, 0x0a, 0x0f, 0x4e, 0x65, 0x77, 0x50, 0x6f, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x70, 0x6f, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x50,
	0x6f, 0x73, 0x74, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x22, 0x31, 0x0a, 0x07, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x68, 0x0a, 0x14,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x6e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x27, 0x0a,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x40, 0x0a, 0x15, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x4f, 0x6e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x27, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x32, 0xd1, 0x03, 0x0a, 0x0b, 0x50, 0x6f, 0x73,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x14, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b,
	0x12, 0x19, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x70, 0x6f, 0x73,
	0x74, 0x2f, 0x7b, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x42, 0x0a, 0x06, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x13, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x6f, 0x73,
	0x74, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x0d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x07, 0x12, 0x05, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x12,
	0x57, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x6f, 0x73, 0x74,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x12, 0x54, 0x0a, 0x06, 0x49, 0x6e, 0x73, 0x65,
	0x72, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x4e, 0x65, 0x77, 0x50, 0x6f, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e,
	0x4e, 0x65, 0x77, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x0f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x3a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x12, 0x7c,
	0x0a, 0x0d, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x1a, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x6e,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x6e, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c,
	0x22, 0x21, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x70, 0x6f, 0x73,
	0x74, 0x2f, 0x7b, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x3a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x4c, 0x5a, 0x4a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x58, 0x57, 0x53, 0x2d, 0x44,
	0x69, 0x73, 0x6c, 0x69, 0x6e, 0x6b, 0x74, 0x2d, 0x54, 0x65, 0x61, 0x6d, 0x2d, 0x34, 0x31, 0x2f,
	0x64, 0x69, 0x73, 0x6c, 0x69, 0x6e, 0x6b, 0x74, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_post_service_post_service_proto_rawDescOnce sync.Once
	file_post_service_post_service_proto_rawDescData = file_post_service_post_service_proto_rawDesc
)

func file_post_service_post_service_proto_rawDescGZIP() []byte {
	file_post_service_post_service_proto_rawDescOnce.Do(func() {
		file_post_service_post_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_post_service_post_service_proto_rawDescData)
	})
	return file_post_service_post_service_proto_rawDescData
}

var file_post_service_post_service_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_post_service_post_service_proto_goTypes = []interface{}{
	(*GetRequest)(nil),            // 0: post.GetRequest
	(*GetResponse)(nil),           // 1: post.GetResponse
	(*GetAllRequest)(nil),         // 2: post.GetAllRequest
	(*GetAllResponse)(nil),        // 3: post.GetAllResponse
	(*Post)(nil),                  // 4: post.Post
	(*GetPostRequest)(nil),        // 5: post.GetPostRequest
	(*GetPostResponse)(nil),       // 6: post.GetPostResponse
	(*NewPostRequest)(nil),        // 7: post.NewPostRequest
	(*NewPostResponse)(nil),       // 8: post.NewPostResponse
	(*Comment)(nil),               // 9: post.Comment
	(*CommentOnPostRequest)(nil),  // 10: post.CommentOnPostRequest
	(*CommentOnPostResponse)(nil), // 11: post.CommentOnPostResponse
}
var file_post_service_post_service_proto_depIdxs = []int32{
	4,  // 0: post.GetResponse.post:type_name -> post.Post
	4,  // 1: post.GetAllResponse.posts:type_name -> post.Post
	9,  // 2: post.Post.comments:type_name -> post.Comment
	4,  // 3: post.GetPostResponse.post:type_name -> post.Post
	4,  // 4: post.NewPostRequest.post:type_name -> post.Post
	4,  // 5: post.NewPostResponse.post:type_name -> post.Post
	9,  // 6: post.CommentOnPostRequest.comment:type_name -> post.Comment
	9,  // 7: post.CommentOnPostResponse.comment:type_name -> post.Comment
	5,  // 8: post.PostService.Get:input_type -> post.GetPostRequest
	2,  // 9: post.PostService.GetAll:input_type -> post.GetAllRequest
	0,  // 10: post.PostService.GetAllFromCollection:input_type -> post.GetRequest
	7,  // 11: post.PostService.Insert:input_type -> post.NewPostRequest
	10, // 12: post.PostService.InsertComment:input_type -> post.CommentOnPostRequest
	1,  // 13: post.PostService.Get:output_type -> post.GetResponse
	3,  // 14: post.PostService.GetAll:output_type -> post.GetAllResponse
	3,  // 15: post.PostService.GetAllFromCollection:output_type -> post.GetAllResponse
	8,  // 16: post.PostService.Insert:output_type -> post.NewPostResponse
	11, // 17: post.PostService.InsertComment:output_type -> post.CommentOnPostResponse
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_post_service_post_service_proto_init() }
func file_post_service_post_service_proto_init() {
	if File_post_service_post_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_post_service_post_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_post_service_post_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_post_service_post_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllRequest); i {
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
		file_post_service_post_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllResponse); i {
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
		file_post_service_post_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_post_service_post_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostRequest); i {
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
		file_post_service_post_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostResponse); i {
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
		file_post_service_post_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewPostRequest); i {
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
		file_post_service_post_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewPostResponse); i {
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
		file_post_service_post_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
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
		file_post_service_post_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentOnPostRequest); i {
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
		file_post_service_post_service_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentOnPostResponse); i {
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
			RawDescriptor: file_post_service_post_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_post_service_post_service_proto_goTypes,
		DependencyIndexes: file_post_service_post_service_proto_depIdxs,
		MessageInfos:      file_post_service_post_service_proto_msgTypes,
	}.Build()
	File_post_service_post_service_proto = out.File
	file_post_service_post_service_proto_rawDesc = nil
	file_post_service_post_service_proto_goTypes = nil
	file_post_service_post_service_proto_depIdxs = nil
}