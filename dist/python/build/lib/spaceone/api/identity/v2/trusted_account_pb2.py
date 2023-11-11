# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/identity/v2/trusted_account.proto
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


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n.spaceone/api/identity/v2/trusted_account.proto\x12\x18spaceone.api.identity.v2\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v1/query.proto\"\xbc\x02\n\"CreateTrustedServiceAccountRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12%\n\x04\x64\x61ta\x18\x02 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x10\n\x08provider\x18\x03 \x01(\t\x12%\n\x04tags\x18\x04 \x01(\x0b\x32\x17.google.protobuf.Struct\x12Q\n\x05scope\x18\x05 \x01(\x0e\x32\x42.spaceone.api.identity.v2.CreateTrustedServiceAccountRequest.Scope\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x16 \x01(\t\",\n\x05Scope\x12\x08\n\x04NONE\x10\x00\x12\n\n\x06\x44OMAIN\x10\x01\x12\r\n\tWORKSPACE\x10\x02\"\xcd\x01\n\"UpdateTrustedServiceAccountRequest\x12\"\n\x1atrusted_service_account_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12%\n\x04\x64\x61ta\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12%\n\x04tags\x18\x04 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x16 \x01(\t\"k\n\x1cTrustedServiceAccountRequest\x12\"\n\x1atrusted_service_account_id\x18\x01 \x01(\t\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x16 \x01(\t\"\xba\x02\n TrustedServiceAccountSearchQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v1.Query\x12\"\n\x1atrusted_service_account_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x10\n\x08provider\x18\x04 \x01(\t\x12O\n\x05scope\x18\x08 \x01(\x0e\x32@.spaceone.api.identity.v2.TrustedServiceAccountSearchQuery.Scope\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x16 \x01(\t\",\n\x05Scope\x12\x08\n\x04NONE\x10\x00\x12\n\n\x06\x44OMAIN\x10\x01\x12\r\n\tWORKSPACE\x10\x02\"\xe2\x02\n\x19TrustedServiceAccountInfo\x12\"\n\x1atrusted_service_account_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12%\n\x04\x64\x61ta\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x10\n\x08provider\x18\x04 \x01(\t\x12%\n\x04tags\x18\x05 \x01(\x0b\x32\x17.google.protobuf.Struct\x12H\n\x05scope\x18\x06 \x01(\x0e\x32\x39.spaceone.api.identity.v2.TrustedServiceAccountInfo.Scope\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x16 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\",\n\x05Scope\x12\x08\n\x04NONE\x10\x00\x12\n\n\x06\x44OMAIN\x10\x01\x12\r\n\tWORKSPACE\x10\x02\"w\n\x1aTrustedServiceAccountsInfo\x12\x44\n\x07results\x18\x01 \x03(\x0b\x32\x33.spaceone.api.identity.v2.TrustedServiceAccountInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"\x7f\n\x1eTrustedServiceAccountStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v1.StatisticsQuery\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x16 \x01(\t2\xd3\x07\n\x15TrustedServiceAccount\x12\xab\x01\n\x06\x63reate\x12<.spaceone.api.identity.v2.CreateTrustedServiceAccountRequest\x1a\x33.spaceone.api.identity.v2.TrustedServiceAccountInfo\".\x82\xd3\xe4\x93\x02(\"#/identity/v2/trusted-account/create:\x01*\x12\xab\x01\n\x06update\x12<.spaceone.api.identity.v2.UpdateTrustedServiceAccountRequest\x1a\x33.spaceone.api.identity.v2.TrustedServiceAccountInfo\".\x82\xd3\xe4\x93\x02(\"#/identity/v2/trusted-account/update:\x01*\x12\x88\x01\n\x06\x64\x65lete\x12\x36.spaceone.api.identity.v2.TrustedServiceAccountRequest\x1a\x16.google.protobuf.Empty\".\x82\xd3\xe4\x93\x02(\"#/identity/v2/trusted-account/delete:\x01*\x12\x9f\x01\n\x03get\x12\x36.spaceone.api.identity.v2.TrustedServiceAccountRequest\x1a\x33.spaceone.api.identity.v2.TrustedServiceAccountInfo\"+\x82\xd3\xe4\x93\x02%\" /identity/v2/trusted-account/get:\x01*\x12\xa6\x01\n\x04list\x12:.spaceone.api.identity.v2.TrustedServiceAccountSearchQuery\x1a\x34.spaceone.api.identity.v2.TrustedServiceAccountsInfo\",\x82\xd3\xe4\x93\x02&\"!/identity/v2/trusted-account/list:\x01*\x12\x87\x01\n\x04stat\x12\x38.spaceone.api.identity.v2.TrustedServiceAccountStatQuery\x1a\x17.google.protobuf.Struct\",\x82\xd3\xe4\x93\x02&\"!/identity/v2/trusted-account/stat:\x01*B?Z=github.com/cloudforet-io/api/dist/go/spaceone/api/identity/v2b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.identity.v2.trusted_account_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z=github.com/cloudforet-io/api/dist/go/spaceone/api/identity/v2'
  _TRUSTEDSERVICEACCOUNT.methods_by_name['create']._options = None
  _TRUSTEDSERVICEACCOUNT.methods_by_name['create']._serialized_options = b'\202\323\344\223\002(\"#/identity/v2/trusted-account/create:\001*'
  _TRUSTEDSERVICEACCOUNT.methods_by_name['update']._options = None
  _TRUSTEDSERVICEACCOUNT.methods_by_name['update']._serialized_options = b'\202\323\344\223\002(\"#/identity/v2/trusted-account/update:\001*'
  _TRUSTEDSERVICEACCOUNT.methods_by_name['delete']._options = None
  _TRUSTEDSERVICEACCOUNT.methods_by_name['delete']._serialized_options = b'\202\323\344\223\002(\"#/identity/v2/trusted-account/delete:\001*'
  _TRUSTEDSERVICEACCOUNT.methods_by_name['get']._options = None
  _TRUSTEDSERVICEACCOUNT.methods_by_name['get']._serialized_options = b'\202\323\344\223\002%\" /identity/v2/trusted-account/get:\001*'
  _TRUSTEDSERVICEACCOUNT.methods_by_name['list']._options = None
  _TRUSTEDSERVICEACCOUNT.methods_by_name['list']._serialized_options = b'\202\323\344\223\002&\"!/identity/v2/trusted-account/list:\001*'
  _TRUSTEDSERVICEACCOUNT.methods_by_name['stat']._options = None
  _TRUSTEDSERVICEACCOUNT.methods_by_name['stat']._serialized_options = b'\202\323\344\223\002&\"!/identity/v2/trusted-account/stat:\001*'
  _globals['_CREATETRUSTEDSERVICEACCOUNTREQUEST']._serialized_start=200
  _globals['_CREATETRUSTEDSERVICEACCOUNTREQUEST']._serialized_end=516
  _globals['_CREATETRUSTEDSERVICEACCOUNTREQUEST_SCOPE']._serialized_start=472
  _globals['_CREATETRUSTEDSERVICEACCOUNTREQUEST_SCOPE']._serialized_end=516
  _globals['_UPDATETRUSTEDSERVICEACCOUNTREQUEST']._serialized_start=519
  _globals['_UPDATETRUSTEDSERVICEACCOUNTREQUEST']._serialized_end=724
  _globals['_TRUSTEDSERVICEACCOUNTREQUEST']._serialized_start=726
  _globals['_TRUSTEDSERVICEACCOUNTREQUEST']._serialized_end=833
  _globals['_TRUSTEDSERVICEACCOUNTSEARCHQUERY']._serialized_start=836
  _globals['_TRUSTEDSERVICEACCOUNTSEARCHQUERY']._serialized_end=1150
  _globals['_TRUSTEDSERVICEACCOUNTSEARCHQUERY_SCOPE']._serialized_start=472
  _globals['_TRUSTEDSERVICEACCOUNTSEARCHQUERY_SCOPE']._serialized_end=516
  _globals['_TRUSTEDSERVICEACCOUNTINFO']._serialized_start=1153
  _globals['_TRUSTEDSERVICEACCOUNTINFO']._serialized_end=1507
  _globals['_TRUSTEDSERVICEACCOUNTINFO_SCOPE']._serialized_start=472
  _globals['_TRUSTEDSERVICEACCOUNTINFO_SCOPE']._serialized_end=516
  _globals['_TRUSTEDSERVICEACCOUNTSINFO']._serialized_start=1509
  _globals['_TRUSTEDSERVICEACCOUNTSINFO']._serialized_end=1628
  _globals['_TRUSTEDSERVICEACCOUNTSTATQUERY']._serialized_start=1630
  _globals['_TRUSTEDSERVICEACCOUNTSTATQUERY']._serialized_end=1757
  _globals['_TRUSTEDSERVICEACCOUNT']._serialized_start=1760
  _globals['_TRUSTEDSERVICEACCOUNT']._serialized_end=2739
# @@protoc_insertion_point(module_scope)
