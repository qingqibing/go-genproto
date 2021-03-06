// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/osconfig/agentendpoint/v1beta/tasks.proto

package agentendpoint

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Specifies the current agent behavior.
type TaskDirective int32

const (
	// Unspecified is invalid.
	TaskDirective_TASK_DIRECTIVE_UNSPECIFIED TaskDirective = 0
	// The task should continue to progress.
	TaskDirective_CONTINUE TaskDirective = 1
	// Task should not be started, or if already in progress, should stop
	// at first safe stopping point.  Task should be considered done and will
	// never repeat.
	TaskDirective_STOP TaskDirective = 2
)

var TaskDirective_name = map[int32]string{
	0: "TASK_DIRECTIVE_UNSPECIFIED",
	1: "CONTINUE",
	2: "STOP",
}

var TaskDirective_value = map[string]int32{
	"TASK_DIRECTIVE_UNSPECIFIED": 0,
	"CONTINUE":                   1,
	"STOP":                       2,
}

func (x TaskDirective) String() string {
	return proto.EnumName(TaskDirective_name, int32(x))
}

func (TaskDirective) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_282371b6fd36c625, []int{0}
}

// Specifies the type of task to perform.
type TaskType int32

const (
	// Unspecified is invalid.
	TaskType_TASK_TYPE_UNSPECIFIED TaskType = 0
	// The apply patches task.
	TaskType_APPLY_PATCHES TaskType = 1
	// The exec step task.
	TaskType_EXEC_STEP_TASK TaskType = 2
)

var TaskType_name = map[int32]string{
	0: "TASK_TYPE_UNSPECIFIED",
	1: "APPLY_PATCHES",
	2: "EXEC_STEP_TASK",
}

var TaskType_value = map[string]int32{
	"TASK_TYPE_UNSPECIFIED": 0,
	"APPLY_PATCHES":         1,
	"EXEC_STEP_TASK":        2,
}

func (x TaskType) String() string {
	return proto.EnumName(TaskType_name, int32(x))
}

func (TaskType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_282371b6fd36c625, []int{1}
}

// The intermediate states of applying patches.
type ApplyPatchesTaskProgress_State int32

const (
	// Unspecified is invalid.
	ApplyPatchesTaskProgress_STATE_UNSPECIFIED ApplyPatchesTaskProgress_State = 0
	// The agent has started the patch task.
	ApplyPatchesTaskProgress_STARTED ApplyPatchesTaskProgress_State = 4
	// The agent is currently downloading patches.
	ApplyPatchesTaskProgress_DOWNLOADING_PATCHES ApplyPatchesTaskProgress_State = 1
	// The agent is currently applying patches.
	ApplyPatchesTaskProgress_APPLYING_PATCHES ApplyPatchesTaskProgress_State = 2
	// The agent is currently rebooting the VM instance.
	ApplyPatchesTaskProgress_REBOOTING ApplyPatchesTaskProgress_State = 3
)

var ApplyPatchesTaskProgress_State_name = map[int32]string{
	0: "STATE_UNSPECIFIED",
	4: "STARTED",
	1: "DOWNLOADING_PATCHES",
	2: "APPLYING_PATCHES",
	3: "REBOOTING",
}

var ApplyPatchesTaskProgress_State_value = map[string]int32{
	"STATE_UNSPECIFIED":   0,
	"STARTED":             4,
	"DOWNLOADING_PATCHES": 1,
	"APPLYING_PATCHES":    2,
	"REBOOTING":           3,
}

func (x ApplyPatchesTaskProgress_State) String() string {
	return proto.EnumName(ApplyPatchesTaskProgress_State_name, int32(x))
}

func (ApplyPatchesTaskProgress_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_282371b6fd36c625, []int{2, 0}
}

// The final states of applying patches.
type ApplyPatchesTaskOutput_State int32

const (
	// Unspecified is invalid.
	ApplyPatchesTaskOutput_STATE_UNSPECIFIED ApplyPatchesTaskOutput_State = 0
	// Applying patches completed successfully.
	ApplyPatchesTaskOutput_SUCCEEDED ApplyPatchesTaskOutput_State = 1
	// Applying patches completed successfully, but a reboot is required.
	ApplyPatchesTaskOutput_SUCCEEDED_REBOOT_REQUIRED ApplyPatchesTaskOutput_State = 2
	// Applying patches failed.
	ApplyPatchesTaskOutput_FAILED ApplyPatchesTaskOutput_State = 3
)

