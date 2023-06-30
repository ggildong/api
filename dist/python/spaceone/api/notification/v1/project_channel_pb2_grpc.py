# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from spaceone.api.notification.v1 import project_channel_pb2 as spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2


class ProjectChannelStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.create = channel.unary_unary(
                '/spaceone.api.notification.v1.ProjectChannel/create',
                request_serializer=spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.CreateProjectChannelRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.ProjectChannelInfo.FromString,
                )
        self.update = channel.unary_unary(
                '/spaceone.api.notification.v1.ProjectChannel/update',
                request_serializer=spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.UpdateProjectChannelRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.ProjectChannelInfo.FromString,
                )
        self.set_schedule = channel.unary_unary(
                '/spaceone.api.notification.v1.ProjectChannel/set_schedule',
                request_serializer=spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.UpdateProjectChannelScheduleRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.ProjectChannelInfo.FromString,
                )
        self.set_subscription = channel.unary_unary(
                '/spaceone.api.notification.v1.ProjectChannel/set_subscription',
                request_serializer=spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.UpdateProjectChannelSubscriptionRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.ProjectChannelInfo.FromString,
                )
        self.enable = channel.unary_unary(
                '/spaceone.api.notification.v1.ProjectChannel/enable',
                request_serializer=spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.ProjectChannelRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.ProjectChannelInfo.FromString,
                )
        self.disable = channel.unary_unary(
                '/spaceone.api.notification.v1.ProjectChannel/disable',
                request_serializer=spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.ProjectChannelRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.ProjectChannelInfo.FromString,
                )
        self.delete = channel.unary_unary(
                '/spaceone.api.notification.v1.ProjectChannel/delete',
                request_serializer=spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.ProjectChannelRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )
        self.get = channel.unary_unary(
                '/spaceone.api.notification.v1.ProjectChannel/get',
                request_serializer=spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.GetProjectChannelRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.ProjectChannelInfo.FromString,
                )
        self.list = channel.unary_unary(
                '/spaceone.api.notification.v1.ProjectChannel/list',
                request_serializer=spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.ProjectChannelQuery.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.ProjectChannelsInfo.FromString,
                )
        self.stat = channel.unary_unary(
                '/spaceone.api.notification.v1.ProjectChannel/stat',
                request_serializer=spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.ProjectChannelStatQuery.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_struct__pb2.Struct.FromString,
                )


