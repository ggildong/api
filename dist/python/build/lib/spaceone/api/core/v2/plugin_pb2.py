# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/core/v2/plugin.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n!spaceone/api/core/v2/plugin.proto\x12\x14spaceone.api.core.v2\x1a\x1cgoogle/protobuf/struct.proto\"\x91\x02\n\rPluginRequest\x12\x11\n\tplugin_id\x18\x01 \x01(\t\x12\x0f\n\x07version\x18\x02 \x01(\t\x12\x45\n\x0cupgrade_mode\x18\x03 \x01(\x0e\x32/.spaceone.api.core.v2.PluginRequest.UpgradeMode\x12(\n\x07options\x18\x04 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x0e\n\x06schema\x18\x05 \x01(\t\x12,\n\x0bsecret_data\x18\x06 \x01(\x0b\x32\x17.google.protobuf.Struct\"-\n\x0bUpgradeMode\x12\x08\n\x04NONE\x10\x00\x12\n\n\x06MANUAL\x10\x01\x12\x08\n\x04\x41UTO\x10\x02\"\x9b\x02\n\nPluginInfo\x12\x11\n\tplugin_id\x18\x01 \x01(\t\x12\x0f\n\x07version\x18\x02 \x01(\t\x12\x42\n\x0cupgrade_mode\x18\x03 \x01(\x0e\x32,.spaceone.api.core.v2.PluginInfo.UpgradeMode\x12(\n\x07options\x18\x04 \x01(\x0b\x32\x17.google.protobuf.Struct\x12)\n\x08metadata\x18\x05 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x0e\n\x06schema\x18\x06 \x01(\t\x12\x11\n\tsecret_id\x18\x07 \x01(\t\"-\n\x0bUpgradeMode\x12\x08\n\x04NONE\x10\x00\x12\n\n\x06MANUAL\x10\x01\x12\x08\n\x04\x41UTO\x10\x02\x42;Z9github.com/cloudforet-io/api/dist/go/spaceone/api/core/v2b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.core.v2.plugin_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z9github.com/cloudforet-io/api/dist/go/spaceone/api/core/v2'
  _globals['_PLUGINREQUEST']._serialized_start=90
  _globals['_PLUGINREQUEST']._serialized_end=363
  _globals['_PLUGINREQUEST_UPGRADEMODE']._serialized_start=318
  _globals['_PLUGINREQUEST_UPGRADEMODE']._serialized_end=363
  _globals['_PLUGININFO']._serialized_start=366
  _globals['_PLUGININFO']._serialized_end=649
  _globals['_PLUGININFO_UPGRADEMODE']._serialized_start=318
  _globals['_PLUGININFO_UPGRADEMODE']._serialized_end=363
# @@protoc_insertion_point(module_scope)
