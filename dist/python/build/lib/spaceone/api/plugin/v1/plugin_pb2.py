# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/plugin/v1/plugin.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n#spaceone/api/plugin/v1/plugin.proto\x12\x16spaceone.api.plugin.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\"\xed\x01\n\x15PluginEndpointRequest\x12\x11\n\tplugin_id\x18\x01 \x01(\t\x12\x0f\n\x07version\x18\x02 \x01(\t\x12\'\n\x06labels\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x04 \x01(\t\x12O\n\x0cupgrade_mode\x18\x05 \x01(\x0e\x32\x39.spaceone.api.plugin.v1.PluginEndpointRequest.UpgradeMode\"#\n\x0bUpgradeMode\x12\n\n\x06MANUAL\x10\x00\x12\x08\n\x04\x41UTO\x10\x01\"\xee\x01\n\x15PluginMetadataRequest\x12\x11\n\tplugin_id\x18\x01 \x01(\t\x12\x0f\n\x07version\x18\x02 \x01(\t\x12(\n\x07options\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x04 \x01(\t\x12O\n\x0cupgrade_mode\x18\x05 \x01(\x0e\x32\x39.spaceone.api.plugin.v1.PluginMetadataRequest.UpgradeMode\"#\n\x0bUpgradeMode\x12\n\n\x06MANUAL\x10\x00\x12\x08\n\x04\x41UTO\x10\x01\"d\n\x14PluginFailureRequest\x12\x15\n\rsupervisor_id\x18\x01 \x01(\t\x12\x11\n\tplugin_id\x18\x02 \x01(\t\x12\x0f\n\x07version\x18\x03 \x01(\t\x12\x11\n\tdomain_id\x18\x04 \x01(\t\"Q\n\x0ePluginEndpoint\x12\x10\n\x08\x65ndpoint\x18\x01 \x01(\t\x12\x14\n\x0c\x61\x63\x63\x65ss_token\x18\x02 \x01(\t\x12\x17\n\x0fupdated_version\x18\x03 \x01(\t\";\n\x0ePluginMetadata\x12)\n\x08metadata\x18\x01 \x01(\x0b\x32\x17.google.protobuf.Struct2\xc9\x03\n\x06Plugin\x12\x9e\x01\n\x13get_plugin_endpoint\x12-.spaceone.api.plugin.v1.PluginEndpointRequest\x1a&.spaceone.api.plugin.v1.PluginEndpoint\"0\x82\xd3\xe4\x93\x02*\"%/plugin/v1/plugin/get-plugin-endpoint:\x01*\x12\x97\x01\n\x13get_plugin_metadata\x12-.spaceone.api.plugin.v1.PluginMetadataRequest\x1a&.spaceone.api.plugin.v1.PluginMetadata\")\x82\xd3\xe4\x93\x02#\"\x1e/plugin/v1/plugin/get-metadata:\x01*\x12\x83\x01\n\x0enotify_failure\x12,.spaceone.api.plugin.v1.PluginFailureRequest\x1a\x16.google.protobuf.Empty\"+\x82\xd3\xe4\x93\x02%\" /plugin/v1/plugin/notify-failure:\x01*B=Z;github.com/cloudforet-io/api/dist/go/spaceone/api/plugin/v1b\x06proto3')



_PLUGINENDPOINTREQUEST = DESCRIPTOR.message_types_by_name['PluginEndpointRequest']
_PLUGINMETADATAREQUEST = DESCRIPTOR.message_types_by_name['PluginMetadataRequest']
_PLUGINFAILUREREQUEST = DESCRIPTOR.message_types_by_name['PluginFailureRequest']
_PLUGINENDPOINT = DESCRIPTOR.message_types_by_name['PluginEndpoint']
_PLUGINMETADATA = DESCRIPTOR.message_types_by_name['PluginMetadata']
_PLUGINENDPOINTREQUEST_UPGRADEMODE = _PLUGINENDPOINTREQUEST.enum_types_by_name['UpgradeMode']
_PLUGINMETADATAREQUEST_UPGRADEMODE = _PLUGINMETADATAREQUEST.enum_types_by_name['UpgradeMode']
PluginEndpointRequest = _reflection.GeneratedProtocolMessageType('PluginEndpointRequest', (_message.Message,), {
  'DESCRIPTOR' : _PLUGINENDPOINTREQUEST,
  '__module__' : 'spaceone.api.plugin.v1.plugin_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.plugin.v1.PluginEndpointRequest)
  })
