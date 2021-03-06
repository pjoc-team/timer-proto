// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: scheduler/scheduler.proto

package scheduler

import (
	context "context"
	any "github.com/golang/protobuf/ptypes/any"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	common "github.com/pjoc-team/timer-proto/common"
	job "github.com/pjoc-team/timer-proto/job"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// HeartbeatRequest 心跳请求
type HeartbeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// task_id 任务id
	TaskId string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	// heartbeat_time 心跳时间
	HeartbeatTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=heartbeat_time,json=heartbeatTime,proto3" json:"heartbeat_time,omitempty"`
}

func (x *HeartbeatRequest) Reset() {
	*x = HeartbeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_scheduler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatRequest) ProtoMessage() {}

func (x *HeartbeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_scheduler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatRequest.ProtoReflect.Descriptor instead.
func (*HeartbeatRequest) Descriptor() ([]byte, []int) {
	return file_scheduler_scheduler_proto_rawDescGZIP(), []int{0}
}

func (x *HeartbeatRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *HeartbeatRequest) GetHeartbeatTime() *timestamp.Timestamp {
	if x != nil {
		return x.HeartbeatTime
	}
	return nil
}

// HeartbeatResponse 心跳响应
type HeartbeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HeartbeatResponse) Reset() {
	*x = HeartbeatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_scheduler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatResponse) ProtoMessage() {}

func (x *HeartbeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_scheduler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatResponse.ProtoReflect.Descriptor instead.
func (*HeartbeatResponse) Descriptor() ([]byte, []int) {
	return file_scheduler_scheduler_proto_rawDescGZIP(), []int{1}
}

type ScheduleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// schedule_id 调度id，全局唯一。worker需要定时上报该id的状态
	ScheduleId string `protobuf:"bytes,1,opt,name=schedule_id,json=scheduleId,proto3" json:"schedule_id,omitempty"`
	// task 任务内容
	Task *job.Task `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
	// heartbeat_interval 心跳间隔
	HeartbeatInterval *duration.Duration `protobuf:"bytes,3,opt,name=heartbeat_interval,json=heartbeatInterval,proto3" json:"heartbeat_interval,omitempty"`
	// heartbeat_timeout 心跳超时时间
	HeartbeatTimeout *duration.Duration `protobuf:"bytes,4,opt,name=heartbeat_timeout,json=heartbeatTimeout,proto3" json:"heartbeat_timeout,omitempty"`
	// report 心跳或者上报的目标地址
	Report *common.Target `protobuf:"bytes,5,opt,name=report,proto3" json:"report,omitempty"`
	// last_task_result 上一个task的执行结果，如果是第一个task，则该值会为空
	LastTaskResult *any.Any `protobuf:"bytes,6,opt,name=last_task_result,json=lastTaskResult,proto3" json:"last_task_result,omitempty"`
}

func (x *ScheduleRequest) Reset() {
	*x = ScheduleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_scheduler_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleRequest) ProtoMessage() {}

func (x *ScheduleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_scheduler_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleRequest.ProtoReflect.Descriptor instead.
func (*ScheduleRequest) Descriptor() ([]byte, []int) {
	return file_scheduler_scheduler_proto_rawDescGZIP(), []int{2}
}

func (x *ScheduleRequest) GetScheduleId() string {
	if x != nil {
		return x.ScheduleId
	}
	return ""
}

func (x *ScheduleRequest) GetTask() *job.Task {
	if x != nil {
		return x.Task
	}
	return nil
}

func (x *ScheduleRequest) GetHeartbeatInterval() *duration.Duration {
	if x != nil {
		return x.HeartbeatInterval
	}
	return nil
}

func (x *ScheduleRequest) GetHeartbeatTimeout() *duration.Duration {
	if x != nil {
		return x.HeartbeatTimeout
	}
	return nil
}

func (x *ScheduleRequest) GetReport() *common.Target {
	if x != nil {
		return x.Report
	}
	return nil
}

func (x *ScheduleRequest) GetLastTaskResult() *any.Any {
	if x != nil {
		return x.LastTaskResult
	}
	return nil
}

type ScheduleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ScheduleResponse) Reset() {
	*x = ScheduleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_scheduler_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleResponse) ProtoMessage() {}

func (x *ScheduleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_scheduler_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleResponse.ProtoReflect.Descriptor instead.
func (*ScheduleResponse) Descriptor() ([]byte, []int) {
	return file_scheduler_scheduler_proto_rawDescGZIP(), []int{3}
}

type CompleteTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// task_id 任务id
	TaskId string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	// now 完成时间
	Now *timestamp.Timestamp `protobuf:"bytes,2,opt,name=now,proto3" json:"now,omitempty"`
	// result 任务结果
	Result *any.Any `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
	// task_name 任务name
	TaskName string `protobuf:"bytes,4,opt,name=task_name,json=taskName,proto3" json:"task_name,omitempty"`
}

