
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
//由Protoc Gen Go生成的代码。不要编辑。
//来源：peer/admin.proto

package peer //导入“github.com/hyperledger/fabric/protos/peer”

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import common "github.com/hyperledger/fabric/protos/common"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

//引用导入以禁止错误（如果未使用）。
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

//这是一个编译时断言，以确保生成的文件
//与正在编译的proto包兼容。
//此行的编译错误可能意味着您的
//需要更新proto包。
const _ = proto.ProtoPackageIsVersion2 //请升级proto包

type ServerStatus_StatusCode int32

const (
	ServerStatus_UNDEFINED ServerStatus_StatusCode = 0
	ServerStatus_STARTED   ServerStatus_StatusCode = 1
	ServerStatus_STOPPED   ServerStatus_StatusCode = 2
	ServerStatus_PAUSED    ServerStatus_StatusCode = 3
	ServerStatus_ERROR     ServerStatus_StatusCode = 4
	ServerStatus_UNKNOWN   ServerStatus_StatusCode = 5
)

var ServerStatus_StatusCode_name = map[int32]string{
	0: "UNDEFINED",
	1: "STARTED",
	2: "STOPPED",
	3: "PAUSED",
	4: "ERROR",
	5: "UNKNOWN",
}
var ServerStatus_StatusCode_value = map[string]int32{
	"UNDEFINED": 0,
	"STARTED":   1,
	"STOPPED":   2,
	"PAUSED":    3,
	"ERROR":     4,
	"UNKNOWN":   5,
}

func (x ServerStatus_StatusCode) String() string {
	return proto.EnumName(ServerStatus_StatusCode_name, int32(x))
}
func (ServerStatus_StatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_admin_b2904393863b6bc5, []int{0, 0}
}