var ApplyPatchesTaskOutput_State_name = map[int32]string{
	0: "STATE_UNSPECIFIED",
	1: "SUCCEEDED",
	2: "SUCCEEDED_REBOOT_REQUIRED",
	3: "FAILED",
}

var ApplyPatchesTaskOutput_State_value = map[string]int32{
	"STATE_UNSPECIFIED":         0,
	"SUCCEEDED":                 1,
	"SUCCEEDED_REBOOT_REQUIRED": 2,
	"FAILED":                    3,
}

func (x ApplyPatchesTaskOutput_State) String() string {
	return proto.EnumName(ApplyPatchesTaskOutput_State_name, int32(x))
}

func (ApplyPatchesTaskOutput_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_282371b6fd36c625, []int{3, 0}
}

// The intermediate states of exec steps.
type ExecStepTaskProgress_State int32

const (
	// Unspecified is invalid.
	ExecStepTaskProgress_STATE_UNSPECIFIED ExecStepTaskProgress_State = 0
	// The agent has started the exec step task.
	ExecStepTaskProgress_STARTED ExecStepTaskProgress_State = 1
)

var ExecStepTaskProgress_State_name = map[int32]string{
	0: "STATE_UNSPECIFIED",
	1: "STARTED",
}

var ExecStepTaskProgress_State_value = map[string]int32{
	"STATE_UNSPECIFIED": 0,
	"STARTED":           1,
}

func (x ExecStepTaskProgress_State) String() string {
	return proto.EnumName(ExecStepTaskProgress_State_name, int32(x))
}

func (ExecStepTaskProgress_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_282371b6fd36c625, []int{5, 0}
}

// The final states of exec steps.
type ExecStepTaskOutput_State int32

const (
	// Unspecified is invalid.
	ExecStepTaskOutput_STATE_UNSPECIFIED ExecStepTaskOutput_State = 0
	// The exec step completed normally.
	ExecStepTaskOutput_COMPLETED ExecStepTaskOutput_State = 1
	// The exec step was terminated because it took too long.
	ExecStepTaskOutput_TIMED_OUT ExecStepTaskOutput_State = 2
	// The exec step task was cancelled before it started.
	ExecStepTaskOutput_CANCELLED ExecStepTaskOutput_State = 3
)

var ExecStepTaskOutput_State_name = map[int32]string{
	0: "STATE_UNSPECIFIED",
	1: "COMPLETED",
	2: "TIMED_OUT",
	3: "CANCELLED",
}

var ExecStepTaskOutput_State_value = map[string]int32{
	"STATE_UNSPECIFIED": 0,
	"COMPLETED":         1,
	"TIMED_OUT":         2,
	"CANCELLED":         3,
}

func (x ExecStepTaskOutput_State) String() string {
	return proto.EnumName(ExecStepTaskOutput_State_name, int32(x))
}

func (ExecStepTaskOutput_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_282371b6fd36c625, []int{6, 0}
}

