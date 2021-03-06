// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/services/bidding_strategy_service.proto

package services // import "google.golang.org/genproto/googleapis/ads/googleads/v0/services"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import resources "google.golang.org/genproto/googleapis/ads/googleads/v0/resources"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import field_mask "google.golang.org/genproto/protobuf/field_mask"

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

// Request message for [BiddingStrategyService.GetBiddingStrategy][google.ads.googleads.v0.services.BiddingStrategyService.GetBiddingStrategy].
type GetBiddingStrategyRequest struct {
	// The resource name of the bidding strategy to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBiddingStrategyRequest) Reset()         { *m = GetBiddingStrategyRequest{} }
func (m *GetBiddingStrategyRequest) String() string { return proto.CompactTextString(m) }
func (*GetBiddingStrategyRequest) ProtoMessage()    {}
func (*GetBiddingStrategyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bidding_strategy_service_e353af7be94060a5, []int{0}
}
func (m *GetBiddingStrategyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBiddingStrategyRequest.Unmarshal(m, b)
}
func (m *GetBiddingStrategyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBiddingStrategyRequest.Marshal(b, m, deterministic)
}
func (dst *GetBiddingStrategyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBiddingStrategyRequest.Merge(dst, src)
}
func (m *GetBiddingStrategyRequest) XXX_Size() int {
	return xxx_messageInfo_GetBiddingStrategyRequest.Size(m)
}
func (m *GetBiddingStrategyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBiddingStrategyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBiddingStrategyRequest proto.InternalMessageInfo

func (m *GetBiddingStrategyRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [BiddingStrategyService.MutateBiddingStrategies][google.ads.googleads.v0.services.BiddingStrategyService.MutateBiddingStrategies].
type MutateBiddingStrategiesRequest struct {
	// The ID of the customer whose bidding strategies are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on individual bidding strategies.
	Operations           []*BiddingStrategyOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *MutateBiddingStrategiesRequest) Reset()         { *m = MutateBiddingStrategiesRequest{} }
func (m *MutateBiddingStrategiesRequest) String() string { return proto.CompactTextString(m) }
func (*MutateBiddingStrategiesRequest) ProtoMessage()    {}
func (*MutateBiddingStrategiesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bidding_strategy_service_e353af7be94060a5, []int{1}
}
func (m *MutateBiddingStrategiesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateBiddingStrategiesRequest.Unmarshal(m, b)
}
func (m *MutateBiddingStrategiesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateBiddingStrategiesRequest.Marshal(b, m, deterministic)
}
func (dst *MutateBiddingStrategiesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateBiddingStrategiesRequest.Merge(dst, src)
}
func (m *MutateBiddingStrategiesRequest) XXX_Size() int {
	return xxx_messageInfo_MutateBiddingStrategiesRequest.Size(m)
}
func (m *MutateBiddingStrategiesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateBiddingStrategiesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateBiddingStrategiesRequest proto.InternalMessageInfo

func (m *MutateBiddingStrategiesRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateBiddingStrategiesRequest) GetOperations() []*BiddingStrategyOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

// A single operation (create, update, remove) on a bidding strategy.
type BiddingStrategyOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*BiddingStrategyOperation_Create
	//	*BiddingStrategyOperation_Update
	//	*BiddingStrategyOperation_Remove
	Operation            isBiddingStrategyOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *BiddingStrategyOperation) Reset()         { *m = BiddingStrategyOperation{} }
func (m *BiddingStrategyOperation) String() string { return proto.CompactTextString(m) }
func (*BiddingStrategyOperation) ProtoMessage()    {}
func (*BiddingStrategyOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_bidding_strategy_service_e353af7be94060a5, []int{2}
}
func (m *BiddingStrategyOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BiddingStrategyOperation.Unmarshal(m, b)
}
func (m *BiddingStrategyOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BiddingStrategyOperation.Marshal(b, m, deterministic)
}
func (dst *BiddingStrategyOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BiddingStrategyOperation.Merge(dst, src)
}
func (m *BiddingStrategyOperation) XXX_Size() int {
	return xxx_messageInfo_BiddingStrategyOperation.Size(m)
}
func (m *BiddingStrategyOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_BiddingStrategyOperation.DiscardUnknown(m)
}

var xxx_messageInfo_BiddingStrategyOperation proto.InternalMessageInfo

func (m *BiddingStrategyOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isBiddingStrategyOperation_Operation interface {
	isBiddingStrategyOperation_Operation()
}

type BiddingStrategyOperation_Create struct {
	Create *resources.BiddingStrategy `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type BiddingStrategyOperation_Update struct {
	Update *resources.BiddingStrategy `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type BiddingStrategyOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*BiddingStrategyOperation_Create) isBiddingStrategyOperation_Operation() {}

func (*BiddingStrategyOperation_Update) isBiddingStrategyOperation_Operation() {}

func (*BiddingStrategyOperation_Remove) isBiddingStrategyOperation_Operation() {}

func (m *BiddingStrategyOperation) GetOperation() isBiddingStrategyOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *BiddingStrategyOperation) GetCreate() *resources.BiddingStrategy {
	if x, ok := m.GetOperation().(*BiddingStrategyOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *BiddingStrategyOperation) GetUpdate() *resources.BiddingStrategy {
	if x, ok := m.GetOperation().(*BiddingStrategyOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *BiddingStrategyOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*BiddingStrategyOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*BiddingStrategyOperation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _BiddingStrategyOperation_OneofMarshaler, _BiddingStrategyOperation_OneofUnmarshaler, _BiddingStrategyOperation_OneofSizer, []interface{}{
		(*BiddingStrategyOperation_Create)(nil),
		(*BiddingStrategyOperation_Update)(nil),
		(*BiddingStrategyOperation_Remove)(nil),
	}
}

func _BiddingStrategyOperation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*BiddingStrategyOperation)
	// operation
	switch x := m.Operation.(type) {
	case *BiddingStrategyOperation_Create:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Create); err != nil {
			return err
		}
	case *BiddingStrategyOperation_Update:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Update); err != nil {
			return err
		}
	case *BiddingStrategyOperation_Remove:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Remove)
	case nil:
	default:
		return fmt.Errorf("BiddingStrategyOperation.Operation has unexpected type %T", x)
	}
	return nil
}