type ServerStatus struct {
	Status               ServerStatus_StatusCode `protobuf:"varint,1,opt,name=status,proto3,enum=protos.ServerStatus_StatusCode" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ServerStatus) Reset()         { *m = ServerStatus{} }
func (m *ServerStatus) String() string { return proto.CompactTextString(m) }
func (*ServerStatus) ProtoMessage()    {}
func (*ServerStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_admin_b2904393863b6bc5, []int{0}
}
func (m *ServerStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerStatus.Unmarshal(m, b)
}
func (m *ServerStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerStatus.Marshal(b, m, deterministic)
}
func (dst *ServerStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerStatus.Merge(dst, src)
}
func (m *ServerStatus) XXX_Size() int {
	return xxx_messageInfo_ServerStatus.Size(m)
}
func (m *ServerStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ServerStatus proto.InternalMessageInfo

func (m *ServerStatus) GetStatus() ServerStatus_StatusCode {
	if m != nil {
		return m.Status
	}
	return ServerStatus_UNDEFINED
}

type LogLevelRequest struct {
	LogModule            string   `protobuf:"bytes,1,opt,name=log_module,json=logModule,proto3" json:"log_module,omitempty"`
	LogLevel             string   `protobuf:"bytes,2,opt,name=log_level,json=logLevel,proto3" json:"log_level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogLevelRequest) Reset()         { *m = LogLevelRequest{} }
func (m *LogLevelRequest) String() string { return proto.CompactTextString(m) }
func (*LogLevelRequest) ProtoMessage()    {}
func (*LogLevelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_admin_b2904393863b6bc5, []int{1}
}
func (m *LogLevelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogLevelRequest.Unmarshal(m, b)
}
func (m *LogLevelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogLevelRequest.Marshal(b, m, deterministic)
}
func (dst *LogLevelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogLevelRequest.Merge(dst, src)
}
func (m *LogLevelRequest) XXX_Size() int {
	return xxx_messageInfo_LogLevelRequest.Size(m)
}
func (m *LogLevelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogLevelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogLevelRequest proto.InternalMessageInfo

func (m *LogLevelRequest) GetLogModule() string {
	if m != nil {
		return m.LogModule
	}
	return ""
}

func (m *LogLevelRequest) GetLogLevel() string {
	if m != nil {
		return m.LogLevel
	}
	return ""
}

type LogLevelResponse struct {
	LogModule            string   `protobuf:"bytes,1,opt,name=log_module,json=logModule,proto3" json:"log_module,omitempty"`
	LogLevel             string   `protobuf:"bytes,2,opt,name=log_level,json=logLevel,proto3" json:"log_level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogLevelResponse) Reset()         { *m = LogLevelResponse{} }
func (m *LogLevelResponse) String() string { return proto.CompactTextString(m) }
func (*LogLevelResponse) ProtoMessage()    {}
func (*LogLevelResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_admin_b2904393863b6bc5, []int{2}
}
func (m *LogLevelResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogLevelResponse.Unmarshal(m, b)
}
func (m *LogLevelResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogLevelResponse.Marshal(b, m, deterministic)
}
func (dst *LogLevelResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogLevelResponse.Merge(dst, src)
}
func (m *LogLevelResponse) XXX_Size() int {
	return xxx_messageInfo_LogLevelResponse.Size(m)
}
func (m *LogLevelResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogLevelResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogLevelResponse proto.InternalMessageInfo

func (m *LogLevelResponse) GetLogModule() string {
	if m != nil {
		return m.LogModule
	}
	return ""
}

func (m *LogLevelResponse) GetLogLevel() string {
	if m != nil {
		return m.LogLevel
	}
	return ""
}

type LogSpecRequest struct {
	LogSpec              string   `protobuf:"bytes,1,opt,name=log_spec,json=logSpec,proto3" json:"log_spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogSpecRequest) Reset()         { *m = LogSpecRequest{} }
func (m *LogSpecRequest) String() string { return proto.CompactTextString(m) }
func (*LogSpecRequest) ProtoMessage()    {}
func (*LogSpecRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_admin_b2904393863b6bc5, []int{3}
}
func (m *LogSpecRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogSpecRequest.Unmarshal(m, b)
}
func (m *LogSpecRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogSpecRequest.Marshal(b, m, deterministic)
}
func (dst *LogSpecRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogSpecRequest.Merge(dst, src)
}
func (m *LogSpecRequest) XXX_Size() int {
	return xxx_messageInfo_LogSpecRequest.Size(m)
}
func (m *LogSpecRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogSpecRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogSpecRequest proto.InternalMessageInfo

func (m *LogSpecRequest) GetLogSpec() string {
	if m != nil {
		return m.LogSpec
	}
	return ""
}

type LogSpecResponse struct {
	LogSpec              string   `protobuf:"bytes,1,opt,name=log_spec,json=logSpec,proto3" json:"log_spec,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogSpecResponse) Reset()         { *m = LogSpecResponse{} }
func (m *LogSpecResponse) String() string { return proto.CompactTextString(m) }
func (*LogSpecResponse) ProtoMessage()    {}
func (*LogSpecResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_admin_b2904393863b6bc5, []int{4}
}
func (m *LogSpecResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogSpecResponse.Unmarshal(m, b)
}
func (m *LogSpecResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogSpecResponse.Marshal(b, m, deterministic)
}
func (dst *LogSpecResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogSpecResponse.Merge(dst, src)
}
func (m *LogSpecResponse) XXX_Size() int {
	return xxx_messageInfo_LogSpecResponse.Size(m)
}
func (m *LogSpecResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogSpecResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogSpecResponse proto.InternalMessageInfo

func (m *LogSpecResponse) GetLogSpec() string {
	if m != nil {
		return m.LogSpec
	}
	return ""
}

func (m *LogSpecResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type AdminOperation struct {
//有效分配给内容的类型：
//*管理员操作日志
//*管理员操作日志规范
	Content              isAdminOperation_Content `protobuf_oneof:"content"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *AdminOperation) Reset()         { *m = AdminOperation{} }
func (m *AdminOperation) String() string { return proto.CompactTextString(m) }
func (*AdminOperation) ProtoMessage()    {}
func (*AdminOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_admin_b2904393863b6bc5, []int{5}
}
func (m *AdminOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdminOperation.Unmarshal(m, b)
}
func (m *AdminOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdminOperation.Marshal(b, m, deterministic)
}
func (dst *AdminOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdminOperation.Merge(dst, src)
}
func (m *AdminOperation) XXX_Size() int {
	return xxx_messageInfo_AdminOperation.Size(m)
}
func (m *AdminOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_AdminOperation.DiscardUnknown(m)
}

var xxx_messageInfo_AdminOperation proto.InternalMessageInfo

type isAdminOperation_Content interface {
	isAdminOperation_Content()
}

type AdminOperation_LogReq struct {
	LogReq *LogLevelRequest `protobuf:"bytes,1,opt,name=logReq,proto3,oneof"`
}

type AdminOperation_LogSpecReq struct {
	LogSpecReq *LogSpecRequest `protobuf:"bytes,2,opt,name=logSpecReq,proto3,oneof"`
}

func (*AdminOperation_LogReq) isAdminOperation_Content() {}

func (*AdminOperation_LogSpecReq) isAdminOperation_Content() {}

func (m *AdminOperation) GetContent() isAdminOperation_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *AdminOperation) GetLogReq() *LogLevelRequest {
	if x, ok := m.GetContent().(*AdminOperation_LogReq); ok {
		return x.LogReq
	}
	return nil
}

func (m *AdminOperation) GetLogSpecReq() *LogSpecRequest {
	if x, ok := m.GetContent().(*AdminOperation_LogSpecReq); ok {
		return x.LogSpecReq
	}
	return nil
}

//xxxoneoffuncs用于Proto包的内部使用。
func (*AdminOperation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AdminOperation_OneofMarshaler, _AdminOperation_OneofUnmarshaler, _AdminOperation_OneofSizer, []interface{}{
		(*AdminOperation_LogReq)(nil),
		(*AdminOperation_LogSpecReq)(nil),
	}
}

func _AdminOperation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AdminOperation)
//内容
	switch x := m.Content.(type) {
	case *AdminOperation_LogReq:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LogReq); err != nil {
			return err
		}
	case *AdminOperation_LogSpecReq:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LogSpecReq); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("AdminOperation.Content has unexpected type %T", x)
	}
	return nil
}

