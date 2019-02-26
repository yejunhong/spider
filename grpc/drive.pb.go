// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/drive.proto

package grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	ConfigName           string   `protobuf:"bytes,2,opt,name=config_name,json=configName,proto3" json:"config_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_drive_3a6906e4ce667da7, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Request) GetConfigName() string {
	if m != nil {
		return m.ConfigName
	}
	return ""
}

type List struct {
	Tags                 string   `protobuf:"bytes,1,opt,name=tags,proto3" json:"tags,omitempty"`
	Detail               string   `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
	ResourceName         string   `protobuf:"bytes,3,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	ResourceUrl          string   `protobuf:"bytes,4,opt,name=resource_url,json=resourceUrl,proto3" json:"resource_url,omitempty"`
	ResourceImgUrl       string   `protobuf:"bytes,5,opt,name=resource_img_url,json=resourceImgUrl,proto3" json:"resource_img_url,omitempty"`
	Author               string   `protobuf:"bytes,6,opt,name=author,proto3" json:"author,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *List) Reset()         { *m = List{} }
func (m *List) String() string { return proto.CompactTextString(m) }
func (*List) ProtoMessage()    {}
func (*List) Descriptor() ([]byte, []int) {
	return fileDescriptor_drive_3a6906e4ce667da7, []int{1}
}
func (m *List) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_List.Unmarshal(m, b)
}
func (m *List) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_List.Marshal(b, m, deterministic)
}
func (dst *List) XXX_Merge(src proto.Message) {
	xxx_messageInfo_List.Merge(dst, src)
}
func (m *List) XXX_Size() int {
	return xxx_messageInfo_List.Size(m)
}
func (m *List) XXX_DiscardUnknown() {
	xxx_messageInfo_List.DiscardUnknown(m)
}

var xxx_messageInfo_List proto.InternalMessageInfo

func (m *List) GetTags() string {
	if m != nil {
		return m.Tags
	}
	return ""
}

func (m *List) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

func (m *List) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *List) GetResourceUrl() string {
	if m != nil {
		return m.ResourceUrl
	}
	return ""
}

func (m *List) GetResourceImgUrl() string {
	if m != nil {
		return m.ResourceImgUrl
	}
	return ""
}

func (m *List) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