func _BiddingStrategyOperation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*BiddingStrategyOperation)
	switch tag {
	case 1: // operation.create
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(resources.BiddingStrategy)
		err := b.DecodeMessage(msg)
		m.Operation = &BiddingStrategyOperation_Create{msg}
		return true, err
	case 2: // operation.update
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(resources.BiddingStrategy)
		err := b.DecodeMessage(msg)
		m.Operation = &BiddingStrategyOperation_Update{msg}
		return true, err
	case 3: // operation.remove
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Operation = &BiddingStrategyOperation_Remove{x}
		return true, err
	default:
		return false, nil
	}
}

func _BiddingStrategyOperation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*BiddingStrategyOperation)
	// operation
	switch x := m.Operation.(type) {
	case *BiddingStrategyOperation_Create:
		s := proto.Size(x.Create)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *BiddingStrategyOperation_Update:
		s := proto.Size(x.Update)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *BiddingStrategyOperation_Remove:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Remove)))
		n += len(x.Remove)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Response message for bidding strategy mutate.
type MutateBiddingStrategiesResponse struct {
	// All results for the mutate.
	Results              []*MutateBiddingStrategyResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *MutateBiddingStrategiesResponse) Reset()         { *m = MutateBiddingStrategiesResponse{} }
func (m *MutateBiddingStrategiesResponse) String() string { return proto.CompactTextString(m) }
func (*MutateBiddingStrategiesResponse) ProtoMessage()    {}
func (*MutateBiddingStrategiesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bidding_strategy_service_e353af7be94060a5, []int{3}
}
func (m *MutateBiddingStrategiesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateBiddingStrategiesResponse.Unmarshal(m, b)
}
func (m *MutateBiddingStrategiesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateBiddingStrategiesResponse.Marshal(b, m, deterministic)
}
func (dst *MutateBiddingStrategiesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateBiddingStrategiesResponse.Merge(dst, src)
}
func (m *MutateBiddingStrategiesResponse) XXX_Size() int {
	return xxx_messageInfo_MutateBiddingStrategiesResponse.Size(m)
}
func (m *MutateBiddingStrategiesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateBiddingStrategiesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateBiddingStrategiesResponse proto.InternalMessageInfo

func (m *MutateBiddingStrategiesResponse) GetResults() []*MutateBiddingStrategyResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the bidding strategy mutate.
type MutateBiddingStrategyResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateBiddingStrategyResult) Reset()         { *m = MutateBiddingStrategyResult{} }
func (m *MutateBiddingStrategyResult) String() string { return proto.CompactTextString(m) }
func (*MutateBiddingStrategyResult) ProtoMessage()    {}
func (*MutateBiddingStrategyResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_bidding_strategy_service_e353af7be94060a5, []int{4}
}
func (m *MutateBiddingStrategyResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateBiddingStrategyResult.Unmarshal(m, b)
}
func (m *MutateBiddingStrategyResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateBiddingStrategyResult.Marshal(b, m, deterministic)
}
func (dst *MutateBiddingStrategyResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateBiddingStrategyResult.Merge(dst, src)
}
func (m *MutateBiddingStrategyResult) XXX_Size() int {
	return xxx_messageInfo_MutateBiddingStrategyResult.Size(m)
}
func (m *MutateBiddingStrategyResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateBiddingStrategyResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateBiddingStrategyResult proto.InternalMessageInfo

func (m *MutateBiddingStrategyResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetBiddingStrategyRequest)(nil), "google.ads.googleads.v0.services.GetBiddingStrategyRequest")
	proto.RegisterType((*MutateBiddingStrategiesRequest)(nil), "google.ads.googleads.v0.services.MutateBiddingStrategiesRequest")
	proto.RegisterType((*BiddingStrategyOperation)(nil), "google.ads.googleads.v0.services.BiddingStrategyOperation")
	proto.RegisterType((*MutateBiddingStrategiesResponse)(nil), "google.ads.googleads.v0.services.MutateBiddingStrategiesResponse")
	proto.RegisterType((*MutateBiddingStrategyResult)(nil), "google.ads.googleads.v0.services.MutateBiddingStrategyResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BiddingStrategyServiceClient is the client API for BiddingStrategyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BiddingStrategyServiceClient interface {
	// Returns the requested bidding strategy in full detail.
	GetBiddingStrategy(ctx context.Context, in *GetBiddingStrategyRequest, opts ...grpc.CallOption) (*resources.BiddingStrategy, error)
	// Creates, updates, or removes bidding strategies. Operation statuses are
	// returned.
	MutateBiddingStrategies(ctx context.Context, in *MutateBiddingStrategiesRequest, opts ...grpc.CallOption) (*MutateBiddingStrategiesResponse, error)
}