func _AdminOperation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AdminOperation)
	switch tag {
case 1: //洛格雷克
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LogLevelRequest)
		err := b.DecodeMessage(msg)
		m.Content = &AdminOperation_LogReq{msg}
		return true, err
case 2: //content.logspecreq
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LogSpecRequest)
		err := b.DecodeMessage(msg)
		m.Content = &AdminOperation_LogSpecReq{msg}
		return true, err
	default:
		return false, nil
	}
}

func _AdminOperation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AdminOperation)
//内容
	switch x := m.Content.(type) {
	case *AdminOperation_LogReq:
		s := proto.Size(x.LogReq)
n += 1 //标签和电线
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AdminOperation_LogSpecReq:
		s := proto.Size(x.LogSpecReq)
n += 1 //标签和电线
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*ServerStatus)(nil), "protos.ServerStatus")
	proto.RegisterType((*LogLevelRequest)(nil), "protos.LogLevelRequest")
	proto.RegisterType((*LogLevelResponse)(nil), "protos.LogLevelResponse")
	proto.RegisterType((*LogSpecRequest)(nil), "protos.LogSpecRequest")
	proto.RegisterType((*LogSpecResponse)(nil), "protos.LogSpecResponse")
	proto.RegisterType((*AdminOperation)(nil), "protos.AdminOperation")
	proto.RegisterEnum("protos.ServerStatus_StatusCode", ServerStatus_StatusCode_name, ServerStatus_StatusCode_value)
}

//引用导入以禁止错误（如果未使用）。
var _ context.Context
var _ grpc.ClientConn

//这是一个编译时断言，以确保生成的文件
//与正在编译的GRPC包兼容。
const _ = grpc.SupportPackageIsVersion4

//admin client是管理服务的客户端API。
//
//有关CTX使用和关闭/结束流式RPC的语义，请参阅https://godoc.org/google.golang.org/grpc clientconn.newstream。
type AdminClient interface {
	GetStatus(ctx context.Context, in *common.Envelope, opts ...grpc.CallOption) (*ServerStatus, error)
	StartServer(ctx context.Context, in *common.Envelope, opts ...grpc.CallOption) (*ServerStatus, error)
	GetModuleLogLevel(ctx context.Context, in *common.Envelope, opts ...grpc.CallOption) (*LogLevelResponse, error)
	SetModuleLogLevel(ctx context.Context, in *common.Envelope, opts ...grpc.CallOption) (*LogLevelResponse, error)
	RevertLogLevels(ctx context.Context, in *common.Envelope, opts ...grpc.CallOption) (*empty.Empty, error)
	GetLogSpec(ctx context.Context, in *common.Envelope, opts ...grpc.CallOption) (*LogSpecResponse, error)
	SetLogSpec(ctx context.Context, in *common.Envelope, opts ...grpc.CallOption) (*LogSpecResponse, error)
}