_sym_db.RegisterMessage(PluginEndpointRequest)

PluginMetadataRequest = _reflection.GeneratedProtocolMessageType('PluginMetadataRequest', (_message.Message,), {
  'DESCRIPTOR' : _PLUGINMETADATAREQUEST,
  '__module__' : 'spaceone.api.plugin.v1.plugin_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.plugin.v1.PluginMetadataRequest)
  })
_sym_db.RegisterMessage(PluginMetadataRequest)

PluginFailureRequest = _reflection.GeneratedProtocolMessageType('PluginFailureRequest', (_message.Message,), {
  'DESCRIPTOR' : _PLUGINFAILUREREQUEST,
  '__module__' : 'spaceone.api.plugin.v1.plugin_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.plugin.v1.PluginFailureRequest)
  })
_sym_db.RegisterMessage(PluginFailureRequest)

PluginEndpoint = _reflection.GeneratedProtocolMessageType('PluginEndpoint', (_message.Message,), {
  'DESCRIPTOR' : _PLUGINENDPOINT,
  '__module__' : 'spaceone.api.plugin.v1.plugin_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.plugin.v1.PluginEndpoint)
  })
_sym_db.RegisterMessage(PluginEndpoint)

PluginMetadata = _reflection.GeneratedProtocolMessageType('PluginMetadata', (_message.Message,), {
  'DESCRIPTOR' : _PLUGINMETADATA,
  '__module__' : 'spaceone.api.plugin.v1.plugin_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.plugin.v1.PluginMetadata)
  })
_sym_db.RegisterMessage(PluginMetadata)

_PLUGIN = DESCRIPTOR.services_by_name['Plugin']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z;github.com/cloudforet-io/api/dist/go/spaceone/api/plugin/v1'
  _PLUGIN.methods_by_name['get_plugin_endpoint']._options = None
  _PLUGIN.methods_by_name['get_plugin_endpoint']._serialized_options = b'\202\323\344\223\002*\"%/plugin/v1/plugin/get-plugin-endpoint:\001*'
  _PLUGIN.methods_by_name['get_plugin_metadata']._options = None
  _PLUGIN.methods_by_name['get_plugin_metadata']._serialized_options = b'\202\323\344\223\002#\"\036/plugin/v1/plugin/get-metadata:\001*'
  _PLUGIN.methods_by_name['notify_failure']._options = None
  _PLUGIN.methods_by_name['notify_failure']._serialized_options = b'\202\323\344\223\002%\" /plugin/v1/plugin/notify-failure:\001*'
  _PLUGINENDPOINTREQUEST._serialized_start=153
  _PLUGINENDPOINTREQUEST._serialized_end=390
  _PLUGINENDPOINTREQUEST_UPGRADEMODE._serialized_start=355
  _PLUGINENDPOINTREQUEST_UPGRADEMODE._serialized_end=390
  _PLUGINMETADATAREQUEST._serialized_start=393
  _PLUGINMETADATAREQUEST._serialized_end=631
  _PLUGINMETADATAREQUEST_UPGRADEMODE._serialized_start=355
  _PLUGINMETADATAREQUEST_UPGRADEMODE._serialized_end=390
  _PLUGINFAILUREREQUEST._serialized_start=633
  _PLUGINFAILUREREQUEST._serialized_end=733
  _PLUGINENDPOINT._serialized_start=735
  _PLUGINENDPOINT._serialized_end=816
  _PLUGINMETADATA._serialized_start=818
  _PLUGINMETADATA._serialized_end=877
  _PLUGIN._serialized_start=880
  _PLUGIN._serialized_end=1337
# @@protoc_insertion_point(module_scope)