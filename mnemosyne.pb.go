// Code generated by protoc-gen-go.
// source: mnemosyne.proto
// DO NOT EDIT!

package mnemosyne

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

// Ignoring public import of Timestamp from timestamp.proto

// Token represents identifier of single session. It consist of partition key and a hash.
type Token struct {
	Key  string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Hash string `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}

type Session struct {
	Token    *Token            `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	Data     map[string]string `protobuf:"bytes,2,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	ExpireAt *Timestamp        `protobuf:"bytes,3,opt,name=expire_at" json:"expire_at,omitempty"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}

func (m *Session) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *Session) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Session) GetExpireAt() *Timestamp {
	if m != nil {
		return m.ExpireAt
	}
	return nil
}

type GetRequest struct {
	Token *Token `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}

func (m *GetRequest) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type GetResponse struct {
	Session *Session `protobuf:"bytes,1,opt,name=session" json:"session,omitempty"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}

func (m *GetResponse) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

type ListRequest struct {
	Offset       int64      `protobuf:"varint,1,opt,name=offset" json:"offset,omitempty"`
	Limit        int64      `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
	ExpireAtFrom *Timestamp `protobuf:"bytes,3,opt,name=expire_at_from" json:"expire_at_from,omitempty"`
	ExpireAtTo   *Timestamp `protobuf:"bytes,4,opt,name=expire_at_to" json:"expire_at_to,omitempty"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}

func (m *ListRequest) GetExpireAtFrom() *Timestamp {
	if m != nil {
		return m.ExpireAtFrom
	}
	return nil
}

func (m *ListRequest) GetExpireAtTo() *Timestamp {
	if m != nil {
		return m.ExpireAtTo
	}
	return nil
}

type ListResponse struct {
	Sessions []*Session `protobuf:"bytes,1,rep,name=sessions" json:"sessions,omitempty"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}

func (m *ListResponse) GetSessions() []*Session {
	if m != nil {
		return m.Sessions
	}
	return nil
}

type ExistsRequest struct {
	Token *Token `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *ExistsRequest) Reset()         { *m = ExistsRequest{} }
func (m *ExistsRequest) String() string { return proto.CompactTextString(m) }
func (*ExistsRequest) ProtoMessage()    {}

func (m *ExistsRequest) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type ExistsResponse struct {
	Exists bool `protobuf:"varint,1,opt,name=exists" json:"exists,omitempty"`
}

func (m *ExistsResponse) Reset()         { *m = ExistsResponse{} }
func (m *ExistsResponse) String() string { return proto.CompactTextString(m) }
func (*ExistsResponse) ProtoMessage()    {}

type CreateRequest struct {
	Data map[string]string `protobuf:"bytes,1,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}

func (m *CreateRequest) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

type CreateResponse struct {
	Session *Session `protobuf:"bytes,1,opt,name=session" json:"session,omitempty"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}

func (m *CreateResponse) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

type AbandonRequest struct {
	Token *Token `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *AbandonRequest) Reset()         { *m = AbandonRequest{} }
func (m *AbandonRequest) String() string { return proto.CompactTextString(m) }
func (*AbandonRequest) ProtoMessage()    {}

func (m *AbandonRequest) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type AbandonResponse struct {
	Abandoned bool `protobuf:"varint,1,opt,name=abandoned" json:"abandoned,omitempty"`
}

func (m *AbandonResponse) Reset()         { *m = AbandonResponse{} }
func (m *AbandonResponse) String() string { return proto.CompactTextString(m) }
func (*AbandonResponse) ProtoMessage()    {}

type SetDataRequest struct {
	Token *Token `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	Key   string `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
}

func (m *SetDataRequest) Reset()         { *m = SetDataRequest{} }
func (m *SetDataRequest) String() string { return proto.CompactTextString(m) }
func (*SetDataRequest) ProtoMessage()    {}

func (m *SetDataRequest) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type SetDataResponse struct {
	Session *Session `protobuf:"bytes,1,opt,name=session" json:"session,omitempty"`
}

func (m *SetDataResponse) Reset()         { *m = SetDataResponse{} }
func (m *SetDataResponse) String() string { return proto.CompactTextString(m) }
func (*SetDataResponse) ProtoMessage()    {}

func (m *SetDataResponse) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

type DeleteRequest struct {
	Token        *Token     `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	ExpireAtFrom *Timestamp `protobuf:"bytes,2,opt,name=expire_at_from" json:"expire_at_from,omitempty"`
	ExpireAtTo   *Timestamp `protobuf:"bytes,3,opt,name=expire_at_to" json:"expire_at_to,omitempty"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}

func (m *DeleteRequest) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *DeleteRequest) GetExpireAtFrom() *Timestamp {
	if m != nil {
		return m.ExpireAtFrom
	}
	return nil
}

func (m *DeleteRequest) GetExpireAtTo() *Timestamp {
	if m != nil {
		return m.ExpireAtTo
	}
	return nil
}

type DeleteResponse struct {
	Count int64 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for RPC service

type RPCClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Exists(ctx context.Context, in *ExistsRequest, opts ...grpc.CallOption) (*ExistsResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Abandon(ctx context.Context, in *AbandonRequest, opts ...grpc.CallOption) (*AbandonResponse, error)
	SetData(ctx context.Context, in *SetDataRequest, opts ...grpc.CallOption) (*SetDataResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type rPCClient struct {
	cc *grpc.ClientConn
}

func NewRPCClient(cc *grpc.ClientConn) RPCClient {
	return &rPCClient{cc}
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

func (c *rPCClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := grpc.Invoke(ctx, "/mnemosyne.RPC/Create", in, out, c.cc, opts...)
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

func (c *rPCClient) SetData(ctx context.Context, in *SetDataRequest, opts ...grpc.CallOption) (*SetDataResponse, error) {
	out := new(SetDataResponse)
	err := grpc.Invoke(ctx, "/mnemosyne.RPC/SetData", in, out, c.cc, opts...)
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
	Get(context.Context, *GetRequest) (*GetResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	Exists(context.Context, *ExistsRequest) (*ExistsResponse, error)
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Abandon(context.Context, *AbandonRequest) (*AbandonResponse, error)
	SetData(context.Context, *SetDataRequest) (*SetDataResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
}

func RegisterRPCServer(s *grpc.Server, srv RPCServer) {
	s.RegisterService(&_RPC_serviceDesc, srv)
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

func _RPC_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(RPCServer).Create(ctx, in)
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

func _RPC_SetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(SetDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(RPCServer).SetData(ctx, in)
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
			MethodName: "Create",
			Handler:    _RPC_Create_Handler,
		},
		{
			MethodName: "Abandon",
			Handler:    _RPC_Abandon_Handler,
		},
		{
			MethodName: "SetData",
			Handler:    _RPC_SetData_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _RPC_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
