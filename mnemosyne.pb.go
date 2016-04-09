// Code generated by protoc-gen-go.
// source: mnemosyne.proto
// DO NOT EDIT!

/*
Package mnemosyne is a generated protocol buffer package.

It is generated from these files:
	mnemosyne.proto

It has these top-level messages:
	Empty
	AccessToken
	Session
	GetRequest
	GetResponse
	ListRequest
	ListResponse
	ExistsRequest
	ExistsResponse
	StartRequest
	StartResponse
	AbandonRequest
	AbandonResponse
	SetValueRequest
	SetValueResponse
	DeleteValueRequest
	DeleteValueResponse
	ClearRequest
	ClearResponse
	DeleteRequest
	DeleteResponse
*/
package mnemosyne

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

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
const _ = proto.ProtoPackageIsVersion1

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// AccessToken represents identifier of single session. It consist of partition key and a hash.
type AccessToken struct {
	Key  []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Hash []byte `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *AccessToken) Reset()                    { *m = AccessToken{} }
func (m *AccessToken) String() string            { return proto.CompactTextString(m) }
func (*AccessToken) ProtoMessage()               {}
func (*AccessToken) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Session struct {
	AccessToken *AccessToken               `protobuf:"bytes,1,opt,name=access_token" json:"access_token,omitempty"`
	SubjectId   string                     `protobuf:"bytes,2,opt,name=subject_id" json:"subject_id,omitempty"`
	Bag         map[string]string          `protobuf:"bytes,3,rep,name=bag" json:"bag,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	ExpireAt    *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=expire_at" json:"expire_at,omitempty"`
}

func (m *Session) Reset()                    { *m = Session{} }
func (m *Session) String() string            { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()               {}
func (*Session) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Session) GetAccessToken() *AccessToken {
	if m != nil {
		return m.AccessToken
	}
	return nil
}

func (m *Session) GetBag() map[string]string {
	if m != nil {
		return m.Bag
	}
	return nil
}

func (m *Session) GetExpireAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.ExpireAt
	}
	return nil
}

type GetRequest struct {
	AccessToken *AccessToken `protobuf:"bytes,1,opt,name=access_token" json:"access_token,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetRequest) GetAccessToken() *AccessToken {
	if m != nil {
		return m.AccessToken
	}
	return nil
}

type GetResponse struct {
	Session *Session `protobuf:"bytes,1,opt,name=session" json:"session,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetResponse) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

type ListRequest struct {
	Offset       int64                      `protobuf:"varint,1,opt,name=offset" json:"offset,omitempty"`
	Limit        int64                      `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
	ExpireAtFrom *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=expire_at_from" json:"expire_at_from,omitempty"`
	ExpireAtTo   *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=expire_at_to" json:"expire_at_to,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ListRequest) GetExpireAtFrom() *google_protobuf.Timestamp {
	if m != nil {
		return m.ExpireAtFrom
	}
	return nil
}

func (m *ListRequest) GetExpireAtTo() *google_protobuf.Timestamp {
	if m != nil {
		return m.ExpireAtTo
	}
	return nil
}

