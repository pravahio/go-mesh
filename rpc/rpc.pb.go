// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

package in_soket_rpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type PeerTopicInfo struct {
	Peer                 string   `protobuf:"bytes,1,opt,name=peer,proto3" json:"peer,omitempty"`
	Topics               []string `protobuf:"bytes,2,rep,name=topics,proto3" json:"topics,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeerTopicInfo) Reset()         { *m = PeerTopicInfo{} }
func (m *PeerTopicInfo) String() string { return proto.CompactTextString(m) }
func (*PeerTopicInfo) ProtoMessage()    {}
func (*PeerTopicInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{0}
}

func (m *PeerTopicInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerTopicInfo.Unmarshal(m, b)
}
func (m *PeerTopicInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerTopicInfo.Marshal(b, m, deterministic)
}
func (m *PeerTopicInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerTopicInfo.Merge(m, src)
}
func (m *PeerTopicInfo) XXX_Size() int {
	return xxx_messageInfo_PeerTopicInfo.Size(m)
}
func (m *PeerTopicInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerTopicInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PeerTopicInfo proto.InternalMessageInfo

func (m *PeerTopicInfo) GetPeer() string {
	if m != nil {
		return m.Peer
	}
	return ""
}

func (m *PeerTopicInfo) GetTopics() []string {
	if m != nil {
		return m.Topics
	}
	return nil
}

type PublishData struct {
	Info                 *PeerTopicInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Data                 *Data          `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PublishData) Reset()         { *m = PublishData{} }
func (m *PublishData) String() string { return proto.CompactTextString(m) }
func (*PublishData) ProtoMessage()    {}
func (*PublishData) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{1}
}

func (m *PublishData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishData.Unmarshal(m, b)
}
func (m *PublishData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishData.Marshal(b, m, deterministic)
}
func (m *PublishData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishData.Merge(m, src)
}
func (m *PublishData) XXX_Size() int {
	return xxx_messageInfo_PublishData.Size(m)
}
func (m *PublishData) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishData.DiscardUnknown(m)
}

var xxx_messageInfo_PublishData proto.InternalMessageInfo

