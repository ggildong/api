# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/inventory/v1/collector.proto
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
from spaceone.api.core.v1 import query_pb2 as spaceone_dot_api_dot_core_dot_v1_dot_query__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n)spaceone/api/inventory/v1/collector.proto\x12\x19spaceone.api.inventory.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v1/query.proto\"\x95\x01\n\tScheduled\x12\x42\n\x05state\x18\x01 \x01(\x0e\x32\x33.spaceone.api.inventory.v1.Scheduled.ScheduledState\x12\r\n\x05hours\x18\x02 \x03(\x05\"5\n\x0eScheduledState\x12\x08\n\x04NONE\x10\x00\x12\x0b\n\x07\x45NABLED\x10\x01\x12\x0c\n\x08\x44ISABLED\x10\x02\"\xce\x01\n\x0cSecretFilter\x12H\n\x05state\x18\x01 \x01(\x0e\x32\x39.spaceone.api.inventory.v1.SecretFilter.SecretFilterState\x12\x0f\n\x07secrets\x18\x02 \x03(\t\x12\x18\n\x10service_accounts\x18\x03 \x03(\t\x12\x0f\n\x07schemas\x18\x04 \x03(\t\"8\n\x11SecretFilterState\x12\x08\n\x04NONE\x10\x00\x12\x0b\n\x07\x45NABLED\x10\x01\x12\x0c\n\x08\x44ISABLED\x10\x02\"\xfd\x01\n\nPluginInfo\x12\x11\n\tplugin_id\x18\x01 \x01(\t\x12\x0f\n\x07version\x18\x02 \x01(\t\x12(\n\x07options\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12)\n\x08metadata\x18\n \x01(\x0b\x32\x17.google.protobuf.Struct\x12G\n\x0cupgrade_mode\x18\x0b \x01(\x0e\x32\x31.spaceone.api.inventory.v1.PluginInfo.UpgradeMode\"-\n\x0bUpgradeMode\x12\x08\n\x04NONE\x10\x00\x12\n\n\x06MANUAL\x10\x01\x12\x08\n\x04\x41UTO\x10\x02\"\xa6\x02\n\x16\x43reateCollectorRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12:\n\x0bplugin_info\x18\x02 \x01(\x0b\x32%.spaceone.api.inventory.v1.PluginInfo\x12\x36\n\x08schedule\x18\x03 \x01(\x0b\x32$.spaceone.api.inventory.v1.Scheduled\x12\x10\n\x08provider\x18\x04 \x01(\t\x12%\n\x04tags\x18\x05 \x01(\x0b\x32\x17.google.protobuf.Struct\x12>\n\rsecret_filter\x18\x06 \x01(\x0b\x32\'.spaceone.api.inventory.v1.SecretFilter\x12\x11\n\tdomain_id\x18\x0b \x01(\t\"\xee\x01\n\x16UpdateCollectorRequest\x12\x14\n\x0c\x63ollector_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x36\n\x08schedule\x18\x05 \x01(\x0b\x32$.spaceone.api.inventory.v1.Scheduled\x12>\n\rsecret_filter\x18\x06 \x01(\x0b\x32\'.spaceone.api.inventory.v1.SecretFilter\x12%\n\x04tags\x18\x07 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x15 \x01(\t\"\xfa\x01\n\x13UpdatePluginRequest\x12\x14\n\x0c\x63ollector_id\x18\x01 \x01(\t\x12\x0f\n\x07version\x18\x02 \x01(\t\x12(\n\x07options\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x04 \x01(\t\x12P\n\x0cupgrade_mode\x18\x05 \x01(\x0e\x32:.spaceone.api.inventory.v1.UpdatePluginRequest.UpgradeMode\"-\n\x0bUpgradeMode\x12\x08\n\x04NONE\x10\x00\x12\n\n\x06MANUAL\x10\x01\x12\x08\n\x04\x41UTO\x10\x02\";\n\x10\x43ollectorRequest\x12\x14\n\x0c\x63ollector_id\x18\x01 \x01(\t\x12\x11\n\tdomain_id\x18\x02 \x01(\t\"L\n\x13GetCollectorRequest\x12\x14\n\x0c\x63ollector_id\x18\x01 \x01(\t\x12\x11\n\tdomain_id\x18\x02 \x01(\t\x12\x0c\n\x04only\x18\x03 \x03(\t\"\x86\x02\n\x0e\x43ollectorQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v1.Query\x12\x14\n\x0c\x63ollector_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12>\n\x05state\x18\x04 \x01(\x0e\x32/.spaceone.api.inventory.v1.CollectorQuery.State\x12\x10\n\x08priority\x18\x05 \x01(\x05\x12\x11\n\tplugin_id\x18\x06 \x01(\t\x12\x11\n\tdomain_id\x18\x07 \x01(\t\",\n\x05State\x12\x08\n\x04NONE\x10\x00\x12\x0b\n\x07\x45NABLED\x10\x01\x12\x0c\n\x08\x44ISABLED\x10\x02\"\x9e\x01\n\x0e\x43ollectRequest\x12\x14\n\x0c\x63ollector_id\x18\x01 \x01(\t\x12\'\n\x06\x66ilter\x18\x02 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tsecret_id\x18\x03 \x01(\t\x12\x14\n\x0c\x63ollect_mode\x18\x04 \x01(\t\x12\x11\n\tdomain_id\x18\x05 \x01(\t\x12\x11\n\tuse_cache\x18\x06 \x01(\x08\"Q\n\x13VerifyPluginRequest\x12\x14\n\x0c\x63ollector_id\x18\x01 \x01(\t\x12\x11\n\tsecret_id\x18\x02 \x01(\t\x12\x11\n\tdomain_id\x18\x03 \x01(\t\"\xbd\x03\n\rCollectorInfo\x12\x14\n\x0c\x63ollector_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12:\n\x0bplugin_info\x18\x04 \x01(\x0b\x32%.spaceone.api.inventory.v1.PluginInfo\x12%\n\x04tags\x18\x05 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x10\n\x08provider\x18\x06 \x01(\t\x12+\n\ncapability\x18\x07 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x36\n\x08schedule\x18\x08 \x01(\x0b\x32$.spaceone.api.inventory.v1.Scheduled\x12>\n\rsecret_filter\x18\t \x01(\x0b\x32\'.spaceone.api.inventory.v1.SecretFilter\x12\x12\n\ncreated_at\x18\x0b \x01(\t\x12\x19\n\x11last_collected_at\x18\x0c \x01(\t\x12\x11\n\tdomain_id\x18\x15 \x01(\t\",\n\x05State\x12\x08\n\x04NONE\x10\x00\x12\x0b\n\x07\x45NABLED\x10\x01\x12\x0c\n\x08\x44ISABLED\x10\x02\"`\n\x0e\x43ollectorsInfo\x12\x39\n\x07results\x18\x01 \x03(\x0b\x32(.spaceone.api.inventory.v1.CollectorInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"]\n\x12\x43ollectorStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v1.StatisticsQuery\x12\x11\n\tdomain_id\x18\x02 \x01(\t\"\x1c\n\nVerifyInfo\x12\x0e\n\x06status\x18\x01 \x01(\x08\"]\n\tErrorInfo\x12\x12\n\nerror_code\x18\x01 \x01(\t\x12\x0f\n\x07message\x18\x02 \x01(\t\x12+\n\nadditional\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\"\xeb\x03\n\x07JobInfo\x12\x0e\n\x06job_id\x18\x01 \x01(\t\x12\x39\n\x06status\x18\x02 \x01(\x0e\x32).spaceone.api.inventory.v1.JobInfo.Status\x12\'\n\x06\x66ilter\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x13\n\x0btotal_tasks\x18\x04 \x01(\x05\x12\x16\n\x0eremained_tasks\x18\x05 \x01(\x05\x12\x34\n\x06\x65rrors\x18\x06 \x03(\x0b\x32$.spaceone.api.inventory.v1.ErrorInfo\x12@\n\x0e\x63ollector_info\x18\x0b \x01(\x0b\x32(.spaceone.api.inventory.v1.CollectorInfo\x12\x12\n\nproject_id\x18\x0c \x01(\t\x12\x11\n\tdomain_id\x18\r \x01(\t\x12\x12\n\ncreated_at\x18\x15 \x01(\t\x12\x12\n\nupdated_at\x18\x16 \x01(\t\x12\x13\n\x0b\x66inished_at\x18\x17 \x01(\t\"c\n\x06Status\x12\x08\n\x04NONE\x10\x00\x12\x0b\n\x07\x43REATED\x10\x01\x12\x0c\n\x08\x43\x41NCELED\x10\x02\x12\x0f\n\x0bIN_PROGRESS\x10\x03\x12\x0b\n\x07SUCCESS\x10\x04\x12\t\n\x05\x45RROR\x10\x05\x12\x0b\n\x07TIMEOUT\x10\x06\"\xc5\x01\n\x15\x43reateScheduleRequest\x12\x11\n\tdomain_id\x18\x01 \x01(\t\x12\x14\n\x0c\x63ollector_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x14\n\x0c\x63ollect_mode\x18\x04 \x01(\t\x12\x36\n\x08schedule\x18\x05 \x01(\x0b\x32$.spaceone.api.inventory.v1.Scheduled\x12\'\n\x06\x66ilter\x18\x06 \x01(\x0b\x32\x17.google.protobuf.Struct\"\xda\x01\n\x15UpdateScheduleRequest\x12\x11\n\tdomain_id\x18\x01 \x01(\t\x12\x13\n\x0bschedule_id\x18\x02 \x01(\t\x12\x14\n\x0c\x63ollector_id\x18\x03 \x01(\t\x12\x0c\n\x04name\x18\x04 \x01(\t\x12\x14\n\x0c\x63ollect_mode\x18\x05 \x01(\t\x12\x36\n\x08schedule\x18\x06 \x01(\x0b\x32$.spaceone.api.inventory.v1.Scheduled\x12\'\n\x06\x66ilter\x18\x07 \x01(\x0b\x32\x17.google.protobuf.Struct\"O\n\x0fScheduleRequest\x12\x11\n\tdomain_id\x18\x01 \x01(\t\x12\x13\n\x0bschedule_id\x18\x02 \x01(\t\x12\x14\n\x0c\x63ollector_id\x18\x03 \x01(\t\"U\n\x15\x44\x65leteScheduleRequest\x12\x11\n\tdomain_id\x18\x01 \x01(\t\x12\x13\n\x0bschedule_id\x18\x02 \x01(\t\x12\x14\n\x0c\x63ollector_id\x18\x03 \x01(\t\"y\n\rScheduleQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v1.Query\x12\x14\n\x0c\x63ollector_id\x18\x02 \x01(\t\x12\x13\n\x0bschedule_id\x18\x03 \x01(\t\x12\x11\n\tdomain_id\x18\x04 \x01(\t\"\xac\x02\n\x0cScheduleInfo\x12\x11\n\tdomain_id\x18\x01 \x01(\t\x12\x13\n\x0bschedule_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x14\n\x0c\x63ollect_mode\x18\x04 \x01(\t\x12\x36\n\x08schedule\x18\x05 \x01(\x0b\x32$.spaceone.api.inventory.v1.Scheduled\x12\x12\n\ncreated_at\x18\n \x01(\t\x12\x19\n\x11last_scheduled_at\x18\x0b \x01(\t\x12@\n\x0e\x63ollector_info\x18\x14 \x01(\x0b\x32(.spaceone.api.inventory.v1.CollectorInfo\x12\'\n\x06\x66ilter\x18\x15 \x01(\x0b\x32\x17.google.protobuf.Struct\"^\n\rSchedulesInfo\x12\x38\n\x07results\x18\x01 \x03(\x0b\x32\'.spaceone.api.inventory.v1.ScheduleInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\x32\xe7\t\n\tCollector\x12\x90\x01\n\x06\x63reate\x12\x31.spaceone.api.inventory.v1.CreateCollectorRequest\x1a(.spaceone.api.inventory.v1.CollectorInfo\")\x82\xd3\xe4\x93\x02#\"\x1e/inventory/v1/collector/create:\x01*\x12\x90\x01\n\x06update\x12\x31.spaceone.api.inventory.v1.UpdateCollectorRequest\x1a(.spaceone.api.inventory.v1.CollectorInfo\")\x82\xd3\xe4\x93\x02#\"\x1e/inventory/v1/collector/update:\x01*\x12\x9b\x01\n\rupdate_plugin\x12..spaceone.api.inventory.v1.UpdatePluginRequest\x1a(.spaceone.api.inventory.v1.CollectorInfo\"0\x82\xd3\xe4\x93\x02*\"%/inventory/v1/collector/update-plugin:\x01*\x12\x89\x01\n\rverify_plugin\x12..spaceone.api.inventory.v1.VerifyPluginRequest\x1a\x16.google.protobuf.Empty\"0\x82\xd3\xe4\x93\x02*\"%/inventory/v1/collector/verify-plugin:\x01*\x12x\n\x06\x64\x65lete\x12+.spaceone.api.inventory.v1.CollectorRequest\x1a\x16.google.protobuf.Empty\")\x82\xd3\xe4\x93\x02#\"\x1e/inventory/v1/collector/delete:\x01*\x12\x87\x01\n\x03get\x12..spaceone.api.inventory.v1.GetCollectorRequest\x1a(.spaceone.api.inventory.v1.CollectorInfo\"&\x82\xd3\xe4\x93\x02 \"\x1b/inventory/v1/collector/get:\x01*\x12\x85\x01\n\x04list\x12).spaceone.api.inventory.v1.CollectorQuery\x1a).spaceone.api.inventory.v1.CollectorsInfo\"\'\x82\xd3\xe4\x93\x02!\"\x1c/inventory/v1/collector/list:\x01*\x12w\n\x04stat\x12-.spaceone.api.inventory.v1.CollectorStatQuery\x1a\x17.google.protobuf.Struct\"\'\x82\xd3\xe4\x93\x02!\"\x1c/inventory/v1/collector/stat:\x01*\x12\x84\x01\n\x07\x63ollect\x12).spaceone.api.inventory.v1.CollectRequest\x1a\".spaceone.api.inventory.v1.JobInfo\"*\x82\xd3\xe4\x93\x02$\"\x1f/inventory/v1/collector/collect:\x01*B@Z>github.com/cloudforet-io/api/dist/go/spaceone/api/inventory/v1b\x06proto3')



_SCHEDULED = DESCRIPTOR.message_types_by_name['Scheduled']
_SECRETFILTER = DESCRIPTOR.message_types_by_name['SecretFilter']
_PLUGININFO = DESCRIPTOR.message_types_by_name['PluginInfo']
_CREATECOLLECTORREQUEST = DESCRIPTOR.message_types_by_name['CreateCollectorRequest']
_UPDATECOLLECTORREQUEST = DESCRIPTOR.message_types_by_name['UpdateCollectorRequest']
_UPDATEPLUGINREQUEST = DESCRIPTOR.message_types_by_name['UpdatePluginRequest']
_COLLECTORREQUEST = DESCRIPTOR.message_types_by_name['CollectorRequest']
_GETCOLLECTORREQUEST = DESCRIPTOR.message_types_by_name['GetCollectorRequest']
_COLLECTORQUERY = DESCRIPTOR.message_types_by_name['CollectorQuery']
_COLLECTREQUEST = DESCRIPTOR.message_types_by_name['CollectRequest']
_VERIFYPLUGINREQUEST = DESCRIPTOR.message_types_by_name['VerifyPluginRequest']
_COLLECTORINFO = DESCRIPTOR.message_types_by_name['CollectorInfo']
_COLLECTORSINFO = DESCRIPTOR.message_types_by_name['CollectorsInfo']
_COLLECTORSTATQUERY = DESCRIPTOR.message_types_by_name['CollectorStatQuery']
_VERIFYINFO = DESCRIPTOR.message_types_by_name['VerifyInfo']
_ERRORINFO = DESCRIPTOR.message_types_by_name['ErrorInfo']
_JOBINFO = DESCRIPTOR.message_types_by_name['JobInfo']
_CREATESCHEDULEREQUEST = DESCRIPTOR.message_types_by_name['CreateScheduleRequest']
_UPDATESCHEDULEREQUEST = DESCRIPTOR.message_types_by_name['UpdateScheduleRequest']
_SCHEDULEREQUEST = DESCRIPTOR.message_types_by_name['ScheduleRequest']
_DELETESCHEDULEREQUEST = DESCRIPTOR.message_types_by_name['DeleteScheduleRequest']
_SCHEDULEQUERY = DESCRIPTOR.message_types_by_name['ScheduleQuery']
_SCHEDULEINFO = DESCRIPTOR.message_types_by_name['ScheduleInfo']
_SCHEDULESINFO = DESCRIPTOR.message_types_by_name['SchedulesInfo']
_SCHEDULED_SCHEDULEDSTATE = _SCHEDULED.enum_types_by_name['ScheduledState']
_SECRETFILTER_SECRETFILTERSTATE = _SECRETFILTER.enum_types_by_name['SecretFilterState']
_PLUGININFO_UPGRADEMODE = _PLUGININFO.enum_types_by_name['UpgradeMode']
_UPDATEPLUGINREQUEST_UPGRADEMODE = _UPDATEPLUGINREQUEST.enum_types_by_name['UpgradeMode']
_COLLECTORQUERY_STATE = _COLLECTORQUERY.enum_types_by_name['State']
_COLLECTORINFO_STATE = _COLLECTORINFO.enum_types_by_name['State']
_JOBINFO_STATUS = _JOBINFO.enum_types_by_name['Status']
Scheduled = _reflection.GeneratedProtocolMessageType('Scheduled', (_message.Message,), {
  'DESCRIPTOR' : _SCHEDULED,
  '__module__' : 'spaceone.api.inventory.v1.collector_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.inventory.v1.Scheduled)
  })
_sym_db.RegisterMessage(Scheduled)

SecretFilter = _reflection.GeneratedProtocolMessageType('SecretFilter', (_message.Message,), {
  'DESCRIPTOR' : _SECRETFILTER,
  '__module__' : 'spaceone.api.inventory.v1.collector_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.inventory.v1.SecretFilter)
  })
_sym_db.RegisterMessage(SecretFilter)

PluginInfo = _reflection.GeneratedProtocolMessageType('PluginInfo', (_message.Message,), {
  'DESCRIPTOR' : _PLUGININFO,
  '__module__' : 'spaceone.api.inventory.v1.collector_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.inventory.v1.PluginInfo)
  })
_sym_db.RegisterMessage(PluginInfo)

CreateCollectorRequest = _reflection.GeneratedProtocolMessageType('CreateCollectorRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATECOLLECTORREQUEST,
  '__module__' : 'spaceone.api.inventory.v1.collector_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.inventory.v1.CreateCollectorRequest)
  })
