// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: api/go_load.proto

package go_load

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DownloadType int32

const (
	DownloadType_DOWNLOAD_TYPE_UNSPECIFIED DownloadType = 0
	DownloadType_DOWNLOAD_TYPE_HTTP        DownloadType = 1
)

// Enum value maps for DownloadType.
var (
	DownloadType_name = map[int32]string{
		0: "DOWNLOAD_TYPE_UNSPECIFIED",
		1: "DOWNLOAD_TYPE_HTTP",
	}
	DownloadType_value = map[string]int32{
		"DOWNLOAD_TYPE_UNSPECIFIED": 0,
		"DOWNLOAD_TYPE_HTTP":        1,
	}
)

func (x DownloadType) Enum() *DownloadType {
	p := new(DownloadType)
	*p = x
	return p
}

func (x DownloadType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DownloadType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_go_load_proto_enumTypes[0].Descriptor()
}

func (DownloadType) Type() protoreflect.EnumType {
	return &file_api_go_load_proto_enumTypes[0]
}

func (x DownloadType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DownloadType.Descriptor instead.
func (DownloadType) EnumDescriptor() ([]byte, []int) {
	return file_api_go_load_proto_rawDescGZIP(), []int{0}
}

type DownloadStatus int32

const (
	DownloadStatus_DOWNLOAD_STATUS_UNSPECIFIED DownloadStatus = 0
	DownloadStatus_DOWNLOAD_STATUS_PENDING     DownloadStatus = 1
	DownloadStatus_DOWNLOAD_STATUS_DOWNLOADING DownloadStatus = 2
	DownloadStatus_DOWNLOAD_STATUS_FAILED      DownloadStatus = 3
	DownloadStatus_DOWNLOAD_STATUS_SUCCESS     DownloadStatus = 4
)

// Enum value maps for DownloadStatus.
var (
	DownloadStatus_name = map[int32]string{
		0: "DOWNLOAD_STATUS_UNSPECIFIED",
		1: "DOWNLOAD_STATUS_PENDING",
		2: "DOWNLOAD_STATUS_DOWNLOADING",
		3: "DOWNLOAD_STATUS_FAILED",
		4: "DOWNLOAD_STATUS_SUCCESS",
	}
	DownloadStatus_value = map[string]int32{
		"DOWNLOAD_STATUS_UNSPECIFIED": 0,
		"DOWNLOAD_STATUS_PENDING":     1,
		"DOWNLOAD_STATUS_DOWNLOADING": 2,
		"DOWNLOAD_STATUS_FAILED":      3,
		"DOWNLOAD_STATUS_SUCCESS":     4,
	}
)

func (x DownloadStatus) Enum() *DownloadStatus {
	p := new(DownloadStatus)
	*p = x
	return p
}

func (x DownloadStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DownloadStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_go_load_proto_enumTypes[1].Descriptor()
}

func (DownloadStatus) Type() protoreflect.EnumType {
	return &file_api_go_load_proto_enumTypes[1]
}

func (x DownloadStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DownloadStatus.Descriptor instead.
func (DownloadStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_go_load_proto_rawDescGZIP(), []int{1}
}

type Account struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AccountName   string                 `protobuf:"bytes,2,opt,name=account_name,json=accountName,proto3" json:"account_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Account) Reset() {
	*x = Account{}
	mi := &file_api_go_load_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_api_go_load_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_api_go_load_proto_rawDescGZIP(), []int{0}
}

func (x *Account) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Account) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

type DownloadTask struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Id             uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OfAccount      *Account               `protobuf:"bytes,2,opt,name=of_account,json=ofAccount,proto3" json:"of_account,omitempty"`
	DownloadType   DownloadType           `protobuf:"varint,3,opt,name=download_type,json=downloadType,proto3,enum=go_load.DownloadType" json:"download_type,omitempty"`
	Url            string                 `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	DownloadStatus DownloadStatus         `protobuf:"varint,5,opt,name=download_status,json=downloadStatus,proto3,enum=go_load.DownloadStatus" json:"download_status,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *DownloadTask) Reset() {
	*x = DownloadTask{}
	mi := &file_api_go_load_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadTask) ProtoMessage() {}

func (x *DownloadTask) ProtoReflect() protoreflect.Message {
	mi := &file_api_go_load_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadTask.ProtoReflect.Descriptor instead.
func (*DownloadTask) Descriptor() ([]byte, []int) {
	return file_api_go_load_proto_rawDescGZIP(), []int{1}
}

func (x *DownloadTask) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DownloadTask) GetOfAccount() *Account {
	if x != nil {
		return x.OfAccount
	}
	return nil
}

func (x *DownloadTask) GetDownloadType() DownloadType {
	if x != nil {
		return x.DownloadType
	}
	return DownloadType_DOWNLOAD_TYPE_UNSPECIFIED
}

func (x *DownloadTask) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *DownloadTask) GetDownloadStatus() DownloadStatus {
	if x != nil {
		return x.DownloadStatus
	}
	return DownloadStatus_DOWNLOAD_STATUS_UNSPECIFIED
}

type CreateAccountRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAccountRequest) Reset() {
	*x = CreateAccountRequest{}
	mi := &file_api_go_load_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountRequest) ProtoMessage() {}

func (x *CreateAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_go_load_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountRequest.ProtoReflect.Descriptor instead.
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return file_api_go_load_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAccountRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateAccountRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CreateAccountResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        uint64                 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAccountResponse) Reset() {
	*x = CreateAccountResponse{}
	mi := &file_api_go_load_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountResponse) ProtoMessage() {}

func (x *CreateAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_go_load_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountResponse.ProtoReflect.Descriptor instead.
func (*CreateAccountResponse) Descriptor() ([]byte, []int) {
	return file_api_go_load_proto_rawDescGZIP(), []int{3}
}

func (x *CreateAccountResponse) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateAccountResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type CreateSessionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateSessionRequest) Reset() {
	*x = CreateSessionRequest{}
	mi := &file_api_go_load_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSessionRequest) ProtoMessage() {}

func (x *CreateSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_go_load_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSessionRequest.ProtoReflect.Descriptor instead.
func (*CreateSessionRequest) Descriptor() ([]byte, []int) {
	return file_api_go_load_proto_rawDescGZIP(), []int{4}
}

func (x *CreateSessionRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateSessionRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CreateSessionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateSessionResponse) Reset() {
	*x = CreateSessionResponse{}
	mi := &file_api_go_load_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSessionResponse) ProtoMessage() {}

func (x *CreateSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_go_load_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSessionResponse.ProtoReflect.Descriptor instead.
func (*CreateSessionResponse) Descriptor() ([]byte, []int) {
	return file_api_go_load_proto_rawDescGZIP(), []int{5}
}

func (x *CreateSessionResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type CreateDownloadTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	DownloadType  DownloadType           `protobuf:"varint,2,opt,name=download_type,json=downloadType,proto3,enum=go_load.DownloadType" json:"download_type,omitempty"`
	Url           string                 `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateDownloadTaskRequest) Reset() {
	*x = CreateDownloadTaskRequest{}
	mi := &file_api_go_load_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateDownloadTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDownloadTaskRequest) ProtoMessage() {}

func (x *CreateDownloadTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_go_load_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDownloadTaskRequest.ProtoReflect.Descriptor instead.
func (*CreateDownloadTaskRequest) Descriptor() ([]byte, []int) {
	return file_api_go_load_proto_rawDescGZIP(), []int{6}
}

func (x *CreateDownloadTaskRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CreateDownloadTaskRequest) GetDownloadType() DownloadType {
	if x != nil {
		return x.DownloadType
	}
	return DownloadType_DOWNLOAD_TYPE_UNSPECIFIED
}

func (x *CreateDownloadTaskRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type CreateDownloadTaskResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DownloadTask  *DownloadTask          `protobuf:"bytes,1,opt,name=download_task,json=downloadTask,proto3" json:"download_task,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateDownloadTaskResponse) Reset() {
	*x = CreateDownloadTaskResponse{}
	mi := &file_api_go_load_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateDownloadTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDownloadTaskResponse) ProtoMessage() {}

func (x *CreateDownloadTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_go_load_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDownloadTaskResponse.ProtoReflect.Descriptor instead.
func (*CreateDownloadTaskResponse) Descriptor() ([]byte, []int) {
	return file_api_go_load_proto_rawDescGZIP(), []int{7}
}

func (x *CreateDownloadTaskResponse) GetDownloadTask() *DownloadTask {
	if x != nil {
		return x.DownloadTask
	}
	return nil
}

type GetDownloadTaskListRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Offset        uint64                 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit         uint64                 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetDownloadTaskListRequest) Reset() {
	*x = GetDownloadTaskListRequest{}
	mi := &file_api_go_load_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDownloadTaskListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDownloadTaskListRequest) ProtoMessage() {}

func (x *GetDownloadTaskListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_go_load_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDownloadTaskListRequest.ProtoReflect.Descriptor instead.
func (*GetDownloadTaskListRequest) Descriptor() ([]byte, []int) {
	return file_api_go_load_proto_rawDescGZIP(), []int{8}
}

func (x *GetDownloadTaskListRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetDownloadTaskListRequest) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetDownloadTaskListRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetDownloadTaskListResponse struct {
	state                  protoimpl.MessageState `protogen:"open.v1"`
	DownloadTaskList       []*DownloadTask        `protobuf:"bytes,1,rep,name=download_task_list,json=downloadTaskList,proto3" json:"download_task_list,omitempty"`
	TotalDownloadTaskCount uint64                 `protobuf:"varint,2,opt,name=total_download_task_count,json=totalDownloadTaskCount,proto3" json:"total_download_task_count,omitempty"`
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *GetDownloadTaskListResponse) Reset() {
	*x = GetDownloadTaskListResponse{}
	mi := &file_api_go_load_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDownloadTaskListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDownloadTaskListResponse) ProtoMessage() {}

func (x *GetDownloadTaskListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_go_load_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDownloadTaskListResponse.ProtoReflect.Descriptor instead.
func (*GetDownloadTaskListResponse) Descriptor() ([]byte, []int) {
	return file_api_go_load_proto_rawDescGZIP(), []int{9}
}

func (x *GetDownloadTaskListResponse) GetDownloadTaskList() []*DownloadTask {
	if x != nil {
		return x.DownloadTaskList
	}
	return nil
}

func (x *GetDownloadTaskListResponse) GetTotalDownloadTaskCount() uint64 {
	if x != nil {
		return x.TotalDownloadTaskCount
	}
	return 0
}

type UpdateDownloadTaskRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	DownloadTaskId uint64                 `protobuf:"varint,1,opt,name=download_task_id,json=downloadTaskId,proto3" json:"download_task_id,omitempty"`
	Url            string                 `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *UpdateDownloadTaskRequest) Reset() {
	*x = UpdateDownloadTaskRequest{}
	mi := &file_api_go_load_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateDownloadTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDownloadTaskRequest) ProtoMessage() {}

func (x *UpdateDownloadTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_go_load_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDownloadTaskRequest.ProtoReflect.Descriptor instead.
func (*UpdateDownloadTaskRequest) Descriptor() ([]byte, []int) {
	return file_api_go_load_proto_rawDescGZIP(), []int{10}
}

func (x *UpdateDownloadTaskRequest) GetDownloadTaskId() uint64 {
	if x != nil {
		return x.DownloadTaskId
	}
	return 0
}

func (x *UpdateDownloadTaskRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type UpdateDownloadTaskResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DownloadTask  *DownloadTask          `protobuf:"bytes,1,opt,name=download_task,json=downloadTask,proto3" json:"download_task,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateDownloadTaskResponse) Reset() {
	*x = UpdateDownloadTaskResponse{}
	mi := &file_api_go_load_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateDownloadTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDownloadTaskResponse) ProtoMessage() {}

func (x *UpdateDownloadTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_go_load_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDownloadTaskResponse.ProtoReflect.Descriptor instead.
func (*UpdateDownloadTaskResponse) Descriptor() ([]byte, []int) {
	return file_api_go_load_proto_rawDescGZIP(), []int{11}
}

func (x *UpdateDownloadTaskResponse) GetDownloadTask() *DownloadTask {
	if x != nil {
		return x.DownloadTask
	}
	return nil
}

type DeleteDownloadTaskRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	DownloadTaskId uint64                 `protobuf:"varint,1,opt,name=download_task_id,json=downloadTaskId,proto3" json:"download_task_id,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *DeleteDownloadTaskRequest) Reset() {
	*x = DeleteDownloadTaskRequest{}
	mi := &file_api_go_load_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteDownloadTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDownloadTaskRequest) ProtoMessage() {}

func (x *DeleteDownloadTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_go_load_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDownloadTaskRequest.ProtoReflect.Descriptor instead.
func (*DeleteDownloadTaskRequest) Descriptor() ([]byte, []int) {
	return file_api_go_load_proto_rawDescGZIP(), []int{12}
}

func (x *DeleteDownloadTaskRequest) GetDownloadTaskId() uint64 {
	if x != nil {
		return x.DownloadTaskId
	}
	return 0
}

type DeleteDownloadTaskResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteDownloadTaskResponse) Reset() {
	*x = DeleteDownloadTaskResponse{}
	mi := &file_api_go_load_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteDownloadTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDownloadTaskResponse) ProtoMessage() {}

func (x *DeleteDownloadTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_go_load_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDownloadTaskResponse.ProtoReflect.Descriptor instead.
func (*DeleteDownloadTaskResponse) Descriptor() ([]byte, []int) {
	return file_api_go_load_proto_rawDescGZIP(), []int{13}
}

type GetDownloadTaskFileRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	DownloadTaskId uint64                 `protobuf:"varint,1,opt,name=download_task_id,json=downloadTaskId,proto3" json:"download_task_id,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *GetDownloadTaskFileRequest) Reset() {
	*x = GetDownloadTaskFileRequest{}
	mi := &file_api_go_load_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDownloadTaskFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDownloadTaskFileRequest) ProtoMessage() {}

func (x *GetDownloadTaskFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_go_load_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDownloadTaskFileRequest.ProtoReflect.Descriptor instead.
func (*GetDownloadTaskFileRequest) Descriptor() ([]byte, []int) {
	return file_api_go_load_proto_rawDescGZIP(), []int{14}
}

func (x *GetDownloadTaskFileRequest) GetDownloadTaskId() uint64 {
	if x != nil {
		return x.DownloadTaskId
	}
	return 0
}

type GetDownloadTaskFileResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          []byte                 `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetDownloadTaskFileResponse) Reset() {
	*x = GetDownloadTaskFileResponse{}
	mi := &file_api_go_load_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDownloadTaskFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDownloadTaskFileResponse) ProtoMessage() {}

func (x *GetDownloadTaskFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_go_load_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDownloadTaskFileResponse.ProtoReflect.Descriptor instead.
func (*GetDownloadTaskFileResponse) Descriptor() ([]byte, []int) {
	return file_api_go_load_proto_rawDescGZIP(), []int{15}
}

func (x *GetDownloadTaskFileResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_api_go_load_proto protoreflect.FileDescriptor

const file_api_go_load_proto_rawDesc = "" +
	"\n" +
	"\x11api/go_load.proto\x12\ago_load\"<\n" +
	"\aAccount\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x04R\x02id\x12!\n" +
	"\faccount_name\x18\x02 \x01(\tR\vaccountName\"\xdf\x01\n" +
	"\fDownloadTask\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x04R\x02id\x12/\n" +
	"\n" +
	"of_account\x18\x02 \x01(\v2\x10.go_load.AccountR\tofAccount\x12:\n" +
	"\rdownload_type\x18\x03 \x01(\x0e2\x15.go_load.DownloadTypeR\fdownloadType\x12\x10\n" +
	"\x03url\x18\x04 \x01(\tR\x03url\x12@\n" +
	"\x0fdownload_status\x18\x05 \x01(\x0e2\x17.go_load.DownloadStatusR\x0edownloadStatus\"N\n" +
	"\x14CreateAccountRequest\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\"L\n" +
	"\x15CreateAccountResponse\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\x04R\x06userId\x12\x1a\n" +
	"\busername\x18\x02 \x01(\tR\busername\"N\n" +
	"\x14CreateSessionRequest\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\"-\n" +
	"\x15CreateSessionResponse\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\"\x7f\n" +
	"\x19CreateDownloadTaskRequest\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\x12:\n" +
	"\rdownload_type\x18\x02 \x01(\x0e2\x15.go_load.DownloadTypeR\fdownloadType\x12\x10\n" +
	"\x03url\x18\x03 \x01(\tR\x03url\"X\n" +
	"\x1aCreateDownloadTaskResponse\x12:\n" +
	"\rdownload_task\x18\x01 \x01(\v2\x15.go_load.DownloadTaskR\fdownloadTask\"`\n" +
	"\x1aGetDownloadTaskListRequest\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\x12\x16\n" +
	"\x06offset\x18\x02 \x01(\x04R\x06offset\x12\x14\n" +
	"\x05limit\x18\x03 \x01(\x04R\x05limit\"\x9d\x01\n" +
	"\x1bGetDownloadTaskListResponse\x12C\n" +
	"\x12download_task_list\x18\x01 \x03(\v2\x15.go_load.DownloadTaskR\x10downloadTaskList\x129\n" +
	"\x19total_download_task_count\x18\x02 \x01(\x04R\x16totalDownloadTaskCount\"W\n" +
	"\x19UpdateDownloadTaskRequest\x12(\n" +
	"\x10download_task_id\x18\x01 \x01(\x04R\x0edownloadTaskId\x12\x10\n" +
	"\x03url\x18\x02 \x01(\tR\x03url\"X\n" +
	"\x1aUpdateDownloadTaskResponse\x12:\n" +
	"\rdownload_task\x18\x01 \x01(\v2\x15.go_load.DownloadTaskR\fdownloadTask\"E\n" +
	"\x19DeleteDownloadTaskRequest\x12(\n" +
	"\x10download_task_id\x18\x01 \x01(\x04R\x0edownloadTaskId\"\x1c\n" +
	"\x1aDeleteDownloadTaskResponse\"F\n" +
	"\x1aGetDownloadTaskFileRequest\x12(\n" +
	"\x10download_task_id\x18\x01 \x01(\x04R\x0edownloadTaskId\"1\n" +
	"\x1bGetDownloadTaskFileResponse\x12\x12\n" +
	"\x04data\x18\x01 \x01(\fR\x04data*E\n" +
	"\fDownloadType\x12\x1d\n" +
	"\x19DOWNLOAD_TYPE_UNSPECIFIED\x10\x00\x12\x16\n" +
	"\x12DOWNLOAD_TYPE_HTTP\x10\x01*\xa8\x01\n" +
	"\x0eDownloadStatus\x12\x1f\n" +
	"\x1bDOWNLOAD_STATUS_UNSPECIFIED\x10\x00\x12\x1b\n" +
	"\x17DOWNLOAD_STATUS_PENDING\x10\x01\x12\x1f\n" +
	"\x1bDOWNLOAD_STATUS_DOWNLOADING\x10\x02\x12\x1a\n" +
	"\x16DOWNLOAD_STATUS_FAILED\x10\x03\x12\x1b\n" +
	"\x17DOWNLOAD_STATUS_SUCCESS\x10\x042\xa0\x05\n" +
	"\rGoLoadService\x12P\n" +
	"\rCreateAccount\x12\x1d.go_load.CreateAccountRequest\x1a\x1e.go_load.CreateAccountResponse\"\x00\x12P\n" +
	"\rCreateSession\x12\x1d.go_load.CreateSessionRequest\x1a\x1e.go_load.CreateSessionResponse\"\x00\x12_\n" +
	"\x12CreateDownloadTask\x12\".go_load.CreateDownloadTaskRequest\x1a#.go_load.CreateDownloadTaskResponse\"\x00\x12b\n" +
	"\x13GetDownloadTaskList\x12#.go_load.GetDownloadTaskListRequest\x1a$.go_load.GetDownloadTaskListResponse\"\x00\x12_\n" +
	"\x12UpdateDownloadTask\x12\".go_load.UpdateDownloadTaskRequest\x1a#.go_load.UpdateDownloadTaskResponse\"\x00\x12_\n" +
	"\x12DeleteDownloadTask\x12\".go_load.DeleteDownloadTaskRequest\x1a#.go_load.DeleteDownloadTaskResponse\"\x00\x12d\n" +
	"\x13GetDownloadTaskFile\x12#.go_load.GetDownloadTaskFileRequest\x1a$.go_load.GetDownloadTaskFileResponse\"\x000\x01B\x0eZ\fgrpc/go_loadb\x06proto3"

var (
	file_api_go_load_proto_rawDescOnce sync.Once
	file_api_go_load_proto_rawDescData []byte
)

func file_api_go_load_proto_rawDescGZIP() []byte {
	file_api_go_load_proto_rawDescOnce.Do(func() {
		file_api_go_load_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_go_load_proto_rawDesc), len(file_api_go_load_proto_rawDesc)))
	})
	return file_api_go_load_proto_rawDescData
}

var file_api_go_load_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_go_load_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_api_go_load_proto_goTypes = []any{
	(DownloadType)(0),                   // 0: go_load.DownloadType
	(DownloadStatus)(0),                 // 1: go_load.DownloadStatus
	(*Account)(nil),                     // 2: go_load.Account
	(*DownloadTask)(nil),                // 3: go_load.DownloadTask
	(*CreateAccountRequest)(nil),        // 4: go_load.CreateAccountRequest
	(*CreateAccountResponse)(nil),       // 5: go_load.CreateAccountResponse
	(*CreateSessionRequest)(nil),        // 6: go_load.CreateSessionRequest
	(*CreateSessionResponse)(nil),       // 7: go_load.CreateSessionResponse
	(*CreateDownloadTaskRequest)(nil),   // 8: go_load.CreateDownloadTaskRequest
	(*CreateDownloadTaskResponse)(nil),  // 9: go_load.CreateDownloadTaskResponse
	(*GetDownloadTaskListRequest)(nil),  // 10: go_load.GetDownloadTaskListRequest
	(*GetDownloadTaskListResponse)(nil), // 11: go_load.GetDownloadTaskListResponse
	(*UpdateDownloadTaskRequest)(nil),   // 12: go_load.UpdateDownloadTaskRequest
	(*UpdateDownloadTaskResponse)(nil),  // 13: go_load.UpdateDownloadTaskResponse
	(*DeleteDownloadTaskRequest)(nil),   // 14: go_load.DeleteDownloadTaskRequest
	(*DeleteDownloadTaskResponse)(nil),  // 15: go_load.DeleteDownloadTaskResponse
	(*GetDownloadTaskFileRequest)(nil),  // 16: go_load.GetDownloadTaskFileRequest
	(*GetDownloadTaskFileResponse)(nil), // 17: go_load.GetDownloadTaskFileResponse
}
var file_api_go_load_proto_depIdxs = []int32{
	2,  // 0: go_load.DownloadTask.of_account:type_name -> go_load.Account
	0,  // 1: go_load.DownloadTask.download_type:type_name -> go_load.DownloadType
	1,  // 2: go_load.DownloadTask.download_status:type_name -> go_load.DownloadStatus
	0,  // 3: go_load.CreateDownloadTaskRequest.download_type:type_name -> go_load.DownloadType
	3,  // 4: go_load.CreateDownloadTaskResponse.download_task:type_name -> go_load.DownloadTask
	3,  // 5: go_load.GetDownloadTaskListResponse.download_task_list:type_name -> go_load.DownloadTask
	3,  // 6: go_load.UpdateDownloadTaskResponse.download_task:type_name -> go_load.DownloadTask
	4,  // 7: go_load.GoLoadService.CreateAccount:input_type -> go_load.CreateAccountRequest
	6,  // 8: go_load.GoLoadService.CreateSession:input_type -> go_load.CreateSessionRequest
	8,  // 9: go_load.GoLoadService.CreateDownloadTask:input_type -> go_load.CreateDownloadTaskRequest
	10, // 10: go_load.GoLoadService.GetDownloadTaskList:input_type -> go_load.GetDownloadTaskListRequest
	12, // 11: go_load.GoLoadService.UpdateDownloadTask:input_type -> go_load.UpdateDownloadTaskRequest
	14, // 12: go_load.GoLoadService.DeleteDownloadTask:input_type -> go_load.DeleteDownloadTaskRequest
	16, // 13: go_load.GoLoadService.GetDownloadTaskFile:input_type -> go_load.GetDownloadTaskFileRequest
	5,  // 14: go_load.GoLoadService.CreateAccount:output_type -> go_load.CreateAccountResponse
	7,  // 15: go_load.GoLoadService.CreateSession:output_type -> go_load.CreateSessionResponse
	9,  // 16: go_load.GoLoadService.CreateDownloadTask:output_type -> go_load.CreateDownloadTaskResponse
	11, // 17: go_load.GoLoadService.GetDownloadTaskList:output_type -> go_load.GetDownloadTaskListResponse
	13, // 18: go_load.GoLoadService.UpdateDownloadTask:output_type -> go_load.UpdateDownloadTaskResponse
	15, // 19: go_load.GoLoadService.DeleteDownloadTask:output_type -> go_load.DeleteDownloadTaskResponse
	17, // 20: go_load.GoLoadService.GetDownloadTaskFile:output_type -> go_load.GetDownloadTaskFileResponse
	14, // [14:21] is the sub-list for method output_type
	7,  // [7:14] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_api_go_load_proto_init() }
func file_api_go_load_proto_init() {
	if File_api_go_load_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_go_load_proto_rawDesc), len(file_api_go_load_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_go_load_proto_goTypes,
		DependencyIndexes: file_api_go_load_proto_depIdxs,
		EnumInfos:         file_api_go_load_proto_enumTypes,
		MessageInfos:      file_api_go_load_proto_msgTypes,
	}.Build()
	File_api_go_load_proto = out.File
	file_api_go_load_proto_goTypes = nil
	file_api_go_load_proto_depIdxs = nil
}
