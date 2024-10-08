// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.2
// source: policy.proto

package policy

import (
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

type RuleOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID     string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Nickname   string `protobuf:"bytes,2,opt,name=Nickname,proto3" json:"Nickname,omitempty"`
	Telephone  string `protobuf:"bytes,3,opt,name=Telephone,proto3" json:"Telephone,omitempty"`
	Birthday   string `protobuf:"bytes,4,opt,name=Birthday,proto3" json:"Birthday,omitempty"`
	Gender     int64  `protobuf:"varint,5,opt,name=Gender,proto3" json:"Gender,omitempty"`
	Score      int64  `protobuf:"varint,6,opt,name=Score,proto3" json:"Score,omitempty"`
	Balance    int64  `protobuf:"varint,7,opt,name=Balance,proto3" json:"Balance,omitempty"`
	Channel    string `protobuf:"bytes,8,opt,name=Channel,proto3" json:"Channel,omitempty"`
	AppId      string `protobuf:"bytes,9,opt,name=AppId,proto3" json:"AppId,omitempty"`
	HavePay    bool   `protobuf:"varint,10,opt,name=HavePay,proto3" json:"HavePay,omitempty"`
	ReadTime   int64  `protobuf:"varint,11,opt,name=ReadTime,proto3" json:"ReadTime,omitempty"`
	ListenTime int64  `protobuf:"varint,12,opt,name=ListenTime,proto3" json:"ListenTime,omitempty"`
}

func (x *RuleOptions) Reset() {
	*x = RuleOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RuleOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuleOptions) ProtoMessage() {}

func (x *RuleOptions) ProtoReflect() protoreflect.Message {
	mi := &file_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuleOptions.ProtoReflect.Descriptor instead.
func (*RuleOptions) Descriptor() ([]byte, []int) {
	return file_policy_proto_rawDescGZIP(), []int{0}
}

func (x *RuleOptions) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *RuleOptions) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *RuleOptions) GetTelephone() string {
	if x != nil {
		return x.Telephone
	}
	return ""
}

func (x *RuleOptions) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *RuleOptions) GetGender() int64 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *RuleOptions) GetScore() int64 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *RuleOptions) GetBalance() int64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *RuleOptions) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *RuleOptions) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *RuleOptions) GetHavePay() bool {
	if x != nil {
		return x.HavePay
	}
	return false
}

func (x *RuleOptions) GetReadTime() int64 {
	if x != nil {
		return x.ReadTime
	}
	return 0
}

func (x *RuleOptions) GetListenTime() int64 {
	if x != nil {
		return x.ListenTime
	}
	return 0
}

type Input struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cate  string `protobuf:"bytes,1,opt,name=cate,proto3" json:"cate,omitempty"`
	Attr  string `protobuf:"bytes,2,opt,name=attr,proto3" json:"attr,omitempty"`
	Rule  string `protobuf:"bytes,3,opt,name=rule,proto3" json:"rule,omitempty"`
	ID    *int64 `protobuf:"varint,4,opt,name=ID,proto3,oneof" json:"ID,omitempty"`
	State *int64 `protobuf:"varint,5,opt,name=State,proto3,oneof" json:"State,omitempty"`
}

func (x *Input) Reset() {
	*x = Input{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Input) ProtoMessage() {}

func (x *Input) ProtoReflect() protoreflect.Message {
	mi := &file_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Input.ProtoReflect.Descriptor instead.
func (*Input) Descriptor() ([]byte, []int) {
	return file_policy_proto_rawDescGZIP(), []int{1}
}

func (x *Input) GetCate() string {
	if x != nil {
		return x.Cate
	}
	return ""
}

func (x *Input) GetAttr() string {
	if x != nil {
		return x.Attr
	}
	return ""
}

func (x *Input) GetRule() string {
	if x != nil {
		return x.Rule
	}
	return ""
}

func (x *Input) GetID() int64 {
	if x != nil && x.ID != nil {
		return *x.ID
	}
	return 0
}

func (x *Input) GetState() int64 {
	if x != nil && x.State != nil {
		return *x.State
	}
	return 0
}

type AddReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *AddReq) Reset() {
	*x = AddReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddReq) ProtoMessage() {}

func (x *AddReq) ProtoReflect() protoreflect.Message {
	mi := &file_policy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddReq.ProtoReflect.Descriptor instead.
func (*AddReq) Descriptor() ([]byte, []int) {
	return file_policy_proto_rawDescGZIP(), []int{2}
}

func (x *AddReq) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type AddRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *AddRep) Reset() {
	*x = AddRep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRep) ProtoMessage() {}

func (x *AddRep) ProtoReflect() protoreflect.Message {
	mi := &file_policy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRep.ProtoReflect.Descriptor instead.
func (*AddRep) Descriptor() ([]byte, []int) {
	return file_policy_proto_rawDescGZIP(), []int{3}
}

func (x *AddRep) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type UpdateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateResp) Reset() {
	*x = UpdateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResp) ProtoMessage() {}

func (x *UpdateResp) ProtoReflect() protoreflect.Message {
	mi := &file_policy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResp.ProtoReflect.Descriptor instead.
func (*UpdateResp) Descriptor() ([]byte, []int) {
	return file_policy_proto_rawDescGZIP(), []int{4}
}

type DeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID  int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Msg string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *DeleteReq) Reset() {
	*x = DeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteReq) ProtoMessage() {}

func (x *DeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_policy_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteReq.ProtoReflect.Descriptor instead.
func (*DeleteReq) Descriptor() ([]byte, []int) {
	return file_policy_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteReq) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *DeleteReq) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type CheckReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cate string       `protobuf:"bytes,1,opt,name=cate,proto3" json:"cate,omitempty"`
	User *RuleOptions `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *CheckReq) Reset() {
	*x = CheckReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckReq) ProtoMessage() {}

func (x *CheckReq) ProtoReflect() protoreflect.Message {
	mi := &file_policy_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckReq.ProtoReflect.Descriptor instead.
func (*CheckReq) Descriptor() ([]byte, []int) {
	return file_policy_proto_rawDescGZIP(), []int{6}
}

func (x *CheckReq) GetCate() string {
	if x != nil {
		return x.Cate
	}
	return ""
}

func (x *CheckReq) GetUser() *RuleOptions {
	if x != nil {
		return x.User
	}
	return nil
}

type CheckRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*Input `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *CheckRep) Reset() {
	*x = CheckRep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckRep) ProtoMessage() {}

