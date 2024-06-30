# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/dashboard/v1/public_folder.proto
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


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n-spaceone/api/dashboard/v1/public_folder.proto\x12\x19spaceone.api.dashboard.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\"2\n\x0cSharedFolder\x12\x11\n\tworkspace\x18\x01 \x01(\x08\x12\x0f\n\x07project\x18\x02 \x01(\x08\"\xd5\x02\n\x19\x43reatePublicFolderRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12%\n\x04tags\x18\x02 \x01(\x0b\x32\x17.google.protobuf.Struct\x12+\n\ndashboards\x18\x03 \x03(\x0b\x32\x17.google.protobuf.Struct\x12Z\n\x0eresource_group\x18\x14 \x01(\x0e\x32\x42.spaceone.api.dashboard.v1.CreatePublicFolderRequest.ResourceGroup\x12\x14\n\x0cworkspace_id\x18\x15 \x01(\t\x12\x12\n\nproject_id\x18\x16 \x01(\t\"P\n\rResourceGroup\x12\x17\n\x13RESOURCE_GROUP_NONE\x10\x00\x12\n\n\x06\x44OMAIN\x10\x01\x12\r\n\tWORKSPACE\x10\x02\x12\x0b\n\x07PROJECT\x10\x03\"c\n\x19UpdatePublicFolderRequest\x12\x11\n\tfolder_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12%\n\x04tags\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\"(\n\x13PublicFolderRequest\x12\x11\n\tfolder_id\x18\x01 \x01(\t\"Q\n\x18SharePublicFolderRequest\x12\x11\n\tfolder_id\x18\x01 \x01(\t\x12\x11\n\tworkspace\x18\x02 \x01(\x08\x12\x0f\n\x07project\x18\x03 \x01(\x08\"\x8a\x01\n\x11PublicFolderQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v2.Query\x12\x11\n\tfolder_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x15 \x01(\t\x12\x12\n\nproject_id\x18\x16 \x01(\t\"\x9d\x03\n\x10PublicFolderInfo\x12\x11\n\tfolder_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12%\n\x04tags\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x37\n\x06shared\x18\x13 \x01(\x0b\x32\'.spaceone.api.dashboard.v1.SharedFolder\x12Q\n\x0eresource_group\x18\x14 \x01(\x0e\x32\x39.spaceone.api.dashboard.v1.PublicFolderInfo.ResourceGroup\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x16 \x01(\t\x12\x12\n\nproject_id\x18\x17 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\x12\x12\n\nupdated_at\x18  \x01(\t\"P\n\rResourceGroup\x12\x17\n\x13RESOURCE_GROUP_NONE\x10\x00\x12\n\n\x06\x44OMAIN\x10\x01\x12\r\n\tWORKSPACE\x10\x02\x12\x0b\n\x07PROJECT\x10\x03\"f\n\x11PublicFoldersInfo\x12<\n\x07results\x18\x01 \x03(\x0b\x32+.spaceone.api.dashboard.v1.PublicFolderInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"M\n\x15PublicFolderStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v2.StatisticsQuery2\x86\x08\n\x0cPublicFolder\x12\x9a\x01\n\x06\x63reate\x12\x34.spaceone.api.dashboard.v1.CreatePublicFolderRequest\x1a+.spaceone.api.dashboard.v1.PublicFolderInfo\"-\x82\xd3\xe4\x93\x02\'\"\"/dashboard/v1/public-folder/create:\x01*\x12\x9a\x01\n\x06update\x12\x34.spaceone.api.dashboard.v1.UpdatePublicFolderRequest\x1a+.spaceone.api.dashboard.v1.PublicFolderInfo\"-\x82\xd3\xe4\x93\x02\'\"\"/dashboard/v1/public-folder/update:\x01*\x12\x97\x01\n\x05share\x12\x33.spaceone.api.dashboard.v1.SharePublicFolderRequest\x1a+.spaceone.api.dashboard.v1.PublicFolderInfo\",\x82\xd3\xe4\x93\x02&\"!/dashboard/v1/public-folder/share:\x01*\x12\x7f\n\x06\x64\x65lete\x12..spaceone.api.dashboard.v1.PublicFolderRequest\x1a\x16.google.protobuf.Empty\"-\x82\xd3\xe4\x93\x02\'\"\"/dashboard/v1/public-folder/delete:\x01*\x12\x8e\x01\n\x03get\x12..spaceone.api.dashboard.v1.PublicFolderRequest\x1a+.spaceone.api.dashboard.v1.PublicFolderInfo\"*\x82\xd3\xe4\x93\x02$\"\x1f/dashboard/v1/public-folder/get:\x01*\x12\x8f\x01\n\x04list\x12,.spaceone.api.dashboard.v1.PublicFolderQuery\x1a,.spaceone.api.dashboard.v1.PublicFoldersInfo\"+\x82\xd3\xe4\x93\x02%\" /dashboard/v1/public-folder/list:\x01*\x12~\n\x04stat\x12\x30.spaceone.api.dashboard.v1.PublicFolderStatQuery\x1a\x17.google.protobuf.Struct\"+\x82\xd3\xe4\x93\x02%\" /dashboard/v1/public-folder/stat:\x01*B@Z>github.com/cloudforet-io/api/dist/go/spaceone/api/dashboard/v1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.dashboard.v1.public_folder_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z>github.com/cloudforet-io/api/dist/go/spaceone/api/dashboard/v1'
  _globals['_PUBLICFOLDER'].methods_by_name['create']._loaded_options = None
  _globals['_PUBLICFOLDER'].methods_by_name['create']._serialized_options = b'\202\323\344\223\002\'\"\"/dashboard/v1/public-folder/create:\001*'
  _globals['_PUBLICFOLDER'].methods_by_name['update']._loaded_options = None
  _globals['_PUBLICFOLDER'].methods_by_name['update']._serialized_options = b'\202\323\344\223\002\'\"\"/dashboard/v1/public-folder/update:\001*'
  _globals['_PUBLICFOLDER'].methods_by_name['share']._loaded_options = None
  _globals['_PUBLICFOLDER'].methods_by_name['share']._serialized_options = b'\202\323\344\223\002&\"!/dashboard/v1/public-folder/share:\001*'
  _globals['_PUBLICFOLDER'].methods_by_name['delete']._loaded_options = None
  _globals['_PUBLICFOLDER'].methods_by_name['delete']._serialized_options = b'\202\323\344\223\002\'\"\"/dashboard/v1/public-folder/delete:\001*'
  _globals['_PUBLICFOLDER'].methods_by_name['get']._loaded_options = None
  _globals['_PUBLICFOLDER'].methods_by_name['get']._serialized_options = b'\202\323\344\223\002$\"\037/dashboard/v1/public-folder/get:\001*'
  _globals['_PUBLICFOLDER'].methods_by_name['list']._loaded_options = None
  _globals['_PUBLICFOLDER'].methods_by_name['list']._serialized_options = b'\202\323\344\223\002%\" /dashboard/v1/public-folder/list:\001*'
  _globals['_PUBLICFOLDER'].methods_by_name['stat']._loaded_options = None
  _globals['_PUBLICFOLDER'].methods_by_name['stat']._serialized_options = b'\202\323\344\223\002%\" /dashboard/v1/public-folder/stat:\001*'
  _globals['_SHAREDFOLDER']._serialized_start=199
  _globals['_SHAREDFOLDER']._serialized_end=249
  _globals['_CREATEPUBLICFOLDERREQUEST']._serialized_start=252
  _globals['_CREATEPUBLICFOLDERREQUEST']._serialized_end=593
  _globals['_CREATEPUBLICFOLDERREQUEST_RESOURCEGROUP']._serialized_start=513
  _globals['_CREATEPUBLICFOLDERREQUEST_RESOURCEGROUP']._serialized_end=593
  _globals['_UPDATEPUBLICFOLDERREQUEST']._serialized_start=595
  _globals['_UPDATEPUBLICFOLDERREQUEST']._serialized_end=694
  _globals['_PUBLICFOLDERREQUEST']._serialized_start=696
  _globals['_PUBLICFOLDERREQUEST']._serialized_end=736
  _globals['_SHAREPUBLICFOLDERREQUEST']._serialized_start=738
  _globals['_SHAREPUBLICFOLDERREQUEST']._serialized_end=819
  _globals['_PUBLICFOLDERQUERY']._serialized_start=822
  _globals['_PUBLICFOLDERQUERY']._serialized_end=960
  _globals['_PUBLICFOLDERINFO']._serialized_start=963
  _globals['_PUBLICFOLDERINFO']._serialized_end=1376
  _globals['_PUBLICFOLDERINFO_RESOURCEGROUP']._serialized_start=513
  _globals['_PUBLICFOLDERINFO_RESOURCEGROUP']._serialized_end=593
  _globals['_PUBLICFOLDERSINFO']._serialized_start=1378
  _globals['_PUBLICFOLDERSINFO']._serialized_end=1480
  _globals['_PUBLICFOLDERSTATQUERY']._serialized_start=1482
  _globals['_PUBLICFOLDERSTATQUERY']._serialized_end=1559
  _globals['_PUBLICFOLDER']._serialized_start=1562
  _globals['_PUBLICFOLDER']._serialized_end=2592
# @@protoc_insertion_point(module_scope)