_sym_db.RegisterMessage(CreateCollectorRequest)

UpdateCollectorRequest = _reflection.GeneratedProtocolMessageType('UpdateCollectorRequest', (_message.Message,), {
  'DESCRIPTOR' : _UPDATECOLLECTORREQUEST,
  '__module__' : 'spaceone.api.inventory.v1.collector_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.inventory.v1.UpdateCollectorRequest)
  })
_sym_db.RegisterMessage(UpdateCollectorRequest)

UpdatePluginRequest = _reflection.GeneratedProtocolMessageType('UpdatePluginRequest', (_message.Message,), {
  'DESCRIPTOR' : _UPDATEPLUGINREQUEST,
  '__module__' : 'spaceone.api.inventory.v1.collector_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.inventory.v1.UpdatePluginRequest)
  })
_sym_db.RegisterMessage(UpdatePluginRequest)

CollectorRequest = _reflection.GeneratedProtocolMessageType('CollectorRequest', (_message.Message,), {
  'DESCRIPTOR' : _COLLECTORREQUEST,
  '__module__' : 'spaceone.api.inventory.v1.collector_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.inventory.v1.CollectorRequest)
  })
_sym_db.RegisterMessage(CollectorRequest)

GetCollectorRequest = _reflection.GeneratedProtocolMessageType('GetCollectorRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETCOLLECTORREQUEST,
  '__module__' : 'spaceone.api.inventory.v1.collector_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.inventory.v1.GetCollectorRequest)
  })
