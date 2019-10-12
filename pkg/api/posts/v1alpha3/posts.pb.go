// Copyright 2019 Publr Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/posts/v1alpha3/posts.proto

package posts

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Post struct {
	Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Title                string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Slug                 string               `protobuf:"bytes,3,opt,name=slug,proto3" json:"slug,omitempty"`
	Html                 string               `protobuf:"bytes,4,opt,name=html,proto3" json:"html,omitempty"`
	Image                string               `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
	Published            bool                 `protobuf:"varint,6,opt,name=published,proto3" json:"published,omitempty"`
	CreateTime           *timestamp.Timestamp `protobuf:"bytes,7,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime           *timestamp.Timestamp `protobuf:"bytes,8,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	PublishTime          *timestamp.Timestamp `protobuf:"bytes,9,opt,name=publish_time,json=publishTime,proto3" json:"publish_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Post) Reset()         { *m = Post{} }
func (m *Post) String() string { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()    {}
func (*Post) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bca93b1a45ce262, []int{0}
}

func (m *Post) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Post.Unmarshal(m, b)
}
func (m *Post) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Post.Marshal(b, m, deterministic)
}
func (m *Post) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Post.Merge(m, src)
}
func (m *Post) XXX_Size() int {
	return xxx_messageInfo_Post.Size(m)
}
func (m *Post) XXX_DiscardUnknown() {
	xxx_messageInfo_Post.DiscardUnknown(m)
}

var xxx_messageInfo_Post proto.InternalMessageInfo

func (m *Post) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Post) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Post) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *Post) GetHtml() string {
	if m != nil {
		return m.Html
	}
	return ""
}

func (m *Post) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Post) GetPublished() bool {
	if m != nil {
		return m.Published
	}
	return false
}

func (m *Post) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Post) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *Post) GetPublishTime() *timestamp.Timestamp {
	if m != nil {
		return m.PublishTime
	}
	return nil
}

type PostList struct {
	Posts                []*Post  `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostList) Reset()         { *m = PostList{} }
func (m *PostList) String() string { return proto.CompactTextString(m) }
func (*PostList) ProtoMessage()    {}
func (*PostList) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bca93b1a45ce262, []int{1}
}

func (m *PostList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostList.Unmarshal(m, b)
}
func (m *PostList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostList.Marshal(b, m, deterministic)
}
func (m *PostList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostList.Merge(m, src)
}
func (m *PostList) XXX_Size() int {
	return xxx_messageInfo_PostList.Size(m)
}
func (m *PostList) XXX_DiscardUnknown() {
	xxx_messageInfo_PostList.DiscardUnknown(m)
}

var xxx_messageInfo_PostList proto.InternalMessageInfo

func (m *PostList) GetPosts() []*Post {
	if m != nil {
		return m.Posts
	}
	return nil
}

func (m *PostList) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

type ListPostRequest struct {
	Parent               string   `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	Published            bool     `protobuf:"varint,2,opt,name=published,proto3" json:"published,omitempty"`
	OrderBy              string   `protobuf:"bytes,3,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	PageSize             int32    `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken            string   `protobuf:"bytes,5,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPostRequest) Reset()         { *m = ListPostRequest{} }
func (m *ListPostRequest) String() string { return proto.CompactTextString(m) }
func (*ListPostRequest) ProtoMessage()    {}
func (*ListPostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bca93b1a45ce262, []int{2}
}