type biddingStrategyServiceClient struct {
	cc *grpc.ClientConn
}

func NewBiddingStrategyServiceClient(cc *grpc.ClientConn) BiddingStrategyServiceClient {
	return &biddingStrategyServiceClient{cc}
}

func (c *biddingStrategyServiceClient) GetBiddingStrategy(ctx context.Context, in *GetBiddingStrategyRequest, opts ...grpc.CallOption) (*resources.BiddingStrategy, error) {
	out := new(resources.BiddingStrategy)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.BiddingStrategyService/GetBiddingStrategy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *biddingStrategyServiceClient) MutateBiddingStrategies(ctx context.Context, in *MutateBiddingStrategiesRequest, opts ...grpc.CallOption) (*MutateBiddingStrategiesResponse, error) {
	out := new(MutateBiddingStrategiesResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.BiddingStrategyService/MutateBiddingStrategies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BiddingStrategyServiceServer is the server API for BiddingStrategyService service.
type BiddingStrategyServiceServer interface {
	// Returns the requested bidding strategy in full detail.
	GetBiddingStrategy(context.Context, *GetBiddingStrategyRequest) (*resources.BiddingStrategy, error)
	// Creates, updates, or removes bidding strategies. Operation statuses are
	// returned.
	MutateBiddingStrategies(context.Context, *MutateBiddingStrategiesRequest) (*MutateBiddingStrategiesResponse, error)
}

func RegisterBiddingStrategyServiceServer(s *grpc.Server, srv BiddingStrategyServiceServer) {
	s.RegisterService(&_BiddingStrategyService_serviceDesc, srv)
}

func _BiddingStrategyService_GetBiddingStrategy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBiddingStrategyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BiddingStrategyServiceServer).GetBiddingStrategy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.BiddingStrategyService/GetBiddingStrategy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BiddingStrategyServiceServer).GetBiddingStrategy(ctx, req.(*GetBiddingStrategyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BiddingStrategyService_MutateBiddingStrategies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateBiddingStrategiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BiddingStrategyServiceServer).MutateBiddingStrategies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.BiddingStrategyService/MutateBiddingStrategies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BiddingStrategyServiceServer).MutateBiddingStrategies(ctx, req.(*MutateBiddingStrategiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BiddingStrategyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v0.services.BiddingStrategyService",
	HandlerType: (*BiddingStrategyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBiddingStrategy",
			Handler:    _BiddingStrategyService_GetBiddingStrategy_Handler,
		},
		{
			MethodName: "MutateBiddingStrategies",
			Handler:    _BiddingStrategyService_MutateBiddingStrategies_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v0/services/bidding_strategy_service.proto",
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/services/bidding_strategy_service.proto", fileDescriptor_bidding_strategy_service_e353af7be94060a5)
}

var fileDescriptor_bidding_strategy_service_e353af7be94060a5 = []byte{
	// 602 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0xb1, 0x83, 0x8a, 0x7a, 0x86, 0xe5, 0x06, 0x30, 0x29, 0xa2, 0x91, 0x61, 0x88, 0x32,
	0xd8, 0x51, 0x2a, 0x10, 0x4a, 0x14, 0x51, 0x67, 0x20, 0x45, 0xa2, 0x50, 0xb9, 0x52, 0x91, 0xaa,
	0x48, 0xd1, 0x25, 0x7e, 0xb5, 0xac, 0xc6, 0x3e, 0x73, 0x77, 0x8e, 0x54, 0xaa, 0x2e, 0x7c, 0x05,
	0x36, 0x24, 0x16, 0x46, 0xbe, 0x05, 0x8c, 0xac, 0x8c, 0x4c, 0x48, 0x7c, 0x10, 0x64, 0x9f, 0x2f,
	0x69, 0xd3, 0x98, 0xa0, 0x76, 0x7b, 0xb6, 0xdf, 0xfd, 0xde, 0xff, 0xbd, 0xff, 0x3d, 0xa3, 0xe7,
	0x01, 0xa5, 0xc1, 0x04, 0x1c, 0xe2, 0x73, 0x47, 0x86, 0x59, 0x34, 0x6d, 0x3a, 0x1c, 0xd8, 0x34,
	0x1c, 0x03, 0x77, 0x46, 0xa1, 0xef, 0x87, 0x71, 0x30, 0xe4, 0x82, 0x11, 0x01, 0xc1, 0xc9, 0xb0,
	0xf8, 0x62, 0x27, 0x8c, 0x0a, 0x8a, 0x6b, 0xf2, 0x94, 0x4d, 0x7c, 0x6e, 0xcf, 0x00, 0xf6, 0xb4,
	0x69, 0x2b, 0x40, 0xf5, 0x59, 0x59, 0x09, 0x06, 0x9c, 0xa6, 0x6c, 0x59, 0x0d, 0xc9, 0xae, 0x3e,
	0x50, 0x27, 0x93, 0xd0, 0x21, 0x71, 0x4c, 0x05, 0x11, 0x21, 0x8d, 0x79, 0xf1, 0xb5, 0xa8, 0xec,
	0xe4, 0x4f, 0xa3, 0xf4, 0xc8, 0x39, 0x0a, 0x61, 0xe2, 0x0f, 0x23, 0xc2, 0x8f, 0x65, 0x86, 0xb5,
	0x8d, 0xee, 0xf7, 0x41, 0xf4, 0x24, 0x7c, 0xbf, 0x60, 0x7b, 0xf0, 0x2e, 0x05, 0x2e, 0xf0, 0x23,
	0x74, 0x47, 0x09, 0x18, 0xc6, 0x24, 0x02, 0x53, 0xab, 0x69, 0xf5, 0x75, 0xef, 0xb6, 0x7a, 0xf9,
	0x9a, 0x44, 0x60, 0x7d, 0xd6, 0xd0, 0xc3, 0xdd, 0x54, 0x10, 0x01, 0x17, 0x29, 0x21, 0x70, 0xc5,
	0xd9, 0x44, 0xc6, 0x38, 0xe5, 0x82, 0x46, 0xc0, 0x86, 0xa1, 0x5f, 0x50, 0x90, 0x7a, 0xf5, 0xd2,
	0xc7, 0x87, 0x08, 0xd1, 0x04, 0x98, 0xd4, 0x6e, 0xea, 0xb5, 0x4a, 0xdd, 0x68, 0xb5, 0xed, 0x55,
	0x63, 0xb3, 0x17, 0x64, 0xbf, 0x51, 0x08, 0xef, 0x1c, 0xcd, 0xfa, 0xa4, 0x23, 0xb3, 0x2c, 0x11,
	0x77, 0x90, 0x91, 0x26, 0x3e, 0x11, 0x90, 0xcf, 0xc4, 0xbc, 0x59, 0xd3, 0xea, 0x46, 0xab, 0xaa,
	0x2a, 0xab, 0xb1, 0xd9, 0x2f, 0xb2, 0xb1, 0xed, 0x12, 0x7e, 0xec, 0x21, 0x99, 0x9e, 0xc5, 0xf8,
	0x15, 0x5a, 0x1b, 0x33, 0x20, 0x42, 0xce, 0xc5, 0x68, 0xb5, 0x4a, 0x15, 0xcf, 0x6c, 0x5c, 0x94,
	0xbc, 0x73, 0xc3, 0x2b, 0x18, 0x19, 0x4d, 0xb2, 0x4d, 0xfd, 0x3a, 0x34, 0xc9, 0xc0, 0x26, 0x5a,
	0x63, 0x10, 0xd1, 0x29, 0x98, 0x95, 0x6c, 0xda, 0xd9, 0x17, 0xf9, 0xdc, 0x33, 0xd0, 0xfa, 0x6c,
	0x3a, 0xd6, 0x7b, 0xb4, 0x59, 0xea, 0x1d, 0x4f, 0x68, 0xcc, 0x01, 0xbf, 0x45, 0xb7, 0x18, 0xf0,
	0x74, 0x22, 0x94, 0x31, 0xdd, 0xd5, 0xc6, 0x2c, 0x63, 0x9e, 0x78, 0x39, 0xc5, 0x53, 0x34, 0xab,
	0x87, 0x36, 0xfe, 0x91, 0xf7, 0x5f, 0x97, 0xaf, 0xf5, 0xad, 0x82, 0xee, 0x2e, 0x1c, 0xdf, 0x97,
	0x22, 0xf0, 0x77, 0x0d, 0xe1, 0xcb, 0x57, 0x1b, 0x77, 0x56, 0xab, 0x2f, 0x5d, 0x88, 0xea, 0x15,
	0x3c, 0xb1, 0x3a, 0x1f, 0x7e, 0xfe, 0xf9, 0xa8, 0x3f, 0xc1, 0x5b, 0xd9, 0x3e, 0x9f, 0x5e, 0x68,
	0xa9, 0xab, 0x56, 0x80, 0x3b, 0x0d, 0xb5, 0xe0, 0x73, 0x07, 0x9c, 0xc6, 0x19, 0xfe, 0xad, 0xa1,
	0x7b, 0x25, 0x06, 0xe1, 0xed, 0xab, 0xf9, 0x30, 0xdf, 0xcb, 0xaa, 0x7b, 0x0d, 0x82, 0xbc, 0x1d,
	0x96, 0x9b, 0x77, 0xd7, 0xb1, 0x9e, 0x66, 0xdd, 0xcd, 0xdb, 0x39, 0x3d, 0xb7, 0xef, 0xdd, 0xc6,
	0xd9, 0xe5, 0xe6, 0xda, 0x51, 0x0e, 0x6e, 0x6b, 0x8d, 0xde, 0x2f, 0x0d, 0x3d, 0x1e, 0xd3, 0x68,
	0xa5, 0x96, 0xde, 0xc6, 0x72, 0xa7, 0xf7, 0xb2, 0x2d, 0xdd, 0xd3, 0x0e, 0x77, 0x0a, 0x40, 0x40,
	0x27, 0x24, 0x0e, 0x6c, 0xca, 0x02, 0x27, 0x80, 0x38, 0xdf, 0x61, 0xf5, 0x53, 0x4d, 0x42, 0x5e,
	0xfe, 0x1b, 0xef, 0xa8, 0xe0, 0x8b, 0x5e, 0xe9, 0xbb, 0xee, 0x57, 0xbd, 0xd6, 0x97, 0x40, 0xd7,
	0xe7, 0xb6, 0x0c, 0xb3, 0xe8, 0xa0, 0x69, 0x17, 0x85, 0xf9, 0x0f, 0x95, 0x32, 0x70, 0x7d, 0x3e,
	0x98, 0xa5, 0x0c, 0x0e, 0x9a, 0x03, 0x95, 0x32, 0x5a, 0xcb, 0x05, 0x6c, 0xfd, 0x0d, 0x00, 0x00,
	0xff, 0xff, 0x09, 0xdb, 0xf0, 0x1a, 0x46, 0x06, 0x00, 0x00,
}