class ProjectChannelServicer(object):
    """Missing associated documentation comment in .proto file."""

    def create(self, request, context):
        """Creates a new ProjectChannel. ProjectChannel is a channel that delivers a Notification to the Project when the Notification is created. When creating a ProjectChannel, one Protocol must be selected, and a Notification is dispatched via the selected Protocol.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def update(self, request, context):
        """Updates a specific ProjectChannel. A ProjectChannel that has already been configured cannot be changed. Instead, the data required for dispatching Notifications to a ProjectChannel can be updated.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def set_schedule(self, request, context):
        """Sets a schedule for a ProjectChannel. A schedule defines the time to receive a Notification. When a Notification is created, you can set the day of the week and time you want to receive it. When you set the day of the week, you can receive a notification only on the day of the week you set. If you also set the start time and end time with the day of the week, you can receive a Notification only at the scheduled time on the day of the week you set. If there is no schedule set in a ProjectChannel, Notifications will be dispatched at all times via the ProjectChannel.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def set_subscription(self, request, context):
        """Sets a subscription for a ProjectChannel. A subscription is a topic for channels to subscribe to and get notified about. If a ProjectChannel has subscriptions, a Notification is dispatched only if the topic of the Notification is the same as the one set in the subscriptions. If there is no subscription in a ProjectChannel, all Notifications will be dispatched.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def enable(self, request, context):
        """Enables a specific ProjectChannel. If a ProjectChannel is enabled, the ProjectChannel can be used and the Notification can be dispatched. Even if a ProjectChannel is enabled, if the Protocol used in the ProjectChannel is disabled, the Notification is not dispatched.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def disable(self, request, context):
        """Disables a specific ProjectChannel. If a ProjectChannel is disabled, the Notification is not dispatched, even if it is created.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def delete(self, request, context):
        """Deletes a specific ProjectChannel.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def get(self, request, context):
        """Gets a specific ProjectChannel. Prints detailed information about the ProjectChannel.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def list(self, request, context):
        """Gets a list of all ProjectChannels. You can use a query to get a filtered list of ProjectChannels.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def stat(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ProjectChannelServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'create': grpc.unary_unary_rpc_method_handler(
                    servicer.create,
                    request_deserializer=spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.CreateProjectChannelRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.ProjectChannelInfo.SerializeToString,
            ),
            'update': grpc.unary_unary_rpc_method_handler(
                    servicer.update,
                    request_deserializer=spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.UpdateProjectChannelRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.ProjectChannelInfo.SerializeToString,
            ),
            'set_schedule': grpc.unary_unary_rpc_method_handler(
                    servicer.set_schedule,
                    request_deserializer=spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.UpdateProjectChannelScheduleRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.ProjectChannelInfo.SerializeToString,
            ),
            'set_subscription': grpc.unary_unary_rpc_method_handler(
                    servicer.set_subscription,
                    request_deserializer=spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.UpdateProjectChannelSubscriptionRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.ProjectChannelInfo.SerializeToString,
            ),
            'enable': grpc.unary_unary_rpc_method_handler(
                    servicer.enable,
                    request_deserializer=spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.ProjectChannelRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.ProjectChannelInfo.SerializeToString,
            ),
            'disable': grpc.unary_unary_rpc_method_handler(
                    servicer.disable,
                    request_deserializer=spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.ProjectChannelRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.ProjectChannelInfo.SerializeToString,
            ),
            'delete': grpc.unary_unary_rpc_method_handler(
                    servicer.delete,
                    request_deserializer=spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.ProjectChannelRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'get': grpc.unary_unary_rpc_method_handler(
                    servicer.get,
                    request_deserializer=spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.GetProjectChannelRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.ProjectChannelInfo.SerializeToString,
            ),
            'list': grpc.unary_unary_rpc_method_handler(
                    servicer.list,
                    request_deserializer=spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.ProjectChannelQuery.FromString,
                    response_serializer=spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.ProjectChannelsInfo.SerializeToString,
            ),
            'stat': grpc.unary_unary_rpc_method_handler(
                    servicer.stat,
                    request_deserializer=spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.ProjectChannelStatQuery.FromString,
                    response_serializer=google_dot_protobuf_dot_struct__pb2.Struct.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'spaceone.api.notification.v1.ProjectChannel', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class ProjectChannel(object):
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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.notification.v1.ProjectChannel/create',
            spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.CreateProjectChannelRequest.SerializeToString,
            spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.ProjectChannelInfo.FromString,
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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.notification.v1.ProjectChannel/update',
            spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.UpdateProjectChannelRequest.SerializeToString,
            spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.ProjectChannelInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def set_schedule(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.notification.v1.ProjectChannel/set_schedule',
            spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.UpdateProjectChannelScheduleRequest.SerializeToString,
            spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.ProjectChannelInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def set_subscription(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.notification.v1.ProjectChannel/set_subscription',
            spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.UpdateProjectChannelSubscriptionRequest.SerializeToString,
            spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.ProjectChannelInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def enable(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.notification.v1.ProjectChannel/enable',
            spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.ProjectChannelRequest.SerializeToString,
            spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.ProjectChannelInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def disable(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.notification.v1.ProjectChannel/disable',
            spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.ProjectChannelRequest.SerializeToString,
            spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.ProjectChannelInfo.FromString,
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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.notification.v1.ProjectChannel/delete',
            spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.ProjectChannelRequest.SerializeToString,
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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.notification.v1.ProjectChannel/get',
            spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.GetProjectChannelRequest.SerializeToString,
            spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.ProjectChannelInfo.FromString,
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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.notification.v1.ProjectChannel/list',
            spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.ProjectChannelQuery.SerializeToString,
            spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.ProjectChannelsInfo.FromString,
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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.notification.v1.ProjectChannel/stat',
            spaceone_dot_api_dot_notification_dot_v1_dot_project__channel__pb2.ProjectChannelStatQuery.SerializeToString,
            google_dot_protobuf_dot_struct__pb2.Struct.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)