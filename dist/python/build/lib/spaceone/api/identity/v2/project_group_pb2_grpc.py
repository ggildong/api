# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from spaceone.api.identity.v2 import project_group_pb2 as spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2


class ProjectGroupStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.create = channel.unary_unary(
                '/spaceone.api.identity.v2.ProjectGroup/create',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.CreateProjectGroupRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupInfo.FromString,
                )
        self.update = channel.unary_unary(
                '/spaceone.api.identity.v2.ProjectGroup/update',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.UpdateProjectGroupRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupInfo.FromString,
                )
        self.change_parent_group = channel.unary_unary(
                '/spaceone.api.identity.v2.ProjectGroup/change_parent_group',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ChangeParentGroupRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupInfo.FromString,
                )
        self.delete = channel.unary_unary(
                '/spaceone.api.identity.v2.ProjectGroup/delete',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )
        self.get = channel.unary_unary(
                '/spaceone.api.identity.v2.ProjectGroup/get',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupInfo.FromString,
                )
        self.list = channel.unary_unary(
                '/spaceone.api.identity.v2.ProjectGroup/list',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupSearchQuery.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupsInfo.FromString,
                )
        self.stat = channel.unary_unary(
                '/spaceone.api.identity.v2.ProjectGroup/stat',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupStatQuery.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_struct__pb2.Struct.FromString,
                )


class ProjectGroupServicer(object):
    """Missing associated documentation comment in .proto file."""

    def create(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def update(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def change_parent_group(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def delete(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def get(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def list(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def stat(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ProjectGroupServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'create': grpc.unary_unary_rpc_method_handler(
                    servicer.create,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.CreateProjectGroupRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupInfo.SerializeToString,
            ),
            'update': grpc.unary_unary_rpc_method_handler(
                    servicer.update,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.UpdateProjectGroupRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupInfo.SerializeToString,
            ),
            'change_parent_group': grpc.unary_unary_rpc_method_handler(
                    servicer.change_parent_group,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ChangeParentGroupRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupInfo.SerializeToString,
            ),
            'delete': grpc.unary_unary_rpc_method_handler(
                    servicer.delete,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'get': grpc.unary_unary_rpc_method_handler(
                    servicer.get,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupInfo.SerializeToString,
            ),
            'list': grpc.unary_unary_rpc_method_handler(
                    servicer.list,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupSearchQuery.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupsInfo.SerializeToString,
            ),
            'stat': grpc.unary_unary_rpc_method_handler(
                    servicer.stat,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupStatQuery.FromString,
                    response_serializer=google_dot_protobuf_dot_struct__pb2.Struct.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'spaceone.api.identity.v2.ProjectGroup', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class ProjectGroup(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def create(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.identity.v2.ProjectGroup/create',
            spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.CreateProjectGroupRequest.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def update(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.identity.v2.ProjectGroup/update',
            spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.UpdateProjectGroupRequest.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def change_parent_group(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.identity.v2.ProjectGroup/change_parent_group',
            spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ChangeParentGroupRequest.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def delete(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.identity.v2.ProjectGroup/delete',
            spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupRequest.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def get(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.identity.v2.ProjectGroup/get',
            spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupRequest.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def list(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.identity.v2.ProjectGroup/list',
            spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupSearchQuery.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupsInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def stat(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.identity.v2.ProjectGroup/stat',
            spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupStatQuery.SerializeToString,
            google_dot_protobuf_dot_struct__pb2.Struct.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