_sym_db.RegisterMessage(GetCollectorRequest)

CollectorQuery = _reflection.GeneratedProtocolMessageType('CollectorQuery', (_message.Message,), {
  'DESCRIPTOR' : _COLLECTORQUERY,
  '__module__' : 'spaceone.api.inventory.v1.collector_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.inventory.v1.CollectorQuery)
  })
_sym_db.RegisterMessage(CollectorQuery)

CollectRequest = _reflection.GeneratedProtocolMessageType('CollectRequest', (_message.Message,), {
  'DESCRIPTOR' : _COLLECTREQUEST,
  '__module__' : 'spaceone.api.inventory.v1.collector_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.inventory.v1.CollectRequest)
  })
_sym_db.RegisterMessage(CollectRequest)

VerifyPluginRequest = _reflection.GeneratedProtocolMessageType('VerifyPluginRequest', (_message.Message,), {
  'DESCRIPTOR' : _VERIFYPLUGINREQUEST,
  '__module__' : 'spaceone.api.inventory.v1.collector_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.inventory.v1.VerifyPluginRequest)
  })
_sym_db.RegisterMessage(VerifyPluginRequest)

CollectorInfo = _reflection.GeneratedProtocolMessageType('CollectorInfo', (_message.Message,), {
  'DESCRIPTOR' : _COLLECTORINFO,
  '__module__' : 'spaceone.api.inventory.v1.collector_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.inventory.v1.CollectorInfo)
  })
