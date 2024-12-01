# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/alert_manager/v1/alert.proto
# Protobuf Python Version: 5.26.1
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


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n)spaceone/api/alert_manager/v1/alert.proto\x12\x1dspaceone.api.alert_manager.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\"+\n\x07Members\x12\x0c\n\x04USER\x18\x01 \x03(\t\x12\x12\n\nUSER_GROUP\x18\x02 \x03(\t\"\x8b\x01\n\x07Options\"F\n\x15notification_urgency_\x12\x15\n\x11NOTIFICATION_NONE\x10\x00\x12\x07\n\x03\x41LL\x10\x01\x12\r\n\tHIGH_ONLY\x10\x02\"8\n\rrecovery_mode\x12\x11\n\rRECOVERY_NONE\x10\x00\x12\x08\n\x04\x41UTO\x10\x01\x12\n\n\x06MANUAL\x10\x02\"\xaf\x01\n\x12\x41lertCreateRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x13\n\x0bservice_key\x18\x02 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\x12(\n\x07members\x18\x04 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x37\n\x07options\x18\x05 \x01(\x0b\x32&.spaceone.api.alert_manager.v1.Options\"\xa2\x01\n\x12\x41lertUpdateRequest\x12\x12\n\nservice_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\x12\x37\n\x07options\x18\x04 \x01(\x0b\x32&.spaceone.api.alert_manager.v1.Options\x12\x1c\n\x14\x65scalation_policy_id\x18\x05 \x01(\t\"g\n\x18\x41lertChangeMemberRequest\x12\x12\n\nservice_id\x18\x01 \x01(\t\x12\x37\n\x07members\x18\x02 \x01(\x0b\x32&.spaceone.api.alert_manager.v1.Members\"\"\n\x0c\x41lertRequest\x12\x12\n\nservice_id\x18\x01 \x01(\t\"\x89\x01\n\x10\x41lertSearchQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v2.Query\x12\x12\n\nservice_id\x18\x02 \x01(\t\x12\x1c\n\x14\x65scalation_policy_id\x18\x03 \x01(\t\x12\x17\n\x0finclude_details\x18\x04 \x01(\x08\"F\n\x0e\x41lertStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v2.StatisticsQuery\"\xaa\x03\n\x0bServiceInfo\x12\x12\n\nservice_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x13\n\x0bservice_key\x18\x03 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x04 \x01(\t\x12(\n\x07members\x18\x05 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x37\n\x07options\x18\x06 \x01(\x0b\x32&.spaceone.api.alert_manager.v1.Options\x12)\n\x08\x63hannels\x18\x07 \x01(\x0b\x32\x17.google.protobuf.Struct\x12)\n\x08webhooks\x18\x08 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\'\n\x06\x61lerts\x18\t \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x1c\n\x14\x65scalation_policy_id\x18\n \x01(\t\x12\x14\n\x0cworkspace_id\x18\x0b \x01(\t\x12\x11\n\tdomain_id\x18\x0c \x01(\t\x12\x12\n\ncreated_at\x18\r \x01(\t\x12\x12\n\nupdated_at\x18\x0e \x01(\t\"`\n\x0cServicesInfo\x12;\n\x07results\x18\x01 \x03(\x0b\x32*.spaceone.api.alert_manager.v1.ServiceInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"\\\n\x08\x43hannels\x12P\n\x14service_channel_info\x18\x01 \x03(\x0b\x32\x32.spaceone.api.alert_manager.v1.ServiceChannelsInfo\"\xce\x02\n\x13ServiceChannelsInfo\x12\x12\n\nchannel_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\r\n\x05state\x18\x03 \x01(\t\x12\x14\n\x0c\x63hannel_type\x18\x04 \x01(\t\x12%\n\x04\x64\x61ta\x18\x05 \x01(\x0b\x32\x17.google.protobuf.Struct\x12)\n\x08schedule\x18\x06 \x01(\x0b\x32\x17.google.protobuf.Struct\x12%\n\x04tags\x18\x07 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tsecret_id\x18\x08 \x01(\t\x12\x13\n\x0bprotocol_id\x18\t \x01(\t\x12\x12\n\nservice_id\x18\n \x01(\t\x12\x14\n\x0cworkspace_id\x18\x0b \x01(\t\x12\x11\n\tdomain_id\x18\x0c \x01(\t\x12\x12\n\ncreated_at\x18\r \x01(\t\"L\n\x08Webhooks\x12@\n\x0cwebhook_info\x18\x01 \x03(\x0b\x32*.spaceone.api.alert_manager.v1.WebhookInfo\"!\n\x0bWebhookInfo\x12\x12\n\nwebhook_id\x18\x01 \x01(\t\"\xbf\x01\n\x06\x41lerts\x12\x37\n\x05TOTAL\x18\x01 \x01(\x0b\x32(.spaceone.api.alert_manager.v1.AlertInfo\x12<\n\nTRIGGERRED\x18\x02 \x01(\x0b\x32(.spaceone.api.alert_manager.v1.AlertInfo\x12>\n\x0c\x41\x43KNOWLEDGED\x18\x03 \x01(\x0b\x32(.spaceone.api.alert_manager.v1.AlertInfo\"&\n\tAlertInfo\x12\x0c\n\x04high\x18\x01 \x01(\x05\x12\x0b\n\x03low\x18\x02 \x01(\x05\x32\xf7\x08\n\x05\x41lert\x12\x92\x01\n\x06\x63reate\x12\x31.spaceone.api.alert_manager.v1.AlertCreateRequest\x1a*.spaceone.api.alert_manager.v1.ServiceInfo\")\x82\xd3\xe4\x93\x02#\"\x1e/alert-manager/v1/alert/create:\x01*\x12\x92\x01\n\x06update\x12\x31.spaceone.api.alert_manager.v1.AlertUpdateRequest\x1a*.spaceone.api.alert_manager.v1.ServiceInfo\")\x82\xd3\xe4\x93\x02#\"\x1e/alert-manager/v1/alert/update:\x01*\x12x\n\x06\x64\x65lete\x12+.spaceone.api.alert_manager.v1.AlertRequest\x1a\x16.google.protobuf.Empty\")\x82\xd3\xe4\x93\x02#\"\x1e/alert-manager/v1/alert/delete:\x01*\x12\x86\x01\n\x03get\x12+.spaceone.api.alert_manager.v1.AlertRequest\x1a*.spaceone.api.alert_manager.v1.ServiceInfo\"&\x82\xd3\xe4\x93\x02 \"\x1b/alert-manager/v1/alert/get:\x01*\x12\x8d\x01\n\x04list\x12/.spaceone.api.alert_manager.v1.AlertSearchQuery\x1a+.spaceone.api.alert_manager.v1.ServicesInfo\"\'\x82\xd3\xe4\x93\x02!\"\x1c/alert-manager/v1/alert/list:\x01*\x12\x9a\x01\n\x07history\x12\x37.spaceone.api.alert_manager.v1.AlertChangeMemberRequest\x1a*.spaceone.api.alert_manager.v1.ServiceInfo\"*\x82\xd3\xe4\x93\x02$\"\x1f/alert-manager/v1/alert/history:\x01*\x12\x9a\x01\n\x07\x61nalyze\x12\x37.spaceone.api.alert_manager.v1.AlertChangeMemberRequest\x1a*.spaceone.api.alert_manager.v1.ServiceInfo\"*\x82\xd3\xe4\x93\x02$\"\x1f/alert-manager/v1/alert/analyze:\x01*\x12w\n\x04stat\x12-.spaceone.api.alert_manager.v1.AlertStatQuery\x1a\x17.google.protobuf.Struct\"\'\x82\xd3\xe4\x93\x02!\"\x1c/alert_manager/v1/alert/stat:\x01*BDZBgithub.com/cloudforet-io/api/dist/go/spaceone/api/alert-manager/v1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.alert_manager.v1.alert_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'ZBgithub.com/cloudforet-io/api/dist/go/spaceone/api/alert-manager/v1'
  _globals['_ALERT'].methods_by_name['create']._loaded_options = None
  _globals['_ALERT'].methods_by_name['create']._serialized_options = b'\202\323\344\223\002#\"\036/alert-manager/v1/alert/create:\001*'
  _globals['_ALERT'].methods_by_name['update']._loaded_options = None
  _globals['_ALERT'].methods_by_name['update']._serialized_options = b'\202\323\344\223\002#\"\036/alert-manager/v1/alert/update:\001*'
  _globals['_ALERT'].methods_by_name['delete']._loaded_options = None
  _globals['_ALERT'].methods_by_name['delete']._serialized_options = b'\202\323\344\223\002#\"\036/alert-manager/v1/alert/delete:\001*'
  _globals['_ALERT'].methods_by_name['get']._loaded_options = None
  _globals['_ALERT'].methods_by_name['get']._serialized_options = b'\202\323\344\223\002 \"\033/alert-manager/v1/alert/get:\001*'
  _globals['_ALERT'].methods_by_name['list']._loaded_options = None
  _globals['_ALERT'].methods_by_name['list']._serialized_options = b'\202\323\344\223\002!\"\034/alert-manager/v1/alert/list:\001*'
  _globals['_ALERT'].methods_by_name['history']._loaded_options = None
  _globals['_ALERT'].methods_by_name['history']._serialized_options = b'\202\323\344\223\002$\"\037/alert-manager/v1/alert/history:\001*'
  _globals['_ALERT'].methods_by_name['analyze']._loaded_options = None
  _globals['_ALERT'].methods_by_name['analyze']._serialized_options = b'\202\323\344\223\002$\"\037/alert-manager/v1/alert/analyze:\001*'
  _globals['_ALERT'].methods_by_name['stat']._loaded_options = None
  _globals['_ALERT'].methods_by_name['stat']._serialized_options = b'\202\323\344\223\002!\"\034/alert_manager/v1/alert/stat:\001*'
  _globals['_MEMBERS']._serialized_start=199
  _globals['_MEMBERS']._serialized_end=242
  _globals['_OPTIONS']._serialized_start=245
  _globals['_OPTIONS']._serialized_end=384
  _globals['_OPTIONS_NOTIFICATION_URGENCY_']._serialized_start=256
  _globals['_OPTIONS_NOTIFICATION_URGENCY_']._serialized_end=326
  _globals['_OPTIONS_RECOVERY_MODE']._serialized_start=328
  _globals['_OPTIONS_RECOVERY_MODE']._serialized_end=384
  _globals['_ALERTCREATEREQUEST']._serialized_start=387
  _globals['_ALERTCREATEREQUEST']._serialized_end=562
  _globals['_ALERTUPDATEREQUEST']._serialized_start=565
  _globals['_ALERTUPDATEREQUEST']._serialized_end=727
  _globals['_ALERTCHANGEMEMBERREQUEST']._serialized_start=729
  _globals['_ALERTCHANGEMEMBERREQUEST']._serialized_end=832
  _globals['_ALERTREQUEST']._serialized_start=834
  _globals['_ALERTREQUEST']._serialized_end=868
  _globals['_ALERTSEARCHQUERY']._serialized_start=871
  _globals['_ALERTSEARCHQUERY']._serialized_end=1008
  _globals['_ALERTSTATQUERY']._serialized_start=1010
  _globals['_ALERTSTATQUERY']._serialized_end=1080
  _globals['_SERVICEINFO']._serialized_start=1083
  _globals['_SERVICEINFO']._serialized_end=1509
  _globals['_SERVICESINFO']._serialized_start=1511
  _globals['_SERVICESINFO']._serialized_end=1607
  _globals['_CHANNELS']._serialized_start=1609
  _globals['_CHANNELS']._serialized_end=1701
  _globals['_SERVICECHANNELSINFO']._serialized_start=1704
  _globals['_SERVICECHANNELSINFO']._serialized_end=2038
  _globals['_WEBHOOKS']._serialized_start=2040
  _globals['_WEBHOOKS']._serialized_end=2116
  _globals['_WEBHOOKINFO']._serialized_start=2118
  _globals['_WEBHOOKINFO']._serialized_end=2151
  _globals['_ALERTS']._serialized_start=2154
  _globals['_ALERTS']._serialized_end=2345
  _globals['_ALERTINFO']._serialized_start=2347
  _globals['_ALERTINFO']._serialized_end=2385
  _globals['_ALERT']._serialized_start=2388
  _globals['_ALERT']._serialized_end=3531
# @@protoc_insertion_point(module_scope)
