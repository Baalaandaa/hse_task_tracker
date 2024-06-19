// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.0
// source: proto/gateway.proto

package model

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

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" fieldtag:"pk"`
	AuthorId string `protobuf:"bytes,2,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Title    string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Content  string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Status   int32  `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_proto_task_proto_rawDescGZIP(), []int{0}
}

func (x *Task) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Task) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

func (x *Task) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Task) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Task) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type CreateTaskInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorId string `protobuf:"bytes,1,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Title    string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content  string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Status   int32  `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CreateTaskInput) Reset() {
	*x = CreateTaskInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTaskInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskInput) ProtoMessage() {}

func (x *CreateTaskInput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskInput.ProtoReflect.Descriptor instead.
func (*CreateTaskInput) Descriptor() ([]byte, []int) {
	return file_proto_task_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTaskInput) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

func (x *CreateTaskInput) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateTaskInput) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CreateTaskInput) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type CreateTaskOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *CreateTaskOutput) Reset() {
	*x = CreateTaskOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTaskOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskOutput) ProtoMessage() {}

func (x *CreateTaskOutput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskOutput.ProtoReflect.Descriptor instead.
func (*CreateTaskOutput) Descriptor() ([]byte, []int) {
	return file_proto_task_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTaskOutput) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type DeleteTaskInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteTaskInput) Reset() {
	*x = DeleteTaskInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTaskInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskInput) ProtoMessage() {}

func (x *DeleteTaskInput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskInput.ProtoReflect.Descriptor instead.
func (*DeleteTaskInput) Descriptor() ([]byte, []int) {
	return file_proto_task_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteTaskInput) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteTaskOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteTaskOutput) Reset() {
	*x = DeleteTaskOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_task_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTaskOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskOutput) ProtoMessage() {}

func (x *DeleteTaskOutput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskOutput.ProtoReflect.Descriptor instead.
func (*DeleteTaskOutput) Descriptor() ([]byte, []int) {
	return file_proto_task_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteTaskOutput) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetTaskByIDInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTaskByIDInput) Reset() {
	*x = GetTaskByIDInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_task_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTaskByIDInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskByIDInput) ProtoMessage() {}

func (x *GetTaskByIDInput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskByIDInput.ProtoReflect.Descriptor instead.
func (*GetTaskByIDInput) Descriptor() ([]byte, []int) {
	return file_proto_task_proto_rawDescGZIP(), []int{5}
}

func (x *GetTaskByIDInput) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetTaskByIDOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task *Task `protobuf:"bytes,1,opt,name=gateway,proto3,oneof" json:"gateway,omitempty"`
}

func (x *GetTaskByIDOutput) Reset() {
	*x = GetTaskByIDOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_task_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTaskByIDOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskByIDOutput) ProtoMessage() {}

func (x *GetTaskByIDOutput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskByIDOutput.ProtoReflect.Descriptor instead.
func (*GetTaskByIDOutput) Descriptor() ([]byte, []int) {
	return file_proto_task_proto_rawDescGZIP(), []int{6}
}

func (x *GetTaskByIDOutput) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

type GetTasksInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *GetTasksInput) Reset() {
	*x = GetTasksInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_task_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTasksInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTasksInput) ProtoMessage() {}

func (x *GetTasksInput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTasksInput.ProtoReflect.Descriptor instead.
func (*GetTasksInput) Descriptor() ([]byte, []int) {
	return file_proto_task_proto_rawDescGZIP(), []int{7}
}

func (x *GetTasksInput) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetTasksInput) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type GetTasksOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32   `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Tasks []*Task `protobuf:"bytes,2,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *GetTasksOutput) Reset() {
	*x = GetTasksOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_task_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTasksOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTasksOutput) ProtoMessage() {}

func (x *GetTasksOutput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTasksOutput.ProtoReflect.Descriptor instead.
func (*GetTasksOutput) Descriptor() ([]byte, []int) {
	return file_proto_task_proto_rawDescGZIP(), []int{8}
}

func (x *GetTasksOutput) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetTasksOutput) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

var File_proto_task_proto protoreflect.FileDescriptor

