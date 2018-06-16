// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metadata/types/drone.proto

package pbtypes // import "openpitrix.io/openpitrix/pkg/pb/metadata/types"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DroneId struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DroneId) Reset()         { *m = DroneId{} }
func (m *DroneId) String() string { return proto.CompactTextString(m) }
func (*DroneId) ProtoMessage()    {}
func (*DroneId) Descriptor() ([]byte, []int) {
	return fileDescriptor_drone_e1268279a69a094c, []int{0}
}
func (m *DroneId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DroneId.Unmarshal(m, b)
}
func (m *DroneId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DroneId.Marshal(b, m, deterministic)
}
func (dst *DroneId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DroneId.Merge(dst, src)
}
func (m *DroneId) XXX_Size() int {
	return xxx_messageInfo_DroneId.Size(m)
}
func (m *DroneId) XXX_DiscardUnknown() {
	xxx_messageInfo_DroneId.DiscardUnknown(m)
}

var xxx_messageInfo_DroneId proto.InternalMessageInfo

func (m *DroneId) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DroneIdList struct {
	IdList               []string `protobuf:"bytes,1,rep,name=id_list,json=idList,proto3" json:"id_list"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DroneIdList) Reset()         { *m = DroneIdList{} }
func (m *DroneIdList) String() string { return proto.CompactTextString(m) }
func (*DroneIdList) ProtoMessage()    {}
func (*DroneIdList) Descriptor() ([]byte, []int) {
	return fileDescriptor_drone_e1268279a69a094c, []int{1}
}
func (m *DroneIdList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DroneIdList.Unmarshal(m, b)
}
func (m *DroneIdList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DroneIdList.Marshal(b, m, deterministic)
}
func (dst *DroneIdList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DroneIdList.Merge(dst, src)
}
func (m *DroneIdList) XXX_Size() int {
	return xxx_messageInfo_DroneIdList.Size(m)
}
func (m *DroneIdList) XXX_DiscardUnknown() {
	xxx_messageInfo_DroneIdList.DiscardUnknown(m)
}

var xxx_messageInfo_DroneIdList proto.InternalMessageInfo

func (m *DroneIdList) GetIdList() []string {
	if m != nil {
		return m.IdList
	}
	return nil
}

type DroneConfig struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Host                 string   `protobuf:"bytes,2,opt,name=host,proto3" json:"host"`
	ListenPort           int32    `protobuf:"varint,3,opt,name=listen_port,json=listenPort,proto3" json:"listen_port"`
	CmdInfoLogPath       string   `protobuf:"bytes,4,opt,name=cmd_info_log_path,json=cmdInfoLogPath,proto3" json:"cmd_info_log_path"`
	ConfdSelfHost        string   `protobuf:"bytes,5,opt,name=confd_self_host,json=confdSelfHost,proto3" json:"confd_self_host"`
	LogLevel             string   `protobuf:"bytes,6,opt,name=log_level,json=logLevel,proto3" json:"log_level"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DroneConfig) Reset()         { *m = DroneConfig{} }
func (m *DroneConfig) String() string { return proto.CompactTextString(m) }
func (*DroneConfig) ProtoMessage()    {}
func (*DroneConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_drone_e1268279a69a094c, []int{2}
}
func (m *DroneConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DroneConfig.Unmarshal(m, b)
}
func (m *DroneConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DroneConfig.Marshal(b, m, deterministic)
}
func (dst *DroneConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DroneConfig.Merge(dst, src)
}
func (m *DroneConfig) XXX_Size() int {
	return xxx_messageInfo_DroneConfig.Size(m)
}
func (m *DroneConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DroneConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DroneConfig proto.InternalMessageInfo

func (m *DroneConfig) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DroneConfig) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *DroneConfig) GetListenPort() int32 {
	if m != nil {
		return m.ListenPort
	}
	return 0
}

func (m *DroneConfig) GetCmdInfoLogPath() string {
	if m != nil {
		return m.CmdInfoLogPath
	}
	return ""
}