type ListReply struct {
	Data                 []*List  `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Next                 string   `protobuf:"bytes,2,opt,name=next,proto3" json:"next,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListReply) Reset()         { *m = ListReply{} }
func (m *ListReply) String() string { return proto.CompactTextString(m) }
func (*ListReply) ProtoMessage()    {}
func (*ListReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_drive_3a6906e4ce667da7, []int{2}
}
func (m *ListReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListReply.Unmarshal(m, b)
}
func (m *ListReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListReply.Marshal(b, m, deterministic)
}
func (dst *ListReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListReply.Merge(dst, src)
}
func (m *ListReply) XXX_Size() int {
	return xxx_messageInfo_ListReply.Size(m)
}
func (m *ListReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ListReply.DiscardUnknown(m)
}

var xxx_messageInfo_ListReply proto.InternalMessageInfo

func (m *ListReply) GetData() []*List {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ListReply) GetNext() string {
	if m != nil {
		return m.Next
	}
	return ""
}

type Chapter struct {
	IsFree               string   `protobuf:"bytes,1,opt,name=is_free,json=isFree,proto3" json:"is_free,omitempty"`
	ResourceName         string   `protobuf:"bytes,2,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	ResourceUrl          string   `protobuf:"bytes,3,opt,name=resource_url,json=resourceUrl,proto3" json:"resource_url,omitempty"`
	ResourceImgUrl       string   `protobuf:"bytes,4,opt,name=resource_img_url,json=resourceImgUrl,proto3" json:"resource_img_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Chapter) Reset()         { *m = Chapter{} }
func (m *Chapter) String() string { return proto.CompactTextString(m) }
func (*Chapter) ProtoMessage()    {}
func (*Chapter) Descriptor() ([]byte, []int) {
	return fileDescriptor_drive_3a6906e4ce667da7, []int{3}
}
func (m *Chapter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chapter.Unmarshal(m, b)
}
func (m *Chapter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chapter.Marshal(b, m, deterministic)
}
func (dst *Chapter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chapter.Merge(dst, src)
}
func (m *Chapter) XXX_Size() int {
	return xxx_messageInfo_Chapter.Size(m)
}
func (m *Chapter) XXX_DiscardUnknown() {
	xxx_messageInfo_Chapter.DiscardUnknown(m)
}

var xxx_messageInfo_Chapter proto.InternalMessageInfo

func (m *Chapter) GetIsFree() string {
	if m != nil {
		return m.IsFree
	}
	return ""
}

func (m *Chapter) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *Chapter) GetResourceUrl() string {
	if m != nil {
		return m.ResourceUrl
	}
	return ""
}

func (m *Chapter) GetResourceImgUrl() string {
	if m != nil {
		return m.ResourceImgUrl
	}
	return ""
}

type ChapterReply struct {
	Data                 []*Chapter `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Next                 string     `protobuf:"bytes,2,opt,name=next,proto3" json:"next,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ChapterReply) Reset()         { *m = ChapterReply{} }
func (m *ChapterReply) String() string { return proto.CompactTextString(m) }
func (*ChapterReply) ProtoMessage()    {}
func (*ChapterReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_drive_3a6906e4ce667da7, []int{4}
}
func (m *ChapterReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChapterReply.Unmarshal(m, b)
}
func (m *ChapterReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChapterReply.Marshal(b, m, deterministic)
}
func (dst *ChapterReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChapterReply.Merge(dst, src)
}
func (m *ChapterReply) XXX_Size() int {
	return xxx_messageInfo_ChapterReply.Size(m)
}
func (m *ChapterReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ChapterReply.DiscardUnknown(m)
}

var xxx_messageInfo_ChapterReply proto.InternalMessageInfo

func (m *ChapterReply) GetData() []*Chapter {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ChapterReply) GetNext() string {
	if m != nil {
		return m.Next
	}
	return ""
}

type ChapterContent struct {
	ResourceImgUrl       []string `protobuf:"bytes,1,rep,name=resource_img_url,json=resourceImgUrl,proto3" json:"resource_img_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChapterContent) Reset()         { *m = ChapterContent{} }
func (m *ChapterContent) String() string { return proto.CompactTextString(m) }
func (*ChapterContent) ProtoMessage()    {}
func (*ChapterContent) Descriptor() ([]byte, []int) {
	return fileDescriptor_drive_3a6906e4ce667da7, []int{5}
}
func (m *ChapterContent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChapterContent.Unmarshal(m, b)
}
func (m *ChapterContent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChapterContent.Marshal(b, m, deterministic)
}
func (dst *ChapterContent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChapterContent.Merge(dst, src)
}
func (m *ChapterContent) XXX_Size() int {
	return xxx_messageInfo_ChapterContent.Size(m)
}
func (m *ChapterContent) XXX_DiscardUnknown() {
	xxx_messageInfo_ChapterContent.DiscardUnknown(m)
}

var xxx_messageInfo_ChapterContent proto.InternalMessageInfo

func (m *ChapterContent) GetResourceImgUrl() []string {
	if m != nil {
		return m.ResourceImgUrl
	}
	return nil
}

type ChapterContentReply struct {
	Data                 []*ChapterContent `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Next                 string            `protobuf:"bytes,2,opt,name=next,proto3" json:"next,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ChapterContentReply) Reset()         { *m = ChapterContentReply{} }
func (m *ChapterContentReply) String() string { return proto.CompactTextString(m) }
func (*ChapterContentReply) ProtoMessage()    {}
func (*ChapterContentReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_drive_3a6906e4ce667da7, []int{6}
}
func (m *ChapterContentReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChapterContentReply.Unmarshal(m, b)
}
func (m *ChapterContentReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChapterContentReply.Marshal(b, m, deterministic)
}
func (dst *ChapterContentReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChapterContentReply.Merge(dst, src)
}
func (m *ChapterContentReply) XXX_Size() int {
	return xxx_messageInfo_ChapterContentReply.Size(m)
}
func (m *ChapterContentReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ChapterContentReply.DiscardUnknown(m)
}

var xxx_messageInfo_ChapterContentReply proto.InternalMessageInfo

func (m *ChapterContentReply) GetData() []*ChapterContent {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ChapterContentReply) GetNext() string {
	if m != nil {
		return m.Next
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "grpc.Request")
	proto.RegisterType((*List)(nil), "grpc.List")
	proto.RegisterType((*ListReply)(nil), "grpc.ListReply")
	proto.RegisterType((*Chapter)(nil), "grpc.Chapter")
	proto.RegisterType((*ChapterReply)(nil), "grpc.ChapterReply")
	proto.RegisterType((*ChapterContent)(nil), "grpc.ChapterContent")
	proto.RegisterType((*ChapterContentReply)(nil), "grpc.ChapterContentReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BrowserClient is the client API for Browser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BrowserClient interface {
	CrawlList(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ListReply, error)
	CrawlChapter(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ChapterReply, error)
	CrawlChapterContent(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ChapterContentReply, error)
}

type browserClient struct {
	cc *grpc.ClientConn
}

func NewBrowserClient(cc *grpc.ClientConn) BrowserClient {
	return &browserClient{cc}
}

func (c *browserClient) CrawlList(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ListReply, error) {
	out := new(ListReply)
	err := c.cc.Invoke(ctx, "/grpc.browser/CrawlList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *browserClient) CrawlChapter(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ChapterReply, error) {
	out := new(ChapterReply)
	err := c.cc.Invoke(ctx, "/grpc.browser/CrawlChapter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *browserClient) CrawlChapterContent(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ChapterContentReply, error) {
	out := new(ChapterContentReply)
	err := c.cc.Invoke(ctx, "/grpc.browser/CrawlChapterContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BrowserServer is the server API for Browser service.
type BrowserServer interface {
	CrawlList(context.Context, *Request) (*ListReply, error)
	CrawlChapter(context.Context, *Request) (*ChapterReply, error)
	CrawlChapterContent(context.Context, *Request) (*ChapterContentReply, error)
}

func RegisterBrowserServer(s *grpc.Server, srv BrowserServer) {
	s.RegisterService(&_Browser_serviceDesc, srv)
}

func _Browser_CrawlList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrowserServer).CrawlList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.browser/CrawlList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrowserServer).CrawlList(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Browser_CrawlChapter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrowserServer).CrawlChapter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.browser/CrawlChapter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrowserServer).CrawlChapter(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Browser_CrawlChapterContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrowserServer).CrawlChapterContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.browser/CrawlChapterContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrowserServer).CrawlChapterContent(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Browser_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.browser",
	HandlerType: (*BrowserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CrawlList",
			Handler:    _Browser_CrawlList_Handler,
		},
		{
			MethodName: "CrawlChapter",
			Handler:    _Browser_CrawlChapter_Handler,
		},
		{
			MethodName: "CrawlChapterContent",
			Handler:    _Browser_CrawlChapterContent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/drive.proto",
}

func init() { proto.RegisterFile("grpc/drive.proto", fileDescriptor_drive_3a6906e4ce667da7) }

var fileDescriptor_drive_3a6906e4ce667da7 = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcb, 0x6e, 0xda, 0x40,
	0x14, 0x86, 0x6b, 0xec, 0xda, 0xe2, 0x70, 0x29, 0x1a, 0x2a, 0xea, 0x76, 0xd1, 0x82, 0xbb, 0xf1,
	0xa6, 0x54, 0x82, 0x5d, 0x55, 0x29, 0x8a, 0x50, 0x22, 0x45, 0x8a, 0xb2, 0x20, 0xca, 0x1a, 0x0d,
	0x70, 0x30, 0x96, 0x7c, 0xcb, 0x78, 0x1c, 0x92, 0xf7, 0xc8, 0x83, 0xe4, 0x01, 0xf2, 0x70, 0xd1,
	0x5c, 0x8c, 0x30, 0x17, 0x89, 0xdd, 0xcc, 0x99, 0xff, 0x3f, 0xfe, 0xbf, 0x33, 0x63, 0xe8, 0x04,
	0x2c, 0x5b, 0xfc, 0x5d, 0xb2, 0xf0, 0x09, 0x87, 0x19, 0x4b, 0x79, 0x4a, 0x2c, 0x51, 0xf1, 0xfe,
	0x83, 0x33, 0xc5, 0xc7, 0x02, 0x73, 0x4e, 0x3a, 0x60, 0x16, 0x2c, 0x72, 0x8d, 0xbe, 0xe1, 0xd7,
	0xa7, 0x62, 0x49, 0x7e, 0x41, 0x63, 0x91, 0x26, 0xab, 0x30, 0x98, 0x25, 0x34, 0x46, 0xb7, 0x26,
	0x4f, 0x40, 0x95, 0xee, 0x68, 0x8c, 0xde, 0xbb, 0x01, 0xd6, 0x6d, 0x98, 0x73, 0x42, 0xc0, 0xe2,
	0x34, 0xc8, 0xb5, 0x59, 0xae, 0x49, 0x0f, 0xec, 0x25, 0x72, 0x1a, 0x46, 0xda, 0xa8, 0x77, 0xe4,
	0x37, 0xb4, 0x18, 0xe6, 0x69, 0xc1, 0x16, 0xa8, 0xfa, 0x9a, 0xf2, 0xb8, 0x59, 0x16, 0x45, 0x67,
	0x32, 0x80, 0xed, 0x7e, 0x26, 0x52, 0x59, 0x52, 0xd3, 0x28, 0x6b, 0x0f, 0x2c, 0x22, 0x3e, 0x74,
	0xb6, 0x92, 0x30, 0x0e, 0xa4, 0xec, 0xb3, 0x94, 0xb5, 0xcb, 0xfa, 0x4d, 0x1c, 0x08, 0x65, 0x0f,
	0x6c, 0x5a, 0xf0, 0x75, 0xca, 0x5c, 0x5b, 0x25, 0x51, 0x3b, 0xef, 0x02, 0xea, 0x22, 0xfd, 0x14,
	0xb3, 0xe8, 0x85, 0xfc, 0x04, 0x6b, 0x49, 0x39, 0x75, 0x8d, 0xbe, 0xe9, 0x37, 0x46, 0x30, 0x14,
	0xe3, 0x19, 0xca, 0x63, 0x59, 0x17, 0x88, 0x09, 0x3e, 0x73, 0x0d, 0x23, 0xd7, 0xde, 0xab, 0x01,
	0xce, 0x64, 0x4d, 0x33, 0x8e, 0x8c, 0x7c, 0x03, 0x27, 0xcc, 0x67, 0x2b, 0x86, 0xa8, 0xa7, 0x60,
	0x87, 0xf9, 0x35, 0x43, 0x3c, 0xe4, 0xad, 0x9d, 0xc1, 0x6b, 0x9e, 0xc7, 0x6b, 0x1d, 0xe3, 0xf5,
	0xae, 0xa0, 0xa9, 0x53, 0x29, 0xb4, 0x41, 0x05, 0xad, 0xa5, 0xd0, 0x4a, 0xc5, 0x69, 0xba, 0x7f,
	0xd0, 0xd6, 0xa2, 0x49, 0x9a, 0x70, 0x4c, 0xf8, 0xd1, 0x08, 0xa2, 0xe9, 0x61, 0x84, 0x7b, 0xe8,
	0x56, 0xbd, 0x2a, 0x89, 0x5f, 0x49, 0xf2, 0xb5, 0x92, 0xa4, 0x14, 0x9e, 0x0c, 0x34, 0x7a, 0x33,
	0xc0, 0x99, 0xb3, 0x74, 0x93, 0x23, 0x23, 0x7f, 0xa0, 0x3e, 0x61, 0x74, 0x13, 0xc9, 0xe7, 0xa7,
	0x91, 0xf4, 0x4b, 0xfe, 0xf1, 0x65, 0xe7, 0xf2, 0xc4, 0x67, 0xbd, 0x4f, 0x64, 0x0c, 0x4d, 0x29,
	0x2f, 0x6f, 0x6b, 0xcf, 0x41, 0xaa, 0x33, 0xd1, 0xa6, 0x4b, 0xe8, 0xee, 0x9a, 0xca, 0x29, 0xec,
	0x79, 0xbf, 0x1f, 0xa5, 0x50, 0x2d, 0xe6, 0xb6, 0xfc, 0xd9, 0xc6, 0x1f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x2e, 0x41, 0x45, 0x05, 0x80, 0x03, 0x00, 0x00,
}