func (x *CompleteTaskRequest) Reset() {
	*x = CompleteTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_scheduler_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteTaskRequest) ProtoMessage() {}

func (x *CompleteTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_scheduler_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteTaskRequest.ProtoReflect.Descriptor instead.
func (*CompleteTaskRequest) Descriptor() ([]byte, []int) {
	return file_scheduler_scheduler_proto_rawDescGZIP(), []int{4}
}

func (x *CompleteTaskRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *CompleteTaskRequest) GetNow() *timestamp.Timestamp {
	if x != nil {
		return x.Now
	}
	return nil
}

func (x *CompleteTaskRequest) GetResult() *any.Any {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *CompleteTaskRequest) GetTaskName() string {
	if x != nil {
		return x.TaskName
	}
	return ""
}

type CompleteTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CompleteTaskResponse) Reset() {
	*x = CompleteTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_scheduler_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteTaskResponse) ProtoMessage() {}

func (x *CompleteTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_scheduler_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteTaskResponse.ProtoReflect.Descriptor instead.
func (*CompleteTaskResponse) Descriptor() ([]byte, []int) {
	return file_scheduler_scheduler_proto_rawDescGZIP(), []int{5}
}

type SetFailTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// task_id 任务id
	TaskId string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	// now 完成时间
	Now *timestamp.Timestamp `protobuf:"bytes,2,opt,name=now,proto3" json:"now,omitempty"`
	// reason 失败原因
	Reason string `protobuf:"bytes,3,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *SetFailTaskRequest) Reset() {
	*x = SetFailTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_scheduler_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetFailTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetFailTaskRequest) ProtoMessage() {}

func (x *SetFailTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_scheduler_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetFailTaskRequest.ProtoReflect.Descriptor instead.
func (*SetFailTaskRequest) Descriptor() ([]byte, []int) {
	return file_scheduler_scheduler_proto_rawDescGZIP(), []int{6}
}

func (x *SetFailTaskRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *SetFailTaskRequest) GetNow() *timestamp.Timestamp {
	if x != nil {
		return x.Now
	}
	return nil
}

func (x *SetFailTaskRequest) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type SetFailTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetFailTaskResponse) Reset() {
	*x = SetFailTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_scheduler_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetFailTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetFailTaskResponse) ProtoMessage() {}

func (x *SetFailTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_scheduler_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetFailTaskResponse.ProtoReflect.Descriptor instead.
func (*SetFailTaskResponse) Descriptor() ([]byte, []int) {
	return file_scheduler_scheduler_proto_rawDescGZIP(), []int{7}
}

var File_scheduler_scheduler_proto protoreflect.FileDescriptor

var file_scheduler_scheduler_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2f, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0d, 0x6a, 0x6f, 0x62, 0x2f, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x10, 0x48,
	0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x41, 0x0a, 0x0e, 0x68, 0x65, 0x61, 0x72,
	0x74, 0x62, 0x65, 0x61, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x68, 0x65,
	0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x48,
	0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0xcb, 0x02, 0x0a, 0x0f, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x04,
	0x74, 0x61, 0x73, 0x6b, 0x12, 0x48, 0x0a, 0x12, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61,
	0x74, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x68, 0x65, 0x61,
	0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x46,
	0x0a, 0x11, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x26, 0x0a, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x3e,
	0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0e,
	0x6c, 0x61, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x12,
	0x0a, 0x10, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0xa7, 0x01, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61,
	0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73,
	0x6b, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x03, 0x6e, 0x6f, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x6e, 0x6f,
	0x77, 0x12, 0x2c, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x16, 0x0a, 0x14,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x73, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61,
	0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73,
	0x6b, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x03, 0x6e, 0x6f, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x6e, 0x6f,
	0x77, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x15, 0x0a, 0x13, 0x53, 0x65, 0x74,
	0x46, 0x61, 0x69, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0xbf, 0x02, 0x0a, 0x09, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x12, 0x45,
	0x0a, 0x08, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x1a, 0x2e, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x72, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65,
	0x61, 0x74, 0x12, 0x1b, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x48,
	0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x48, 0x65, 0x61, 0x72,
	0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x4e, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x1d,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x46, 0x61,
	0x69, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x46, 0x61, 0x69,
	0x6c, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x51, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12,
	0x1e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x6a, 0x6f, 0x63, 0x2d, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x72,
	0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_scheduler_scheduler_proto_rawDescOnce sync.Once
	file_scheduler_scheduler_proto_rawDescData = file_scheduler_scheduler_proto_rawDesc
)

func file_scheduler_scheduler_proto_rawDescGZIP() []byte {
	file_scheduler_scheduler_proto_rawDescOnce.Do(func() {
		file_scheduler_scheduler_proto_rawDescData = protoimpl.X.CompressGZIP(file_scheduler_scheduler_proto_rawDescData)
	})
	return file_scheduler_scheduler_proto_rawDescData
}

var file_scheduler_scheduler_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_scheduler_scheduler_proto_goTypes = []interface{}{
	(*HeartbeatRequest)(nil),     // 0: scheduler.HeartbeatRequest
	(*HeartbeatResponse)(nil),    // 1: scheduler.HeartbeatResponse
	(*ScheduleRequest)(nil),      // 2: scheduler.ScheduleRequest
	(*ScheduleResponse)(nil),     // 3: scheduler.ScheduleResponse
	(*CompleteTaskRequest)(nil),  // 4: scheduler.CompleteTaskRequest
	(*CompleteTaskResponse)(nil), // 5: scheduler.CompleteTaskResponse
	(*SetFailTaskRequest)(nil),   // 6: scheduler.SetFailTaskRequest
	(*SetFailTaskResponse)(nil),  // 7: scheduler.SetFailTaskResponse
	(*timestamp.Timestamp)(nil),  // 8: google.protobuf.Timestamp
	(*job.Task)(nil),             // 9: job.Task
	(*duration.Duration)(nil),    // 10: google.protobuf.Duration
	(*common.Target)(nil),        // 11: common.Target
	(*any.Any)(nil),              // 12: google.protobuf.Any
}
var file_scheduler_scheduler_proto_depIdxs = []int32{
	8,  // 0: scheduler.HeartbeatRequest.heartbeat_time:type_name -> google.protobuf.Timestamp
	9,  // 1: scheduler.ScheduleRequest.task:type_name -> job.Task
	10, // 2: scheduler.ScheduleRequest.heartbeat_interval:type_name -> google.protobuf.Duration
	10, // 3: scheduler.ScheduleRequest.heartbeat_timeout:type_name -> google.protobuf.Duration
	11, // 4: scheduler.ScheduleRequest.report:type_name -> common.Target
	12, // 5: scheduler.ScheduleRequest.last_task_result:type_name -> google.protobuf.Any
	8,  // 6: scheduler.CompleteTaskRequest.now:type_name -> google.protobuf.Timestamp
	12, // 7: scheduler.CompleteTaskRequest.result:type_name -> google.protobuf.Any
	8,  // 8: scheduler.SetFailTaskRequest.now:type_name -> google.protobuf.Timestamp
	2,  // 9: scheduler.Scheduler.Schedule:input_type -> scheduler.ScheduleRequest
	0,  // 10: scheduler.Scheduler.Heartbeat:input_type -> scheduler.HeartbeatRequest
	6,  // 11: scheduler.Scheduler.SetFailTask:input_type -> scheduler.SetFailTaskRequest
	4,  // 12: scheduler.Scheduler.CompleteTask:input_type -> scheduler.CompleteTaskRequest
	3,  // 13: scheduler.Scheduler.Schedule:output_type -> scheduler.ScheduleResponse
	1,  // 14: scheduler.Scheduler.Heartbeat:output_type -> scheduler.HeartbeatResponse
	7,  // 15: scheduler.Scheduler.SetFailTask:output_type -> scheduler.SetFailTaskResponse
	5,  // 16: scheduler.Scheduler.CompleteTask:output_type -> scheduler.CompleteTaskResponse
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_scheduler_scheduler_proto_init() }
func file_scheduler_scheduler_proto_init() {
	if File_scheduler_scheduler_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_scheduler_scheduler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatRequest); i {
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
		file_scheduler_scheduler_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatResponse); i {
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
		file_scheduler_scheduler_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleRequest); i {
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
		file_scheduler_scheduler_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleResponse); i {
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
		file_scheduler_scheduler_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteTaskRequest); i {
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
		file_scheduler_scheduler_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteTaskResponse); i {
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
		file_scheduler_scheduler_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetFailTaskRequest); i {
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
		file_scheduler_scheduler_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetFailTaskResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_scheduler_scheduler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_scheduler_scheduler_proto_goTypes,
		DependencyIndexes: file_scheduler_scheduler_proto_depIdxs,
		MessageInfos:      file_scheduler_scheduler_proto_msgTypes,
	}.Build()
	File_scheduler_scheduler_proto = out.File
	file_scheduler_scheduler_proto_rawDesc = nil
	file_scheduler_scheduler_proto_goTypes = nil
	file_scheduler_scheduler_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SchedulerClient is the client API for Scheduler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SchedulerClient interface {
	// Schedule 执行任务
	Schedule(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*ScheduleResponse, error)
	// Heartbeat 心跳
	Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatResponse, error)
	// SetFailTask 任务失败
	SetFailTask(ctx context.Context, in *SetFailTaskRequest, opts ...grpc.CallOption) (*SetFailTaskResponse, error)
	// CompleteTask 完成任务
	CompleteTask(ctx context.Context, in *CompleteTaskRequest, opts ...grpc.CallOption) (*CompleteTaskResponse, error)
}

type schedulerClient struct {
	cc grpc.ClientConnInterface
}

func NewSchedulerClient(cc grpc.ClientConnInterface) SchedulerClient {
	return &schedulerClient{cc}
}

func (c *schedulerClient) Schedule(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*ScheduleResponse, error) {
	out := new(ScheduleResponse)
	err := c.cc.Invoke(ctx, "/scheduler.Scheduler/Schedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatResponse, error) {
	out := new(HeartbeatResponse)
	err := c.cc.Invoke(ctx, "/scheduler.Scheduler/Heartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) SetFailTask(ctx context.Context, in *SetFailTaskRequest, opts ...grpc.CallOption) (*SetFailTaskResponse, error) {
	out := new(SetFailTaskResponse)
	err := c.cc.Invoke(ctx, "/scheduler.Scheduler/SetFailTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) CompleteTask(ctx context.Context, in *CompleteTaskRequest, opts ...grpc.CallOption) (*CompleteTaskResponse, error) {
	out := new(CompleteTaskResponse)
	err := c.cc.Invoke(ctx, "/scheduler.Scheduler/CompleteTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchedulerServer is the server API for Scheduler service.
type SchedulerServer interface {
	// Schedule 执行任务
	Schedule(context.Context, *ScheduleRequest) (*ScheduleResponse, error)
	// Heartbeat 心跳
	Heartbeat(context.Context, *HeartbeatRequest) (*HeartbeatResponse, error)
	// SetFailTask 任务失败
	SetFailTask(context.Context, *SetFailTaskRequest) (*SetFailTaskResponse, error)
	// CompleteTask 完成任务
	CompleteTask(context.Context, *CompleteTaskRequest) (*CompleteTaskResponse, error)
}

// UnimplementedSchedulerServer can be embedded to have forward compatible implementations.
type UnimplementedSchedulerServer struct {
}

func (*UnimplementedSchedulerServer) Schedule(context.Context, *ScheduleRequest) (*ScheduleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Schedule not implemented")
}
func (*UnimplementedSchedulerServer) Heartbeat(context.Context, *HeartbeatRequest) (*HeartbeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heartbeat not implemented")
}
func (*UnimplementedSchedulerServer) SetFailTask(context.Context, *SetFailTaskRequest) (*SetFailTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetFailTask not implemented")
}
func (*UnimplementedSchedulerServer) CompleteTask(context.Context, *CompleteTaskRequest) (*CompleteTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteTask not implemented")
}

func RegisterSchedulerServer(s *grpc.Server, srv SchedulerServer) {
	s.RegisterService(&_Scheduler_serviceDesc, srv)
}

func _Scheduler_Schedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).Schedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.Scheduler/Schedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).Schedule(ctx, req.(*ScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartbeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.Scheduler/Heartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).Heartbeat(ctx, req.(*HeartbeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_SetFailTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetFailTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).SetFailTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.Scheduler/SetFailTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).SetFailTask(ctx, req.(*SetFailTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_CompleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).CompleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.Scheduler/CompleteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).CompleteTask(ctx, req.(*CompleteTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Scheduler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scheduler.Scheduler",
	HandlerType: (*SchedulerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Schedule",
			Handler:    _Scheduler_Schedule_Handler,
		},
		{
			MethodName: "Heartbeat",
			Handler:    _Scheduler_Heartbeat_Handler,
		},
		{
			MethodName: "SetFailTask",
			Handler:    _Scheduler_SetFailTask_Handler,
		},
		{
			MethodName: "CompleteTask",
			Handler:    _Scheduler_CompleteTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scheduler/scheduler.proto",
}