var file_proto_task_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x22, 0x7b, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x76, 0x0a,
	0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2c, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x22, 0x21, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73,
	0x6b, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2c, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x22, 0x22, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x42,
	0x79, 0x49, 0x44, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x49, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x54,
	0x61, 0x73, 0x6b, 0x42, 0x79, 0x49, 0x44, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x2b, 0x0a,
	0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x61,
	0x73, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x48,
	0x00, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x74,
	0x61, 0x73, 0x6b, 0x22, 0x3d, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x22, 0x50, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x05, 0x74, 0x61,
	0x73, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x61, 0x73, 0x6b,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x74,
	0x61, 0x73, 0x6b, 0x73, 0x32, 0xfd, 0x02, 0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x12, 0x1d, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x1a, 0x1e, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73,
	0x6b, 0x12, 0x12, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x1a, 0x12, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0a, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x1d, 0x2e, 0x74, 0x61, 0x73, 0x6b,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x1e, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1e, 0x2e, 0x74, 0x61, 0x73, 0x6b,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b,
	0x42, 0x79, 0x49, 0x44, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x1e, 0x2e, 0x74, 0x61, 0x73, 0x6b,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b,
	0x42, 0x79, 0x49, 0x44, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x1b, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x1a, 0x1c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_task_proto_rawDescOnce sync.Once
	file_proto_task_proto_rawDescData = file_proto_task_proto_rawDesc
)

func file_proto_task_proto_rawDescGZIP() []byte {
	file_proto_task_proto_rawDescOnce.Do(func() {
		file_proto_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_task_proto_rawDescData)
	})
	return file_proto_task_proto_rawDescData
}

var file_proto_task_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_task_proto_goTypes = []interface{}{
	(*Task)(nil),              // 0: task_service.Task
	(*CreateTaskInput)(nil),   // 1: task_service.CreateTaskInput
	(*CreateTaskOutput)(nil),  // 2: task_service.CreateTaskOutput
	(*DeleteTaskInput)(nil),   // 3: task_service.DeleteTaskInput
	(*DeleteTaskOutput)(nil),  // 4: task_service.DeleteTaskOutput
	(*GetTaskByIDInput)(nil),  // 5: task_service.GetTaskByIDInput
	(*GetTaskByIDOutput)(nil), // 6: task_service.GetTaskByIDOutput
	(*GetTasksInput)(nil),     // 7: task_service.GetTasksInput
	(*GetTasksOutput)(nil),    // 8: task_service.GetTasksOutput
}
var file_proto_task_proto_depIdxs = []int32{
	0, // 0: task_service.GetTaskByIDOutput.gateway:type_name -> task_service.Task
	0, // 1: task_service.GetTasksOutput.tasks:type_name -> task_service.Task
	1, // 2: task_service.TaskService.CreateTask:input_type -> task_service.CreateTaskInput
	0, // 3: task_service.TaskService.UpdateTask:input_type -> task_service.Task
	3, // 4: task_service.TaskService.DeleteTask:input_type -> task_service.DeleteTaskInput
	5, // 5: task_service.TaskService.GetTaskByID:input_type -> task_service.GetTaskByIDInput
	7, // 6: task_service.TaskService.GetTasks:input_type -> task_service.GetTasksInput
	2, // 7: task_service.TaskService.CreateTask:output_type -> task_service.CreateTaskOutput
	0, // 8: task_service.TaskService.UpdateTask:output_type -> task_service.Task
	4, // 9: task_service.TaskService.DeleteTask:output_type -> task_service.DeleteTaskOutput
	5, // 10: task_service.TaskService.GetTaskByID:output_type -> task_service.GetTaskByIDInput
	8, // 11: task_service.TaskService.GetTasks:output_type -> task_service.GetTasksOutput
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_task_proto_init() }
func file_proto_task_proto_init() {
	if File_proto_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_proto_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTaskInput); i {
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
		file_proto_task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTaskOutput); i {
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
		file_proto_task_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTaskInput); i {
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
		file_proto_task_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTaskOutput); i {
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
		file_proto_task_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTaskByIDInput); i {
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
		file_proto_task_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTaskByIDOutput); i {
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
		file_proto_task_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTasksInput); i {
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
		file_proto_task_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTasksOutput); i {
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
	file_proto_task_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_task_proto_goTypes,
		DependencyIndexes: file_proto_task_proto_depIdxs,
		MessageInfos:      file_proto_task_proto_msgTypes,
	}.Build()
	File_proto_task_proto = out.File
	file_proto_task_proto_rawDesc = nil
	file_proto_task_proto_goTypes = nil
	file_proto_task_proto_depIdxs = nil
}
