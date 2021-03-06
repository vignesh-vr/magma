// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lte/protos/abort_session.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type AbortSessionResult_Code int32

const (
	AbortSessionResult_SESSION_REMOVED     AbortSessionResult_Code = 0
	AbortSessionResult_SESSION_NOT_FOUND   AbortSessionResult_Code = 1
	AbortSessionResult_USER_NOT_FOUND      AbortSessionResult_Code = 2
	AbortSessionResult_GATEWAY_NOT_FOUND   AbortSessionResult_Code = 3
	AbortSessionResult_RADIUS_SERVER_ERROR AbortSessionResult_Code = 4
)

var AbortSessionResult_Code_name = map[int32]string{
	0: "SESSION_REMOVED",
	1: "SESSION_NOT_FOUND",
	2: "USER_NOT_FOUND",
	3: "GATEWAY_NOT_FOUND",
	4: "RADIUS_SERVER_ERROR",
}

var AbortSessionResult_Code_value = map[string]int32{
	"SESSION_REMOVED":     0,
	"SESSION_NOT_FOUND":   1,
	"USER_NOT_FOUND":      2,
	"GATEWAY_NOT_FOUND":   3,
	"RADIUS_SERVER_ERROR": 4,
}

func (x AbortSessionResult_Code) String() string {
	return proto.EnumName(AbortSessionResult_Code_name, int32(x))
}

func (AbortSessionResult_Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_80dacefcd8bf4976, []int{1, 0}
}

type AbortSessionRequest struct {
	SessionId            string   `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AbortSessionRequest) Reset()         { *m = AbortSessionRequest{} }
func (m *AbortSessionRequest) String() string { return proto.CompactTextString(m) }
func (*AbortSessionRequest) ProtoMessage()    {}
func (*AbortSessionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_80dacefcd8bf4976, []int{0}
}

func (m *AbortSessionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AbortSessionRequest.Unmarshal(m, b)
}
func (m *AbortSessionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AbortSessionRequest.Marshal(b, m, deterministic)
}
func (m *AbortSessionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AbortSessionRequest.Merge(m, src)
}
func (m *AbortSessionRequest) XXX_Size() int {
	return xxx_messageInfo_AbortSessionRequest.Size(m)
}
func (m *AbortSessionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AbortSessionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AbortSessionRequest proto.InternalMessageInfo

func (m *AbortSessionRequest) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *AbortSessionRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

type AbortSessionResult struct {
	Code                 AbortSessionResult_Code `protobuf:"varint,1,opt,name=code,proto3,enum=magma.lte.AbortSessionResult_Code" json:"code,omitempty"`
	ErrorMessage         string                  `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *AbortSessionResult) Reset()         { *m = AbortSessionResult{} }
func (m *AbortSessionResult) String() string { return proto.CompactTextString(m) }
func (*AbortSessionResult) ProtoMessage()    {}
func (*AbortSessionResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_80dacefcd8bf4976, []int{1}
}

func (m *AbortSessionResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AbortSessionResult.Unmarshal(m, b)
}
func (m *AbortSessionResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AbortSessionResult.Marshal(b, m, deterministic)
}
func (m *AbortSessionResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AbortSessionResult.Merge(m, src)
}
func (m *AbortSessionResult) XXX_Size() int {
	return xxx_messageInfo_AbortSessionResult.Size(m)
}
func (m *AbortSessionResult) XXX_DiscardUnknown() {
	xxx_messageInfo_AbortSessionResult.DiscardUnknown(m)
}

var xxx_messageInfo_AbortSessionResult proto.InternalMessageInfo

func (m *AbortSessionResult) GetCode() AbortSessionResult_Code {
	if m != nil {
		return m.Code
	}
	return AbortSessionResult_SESSION_REMOVED
}

func (m *AbortSessionResult) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func init() {
	proto.RegisterEnum("magma.lte.AbortSessionResult_Code", AbortSessionResult_Code_name, AbortSessionResult_Code_value)
	proto.RegisterType((*AbortSessionRequest)(nil), "magma.lte.AbortSessionRequest")
	proto.RegisterType((*AbortSessionResult)(nil), "magma.lte.AbortSessionResult")
}

func init() { proto.RegisterFile("lte/protos/abort_session.proto", fileDescriptor_80dacefcd8bf4976) }

