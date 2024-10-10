// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.6.1
// source: spaceone/api/cost_analysis/v1/cost_report_data.proto

package v1

import (
	v2 "github.com/cloudforet-io/api/dist/go/spaceone/api/core/v2"
	_ "github.com/golang/protobuf/ptypes/empty"
	_struct "github.com/golang/protobuf/ptypes/struct"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CostReportDataQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// +optional
	Query *v2.Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	CostReportDataId string `protobuf:"bytes,2,opt,name=cost_report_data_id,json=costReportDataId,proto3" json:"cost_report_data_id,omitempty"`
	// +optional
	Product string `protobuf:"bytes,3,opt,name=product,proto3" json:"product,omitempty"`
	// +optional
	Provider string `protobuf:"bytes,4,opt,name=provider,proto3" json:"provider,omitempty"`
	// +optional
	IsConfirmed bool `protobuf:"varint,5,opt,name=is_confirmed,json=isConfirmed,proto3" json:"is_confirmed,omitempty"`
	// +optional
	WorkspaceId string `protobuf:"bytes,21,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	// +optional
	CostReportConfigId string `protobuf:"bytes,22,opt,name=cost_report_config_id,json=costReportConfigId,proto3" json:"cost_report_config_id,omitempty"`
	// +optional
	CostReportId string `protobuf:"bytes,23,opt,name=cost_report_id,json=costReportId,proto3" json:"cost_report_id,omitempty"`
	// +optional
	DataSourceId string `protobuf:"bytes,24,opt,name=data_source_id,json=dataSourceId,proto3" json:"data_source_id,omitempty"`
}

func (x *CostReportDataQuery) Reset() {
	*x = CostReportDataQuery{}
	mi := &file_spaceone_api_cost_analysis_v1_cost_report_data_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CostReportDataQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CostReportDataQuery) ProtoMessage() {}

func (x *CostReportDataQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_cost_report_data_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CostReportDataQuery.ProtoReflect.Descriptor instead.
func (*CostReportDataQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_cost_report_data_proto_rawDescGZIP(), []int{0}
}

func (x *CostReportDataQuery) GetQuery() *v2.Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *CostReportDataQuery) GetCostReportDataId() string {
	if x != nil {
		return x.CostReportDataId
	}
	return ""
}

func (x *CostReportDataQuery) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

func (x *CostReportDataQuery) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *CostReportDataQuery) GetIsConfirmed() bool {
	if x != nil {
		return x.IsConfirmed
	}
	return false
}

func (x *CostReportDataQuery) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *CostReportDataQuery) GetCostReportConfigId() string {
	if x != nil {
		return x.CostReportConfigId
	}
	return ""
}

func (x *CostReportDataQuery) GetCostReportId() string {
	if x != nil {
		return x.CostReportId
	}
	return ""
}

func (x *CostReportDataQuery) GetDataSourceId() string {
	if x != nil {
		return x.DataSourceId
	}
	return ""
}

type CostReportDataAnalyzeQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// +optional
	Query *v2.TimeSeriesAnalyzeQuery `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	Product string `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
	// +optional
	Provider string `protobuf:"bytes,3,opt,name=provider,proto3" json:"provider,omitempty"`
	// +optional
	IsConfirmed bool `protobuf:"varint,4,opt,name=is_confirmed,json=isConfirmed,proto3" json:"is_confirmed,omitempty"`
	// +optional
	WorkspaceId string `protobuf:"bytes,21,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	// +optional
	ProjectId string `protobuf:"bytes,22,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// +optional
	CostReportConfigId string `protobuf:"bytes,23,opt,name=cost_report_config_id,json=costReportConfigId,proto3" json:"cost_report_config_id,omitempty"`
	// +optional
	CostReportId string `protobuf:"bytes,24,opt,name=cost_report_id,json=costReportId,proto3" json:"cost_report_id,omitempty"`
	// +optional
	DataSourceId string `protobuf:"bytes,25,opt,name=data_source_id,json=dataSourceId,proto3" json:"data_source_id,omitempty"`
}

func (x *CostReportDataAnalyzeQuery) Reset() {
	*x = CostReportDataAnalyzeQuery{}
	mi := &file_spaceone_api_cost_analysis_v1_cost_report_data_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CostReportDataAnalyzeQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CostReportDataAnalyzeQuery) ProtoMessage() {}

func (x *CostReportDataAnalyzeQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_cost_report_data_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CostReportDataAnalyzeQuery.ProtoReflect.Descriptor instead.
func (*CostReportDataAnalyzeQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_cost_report_data_proto_rawDescGZIP(), []int{1}
}

func (x *CostReportDataAnalyzeQuery) GetQuery() *v2.TimeSeriesAnalyzeQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *CostReportDataAnalyzeQuery) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

func (x *CostReportDataAnalyzeQuery) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *CostReportDataAnalyzeQuery) GetIsConfirmed() bool {
	if x != nil {
		return x.IsConfirmed
	}
	return false
}

func (x *CostReportDataAnalyzeQuery) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *CostReportDataAnalyzeQuery) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *CostReportDataAnalyzeQuery) GetCostReportConfigId() string {
	if x != nil {
		return x.CostReportConfigId
	}
	return ""
}

func (x *CostReportDataAnalyzeQuery) GetCostReportId() string {
	if x != nil {
		return x.CostReportId
	}
	return ""
}

func (x *CostReportDataAnalyzeQuery) GetDataSourceId() string {
	if x != nil {
		return x.DataSourceId
	}
	return ""
}

type CostReportDataInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CostReportDataId   string          `protobuf:"bytes,1,opt,name=cost_report_data_id,json=costReportDataId,proto3" json:"cost_report_data_id,omitempty"`
	Cost               *_struct.Struct `protobuf:"bytes,2,opt,name=cost,proto3" json:"cost,omitempty"`
	CostReportName     string          `protobuf:"bytes,3,opt,name=cost_report_name,json=costReportName,proto3" json:"cost_report_name,omitempty"`
	IssueDate          string          `protobuf:"bytes,4,opt,name=issue_date,json=issueDate,proto3" json:"issue_date,omitempty"`
	ReportYear         string          `protobuf:"bytes,5,opt,name=report_year,json=reportYear,proto3" json:"report_year,omitempty"`
	ReportMonth        string          `protobuf:"bytes,6,opt,name=report_month,json=reportMonth,proto3" json:"report_month,omitempty"`
	IsConfirmed        bool            `protobuf:"varint,7,opt,name=is_confirmed,json=isConfirmed,proto3" json:"is_confirmed,omitempty"`
	Provider           string          `protobuf:"bytes,8,opt,name=provider,proto3" json:"provider,omitempty"`
	Product            string          `protobuf:"bytes,9,opt,name=product,proto3" json:"product,omitempty"`
	ServiceAccountName string          `protobuf:"bytes,10,opt,name=service_account_name,json=serviceAccountName,proto3" json:"service_account_name,omitempty"`
	DataSourceName     string          `protobuf:"bytes,11,opt,name=data_source_name,json=dataSourceName,proto3" json:"data_source_name,omitempty"`
	ProjectName        string          `protobuf:"bytes,12,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	WorkspaceName      string          `protobuf:"bytes,13,opt,name=workspace_name,json=workspaceName,proto3" json:"workspace_name,omitempty"`
	DomainId           string          `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	WorkspaceId        string          `protobuf:"bytes,22,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	ProjectId          string          `protobuf:"bytes,23,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	CostReportConfigId string          `protobuf:"bytes,24,opt,name=cost_report_config_id,json=costReportConfigId,proto3" json:"cost_report_config_id,omitempty"`
	CostReportId       string          `protobuf:"bytes,25,opt,name=cost_report_id,json=costReportId,proto3" json:"cost_report_id,omitempty"`
	DataSourceId       string          `protobuf:"bytes,26,opt,name=data_source_id,json=dataSourceId,proto3" json:"data_source_id,omitempty"`
	ServiceAccountId   string          `protobuf:"bytes,27,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty"`
	CreatedAt          string          `protobuf:"bytes,31,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *CostReportDataInfo) Reset() {
	*x = CostReportDataInfo{}
	mi := &file_spaceone_api_cost_analysis_v1_cost_report_data_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CostReportDataInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CostReportDataInfo) ProtoMessage() {}

func (x *CostReportDataInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_cost_report_data_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CostReportDataInfo.ProtoReflect.Descriptor instead.
func (*CostReportDataInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_cost_report_data_proto_rawDescGZIP(), []int{2}
}

func (x *CostReportDataInfo) GetCostReportDataId() string {
	if x != nil {
		return x.CostReportDataId
	}
	return ""
}

func (x *CostReportDataInfo) GetCost() *_struct.Struct {
	if x != nil {
		return x.Cost
	}
	return nil
}

func (x *CostReportDataInfo) GetCostReportName() string {
	if x != nil {
		return x.CostReportName
	}
	return ""
}

func (x *CostReportDataInfo) GetIssueDate() string {
	if x != nil {
		return x.IssueDate
	}
	return ""
}

func (x *CostReportDataInfo) GetReportYear() string {
	if x != nil {
		return x.ReportYear
	}
	return ""
}

func (x *CostReportDataInfo) GetReportMonth() string {
	if x != nil {
		return x.ReportMonth
	}
	return ""
}

func (x *CostReportDataInfo) GetIsConfirmed() bool {
	if x != nil {
		return x.IsConfirmed
	}
	return false
}

func (x *CostReportDataInfo) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *CostReportDataInfo) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

func (x *CostReportDataInfo) GetServiceAccountName() string {
	if x != nil {
		return x.ServiceAccountName
	}
	return ""
}

func (x *CostReportDataInfo) GetDataSourceName() string {
	if x != nil {
		return x.DataSourceName
	}
	return ""
}

func (x *CostReportDataInfo) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *CostReportDataInfo) GetWorkspaceName() string {
	if x != nil {
		return x.WorkspaceName
	}
	return ""
}

func (x *CostReportDataInfo) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *CostReportDataInfo) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *CostReportDataInfo) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *CostReportDataInfo) GetCostReportConfigId() string {
	if x != nil {
		return x.CostReportConfigId
	}
	return ""
}

func (x *CostReportDataInfo) GetCostReportId() string {
	if x != nil {
		return x.CostReportId
	}
	return ""
}

func (x *CostReportDataInfo) GetDataSourceId() string {
	if x != nil {
		return x.DataSourceId
	}
	return ""
}

func (x *CostReportDataInfo) GetServiceAccountId() string {
	if x != nil {
		return x.ServiceAccountId
	}
	return ""
}

func (x *CostReportDataInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type CostReportsDataInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results    []*CostReportDataInfo `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	TotalCount int32                 `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *CostReportsDataInfo) Reset() {
	*x = CostReportsDataInfo{}
	mi := &file_spaceone_api_cost_analysis_v1_cost_report_data_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CostReportsDataInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CostReportsDataInfo) ProtoMessage() {}

func (x *CostReportsDataInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_cost_report_data_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CostReportsDataInfo.ProtoReflect.Descriptor instead.
func (*CostReportsDataInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_cost_report_data_proto_rawDescGZIP(), []int{3}
}

func (x *CostReportsDataInfo) GetResults() []*CostReportDataInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *CostReportsDataInfo) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type CostReportDataStatQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query *v2.StatisticsQuery `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	CostReportId string `protobuf:"bytes,21,opt,name=cost_report_id,json=costReportId,proto3" json:"cost_report_id,omitempty"`
	// +optional
	CostReportConfigId string `protobuf:"bytes,22,opt,name=cost_report_config_id,json=costReportConfigId,proto3" json:"cost_report_config_id,omitempty"`
}

