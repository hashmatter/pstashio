// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/p2p.proto

package pstashio

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

type RegisterPeerRequest struct {
	Ip                   string   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Location             string   `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	CachedBlocks         []*Block `protobuf:"bytes,3,rep,name=cached_blocks,json=cachedBlocks,proto3" json:"cached_blocks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterPeerRequest) Reset()         { *m = RegisterPeerRequest{} }
func (m *RegisterPeerRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterPeerRequest) ProtoMessage()    {}
func (*RegisterPeerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_faf6959bb89dc13d, []int{0}
}

func (m *RegisterPeerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterPeerRequest.Unmarshal(m, b)
}
func (m *RegisterPeerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterPeerRequest.Marshal(b, m, deterministic)
}
func (m *RegisterPeerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterPeerRequest.Merge(m, src)
}
func (m *RegisterPeerRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterPeerRequest.Size(m)
}
func (m *RegisterPeerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterPeerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterPeerRequest proto.InternalMessageInfo

func (m *RegisterPeerRequest) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *RegisterPeerRequest) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *RegisterPeerRequest) GetCachedBlocks() []*Block {
	if m != nil {
		return m.CachedBlocks
	}
	return nil
}

type RegisterPeerResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterPeerResponse) Reset()         { *m = RegisterPeerResponse{} }
func (m *RegisterPeerResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterPeerResponse) ProtoMessage()    {}
func (*RegisterPeerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_faf6959bb89dc13d, []int{1}
}