_sym_db.RegisterMessage(CollectorInfo)

CollectorsInfo = _reflection.GeneratedProtocolMessageType('CollectorsInfo', (_message.Message,), {
  'DESCRIPTOR' : _COLLECTORSINFO,
  '__module__' : 'spaceone.api.inventory.v1.collector_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.inventory.v1.CollectorsInfo)
  })
_sym_db.RegisterMessage(CollectorsInfo)

CollectorStatQuery = _reflection.GeneratedProtocolMessageType('CollectorStatQuery', (_message.Message,), {
  'DESCRIPTOR' : _COLLECTORSTATQUERY,
  '__module__' : 'spaceone.api.inventory.v1.collector_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.inventory.v1.CollectorStatQuery)
  })
_sym_db.RegisterMessage(CollectorStatQuery)

VerifyInfo = _reflection.GeneratedProtocolMessageType('VerifyInfo', (_message.Message,), {
  'DESCRIPTOR' : _VERIFYINFO,
  '__module__' : 'spaceone.api.inventory.v1.collector_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.inventory.v1.VerifyInfo)
  })
_sym_db.RegisterMessage(VerifyInfo)

ErrorInfo = _reflection.GeneratedProtocolMessageType('ErrorInfo', (_message.Message,), {
  'DESCRIPTOR' : _ERRORINFO,
  '__module__' : 'spaceone.api.inventory.v1.collector_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.inventory.v1.ErrorInfo)
  })