// A unit of work to be performed by the agent.
type Task struct {
	// Unique task id.
	TaskId string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	// The type of task to perform.
	//
	// Task details must include the appropriate message based on this enum as
	// specified below:
	// APPLY_PATCHES = ApplyPatchesTask
	// EXEC_STEP = ExecStepTask;
	TaskType TaskType `protobuf:"varint,2,opt,name=task_type,json=taskType,proto3,enum=google.cloud.osconfig.agentendpoint.v1beta.TaskType" json:"task_type,omitempty"`
	// Current directive to the agent.
	TaskDirective TaskDirective `protobuf:"varint,3,opt,name=task_directive,json=taskDirective,proto3,enum=google.cloud.osconfig.agentendpoint.v1beta.TaskDirective" json:"task_directive,omitempty"`
	// Specific details about the current task to perform.
	//
	// Types that are valid to be assigned to TaskDetails:
	//	*Task_ApplyPatchesTask
	//	*Task_ExecStepTask
	TaskDetails isTask_TaskDetails `protobuf_oneof:"task_details"`
	// Labels describing the task.  Used for logging by the agent.
	ServiceLabels        map[string]string `protobuf:"bytes,6,rep,name=service_labels,json=serviceLabels,proto3" json:"service_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_282371b6fd36c625, []int{0}
}

func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *Task) GetTaskType() TaskType {
	if m != nil {
		return m.TaskType
	}
	return TaskType_TASK_TYPE_UNSPECIFIED
}

func (m *Task) GetTaskDirective() TaskDirective {
	if m != nil {
		return m.TaskDirective
	}
	return TaskDirective_TASK_DIRECTIVE_UNSPECIFIED
}

type isTask_TaskDetails interface {
	isTask_TaskDetails()
}

type Task_ApplyPatchesTask struct {
	ApplyPatchesTask *ApplyPatchesTask `protobuf:"bytes,4,opt,name=apply_patches_task,json=applyPatchesTask,proto3,oneof"`
}

type Task_ExecStepTask struct {
	ExecStepTask *ExecStepTask `protobuf:"bytes,5,opt,name=exec_step_task,json=execStepTask,proto3,oneof"`
}

func (*Task_ApplyPatchesTask) isTask_TaskDetails() {}

func (*Task_ExecStepTask) isTask_TaskDetails() {}

func (m *Task) GetTaskDetails() isTask_TaskDetails {
	if m != nil {
		return m.TaskDetails
	}
	return nil
}

func (m *Task) GetApplyPatchesTask() *ApplyPatchesTask {
	if x, ok := m.GetTaskDetails().(*Task_ApplyPatchesTask); ok {
		return x.ApplyPatchesTask
	}
	return nil
}

func (m *Task) GetExecStepTask() *ExecStepTask {
	if x, ok := m.GetTaskDetails().(*Task_ExecStepTask); ok {
		return x.ExecStepTask
	}
	return nil
}

func (m *Task) GetServiceLabels() map[string]string {
	if m != nil {
		return m.ServiceLabels
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Task) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Task_ApplyPatchesTask)(nil),
		(*Task_ExecStepTask)(nil),
	}
}

// Message which instructs agent to apply patches.
type ApplyPatchesTask struct {
	// Specific information about how patches should be applied.
	PatchConfig *PatchConfig `protobuf:"bytes,1,opt,name=patch_config,json=patchConfig,proto3" json:"patch_config,omitempty"`
	// If true, the agent will report its status as it goes through the motions
	// but won't actually run any updates or perform any reboots.
	DryRun               bool     `protobuf:"varint,3,opt,name=dry_run,json=dryRun,proto3" json:"dry_run,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplyPatchesTask) Reset()         { *m = ApplyPatchesTask{} }
func (m *ApplyPatchesTask) String() string { return proto.CompactTextString(m) }
func (*ApplyPatchesTask) ProtoMessage()    {}
func (*ApplyPatchesTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_282371b6fd36c625, []int{1}
}

func (m *ApplyPatchesTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyPatchesTask.Unmarshal(m, b)
}
func (m *ApplyPatchesTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyPatchesTask.Marshal(b, m, deterministic)
}
func (m *ApplyPatchesTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyPatchesTask.Merge(m, src)
}
func (m *ApplyPatchesTask) XXX_Size() int {
	return xxx_messageInfo_ApplyPatchesTask.Size(m)
}
func (m *ApplyPatchesTask) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyPatchesTask.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyPatchesTask proto.InternalMessageInfo

func (m *ApplyPatchesTask) GetPatchConfig() *PatchConfig {
	if m != nil {
		return m.PatchConfig
	}
	return nil
}

func (m *ApplyPatchesTask) GetDryRun() bool {
	if m != nil {
		return m.DryRun
	}
	return false
}

// Information reported from the agent about applying patches execution.
type ApplyPatchesTaskProgress struct {
	// Required. The current state of this patch execution.
	State                ApplyPatchesTaskProgress_State `protobuf:"varint,1,opt,name=state,proto3,enum=google.cloud.osconfig.agentendpoint.v1beta.ApplyPatchesTaskProgress_State" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *ApplyPatchesTaskProgress) Reset()         { *m = ApplyPatchesTaskProgress{} }
func (m *ApplyPatchesTaskProgress) String() string { return proto.CompactTextString(m) }
func (*ApplyPatchesTaskProgress) ProtoMessage()    {}
func (*ApplyPatchesTaskProgress) Descriptor() ([]byte, []int) {
	return fileDescriptor_282371b6fd36c625, []int{2}
}

func (m *ApplyPatchesTaskProgress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyPatchesTaskProgress.Unmarshal(m, b)
}
func (m *ApplyPatchesTaskProgress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyPatchesTaskProgress.Marshal(b, m, deterministic)
}
func (m *ApplyPatchesTaskProgress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyPatchesTaskProgress.Merge(m, src)
}
func (m *ApplyPatchesTaskProgress) XXX_Size() int {
	return xxx_messageInfo_ApplyPatchesTaskProgress.Size(m)
}
func (m *ApplyPatchesTaskProgress) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyPatchesTaskProgress.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyPatchesTaskProgress proto.InternalMessageInfo

func (m *ApplyPatchesTaskProgress) GetState() ApplyPatchesTaskProgress_State {
	if m != nil {
		return m.State
	}
	return ApplyPatchesTaskProgress_STATE_UNSPECIFIED
}

// Information reported from the agent about applying patches execution.
type ApplyPatchesTaskOutput struct {
	// Required. The final state of this task.
	State                ApplyPatchesTaskOutput_State `protobuf:"varint,1,opt,name=state,proto3,enum=google.cloud.osconfig.agentendpoint.v1beta.ApplyPatchesTaskOutput_State" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ApplyPatchesTaskOutput) Reset()         { *m = ApplyPatchesTaskOutput{} }
func (m *ApplyPatchesTaskOutput) String() string { return proto.CompactTextString(m) }
func (*ApplyPatchesTaskOutput) ProtoMessage()    {}
func (*ApplyPatchesTaskOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_282371b6fd36c625, []int{3}
}

func (m *ApplyPatchesTaskOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyPatchesTaskOutput.Unmarshal(m, b)
}
func (m *ApplyPatchesTaskOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyPatchesTaskOutput.Marshal(b, m, deterministic)
}
func (m *ApplyPatchesTaskOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyPatchesTaskOutput.Merge(m, src)
}
func (m *ApplyPatchesTaskOutput) XXX_Size() int {
	return xxx_messageInfo_ApplyPatchesTaskOutput.Size(m)
}
func (m *ApplyPatchesTaskOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyPatchesTaskOutput.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyPatchesTaskOutput proto.InternalMessageInfo

func (m *ApplyPatchesTaskOutput) GetState() ApplyPatchesTaskOutput_State {
	if m != nil {
		return m.State
	}
	return ApplyPatchesTaskOutput_STATE_UNSPECIFIED
}

// Message which instructs agent to execute the following command.
type ExecStepTask struct {
	// Details of the exec step to run.
	ExecStep             *ExecStep `protobuf:"bytes,1,opt,name=exec_step,json=execStep,proto3" json:"exec_step,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ExecStepTask) Reset()         { *m = ExecStepTask{} }
func (m *ExecStepTask) String() string { return proto.CompactTextString(m) }
func (*ExecStepTask) ProtoMessage()    {}
func (*ExecStepTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_282371b6fd36c625, []int{4}
}

func (m *ExecStepTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecStepTask.Unmarshal(m, b)
}
func (m *ExecStepTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecStepTask.Marshal(b, m, deterministic)
}
func (m *ExecStepTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecStepTask.Merge(m, src)
}
func (m *ExecStepTask) XXX_Size() int {
	return xxx_messageInfo_ExecStepTask.Size(m)
}
func (m *ExecStepTask) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecStepTask.DiscardUnknown(m)
}

var xxx_messageInfo_ExecStepTask proto.InternalMessageInfo

func (m *ExecStepTask) GetExecStep() *ExecStep {
	if m != nil {
		return m.ExecStep
	}
	return nil
}

// Information reported from the agent about the exec step execution.
type ExecStepTaskProgress struct {
	// Required. The current state of this exec step.
	State                ExecStepTaskProgress_State `protobuf:"varint,1,opt,name=state,proto3,enum=google.cloud.osconfig.agentendpoint.v1beta.ExecStepTaskProgress_State" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ExecStepTaskProgress) Reset()         { *m = ExecStepTaskProgress{} }
func (m *ExecStepTaskProgress) String() string { return proto.CompactTextString(m) }
func (*ExecStepTaskProgress) ProtoMessage()    {}
func (*ExecStepTaskProgress) Descriptor() ([]byte, []int) {
	return fileDescriptor_282371b6fd36c625, []int{5}
}

func (m *ExecStepTaskProgress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecStepTaskProgress.Unmarshal(m, b)
}
func (m *ExecStepTaskProgress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecStepTaskProgress.Marshal(b, m, deterministic)
}
func (m *ExecStepTaskProgress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecStepTaskProgress.Merge(m, src)
}
func (m *ExecStepTaskProgress) XXX_Size() int {
	return xxx_messageInfo_ExecStepTaskProgress.Size(m)
}
func (m *ExecStepTaskProgress) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecStepTaskProgress.DiscardUnknown(m)
}

var xxx_messageInfo_ExecStepTaskProgress proto.InternalMessageInfo

func (m *ExecStepTaskProgress) GetState() ExecStepTaskProgress_State {
	if m != nil {
		return m.State
	}
	return ExecStepTaskProgress_STATE_UNSPECIFIED
}

// Information reported from the agent about the exec step execution.
type ExecStepTaskOutput struct {
	// Required. The final state of the exec step.
	State ExecStepTaskOutput_State `protobuf:"varint,1,opt,name=state,proto3,enum=google.cloud.osconfig.agentendpoint.v1beta.ExecStepTaskOutput_State" json:"state,omitempty"`
	// Required. The exit code received from the script which ran as part of the exec step.
	ExitCode             int32    `protobuf:"varint,2,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecStepTaskOutput) Reset()         { *m = ExecStepTaskOutput{} }
func (m *ExecStepTaskOutput) String() string { return proto.CompactTextString(m) }
func (*ExecStepTaskOutput) ProtoMessage()    {}
func (*ExecStepTaskOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_282371b6fd36c625, []int{6}
}

func (m *ExecStepTaskOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecStepTaskOutput.Unmarshal(m, b)
}
func (m *ExecStepTaskOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecStepTaskOutput.Marshal(b, m, deterministic)
}
func (m *ExecStepTaskOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecStepTaskOutput.Merge(m, src)
}
func (m *ExecStepTaskOutput) XXX_Size() int {
	return xxx_messageInfo_ExecStepTaskOutput.Size(m)
}
func (m *ExecStepTaskOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecStepTaskOutput.DiscardUnknown(m)
}

var xxx_messageInfo_ExecStepTaskOutput proto.InternalMessageInfo

func (m *ExecStepTaskOutput) GetState() ExecStepTaskOutput_State {
	if m != nil {
		return m.State
	}
	return ExecStepTaskOutput_STATE_UNSPECIFIED
}

func (m *ExecStepTaskOutput) GetExitCode() int32 {
	if m != nil {
		return m.ExitCode
	}
	return 0
}

func init() {
	proto.RegisterEnum("google.cloud.osconfig.agentendpoint.v1beta.TaskDirective", TaskDirective_name, TaskDirective_value)
	proto.RegisterEnum("google.cloud.osconfig.agentendpoint.v1beta.TaskType", TaskType_name, TaskType_value)
	proto.RegisterEnum("google.cloud.osconfig.agentendpoint.v1beta.ApplyPatchesTaskProgress_State", ApplyPatchesTaskProgress_State_name, ApplyPatchesTaskProgress_State_value)
	proto.RegisterEnum("google.cloud.osconfig.agentendpoint.v1beta.ApplyPatchesTaskOutput_State", ApplyPatchesTaskOutput_State_name, ApplyPatchesTaskOutput_State_value)
	proto.RegisterEnum("google.cloud.osconfig.agentendpoint.v1beta.ExecStepTaskProgress_State", ExecStepTaskProgress_State_name, ExecStepTaskProgress_State_value)
	proto.RegisterEnum("google.cloud.osconfig.agentendpoint.v1beta.ExecStepTaskOutput_State", ExecStepTaskOutput_State_name, ExecStepTaskOutput_State_value)
	proto.RegisterType((*Task)(nil), "google.cloud.osconfig.agentendpoint.v1beta.Task")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.osconfig.agentendpoint.v1beta.Task.ServiceLabelsEntry")
	proto.RegisterType((*ApplyPatchesTask)(nil), "google.cloud.osconfig.agentendpoint.v1beta.ApplyPatchesTask")
	proto.RegisterType((*ApplyPatchesTaskProgress)(nil), "google.cloud.osconfig.agentendpoint.v1beta.ApplyPatchesTaskProgress")
	proto.RegisterType((*ApplyPatchesTaskOutput)(nil), "google.cloud.osconfig.agentendpoint.v1beta.ApplyPatchesTaskOutput")
	proto.RegisterType((*ExecStepTask)(nil), "google.cloud.osconfig.agentendpoint.v1beta.ExecStepTask")
	proto.RegisterType((*ExecStepTaskProgress)(nil), "google.cloud.osconfig.agentendpoint.v1beta.ExecStepTaskProgress")
	proto.RegisterType((*ExecStepTaskOutput)(nil), "google.cloud.osconfig.agentendpoint.v1beta.ExecStepTaskOutput")
}

func init() {
	proto.RegisterFile("google/cloud/osconfig/agentendpoint/v1beta/tasks.proto", fileDescriptor_282371b6fd36c625)
}

var fileDescriptor_282371b6fd36c625 = []byte{
	// 864 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0xcd, 0x92, 0xdb, 0x44,
	0x10, 0xc7, 0x23, 0x7f, 0xc5, 0x6e, 0x7f, 0x94, 0x32, 0x6c, 0x88, 0xb3, 0x55, 0x80, 0xcb, 0xa7,
	0xad, 0xa5, 0x4a, 0x2e, 0x4c, 0x0a, 0x02, 0xe1, 0x80, 0x2c, 0xcd, 0xee, 0x8a, 0x38, 0xb6, 0x22,
	0xc9, 0xe4, 0xe3, 0x32, 0xc8, 0xd2, 0x44, 0x51, 0x56, 0x48, 0x2a, 0x69, 0xec, 0x8a, 0x9f, 0x80,
	0x2b, 0x4f, 0xc1, 0x2b, 0x71, 0xa5, 0x28, 0x8e, 0x3c, 0x04, 0xa5, 0x19, 0x7b, 0xb1, 0xcd, 0x47,
	0xd6, 0x9b, 0x9b, 0xba, 0x47, 0xfd, 0xeb, 0xee, 0x99, 0x7f, 0xcf, 0xc0, 0x17, 0x41, 0x92, 0x04,
	0x11, 0x1d, 0x78, 0x51, 0xb2, 0xf0, 0x07, 0x49, 0xee, 0x25, 0xf1, 0xab, 0x30, 0x18, 0xb8, 0x01,
	0x8d, 0x19, 0x8d, 0xfd, 0x34, 0x09, 0x63, 0x36, 0x58, 0x7e, 0x36, 0xa7, 0xcc, 0x1d, 0x30, 0x37,
	0xbf, 0xcc, 0x95, 0x34, 0x4b, 0x58, 0x82, 0x4e, 0x45, 0x9c, 0xc2, 0xe3, 0x94, 0x4d, 0x9c, 0xb2,
	0x13, 0xa7, 0x88, 0xb8, 0xe3, 0x4f, 0xd6, 0x39, 0xdc, 0x34, 0x1c, 0xbc, 0x0a, 0x69, 0xe4, 0x93,
	0x39, 0x7d, 0xed, 0x2e, 0xc3, 0x24, 0x13, 0xb0, 0xe3, 0x47, 0x07, 0x14, 0x91, 0xba, 0xcc, 0x7b,
	0x4d, 0xde, 0x24, 0xf3, 0x75, 0x25, 0xfd, 0xdf, 0x2b, 0x50, 0x71, 0xdc, 0xfc, 0x12, 0xdd, 0x83,
	0xdb, 0x45, 0x85, 0x24, 0xf4, 0xbb, 0x52, 0x4f, 0x3a, 0x69, 0x58, 0xb5, 0xc2, 0x34, 0x7c, 0xf4,
	0x14, 0x1a, 0x7c, 0x81, 0xad, 0x52, 0xda, 0x2d, 0xf5, 0xa4, 0x93, 0xce, 0xf0, 0x81, 0x72, 0xfd,
	0xfa, 0x95, 0x82, 0xee, 0xac, 0x52, 0x6a, 0xd5, 0xd9, 0xfa, 0x0b, 0xfd, 0x00, 0x1d, 0x8e, 0xf4,
	0xc3, 0x8c, 0x7a, 0x2c, 0x5c, 0xd2, 0x6e, 0x99, 0x73, 0xbf, 0x3a, 0x94, 0xab, 0x6f, 0x00, 0x56,
	0x9b, 0x6d, 0x9b, 0x28, 0x02, 0xe4, 0xa6, 0x69, 0xb4, 0x22, 0xbc, 0x61, 0x9a, 0x93, 0x62, 0xb9,
	0x5b, 0xe9, 0x49, 0x27, 0xcd, 0xe1, 0x37, 0x87, 0x64, 0x51, 0x0b, 0x8a, 0x29, 0x20, 0x45, 0xc6,
	0x8b, 0x5b, 0x96, 0xec, 0xee, 0xf9, 0x8a, 0x7e, 0xe8, 0x5b, 0xea, 0x91, 0x9c, 0xd1, 0x54, 0x64,
	0xaa, 0xf2, 0x4c, 0x0f, 0x0f, 0xc9, 0x84, 0xdf, 0x52, 0xcf, 0x66, 0x34, 0x5d, 0x67, 0x69, 0xd1,
	0x2d, 0x1b, 0xbd, 0x81, 0x4e, 0x4e, 0xb3, 0x65, 0xe8, 0x51, 0x12, 0xb9, 0x73, 0x1a, 0xe5, 0xdd,
	0x5a, 0xaf, 0x7c, 0xd2, 0x1c, 0x6a, 0x87, 0xee, 0x98, 0x62, 0x0b, 0xcc, 0x98, 0x53, 0x70, 0xcc,
	0xb2, 0x95, 0xd5, 0xce, 0xb7, 0x7d, 0xc7, 0xdf, 0x02, 0xfa, 0xe7, 0x4f, 0x48, 0x86, 0xf2, 0x25,
	0x5d, 0xad, 0xb5, 0x51, 0x7c, 0xa2, 0x23, 0xa8, 0x2e, 0xdd, 0x68, 0x21, 0x44, 0xd1, 0xb0, 0x84,
	0xf1, 0x75, 0xe9, 0xa1, 0x34, 0xea, 0x40, 0x4b, 0x9c, 0x2f, 0x65, 0x6e, 0x18, 0xe5, 0xfd, 0x9f,
	0x24, 0x90, 0xf7, 0x37, 0x12, 0xbd, 0x84, 0x96, 0x50, 0xa3, 0x28, 0x99, 0x93, 0x9b, 0xc3, 0x2f,
	0x0f, 0x69, 0x88, 0xe3, 0x34, 0xfe, 0x83, 0xd5, 0x4c, 0xff, 0x36, 0x0a, 0x31, 0xfb, 0xd9, 0x8a,
	0x64, 0x8b, 0x98, 0x2b, 0xab, 0x6e, 0xd5, 0xfc, 0x6c, 0x65, 0x2d, 0xe2, 0xfe, 0x9f, 0x12, 0x74,
	0xf7, 0x2b, 0x31, 0xb3, 0x24, 0xc8, 0x68, 0x9e, 0x23, 0x0a, 0xd5, 0x9c, 0xb9, 0x8c, 0xf2, 0x52,
	0x3a, 0xc3, 0xef, 0xde, 0x47, 0x27, 0x1b, 0xa8, 0x62, 0x17, 0xc4, 0x51, 0xf9, 0x37, 0xb5, 0x64,
	0x09, 0x7a, 0x3f, 0x84, 0x2a, 0x77, 0xa2, 0xbb, 0x70, 0xc7, 0x76, 0x54, 0x07, 0x93, 0xd9, 0xc4,
	0x36, 0xb1, 0x66, 0x9c, 0x19, 0x58, 0x97, 0x6f, 0xa1, 0x26, 0xdc, 0xb6, 0x1d, 0xd5, 0x72, 0xb0,
	0x2e, 0x57, 0xd0, 0x3d, 0xf8, 0x40, 0x9f, 0x3e, 0x9b, 0x8c, 0xa7, 0xaa, 0x6e, 0x4c, 0xce, 0x89,
	0xa9, 0x3a, 0xda, 0x05, 0xb6, 0x65, 0x09, 0x1d, 0x81, 0xac, 0x9a, 0xe6, 0xf8, 0xc5, 0xb6, 0xb7,
	0x84, 0xda, 0xd0, 0xb0, 0xf0, 0x68, 0x3a, 0x75, 0x8c, 0xc9, 0xb9, 0x5c, 0xee, 0xff, 0x2a, 0xc1,
	0x87, 0xfb, 0x95, 0x4d, 0x17, 0x2c, 0x5d, 0x30, 0xe4, 0xed, 0x36, 0x7b, 0xf1, 0x3e, 0xcd, 0x0a,
	0xe4, 0xbf, 0xb4, 0xfa, 0xfc, 0x1d, 0xad, 0xb6, 0xa1, 0x61, 0xcf, 0x34, 0x0d, 0x63, 0x1d, 0xeb,
	0xb2, 0x84, 0x3e, 0x82, 0xfb, 0x57, 0x26, 0x11, 0x7d, 0x10, 0x0b, 0x3f, 0x9d, 0x19, 0x16, 0xd6,
	0xe5, 0x12, 0x02, 0xa8, 0x9d, 0xa9, 0xc6, 0x18, 0xeb, 0x72, 0xb9, 0xef, 0x42, 0x6b, 0x7b, 0x60,
	0x8a, 0x5b, 0xea, 0x6a, 0x04, 0xd7, 0x52, 0x7a, 0x70, 0x93, 0xe9, 0xb3, 0xea, 0x9b, 0xb9, 0xeb,
	0xff, 0x22, 0xc1, 0xd1, 0x76, 0x8e, 0x2b, 0x9d, 0xb8, 0xbb, 0x5b, 0x77, 0x76, 0xd3, 0x29, 0xff,
	0x1f, 0x8d, 0x7c, 0x7a, 0x7d, 0x8d, 0x48, 0xfd, 0x3f, 0x24, 0x40, 0xdb, 0xdc, 0xf5, 0x09, 0x93,
	0xdd, 0x32, 0xf5, 0x9b, 0x96, 0xf9, 0x5f, 0xa7, 0x8b, 0x7a, 0xc5, 0x9e, 0x87, 0x8c, 0x78, 0x89,
	0x2f, 0x2e, 0x81, 0xaa, 0x58, 0xae, 0x17, 0x5e, 0x2d, 0xf1, 0x69, 0xff, 0xf1, 0xbb, 0xcf, 0x5f,
	0x9b, 0x3e, 0x31, 0xc7, 0x98, 0x37, 0x52, 0x98, 0x8e, 0xf1, 0x04, 0xeb, 0x64, 0x3a, 0x73, 0x84,
	0x98, 0x35, 0x75, 0xa2, 0xe1, 0x31, 0x3f, 0xf2, 0xd3, 0x73, 0x68, 0xef, 0xdc, 0xf9, 0xe8, 0x63,
	0x38, 0x76, 0x54, 0xfb, 0x31, 0xd1, 0x0d, 0x0b, 0x6b, 0x8e, 0xf1, 0xfd, 0x3e, 0xbd, 0x05, 0x75,
	0x6d, 0x3a, 0x71, 0x8c, 0xc9, 0x0c, 0xcb, 0x12, 0xaa, 0x43, 0xc5, 0x76, 0xa6, 0xa6, 0x5c, 0x3a,
	0x1d, 0x43, 0x7d, 0xf3, 0x28, 0xa1, 0xfb, 0x70, 0x97, 0x33, 0x9c, 0x17, 0xe6, 0x7e, 0xf8, 0x1d,
	0x68, 0xf3, 0x09, 0xdb, 0x1a, 0x3a, 0x04, 0x1d, 0xfc, 0x1c, 0x6b, 0xc4, 0x76, 0xb0, 0x49, 0x8a,
	0x38, 0xb9, 0x34, 0xfa, 0x59, 0x02, 0xc5, 0x4b, 0x7e, 0x3c, 0x60, 0x77, 0x47, 0xd5, 0x22, 0x7d,
	0xfe, 0xf2, 0xd9, 0x3a, 0x24, 0x48, 0x22, 0x37, 0x0e, 0x94, 0x24, 0x0b, 0x06, 0x01, 0x8d, 0xf9,
	0xbb, 0x3c, 0x10, 0x4b, 0x6e, 0x1a, 0xe6, 0xd7, 0x79, 0xd7, 0x1f, 0xed, 0x38, 0xe7, 0x35, 0xce,
	0xf8, 0xfc, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8b, 0xab, 0x36, 0xde, 0x9f, 0x08, 0x00, 0x00,
}