type adminClient struct {
	cc *grpc.ClientConn
}

func NewAdminClient(cc *grpc.ClientConn) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) GetStatus(ctx context.Context, in *common.Envelope, opts ...grpc.CallOption) (*ServerStatus, error) {
	out := new(ServerStatus)
	err := c.cc.Invoke(ctx, "/protos.Admin/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) StartServer(ctx context.Context, in *common.Envelope, opts ...grpc.CallOption) (*ServerStatus, error) {
	out := new(ServerStatus)
	err := c.cc.Invoke(ctx, "/protos.Admin/StartServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) GetModuleLogLevel(ctx context.Context, in *common.Envelope, opts ...grpc.CallOption) (*LogLevelResponse, error) {
	out := new(LogLevelResponse)
	err := c.cc.Invoke(ctx, "/protos.Admin/GetModuleLogLevel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) SetModuleLogLevel(ctx context.Context, in *common.Envelope, opts ...grpc.CallOption) (*LogLevelResponse, error) {
	out := new(LogLevelResponse)
	err := c.cc.Invoke(ctx, "/protos.Admin/SetModuleLogLevel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) RevertLogLevels(ctx context.Context, in *common.Envelope, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/protos.Admin/RevertLogLevels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) GetLogSpec(ctx context.Context, in *common.Envelope, opts ...grpc.CallOption) (*LogSpecResponse, error) {
	out := new(LogSpecResponse)
	err := c.cc.Invoke(ctx, "/protos.Admin/GetLogSpec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) SetLogSpec(ctx context.Context, in *common.Envelope, opts ...grpc.CallOption) (*LogSpecResponse, error) {
	out := new(LogSpecResponse)
	err := c.cc.Invoke(ctx, "/protos.Admin/SetLogSpec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

//admin server是管理服务的服务器API。
type AdminServer interface {
	GetStatus(context.Context, *common.Envelope) (*ServerStatus, error)
	StartServer(context.Context, *common.Envelope) (*ServerStatus, error)
	GetModuleLogLevel(context.Context, *common.Envelope) (*LogLevelResponse, error)
	SetModuleLogLevel(context.Context, *common.Envelope) (*LogLevelResponse, error)
	RevertLogLevels(context.Context, *common.Envelope) (*empty.Empty, error)
	GetLogSpec(context.Context, *common.Envelope) (*LogSpecResponse, error)
	SetLogSpec(context.Context, *common.Envelope) (*LogSpecResponse, error)
}

func RegisterAdminServer(s *grpc.Server, srv AdminServer) {
	s.RegisterService(&_Admin_serviceDesc, srv)
}

func _Admin_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Envelope)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Admin/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).GetStatus(ctx, req.(*common.Envelope))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_StartServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Envelope)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).StartServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Admin/StartServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).StartServer(ctx, req.(*common.Envelope))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_GetModuleLogLevel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Envelope)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).GetModuleLogLevel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Admin/GetModuleLogLevel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).GetModuleLogLevel(ctx, req.(*common.Envelope))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_SetModuleLogLevel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Envelope)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).SetModuleLogLevel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Admin/SetModuleLogLevel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).SetModuleLogLevel(ctx, req.(*common.Envelope))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_RevertLogLevels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Envelope)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).RevertLogLevels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Admin/RevertLogLevels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).RevertLogLevels(ctx, req.(*common.Envelope))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_GetLogSpec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Envelope)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).GetLogSpec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Admin/GetLogSpec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).GetLogSpec(ctx, req.(*common.Envelope))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_SetLogSpec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Envelope)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).SetLogSpec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Admin/SetLogSpec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).SetLogSpec(ctx, req.(*common.Envelope))
	}
	return interceptor(ctx, in, info, handler)
}