_sym_db.RegisterMessage(ErrorInfo)

JobInfo = _reflection.GeneratedProtocolMessageType('JobInfo', (_message.Message,), {
  'DESCRIPTOR' : _JOBINFO,
  '__module__' : 'spaceone.api.inventory.v1.collector_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.inventory.v1.JobInfo)
  })
_sym_db.RegisterMessage(JobInfo)

CreateScheduleRequest = _reflection.GeneratedProtocolMessageType('CreateScheduleRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATESCHEDULEREQUEST,
  '__module__' : 'spaceone.api.inventory.v1.collector_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.inventory.v1.CreateScheduleRequest)
  })
_sym_db.RegisterMessage(CreateScheduleRequest)

UpdateScheduleRequest = _reflection.GeneratedProtocolMessageType('UpdateScheduleRequest', (_message.Message,), {
  'DESCRIPTOR' : _UPDATESCHEDULEREQUEST,
  '__module__' : 'spaceone.api.inventory.v1.collector_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.inventory.v1.UpdateScheduleRequest)
  })
_sym_db.RegisterMessage(UpdateScheduleRequest)

ScheduleRequest = _reflection.GeneratedProtocolMessageType('ScheduleRequest', (_message.Message,), {
  'DESCRIPTOR' : _SCHEDULEREQUEST,
  '__module__' : 'spaceone.api.inventory.v1.collector_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.inventory.v1.ScheduleRequest)
  })
