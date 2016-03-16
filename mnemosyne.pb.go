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
import protot "github.com/piotrkowalczuk/protot"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

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
	AccessToken *AccessToken      `protobuf:"bytes,1,opt,name=access_token" json:"access_token,omitempty"`
	SubjectId   string            `protobuf:"bytes,2,opt,name=subject_id" json:"subject_id,omitempty"`
	Bag         map[string]string `protobuf:"bytes,3,rep,name=bag" json:"bag,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	ExpireAt    *protot.Timestamp `protobuf:"bytes,4,opt,name=expire_at" json:"expire_at,omitempty"`
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

func (m *Session) GetExpireAt() *protot.Timestamp {
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
	Offset       int64             `protobuf:"varint,1,opt,name=offset" json:"offset,omitempty"`
	Limit        int64             `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
	ExpireAtFrom *protot.Timestamp `protobuf:"bytes,3,opt,name=expire_at_from" json:"expire_at_from,omitempty"`
	ExpireAtTo   *protot.Timestamp `protobuf:"bytes,4,opt,name=expire_at_to" json:"expire_at_to,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ListRequest) GetExpireAtFrom() *protot.Timestamp {
	if m != nil {
		return m.ExpireAtFrom
	}
	return nil
}

func (m *ListRequest) GetExpireAtTo() *protot.Timestamp {
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
	AccessToken  *AccessToken      `protobuf:"bytes,1,opt,name=access_token" json:"access_token,omitempty"`
	ExpireAtFrom *protot.Timestamp `protobuf:"bytes,2,opt,name=expire_at_from" json:"expire_at_from,omitempty"`
	ExpireAtTo   *protot.Timestamp `protobuf:"bytes,3,opt,name=expire_at_to" json:"expire_at_to,omitempty"`
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

func (m *DeleteRequest) GetExpireAtFrom() *protot.Timestamp {
	if m != nil {
		return m.ExpireAtFrom
	}
	return nil
}

func (m *DeleteRequest) GetExpireAtTo() *protot.Timestamp {
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
	// 676 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x55, 0x5f, 0x6f, 0xd3, 0x3e,
	0x14, 0x5d, 0x9a, 0x6d, 0x6d, 0x6e, 0xd3, 0x76, 0xf3, 0x4f, 0xbf, 0x2d, 0x64, 0x0f, 0x54, 0xa6,
	0x12, 0x03, 0x41, 0x25, 0xc6, 0x84, 0x60, 0x1a, 0xa0, 0x6d, 0x54, 0xbc, 0x20, 0x81, 0xb6, 0x89,
	0x17, 0x1e, 0xaa, 0xb4, 0x75, 0xb7, 0xb0, 0x26, 0xee, 0x62, 0x17, 0xb5, 0x0f, 0x48, 0x7c, 0x05,
	0x3e, 0x15, 0xdf, 0x88, 0x67, 0x12, 0xdb, 0x4d, 0x1d, 0xda, 0x22, 0x1a, 0xf1, 0x18, 0xdb, 0xe7,
	0xde, 0x73, 0xce, 0xfd, 0x13, 0xa8, 0x05, 0x21, 0x09, 0x28, 0x9b, 0x84, 0xa4, 0x39, 0x8c, 0x28,
	0xa7, 0xc8, 0x4a, 0x0f, 0x5c, 0x5b, 0x9c, 0x70, 0x79, 0x81, 0x8b, 0xb0, 0xd1, 0x0a, 0x86, 0x7c,
	0x82, 0xf7, 0xa1, 0x7c, 0xd2, 0xed, 0x12, 0xc6, 0x2e, 0xe9, 0x0d, 0x09, 0x51, 0x19, 0xcc, 0x1b,
	0x32, 0x71, 0x8c, 0xba, 0xb1, 0x6f, 0x23, 0x1b, 0xd6, 0xaf, 0x3d, 0x76, 0xed, 0x14, 0x92, 0x2f,
	0xfc, 0xc3, 0x80, 0xe2, 0x45, 0xfc, 0xd0, 0xa7, 0x21, 0x7a, 0x04, 0xb6, 0x27, 0x50, 0x6d, 0x9e,
	0xc0, 0xc4, 0xfb, 0xf2, 0xc1, 0x4e, 0x73, 0x96, 0x5f, 0x0f, 0x8a, 0x00, 0xd8, 0xa8, 0xf3, 0x99,
	0x74, 0x79, 0xdb, 0xef, 0x89, 0x68, 0x16, 0xda, 0x07, 0xb3, 0xe3, 0x5d, 0x39, 0x66, 0xdd, 0x8c,
	0x81, 0x7b, 0x1a, 0x50, 0xa5, 0x68, 0x9e, 0x7a, 0x57, 0xad, 0x90, 0x47, 0x13, 0xd4, 0x00, 0x8b,
	0x8c, 0x87, 0x7e, 0x44, 0xda, 0x1e, 0x77, 0xd6, 0x45, 0xa2, 0xed, 0xa6, 0x12, 0x73, 0xe9, 0x07,
	0x84, 0x71, 0x2f, 0x18, 0xba, 0x0f, 0xa1, 0x94, 0x22, 0x34, 0x11, 0x16, 0xaa, 0xc0, 0xc6, 0x17,
	0x6f, 0x30, 0x22, 0x32, 0xef, 0x51, 0xe1, 0xb9, 0x81, 0x8f, 0x00, 0xde, 0x12, 0x7e, 0x4e, 0x6e,
	0x47, 0x31, 0x78, 0x35, 0x2d, 0xf8, 0x00, 0xca, 0x02, 0xcb, 0x86, 0x34, 0x64, 0x04, 0xdd, 0x83,
	0x22, 0x93, 0x84, 0x15, 0x0e, 0xcd, 0x4b, 0xc1, 0xdf, 0x0c, 0x28, 0xbf, 0xf3, 0x59, 0x9a, 0xb1,
	0x0a, 0x9b, 0xb4, 0xdf, 0x67, 0x84, 0x0b, 0x8c, 0x99, 0x50, 0x1c, 0xf8, 0x81, 0xcf, 0x05, 0x45,
	0x13, 0x3d, 0x80, 0x6a, 0x2a, 0xb8, 0xdd, 0x8f, 0x68, 0x10, 0xbb, 0xb4, 0x58, 0x35, 0xba, 0x0f,
	0xf6, 0xec, 0x29, 0xa7, 0x4b, 0xed, 0xc1, 0x87, 0x60, 0x4b, 0x06, 0x8a, 0x77, 0x03, 0x4a, 0x8a,
	0x37, 0x8b, 0x49, 0x98, 0x4b, 0x88, 0xbf, 0x84, 0x4a, 0x6b, 0x1c, 0xc3, 0x58, 0x3e, 0xaf, 0xea,
	0x50, 0x9d, 0xc2, 0x55, 0xda, 0x58, 0x39, 0x11, 0x27, 0x02, 0x59, 0xc2, 0x5f, 0xc1, 0xbe, 0xe0,
	0x5e, 0x94, 0x3a, 0x93, 0xed, 0x14, 0x59, 0xc0, 0xc7, 0xb2, 0x53, 0x0a, 0x82, 0x65, 0x5d, 0x67,
	0xa9, 0x21, 0xd3, 0x76, 0x59, 0xa9, 0x11, 0x0e, 0xa1, 0xa2, 0x82, 0xac, 0x52, 0xce, 0x57, 0x50,
	0x3d, 0xe9, 0x78, 0x61, 0x8f, 0x86, 0xf9, 0x6c, 0x69, 0x40, 0x2d, 0xc5, 0xab, 0xbc, 0xdb, 0x60,
	0x79, 0xf2, 0x88, 0xf4, 0x94, 0x35, 0x9f, 0xa0, 0x76, 0x41, 0xf8, 0xc7, 0x84, 0x71, 0xae, 0x34,
	0x53, 0xf1, 0x85, 0xac, 0xf8, 0xa4, 0x95, 0x2c, 0x7c, 0x0b, 0x5b, 0xb3, 0xe0, 0x8a, 0xc3, 0x13,
	0xe9, 0xb3, 0xec, 0x86, 0x46, 0x46, 0x77, 0xf6, 0x65, 0x3e, 0xaf, 0xdf, 0x03, 0x7a, 0x43, 0x06,
	0x84, 0x93, 0x7f, 0x24, 0x29, 0x9e, 0xe2, 0xff, 0x32, 0x01, 0x57, 0x29, 0xe1, 0x31, 0xd8, 0x67,
	0x03, 0xe2, 0x45, 0xf9, 0x0a, 0x58, 0x83, 0x8a, 0x42, 0xcb, 0x9c, 0xf8, 0xbb, 0x01, 0x15, 0xc9,
	0x25, 0x9f, 0xae, 0xf9, 0x89, 0x2f, 0xfc, 0xed, 0xc4, 0x2f, 0x5b, 0x0d, 0xf8, 0x2e, 0x54, 0xa7,
	0x94, 0x94, 0x33, 0x71, 0x51, 0xba, 0x74, 0x14, 0xaa, 0xad, 0x73, 0xf0, 0xd3, 0x04, 0xf3, 0xfc,
	0xc3, 0x59, 0x5c, 0xf7, 0xe2, 0x19, 0x0d, 0x39, 0x19, 0x73, 0xb4, 0xa5, 0xf1, 0x13, 0xbf, 0x07,
	0x77, 0x91, 0x79, 0x6b, 0xe8, 0x19, 0x98, 0xf1, 0x12, 0x44, 0xff, 0x6b, 0x97, 0xb3, 0x85, 0xea,
	0xee, 0xfc, 0x7e, 0xac, 0x5c, 0x5a, 0x43, 0x2f, 0x60, 0x3d, 0xd9, 0x42, 0x48, 0x7f, 0xa1, 0x2d,
	0x46, 0x77, 0x77, 0xee, 0x3c, 0x85, 0xbe, 0x86, 0x4d, 0xb9, 0x4b, 0x90, 0xa3, 0x93, 0xd4, 0xb7,
	0x93, 0x7b, 0x67, 0xc1, 0x4d, 0x1a, 0xe0, 0x18, 0x36, 0xc4, 0xac, 0xa3, 0xdd, 0x25, 0x2b, 0xc4,
	0x75, 0xe6, 0x2f, 0x52, 0xf4, 0x29, 0x14, 0xd5, 0xcc, 0x22, 0x3d, 0x4b, 0x76, 0x0f, 0xb8, 0xee,
	0xa2, 0xab, 0x34, 0x46, 0x0b, 0x4a, 0xd3, 0x51, 0x42, 0xee, 0xc2, 0xf9, 0x92, 0x51, 0xf6, 0xfe,
	0x30, 0x7b, 0xd2, 0x09, 0x59, 0xd8, 0x8c, 0x13, 0x99, 0xf6, 0xcb, 0x38, 0x91, 0xed, 0x02, 0xbc,
	0xd6, 0xd9, 0x14, 0xbd, 0xf2, 0xf4, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xea, 0x4d, 0x70, 0x9e,
	0x2e, 0x08, 0x00, 0x00,
}