func (m *RegisterPeerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterPeerResponse.Unmarshal(m, b)
}
func (m *RegisterPeerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterPeerResponse.Marshal(b, m, deterministic)
}
func (m *RegisterPeerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterPeerResponse.Merge(m, src)
}
func (m *RegisterPeerResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterPeerResponse.Size(m)
}
func (m *RegisterPeerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterPeerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterPeerResponse proto.InternalMessageInfo

func (m *RegisterPeerResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type GetProvidersRequest struct {
	ResourceId           string   `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProvidersRequest) Reset()         { *m = GetProvidersRequest{} }
func (m *GetProvidersRequest) String() string { return proto.CompactTextString(m) }
func (*GetProvidersRequest) ProtoMessage()    {}
func (*GetProvidersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_faf6959bb89dc13d, []int{2}
}

func (m *GetProvidersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProvidersRequest.Unmarshal(m, b)
}
func (m *GetProvidersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProvidersRequest.Marshal(b, m, deterministic)
}
func (m *GetProvidersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProvidersRequest.Merge(m, src)
}
func (m *GetProvidersRequest) XXX_Size() int {
	return xxx_messageInfo_GetProvidersRequest.Size(m)
}
func (m *GetProvidersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProvidersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProvidersRequest proto.InternalMessageInfo

func (m *GetProvidersRequest) GetResourceId() string {
	if m != nil {
		return m.ResourceId
	}
	return ""
}

type GetProvidersResponse struct {
	Manifest             *ReplicationManifest `protobuf:"bytes,1,opt,name=manifest,proto3" json:"manifest,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetProvidersResponse) Reset()         { *m = GetProvidersResponse{} }
func (m *GetProvidersResponse) String() string { return proto.CompactTextString(m) }
func (*GetProvidersResponse) ProtoMessage()    {}
func (*GetProvidersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_faf6959bb89dc13d, []int{3}
}

func (m *GetProvidersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProvidersResponse.Unmarshal(m, b)
}
func (m *GetProvidersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProvidersResponse.Marshal(b, m, deterministic)
}
func (m *GetProvidersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProvidersResponse.Merge(m, src)
}
func (m *GetProvidersResponse) XXX_Size() int {
	return xxx_messageInfo_GetProvidersResponse.Size(m)
}
func (m *GetProvidersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProvidersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetProvidersResponse proto.InternalMessageInfo

func (m *GetProvidersResponse) GetManifest() *ReplicationManifest {
	if m != nil {
		return m.Manifest
	}
	return nil
}

// a resource represents data to be requested, replicated and exchanged by the
// peers in the network. a resource contains a set of all the blocks that are
// the data itself.
type Resource struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Size                 int32    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Blocks               []*Block `protobuf:"bytes,3,rep,name=blocks,proto3" json:"blocks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_faf6959bb89dc13d, []int{4}
}

func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Resource) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Resource) GetBlocks() []*Block {
	if m != nil {
		return m.Blocks
	}
	return nil
}

// a block is a "piece" of a resource. blocks in a given resource must have the
// same size (apart from the last one)
type Block struct {
	ResourceId           string   `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	Id                   int32    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Size                 int32    `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_faf6959bb89dc13d, []int{5}
}

func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (m *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(m, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetResourceId() string {
	if m != nil {
		return m.ResourceId
	}
	return ""
}

func (m *Block) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Block) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

// a peer block is a tuple containing information for a peer to fetch and
// validate a block from a resources stored in a given peer.
type PeerBlockRequest struct {
	ResourceId           string   `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	BlockId              string   `protobuf:"bytes,2,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	Addr                 string   `protobuf:"bytes,3,opt,name=addr,proto3" json:"addr,omitempty"`
	PeerPubkey           string   `protobuf:"bytes,4,opt,name=peer_pubkey,json=peerPubkey,proto3" json:"peer_pubkey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeerBlockRequest) Reset()         { *m = PeerBlockRequest{} }
func (m *PeerBlockRequest) String() string { return proto.CompactTextString(m) }
func (*PeerBlockRequest) ProtoMessage()    {}
func (*PeerBlockRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_faf6959bb89dc13d, []int{6}
}

func (m *PeerBlockRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerBlockRequest.Unmarshal(m, b)
}
func (m *PeerBlockRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerBlockRequest.Marshal(b, m, deterministic)
}
func (m *PeerBlockRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerBlockRequest.Merge(m, src)
}
func (m *PeerBlockRequest) XXX_Size() int {
	return xxx_messageInfo_PeerBlockRequest.Size(m)
}
func (m *PeerBlockRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerBlockRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PeerBlockRequest proto.InternalMessageInfo

func (m *PeerBlockRequest) GetResourceId() string {
	if m != nil {
		return m.ResourceId
	}
	return ""
}

func (m *PeerBlockRequest) GetBlockId() string {
	if m != nil {
		return m.BlockId
	}
	return ""
}

func (m *PeerBlockRequest) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *PeerBlockRequest) GetPeerPubkey() string {
	if m != nil {
		return m.PeerPubkey
	}
	return ""
}

// a replication manifest is a list of peer block requests which must be issued
type ReplicationManifest struct {
	BlockReq             []*PeerBlockRequest `protobuf:"bytes,1,rep,name=block_req,json=blockReq,proto3" json:"block_req,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ReplicationManifest) Reset()         { *m = ReplicationManifest{} }
func (m *ReplicationManifest) String() string { return proto.CompactTextString(m) }
func (*ReplicationManifest) ProtoMessage()    {}
func (*ReplicationManifest) Descriptor() ([]byte, []int) {
	return fileDescriptor_faf6959bb89dc13d, []int{7}
}

func (m *ReplicationManifest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplicationManifest.Unmarshal(m, b)
}
func (m *ReplicationManifest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplicationManifest.Marshal(b, m, deterministic)
}
func (m *ReplicationManifest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplicationManifest.Merge(m, src)
}
func (m *ReplicationManifest) XXX_Size() int {
	return xxx_messageInfo_ReplicationManifest.Size(m)
}
func (m *ReplicationManifest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplicationManifest.DiscardUnknown(m)
}

var xxx_messageInfo_ReplicationManifest proto.InternalMessageInfo

func (m *ReplicationManifest) GetBlockReq() []*PeerBlockRequest {
	if m != nil {
		return m.BlockReq
	}
	return nil
}

func init() {
	proto.RegisterType((*RegisterPeerRequest)(nil), "pstashio.RegisterPeerRequest")
	proto.RegisterType((*RegisterPeerResponse)(nil), "pstashio.RegisterPeerResponse")
	proto.RegisterType((*GetProvidersRequest)(nil), "pstashio.GetProvidersRequest")
	proto.RegisterType((*GetProvidersResponse)(nil), "pstashio.GetProvidersResponse")
	proto.RegisterType((*Resource)(nil), "pstashio.Resource")
	proto.RegisterType((*Block)(nil), "pstashio.Block")
	proto.RegisterType((*PeerBlockRequest)(nil), "pstashio.PeerBlockRequest")
	proto.RegisterType((*ReplicationManifest)(nil), "pstashio.ReplicationManifest")
}

func init() { proto.RegisterFile("pb/p2p.proto", fileDescriptor_faf6959bb89dc13d) }

var fileDescriptor_faf6959bb89dc13d = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x4f, 0x8f, 0xd3, 0x30,
	0x10, 0xc5, 0x37, 0xed, 0x6e, 0x49, 0xa6, 0xe5, 0x8f, 0xbc, 0x2b, 0x14, 0x2a, 0x01, 0x2b, 0x5f,
	0xd8, 0x53, 0x90, 0x02, 0x02, 0x71, 0xe5, 0x82, 0x2a, 0xf1, 0xa7, 0x98, 0x03, 0xc7, 0x28, 0x89,
	0x07, 0x6a, 0x35, 0xd4, 0xae, 0xed, 0x80, 0xe0, 0xca, 0xe7, 0xe1, 0x3b, 0xa2, 0xd8, 0x6e, 0x9b,
	0xd2, 0x02, 0x7b, 0xf3, 0x8c, 0xed, 0xf7, 0x1b, 0xbf, 0xbc, 0xc0, 0x44, 0x55, 0x8f, 0x55, 0xae,
	0x32, 0xa5, 0xa5, 0x95, 0x24, 0x56, 0xc6, 0x96, 0x66, 0x21, 0x24, 0xfd, 0x06, 0xe7, 0x0c, 0x3f,
	0x0b, 0x63, 0x51, 0xcf, 0x11, 0x35, 0xc3, 0x75, 0x8b, 0xc6, 0x92, 0x5b, 0x30, 0x10, 0x2a, 0x8d,
	0x2e, 0xa3, 0xab, 0x84, 0x0d, 0x84, 0x22, 0x53, 0x88, 0x1b, 0x59, 0x97, 0x56, 0xc8, 0x55, 0x3a,
	0x70, 0xdd, 0x6d, 0x4d, 0x9e, 0xc2, 0xcd, 0xba, 0xac, 0x17, 0xc8, 0x8b, 0xaa, 0x91, 0xf5, 0xd2,
	0xa4, 0xc3, 0xcb, 0xe1, 0xd5, 0x38, 0xbf, 0x9d, 0x6d, 0x20, 0xd9, 0xcb, 0xae, 0xcf, 0x26, 0xfe,
	0x94, 0x2b, 0x0c, 0xcd, 0xe0, 0x62, 0x1f, 0x6c, 0x94, 0x5c, 0x19, 0x24, 0x77, 0x61, 0x64, 0x6c,
	0x69, 0x5b, 0x13, 0xe8, 0xa1, 0xa2, 0xcf, 0xe0, 0xfc, 0x15, 0xda, 0xb9, 0x96, 0x5f, 0x05, 0x47,
	0x6d, 0x36, 0x83, 0x3e, 0x84, 0xb1, 0x46, 0x23, 0x5b, 0x5d, 0x63, 0x21, 0x78, 0xb8, 0x03, 0x9b,
	0xd6, 0x8c, 0xd3, 0xf7, 0x70, 0xb1, 0x7f, 0x2f, 0x70, 0x5e, 0x40, 0xfc, 0xa5, 0x5c, 0x89, 0x4f,
	0x68, 0xac, 0xbb, 0x35, 0xce, 0xef, 0xef, 0x06, 0x66, 0xa8, 0x1a, 0xe1, 0x9f, 0xf7, 0x26, 0x1c,
	0x62, 0xdb, 0xe3, 0xf4, 0x23, 0xc4, 0x2c, 0x00, 0x9c, 0x51, 0x7c, 0x6b, 0x14, 0x27, 0x04, 0x4e,
	0x8d, 0xf8, 0x81, 0xce, 0xa4, 0x33, 0xe6, 0xd6, 0xe4, 0x11, 0x8c, 0xfe, 0xed, 0x4c, 0xd8, 0xa6,
	0xaf, 0xe1, 0xcc, 0x35, 0xfe, 0xfb, 0xaa, 0x80, 0xf5, 0x90, 0x3e, 0x76, 0xb8, 0xc3, 0xd2, 0x9f,
	0x11, 0xdc, 0xe9, 0xac, 0xf5, 0x8c, 0x6b, 0xfa, 0x45, 0xee, 0x41, 0xec, 0xa6, 0x29, 0x82, 0x7e,
	0xc2, 0x6e, 0xb8, 0x7a, 0xe6, 0x20, 0x25, 0xe7, 0xda, 0x41, 0x12, 0xe6, 0xd6, 0x9d, 0x9e, 0x42,
	0xd4, 0x85, 0x6a, 0xab, 0x25, 0x7e, 0x4f, 0x4f, 0xbd, 0x5e, 0xd7, 0x9a, 0xbb, 0x0e, 0x7d, 0xdb,
	0x05, 0xec, 0xc0, 0x4d, 0xf2, 0x1c, 0x12, 0x8f, 0xd1, 0xb8, 0x4e, 0x23, 0x67, 0xcb, 0x74, 0x67,
	0xcb, 0x9f, 0x63, 0x33, 0x3f, 0x13, 0xc3, 0x75, 0xfe, 0x2b, 0x82, 0xe4, 0x43, 0x97, 0xa3, 0xb6,
	0x41, 0x4d, 0xde, 0xc1, 0xa4, 0x9f, 0x22, 0xb2, 0xf7, 0x0d, 0x0f, 0x62, 0x3d, 0x7d, 0xf0, 0xb7,
	0x6d, 0x1f, 0x0a, 0x7a, 0xd2, 0x09, 0xf6, 0xe3, 0xd2, 0x17, 0x3c, 0x12, 0xbf, 0xbe, 0xe0, 0xb1,
	0x94, 0xd1, 0x93, 0x6a, 0xe4, 0xfe, 0xb8, 0x27, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x83, 0x16,
	0x10, 0xae, 0x81, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SchedulerClient is the client API for Scheduler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SchedulerClient interface {
	RegisterPeer(ctx context.Context, in *RegisterPeerRequest, opts ...grpc.CallOption) (*RegisterPeerResponse, error)
	GetProviders(ctx context.Context, in *GetProvidersRequest, opts ...grpc.CallOption) (*GetProvidersResponse, error)
}

type schedulerClient struct {
	cc *grpc.ClientConn
}

func NewSchedulerClient(cc *grpc.ClientConn) SchedulerClient {
	return &schedulerClient{cc}
}

func (c *schedulerClient) RegisterPeer(ctx context.Context, in *RegisterPeerRequest, opts ...grpc.CallOption) (*RegisterPeerResponse, error) {
	out := new(RegisterPeerResponse)
	err := c.cc.Invoke(ctx, "/pstashio.Scheduler/RegisterPeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) GetProviders(ctx context.Context, in *GetProvidersRequest, opts ...grpc.CallOption) (*GetProvidersResponse, error) {
	out := new(GetProvidersResponse)
	err := c.cc.Invoke(ctx, "/pstashio.Scheduler/GetProviders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchedulerServer is the server API for Scheduler service.
type SchedulerServer interface {
	RegisterPeer(context.Context, *RegisterPeerRequest) (*RegisterPeerResponse, error)
	GetProviders(context.Context, *GetProvidersRequest) (*GetProvidersResponse, error)
}

// UnimplementedSchedulerServer can be embedded to have forward compatible implementations.
type UnimplementedSchedulerServer struct {
}

func (*UnimplementedSchedulerServer) RegisterPeer(ctx context.Context, req *RegisterPeerRequest) (*RegisterPeerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterPeer not implemented")
}
func (*UnimplementedSchedulerServer) GetProviders(ctx context.Context, req *GetProvidersRequest) (*GetProvidersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProviders not implemented")
}

func RegisterSchedulerServer(s *grpc.Server, srv SchedulerServer) {
	s.RegisterService(&_Scheduler_serviceDesc, srv)
}

func _Scheduler_RegisterPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterPeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).RegisterPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pstashio.Scheduler/RegisterPeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).RegisterPeer(ctx, req.(*RegisterPeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_GetProviders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProvidersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).GetProviders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pstashio.Scheduler/GetProviders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).GetProviders(ctx, req.(*GetProvidersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Scheduler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pstashio.Scheduler",
	HandlerType: (*SchedulerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterPeer",
			Handler:    _Scheduler_RegisterPeer_Handler,
		},
		{
			MethodName: "GetProviders",
			Handler:    _Scheduler_GetProviders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/p2p.proto",
}
