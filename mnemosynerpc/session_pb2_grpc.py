# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc
from grpc.framework.common import cardinality
from grpc.framework.interfaces.face import utilities as face_utilities

import google.protobuf.empty_pb2 as google_dot_protobuf_dot_empty__pb2
import session_pb2 as session__pb2


class SessionManagerStub(object):

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.Context = channel.unary_unary(
        '/mnemosynerpc.SessionManager/Context',
        request_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
        response_deserializer=session__pb2.ContextResponse.FromString,
        )
    self.Get = channel.unary_unary(
        '/mnemosynerpc.SessionManager/Get',
        request_serializer=session__pb2.GetRequest.SerializeToString,
        response_deserializer=session__pb2.GetResponse.FromString,
        )
    self.List = channel.unary_unary(
        '/mnemosynerpc.SessionManager/List',
        request_serializer=session__pb2.ListRequest.SerializeToString,
        response_deserializer=session__pb2.ListResponse.FromString,
        )
    self.Exists = channel.unary_unary(
        '/mnemosynerpc.SessionManager/Exists',
        request_serializer=session__pb2.ExistsRequest.SerializeToString,
        response_deserializer=session__pb2.ExistsResponse.FromString,
        )
    self.Start = channel.unary_unary(
        '/mnemosynerpc.SessionManager/Start',
        request_serializer=session__pb2.StartRequest.SerializeToString,
        response_deserializer=session__pb2.StartResponse.FromString,
        )
    self.Abandon = channel.unary_unary(
        '/mnemosynerpc.SessionManager/Abandon',
        request_serializer=session__pb2.AbandonRequest.SerializeToString,
        response_deserializer=session__pb2.AbandonResponse.FromString,
        )
    self.SetValue = channel.unary_unary(
        '/mnemosynerpc.SessionManager/SetValue',
        request_serializer=session__pb2.SetValueRequest.SerializeToString,
        response_deserializer=session__pb2.SetValueResponse.FromString,
        )
    self.Delete = channel.unary_unary(
        '/mnemosynerpc.SessionManager/Delete',
        request_serializer=session__pb2.DeleteRequest.SerializeToString,
        response_deserializer=session__pb2.DeleteResponse.FromString,
        )


class SessionManagerServicer(object):

  def Context(self, request, context):
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def Get(self, request, context):
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def List(self, request, context):
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def Exists(self, request, context):
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def Start(self, request, context):
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def Abandon(self, request, context):
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def SetValue(self, request, context):
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def Delete(self, request, context):
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_SessionManagerServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'Context': grpc.unary_unary_rpc_method_handler(
          servicer.Context,
          request_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
          response_serializer=session__pb2.ContextResponse.SerializeToString,
      ),
      'Get': grpc.unary_unary_rpc_method_handler(
          servicer.Get,
          request_deserializer=session__pb2.GetRequest.FromString,
          response_serializer=session__pb2.GetResponse.SerializeToString,
      ),
      'List': grpc.unary_unary_rpc_method_handler(
          servicer.List,
          request_deserializer=session__pb2.ListRequest.FromString,
          response_serializer=session__pb2.ListResponse.SerializeToString,
      ),
      'Exists': grpc.unary_unary_rpc_method_handler(
          servicer.Exists,
          request_deserializer=session__pb2.ExistsRequest.FromString,
          response_serializer=session__pb2.ExistsResponse.SerializeToString,
      ),
      'Start': grpc.unary_unary_rpc_method_handler(
          servicer.Start,
          request_deserializer=session__pb2.StartRequest.FromString,
          response_serializer=session__pb2.StartResponse.SerializeToString,
      ),
      'Abandon': grpc.unary_unary_rpc_method_handler(
          servicer.Abandon,
          request_deserializer=session__pb2.AbandonRequest.FromString,
          response_serializer=session__pb2.AbandonResponse.SerializeToString,
      ),
      'SetValue': grpc.unary_unary_rpc_method_handler(
          servicer.SetValue,
          request_deserializer=session__pb2.SetValueRequest.FromString,
          response_serializer=session__pb2.SetValueResponse.SerializeToString,
      ),
      'Delete': grpc.unary_unary_rpc_method_handler(
          servicer.Delete,
          request_deserializer=session__pb2.DeleteRequest.FromString,
          response_serializer=session__pb2.DeleteResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'mnemosynerpc.SessionManager', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