func (m *ListPostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPostRequest.Unmarshal(m, b)
}
func (m *ListPostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPostRequest.Marshal(b, m, deterministic)
}
func (m *ListPostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPostRequest.Merge(m, src)
}
func (m *ListPostRequest) XXX_Size() int {
	return xxx_messageInfo_ListPostRequest.Size(m)
}
func (m *ListPostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListPostRequest proto.InternalMessageInfo

func (m *ListPostRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListPostRequest) GetPublished() bool {
	if m != nil {
		return m.Published
	}
	return false
}

func (m *ListPostRequest) GetOrderBy() string {
	if m != nil {
		return m.OrderBy
	}
	return ""
}

func (m *ListPostRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListPostRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

type CreatePostRequest struct {
	Parent               string   `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	Post                 *Post    `protobuf:"bytes,2,opt,name=post,proto3" json:"post,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePostRequest) Reset()         { *m = CreatePostRequest{} }
func (m *CreatePostRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePostRequest) ProtoMessage()    {}
func (*CreatePostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bca93b1a45ce262, []int{3}
}

func (m *CreatePostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePostRequest.Unmarshal(m, b)
}
func (m *CreatePostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePostRequest.Marshal(b, m, deterministic)
}
func (m *CreatePostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePostRequest.Merge(m, src)
}
func (m *CreatePostRequest) XXX_Size() int {
	return xxx_messageInfo_CreatePostRequest.Size(m)
}
func (m *CreatePostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePostRequest proto.InternalMessageInfo

func (m *CreatePostRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreatePostRequest) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

type GetPostRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPostRequest) Reset()         { *m = GetPostRequest{} }
func (m *GetPostRequest) String() string { return proto.CompactTextString(m) }
func (*GetPostRequest) ProtoMessage()    {}
func (*GetPostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bca93b1a45ce262, []int{4}
}

func (m *GetPostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPostRequest.Unmarshal(m, b)
}
func (m *GetPostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPostRequest.Marshal(b, m, deterministic)
}
func (m *GetPostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPostRequest.Merge(m, src)
}
func (m *GetPostRequest) XXX_Size() int {
	return xxx_messageInfo_GetPostRequest.Size(m)
}
func (m *GetPostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPostRequest proto.InternalMessageInfo

func (m *GetPostRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type UpdatePostRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Post                 *Post    `protobuf:"bytes,2,opt,name=post,proto3" json:"post,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePostRequest) Reset()         { *m = UpdatePostRequest{} }
func (m *UpdatePostRequest) String() string { return proto.CompactTextString(m) }
func (*UpdatePostRequest) ProtoMessage()    {}
func (*UpdatePostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bca93b1a45ce262, []int{5}
}

func (m *UpdatePostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePostRequest.Unmarshal(m, b)
}
func (m *UpdatePostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePostRequest.Marshal(b, m, deterministic)
}
func (m *UpdatePostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePostRequest.Merge(m, src)
}
func (m *UpdatePostRequest) XXX_Size() int {
	return xxx_messageInfo_UpdatePostRequest.Size(m)
}
func (m *UpdatePostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePostRequest proto.InternalMessageInfo

func (m *UpdatePostRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdatePostRequest) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

type DeletePostRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePostRequest) Reset()         { *m = DeletePostRequest{} }
func (m *DeletePostRequest) String() string { return proto.CompactTextString(m) }
func (*DeletePostRequest) ProtoMessage()    {}
func (*DeletePostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bca93b1a45ce262, []int{6}
}

func (m *DeletePostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePostRequest.Unmarshal(m, b)
}
func (m *DeletePostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePostRequest.Marshal(b, m, deterministic)
}
func (m *DeletePostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePostRequest.Merge(m, src)
}
func (m *DeletePostRequest) XXX_Size() int {
	return xxx_messageInfo_DeletePostRequest.Size(m)
}
func (m *DeletePostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePostRequest proto.InternalMessageInfo

func (m *DeletePostRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type SearchPostRequest struct {
	Parent               string   `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	Query                string   `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	Published            bool     `protobuf:"varint,3,opt,name=published,proto3" json:"published,omitempty"`
	OrderBy              string   `protobuf:"bytes,4,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	PageSize             int32    `protobuf:"varint,5,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken            string   `protobuf:"bytes,6,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchPostRequest) Reset()         { *m = SearchPostRequest{} }
func (m *SearchPostRequest) String() string { return proto.CompactTextString(m) }
func (*SearchPostRequest) ProtoMessage()    {}
func (*SearchPostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bca93b1a45ce262, []int{7}
}

func (m *SearchPostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchPostRequest.Unmarshal(m, b)
}
func (m *SearchPostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchPostRequest.Marshal(b, m, deterministic)
}
func (m *SearchPostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchPostRequest.Merge(m, src)
}
func (m *SearchPostRequest) XXX_Size() int {
	return xxx_messageInfo_SearchPostRequest.Size(m)
}
func (m *SearchPostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchPostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchPostRequest proto.InternalMessageInfo

func (m *SearchPostRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *SearchPostRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *SearchPostRequest) GetPublished() bool {
	if m != nil {
		return m.Published
	}
	return false
}

func (m *SearchPostRequest) GetOrderBy() string {
	if m != nil {
		return m.OrderBy
	}
	return ""
}

func (m *SearchPostRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *SearchPostRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func init() {
	proto.RegisterType((*Post)(nil), "publr.posts.v1alpha3.Post")
	proto.RegisterType((*PostList)(nil), "publr.posts.v1alpha3.PostList")
	proto.RegisterType((*ListPostRequest)(nil), "publr.posts.v1alpha3.ListPostRequest")
	proto.RegisterType((*CreatePostRequest)(nil), "publr.posts.v1alpha3.CreatePostRequest")
	proto.RegisterType((*GetPostRequest)(nil), "publr.posts.v1alpha3.GetPostRequest")
	proto.RegisterType((*UpdatePostRequest)(nil), "publr.posts.v1alpha3.UpdatePostRequest")
	proto.RegisterType((*DeletePostRequest)(nil), "publr.posts.v1alpha3.DeletePostRequest")
	proto.RegisterType((*SearchPostRequest)(nil), "publr.posts.v1alpha3.SearchPostRequest")
}

func init() { proto.RegisterFile("api/posts/v1alpha3/posts.proto", fileDescriptor_3bca93b1a45ce262) }

var fileDescriptor_3bca93b1a45ce262 = []byte{
	// 747 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0x95, 0xd3, 0xfc, 0xde, 0x7c, 0x1f, 0x55, 0x46, 0x55, 0x65, 0xd2, 0x52, 0x22, 0xab, 0xd0,
	0x2a, 0xaa, 0x6c, 0x9a, 0x08, 0x09, 0xb5, 0xea, 0xa6, 0x80, 0xd8, 0xb0, 0xa8, 0xd2, 0x02, 0x52,
	0x91, 0x88, 0x26, 0xc9, 0xc5, 0xb1, 0xea, 0xbf, 0xda, 0xe3, 0x8a, 0x14, 0x75, 0x01, 0x5b, 0x96,
	0xec, 0xd9, 0xb0, 0xe2, 0x05, 0xd8, 0x20, 0x96, 0x3c, 0x01, 0xaf, 0xc0, 0x83, 0xa0, 0x99, 0x71,
	0x62, 0xf2, 0x9f, 0xec, 0x7c, 0xaf, 0xcf, 0xb5, 0xcf, 0x9c, 0x73, 0xae, 0x13, 0xd8, 0xa2, 0xbe,
	0x65, 0xf8, 0x5e, 0xc8, 0x42, 0xe3, 0x6a, 0x9f, 0xda, 0x7e, 0x97, 0xd6, 0x65, 0xa9, 0xfb, 0x81,
	0xc7, 0x3c, 0xb2, 0xe6, 0x47, 0x2d, 0x3b, 0xd0, 0x65, 0xab, 0x8f, 0x28, 0x6f, 0x9a, 0x9e, 0x67,
	0xda, 0x68, 0xf0, 0x61, 0xea, 0xba, 0x1e, 0xa3, 0xcc, 0xf2, 0xdc, 0x78, 0xa6, 0xbc, 0x11, 0xdf,
	0x15, 0x55, 0x2b, 0x7a, 0x6b, 0xa0, 0xe3, 0xb3, 0x5e, 0x7c, 0xf3, 0xee, 0xe8, 0x4d, 0x66, 0x39,
	0x18, 0x32, 0xea, 0xf8, 0x12, 0xa0, 0xfd, 0x4a, 0x41, 0xfa, 0xc4, 0x0b, 0x19, 0x21, 0x90, 0x76,
	0xa9, 0x83, 0xaa, 0x52, 0x51, 0x76, 0x0b, 0x0d, 0x71, 0x4d, 0xd6, 0x20, 0xc3, 0x2c, 0x66, 0xa3,
	0x9a, 0x12, 0x4d, 0x59, 0x70, 0x64, 0x68, 0x47, 0xa6, 0xba, 0x22, 0x91, 0xfc, 0x9a, 0xf7, 0xba,
	0xcc, 0xb1, 0xd5, 0xb4, 0xec, 0xf1, 0x6b, 0x3e, 0x6d, 0x39, 0xd4, 0x44, 0x35, 0x23, 0xa7, 0x45,
	0x41, 0x36, 0xa1, 0xc0, 0x0f, 0x69, 0x85, 0x5d, 0xec, 0xa8, 0xd9, 0x8a, 0xb2, 0x9b, 0x6f, 0x24,
	0x0d, 0x72, 0x08, 0xc5, 0x76, 0x80, 0x94, 0x61, 0x93, 0x13, 0x55, 0x73, 0x15, 0x65, 0xb7, 0x58,
	0x2b, 0xeb, 0xf2, 0x14, 0x7a, 0xff, 0x14, 0xfa, 0x59, 0xff, 0x14, 0x0d, 0x90, 0x70, 0xde, 0xe0,
	0xc3, 0x91, 0xdf, 0x19, 0x0c, 0xe7, 0xe7, 0x0f, 0x4b, 0xb8, 0x18, 0x3e, 0x82, 0xff, 0x62, 0x1a,
	0x72, 0xba, 0x30, 0x77, 0xba, 0x18, 0xe3, 0x79, 0x47, 0xeb, 0x40, 0x9e, 0xcb, 0xf8, 0xdc, 0x0a,
	0x19, 0x79, 0x00, 0x19, 0xe1, 0xa0, 0xaa, 0x54, 0x56, 0xc4, 0x33, 0x26, 0xb9, 0xaa, 0x73, 0x78,
	0x43, 0x02, 0xc9, 0x7d, 0x58, 0x75, 0xf1, 0x1d, 0x6b, 0xfa, 0xd4, 0xc4, 0x26, 0xf3, 0x2e, 0xd0,
	0x8d, 0x25, 0xff, 0x9f, 0xb7, 0x4f, 0xa8, 0x89, 0x67, 0xbc, 0xa9, 0x7d, 0x51, 0x60, 0x95, 0xbf,
	0x42, 0xcc, 0xe2, 0x65, 0x84, 0x21, 0x23, 0xeb, 0x90, 0xf5, 0x69, 0x80, 0x2e, 0x8b, 0xad, 0x8b,
	0xab, 0x61, 0xa1, 0x53, 0xa3, 0x42, 0xdf, 0x86, 0xbc, 0x17, 0x74, 0x30, 0x68, 0xb6, 0x7a, 0xb1,
	0x91, 0x39, 0x51, 0x1f, 0xf7, 0xc8, 0x06, 0x14, 0x04, 0x8f, 0xd0, 0xba, 0x46, 0x61, 0x68, 0xa6,
	0x91, 0xe7, 0x8d, 0x53, 0xeb, 0x1a, 0xc9, 0x1d, 0x80, 0x7f, 0x48, 0x4a, 0x67, 0x05, 0x5c, 0x12,
	0x7c, 0x0d, 0xa5, 0xc7, 0xc2, 0x90, 0x45, 0x18, 0xea, 0x90, 0xe6, 0xc7, 0x17, 0xe4, 0x66, 0xcb,
	0x24, 0x70, 0xda, 0x36, 0xdc, 0x7a, 0x86, 0x43, 0x67, 0x9f, 0x10, 0x5a, 0xed, 0x15, 0x94, 0x5e,
	0x08, 0x5b, 0xe7, 0x00, 0x97, 0x7e, 0xfd, 0x0e, 0x94, 0x9e, 0xa0, 0x8d, 0x73, 0x1f, 0xac, 0x7d,
	0x57, 0xa0, 0x74, 0x8a, 0x34, 0x68, 0x77, 0x17, 0x51, 0x61, 0x0d, 0x32, 0x97, 0x11, 0x06, 0xbd,
	0xfe, 0x92, 0x89, 0x62, 0xd8, 0xbd, 0x95, 0x59, 0xee, 0xa5, 0x67, 0xb8, 0x97, 0x99, 0xe9, 0x5e,
	0x76, 0xc4, 0xbd, 0xda, 0xcf, 0x1c, 0x14, 0x39, 0xe5, 0x53, 0x0c, 0xae, 0xac, 0x36, 0x92, 0x6f,
	0x0a, 0xe4, 0xfb, 0x71, 0x23, 0xf7, 0x26, 0x0b, 0x34, 0x12, 0xc7, 0xf2, 0xd6, 0x74, 0x1d, 0x39,
	0x54, 0x7b, 0xf9, 0xf1, 0xf7, 0x9f, 0xcf, 0xa9, 0x13, 0x52, 0x49, 0xbe, 0x80, 0xef, 0xa5, 0x12,
	0x47, 0xa1, 0xc5, 0x30, 0x34, 0xaa, 0x37, 0xf2, 0x93, 0x78, 0xbe, 0x47, 0xaa, 0x53, 0x31, 0x06,
	0x8d, 0x58, 0xd7, 0x0b, 0x12, 0x34, 0xf9, 0xa4, 0x00, 0x24, 0xd1, 0x23, 0x3b, 0x93, 0x69, 0x8c,
	0x85, 0xb3, 0x3c, 0xc3, 0x77, 0xed, 0x91, 0xe0, 0x5a, 0xd3, 0x96, 0xe0, 0x71, 0x20, 0xb2, 0x42,
	0xbe, 0x2a, 0x90, 0x8b, 0xb3, 0x4a, 0xb6, 0x27, 0xbf, 0x61, 0x38, 0xca, 0x33, 0x79, 0x4c, 0xd2,
	0x8c, 0x27, 0x6d, 0xc0, 0x42, 0xfe, 0xa4, 0x54, 0x6f, 0x46, 0x34, 0x1b, 0xc2, 0x0c, 0x98, 0x0e,
	0xd0, 0x42, 0xb3, 0x64, 0x57, 0xa6, 0x69, 0x36, 0xb6, 0x4d, 0x8b, 0x68, 0x56, 0x5b, 0x82, 0x47,
	0xac, 0xd9, 0x07, 0x05, 0x20, 0x59, 0xb0, 0x69, 0x6c, 0xc6, 0x56, 0xb0, 0xbc, 0x3e, 0xf6, 0x8d,
	0x7e, 0xca, 0x7f, 0x01, 0xb5, 0x9a, 0x60, 0xb2, 0x57, 0x5d, 0x46, 0x91, 0x1f, 0x0a, 0x40, 0xb2,
	0xba, 0xd3, 0x38, 0x8c, 0x2d, 0xf7, 0xdc, 0xd4, 0x77, 0x04, 0x97, 0x37, 0x64, 0x67, 0x5e, 0xea,
	0x0f, 0x42, 0xf1, 0xec, 0xf3, 0x3a, 0xd9, 0x5f, 0x22, 0x74, 0x72, 0xe8, 0xf8, 0xe1, 0x79, 0xdd,
	0xb4, 0x58, 0x37, 0x6a, 0xe9, 0x6d, 0xcf, 0x31, 0xfc, 0xe0, 0x22, 0x8c, 0x0c, 0xc1, 0xcb, 0xf0,
	0x2f, 0x4c, 0x63, 0xfc, 0xbf, 0xc7, 0xa1, 0x28, 0x5b, 0x59, 0x21, 0x5c, 0xfd, 0x6f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xcb, 0x7f, 0x09, 0x1b, 0x9e, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PostServiceClient is the client API for PostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PostServiceClient interface {
	ListPost(ctx context.Context, in *ListPostRequest, opts ...grpc.CallOption) (*PostList, error)
	CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*Post, error)
	GetPost(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*Post, error)
	UpdatePost(ctx context.Context, in *UpdatePostRequest, opts ...grpc.CallOption) (*Post, error)
	DeletePost(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	SearchPost(ctx context.Context, in *SearchPostRequest, opts ...grpc.CallOption) (*PostList, error)
}

type postServiceClient struct {
	cc *grpc.ClientConn
}

func NewPostServiceClient(cc *grpc.ClientConn) PostServiceClient {
	return &postServiceClient{cc}
}

func (c *postServiceClient) ListPost(ctx context.Context, in *ListPostRequest, opts ...grpc.CallOption) (*PostList, error) {
	out := new(PostList)
	err := c.cc.Invoke(ctx, "/publr.posts.v1alpha3.PostService/ListPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := c.cc.Invoke(ctx, "/publr.posts.v1alpha3.PostService/CreatePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetPost(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := c.cc.Invoke(ctx, "/publr.posts.v1alpha3.PostService/GetPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) UpdatePost(ctx context.Context, in *UpdatePostRequest, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := c.cc.Invoke(ctx, "/publr.posts.v1alpha3.PostService/UpdatePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) DeletePost(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/publr.posts.v1alpha3.PostService/DeletePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) SearchPost(ctx context.Context, in *SearchPostRequest, opts ...grpc.CallOption) (*PostList, error) {
	out := new(PostList)
	err := c.cc.Invoke(ctx, "/publr.posts.v1alpha3.PostService/SearchPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostServiceServer is the server API for PostService service.
type PostServiceServer interface {
	ListPost(context.Context, *ListPostRequest) (*PostList, error)
	CreatePost(context.Context, *CreatePostRequest) (*Post, error)
	GetPost(context.Context, *GetPostRequest) (*Post, error)
	UpdatePost(context.Context, *UpdatePostRequest) (*Post, error)
	DeletePost(context.Context, *DeletePostRequest) (*empty.Empty, error)
	SearchPost(context.Context, *SearchPostRequest) (*PostList, error)
}

// UnimplementedPostServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPostServiceServer struct {
}

func (*UnimplementedPostServiceServer) ListPost(ctx context.Context, req *ListPostRequest) (*PostList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPost not implemented")
}
func (*UnimplementedPostServiceServer) CreatePost(ctx context.Context, req *CreatePostRequest) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (*UnimplementedPostServiceServer) GetPost(ctx context.Context, req *GetPostRequest) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPost not implemented")
}
func (*UnimplementedPostServiceServer) UpdatePost(ctx context.Context, req *UpdatePostRequest) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePost not implemented")
}
func (*UnimplementedPostServiceServer) DeletePost(ctx context.Context, req *DeletePostRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}
func (*UnimplementedPostServiceServer) SearchPost(ctx context.Context, req *SearchPostRequest) (*PostList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchPost not implemented")
}

func RegisterPostServiceServer(s *grpc.Server, srv PostServiceServer) {
	s.RegisterService(&_PostService_serviceDesc, srv)
}

func _PostService_ListPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).ListPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/publr.posts.v1alpha3.PostService/ListPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).ListPost(ctx, req.(*ListPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/publr.posts.v1alpha3.PostService/CreatePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).CreatePost(ctx, req.(*CreatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/publr.posts.v1alpha3.PostService/GetPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetPost(ctx, req.(*GetPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_UpdatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).UpdatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/publr.posts.v1alpha3.PostService/UpdatePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).UpdatePost(ctx, req.(*UpdatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_DeletePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).DeletePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/publr.posts.v1alpha3.PostService/DeletePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).DeletePost(ctx, req.(*DeletePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_SearchPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).SearchPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/publr.posts.v1alpha3.PostService/SearchPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).SearchPost(ctx, req.(*SearchPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PostService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "publr.posts.v1alpha3.PostService",
	HandlerType: (*PostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListPost",
			Handler:    _PostService_ListPost_Handler,
		},
		{
			MethodName: "CreatePost",
			Handler:    _PostService_CreatePost_Handler,
		},
		{
			MethodName: "GetPost",
			Handler:    _PostService_GetPost_Handler,
		},
		{
			MethodName: "UpdatePost",
			Handler:    _PostService_UpdatePost_Handler,
		},
		{
			MethodName: "DeletePost",
			Handler:    _PostService_DeletePost_Handler,
		},
		{
			MethodName: "SearchPost",
			Handler:    _PostService_SearchPost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/posts/v1alpha3/posts.proto",
}