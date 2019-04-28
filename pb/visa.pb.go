// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/visa.proto

package visa

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

// Address card address
type Address struct {
	// @inject_tag: json:"city,omitempty"
	City string `protobuf:"bytes,1,opt,name=City,proto3" json:"City,omitempty"`
	// @inject_tag: json:"country,omitempty"
	Country string `protobuf:"bytes,2,opt,name=Country,proto3" json:"Country,omitempty"`
	// @inject_tag: json:"county,omitempty"
	County string `protobuf:"bytes,3,opt,name=County,proto3" json:"County,omitempty"`
	// @inject_tag: json:"state,omitempty"
	State string `protobuf:"bytes,4,opt,name=State,proto3" json:"State,omitempty"`
	// @inject_tag: json:"zipCode,omitempty"
	ZipCode              string   `protobuf:"bytes,5,opt,name=ZipCode,proto3" json:"ZipCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb0a7d9e34538282, []int{0}
}

func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Address) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Address) GetCounty() string {
	if m != nil {
		return m.County
	}
	return ""
}

func (m *Address) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Address) GetZipCode() string {
	if m != nil {
		return m.ZipCode
	}
	return ""
}

// CardAcceptor id
type CardAcceptor struct {
	// @inject_tag: json:"adress,omitempty"
	Address *Address `protobuf:"bytes,1,opt,name=Address,proto3" json:"Address,omitempty"`
	// @inject_tag: json:"idcode,omitempty"
	IDCode string `protobuf:"bytes,2,opt,name=IDCode,proto3" json:"IDCode,omitempty"`
	// @inject_tag: json:"name,omitempty"
	Name string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	// @inject_tag: json:"terminalId,omitempty"
	TerminalID           string   `protobuf:"bytes,4,opt,name=TerminalID,proto3" json:"TerminalID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CardAcceptor) Reset()         { *m = CardAcceptor{} }
func (m *CardAcceptor) String() string { return proto.CompactTextString(m) }
func (*CardAcceptor) ProtoMessage()    {}
func (*CardAcceptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb0a7d9e34538282, []int{1}
}

func (m *CardAcceptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CardAcceptor.Unmarshal(m, b)
}
func (m *CardAcceptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CardAcceptor.Marshal(b, m, deterministic)
}
func (m *CardAcceptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CardAcceptor.Merge(m, src)
}
func (m *CardAcceptor) XXX_Size() int {
	return xxx_messageInfo_CardAcceptor.Size(m)
}
func (m *CardAcceptor) XXX_DiscardUnknown() {
	xxx_messageInfo_CardAcceptor.DiscardUnknown(m)
}

var xxx_messageInfo_CardAcceptor proto.InternalMessageInfo

func (m *CardAcceptor) GetAddress() *Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *CardAcceptor) GetIDCode() string {
	if m != nil {
		return m.IDCode
	}
	return ""
}

func (m *CardAcceptor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CardAcceptor) GetTerminalID() string {
	if m != nil {
		return m.TerminalID
	}
	return ""
}