_sym_db.RegisterMessage(ScheduleRequest)

DeleteScheduleRequest = _reflection.GeneratedProtocolMessageType('DeleteScheduleRequest', (_message.Message,), {
  'DESCRIPTOR' : _DELETESCHEDULEREQUEST,
  '__module__' : 'spaceone.api.inventory.v1.collector_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.inventory.v1.DeleteScheduleRequest)
  })
_sym_db.RegisterMessage(DeleteScheduleRequest)

ScheduleQuery = _reflection.GeneratedProtocolMessageType('ScheduleQuery', (_message.Message,), {
  'DESCRIPTOR' : _SCHEDULEQUERY,
  '__module__' : 'spaceone.api.inventory.v1.collector_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.inventory.v1.ScheduleQuery)
  })
_sym_db.RegisterMessage(ScheduleQuery)

ScheduleInfo = _reflection.GeneratedProtocolMessageType('ScheduleInfo', (_message.Message,), {
  'DESCRIPTOR' : _SCHEDULEINFO,
  '__module__' : 'spaceone.api.inventory.v1.collector_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.inventory.v1.ScheduleInfo)
  })
_sym_db.RegisterMessage(ScheduleInfo)

SchedulesInfo = _reflection.GeneratedProtocolMessageType('SchedulesInfo', (_message.Message,), {
  'DESCRIPTOR' : _SCHEDULESINFO,
  '__module__' : 'spaceone.api.inventory.v1.collector_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.inventory.v1.SchedulesInfo)
  })
