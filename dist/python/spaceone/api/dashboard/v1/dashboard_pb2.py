# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/dashboard/v1/dashboard.proto
# Protobuf Python Version: 4.25.0
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
from spaceone.api.core.v1 import query_pb2 as spaceone_dot_api_dot_core_dot_v1_dot_query__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n)spaceone/api/dashboard/v1/dashboard.proto\x12\x19spaceone.api.dashboard.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v1/query.proto\"\x8b\x05\n\x16\x43reateDashboardRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12W\n\x0e\x64\x61shboard_type\x18\x02 \x01(\x0e\x32?.spaceone.api.dashboard.v1.CreateDashboardRequest.DashboardType\x12+\n\x07layouts\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.ListValue\x12*\n\tvariables\x18\x05 \x01(\x0b\x32\x17.google.protobuf.Struct\x12)\n\x08settings\x18\x06 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x31\n\x10variables_schema\x18\x07 \x01(\x0b\x32\x17.google.protobuf.Struct\x12*\n\x06labels\x18\x08 \x01(\x0b\x32\x1a.google.protobuf.ListValue\x12%\n\x04tags\x18\t \x01(\x0b\x32\x17.google.protobuf.Struct\x12W\n\x0eresource_group\x18\x14 \x01(\x0e\x32?.spaceone.api.dashboard.v1.CreateDashboardRequest.ResourceGroup\x12\x12\n\nproject_id\x18\x15 \x01(\t\"A\n\rDashboardType\x12\x17\n\x13\x44\x41SHBOARD_TYPE_NONE\x10\x00\x12\n\n\x06PUBLIC\x10\x01\x12\x0b\n\x07PRIVATE\x10\x02\"P\n\rResourceGroup\x12\x17\n\x13RESOURCE_GROUP_NONE\x10\x00\x12\n\n\x06\x44OMAIN\x10\x01\x12\r\n\tWORKSPACE\x10\x02\x12\x0b\n\x07PROJECT\x10\x03\"\xc6\x02\n\x16UpdateDashboardRequest\x12\x14\n\x0c\x64\x61shboard_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12+\n\x07layouts\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.ListValue\x12*\n\tvariables\x18\x05 \x01(\x0b\x32\x17.google.protobuf.Struct\x12)\n\x08settings\x18\x06 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x31\n\x10variables_schema\x18\x07 \x01(\x0b\x32\x17.google.protobuf.Struct\x12*\n\x06labels\x18\x08 \x01(\x0b\x32\x1a.google.protobuf.ListValue\x12%\n\x04tags\x18\t \x01(\x0b\x32\x17.google.protobuf.Struct\"(\n\x10\x44\x61shboardRequest\x12\x14\n\x0c\x64\x61shboard_id\x18\x01 \x01(\t\"@\n\x17\x44\x61shboardVersionRequest\x12\x14\n\x0c\x64\x61shboard_id\x18\x01 \x01(\t\x12\x0f\n\x07version\x18\x04 \x01(\x05\"p\n\x1b\x44\x61shboardVersionSearchQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v1.Query\x12\x14\n\x0c\x64\x61shboard_id\x18\x02 \x01(\t\x12\x0f\n\x07version\x18\x03 \x01(\x05\"\xaf\x02\n\x0e\x44\x61shboardQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v1.Query\x12\x14\n\x0c\x64\x61shboard_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12O\n\x0e\x64\x61shboard_type\x18\x04 \x01(\x0e\x32\x37.spaceone.api.dashboard.v1.DashboardQuery.DashboardType\x12\x14\n\x0cworkspace_id\x18\x15 \x01(\t\x12\x12\n\nproject_id\x18\x16 \x01(\t\x12\x0f\n\x07user_id\x18\x17 \x01(\t\"A\n\rDashboardType\x12\x17\n\x13\x44\x41SHBOARD_TYPE_NONE\x10\x00\x12\n\n\x06PUBLIC\x10\x01\x12\x0b\n\x07PRIVATE\x10\x02\"\xf9\x05\n\rDashboardInfo\x12\x14\n\x0c\x64\x61shboard_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12N\n\x0e\x64\x61shboard_type\x18\x03 \x01(\x0e\x32\x36.spaceone.api.dashboard.v1.DashboardInfo.DashboardType\x12\x0f\n\x07version\x18\x04 \x01(\x05\x12+\n\x07layouts\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.ListValue\x12*\n\tvariables\x18\x06 \x01(\x0b\x32\x17.google.protobuf.Struct\x12)\n\x08settings\x18\x07 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x31\n\x10variables_schema\x18\x08 \x01(\x0b\x32\x17.google.protobuf.Struct\x12*\n\x06labels\x18\t \x01(\x0b\x32\x1a.google.protobuf.ListValue\x12%\n\x04tags\x18\n \x01(\x0b\x32\x17.google.protobuf.Struct\x12N\n\x0eresource_group\x18\x14 \x01(\x0e\x32\x36.spaceone.api.dashboard.v1.DashboardInfo.ResourceGroup\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x16 \x01(\t\x12\x12\n\nproject_id\x18\x17 \x01(\t\x12\x0f\n\x07user_id\x18\x18 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\x12\x12\n\nupdated_at\x18  \x01(\t\"A\n\rDashboardType\x12\x17\n\x13\x44\x41SHBOARD_TYPE_NONE\x10\x00\x12\n\n\x06PUBLIC\x10\x01\x12\x0b\n\x07PRIVATE\x10\x02\"P\n\rResourceGroup\x12\x17\n\x13RESOURCE_GROUP_NONE\x10\x00\x12\n\n\x06\x44OMAIN\x10\x01\x12\r\n\tWORKSPACE\x10\x02\x12\x0b\n\x07PROJECT\x10\x03\"`\n\x0e\x44\x61shboardsInfo\x12\x39\n\x07results\x18\x01 \x03(\x0b\x32(.spaceone.api.dashboard.v1.DashboardInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"J\n\x12\x44\x61shboardStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v1.StatisticsQuery\"\xab\x02\n\x14\x44\x61shboardVersionInfo\x12\x14\n\x0c\x64\x61shboard_id\x18\x01 \x01(\t\x12\x0f\n\x07version\x18\x02 \x01(\x05\x12\x0e\n\x06latest\x18\x03 \x01(\x08\x12+\n\x07layouts\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.ListValue\x12*\n\tvariables\x18\x05 \x01(\x0b\x32\x17.google.protobuf.Struct\x12)\n\x08settings\x18\x06 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x31\n\x10variables_schema\x18\x07 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\"n\n\x15\x44\x61shboardVersionsInfo\x12@\n\x07results\x18\x01 \x03(\x0b\x32/.spaceone.api.dashboard.v1.DashboardVersionInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\x32\xbc\x0b\n\tDashboard\x12\x90\x01\n\x06\x63reate\x12\x31.spaceone.api.dashboard.v1.CreateDashboardRequest\x1a(.spaceone.api.dashboard.v1.DashboardInfo\")\x82\xd3\xe4\x93\x02#\"\x1e/dashboard/v1/dashboard/create:\x01*\x12\x90\x01\n\x06update\x12\x31.spaceone.api.dashboard.v1.UpdateDashboardRequest\x1a(.spaceone.api.dashboard.v1.DashboardInfo\")\x82\xd3\xe4\x93\x02#\"\x1e/dashboard/v1/dashboard/update:\x01*\x12x\n\x06\x64\x65lete\x12+.spaceone.api.dashboard.v1.DashboardRequest\x1a\x16.google.protobuf.Empty\")\x82\xd3\xe4\x93\x02#\"\x1e/dashboard/v1/dashboard/delete:\x01*\x12\x84\x01\n\x03get\x12+.spaceone.api.dashboard.v1.DashboardRequest\x1a(.spaceone.api.dashboard.v1.DashboardInfo\"&\x82\xd3\xe4\x93\x02 \"\x1b/dashboard/v1/dashboard/get:\x01*\x12\x8f\x01\n\x0e\x64\x65lete_version\x12\x32.spaceone.api.dashboard.v1.DashboardVersionRequest\x1a\x16.google.protobuf.Empty\"1\x82\xd3\xe4\x93\x02+\"&/dashboard/v1/dashboard/delete-version:\x01*\x12\xa1\x01\n\x0erevert_version\x12\x32.spaceone.api.dashboard.v1.DashboardVersionRequest\x1a(.spaceone.api.dashboard.v1.DashboardInfo\"1\x82\xd3\xe4\x93\x02+\"&/dashboard/v1/dashboard/revert-version:\x01*\x12\xa2\x01\n\x0bget_version\x12\x32.spaceone.api.dashboard.v1.DashboardVersionRequest\x1a/.spaceone.api.dashboard.v1.DashboardVersionInfo\".\x82\xd3\xe4\x93\x02(\"#/dashboard/v1/dashboard/get-version:\x01*\x12\xab\x01\n\rlist_versions\x12\x36.spaceone.api.dashboard.v1.DashboardVersionSearchQuery\x1a\x30.spaceone.api.dashboard.v1.DashboardVersionsInfo\"0\x82\xd3\xe4\x93\x02*\"%/dashboard/v1/dashboard/list-versions:\x01*\x12\x85\x01\n\x04list\x12).spaceone.api.dashboard.v1.DashboardQuery\x1a).spaceone.api.dashboard.v1.DashboardsInfo\"\'\x82\xd3\xe4\x93\x02!\"\x1c/dashboard/v1/dashboard/list:\x01*\x12w\n\x04stat\x12-.spaceone.api.dashboard.v1.DashboardStatQuery\x1a\x17.google.protobuf.Struct\"\'\x82\xd3\xe4\x93\x02!\"\x1c/dashboard/v1/dashboard/stat:\x01*B@Z>github.com/cloudforet-io/api/dist/go/spaceone/api/dashboard/v1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.dashboard.v1.dashboard_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z>github.com/cloudforet-io/api/dist/go/spaceone/api/dashboard/v1'
  _globals['_DASHBOARD'].methods_by_name['create']._options = None
  _globals['_DASHBOARD'].methods_by_name['create']._serialized_options = b'\202\323\344\223\002#\"\036/dashboard/v1/dashboard/create:\001*'
  _globals['_DASHBOARD'].methods_by_name['update']._options = None
  _globals['_DASHBOARD'].methods_by_name['update']._serialized_options = b'\202\323\344\223\002#\"\036/dashboard/v1/dashboard/update:\001*'
  _globals['_DASHBOARD'].methods_by_name['delete']._options = None
  _globals['_DASHBOARD'].methods_by_name['delete']._serialized_options = b'\202\323\344\223\002#\"\036/dashboard/v1/dashboard/delete:\001*'
  _globals['_DASHBOARD'].methods_by_name['get']._options = None
  _globals['_DASHBOARD'].methods_by_name['get']._serialized_options = b'\202\323\344\223\002 \"\033/dashboard/v1/dashboard/get:\001*'
  _globals['_DASHBOARD'].methods_by_name['delete_version']._options = None
  _globals['_DASHBOARD'].methods_by_name['delete_version']._serialized_options = b'\202\323\344\223\002+\"&/dashboard/v1/dashboard/delete-version:\001*'
  _globals['_DASHBOARD'].methods_by_name['revert_version']._options = None
  _globals['_DASHBOARD'].methods_by_name['revert_version']._serialized_options = b'\202\323\344\223\002+\"&/dashboard/v1/dashboard/revert-version:\001*'
  _globals['_DASHBOARD'].methods_by_name['get_version']._options = None
  _globals['_DASHBOARD'].methods_by_name['get_version']._serialized_options = b'\202\323\344\223\002(\"#/dashboard/v1/dashboard/get-version:\001*'
  _globals['_DASHBOARD'].methods_by_name['list_versions']._options = None
  _globals['_DASHBOARD'].methods_by_name['list_versions']._serialized_options = b'\202\323\344\223\002*\"%/dashboard/v1/dashboard/list-versions:\001*'
  _globals['_DASHBOARD'].methods_by_name['list']._options = None
  _globals['_DASHBOARD'].methods_by_name['list']._serialized_options = b'\202\323\344\223\002!\"\034/dashboard/v1/dashboard/list:\001*'
  _globals['_DASHBOARD'].methods_by_name['stat']._options = None
  _globals['_DASHBOARD'].methods_by_name['stat']._serialized_options = b'\202\323\344\223\002!\"\034/dashboard/v1/dashboard/stat:\001*'
  _globals['_CREATEDASHBOARDREQUEST']._serialized_start=196
  _globals['_CREATEDASHBOARDREQUEST']._serialized_end=847
  _globals['_CREATEDASHBOARDREQUEST_DASHBOARDTYPE']._serialized_start=700
  _globals['_CREATEDASHBOARDREQUEST_DASHBOARDTYPE']._serialized_end=765
  _globals['_CREATEDASHBOARDREQUEST_RESOURCEGROUP']._serialized_start=767
  _globals['_CREATEDASHBOARDREQUEST_RESOURCEGROUP']._serialized_end=847
  _globals['_UPDATEDASHBOARDREQUEST']._serialized_start=850
  _globals['_UPDATEDASHBOARDREQUEST']._serialized_end=1176
  _globals['_DASHBOARDREQUEST']._serialized_start=1178
  _globals['_DASHBOARDREQUEST']._serialized_end=1218
  _globals['_DASHBOARDVERSIONREQUEST']._serialized_start=1220
  _globals['_DASHBOARDVERSIONREQUEST']._serialized_end=1284
  _globals['_DASHBOARDVERSIONSEARCHQUERY']._serialized_start=1286
  _globals['_DASHBOARDVERSIONSEARCHQUERY']._serialized_end=1398
  _globals['_DASHBOARDQUERY']._serialized_start=1401
  _globals['_DASHBOARDQUERY']._serialized_end=1704
  _globals['_DASHBOARDQUERY_DASHBOARDTYPE']._serialized_start=700
  _globals['_DASHBOARDQUERY_DASHBOARDTYPE']._serialized_end=765
  _globals['_DASHBOARDINFO']._serialized_start=1707
  _globals['_DASHBOARDINFO']._serialized_end=2468
  _globals['_DASHBOARDINFO_DASHBOARDTYPE']._serialized_start=700
  _globals['_DASHBOARDINFO_DASHBOARDTYPE']._serialized_end=765
  _globals['_DASHBOARDINFO_RESOURCEGROUP']._serialized_start=767
  _globals['_DASHBOARDINFO_RESOURCEGROUP']._serialized_end=847
  _globals['_DASHBOARDSINFO']._serialized_start=2470
  _globals['_DASHBOARDSINFO']._serialized_end=2566
  _globals['_DASHBOARDSTATQUERY']._serialized_start=2568
  _globals['_DASHBOARDSTATQUERY']._serialized_end=2642
  _globals['_DASHBOARDVERSIONINFO']._serialized_start=2645
  _globals['_DASHBOARDVERSIONINFO']._serialized_end=2944
  _globals['_DASHBOARDVERSIONSINFO']._serialized_start=2946
  _globals['_DASHBOARDVERSIONSINFO']._serialized_end=3056
  _globals['_DASHBOARD']._serialized_start=3059
  _globals['_DASHBOARD']._serialized_end=4527
# @@protoc_insertion_point(module_scope)