type VisaForexRequest struct {
	// @inject_tag: json:"cardAcceptor"
	CardAcceptor *CardAcceptor `protobuf:"bytes,1,opt,name=CardAcceptor,proto3" json:"CardAcceptor,omitempty"`
	// @inject_tag: json:"destinationCurrencyCode,omitempty"
	DestinationCurrencyCode string `protobuf:"bytes,2,opt,name=DestinationCurrencyCode,proto3" json:"DestinationCurrencyCode,omitempty"`
	// @inject_tag: json:"markUpRate,omitempty"
	MarkUpRate string `protobuf:"bytes,3,opt,name=MarkUpRate,proto3" json:"MarkUpRate,omitempty"`
	// @inject_tag: json:"retrievalReferenceNumber,omitempty"
	RetrievalReferenceNumber string `protobuf:"bytes,4,opt,name=RetrievalReferenceNumber,proto3" json:"RetrievalReferenceNumber,omitempty"`
	// @inject_tag: json:"sourceAmount,omitempty"
	SourceAmount string `protobuf:"bytes,5,opt,name=SourceAmount,proto3" json:"SourceAmount,omitempty"`
	// @inject_tag: json:"sourceCurrencyCode,omitempty"
	SourceCurrencyCode string `protobuf:"bytes,6,opt,name=SourceCurrencyCode,proto3" json:"SourceCurrencyCode,omitempty"`
	// @inject_tag: json:"systemsTraceAuditNumber,omitempty"
	SystemsTraceAuditNumber string   `protobuf:"bytes,7,opt,name=SystemsTraceAuditNumber,proto3" json:"SystemsTraceAuditNumber,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *VisaForexRequest) Reset()         { *m = VisaForexRequest{} }
func (m *VisaForexRequest) String() string { return proto.CompactTextString(m) }
func (*VisaForexRequest) ProtoMessage()    {}
func (*VisaForexRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb0a7d9e34538282, []int{2}
}

func (m *VisaForexRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VisaForexRequest.Unmarshal(m, b)
}
func (m *VisaForexRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VisaForexRequest.Marshal(b, m, deterministic)
}
func (m *VisaForexRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VisaForexRequest.Merge(m, src)
}
func (m *VisaForexRequest) XXX_Size() int {
	return xxx_messageInfo_VisaForexRequest.Size(m)
}
func (m *VisaForexRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VisaForexRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VisaForexRequest proto.InternalMessageInfo

func (m *VisaForexRequest) GetCardAcceptor() *CardAcceptor {
	if m != nil {
		return m.CardAcceptor
	}
	return nil
}

func (m *VisaForexRequest) GetDestinationCurrencyCode() string {
	if m != nil {
		return m.DestinationCurrencyCode
	}
	return ""
}

func (m *VisaForexRequest) GetMarkUpRate() string {
	if m != nil {
		return m.MarkUpRate
	}
	return ""
}

func (m *VisaForexRequest) GetRetrievalReferenceNumber() string {
	if m != nil {
		return m.RetrievalReferenceNumber
	}
	return ""
}

func (m *VisaForexRequest) GetSourceAmount() string {
	if m != nil {
		return m.SourceAmount
	}
	return ""
}

func (m *VisaForexRequest) GetSourceCurrencyCode() string {
	if m != nil {
		return m.SourceCurrencyCode
	}
	return ""
}

func (m *VisaForexRequest) GetSystemsTraceAuditNumber() string {
	if m != nil {
		return m.SystemsTraceAuditNumber
	}
	return ""
}

type VisaForexReply struct {
	// @inject_tag: json:"conversionRate,omitempty"
	ConversionRate float32 `protobuf:"fixed32,1,opt,name=ConversionRate,proto3" json:"ConversionRate,omitempty"`
	// @inject_tag: json:"destinationAmount,omitempty"
	DestinationAmount string `protobuf:"bytes,2,opt,name=DestinationAmount,proto3" json:"DestinationAmount,omitempty"`
	// @inject_tag: json:"markUpRateApplied,omitempty"
	MarkUpRateApplied string `protobuf:"bytes,3,opt,name=MarkUpRateApplied,proto3" json:"MarkUpRateApplied,omitempty"`
	// @inject_tag: json:"originalDestnAmtBeforeMarkUp,omitempty"
	OriginalDestnAmtBeforeMarkUp string   `protobuf:"bytes,4,opt,name=OriginalDestnAmtBeforeMarkUp,proto3" json:"OriginalDestnAmtBeforeMarkUp,omitempty"`
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	XXX_unrecognized             []byte   `json:"-"`
	XXX_sizecache                int32    `json:"-"`
}

func (m *VisaForexReply) Reset()         { *m = VisaForexReply{} }
func (m *VisaForexReply) String() string { return proto.CompactTextString(m) }
func (*VisaForexReply) ProtoMessage()    {}
func (*VisaForexReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb0a7d9e34538282, []int{3}
}

func (m *VisaForexReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VisaForexReply.Unmarshal(m, b)
}
func (m *VisaForexReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VisaForexReply.Marshal(b, m, deterministic)
}
func (m *VisaForexReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VisaForexReply.Merge(m, src)
}
func (m *VisaForexReply) XXX_Size() int {
	return xxx_messageInfo_VisaForexReply.Size(m)
}
func (m *VisaForexReply) XXX_DiscardUnknown() {
	xxx_messageInfo_VisaForexReply.DiscardUnknown(m)
}

var xxx_messageInfo_VisaForexReply proto.InternalMessageInfo

func (m *VisaForexReply) GetConversionRate() float32 {
	if m != nil {
		return m.ConversionRate
	}
	return 0
}

func (m *VisaForexReply) GetDestinationAmount() string {
	if m != nil {
		return m.DestinationAmount
	}
	return ""
}

func (m *VisaForexReply) GetMarkUpRateApplied() string {
	if m != nil {
		return m.MarkUpRateApplied
	}
	return ""
}

func (m *VisaForexReply) GetOriginalDestnAmtBeforeMarkUp() string {
	if m != nil {
		return m.OriginalDestnAmtBeforeMarkUp
	}
	return ""
}

func init() {
	proto.RegisterType((*Address)(nil), "visa.Address")
	proto.RegisterType((*CardAcceptor)(nil), "visa.CardAcceptor")
	proto.RegisterType((*VisaForexRequest)(nil), "visa.VisaForexRequest")
	proto.RegisterType((*VisaForexReply)(nil), "visa.VisaForexReply")
}

func init() { proto.RegisterFile("pb/visa.proto", fileDescriptor_cb0a7d9e34538282) }

var fileDescriptor_cb0a7d9e34538282 = []byte{
	// 464 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x4d, 0x8b, 0xd4, 0x40,
	0x10, 0x75, 0xc6, 0xd9, 0x19, 0xac, 0xfd, 0x40, 0x8b, 0x65, 0x0c, 0x22, 0x22, 0x39, 0xa8, 0x07,
	0x19, 0x61, 0x05, 0x11, 0x6f, 0xd9, 0x0c, 0xca, 0x1e, 0x5c, 0x21, 0x59, 0x3d, 0x78, 0xeb, 0x49,
	0x6a, 0xa5, 0x31, 0x49, 0xc7, 0xea, 0xce, 0x60, 0x4e, 0x1e, 0xfc, 0x93, 0xfe, 0x0a, 0x7f, 0x83,
	0xf4, 0x87, 0xbb, 0x99, 0x1d, 0xc7, 0x5b, 0xd5, 0x7b, 0x95, 0xe2, 0xbd, 0xd7, 0x15, 0x38, 0x6c,
	0x57, 0x2f, 0xd6, 0x52, 0x8b, 0x45, 0xcb, 0xca, 0x28, 0x9c, 0xd8, 0x3a, 0xfe, 0x01, 0xb3, 0xa4,
	0x2c, 0x99, 0xb4, 0x46, 0x84, 0x49, 0x2a, 0x4d, 0x1f, 0x8d, 0x1e, 0x8f, 0x9e, 0xdd, 0xc9, 0x5c,
	0x8d, 0x11, 0xcc, 0x52, 0xd5, 0x35, 0x86, 0xfb, 0x68, 0xec, 0xe0, 0xbf, 0x2d, 0xce, 0x61, 0xea,
	0xca, 0x3e, 0xba, 0xed, 0x88, 0xd0, 0xe1, 0x31, 0xec, 0xe5, 0x46, 0x18, 0x8a, 0x26, 0x0e, 0xf6,
	0x8d, 0xdd, 0xf3, 0x59, 0xb6, 0xa9, 0x2a, 0x29, 0xda, 0xf3, 0x7b, 0x42, 0x1b, 0xff, 0x1c, 0xc1,
	0x41, 0x2a, 0xb8, 0x4c, 0x8a, 0x82, 0x5a, 0xa3, 0x18, 0x9f, 0x5e, 0x29, 0x72, 0x4a, 0xf6, 0x4f,
	0x0e, 0x17, 0x4e, 0x75, 0x00, 0xb3, 0x2b, 0xbd, 0x73, 0x98, 0x9e, 0x2d, 0xdd, 0x4a, 0x2f, 0x2d,
	0x74, 0xd6, 0xc7, 0xb9, 0xa8, 0x29, 0xe8, 0x72, 0x35, 0x3e, 0x02, 0xb8, 0x20, 0xae, 0x65, 0x23,
	0xaa, 0xb3, 0x65, 0x90, 0x36, 0x40, 0xe2, 0xdf, 0x63, 0xb8, 0xfb, 0x49, 0x6a, 0xf1, 0x56, 0x31,
	0x7d, 0xcf, 0xe8, 0x5b, 0x47, 0xda, 0xe0, 0xab, 0x4d, 0x65, 0x41, 0x0e, 0x7a, 0x39, 0x43, 0x26,
	0xdb, 0x74, 0xf0, 0x1a, 0xee, 0x2f, 0x49, 0x1b, 0xd9, 0x08, 0x23, 0x55, 0x93, 0x76, 0xcc, 0xd4,
	0x14, 0xfd, 0x40, 0xe9, 0x2e, 0xda, 0xca, 0x7c, 0x2f, 0xf8, 0xeb, 0xc7, 0x36, 0xb3, 0x09, 0x7a,
	0x03, 0x03, 0x04, 0xdf, 0x40, 0x94, 0x91, 0x61, 0x49, 0x6b, 0x51, 0x65, 0x74, 0x49, 0xf6, 0x4b,
	0x3a, 0xef, 0xea, 0x15, 0x71, 0x30, 0xb5, 0x93, 0xc7, 0x18, 0x0e, 0x72, 0xd5, 0x71, 0x41, 0x49,
	0x6d, 0x5f, 0x2a, 0xbc, 0xc3, 0x06, 0x86, 0x0b, 0x40, 0xdf, 0x6f, 0x88, 0x9e, 0xba, 0xc9, 0x7f,
	0x30, 0xd6, 0x69, 0xde, 0x6b, 0x43, 0xb5, 0xbe, 0x60, 0x51, 0x50, 0xd2, 0x95, 0xd2, 0x04, 0x39,
	0x33, 0xef, 0x74, 0x07, 0x1d, 0xff, 0x1a, 0xc1, 0xd1, 0x20, 0xf0, 0xb6, 0xea, 0xf1, 0x09, 0x1c,
	0xa5, 0xaa, 0x59, 0x13, 0x6b, 0xa9, 0x1a, 0x17, 0x80, 0x0d, 0x7c, 0x9c, 0xdd, 0x40, 0xf1, 0x39,
	0xdc, 0x1b, 0xe4, 0x17, 0xdc, 0xf8, 0x60, 0xb7, 0x09, 0x3b, 0x7d, 0x1d, 0x60, 0xd2, 0xb6, 0x95,
	0xa4, 0x32, 0x24, 0xbb, 0x4d, 0xe0, 0x29, 0x3c, 0xfc, 0xc0, 0xf2, 0x8b, 0xbd, 0x0a, 0xbb, 0xaa,
	0x49, 0x6a, 0x73, 0x4a, 0x97, 0x8a, 0xc9, 0x8f, 0x86, 0x90, 0xff, 0x3b, 0x73, 0x92, 0xc3, 0xbe,
	0x75, 0x96, 0x13, 0xaf, 0x65, 0x41, 0xb8, 0x04, 0x7c, 0x47, 0xc6, 0xf9, 0xbc, 0x36, 0x82, 0x73,
	0x7f, 0x45, 0x37, 0x6f, 0xee, 0xc1, 0xf1, 0x16, 0xde, 0x56, 0x7d, 0x7c, 0x6b, 0x35, 0x75, 0x3f,
	0xed, 0xcb, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x59, 0x1b, 0xde, 0xc5, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VisaServiceClient is the client API for VisaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VisaServiceClient interface {
	GetForexConversion(ctx context.Context, in *VisaForexRequest, opts ...grpc.CallOption) (*VisaForexReply, error)
}

type visaServiceClient struct {
	cc *grpc.ClientConn
}

func NewVisaServiceClient(cc *grpc.ClientConn) VisaServiceClient {
	return &visaServiceClient{cc}
}

func (c *visaServiceClient) GetForexConversion(ctx context.Context, in *VisaForexRequest, opts ...grpc.CallOption) (*VisaForexReply, error) {
	out := new(VisaForexReply)
	err := c.cc.Invoke(ctx, "/visa.VisaService/GetForexConversion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VisaServiceServer is the server API for VisaService service.
type VisaServiceServer interface {
	GetForexConversion(context.Context, *VisaForexRequest) (*VisaForexReply, error)
}

func RegisterVisaServiceServer(s *grpc.Server, srv VisaServiceServer) {
	s.RegisterService(&_VisaService_serviceDesc, srv)
}

func _VisaService_GetForexConversion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VisaForexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VisaServiceServer).GetForexConversion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/visa.VisaService/GetForexConversion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VisaServiceServer).GetForexConversion(ctx, req.(*VisaForexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VisaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "visa.VisaService",
	HandlerType: (*VisaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetForexConversion",
			Handler:    _VisaService_GetForexConversion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/visa.proto",
}