var _Admin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatus",
			Handler:    _Admin_GetStatus_Handler,
		},
		{
			MethodName: "StartServer",
			Handler:    _Admin_StartServer_Handler,
		},
		{
			MethodName: "GetModuleLogLevel",
			Handler:    _Admin_GetModuleLogLevel_Handler,
		},
		{
			MethodName: "SetModuleLogLevel",
			Handler:    _Admin_SetModuleLogLevel_Handler,
		},
		{
			MethodName: "RevertLogLevels",
			Handler:    _Admin_RevertLogLevels_Handler,
		},
		{
			MethodName: "GetLogSpec",
			Handler:    _Admin_GetLogSpec_Handler,
		},
		{
			MethodName: "SetLogSpec",
			Handler:    _Admin_SetLogSpec_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "peer/admin.proto",
}

func init() { proto.RegisterFile("peer/admin.proto", fileDescriptor_admin_b2904393863b6bc5) }

var fileDescriptor_admin_b2904393863b6bc5 = []byte{
//gzip文件描述符或协议的555字节
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x6e, 0x37, 0xda, 0x92, 0xd3, 0xd1, 0x05, 0x33, 0x6d, 0xa5, 0x13, 0x02, 0xe5, 0x0a, 0x84,
	0x94, 0x88, 0x22, 0x34, 0x76, 0xc1, 0x45, 0x4b, 0x43, 0x87, 0xe8, 0xd2, 0x2a, 0x59, 0x85, 0x40,
	0x42, 0x55, 0x9a, 0x9e, 0x79, 0x15, 0x6e, 0x9c, 0x39, 0x6e, 0xa5, 0xbd, 0x01, 0xcf, 0xc1, 0x23,
	0xf0, 0x84, 0x28, 0x76, 0xc2, 0x2a, 0x56, 0x2e, 0xd8, 0xae, 0xec, 0x73, 0xfc, 0x7d, 0xdf, 0xf9,
	0xb3, 0x0d, 0x66, 0x82, 0x28, 0x9c, 0x70, 0xb6, 0x98, 0xc7, 0x76, 0x22, 0xb8, 0xe4, 0xa4, 0xaa,
	0x96, 0xb4, 0x75, 0x48, 0x39, 0xa7, 0x0c, 0x1d, 0x65, 0x4e, 0x97, 0xe7, 0x0e, 0x2e, 0x12, 0x79,
	0xa5, 0x41, 0xad, 0x47, 0x11, 0x5f, 0x2c, 0x78, 0xec, 0xe8, 0x45, 0x3b, 0xad, 0x9f, 0x65, 0xd8,
	0x09, 0x50, 0xac, 0x50, 0x04, 0x32, 0x94, 0xcb, 0x94, 0x1c, 0x41, 0x35, 0x55, 0xbb, 0x66, 0xf9,
	0x59, 0xf9, 0x79, 0xa3, 0xfd, 0x54, 0x03, 0x53, 0x7b, 0x1d, 0x65, 0xeb, 0xe5, 0x3d, 0x9f, 0xa1,
	0x9f, 0xc3, 0xad, 0x2f, 0x00, 0xd7, 0x5e, 0xf2, 0x00, 0x8c, 0xb1, 0xd7, 0x73, 0x3f, 0x7c, 0xf4,
	0xdc, 0x9e, 0x59, 0x22, 0x75, 0xa8, 0x05, 0x67, 0x1d, 0xff, 0xcc, 0xed, 0x99, 0x65, 0x6d, 0x0c,
	0x47, 0x23, 0xb7, 0x67, 0x6e, 0x11, 0x80, 0xea, 0xa8, 0x33, 0x0e, 0xdc, 0x9e, 0xb9, 0x4d, 0x0c,
	0xa8, 0xb8, 0xbe, 0x3f, 0xf4, 0xcd, 0x7b, 0x19, 0x66, 0xec, 0x7d, 0xf2, 0x86, 0x9f, 0x3d, 0xb3,
	0x62, 0x9d, 0xc2, 0xee, 0x80, 0xd3, 0x01, 0xae, 0x90, 0xf9, 0x78, 0xb9, 0xc4, 0x54, 0x92, 0x27,
	0x00, 0x8c, 0xd3, 0xc9, 0x82, 0xcf, 0x96, 0x0c, 0x55, 0xaa, 0x86, 0x6f, 0x30, 0x4e, 0x4f, 0x95,
	0x83, 0x1c, 0x42, 0x66, 0x4c, 0x58, 0x46, 0x69, 0x6e, 0xa9, 0xd3, 0xfb, 0x2c, 0x97, 0xb0, 0x3c,
	0x30, 0xaf, 0xe5, 0xd2, 0x84, 0xc7, 0x29, 0xde, 0x49, 0xef, 0x25, 0x34, 0x06, 0x9c, 0x06, 0x09,
	0x46, 0x45, 0x76, 0x8f, 0x21, 0x3b, 0x9d, 0xa4, 0x09, 0x46, 0xb9, 0x56, 0x8d, 0x69, 0x84, 0xd5,
	0x55, 0xb5, 0x68, 0x70, 0x1e, 0xfb, 0xdf, 0x68, 0xb2, 0x07, 0x15, 0x14, 0x82, 0x8b, 0x3c, 0xa6,
	0x36, 0xac, 0x1f, 0x65, 0x68, 0x74, 0xb2, 0xf1, 0x0f, 0x13, 0x14, 0xa1, 0x9c, 0xf3, 0x98, 0xbc,
	0x82, 0x2a, 0xe3, 0xd4, 0xc7, 0x4b, 0xa5, 0x50, 0x6f, 0x1f, 0x14, 0x63, 0xfb, 0xab, 0x71, 0x27,
	0x25, 0x3f, 0x07, 0x92, 0xb7, 0xaa, 0xe4, 0x3c, 0x6d, 0x15, 0xa0, 0xde, 0xde, 0x5f, 0xa3, 0xad,
	0x15, 0x74, 0x52, 0xf2, 0xd7, 0xb0, 0x5d, 0x03, 0x6a, 0x11, 0x8f, 0x25, 0xc6, 0xb2, 0xfd, 0x6b,
	0x1b, 0x2a, 0x2a, 0x15, 0xf2, 0x06, 0x8c, 0x3e, 0xca, 0xfc, 0x16, 0x99, 0x76, 0x7e, 0xcb, 0xdc,
	0x78, 0x85, 0x8c, 0x27, 0xd8, 0xda, 0xdb, 0x74, 0x8f, 0xac, 0x12, 0x39, 0x82, 0x7a, 0x20, 0x43,
	0x21, 0xb5, 0xfb, 0x3f, 0x88, 0x1d, 0x78, 0xd8, 0x47, 0xa9, 0xe7, 0x53, 0x14, 0xb9, 0x81, 0xde,
	0xbc, 0xd9, 0x08, 0xdd, 0x76, 0x2d, 0x11, 0xdc, 0x51, 0xe2, 0x1d, 0xec, 0xfa, 0xb8, 0x42, 0x21,
	0x8b, 0xb3, 0x4d, 0xb5, 0xef, 0xdb, 0xfa, 0x5d, 0xda, 0xc5, 0xbb, 0xb4, 0xdd, 0xec, 0x5d, 0x5a,
	0x25, 0x72, 0x0c, 0xd0, 0x47, 0x99, 0x37, 0x7b, 0x03, 0xf3, 0xe0, 0xc6, 0x3c, 0xfe, 0x44, 0x3e,
	0x06, 0x08, 0x6e, 0x47, 0xed, 0x7e, 0x03, 0x8b, 0x0b, 0x6a, 0x5f, 0x5c, 0x25, 0x28, 0x18, 0xce,
	0x28, 0x0a, 0xfb, 0x3c, 0x9c, 0x8a, 0x79, 0x54, 0x50, 0xb2, 0x0f, 0xa6, 0xbb, 0xa3, 0xe6, 0x3a,
	0x0a, 0xa3, 0xef, 0x21, 0xc5, 0xaf, 0x2f, 0xe8, 0x5c, 0x5e, 0x2c, 0xa7, 0x59, 0x18, 0x67, 0x8d,
	0xe8, 0x68, 0xa2, 0xfe, 0x71, 0x52, 0x27, 0x23, 0x4e, 0xf5, 0x6f, 0xf4, 0xfa, 0x77, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xb0, 0x77, 0x16, 0x8f, 0xa8, 0x04, 0x00, 0x00,
}
