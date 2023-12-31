// Code generated by protoc-gen-go. DO NOT EDIT.
// source: log.proto

package grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LogOutRequest struct {
	TownCode             string   `protobuf:"bytes,1,opt,name=town_code,json=townCode,proto3" json:"town_code,omitempty"`
	PlatNumber           string   `protobuf:"bytes,2,opt,name=plat_number,json=platNumber,proto3" json:"plat_number,omitempty"`
	TimeOut              string   `protobuf:"bytes,3,opt,name=time_out,json=timeOut,proto3" json:"time_out,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogOutRequest) Reset()         { *m = LogOutRequest{} }
func (m *LogOutRequest) String() string { return proto.CompactTextString(m) }
func (*LogOutRequest) ProtoMessage()    {}
func (*LogOutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_log_6d3cffd1db52ac50, []int{0}
}
func (m *LogOutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogOutRequest.Unmarshal(m, b)
}
func (m *LogOutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogOutRequest.Marshal(b, m, deterministic)
}
func (dst *LogOutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogOutRequest.Merge(dst, src)
}
func (m *LogOutRequest) XXX_Size() int {
	return xxx_messageInfo_LogOutRequest.Size(m)
}
func (m *LogOutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogOutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogOutRequest proto.InternalMessageInfo

func (m *LogOutRequest) GetTownCode() string {
	if m != nil {
		return m.TownCode
	}
	return ""
}

func (m *LogOutRequest) GetPlatNumber() string {
	if m != nil {
		return m.PlatNumber
	}
	return ""
}

func (m *LogOutRequest) GetTimeOut() string {
	if m != nil {
		return m.TimeOut
	}
	return ""
}

type Log struct {
	TownName             string   `protobuf:"bytes,1,opt,name=town_name,json=townName,proto3" json:"town_name,omitempty"`
	TotalTime            float32  `protobuf:"fixed32,2,opt,name=total_time,json=totalTime,proto3" json:"total_time,omitempty"`
	TimeIn               string   `protobuf:"bytes,3,opt,name=time_in,json=timeIn,proto3" json:"time_in,omitempty"`
	TimeOut              string   `protobuf:"bytes,4,opt,name=time_out,json=timeOut,proto3" json:"time_out,omitempty"`
	DateAt               string   `protobuf:"bytes,5,opt,name=date_at,json=dateAt,proto3" json:"date_at,omitempty"`
	PlatNumber           string   `protobuf:"bytes,6,opt,name=plat_number,json=platNumber,proto3" json:"plat_number,omitempty"`
	VehicleType          string   `protobuf:"bytes,7,opt,name=vehicle_type,json=vehicleType,proto3" json:"vehicle_type,omitempty"`
	DateOutAt            string   `protobuf:"bytes,8,opt,name=date_out_at,json=dateOutAt,proto3" json:"date_out_at,omitempty"`
	TownCode             string   `protobuf:"bytes,9,opt,name=town_code,json=townCode,proto3" json:"town_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Log) Reset()         { *m = Log{} }
func (m *Log) String() string { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()    {}
func (*Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_log_6d3cffd1db52ac50, []int{1}
}
func (m *Log) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Log.Unmarshal(m, b)
}
func (m *Log) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Log.Marshal(b, m, deterministic)
}
func (dst *Log) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Log.Merge(dst, src)
}
func (m *Log) XXX_Size() int {
	return xxx_messageInfo_Log.Size(m)
}
func (m *Log) XXX_DiscardUnknown() {
	xxx_messageInfo_Log.DiscardUnknown(m)
}

var xxx_messageInfo_Log proto.InternalMessageInfo

func (m *Log) GetTownName() string {
	if m != nil {
		return m.TownName
	}
	return ""
}

func (m *Log) GetTotalTime() float32 {
	if m != nil {
		return m.TotalTime
	}
	return 0
}

func (m *Log) GetTimeIn() string {
	if m != nil {
		return m.TimeIn
	}
	return ""
}

func (m *Log) GetTimeOut() string {
	if m != nil {
		return m.TimeOut
	}
	return ""
}

func (m *Log) GetDateAt() string {
	if m != nil {
		return m.DateAt
	}
	return ""
}

