package main

import (
	"github.com/go-kit/kit/log"
	"github.com/piotrkowalczuk/mnemosyne"
	"github.com/piotrkowalczuk/sklog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

type rpcServer struct {
	logger  log.Logger
	monitor *monitoring
	storage Storage
	alloc   struct {
		abandon  handlerFunc
		context  handlerFunc
		delete   handlerFunc
		exists   handlerFunc
		get      handlerFunc
		list     handlerFunc
		setValue handlerFunc
		start    handlerFunc
	}
}

func newRPCServer(l log.Logger, s Storage, m *monitoring) *rpcServer {
	return &rpcServer{
		alloc: struct {
			abandon  handlerFunc
			context  handlerFunc
			delete   handlerFunc
			exists   handlerFunc
			get      handlerFunc
			list     handlerFunc
			setValue handlerFunc
			start    handlerFunc
		}{
			abandon:  newHandlerFunc("abandon"),
			context:  newHandlerFunc("context"),
			delete:   newHandlerFunc("delete"),
			exists:   newHandlerFunc("exists"),
			get:      newHandlerFunc("get"),
			list:     newHandlerFunc("list"),
			setValue: newHandlerFunc("set_value"),
			start:    newHandlerFunc("start"),
		},
		logger:  l,
		storage: s,
		monitor: m,
	}
}

// Get implements mnemosyne.RPCServer interface.
func (rs *rpcServer) Context(ctx context.Context, req *mnemosyne.Empty) (*mnemosyne.Session, error) {
	h := rs.alloc.context(rs.loggerBackground(ctx), rs.storage, rs.monitor.rpc)
	h.monitor.requests.Add(1)

	ses, err := h.context(ctx)
	if err != nil {
		h.monitor.errors.Add(1)
		sklog.Error(h.logger, err)

		return nil, rs.error(err)
	}

	sklog.Debug(h.logger, "session has been retrieved (by context)")

	return ses, nil
}

// Get implements mnemosyne.RPCServer interface.
func (rs *rpcServer) Get(ctx context.Context, req *mnemosyne.GetRequest) (*mnemosyne.GetResponse, error) {
	h := rs.alloc.get(rs.loggerBackground(ctx), rs.storage, rs.monitor.rpc)
	h.monitor.requests.Add(1)

	ses, err := h.get(ctx, req)
	if err != nil {
		h.monitor.errors.Add(1)
		sklog.Error(h.logger, err)

		return nil, rs.error(err)
	}

	sklog.Debug(h.logger, "session has been retrieved (by token)")

	return &mnemosyne.GetResponse{
		Session: ses,
	}, nil
}

// List implements mnemosyne.RPCServer interface.
func (rs *rpcServer) List(ctx context.Context, req *mnemosyne.ListRequest) (*mnemosyne.ListResponse, error) {
	h := rs.alloc.list(rs.loggerBackground(ctx), rs.storage, rs.monitor.rpc)
	h.monitor.requests.Add(1)

	sessions, err := h.list(ctx, req)
	if err != nil {
		h.monitor.errors.Add(1)
		sklog.Error(h.logger, err)

		return nil, rs.error(err)
	}

	sklog.Debug(h.logger, "session list has been retrieved")

	return &mnemosyne.ListResponse{
		Sessions: sessions,
	}, nil
}

// Start implements mnemosyne.RPCServer interface.
func (rs *rpcServer) Start(ctx context.Context, req *mnemosyne.StartRequest) (*mnemosyne.StartResponse, error) {
	h := rs.alloc.start(rs.loggerBackground(ctx), rs.storage, rs.monitor.rpc)
	h.monitor.requests.Add(1)

	ses, err := h.start(ctx, req)
	if err != nil {
		h.monitor.errors.Add(1)
		sklog.Error(h.logger, err)

		return nil, rs.error(err)
	}

	sklog.Debug(h.logger, "session has been started")

	return &mnemosyne.StartResponse{
		Session: ses,
	}, nil
}

// Exists implements mnemosyne.RPCServer interface.
func (rs *rpcServer) Exists(ctx context.Context, req *mnemosyne.ExistsRequest) (*mnemosyne.ExistsResponse, error) {
	h := rs.alloc.exists(rs.loggerBackground(ctx), rs.storage, rs.monitor.rpc)
	h.monitor.requests.Add(1)

	exists, err := h.exists(ctx, req)
	if err != nil {
		h.monitor.errors.Add(1)
		sklog.Error(h.logger, err)

		return nil, rs.error(err)
	}

	sklog.Debug(h.logger, "session presence has been checked")

	return &mnemosyne.ExistsResponse{
		Exists: exists,
	}, nil
}

// Abandon implements mnemosyne.RPCServer interface.
func (rs *rpcServer) Abandon(ctx context.Context, req *mnemosyne.AbandonRequest) (*mnemosyne.AbandonResponse, error) {
	h := rs.alloc.abandon(rs.loggerBackground(ctx), rs.storage, rs.monitor.rpc)
	h.monitor.requests.Add(1)

	abandoned, err := h.abandon(ctx, req)
	if err != nil {
		h.monitor.errors.Add(1)
		sklog.Error(h.logger, err)

		return nil, rs.error(err)
	}

	sklog.Debug(h.logger, "session has been abandoned")

	return &mnemosyne.AbandonResponse{
		Abandoned: abandoned,
	}, nil
}

// SetValue implements mnemosyne.RPCServer interface.
func (rs *rpcServer) SetValue(ctx context.Context, req *mnemosyne.SetValueRequest) (*mnemosyne.SetValueResponse, error) {
	h := rs.alloc.setValue(rs.loggerBackground(ctx), rs.storage, rs.monitor.rpc)
	h.monitor.requests.Add(1)

	bag, err := h.setValue(ctx, req)
	if err != nil {
		h.monitor.errors.Add(1)
		sklog.Error(h.logger, err)

		return nil, rs.error(err)
	}

	sklog.Debug(h.logger, "session bag value has been set")

	return &mnemosyne.SetValueResponse{
		Bag: bag,
	}, nil
}

// Delete implements mnemosyne.RPCServer interface.
func (rs *rpcServer) Delete(ctx context.Context, req *mnemosyne.DeleteRequest) (*mnemosyne.DeleteResponse, error) {
	h := rs.alloc.delete(rs.loggerBackground(ctx), rs.storage, rs.monitor.rpc)
	h.monitor.requests.Add(1)

	affected, err := h.delete(ctx, req)
	if err != nil {
		h.monitor.errors.Add(1)
		sklog.Error(h.logger, err)

		return nil, rs.error(err)
	}

	sklog.Debug(h.logger, "session value has been deleted")

	return &mnemosyne.DeleteResponse{
		Count: affected,
	}, nil
}

func (rs *rpcServer) error(err error) error {
	if err == nil {
		return nil
	}

	switch err {
	case errSessionNotFound:
		return mnemosyne.ErrSessionNotFound
	default:
		return grpc.Errorf(codes.Internal, err.Error())
	}
}

func (rs *rpcServer) loggerBackground(ctx context.Context, keyval ...interface{}) log.Logger {
	l := log.NewContext(rs.logger).With(keyval...)
	if md, ok := metadata.FromContext(ctx); ok {
		if rid, ok := md["request_id"]; ok && len(rid) >= 1 {
			l = l.With("request_id", rid[0])
		}
	}

	if p, ok := peer.FromContext(ctx); ok {
		l = l.With("peer_address", p.Addr.String())
	}

	return l
}