func (m *DroneConfig) GetConfdSelfHost() string {
	if m != nil {
		return m.ConfdSelfHost
	}
	return ""
}

func (m *DroneConfig) GetLogLevel() string {
	if m != nil {
		return m.LogLevel
	}
	return ""
}

type DroneEndpoint struct {
	FrontgateId          string   `protobuf:"bytes,1,opt,name=frontgate_id,json=frontgateId,proto3" json:"frontgate_id"`
	DroneIp              string   `protobuf:"bytes,2,opt,name=drone_ip,json=droneIp,proto3" json:"drone_ip"`
	DronePort            int32    `protobuf:"varint,3,opt,name=drone_port,json=dronePort,proto3" json:"drone_port"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DroneEndpoint) Reset()         { *m = DroneEndpoint{} }
func (m *DroneEndpoint) String() string { return proto.CompactTextString(m) }
func (*DroneEndpoint) ProtoMessage()    {}
func (*DroneEndpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_drone_e1268279a69a094c, []int{3}
}
func (m *DroneEndpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DroneEndpoint.Unmarshal(m, b)
}
func (m *DroneEndpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DroneEndpoint.Marshal(b, m, deterministic)
}
func (dst *DroneEndpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DroneEndpoint.Merge(dst, src)
}
func (m *DroneEndpoint) XXX_Size() int {
	return xxx_messageInfo_DroneEndpoint.Size(m)
}
func (m *DroneEndpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_DroneEndpoint.DiscardUnknown(m)
}

var xxx_messageInfo_DroneEndpoint proto.InternalMessageInfo

func (m *DroneEndpoint) GetFrontgateId() string {
	if m != nil {
		return m.FrontgateId
	}
	return ""
}

func (m *DroneEndpoint) GetDroneIp() string {
	if m != nil {
		return m.DroneIp
	}
	return ""
}

func (m *DroneEndpoint) GetDronePort() int32 {
	if m != nil {
		return m.DronePort
	}
	return 0
}

type SetDroneConfigRequest struct {
	Endpoint             *DroneEndpoint `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint"`
	Config               *DroneConfig   `protobuf:"bytes,2,opt,name=config,proto3" json:"config"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SetDroneConfigRequest) Reset()         { *m = SetDroneConfigRequest{} }
func (m *SetDroneConfigRequest) String() string { return proto.CompactTextString(m) }
func (*SetDroneConfigRequest) ProtoMessage()    {}
func (*SetDroneConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_drone_e1268279a69a094c, []int{4}
}
func (m *SetDroneConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetDroneConfigRequest.Unmarshal(m, b)
}
func (m *SetDroneConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetDroneConfigRequest.Marshal(b, m, deterministic)
}
func (dst *SetDroneConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetDroneConfigRequest.Merge(dst, src)
}
func (m *SetDroneConfigRequest) XXX_Size() int {
	return xxx_messageInfo_SetDroneConfigRequest.Size(m)
}
func (m *SetDroneConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetDroneConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetDroneConfigRequest proto.InternalMessageInfo

func (m *SetDroneConfigRequest) GetEndpoint() *DroneEndpoint {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

func (m *SetDroneConfigRequest) GetConfig() *DroneConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

type RunCommandOnDroneRequest struct {
	Endpoint             *DroneEndpoint `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint"`
	Command              string         `protobuf:"bytes,2,opt,name=command,proto3" json:"command"`
	TimeoutSeconds       int32          `protobuf:"varint,3,opt,name=timeout_seconds,json=timeoutSeconds,proto3" json:"timeout_seconds"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *RunCommandOnDroneRequest) Reset()         { *m = RunCommandOnDroneRequest{} }
func (m *RunCommandOnDroneRequest) String() string { return proto.CompactTextString(m) }
func (*RunCommandOnDroneRequest) ProtoMessage()    {}
func (*RunCommandOnDroneRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_drone_e1268279a69a094c, []int{5}
}
func (m *RunCommandOnDroneRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunCommandOnDroneRequest.Unmarshal(m, b)
}
func (m *RunCommandOnDroneRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunCommandOnDroneRequest.Marshal(b, m, deterministic)
}
func (dst *RunCommandOnDroneRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunCommandOnDroneRequest.Merge(dst, src)
}
func (m *RunCommandOnDroneRequest) XXX_Size() int {
	return xxx_messageInfo_RunCommandOnDroneRequest.Size(m)
}
func (m *RunCommandOnDroneRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RunCommandOnDroneRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RunCommandOnDroneRequest proto.InternalMessageInfo

func (m *RunCommandOnDroneRequest) GetEndpoint() *DroneEndpoint {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

func (m *RunCommandOnDroneRequest) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *RunCommandOnDroneRequest) GetTimeoutSeconds() int32 {
	if m != nil {
		return m.TimeoutSeconds
	}
	return 0
}

func init() {
	proto.RegisterType((*DroneId)(nil), "metadata.types.DroneId")
	proto.RegisterType((*DroneIdList)(nil), "metadata.types.DroneIdList")
	proto.RegisterType((*DroneConfig)(nil), "metadata.types.DroneConfig")
	proto.RegisterType((*DroneEndpoint)(nil), "metadata.types.DroneEndpoint")
	proto.RegisterType((*SetDroneConfigRequest)(nil), "metadata.types.SetDroneConfigRequest")
	proto.RegisterType((*RunCommandOnDroneRequest)(nil), "metadata.types.RunCommandOnDroneRequest")
}

func init() { proto.RegisterFile("metadata/types/drone.proto", fileDescriptor_drone_e1268279a69a094c) }

var fileDescriptor_drone_e1268279a69a094c = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0xd3, 0x36, 0x1f, 0x63, 0x9a, 0x8a, 0x95, 0x10, 0x6e, 0xab, 0x8a, 0xe0, 0x43, 0x09,
	0x17, 0x5b, 0x6a, 0x25, 0x04, 0xe2, 0x46, 0x41, 0x22, 0x52, 0x24, 0x2a, 0xe7, 0xc6, 0x65, 0xe5,
	0x78, 0xd7, 0xce, 0x0a, 0x7b, 0x67, 0xb1, 0x27, 0x08, 0x7e, 0x01, 0x7f, 0x81, 0xdf, 0xc3, 0x2f,
	0x43, 0x1e, 0xbb, 0xa1, 0xa9, 0x38, 0x72, 0xb2, 0xe7, 0xbd, 0xb7, 0x6f, 0xe6, 0xcd, 0x2e, 0x9c,
	0x55, 0x9a, 0x52, 0x95, 0x52, 0x1a, 0xd3, 0x0f, 0xa7, 0x9b, 0x58, 0xd5, 0x68, 0x75, 0xe4, 0x6a,
	0x24, 0x14, 0xd3, 0x3b, 0x2e, 0x62, 0xee, 0xec, 0xa1, 0x36, 0x43, 0x9b, 0xab, 0x4e, 0x1b, 0x9e,
	0xc2, 0xe8, 0x7d, 0x7b, 0x74, 0xa1, 0xc4, 0x14, 0x06, 0x46, 0x05, 0xde, 0xcc, 0x9b, 0x4f, 0x92,
	0x81, 0x51, 0xe1, 0x25, 0xf8, 0x3d, 0xb5, 0x34, 0x0d, 0x89, 0xa7, 0x30, 0x32, 0x4a, 0x96, 0xa6,
	0xa1, 0xc0, 0x9b, 0x1d, 0xcc, 0x27, 0xc9, 0xd0, 0x30, 0x11, 0xfe, 0xf6, 0x7a, 0xe1, 0x0d, 0xda,
	0xdc, 0x14, 0x0f, 0x7d, 0x84, 0x80, 0xc3, 0x0d, 0x36, 0x14, 0x0c, 0x18, 0xe1, 0x7f, 0xf1, 0x0c,
	0xfc, 0xd6, 0x49, 0x5b, 0xe9, 0xb0, 0xa6, 0xe0, 0x60, 0xe6, 0xcd, 0x8f, 0x12, 0xe8, 0xa0, 0x5b,
	0xac, 0x49, 0xbc, 0x84, 0xc7, 0x59, 0xa5, 0xa4, 0xb1, 0x39, 0xca, 0x12, 0x0b, 0xe9, 0x52, 0xda,
	0x04, 0x87, 0xec, 0x30, 0xcd, 0x2a, 0xb5, 0xb0, 0x39, 0x2e, 0xb1, 0xb8, 0x4d, 0x69, 0x23, 0x2e,
	0xe1, 0x84, 0x13, 0xc9, 0x46, 0x97, 0xb9, 0xe4, 0x56, 0x47, 0x2c, 0x3c, 0x66, 0x78, 0xa5, 0xcb,
	0xfc, 0x63, 0xdb, 0xf3, 0x1c, 0x26, 0xad, 0x53, 0xa9, 0xbf, 0xe9, 0x32, 0x18, 0xb2, 0x62, 0x5c,
	0x62, 0xb1, 0x6c, 0xeb, 0xb0, 0x84, 0x63, 0xce, 0xf0, 0xc1, 0x2a, 0x87, 0xc6, 0x92, 0x78, 0x0e,
	0x8f, 0xf2, 0x1a, 0x2d, 0x15, 0x29, 0x69, 0xb9, 0xcb, 0xe3, 0xef, 0xb0, 0x85, 0x12, 0xa7, 0x30,
	0xe6, 0xb5, 0x4b, 0xe3, 0xfa, 0x70, 0x23, 0xae, 0x17, 0x4e, 0x5c, 0x00, 0x74, 0xd4, 0xbd, 0x78,
	0x13, 0x46, 0xda, 0x74, 0xe1, 0x4f, 0x0f, 0x9e, 0xac, 0x34, 0xdd, 0xdb, 0x5a, 0xa2, 0xbf, 0x6e,
	0x75, 0x43, 0xe2, 0x0d, 0x8c, 0x75, 0x3f, 0x02, 0xb7, 0xf4, 0xaf, 0x2e, 0xa2, 0xfd, 0xeb, 0x8c,
	0xf6, 0xe6, 0x4c, 0x76, 0x72, 0x71, 0x0d, 0xc3, 0x8c, 0xbd, 0x78, 0x18, 0xff, 0xea, 0xfc, 0x9f,
	0x07, 0xfb, 0x76, 0xbd, 0x34, 0xfc, 0xe5, 0x41, 0x90, 0x6c, 0xed, 0x0d, 0x56, 0x55, 0x6a, 0xd5,
	0x27, 0xcb, 0x9a, 0xff, 0x30, 0x4c, 0x00, 0xa3, 0xac, 0xf3, 0xbc, 0x5b, 0x4d, 0x5f, 0x8a, 0x17,
	0x70, 0x42, 0xa6, 0xd2, 0xb8, 0x25, 0xd9, 0xe8, 0x0c, 0xad, 0x6a, 0xfa, 0xfd, 0x4c, 0x7b, 0x78,
	0xd5, 0xa1, 0xef, 0x5e, 0x7f, 0x7e, 0x85, 0x4e, 0x5b, 0x67, 0xa8, 0x36, 0xdf, 0x23, 0x83, 0xf1,
	0xdf, 0x2a, 0x76, 0x5f, 0x8a, 0xd8, 0xad, 0xe3, 0xfd, 0x87, 0xfd, 0xd6, 0xad, 0xf9, 0xbb, 0x1e,
	0xf2, 0xdb, 0xbe, 0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0x47, 0xac, 0x6f, 0xba, 0x25, 0x03, 0x00,
	0x00,
}