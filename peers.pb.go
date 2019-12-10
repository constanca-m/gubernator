// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peers.proto

package gubernator

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

type GetPeerRateLimitsReq struct {
	// Must specify at least one RateLimit. The peer that recives this request MUST be authoritative for
	// each rate_limit[x].unique_key provided, as the peer will not forward the request to any other peers
	Requests []*RateLimitReq `protobuf:"bytes,1,rep,name=requests" json:"requests,omitempty"`
}

func (m *GetPeerRateLimitsReq) Reset()                    { *m = GetPeerRateLimitsReq{} }
func (m *GetPeerRateLimitsReq) String() string            { return proto.CompactTextString(m) }
func (*GetPeerRateLimitsReq) ProtoMessage()               {}
func (*GetPeerRateLimitsReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *GetPeerRateLimitsReq) GetRequests() []*RateLimitReq {
	if m != nil {
		return m.Requests
	}
	return nil
}

type GetPeerRateLimitsResp struct {
	// Responses are in the same order as they appeared in the PeerRateLimitRequests
	RateLimits []*RateLimitResp `protobuf:"bytes,1,rep,name=rate_limits,json=rateLimits" json:"rate_limits,omitempty"`
}

func (m *GetPeerRateLimitsResp) Reset()                    { *m = GetPeerRateLimitsResp{} }
func (m *GetPeerRateLimitsResp) String() string            { return proto.CompactTextString(m) }
func (*GetPeerRateLimitsResp) ProtoMessage()               {}
func (*GetPeerRateLimitsResp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *GetPeerRateLimitsResp) GetRateLimits() []*RateLimitResp {
	if m != nil {
		return m.RateLimits
	}
	return nil
}

type UpdatePeerGlobalsReq struct {
	// Must specify at least one RateLimit
	Globals []*UpdatePeerGlobal `protobuf:"bytes,1,rep,name=globals" json:"globals,omitempty"`
}

func (m *UpdatePeerGlobalsReq) Reset()                    { *m = UpdatePeerGlobalsReq{} }
func (m *UpdatePeerGlobalsReq) String() string            { return proto.CompactTextString(m) }
func (*UpdatePeerGlobalsReq) ProtoMessage()               {}
func (*UpdatePeerGlobalsReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *UpdatePeerGlobalsReq) GetGlobals() []*UpdatePeerGlobal {
	if m != nil {
		return m.Globals
	}
	return nil
}

type UpdatePeerGlobal struct {
	Key    string         `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Status *RateLimitResp `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
}

func (m *UpdatePeerGlobal) Reset()                    { *m = UpdatePeerGlobal{} }
func (m *UpdatePeerGlobal) String() string            { return proto.CompactTextString(m) }
func (*UpdatePeerGlobal) ProtoMessage()               {}
func (*UpdatePeerGlobal) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *UpdatePeerGlobal) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *UpdatePeerGlobal) GetStatus() *RateLimitResp {
	if m != nil {
		return m.Status
	}
	return nil
}

type UpdatePeerGlobalsResp struct {
}