func (m *Log) GetPlatNumber() string {
	if m != nil {
		return m.PlatNumber
	}
	return ""
}

func (m *Log) GetVehicleType() string {
	if m != nil {
		return m.VehicleType
	}
	return ""
}

func (m *Log) GetDateOutAt() string {
	if m != nil {
		return m.DateOutAt
	}
	return ""
}

func (m *Log) GetTownCode() string {
	if m != nil {
		return m.TownCode
	}
	return ""
}

type GetLogReport struct {
	Date                 string   `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	TownCode             string   `protobuf:"bytes,2,opt,name=town_code,json=townCode,proto3" json:"town_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLogReport) Reset()         { *m = GetLogReport{} }
func (m *GetLogReport) String() string { return proto.CompactTextString(m) }
func (*GetLogReport) ProtoMessage()    {}
func (*GetLogReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_log_6d3cffd1db52ac50, []int{2}
}
func (m *GetLogReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLogReport.Unmarshal(m, b)
}
func (m *GetLogReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLogReport.Marshal(b, m, deterministic)
}
func (dst *GetLogReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLogReport.Merge(dst, src)
}
func (m *GetLogReport) XXX_Size() int {
	return xxx_messageInfo_GetLogReport.Size(m)
}
func (m *GetLogReport) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLogReport.DiscardUnknown(m)
}

var xxx_messageInfo_GetLogReport proto.InternalMessageInfo

func (m *GetLogReport) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *GetLogReport) GetTownCode() string {
	if m != nil {
		return m.TownCode
	}
	return ""
}