func (m *PublishData) GetInfo() *PeerTopicInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *PublishData) GetData() *Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type Response struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Data struct {
	Topic                []string `protobuf:"bytes,1,rep,name=topic,proto3" json:"topic,omitempty"`
	Raw                  []byte   `protobuf:"bytes,2,opt,name=raw,proto3" json:"raw,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{3}
}

func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetTopic() []string {
	if m != nil {
		return m.Topic
	}
	return nil
}

func (m *Data) GetRaw() []byte {
	if m != nil {
		return m.Raw
	}
	return nil
}

func init() {
	proto.RegisterType((*PeerTopicInfo)(nil), "in.soket.rpc.PeerTopicInfo")
	proto.RegisterType((*PublishData)(nil), "in.soket.rpc.PublishData")
	proto.RegisterType((*Response)(nil), "in.soket.rpc.Response")
	proto.RegisterType((*Data)(nil), "in.soket.rpc.Data")
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor_77a6da22d6a3feb1) }

var fileDescriptor_77a6da22d6a3feb1 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xc1, 0x4a, 0x3b, 0x31,
	0x10, 0xc6, 0xbb, 0xed, 0xfe, 0xdb, 0xff, 0x4e, 0x2b, 0xe8, 0x28, 0x65, 0xd5, 0x4b, 0x09, 0x22,
	0x3d, 0x45, 0xa9, 0x47, 0x41, 0x3c, 0xf4, 0xa2, 0x28, 0x94, 0x68, 0x1f, 0x20, 0xbb, 0x4e, 0xdb,
	0xa0, 0x26, 0x21, 0x49, 0xf1, 0x95, 0x7d, 0x0c, 0xd9, 0xd8, 0xa2, 0x15, 0x5d, 0x3c, 0x78, 0x9b,
	0x61, 0xe6, 0xfb, 0xbe, 0x5f, 0x86, 0x40, 0xe6, 0x6c, 0xc9, 0xad, 0x33, 0xc1, 0x60, 0x4f, 0x69,
	0xee, 0xcd, 0x23, 0x05, 0xee, 0x6c, 0xc9, 0xce, 0x61, 0x6b, 0x42, 0xe4, 0xee, 0x8d, 0x55, 0xe5,
	0x95, 0x9e, 0x19, 0x44, 0x48, 0x2d, 0x91, 0xcb, 0x93, 0x41, 0x32, 0xcc, 0x44, 0xac, 0xb1, 0x0f,
	0xed, 0x50, 0x2d, 0xf8, 0xbc, 0x39, 0x68, 0x0d, 0x33, 0xb1, 0xea, 0xd8, 0x0c, 0xba, 0x93, 0x65,
	0xf1, 0xa4, 0xfc, 0x62, 0x2c, 0x83, 0xc4, 0x13, 0x48, 0x95, 0x9e, 0x99, 0x28, 0xed, 0x8e, 0x0e,
	0xf9, 0xe7, 0x20, 0xbe, 0x91, 0x22, 0xe2, 0x22, 0x1e, 0x43, 0xfa, 0x20, 0x83, 0xcc, 0x9b, 0x51,
	0x80, 0x9b, 0x82, 0xca, 0x52, 0xc4, 0x39, 0x3b, 0x82, 0xff, 0x82, 0xbc, 0x35, 0xda, 0x13, 0xe6,
	0xd0, 0x79, 0x26, 0xef, 0xe5, 0x9c, 0x56, 0x88, 0xeb, 0x96, 0x71, 0x48, 0x23, 0xc6, 0x1e, 0xfc,
	0x8b, 0x7c, 0x79, 0x12, 0x61, 0xdf, 0x1b, 0xdc, 0x86, 0x96, 0x93, 0x2f, 0x31, 0xaa, 0x27, 0xaa,
	0x72, 0xf4, 0xda, 0x84, 0xf4, 0x96, 0xfc, 0x02, 0xaf, 0x61, 0x47, 0xd0, 0x5c, 0xf9, 0x50, 0x11,
	0xae, 0x1e, 0x84, 0x75, 0xf8, 0x07, 0xfd, 0xcd, 0xe1, 0x1a, 0x8e, 0x35, 0xf0, 0x06, 0x76, 0xa7,
	0xda, 0xfd, 0x95, 0xdb, 0x05, 0x74, 0xd6, 0x0e, 0xfb, 0x5f, 0x1c, 0x3e, 0xee, 0x5e, 0xa3, 0xbf,
	0x84, 0xec, 0x6e, 0x59, 0xf8, 0xd2, 0xa9, 0x82, 0xea, 0x19, 0xbe, 0x39, 0x3e, 0x6b, 0x9c, 0x26,
	0x38, 0x86, 0xee, 0x54, 0xfb, 0xdf, 0x79, 0xfc, 0xc8, 0x51, 0xb4, 0xe3, 0xd7, 0x3b, 0x7b, 0x0b,
	0x00, 0x00, 0xff, 0xff, 0x7e, 0x42, 0xa4, 0xfe, 0x87, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MeshClient is the client API for Mesh service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MeshClient interface {
	RegisterToPublish(ctx context.Context, in *PeerTopicInfo, opts ...grpc.CallOption) (*Response, error)
	UnregisterToPublish(ctx context.Context, in *PeerTopicInfo, opts ...grpc.CallOption) (*Response, error)
	Publish(ctx context.Context, in *PublishData, opts ...grpc.CallOption) (*Response, error)
	Subscribe(ctx context.Context, in *PeerTopicInfo, opts ...grpc.CallOption) (Mesh_SubscribeClient, error)
	Unsubscribe(ctx context.Context, in *PeerTopicInfo, opts ...grpc.CallOption) (*Response, error)
}

type meshClient struct {
	cc *grpc.ClientConn
}

func NewMeshClient(cc *grpc.ClientConn) MeshClient {
	return &meshClient{cc}
}

func (c *meshClient) RegisterToPublish(ctx context.Context, in *PeerTopicInfo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/in.soket.rpc.Mesh/RegisterToPublish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshClient) UnregisterToPublish(ctx context.Context, in *PeerTopicInfo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/in.soket.rpc.Mesh/UnregisterToPublish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshClient) Publish(ctx context.Context, in *PublishData, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/in.soket.rpc.Mesh/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshClient) Subscribe(ctx context.Context, in *PeerTopicInfo, opts ...grpc.CallOption) (Mesh_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Mesh_serviceDesc.Streams[0], "/in.soket.rpc.Mesh/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &meshSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Mesh_SubscribeClient interface {
	Recv() (*Data, error)
	grpc.ClientStream
}

type meshSubscribeClient struct {
	grpc.ClientStream
}

func (x *meshSubscribeClient) Recv() (*Data, error) {
	m := new(Data)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *meshClient) Unsubscribe(ctx context.Context, in *PeerTopicInfo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/in.soket.rpc.Mesh/Unsubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MeshServer is the server API for Mesh service.
type MeshServer interface {
	RegisterToPublish(context.Context, *PeerTopicInfo) (*Response, error)
	UnregisterToPublish(context.Context, *PeerTopicInfo) (*Response, error)
	Publish(context.Context, *PublishData) (*Response, error)
	Subscribe(*PeerTopicInfo, Mesh_SubscribeServer) error
	Unsubscribe(context.Context, *PeerTopicInfo) (*Response, error)
}

func RegisterMeshServer(s *grpc.Server, srv MeshServer) {
	s.RegisterService(&_Mesh_serviceDesc, srv)
}

func _Mesh_RegisterToPublish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeerTopicInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshServer).RegisterToPublish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/in.soket.rpc.Mesh/RegisterToPublish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshServer).RegisterToPublish(ctx, req.(*PeerTopicInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mesh_UnregisterToPublish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeerTopicInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshServer).UnregisterToPublish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/in.soket.rpc.Mesh/UnregisterToPublish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshServer).UnregisterToPublish(ctx, req.(*PeerTopicInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mesh_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/in.soket.rpc.Mesh/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshServer).Publish(ctx, req.(*PublishData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mesh_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PeerTopicInfo)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MeshServer).Subscribe(m, &meshSubscribeServer{stream})
}

type Mesh_SubscribeServer interface {
	Send(*Data) error
	grpc.ServerStream
}

type meshSubscribeServer struct {
	grpc.ServerStream
}

func (x *meshSubscribeServer) Send(m *Data) error {
	return x.ServerStream.SendMsg(m)
}

func _Mesh_Unsubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeerTopicInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshServer).Unsubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/in.soket.rpc.Mesh/Unsubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshServer).Unsubscribe(ctx, req.(*PeerTopicInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _Mesh_serviceDesc = grpc.ServiceDesc{
	ServiceName: "in.soket.rpc.Mesh",
	HandlerType: (*MeshServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterToPublish",
			Handler:    _Mesh_RegisterToPublish_Handler,
		},
		{
			MethodName: "UnregisterToPublish",
			Handler:    _Mesh_UnregisterToPublish_Handler,
		},
		{
			MethodName: "Publish",
			Handler:    _Mesh_Publish_Handler,
		},
		{
			MethodName: "Unsubscribe",
			Handler:    _Mesh_Unsubscribe_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _Mesh_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "rpc.proto",
}