func (m *UpdatePeerGlobalsResp) Reset()                    { *m = UpdatePeerGlobalsResp{} }
func (m *UpdatePeerGlobalsResp) String() string            { return proto.CompactTextString(m) }
func (*UpdatePeerGlobalsResp) ProtoMessage()               {}
func (*UpdatePeerGlobalsResp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func init() {
	proto.RegisterType((*GetPeerRateLimitsReq)(nil), "pb.gubernator.GetPeerRateLimitsReq")
	proto.RegisterType((*GetPeerRateLimitsResp)(nil), "pb.gubernator.GetPeerRateLimitsResp")
	proto.RegisterType((*UpdatePeerGlobalsReq)(nil), "pb.gubernator.UpdatePeerGlobalsReq")
	proto.RegisterType((*UpdatePeerGlobal)(nil), "pb.gubernator.UpdatePeerGlobal")
	proto.RegisterType((*UpdatePeerGlobalsResp)(nil), "pb.gubernator.UpdatePeerGlobalsResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PeersV1 service

type PeersV1Client interface {
	// Used by peers to relay batches of requests to an authoritative peer
	GetPeerRateLimits(ctx context.Context, in *GetPeerRateLimitsReq, opts ...grpc.CallOption) (*GetPeerRateLimitsResp, error)
	// Used by peers send global rate limit updates to other peers
	UpdatePeerGlobals(ctx context.Context, in *UpdatePeerGlobalsReq, opts ...grpc.CallOption) (*UpdatePeerGlobalsResp, error)
}

type peersV1Client struct {
	cc *grpc.ClientConn
}

func NewPeersV1Client(cc *grpc.ClientConn) PeersV1Client {
	return &peersV1Client{cc}
}

func (c *peersV1Client) GetPeerRateLimits(ctx context.Context, in *GetPeerRateLimitsReq, opts ...grpc.CallOption) (*GetPeerRateLimitsResp, error) {
	out := new(GetPeerRateLimitsResp)
	err := grpc.Invoke(ctx, "/pb.gubernator.PeersV1/GetPeerRateLimits", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peersV1Client) UpdatePeerGlobals(ctx context.Context, in *UpdatePeerGlobalsReq, opts ...grpc.CallOption) (*UpdatePeerGlobalsResp, error) {
	out := new(UpdatePeerGlobalsResp)
	err := grpc.Invoke(ctx, "/pb.gubernator.PeersV1/UpdatePeerGlobals", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PeersV1 service

type PeersV1Server interface {
	// Used by peers to relay batches of requests to an authoritative peer
	GetPeerRateLimits(context.Context, *GetPeerRateLimitsReq) (*GetPeerRateLimitsResp, error)
	// Used by peers send global rate limit updates to other peers
	UpdatePeerGlobals(context.Context, *UpdatePeerGlobalsReq) (*UpdatePeerGlobalsResp, error)
}

func RegisterPeersV1Server(s *grpc.Server, srv PeersV1Server) {
	s.RegisterService(&_PeersV1_serviceDesc, srv)
}

func _PeersV1_GetPeerRateLimits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPeerRateLimitsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeersV1Server).GetPeerRateLimits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.gubernator.PeersV1/GetPeerRateLimits",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeersV1Server).GetPeerRateLimits(ctx, req.(*GetPeerRateLimitsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeersV1_UpdatePeerGlobals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePeerGlobalsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeersV1Server).UpdatePeerGlobals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.gubernator.PeersV1/UpdatePeerGlobals",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeersV1Server).UpdatePeerGlobals(ctx, req.(*UpdatePeerGlobalsReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _PeersV1_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.gubernator.PeersV1",
	HandlerType: (*PeersV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPeerRateLimits",
			Handler:    _PeersV1_GetPeerRateLimits_Handler,
		},
		{
			MethodName: "UpdatePeerGlobals",
			Handler:    _PeersV1_UpdatePeerGlobals_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "peers.proto",
}

func init() { proto.RegisterFile("peers.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xbb, 0x16, 0x5a, 0x9d, 0x20, 0xc6, 0xa5, 0xc5, 0x50, 0x05, 0xc3, 0xea, 0xa1, 0xa7,
	0x80, 0x55, 0x10, 0x0f, 0x5e, 0xbc, 0xf4, 0x22, 0xa8, 0x0b, 0xf6, 0xd0, 0x8b, 0x6e, 0x70, 0x28,
	0xc5, 0x68, 0x36, 0x3b, 0x9b, 0x83, 0x37, 0x5f, 0xd0, 0x77, 0x92, 0xa4, 0x31, 0xc5, 0x64, 0x4b,
	0x6f, 0x93, 0xc9, 0x97, 0x6f, 0xfe, 0xcc, 0x2e, 0x78, 0x1a, 0xd1, 0x50, 0xa4, 0x4d, 0x6a, 0x53,
	0xbe, 0xaf, 0xe3, 0x68, 0x91, 0xc7, 0x68, 0x3e, 0x95, 0x4d, 0xcd, 0xc8, 0x5f, 0xd7, 0x2b, 0x40,
	0x3c, 0xc0, 0x60, 0x8a, 0xf6, 0x11, 0xd1, 0x48, 0x65, 0xf1, 0x7e, 0xf9, 0xb1, 0xb4, 0x24, 0x31,
	0xe3, 0xd7, 0xb0, 0x6b, 0x30, 0xcb, 0x91, 0x2c, 0x05, 0x2c, 0xec, 0x8e, 0xbd, 0xc9, 0x71, 0xf4,
	0xcf, 0x15, 0xd5, 0xbc, 0xc4, 0x4c, 0xd6, 0xb0, 0x98, 0xc1, 0xd0, 0x21, 0x24, 0xcd, 0x6f, 0xc1,
	0x33, 0xca, 0xe2, 0x4b, 0x52, 0xb6, 0x2a, 0xe9, 0xc9, 0x66, 0x29, 0x69, 0x09, 0xa6, 0x56, 0x88,
	0x27, 0x18, 0x3c, 0xeb, 0x37, 0x65, 0xb1, 0x50, 0x4f, 0x93, 0x34, 0x56, 0x49, 0x19, 0xf4, 0x06,
	0xfa, 0x8b, 0xd5, 0x53, 0xa5, 0x3c, 0x6d, 0x28, 0x9b, 0x5f, 0xc9, 0x3f, 0x5e, 0xcc, 0xc1, 0x6f,
	0xbe, 0xe4, 0x3e, 0x74, 0xdf, 0xf1, 0x2b, 0x60, 0x21, 0x1b, 0xef, 0xc9, 0xa2, 0xe4, 0x57, 0xd0,
	0x23, 0xab, 0x6c, 0x4e, 0xc1, 0x4e, 0xc8, 0xb6, 0x46, 0xae, 0x58, 0x71, 0x04, 0x43, 0x47, 0x5c,
	0xd2, 0x93, 0x1f, 0x06, 0xfd, 0xa2, 0x47, 0xb3, 0x0b, 0xfe, 0x0a, 0x87, 0xad, 0x5d, 0xf1, 0xb3,
	0x86, 0xdf, 0x75, 0x3c, 0xa3, 0xf3, 0xed, 0x10, 0x69, 0xd1, 0x29, 0x26, 0xb4, 0x62, 0xb4, 0x26,
	0xb8, 0xf6, 0xda, 0x9a, 0xe0, 0xfc, 0x1b, 0xd1, 0xb9, 0x3b, 0x98, 0xc3, 0x9a, 0xfa, 0x66, 0x2c,
	0xee, 0x95, 0x17, 0xeb, 0xf2, 0x37, 0x00, 0x00, 0xff, 0xff, 0x19, 0x46, 0x07, 0x4d, 0x88, 0x02,
	0x00, 0x00,
}