var fileDescriptor_80dacefcd8bf4976 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x4d, 0x4f, 0xf2, 0x40,
	0x10, 0x80, 0x29, 0x2f, 0x79, 0x63, 0x27, 0x88, 0xb8, 0x84, 0x88, 0x12, 0x88, 0xa9, 0x17, 0x4f,
	0x6d, 0x82, 0x89, 0xf7, 0x2a, 0xab, 0xe1, 0x40, 0x1b, 0x77, 0x01, 0xa3, 0x97, 0x4d, 0xa1, 0x13,
	0x24, 0x69, 0x59, 0xdc, 0x6d, 0xfd, 0xcf, 0xfe, 0x0b, 0xd3, 0x05, 0x4c, 0xd5, 0xe8, 0x69, 0x93,
	0x67, 0x9e, 0x99, 0x9d, 0x0f, 0xe8, 0x27, 0x19, 0x7a, 0x1b, 0x25, 0x33, 0xa9, 0xbd, 0x68, 0x2e,
	0x55, 0x26, 0x34, 0x6a, 0xbd, 0x92, 0x6b, 0xd7, 0x40, 0x62, 0xa7, 0xd1, 0x32, 0x8d, 0xdc, 0x24,
	0x43, 0xe7, 0x01, 0x5a, 0x7e, 0x61, 0xf0, 0xad, 0xc0, 0xf0, 0x35, 0x47, 0x9d, 0x91, 0x1e, 0xc0,
	0x2e, 0x45, 0xac, 0xe2, 0x8e, 0x75, 0x6e, 0x5d, 0xda, 0xcc, 0xde, 0x91, 0x51, 0x4c, 0xba, 0x60,
	0xe7, 0x1a, 0x95, 0x58, 0x47, 0x29, 0x76, 0xaa, 0x26, 0x7a, 0x50, 0x80, 0x20, 0x4a, 0xd1, 0x79,
	0xb7, 0x80, 0x7c, 0xad, 0xa9, 0xf3, 0x24, 0x23, 0xd7, 0x50, 0x5b, 0xc8, 0x18, 0x4d, 0xb1, 0xc6,
	0xc0, 0x71, 0x3f, 0x7b, 0x70, 0x7f, 0xca, 0xee, 0xad, 0x8c, 0x91, 0x19, 0x9f, 0x5c, 0xc0, 0x21,
	0x2a, 0x25, 0x95, 0x48, 0x51, 0xeb, 0x68, 0xb9, 0xff, 0xaf, 0x6e, 0xe0, 0x78, 0xcb, 0x9c, 0x37,
	0xa8, 0x15, 0x29, 0xa4, 0x05, 0x47, 0x9c, 0x72, 0x3e, 0x0a, 0x03, 0xc1, 0xe8, 0x38, 0x9c, 0xd1,
	0x61, 0xb3, 0x42, 0xda, 0x70, 0xbc, 0x87, 0x41, 0x38, 0x11, 0x77, 0xe1, 0x34, 0x18, 0x36, 0x2d,
	0x42, 0xa0, 0x31, 0xe5, 0x94, 0x95, 0x58, 0xb5, 0x50, 0xef, 0xfd, 0x09, 0x7d, 0xf4, 0x9f, 0x4a,
	0xf8, 0x1f, 0x39, 0x81, 0x16, 0xf3, 0x87, 0xa3, 0x29, 0x17, 0x9c, 0xb2, 0x19, 0x65, 0x82, 0x32,
	0x16, 0xb2, 0x66, 0x6d, 0xf0, 0x02, 0xed, 0x6f, 0xdd, 0x6f, 0xe4, 0x3a, 0x46, 0x45, 0x42, 0xa8,
	0x97, 0x03, 0xa4, 0xff, 0xeb, 0xbc, 0x66, 0xe1, 0x67, 0xbd, 0x3f, 0xf7, 0xe1, 0x54, 0x6e, 0xba,
	0xcf, 0xa7, 0xc6, 0xf0, 0x8a, 0xdb, 0x2e, 0x12, 0x99, 0xc7, 0xde, 0x52, 0xee, 0x8e, 0x3c, 0xff,
	0x6f, 0xde, 0xab, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x7f, 0x18, 0x79, 0xf9, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AbortSessionResponderClient is the client API for AbortSessionResponder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AbortSessionResponderClient interface {
	AbortSession(ctx context.Context, in *AbortSessionRequest, opts ...grpc.CallOption) (*AbortSessionResult, error)
}

type abortSessionResponderClient struct {
	cc grpc.ClientConnInterface
}

func NewAbortSessionResponderClient(cc grpc.ClientConnInterface) AbortSessionResponderClient {
	return &abortSessionResponderClient{cc}
}

func (c *abortSessionResponderClient) AbortSession(ctx context.Context, in *AbortSessionRequest, opts ...grpc.CallOption) (*AbortSessionResult, error) {
	out := new(AbortSessionResult)
	err := c.cc.Invoke(ctx, "/magma.lte.AbortSessionResponder/AbortSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AbortSessionResponderServer is the server API for AbortSessionResponder service.
type AbortSessionResponderServer interface {
	AbortSession(context.Context, *AbortSessionRequest) (*AbortSessionResult, error)
}

// UnimplementedAbortSessionResponderServer can be embedded to have forward compatible implementations.
type UnimplementedAbortSessionResponderServer struct {
}

func (*UnimplementedAbortSessionResponderServer) AbortSession(ctx context.Context, req *AbortSessionRequest) (*AbortSessionResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AbortSession not implemented")
}

func RegisterAbortSessionResponderServer(s *grpc.Server, srv AbortSessionResponderServer) {
	s.RegisterService(&_AbortSessionResponder_serviceDesc, srv)
}

func _AbortSessionResponder_AbortSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AbortSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AbortSessionResponderServer).AbortSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.AbortSessionResponder/AbortSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AbortSessionResponderServer).AbortSession(ctx, req.(*AbortSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AbortSessionResponder_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.lte.AbortSessionResponder",
	HandlerType: (*AbortSessionResponderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AbortSession",
			Handler:    _AbortSessionResponder_AbortSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lte/protos/abort_session.proto",
}
