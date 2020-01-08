// Code generated by protoc-gen-go. DO NOT EDIT.
// source: job.proto

package rpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type Job struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Dir                  string   `protobuf:"bytes,2,opt,name=dir,proto3" json:"dir,omitempty"`
	Program              string   `protobuf:"bytes,3,opt,name=program,proto3" json:"program,omitempty"`
	Args                 string   `protobuf:"bytes,4,opt,name=args,proto3" json:"args,omitempty"`
	StdOut               string   `protobuf:"bytes,5,opt,name=std_out,json=stdOut,proto3" json:"std_out,omitempty"`
	StdErr               string   `protobuf:"bytes,6,opt,name=std_err,json=stdErr,proto3" json:"std_err,omitempty"`
	Spec                 string   `protobuf:"bytes,7,opt,name=spec,proto3" json:"spec,omitempty"`
	Timeout              int64    `protobuf:"varint,8,opt,name=timeout,proto3" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Job) Reset()         { *m = Job{} }
func (m *Job) String() string { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()    {}
func (*Job) Descriptor() ([]byte, []int) {
	return fileDescriptor_f32c477d91a04ead, []int{0}
}

func (m *Job) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Job.Unmarshal(m, b)
}
func (m *Job) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Job.Marshal(b, m, deterministic)
}
func (m *Job) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Job.Merge(m, src)
}
func (m *Job) XXX_Size() int {
	return xxx_messageInfo_Job.Size(m)
}
func (m *Job) XXX_DiscardUnknown() {
	xxx_messageInfo_Job.DiscardUnknown(m)
}

var xxx_messageInfo_Job proto.InternalMessageInfo

func (m *Job) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Job) GetDir() string {
	if m != nil {
		return m.Dir
	}
	return ""
}

func (m *Job) GetProgram() string {
	if m != nil {
		return m.Program
	}
	return ""
}

func (m *Job) GetArgs() string {
	if m != nil {
		return m.Args
	}
	return ""
}

func (m *Job) GetStdOut() string {
	if m != nil {
		return m.StdOut
	}
	return ""
}

func (m *Job) GetStdErr() string {
	if m != nil {
		return m.StdErr
	}
	return ""
}

func (m *Job) GetSpec() string {
	if m != nil {
		return m.Spec
	}
	return ""
}

func (m *Job) GetTimeout() int64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

type JobNameRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobNameRequest) Reset()         { *m = JobNameRequest{} }
func (m *JobNameRequest) String() string { return proto.CompactTextString(m) }
func (*JobNameRequest) ProtoMessage()    {}
func (*JobNameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f32c477d91a04ead, []int{1}
}

func (m *JobNameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobNameRequest.Unmarshal(m, b)
}
func (m *JobNameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobNameRequest.Marshal(b, m, deterministic)
}
func (m *JobNameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobNameRequest.Merge(m, src)
}
func (m *JobNameRequest) XXX_Size() int {
	return xxx_messageInfo_JobNameRequest.Size(m)
}
func (m *JobNameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JobNameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JobNameRequest proto.InternalMessageInfo

func (m *JobNameRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type JobResponse struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Job                  *Job     `protobuf:"bytes,2,opt,name=job,proto3" json:"job,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobResponse) Reset()         { *m = JobResponse{} }
func (m *JobResponse) String() string { return proto.CompactTextString(m) }
func (*JobResponse) ProtoMessage()    {}
func (*JobResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f32c477d91a04ead, []int{2}
}

func (m *JobResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobResponse.Unmarshal(m, b)
}
func (m *JobResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobResponse.Marshal(b, m, deterministic)
}
func (m *JobResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobResponse.Merge(m, src)
}
func (m *JobResponse) XXX_Size() int {
	return xxx_messageInfo_JobResponse.Size(m)
}
func (m *JobResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_JobResponse.DiscardUnknown(m)
}

var xxx_messageInfo_JobResponse proto.InternalMessageInfo

func (m *JobResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *JobResponse) GetJob() *Job {
	if m != nil {
		return m.Job
	}
	return nil
}

