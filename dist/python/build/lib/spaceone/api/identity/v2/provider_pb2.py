# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/identity/v2/provider.proto
# Protobuf Python Version: 4.25.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from spaceone.api.core.v2 import query_pb2 as spaceone_dot_api_dot_core_dot_v2_dot_query__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\'spaceone/api/identity/v2/provider.proto\x12\x18spaceone.api.identity.v2\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\"\xd1\x01\n\nPluginInfo\x12\x11\n\tplugin_id\x18\x01 \x01(\t\x12\x0f\n\x07version\x18\x02 \x01(\t\x12\x46\n\x0cupgrade_mode\x18\x03 \x01(\x0e\x32\x30.spaceone.api.identity.v2.PluginInfo.UpgradeMode\x12(\n\x07options\x18\x04 \x01(\x0b\x32\x17.google.protobuf.Struct\"-\n\x0bUpgradeMode\x12\x08\n\x04NONE\x10\x00\x12\n\n\x06MANUAL\x10\x01\x12\x08\n\x04\x41UTO\x10\x02\"\xfe\x01\n\x15\x43reateProviderRequest\x12\x10\n\x08provider\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\r\n\x05\x61lias\x18\x03 \x01(\t\x12\x39\n\x0bplugin_info\x18\x04 \x01(\x0b\x32$.spaceone.api.identity.v2.PluginInfo\x12\r\n\x05\x63olor\x18\x05 \x01(\t\x12\x0c\n\x04icon\x18\x06 \x01(\t\x12\r\n\x05order\x18\x07 \x01(\x05\x12(\n\x07options\x18\x08 \x01(\x0b\x32\x17.google.protobuf.Struct\x12%\n\x04tags\x18\t \x01(\x0b\x32\x17.google.protobuf.Struct\"\xfe\x01\n\x15UpdateProviderRequest\x12\x10\n\x08provider\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\r\n\x05\x61lias\x18\x03 \x01(\t\x12\x39\n\x0bplugin_info\x18\x04 \x01(\x0b\x32$.spaceone.api.identity.v2.PluginInfo\x12\r\n\x05\x63olor\x18\x05 \x01(\t\x12\x0c\n\x04icon\x18\x06 \x01(\t\x12\r\n\x05order\x18\x07 \x01(\x05\x12(\n\x07options\x18\x08 \x01(\x0b\x32\x17.google.protobuf.Struct\x12%\n\x04tags\x18\t \x01(\x0b\x32\x17.google.protobuf.Struct\"\xf2\x01\n\x1bUpdatePluginProviderRequest\x12\x10\n\x08provider\x18\x01 \x01(\t\x12\x0f\n\x07version\x18\x02 \x01(\t\x12(\n\x07options\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12W\n\x0cupgrade_mode\x18\x04 \x01(\x0e\x32\x41.spaceone.api.identity.v2.UpdatePluginProviderRequest.UpgradeMode\"-\n\x0bUpgradeMode\x12\x08\n\x04NONE\x10\x00\x12\n\n\x06MANUAL\x10\x01\x12\x08\n\x04\x41UTO\x10\x02\"#\n\x0fProviderRequest\x12\x10\n\x08provider\x18\x01 \x01(\t\"\xc4\x02\n\x0cProviderInfo\x12\x10\n\x08provider\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\r\n\x05\x61lias\x18\x03 \x01(\t\x12\x39\n\x0bplugin_info\x18\x04 \x01(\x0b\x32$.spaceone.api.identity.v2.PluginInfo\x12\r\n\x05\x63olor\x18\x05 \x01(\t\x12\x0c\n\x04icon\x18\x06 \x01(\t\x12\r\n\x05order\x18\x07 \x01(\x05\x12(\n\x07options\x18\x08 \x01(\x0b\x32\x17.google.protobuf.Struct\x12%\n\x04tags\x18\t \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x12\n\nis_managed\x18\n \x01(\x08\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\x12\x12\n\nupdated_at\x18  \x01(\t\"\x84\x01\n\x13ProviderSearchQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v2.Query\x12\x10\n\x08provider\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\r\n\x05\x61lias\x18\x04 \x01(\t\x12\x12\n\nis_managed\x18\x05 \x01(\x08\"]\n\rProvidersInfo\x12\x37\n\x07results\x18\x01 \x03(\x0b\x32&.spaceone.api.identity.v2.ProviderInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"I\n\x11ProviderStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v2.StatisticsQuery2\xb8\x07\n\x08Provider\x12\x8a\x01\n\x06\x63reate\x12/.spaceone.api.identity.v2.CreateProviderRequest\x1a&.spaceone.api.identity.v2.ProviderInfo\"\'\x82\xd3\xe4\x93\x02!\"\x1c/identity/v2/provider/create:\x01*\x12\x8a\x01\n\x06update\x12/.spaceone.api.identity.v2.UpdateProviderRequest\x1a&.spaceone.api.identity.v2.ProviderInfo\"\'\x82\xd3\xe4\x93\x02!\"\x1c/identity/v2/provider/update:\x01*\x12\x9e\x01\n\rupdate_plugin\x12\x35.spaceone.api.identity.v2.UpdatePluginProviderRequest\x1a&.spaceone.api.identity.v2.ProviderInfo\".\x82\xd3\xe4\x93\x02(\"#/identity/v2/provider/update-plugin:\x01*\x12t\n\x06\x64\x65lete\x12).spaceone.api.identity.v2.ProviderRequest\x1a\x16.google.protobuf.Empty\"\'\x82\xd3\xe4\x93\x02!\"\x1c/identity/v2/provider/delete:\x01*\x12~\n\x03get\x12).spaceone.api.identity.v2.ProviderRequest\x1a&.spaceone.api.identity.v2.ProviderInfo\"$\x82\xd3\xe4\x93\x02\x1e\"\x19/identity/v2/provider/get:\x01*\x12\x85\x01\n\x04list\x12-.spaceone.api.identity.v2.ProviderSearchQuery\x1a\'.spaceone.api.identity.v2.ProvidersInfo\"%\x82\xd3\xe4\x93\x02\x1f\"\x1a/identity/v2/provider/list:\x01*\x12s\n\x04stat\x12+.spaceone.api.identity.v2.ProviderStatQuery\x1a\x17.google.protobuf.Struct\"%\x82\xd3\xe4\x93\x02\x1f\"\x1a/identity/v2/provider/stat:\x01*B?Z=github.com/cloudforet-io/api/dist/go/spaceone/api/identity/v2b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.identity.v2.provider_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z=github.com/cloudforet-io/api/dist/go/spaceone/api/identity/v2'
  _globals['_PROVIDER'].methods_by_name['create']._options = None
  _globals['_PROVIDER'].methods_by_name['create']._serialized_options = b'\202\323\344\223\002!\"\034/identity/v2/provider/create:\001*'
  _globals['_PROVIDER'].methods_by_name['update']._options = None
  _globals['_PROVIDER'].methods_by_name['update']._serialized_options = b'\202\323\344\223\002!\"\034/identity/v2/provider/update:\001*'
  _globals['_PROVIDER'].methods_by_name['update_plugin']._options = None
  _globals['_PROVIDER'].methods_by_name['update_plugin']._serialized_options = b'\202\323\344\223\002(\"#/identity/v2/provider/update-plugin:\001*'
  _globals['_PROVIDER'].methods_by_name['delete']._options = None
  _globals['_PROVIDER'].methods_by_name['delete']._serialized_options = b'\202\323\344\223\002!\"\034/identity/v2/provider/delete:\001*'
  _globals['_PROVIDER'].methods_by_name['get']._options = None
  _globals['_PROVIDER'].methods_by_name['get']._serialized_options = b'\202\323\344\223\002\036\"\031/identity/v2/provider/get:\001*'
  _globals['_PROVIDER'].methods_by_name['list']._options = None
  _globals['_PROVIDER'].methods_by_name['list']._serialized_options = b'\202\323\344\223\002\037\"\032/identity/v2/provider/list:\001*'
  _globals['_PROVIDER'].methods_by_name['stat']._options = None
  _globals['_PROVIDER'].methods_by_name['stat']._serialized_options = b'\202\323\344\223\002\037\"\032/identity/v2/provider/stat:\001*'
  _globals['_PLUGININFO']._serialized_start=193
  _globals['_PLUGININFO']._serialized_end=402
  _globals['_PLUGININFO_UPGRADEMODE']._serialized_start=357
  _globals['_PLUGININFO_UPGRADEMODE']._serialized_end=402
  _globals['_CREATEPROVIDERREQUEST']._serialized_start=405
  _globals['_CREATEPROVIDERREQUEST']._serialized_end=659
  _globals['_UPDATEPROVIDERREQUEST']._serialized_start=662
  _globals['_UPDATEPROVIDERREQUEST']._serialized_end=916
  _globals['_UPDATEPLUGINPROVIDERREQUEST']._serialized_start=919
  _globals['_UPDATEPLUGINPROVIDERREQUEST']._serialized_end=1161
  _globals['_UPDATEPLUGINPROVIDERREQUEST_UPGRADEMODE']._serialized_start=357
  _globals['_UPDATEPLUGINPROVIDERREQUEST_UPGRADEMODE']._serialized_end=402
  _globals['_PROVIDERREQUEST']._serialized_start=1163
  _globals['_PROVIDERREQUEST']._serialized_end=1198
  _globals['_PROVIDERINFO']._serialized_start=1201
  _globals['_PROVIDERINFO']._serialized_end=1525
  _globals['_PROVIDERSEARCHQUERY']._serialized_start=1528
  _globals['_PROVIDERSEARCHQUERY']._serialized_end=1660
  _globals['_PROVIDERSINFO']._serialized_start=1662
  _globals['_PROVIDERSINFO']._serialized_end=1755
  _globals['_PROVIDERSTATQUERY']._serialized_start=1757
  _globals['_PROVIDERSTATQUERY']._serialized_end=1830
  _globals['_PROVIDER']._serialized_start=1833
  _globals['_PROVIDER']._serialized_end=2785
# @@protoc_insertion_point(module_scope)