type ListResponse struct {
	Sessions []*Session `protobuf:"bytes,1,rep,name=sessions" json:"sessions,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ListResponse) GetSessions() []*Session {
	if m != nil {
		return m.Sessions
	}
	return nil
}

type ExistsRequest struct {
	AccessToken *AccessToken `protobuf:"bytes,1,opt,name=access_token" json:"access_token,omitempty"`
}

func (m *ExistsRequest) Reset()                    { *m = ExistsRequest{} }
func (m *ExistsRequest) String() string            { return proto.CompactTextString(m) }
func (*ExistsRequest) ProtoMessage()               {}
func (*ExistsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ExistsRequest) GetAccessToken() *AccessToken {
	if m != nil {
		return m.AccessToken
	}
	return nil
}

type ExistsResponse struct {
	Exists bool `protobuf:"varint,1,opt,name=exists" json:"exists,omitempty"`
}

func (m *ExistsResponse) Reset()                    { *m = ExistsResponse{} }
func (m *ExistsResponse) String() string            { return proto.CompactTextString(m) }
func (*ExistsResponse) ProtoMessage()               {}
func (*ExistsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type StartRequest struct {
	SubjectId string            `protobuf:"bytes,1,opt,name=subject_id" json:"subject_id,omitempty"`
	Bag       map[string]string `protobuf:"bytes,2,rep,name=bag" json:"bag,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *StartRequest) Reset()                    { *m = StartRequest{} }
func (m *StartRequest) String() string            { return proto.CompactTextString(m) }
func (*StartRequest) ProtoMessage()               {}
func (*StartRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *StartRequest) GetBag() map[string]string {
	if m != nil {
		return m.Bag
	}
	return nil
}

type StartResponse struct {
	Session *Session `protobuf:"bytes,1,opt,name=session" json:"session,omitempty"`
}

func (m *StartResponse) Reset()                    { *m = StartResponse{} }
func (m *StartResponse) String() string            { return proto.CompactTextString(m) }
func (*StartResponse) ProtoMessage()               {}
func (*StartResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *StartResponse) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

type AbandonRequest struct {
	AccessToken *AccessToken `protobuf:"bytes,1,opt,name=access_token" json:"access_token,omitempty"`
}

func (m *AbandonRequest) Reset()                    { *m = AbandonRequest{} }
func (m *AbandonRequest) String() string            { return proto.CompactTextString(m) }
func (*AbandonRequest) ProtoMessage()               {}
func (*AbandonRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *AbandonRequest) GetAccessToken() *AccessToken {
	if m != nil {
		return m.AccessToken
	}
	return nil
}

type AbandonResponse struct {
	Abandoned bool `protobuf:"varint,1,opt,name=abandoned" json:"abandoned,omitempty"`
}

func (m *AbandonResponse) Reset()                    { *m = AbandonResponse{} }
func (m *AbandonResponse) String() string            { return proto.CompactTextString(m) }
func (*AbandonResponse) ProtoMessage()               {}
func (*AbandonResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type SetValueRequest struct {
	AccessToken *AccessToken `protobuf:"bytes,1,opt,name=access_token" json:"access_token,omitempty"`
	Key         string       `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Value       string       `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
}

func (m *SetValueRequest) Reset()                    { *m = SetValueRequest{} }
func (m *SetValueRequest) String() string            { return proto.CompactTextString(m) }
func (*SetValueRequest) ProtoMessage()               {}
func (*SetValueRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *SetValueRequest) GetAccessToken() *AccessToken {
	if m != nil {
		return m.AccessToken
	}
	return nil
}

type SetValueResponse struct {
	Bag map[string]string `protobuf:"bytes,1,rep,name=bag" json:"bag,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *SetValueResponse) Reset()                    { *m = SetValueResponse{} }
func (m *SetValueResponse) String() string            { return proto.CompactTextString(m) }
func (*SetValueResponse) ProtoMessage()               {}
func (*SetValueResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *SetValueResponse) GetBag() map[string]string {
	if m != nil {
		return m.Bag
	}
	return nil
}

type DeleteValueRequest struct {
	AccessToken *AccessToken `protobuf:"bytes,1,opt,name=access_token" json:"access_token,omitempty"`
	Key         string       `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
}

func (m *DeleteValueRequest) Reset()                    { *m = DeleteValueRequest{} }
func (m *DeleteValueRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteValueRequest) ProtoMessage()               {}
func (*DeleteValueRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *DeleteValueRequest) GetAccessToken() *AccessToken {
	if m != nil {
		return m.AccessToken
	}
	return nil
}

type DeleteValueResponse struct {
	Session *Session `protobuf:"bytes,1,opt,name=session" json:"session,omitempty"`
}

func (m *DeleteValueResponse) Reset()                    { *m = DeleteValueResponse{} }
func (m *DeleteValueResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteValueResponse) ProtoMessage()               {}
func (*DeleteValueResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *DeleteValueResponse) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

type ClearRequest struct {
	AccessToken *AccessToken `protobuf:"bytes,1,opt,name=access_token" json:"access_token,omitempty"`
}

func (m *ClearRequest) Reset()                    { *m = ClearRequest{} }
func (m *ClearRequest) String() string            { return proto.CompactTextString(m) }
func (*ClearRequest) ProtoMessage()               {}
func (*ClearRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *ClearRequest) GetAccessToken() *AccessToken {
	if m != nil {
		return m.AccessToken
	}
	return nil
}

type ClearResponse struct {
}

func (m *ClearResponse) Reset()                    { *m = ClearResponse{} }
func (m *ClearResponse) String() string            { return proto.CompactTextString(m) }
func (*ClearResponse) ProtoMessage()               {}
func (*ClearResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

type DeleteRequest struct {
	AccessToken  *AccessToken               `protobuf:"bytes,1,opt,name=access_token" json:"access_token,omitempty"`
	ExpireAtFrom *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=expire_at_from" json:"expire_at_from,omitempty"`
	ExpireAtTo   *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=expire_at_to" json:"expire_at_to,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *DeleteRequest) GetAccessToken() *AccessToken {
	if m != nil {
		return m.AccessToken
	}
	return nil
}

func (m *DeleteRequest) GetExpireAtFrom() *google_protobuf.Timestamp {
	if m != nil {
		return m.ExpireAtFrom
	}
	return nil
}

func (m *DeleteRequest) GetExpireAtTo() *google_protobuf.Timestamp {
	if m != nil {
		return m.ExpireAtTo
	}
	return nil
}

type DeleteResponse struct {
	Count int64 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

func init() {
	proto.RegisterType((*Empty)(nil), "mnemosyne.Empty")
	proto.RegisterType((*AccessToken)(nil), "mnemosyne.AccessToken")
	proto.RegisterType((*Session)(nil), "mnemosyne.Session")
	proto.RegisterType((*GetRequest)(nil), "mnemosyne.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "mnemosyne.GetResponse")
	proto.RegisterType((*ListRequest)(nil), "mnemosyne.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "mnemosyne.ListResponse")
	proto.RegisterType((*ExistsRequest)(nil), "mnemosyne.ExistsRequest")
	proto.RegisterType((*ExistsResponse)(nil), "mnemosyne.ExistsResponse")
	proto.RegisterType((*StartRequest)(nil), "mnemosyne.StartRequest")
	proto.RegisterType((*StartResponse)(nil), "mnemosyne.StartResponse")
	proto.RegisterType((*AbandonRequest)(nil), "mnemosyne.AbandonRequest")
	proto.RegisterType((*AbandonResponse)(nil), "mnemosyne.AbandonResponse")
	proto.RegisterType((*SetValueRequest)(nil), "mnemosyne.SetValueRequest")
	proto.RegisterType((*SetValueResponse)(nil), "mnemosyne.SetValueResponse")
	proto.RegisterType((*DeleteValueRequest)(nil), "mnemosyne.DeleteValueRequest")
	proto.RegisterType((*DeleteValueResponse)(nil), "mnemosyne.DeleteValueResponse")
	proto.RegisterType((*ClearRequest)(nil), "mnemosyne.ClearRequest")
	proto.RegisterType((*ClearResponse)(nil), "mnemosyne.ClearResponse")
	proto.RegisterType((*DeleteRequest)(nil), "mnemosyne.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "mnemosyne.DeleteResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion1

// Client API for RPC service

type RPCClient interface {
	Context(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Session, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Exists(ctx context.Context, in *ExistsRequest, opts ...grpc.CallOption) (*ExistsResponse, error)
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error)
	Abandon(ctx context.Context, in *AbandonRequest, opts ...grpc.CallOption) (*AbandonResponse, error)
	SetValue(ctx context.Context, in *SetValueRequest, opts ...grpc.CallOption) (*SetValueResponse, error)
	//    rpc DeleteValue(DeleteValueRequest) returns (DeleteValueResponse) {};
	//    rpc Clear(ClearRequest) returns (ClearResponse) {};
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type rPCClient struct {
	cc *grpc.ClientConn
}

func NewRPCClient(cc *grpc.ClientConn) RPCClient {
	return &rPCClient{cc}
}

func (c *rPCClient) Context(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Session, error) {
	out := new(Session)
	err := grpc.Invoke(ctx, "/mnemosyne.RPC/Context", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/mnemosyne.RPC/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/mnemosyne.RPC/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) Exists(ctx context.Context, in *ExistsRequest, opts ...grpc.CallOption) (*ExistsResponse, error) {
	out := new(ExistsResponse)
	err := grpc.Invoke(ctx, "/mnemosyne.RPC/Exists", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error) {
	out := new(StartResponse)
	err := grpc.Invoke(ctx, "/mnemosyne.RPC/Start", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) Abandon(ctx context.Context, in *AbandonRequest, opts ...grpc.CallOption) (*AbandonResponse, error) {
	out := new(AbandonResponse)
	err := grpc.Invoke(ctx, "/mnemosyne.RPC/Abandon", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) SetValue(ctx context.Context, in *SetValueRequest, opts ...grpc.CallOption) (*SetValueResponse, error) {
	out := new(SetValueResponse)
	err := grpc.Invoke(ctx, "/mnemosyne.RPC/SetValue", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := grpc.Invoke(ctx, "/mnemosyne.RPC/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RPC service

type RPCServer interface {
	Context(context.Context, *Empty) (*Session, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	Exists(context.Context, *ExistsRequest) (*ExistsResponse, error)
	Start(context.Context, *StartRequest) (*StartResponse, error)
	Abandon(context.Context, *AbandonRequest) (*AbandonResponse, error)
	SetValue(context.Context, *SetValueRequest) (*SetValueResponse, error)
	//    rpc DeleteValue(DeleteValueRequest) returns (DeleteValueResponse) {};
	//    rpc Clear(ClearRequest) returns (ClearResponse) {};
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
}

func RegisterRPCServer(s *grpc.Server, srv RPCServer) {
	s.RegisterService(&_RPC_serviceDesc, srv)
}

func _RPC_Context_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(RPCServer).Context(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _RPC_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(RPCServer).Get(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _RPC_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(RPCServer).List(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _RPC_Exists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ExistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(RPCServer).Exists(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _RPC_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(RPCServer).Start(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _RPC_Abandon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(AbandonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(RPCServer).Abandon(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _RPC_SetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(SetValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(RPCServer).SetValue(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _RPC_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(RPCServer).Delete(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _RPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mnemosyne.RPC",
	HandlerType: (*RPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Context",
			Handler:    _RPC_Context_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _RPC_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _RPC_List_Handler,
		},
		{
			MethodName: "Exists",
			Handler:    _RPC_Exists_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _RPC_Start_Handler,
		},
		{
			MethodName: "Abandon",
			Handler:    _RPC_Abandon_Handler,
		},
		{
			MethodName: "SetValue",
			Handler:    _RPC_SetValue_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _RPC_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 709 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x55, 0x41, 0x73, 0x12, 0x31,
	0x18, 0x2d, 0x6c, 0x5b, 0xe0, 0x63, 0x81, 0x1a, 0xc7, 0x16, 0xb7, 0x07, 0x99, 0xd8, 0x43, 0xc7,
	0xb1, 0xa0, 0xd8, 0x71, 0xb4, 0x56, 0x9d, 0xb6, 0x32, 0x5e, 0x9c, 0xd1, 0x69, 0x3b, 0x5e, 0x3c,
	0x30, 0x81, 0x86, 0xed, 0x5a, 0x76, 0x43, 0x49, 0x70, 0xca, 0xc1, 0x3f, 0xe1, 0x5f, 0xf0, 0x0f,
	0xf9, 0x6b, 0x3c, 0x1b, 0x92, 0xb0, 0x64, 0x85, 0xca, 0xc0, 0x78, 0x83, 0x24, 0xef, 0xfb, 0xde,
	0x7b, 0xf9, 0xf2, 0x16, 0x4a, 0x61, 0x44, 0x43, 0xc6, 0x87, 0x11, 0xad, 0xf6, 0xfa, 0x4c, 0x30,
	0x94, 0x8b, 0x17, 0xbc, 0x57, 0x7e, 0x20, 0x2e, 0x07, 0xad, 0x6a, 0x9b, 0x85, 0x35, 0x9f, 0x75,
	0x49, 0xe4, 0xd7, 0xd4, 0x99, 0xd6, 0xa0, 0x53, 0xeb, 0x89, 0x61, 0x8f, 0xf2, 0x9a, 0x08, 0x42,
	0xca, 0x05, 0x09, 0x7b, 0x93, 0x5f, 0xba, 0x0e, 0xce, 0xc0, 0x5a, 0x23, 0x94, 0xa7, 0xf0, 0x2e,
	0xe4, 0x8f, 0xda, 0x6d, 0xca, 0xf9, 0x39, 0xbb, 0xa2, 0x11, 0xca, 0x83, 0x73, 0x45, 0x87, 0xe5,
	0x54, 0x25, 0xb5, 0xeb, 0x22, 0x17, 0x56, 0x2f, 0x09, 0xbf, 0x2c, 0xa7, 0x47, 0xff, 0xf0, 0xaf,
	0x14, 0x64, 0xce, 0xe4, 0xc1, 0x80, 0x45, 0xe8, 0x31, 0xb8, 0x44, 0xa1, 0x9a, 0x62, 0x04, 0x53,
	0xe7, 0xf3, 0xf5, 0xcd, 0xea, 0x84, 0xae, 0x5d, 0x14, 0x01, 0xf0, 0x41, 0xeb, 0x2b, 0x6d, 0x8b,
	0x66, 0x70, 0xa1, 0xaa, 0xe5, 0xd0, 0x2e, 0x38, 0x2d, 0xe2, 0x97, 0x9d, 0x8a, 0x23, 0x81, 0xdb,
	0x16, 0xd0, 0xb4, 0xa8, 0x1e, 0x13, 0xbf, 0x11, 0x89, 0xfe, 0x10, 0xed, 0x41, 0x8e, 0xde, 0xf4,
	0x82, 0x3e, 0x6d, 0x12, 0x51, 0x5e, 0x55, 0x8d, 0xbc, 0xaa, 0xcf, 0x98, 0xdf, 0x35, 0xa6, 0x48,
	0xc1, 0xd5, 0xf3, 0xb1, 0x3e, 0xef, 0x11, 0x64, 0x63, 0xa8, 0xa5, 0x26, 0x87, 0x0a, 0xb0, 0xf6,
	0x8d, 0x74, 0x07, 0x54, 0x13, 0x38, 0x48, 0xbf, 0x48, 0xe1, 0x03, 0x80, 0xf7, 0x54, 0x9c, 0xd2,
	0xeb, 0x81, 0x04, 0x2f, 0x26, 0x0a, 0xd7, 0x21, 0xaf, 0xb0, 0xbc, 0xc7, 0x22, 0x4e, 0xd1, 0x43,
	0xc8, 0x70, 0xcd, 0xdc, 0xe0, 0xd0, 0xb4, 0x26, 0xfc, 0x23, 0x05, 0xf9, 0x0f, 0x01, 0x8f, 0x3b,
	0x16, 0x61, 0x9d, 0x75, 0x3a, 0x9c, 0x0a, 0x85, 0x71, 0x46, 0x14, 0xbb, 0x41, 0x18, 0x08, 0x45,
	0xd1, 0x41, 0x75, 0x28, 0xc6, 0xca, 0x9b, 0x9d, 0x3e, 0x0b, 0xa5, 0x5d, 0x73, 0xe4, 0xa3, 0x27,
	0xe0, 0x4e, 0x30, 0x82, 0xcd, 0x37, 0x0c, 0xef, 0x83, 0xab, 0x39, 0x19, 0x25, 0x3b, 0x90, 0x35,
	0x4a, 0xb8, 0xa4, 0xe5, 0xdc, 0x22, 0xe5, 0x35, 0x14, 0x1a, 0x37, 0x12, 0xc6, 0x97, 0x73, 0xaf,
	0x02, 0xc5, 0x31, 0xdc, 0xb4, 0x95, 0x5e, 0x50, 0xb5, 0xa2, 0x90, 0x59, 0xfc, 0x1d, 0xdc, 0x33,
	0x41, 0xfa, 0xb1, 0x57, 0xc9, 0x21, 0xd2, 0x57, 0xba, 0xa7, 0x87, 0x28, 0xad, 0x58, 0x56, 0x6c,
	0x96, 0x16, 0x32, 0x9e, 0xa4, 0x85, 0x46, 0x63, 0x1f, 0x0a, 0xa6, 0xc8, 0x22, 0x17, 0xfc, 0x06,
	0x8a, 0x47, 0x2d, 0x12, 0x5d, 0xb0, 0x68, 0x39, 0x5b, 0x76, 0xa0, 0x14, 0xe3, 0x4d, 0xdf, 0x3b,
	0x90, 0x23, 0x7a, 0x89, 0x5e, 0x18, 0x6b, 0xbe, 0x40, 0xe9, 0x8c, 0x8a, 0xcf, 0x23, 0xc6, 0x4b,
	0xb5, 0x19, 0x8b, 0x4f, 0x27, 0xc5, 0x8f, 0x86, 0x2b, 0x87, 0xaf, 0x61, 0x63, 0x52, 0xdc, 0x70,
	0x78, 0xaa, 0x7d, 0xd6, 0xd3, 0xb0, 0x93, 0xd0, 0x9d, 0x3c, 0xb9, 0x9c, 0xd7, 0x1f, 0x01, 0xbd,
	0xa3, 0x5d, 0x2a, 0xe8, 0x7f, 0x92, 0x24, 0xdf, 0xf5, 0xdd, 0x44, 0xc1, 0x45, 0xae, 0xf0, 0x10,
	0xdc, 0x93, 0x2e, 0x25, 0xfd, 0xe5, 0x2e, 0xb0, 0x04, 0x05, 0x83, 0xd6, 0x3d, 0xf1, 0xcf, 0x14,
	0x14, 0x34, 0x97, 0xe5, 0x74, 0x4d, 0x67, 0x40, 0x7a, 0xe1, 0x0c, 0x98, 0x9b, 0x1a, 0xf8, 0x01,
	0x14, 0xc7, 0x24, 0x8d, 0x57, 0xf2, 0x9a, 0xda, 0x6c, 0x10, 0x99, 0x64, 0xaa, 0xff, 0x76, 0xc0,
	0x39, 0xfd, 0x74, 0x22, 0x27, 0x21, 0x73, 0xc2, 0x22, 0x41, 0x6f, 0x04, 0xda, 0xb0, 0x18, 0xab,
	0x6f, 0x89, 0x37, 0xcb, 0xce, 0x15, 0xf4, 0x1c, 0x1c, 0x19, 0x94, 0xe8, 0x9e, 0xb5, 0x39, 0x09,
	0x5d, 0x6f, 0xf3, 0xef, 0x65, 0xe3, 0xdb, 0x0a, 0x7a, 0x09, 0xab, 0xa3, 0x5c, 0x42, 0xf6, 0x09,
	0x2b, 0x3c, 0xbd, 0xad, 0xa9, 0xf5, 0x18, 0xfa, 0x16, 0xd6, 0x75, 0xba, 0xa0, 0xb2, 0x4d, 0xd2,
	0xce, 0x2b, 0xef, 0xfe, 0x8c, 0x9d, 0xb8, 0xc0, 0x21, 0xac, 0xa9, 0xd7, 0x8f, 0xb6, 0x6e, 0x09,
	0x15, 0xaf, 0x3c, 0xbd, 0x11, 0xa3, 0x8f, 0x21, 0x63, 0x5e, 0x31, 0xb2, 0xbb, 0x24, 0x93, 0xc1,
	0xf3, 0x66, 0x6d, 0xc5, 0x35, 0x1a, 0x90, 0x1d, 0x3f, 0x2e, 0xe4, 0xcd, 0x7c, 0x71, 0xba, 0xca,
	0xf6, 0x3f, 0x5e, 0xa3, 0x76, 0x42, 0x5f, 0x6c, 0xc2, 0x89, 0xc4, 0x40, 0x26, 0x9c, 0x48, 0x4e,
	0x01, 0x5e, 0x69, 0xad, 0xab, 0x69, 0x79, 0xf6, 0x27, 0x00, 0x00, 0xff, 0xff, 0x21, 0x49, 0xa8,
	0x3e, 0x8a, 0x08, 0x00, 0x00,
}