func (x *CheckRep) ProtoReflect() protoreflect.Message {
	mi := &file_policy_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckRep.ProtoReflect.Descriptor instead.
func (*CheckRep) Descriptor() ([]byte, []int) {
	return file_policy_proto_rawDescGZIP(), []int{7}
}

func (x *CheckRep) GetList() []*Input {
	if x != nil {
		return x.List
	}
	return nil
}

type Placeholder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Placeholder) Reset() {
	*x = Placeholder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Placeholder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Placeholder) ProtoMessage() {}

func (x *Placeholder) ProtoReflect() protoreflect.Message {
	mi := &file_policy_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Placeholder.ProtoReflect.Descriptor instead.
func (*Placeholder) Descriptor() ([]byte, []int) {
	return file_policy_proto_rawDescGZIP(), []int{8}
}

var File_policy_proto protoreflect.FileDescriptor

var file_policy_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0xc9, 0x02, 0x0a, 0x0b, 0x52, 0x75, 0x6c, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a,
	0x0a, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x65,
	0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54,
	0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x42, 0x69, 0x72, 0x74,
	0x68, 0x64, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x42, 0x69, 0x72, 0x74,
	0x68, 0x64, 0x61, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x64, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x48, 0x61, 0x76, 0x65, 0x50, 0x61, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x48,
	0x61, 0x76, 0x65, 0x50, 0x61, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x61, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x52, 0x65, 0x61, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x22, 0x84, 0x01, 0x0a, 0x05, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x61, 0x74, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x61, 0x74, 0x74, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x61, 0x74, 0x74, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x12, 0x13, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x02, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a,
	0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x05,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x49, 0x44, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x18, 0x0a, 0x06, 0x41, 0x64, 0x64,
	0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x49, 0x44, 0x22, 0x18, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x22, 0x0c, 0x0a,
	0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x2d, 0x0a, 0x09, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x22, 0x47, 0x0a, 0x08, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x61, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x61, 0x74, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x22, 0x2d, 0x0a, 0x08, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x12,
	0x21, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x22, 0x0d, 0x0a, 0x0b, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x32, 0xba, 0x01, 0x0a, 0x06, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x24, 0x0a, 0x03,
	0x41, 0x64, 0x64, 0x12, 0x0d, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x41, 0x64, 0x64, 0x52,
	0x65, 0x70, 0x12, 0x2b, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x30, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x12, 0x2b, 0x0a, 0x05, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x10, 0x2e, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_policy_proto_rawDescOnce sync.Once
	file_policy_proto_rawDescData = file_policy_proto_rawDesc
)

func file_policy_proto_rawDescGZIP() []byte {
	file_policy_proto_rawDescOnce.Do(func() {
		file_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_policy_proto_rawDescData)
	})
	return file_policy_proto_rawDescData
}

var file_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_policy_proto_goTypes = []any{
	(*RuleOptions)(nil), // 0: policy.RuleOptions
	(*Input)(nil),       // 1: policy.Input
	(*AddReq)(nil),      // 2: policy.AddReq
	(*AddRep)(nil),      // 3: policy.AddRep
	(*UpdateResp)(nil),  // 4: policy.UpdateResp
	(*DeleteReq)(nil),   // 5: policy.DeleteReq
	(*CheckReq)(nil),    // 6: policy.CheckReq
	(*CheckRep)(nil),    // 7: policy.CheckRep
	(*Placeholder)(nil), // 8: policy.Placeholder
}
var file_policy_proto_depIdxs = []int32{
	0, // 0: policy.CheckReq.user:type_name -> policy.RuleOptions
	1, // 1: policy.CheckRep.list:type_name -> policy.Input
	1, // 2: policy.Policy.Add:input_type -> policy.Input
	1, // 3: policy.Policy.Update:input_type -> policy.Input
	5, // 4: policy.Policy.Delete:input_type -> policy.DeleteReq
	6, // 5: policy.Policy.Check:input_type -> policy.CheckReq
	3, // 6: policy.Policy.Add:output_type -> policy.AddRep
	4, // 7: policy.Policy.Update:output_type -> policy.UpdateResp
	8, // 8: policy.Policy.Delete:output_type -> policy.Placeholder
	7, // 9: policy.Policy.Check:output_type -> policy.CheckRep
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_policy_proto_init() }
func file_policy_proto_init() {
	if File_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_policy_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*RuleOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_policy_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Input); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_policy_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*AddReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_policy_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*AddRep); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_policy_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_policy_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_policy_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*CheckReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_policy_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*CheckRep); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_policy_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*Placeholder); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_policy_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_policy_proto_goTypes,
		DependencyIndexes: file_policy_proto_depIdxs,
		MessageInfos:      file_policy_proto_msgTypes,
	}.Build()
	File_policy_proto = out.File
	file_policy_proto_rawDesc = nil
	file_policy_proto_goTypes = nil
	file_policy_proto_depIdxs = nil
}