func (x *CostReportDataStatQuery) Reset() {
	*x = CostReportDataStatQuery{}
	mi := &file_spaceone_api_cost_analysis_v1_cost_report_data_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CostReportDataStatQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CostReportDataStatQuery) ProtoMessage() {}

func (x *CostReportDataStatQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_cost_report_data_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CostReportDataStatQuery.ProtoReflect.Descriptor instead.
func (*CostReportDataStatQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_cost_report_data_proto_rawDescGZIP(), []int{4}
}

func (x *CostReportDataStatQuery) GetQuery() *v2.StatisticsQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *CostReportDataStatQuery) GetCostReportId() string {
	if x != nil {
		return x.CostReportId
	}
	return ""
}

func (x *CostReportDataStatQuery) GetCostReportConfigId() string {
	if x != nil {
		return x.CostReportConfigId
	}
	return ""
}

var File_spaceone_api_cost_analysis_v1_cost_report_data_proto protoreflect.FileDescriptor

var file_spaceone_api_cost_analysis_v1_cost_report_data_proto_rawDesc = []byte{
	0x0a, 0x34, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x76, 0x32, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x27, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf2, 0x02, 0x0a, 0x13, 0x43, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x31, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x2d, 0x0a, 0x13, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x63, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12,
	0x31, 0x0a, 0x15, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12,
	0x63, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x22, 0xfa,
	0x02, 0x0a, 0x1a, 0x43, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x42, 0x0a,
	0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x32, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x41, 0x6e,
	0x61, 0x6c, 0x79, 0x7a, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69,
	0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x15,
	0x63, 0x6f, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x63, 0x6f, 0x73,
	0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x64, 0x12,
	0x24, 0x0a, 0x0e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64,
	0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x22, 0xa7, 0x06, 0x0a, 0x12,
	0x43, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x2d, 0x0a, 0x13, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x63, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x49,
	0x64, 0x12, 0x2b, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x12, 0x28,
	0x0a, 0x10, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x73, 0x75,
	0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x5f, 0x79, 0x65, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x59, 0x65, 0x61, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x69,
	0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77,
	0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x15, 0x63, 0x6f, 0x73,
	0x74, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f,
	0x69, 0x64, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x63, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e,
	0x63, 0x6f, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x19,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x61,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x1b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x83, 0x01, 0x0a, 0x13, 0x43, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x44, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x4b, 0x0a,
	0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xaf, 0x01, 0x0a, 0x17,
	0x43, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74,
	0x61, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x3b, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x15, 0x63, 0x6f,
	0x73, 0x74, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x5f, 0x69, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x63, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x64, 0x32, 0xda, 0x03,
	0x0a, 0x0e, 0x43, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x12, 0xa2, 0x01, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x32, 0x2e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x32, 0x2e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73,
	0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x44, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x3a, 0x01, 0x2a, 0x22, 0x27, 0x2f, 0x63,
	0x6f, 0x73, 0x74, 0x2d, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x73, 0x74, 0x2d, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2d, 0x64, 0x61, 0x74, 0x61,
	0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x94, 0x01, 0x0a, 0x07, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a,
	0x65, 0x12, 0x39, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x35, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2f, 0x3a, 0x01, 0x2a,
	0x22, 0x2a, 0x2f, 0x63, 0x6f, 0x73, 0x74, 0x2d, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x73, 0x74, 0x2d, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2d,
	0x64, 0x61, 0x74, 0x61, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x12, 0x8b, 0x01, 0x0a,
	0x04, 0x73, 0x74, 0x61, 0x74, 0x12, 0x36, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x61, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x3a, 0x01,
	0x2a, 0x22, 0x27, 0x2f, 0x63, 0x6f, 0x73, 0x74, 0x2d, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x73, 0x74, 0x2d, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2d, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f,
	0x72, 0x65, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x2f,
	0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spaceone_api_cost_analysis_v1_cost_report_data_proto_rawDescOnce sync.Once
	file_spaceone_api_cost_analysis_v1_cost_report_data_proto_rawDescData = file_spaceone_api_cost_analysis_v1_cost_report_data_proto_rawDesc
)