type JobListResponse struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Jobs                 []*Job   `protobuf:"bytes,2,rep,name=jobs,proto3" json:"jobs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobListResponse) Reset()         { *m = JobListResponse{} }
func (m *JobListResponse) String() string { return proto.CompactTextString(m) }
func (*JobListResponse) ProtoMessage()    {}
func (*JobListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f32c477d91a04ead, []int{3}
}

func (m *JobListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobListResponse.Unmarshal(m, b)
}
func (m *JobListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobListResponse.Marshal(b, m, deterministic)
}
func (m *JobListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobListResponse.Merge(m, src)
}
func (m *JobListResponse) XXX_Size() int {
	return xxx_messageInfo_JobListResponse.Size(m)
}
func (m *JobListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_JobListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_JobListResponse proto.InternalMessageInfo

func (m *JobListResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *JobListResponse) GetJobs() []*Job {
	if m != nil {
		return m.Jobs
	}
	return nil
}

func init() {
	proto.RegisterType((*Job)(nil), "Job")
	proto.RegisterType((*JobNameRequest)(nil), "JobNameRequest")
	proto.RegisterType((*JobResponse)(nil), "JobResponse")
	proto.RegisterType((*JobListResponse)(nil), "JobListResponse")
}

func init() { proto.RegisterFile("job.proto", fileDescriptor_f32c477d91a04ead) }

var fileDescriptor_f32c477d91a04ead = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x4f, 0x4b, 0xc3, 0x30,
	0x14, 0x5f, 0x97, 0xae, 0x73, 0x6f, 0xe2, 0x46, 0x0e, 0x1a, 0xea, 0x65, 0x14, 0x19, 0x3b, 0xf5,
	0x30, 0x4f, 0x9e, 0x44, 0x65, 0x08, 0x45, 0x14, 0x0a, 0x5e, 0xbc, 0x48, 0xb3, 0x3c, 0xc6, 0x06,
	0xdd, 0x8b, 0x49, 0x26, 0xf8, 0xb1, 0x3c, 0xfa, 0xed, 0x24, 0x29, 0x1b, 0x53, 0x87, 0xb7, 0xdf,
	0xfb, 0xfd, 0x7b, 0x79, 0x10, 0xe8, 0xad, 0x48, 0xe6, 0xda, 0x90, 0xa3, 0x14, 0x64, 0x65, 0xb1,
	0xc1, 0xd9, 0x57, 0x04, 0xac, 0x20, 0xc9, 0x39, 0xc4, 0xeb, 0xaa, 0x46, 0x11, 0x8d, 0xa2, 0x49,
	0xaf, 0x0c, 0x98, 0x0f, 0x81, 0xa9, 0xa5, 0x11, 0xed, 0x40, 0x79, 0xc8, 0x05, 0x74, 0xb5, 0xa1,
	0x85, 0xa9, 0x6a, 0xc1, 0x02, 0xbb, 0x1d, 0x7d, 0xbe, 0x32, 0x0b, 0x2b, 0xe2, 0x26, 0xef, 0x31,
	0x3f, 0x83, 0xae, 0x75, 0xea, 0x95, 0x36, 0x4e, 0x74, 0x02, 0x9d, 0x58, 0xa7, 0x9e, 0x36, 0x6e,
	0x2b, 0xa0, 0x31, 0x22, 0xd9, 0x09, 0x33, 0x63, 0x7c, 0x8b, 0xd5, 0x38, 0x17, 0xdd, 0xa6, 0xc5,
	0x63, 0xbf, 0xd3, 0x2d, 0x6b, 0xf4, 0x2d, 0x47, 0xa3, 0x68, 0xc2, 0xca, 0xed, 0x98, 0x5d, 0xc0,
	0x49, 0x41, 0xf2, 0xb1, 0xaa, 0xb1, 0xc4, 0xb7, 0x0d, 0x5a, 0x77, 0xe8, 0x8a, 0xec, 0x0a, 0xfa,
	0x05, 0xc9, 0x12, 0xad, 0xa6, 0xb5, 0x45, 0x6f, 0x99, 0x93, 0x6a, 0x2c, 0x9d, 0x32, 0x60, 0x7e,
	0x0a, 0x6c, 0x45, 0x32, 0x1c, 0xda, 0x9f, 0xc6, 0xb9, 0xb7, 0x7b, 0x22, 0xbb, 0x86, 0x41, 0x41,
	0xf2, 0x61, 0x69, 0xdd, 0xbf, 0x71, 0x01, 0xf1, 0x8a, 0xa4, 0x15, 0xed, 0x11, 0xdb, 0xe5, 0x03,
	0x33, 0xfd, 0x8c, 0x20, 0xbe, 0x33, 0xb4, 0xe6, 0x19, 0xc4, 0xbe, 0x86, 0x27, 0xf9, 0xac, 0xd6,
	0xee, 0x23, 0x1d, 0xe6, 0xbf, 0x8a, 0xb3, 0x16, 0x1f, 0x03, 0xbb, 0x47, 0xc7, 0x07, 0xf9, 0xcf,
	0xa3, 0xd2, 0xe3, 0x7c, 0xef, 0xfd, 0x59, 0x8b, 0x0b, 0x60, 0x37, 0x4a, 0xf1, 0xb0, 0x27, 0xed,
	0xe5, 0x7b, 0xca, 0x39, 0x24, 0xcf, 0x5a, 0x55, 0x0e, 0x0f, 0x89, 0x63, 0x48, 0x4a, 0xac, 0xe9,
	0x1d, 0xff, 0x6e, 0xd8, 0xf7, 0xdd, 0x76, 0x5e, 0x98, 0xd1, 0x73, 0x99, 0x84, 0xff, 0x71, 0xf9,
	0x1d, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x24, 0x45, 0x11, 0x38, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CronClient is the client API for Cron service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CronClient interface {
	List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*JobListResponse, error)
	Get(ctx context.Context, in *JobNameRequest, opts ...grpc.CallOption) (*JobResponse, error)
	Add(ctx context.Context, in *Job, opts ...grpc.CallOption) (*Response, error)
	Update(ctx context.Context, in *Job, opts ...grpc.CallOption) (*Response, error)
	Remove(ctx context.Context, in *JobNameRequest, opts ...grpc.CallOption) (*Response, error)
}

type cronClient struct {
	cc *grpc.ClientConn
}

func NewCronClient(cc *grpc.ClientConn) CronClient {
	return &cronClient{cc}
}

func (c *cronClient) List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*JobListResponse, error) {
	out := new(JobListResponse)
	err := c.cc.Invoke(ctx, "/Cron/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cronClient) Get(ctx context.Context, in *JobNameRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, "/Cron/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cronClient) Add(ctx context.Context, in *Job, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Cron/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cronClient) Update(ctx context.Context, in *Job, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Cron/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cronClient) Remove(ctx context.Context, in *JobNameRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Cron/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CronServer is the server API for Cron service.
type CronServer interface {
	List(context.Context, *Empty) (*JobListResponse, error)
	Get(context.Context, *JobNameRequest) (*JobResponse, error)
	Add(context.Context, *Job) (*Response, error)
	Update(context.Context, *Job) (*Response, error)
	Remove(context.Context, *JobNameRequest) (*Response, error)
}

// UnimplementedCronServer can be embedded to have forward compatible implementations.
type UnimplementedCronServer struct {
}

func (*UnimplementedCronServer) List(ctx context.Context, req *Empty) (*JobListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedCronServer) Get(ctx context.Context, req *JobNameRequest) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedCronServer) Add(ctx context.Context, req *Job) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedCronServer) Update(ctx context.Context, req *Job) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedCronServer) Remove(ctx context.Context, req *JobNameRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}

func RegisterCronServer(s *grpc.Server, srv CronServer) {
	s.RegisterService(&_Cron_serviceDesc, srv)
}

func _Cron_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CronServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Cron/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CronServer).List(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cron_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CronServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Cron/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CronServer).Get(ctx, req.(*JobNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cron_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Job)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CronServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Cron/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CronServer).Add(ctx, req.(*Job))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cron_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Job)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CronServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Cron/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CronServer).Update(ctx, req.(*Job))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cron_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CronServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Cron/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CronServer).Remove(ctx, req.(*JobNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cron_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Cron",
	HandlerType: (*CronServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Cron_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Cron_Get_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _Cron_Add_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Cron_Update_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _Cron_Remove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "job.proto",
}