_sym_db.RegisterMessage(SchedulesInfo)

_COLLECTOR = DESCRIPTOR.services_by_name['Collector']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z>github.com/cloudforet-io/api/dist/go/spaceone/api/inventory/v1'
  _COLLECTOR.methods_by_name['create']._options = None
  _COLLECTOR.methods_by_name['create']._serialized_options = b'\202\323\344\223\002#\"\036/inventory/v1/collector/create:\001*'
  _COLLECTOR.methods_by_name['update']._options = None
  _COLLECTOR.methods_by_name['update']._serialized_options = b'\202\323\344\223\002#\"\036/inventory/v1/collector/update:\001*'
  _COLLECTOR.methods_by_name['update_plugin']._options = None
  _COLLECTOR.methods_by_name['update_plugin']._serialized_options = b'\202\323\344\223\002*\"%/inventory/v1/collector/update-plugin:\001*'
  _COLLECTOR.methods_by_name['verify_plugin']._options = None
  _COLLECTOR.methods_by_name['verify_plugin']._serialized_options = b'\202\323\344\223\002*\"%/inventory/v1/collector/verify-plugin:\001*'
  _COLLECTOR.methods_by_name['delete']._options = None
  _COLLECTOR.methods_by_name['delete']._serialized_options = b'\202\323\344\223\002#\"\036/inventory/v1/collector/delete:\001*'
  _COLLECTOR.methods_by_name['get']._options = None
  _COLLECTOR.methods_by_name['get']._serialized_options = b'\202\323\344\223\002 \"\033/inventory/v1/collector/get:\001*'
  _COLLECTOR.methods_by_name['list']._options = None
  _COLLECTOR.methods_by_name['list']._serialized_options = b'\202\323\344\223\002!\"\034/inventory/v1/collector/list:\001*'
  _COLLECTOR.methods_by_name['stat']._options = None
  _COLLECTOR.methods_by_name['stat']._serialized_options = b'\202\323\344\223\002!\"\034/inventory/v1/collector/stat:\001*'
  _COLLECTOR.methods_by_name['collect']._options = None
  _COLLECTOR.methods_by_name['collect']._serialized_options = b'\202\323\344\223\002$\"\037/inventory/v1/collector/collect:\001*'
  _SCHEDULED._serialized_start=196
  _SCHEDULED._serialized_end=345
  _SCHEDULED_SCHEDULEDSTATE._serialized_start=292
  _SCHEDULED_SCHEDULEDSTATE._serialized_end=345
  _SECRETFILTER._serialized_start=348
  _SECRETFILTER._serialized_end=554
  _SECRETFILTER_SECRETFILTERSTATE._serialized_start=498
  _SECRETFILTER_SECRETFILTERSTATE._serialized_end=554
  _PLUGININFO._serialized_start=557
  _PLUGININFO._serialized_end=810
  _PLUGININFO_UPGRADEMODE._serialized_start=765
  _PLUGININFO_UPGRADEMODE._serialized_end=810
  _CREATECOLLECTORREQUEST._serialized_start=813
  _CREATECOLLECTORREQUEST._serialized_end=1107
  _UPDATECOLLECTORREQUEST._serialized_start=1110
  _UPDATECOLLECTORREQUEST._serialized_end=1348
  _UPDATEPLUGINREQUEST._serialized_start=1351
  _UPDATEPLUGINREQUEST._serialized_end=1601
  _UPDATEPLUGINREQUEST_UPGRADEMODE._serialized_start=765
  _UPDATEPLUGINREQUEST_UPGRADEMODE._serialized_end=810
  _COLLECTORREQUEST._serialized_start=1603
  _COLLECTORREQUEST._serialized_end=1662
  _GETCOLLECTORREQUEST._serialized_start=1664
  _GETCOLLECTORREQUEST._serialized_end=1740
  _COLLECTORQUERY._serialized_start=1743
  _COLLECTORQUERY._serialized_end=2005
  _COLLECTORQUERY_STATE._serialized_start=1961
  _COLLECTORQUERY_STATE._serialized_end=2005
  _COLLECTREQUEST._serialized_start=2008
  _COLLECTREQUEST._serialized_end=2166
  _VERIFYPLUGINREQUEST._serialized_start=2168
  _VERIFYPLUGINREQUEST._serialized_end=2249
  _COLLECTORINFO._serialized_start=2252
  _COLLECTORINFO._serialized_end=2697
  _COLLECTORINFO_STATE._serialized_start=1961
  _COLLECTORINFO_STATE._serialized_end=2005
  _COLLECTORSINFO._serialized_start=2699
  _COLLECTORSINFO._serialized_end=2795
  _COLLECTORSTATQUERY._serialized_start=2797
  _COLLECTORSTATQUERY._serialized_end=2890
  _VERIFYINFO._serialized_start=2892
  _VERIFYINFO._serialized_end=2920
  _ERRORINFO._serialized_start=2922
  _ERRORINFO._serialized_end=3015
  _JOBINFO._serialized_start=3018
  _JOBINFO._serialized_end=3509
  _JOBINFO_STATUS._serialized_start=3410
  _JOBINFO_STATUS._serialized_end=3509
  _CREATESCHEDULEREQUEST._serialized_start=3512
  _CREATESCHEDULEREQUEST._serialized_end=3709
  _UPDATESCHEDULEREQUEST._serialized_start=3712
  _UPDATESCHEDULEREQUEST._serialized_end=3930
  _SCHEDULEREQUEST._serialized_start=3932
  _SCHEDULEREQUEST._serialized_end=4011
  _DELETESCHEDULEREQUEST._serialized_start=4013
  _DELETESCHEDULEREQUEST._serialized_end=4098
  _SCHEDULEQUERY._serialized_start=4100
  _SCHEDULEQUERY._serialized_end=4221
  _SCHEDULEINFO._serialized_start=4224
  _SCHEDULEINFO._serialized_end=4524
  _SCHEDULESINFO._serialized_start=4526
  _SCHEDULESINFO._serialized_end=4620
  _COLLECTOR._serialized_start=4623
  _COLLECTOR._serialized_end=5878
# @@protoc_insertion_point(module_scope)