type TownLogReport struct {
	TownCode             string   `protobuf:"bytes,1,opt,name=town_code,json=townCode,proto3" json:"town_code,omitempty"`
	VehicleType          string   `protobuf:"bytes,3,opt,name=vehicle_type,json=vehicleType,proto3" json:"vehicle_type,omitempty"`
	TotalLogged          int64    `protobuf:"varint,2,opt,name=total_logged,json=totalLogged,proto3" json:"total_logged,omitempty"`
	Quantity             int64    `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TownLogReport) Reset()         { *m = TownLogReport{} }
func (m *TownLogReport) String() string { return proto.CompactTextString(m) }
func (*TownLogReport) ProtoMessage()    {}
func (*TownLogReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_log_6d3cffd1db52ac50, []int{3}
}
func (m *TownLogReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TownLogReport.Unmarshal(m, b)
}
func (m *TownLogReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TownLogReport.Marshal(b, m, deterministic)
}
func (dst *TownLogReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TownLogReport.Merge(dst, src)
}
func (m *TownLogReport) XXX_Size() int {
	return xxx_messageInfo_TownLogReport.Size(m)
}
func (m *TownLogReport) XXX_DiscardUnknown() {
	xxx_messageInfo_TownLogReport.DiscardUnknown(m)
}

var xxx_messageInfo_TownLogReport proto.InternalMessageInfo

func (m *TownLogReport) GetTownCode() string {
	if m != nil {
		return m.TownCode
	}
	return ""
}

func (m *TownLogReport) GetVehicleType() string {
	if m != nil {
		return m.VehicleType
	}
	return ""
}

func (m *TownLogReport) GetTotalLogged() int64 {
	if m != nil {
		return m.TotalLogged
	}
	return 0
}

func (m *TownLogReport) GetQuantity() int64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

type TownLogReportList struct {
	Date                 string           `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	List                 []*TownLogReport `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TownLogReportList) Reset()         { *m = TownLogReportList{} }
func (m *TownLogReportList) String() string { return proto.CompactTextString(m) }
func (*TownLogReportList) ProtoMessage()    {}
func (*TownLogReportList) Descriptor() ([]byte, []int) {
	return fileDescriptor_log_6d3cffd1db52ac50, []int{4}
}
func (m *TownLogReportList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TownLogReportList.Unmarshal(m, b)
}
func (m *TownLogReportList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TownLogReportList.Marshal(b, m, deterministic)
}
func (dst *TownLogReportList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TownLogReportList.Merge(dst, src)
}
func (m *TownLogReportList) XXX_Size() int {
	return xxx_messageInfo_TownLogReportList.Size(m)
}
func (m *TownLogReportList) XXX_DiscardUnknown() {
	xxx_messageInfo_TownLogReportList.DiscardUnknown(m)
}

var xxx_messageInfo_TownLogReportList proto.InternalMessageInfo

func (m *TownLogReportList) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *TownLogReportList) GetList() []*TownLogReport {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*LogOutRequest)(nil), "grpc.LogOutRequest")
	proto.RegisterType((*Log)(nil), "grpc.Log")
	proto.RegisterType((*GetLogReport)(nil), "grpc.GetLogReport")
	proto.RegisterType((*TownLogReport)(nil), "grpc.TownLogReport")
	proto.RegisterType((*TownLogReportList)(nil), "grpc.TownLogReportList")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LogsClient is the client API for Logs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogsClient interface {
	LogOutVehicle(ctx context.Context, in *LogOutRequest, opts ...grpc.CallOption) (*Log, error)
	ReportLogged(ctx context.Context, in *GetLogReport, opts ...grpc.CallOption) (*TownLogReportList, error)
}

type logsClient struct {
	cc *grpc.ClientConn
}

func NewLogsClient(cc *grpc.ClientConn) LogsClient {
	return &logsClient{cc}
}

func (c *logsClient) LogOutVehicle(ctx context.Context, in *LogOutRequest, opts ...grpc.CallOption) (*Log, error) {
	out := new(Log)
	err := c.cc.Invoke(ctx, "/grpc.Logs/LogOutVehicle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logsClient) ReportLogged(ctx context.Context, in *GetLogReport, opts ...grpc.CallOption) (*TownLogReportList, error) {
	out := new(TownLogReportList)
	err := c.cc.Invoke(ctx, "/grpc.Logs/ReportLogged", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogsServer is the server API for Logs service.
type LogsServer interface {
	LogOutVehicle(context.Context, *LogOutRequest) (*Log, error)
	ReportLogged(context.Context, *GetLogReport) (*TownLogReportList, error)
}

func RegisterLogsServer(s *grpc.Server, srv LogsServer) {
	s.RegisterService(&_Logs_serviceDesc, srv)
}

func _Logs_LogOutVehicle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogOutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogsServer).LogOutVehicle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Logs/LogOutVehicle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogsServer).LogOutVehicle(ctx, req.(*LogOutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logs_ReportLogged_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLogReport)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogsServer).ReportLogged(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Logs/ReportLogged",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogsServer).ReportLogged(ctx, req.(*GetLogReport))
	}
	return interceptor(ctx, in, info, handler)
}

var _Logs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Logs",
	HandlerType: (*LogsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LogOutVehicle",
			Handler:    _Logs_LogOutVehicle_Handler,
		},
		{
			MethodName: "ReportLogged",
			Handler:    _Logs_ReportLogged_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "log.proto",
}

func init() { proto.RegisterFile("log.proto", fileDescriptor_log_6d3cffd1db52ac50) }

var fileDescriptor_log_6d3cffd1db52ac50 = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x6d, 0x6c, 0x93, 0xc6, 0xe3, 0xf4, 0xc0, 0x72, 0xa8, 0x09, 0x02, 0x5a, 0x5f, 0xe8, 0x29,
	0x12, 0xe5, 0x8c, 0x50, 0xc5, 0x01, 0x21, 0x59, 0x0d, 0xb2, 0x22, 0xae, 0x96, 0x9b, 0x8c, 0xb6,
	0x96, 0x6c, 0x8f, 0x1b, 0xcf, 0x52, 0xf2, 0x13, 0xf0, 0xcb, 0x68, 0xc7, 0x26, 0xed, 0xba, 0x88,
	0x5b, 0xe6, 0xed, 0xbe, 0xf7, 0xe6, 0xbd, 0xac, 0x21, 0xac, 0x48, 0x2f, 0xdb, 0x1d, 0x31, 0xa9,
	0x40, 0xef, 0xda, 0x4d, 0x72, 0x0b, 0x27, 0x29, 0xe9, 0x95, 0xe1, 0x0c, 0xef, 0x0c, 0x76, 0xac,
	0x5e, 0x41, 0xc8, 0x74, 0xdf, 0xe4, 0x1b, 0xda, 0x62, 0x3c, 0x39, 0x9b, 0x5c, 0x84, 0xd9, 0xcc,
	0x02, 0x9f, 0x69, 0x8b, 0xea, 0x2d, 0x44, 0x6d, 0x55, 0x70, 0xde, 0x98, 0xfa, 0x06, 0x77, 0xb1,
	0x27, 0xc7, 0x60, 0xa1, 0x6b, 0x41, 0xd4, 0x4b, 0x98, 0x71, 0x59, 0x63, 0x4e, 0x86, 0x63, 0x5f,
	0x4e, 0x8f, 0xed, 0xbc, 0x32, 0x9c, 0xfc, 0xf6, 0xc0, 0x4f, 0x49, 0x1f, 0x0c, 0x9a, 0xa2, 0x76,
	0x0c, 0xae, 0x8b, 0x1a, 0xd5, 0x6b, 0x00, 0x26, 0x2e, 0xaa, 0xdc, 0xb2, 0x44, 0xdf, 0xcb, 0x42,
	0x41, 0xd6, 0x65, 0x8d, 0xea, 0x14, 0x44, 0x2e, 0x2f, 0x9b, 0x41, 0x7d, 0x6a, 0xc7, 0xaf, 0x8d,
	0xe3, 0x1b, 0x38, 0xbe, 0x96, 0xb3, 0x2d, 0x18, 0xf3, 0x82, 0xe3, 0x67, 0x3d, 0xc7, 0x8e, 0x57,
	0x3c, 0x0e, 0x33, 0x7d, 0x12, 0xe6, 0x1c, 0xe6, 0x3f, 0xf0, 0xb6, 0xdc, 0x54, 0x98, 0xf3, 0xbe,
	0xc5, 0xf8, 0x58, 0x6e, 0x44, 0x03, 0xb6, 0xde, 0xb7, 0xa8, 0xde, 0x40, 0x24, 0xe2, 0x64, 0xd8,
	0x1a, 0xcc, 0xe4, 0x46, 0x68, 0xa1, 0x95, 0xe1, 0xab, 0x51, 0x9b, 0xa1, 0xdb, 0x66, 0xf2, 0x09,
	0xe6, 0x5f, 0x90, 0x53, 0xd2, 0x19, 0xb6, 0xb4, 0x63, 0xa5, 0x20, 0xb0, 0xcc, 0xa1, 0x14, 0xf9,
	0xed, 0x0a, 0x78, 0x23, 0x81, 0x5f, 0x13, 0x38, 0x59, 0xd3, 0x7d, 0xf3, 0x20, 0xf1, 0xdf, 0x7f,
	0x6f, 0x9c, 0xc7, 0x7f, 0x9a, 0xe7, 0x1c, 0xe6, 0x7d, 0xff, 0x15, 0x69, 0x8d, 0x5b, 0x71, 0xf4,
	0xb3, 0x48, 0xb0, 0x54, 0x20, 0xb5, 0x80, 0xd9, 0x9d, 0x29, 0x1a, 0x2e, 0x79, 0x2f, 0x55, 0xfb,
	0xd9, 0x61, 0x4e, 0xbe, 0xc1, 0x73, 0x67, 0x9f, 0xb4, 0xec, 0x1e, 0x62, 0x79, 0x8f, 0x62, 0xbd,
	0x83, 0xa0, 0x2a, 0x3b, 0x8e, 0x27, 0x67, 0xfe, 0x45, 0x74, 0xf9, 0x62, 0x69, 0xdf, 0xe2, 0xd2,
	0xa1, 0x66, 0x72, 0xe1, 0xf2, 0x27, 0x04, 0x29, 0xe9, 0x4e, 0xbd, 0xff, 0xfb, 0x4e, 0xbf, 0xf7,
	0xdb, 0xaa, 0x81, 0xe3, 0x3c, 0xde, 0x45, 0x78, 0x00, 0x93, 0x23, 0xf5, 0x11, 0xe6, 0xc3, 0x16,
	0xfd, 0xe2, 0xaa, 0x3f, 0x7c, 0x5c, 0xf9, 0xe2, 0xf4, 0x1f, 0xce, 0x76, 0xe9, 0xe4, 0xe8, 0x66,
	0x2a, 0x9f, 0xc9, 0x87, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x13, 0x01, 0xa4, 0x33, 0x03,
	0x00, 0x00,
}