func file_spaceone_api_cost_analysis_v1_cost_report_data_proto_rawDescGZIP() []byte {
	file_spaceone_api_cost_analysis_v1_cost_report_data_proto_rawDescOnce.Do(func() {
		file_spaceone_api_cost_analysis_v1_cost_report_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaceone_api_cost_analysis_v1_cost_report_data_proto_rawDescData)
	})
	return file_spaceone_api_cost_analysis_v1_cost_report_data_proto_rawDescData
}

var file_spaceone_api_cost_analysis_v1_cost_report_data_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_spaceone_api_cost_analysis_v1_cost_report_data_proto_goTypes = []any{
	(*CostReportDataQuery)(nil),        // 0: spaceone.api.cost_analysis.v1.CostReportDataQuery
	(*CostReportDataAnalyzeQuery)(nil), // 1: spaceone.api.cost_analysis.v1.CostReportDataAnalyzeQuery
	(*CostReportDataInfo)(nil),         // 2: spaceone.api.cost_analysis.v1.CostReportDataInfo
	(*CostReportsDataInfo)(nil),        // 3: spaceone.api.cost_analysis.v1.CostReportsDataInfo
	(*CostReportDataStatQuery)(nil),    // 4: spaceone.api.cost_analysis.v1.CostReportDataStatQuery
	(*v2.Query)(nil),                   // 5: spaceone.api.core.v2.Query
	(*v2.TimeSeriesAnalyzeQuery)(nil),  // 6: spaceone.api.core.v2.TimeSeriesAnalyzeQuery
	(*_struct.Struct)(nil),             // 7: google.protobuf.Struct
	(*v2.StatisticsQuery)(nil),         // 8: spaceone.api.core.v2.StatisticsQuery
}
var file_spaceone_api_cost_analysis_v1_cost_report_data_proto_depIdxs = []int32{
	5, // 0: spaceone.api.cost_analysis.v1.CostReportDataQuery.query:type_name -> spaceone.api.core.v2.Query
	6, // 1: spaceone.api.cost_analysis.v1.CostReportDataAnalyzeQuery.query:type_name -> spaceone.api.core.v2.TimeSeriesAnalyzeQuery
	7, // 2: spaceone.api.cost_analysis.v1.CostReportDataInfo.cost:type_name -> google.protobuf.Struct
	2, // 3: spaceone.api.cost_analysis.v1.CostReportsDataInfo.results:type_name -> spaceone.api.cost_analysis.v1.CostReportDataInfo
	8, // 4: spaceone.api.cost_analysis.v1.CostReportDataStatQuery.query:type_name -> spaceone.api.core.v2.StatisticsQuery
	0, // 5: spaceone.api.cost_analysis.v1.CostReportData.list:input_type -> spaceone.api.cost_analysis.v1.CostReportDataQuery
	1, // 6: spaceone.api.cost_analysis.v1.CostReportData.analyze:input_type -> spaceone.api.cost_analysis.v1.CostReportDataAnalyzeQuery
	4, // 7: spaceone.api.cost_analysis.v1.CostReportData.stat:input_type -> spaceone.api.cost_analysis.v1.CostReportDataStatQuery
	3, // 8: spaceone.api.cost_analysis.v1.CostReportData.list:output_type -> spaceone.api.cost_analysis.v1.CostReportsDataInfo
	7, // 9: spaceone.api.cost_analysis.v1.CostReportData.analyze:output_type -> google.protobuf.Struct
	7, // 10: spaceone.api.cost_analysis.v1.CostReportData.stat:output_type -> google.protobuf.Struct
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_spaceone_api_cost_analysis_v1_cost_report_data_proto_init() }
func file_spaceone_api_cost_analysis_v1_cost_report_data_proto_init() {
	if File_spaceone_api_cost_analysis_v1_cost_report_data_proto != nil {
		return
	}
	file_spaceone_api_cost_analysis_v1_job_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spaceone_api_cost_analysis_v1_cost_report_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_cost_analysis_v1_cost_report_data_proto_goTypes,
		DependencyIndexes: file_spaceone_api_cost_analysis_v1_cost_report_data_proto_depIdxs,
		MessageInfos:      file_spaceone_api_cost_analysis_v1_cost_report_data_proto_msgTypes,
	}.Build()
	File_spaceone_api_cost_analysis_v1_cost_report_data_proto = out.File
	file_spaceone_api_cost_analysis_v1_cost_report_data_proto_rawDesc = nil
	file_spaceone_api_cost_analysis_v1_cost_report_data_proto_goTypes = nil
	file_spaceone_api_cost_analysis_v1_cost_report_data_proto_depIdxs = nil
}
