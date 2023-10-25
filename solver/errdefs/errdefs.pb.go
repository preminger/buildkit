// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: errdefs.proto

package errdefs

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	pb "github.com/preminger/buildkit/solver/pb"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Vertex struct {
	Digest               string   `protobuf:"bytes,1,opt,name=digest,proto3" json:"digest,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vertex) Reset()         { *m = Vertex{} }
func (m *Vertex) String() string { return proto.CompactTextString(m) }
func (*Vertex) ProtoMessage()    {}
func (*Vertex) Descriptor() ([]byte, []int) {
	return fileDescriptor_689dc58a5060aff5, []int{0}
}
func (m *Vertex) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vertex.Unmarshal(m, b)
}
func (m *Vertex) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vertex.Marshal(b, m, deterministic)
}
func (m *Vertex) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vertex.Merge(m, src)
}
func (m *Vertex) XXX_Size() int {
	return xxx_messageInfo_Vertex.Size(m)
}
func (m *Vertex) XXX_DiscardUnknown() {
	xxx_messageInfo_Vertex.DiscardUnknown(m)
}

var xxx_messageInfo_Vertex proto.InternalMessageInfo

func (m *Vertex) GetDigest() string {
	if m != nil {
		return m.Digest
	}
	return ""
}

type Source struct {
	Info                 *pb.SourceInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Ranges               []*pb.Range    `protobuf:"bytes,2,rep,name=ranges,proto3" json:"ranges,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Source) Reset()         { *m = Source{} }
func (m *Source) String() string { return proto.CompactTextString(m) }
func (*Source) ProtoMessage()    {}
func (*Source) Descriptor() ([]byte, []int) {
	return fileDescriptor_689dc58a5060aff5, []int{1}
}
func (m *Source) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Source.Unmarshal(m, b)
}
func (m *Source) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Source.Marshal(b, m, deterministic)
}
func (m *Source) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Source.Merge(m, src)
}
func (m *Source) XXX_Size() int {
	return xxx_messageInfo_Source.Size(m)
}
func (m *Source) XXX_DiscardUnknown() {
	xxx_messageInfo_Source.DiscardUnknown(m)
}

var xxx_messageInfo_Source proto.InternalMessageInfo

func (m *Source) GetInfo() *pb.SourceInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *Source) GetRanges() []*pb.Range {
	if m != nil {
		return m.Ranges
	}
	return nil
}

type FrontendCap struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FrontendCap) Reset()         { *m = FrontendCap{} }
func (m *FrontendCap) String() string { return proto.CompactTextString(m) }
func (*FrontendCap) ProtoMessage()    {}
func (*FrontendCap) Descriptor() ([]byte, []int) {
	return fileDescriptor_689dc58a5060aff5, []int{2}
}
func (m *FrontendCap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrontendCap.Unmarshal(m, b)
}
func (m *FrontendCap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrontendCap.Marshal(b, m, deterministic)
}
func (m *FrontendCap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrontendCap.Merge(m, src)
}
func (m *FrontendCap) XXX_Size() int {
	return xxx_messageInfo_FrontendCap.Size(m)
}
func (m *FrontendCap) XXX_DiscardUnknown() {
	xxx_messageInfo_FrontendCap.DiscardUnknown(m)
}

var xxx_messageInfo_FrontendCap proto.InternalMessageInfo

func (m *FrontendCap) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Subrequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Subrequest) Reset()         { *m = Subrequest{} }
func (m *Subrequest) String() string { return proto.CompactTextString(m) }
func (*Subrequest) ProtoMessage()    {}
func (*Subrequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_689dc58a5060aff5, []int{3}
}
func (m *Subrequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Subrequest.Unmarshal(m, b)
}
func (m *Subrequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Subrequest.Marshal(b, m, deterministic)
}
func (m *Subrequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subrequest.Merge(m, src)
}
func (m *Subrequest) XXX_Size() int {
	return xxx_messageInfo_Subrequest.Size(m)
}
func (m *Subrequest) XXX_DiscardUnknown() {
	xxx_messageInfo_Subrequest.DiscardUnknown(m)
}

var xxx_messageInfo_Subrequest proto.InternalMessageInfo

func (m *Subrequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Solve struct {
	InputIDs []string `protobuf:"bytes,1,rep,name=inputIDs,proto3" json:"inputIDs,omitempty"`
	MountIDs []string `protobuf:"bytes,2,rep,name=mountIDs,proto3" json:"mountIDs,omitempty"`
	Op       *pb.Op   `protobuf:"bytes,3,opt,name=op,proto3" json:"op,omitempty"`
	// Types that are valid to be assigned to Subject:
	//
	//	*Solve_File
	//	*Solve_Cache
	Subject              isSolve_Subject `protobuf_oneof:"subject"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Solve) Reset()         { *m = Solve{} }
func (m *Solve) String() string { return proto.CompactTextString(m) }
func (*Solve) ProtoMessage()    {}
func (*Solve) Descriptor() ([]byte, []int) {
	return fileDescriptor_689dc58a5060aff5, []int{4}
}
func (m *Solve) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Solve.Unmarshal(m, b)
}
func (m *Solve) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Solve.Marshal(b, m, deterministic)
}
func (m *Solve) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Solve.Merge(m, src)
}
func (m *Solve) XXX_Size() int {
	return xxx_messageInfo_Solve.Size(m)
}
func (m *Solve) XXX_DiscardUnknown() {
	xxx_messageInfo_Solve.DiscardUnknown(m)
}

var xxx_messageInfo_Solve proto.InternalMessageInfo

type isSolve_Subject interface {
	isSolve_Subject()
}

type Solve_File struct {
	File *FileAction `protobuf:"bytes,4,opt,name=file,proto3,oneof" json:"file,omitempty"`
}
type Solve_Cache struct {
	Cache *ContentCache `protobuf:"bytes,5,opt,name=cache,proto3,oneof" json:"cache,omitempty"`
}

func (*Solve_File) isSolve_Subject()  {}
func (*Solve_Cache) isSolve_Subject() {}

func (m *Solve) GetSubject() isSolve_Subject {
	if m != nil {
		return m.Subject
	}
	return nil
}

func (m *Solve) GetInputIDs() []string {
	if m != nil {
		return m.InputIDs
	}
	return nil
}

func (m *Solve) GetMountIDs() []string {
	if m != nil {
		return m.MountIDs
	}
	return nil
}

func (m *Solve) GetOp() *pb.Op {
	if m != nil {
		return m.Op
	}
	return nil
}

func (m *Solve) GetFile() *FileAction {
	if x, ok := m.GetSubject().(*Solve_File); ok {
		return x.File
	}
	return nil
}

func (m *Solve) GetCache() *ContentCache {
	if x, ok := m.GetSubject().(*Solve_Cache); ok {
		return x.Cache
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Solve) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Solve_File)(nil),
		(*Solve_Cache)(nil),
	}
}

type FileAction struct {
	// Index of the file action that failed the exec.
	Index                int64    `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileAction) Reset()         { *m = FileAction{} }
func (m *FileAction) String() string { return proto.CompactTextString(m) }
func (*FileAction) ProtoMessage()    {}
func (*FileAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_689dc58a5060aff5, []int{5}
}
func (m *FileAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileAction.Unmarshal(m, b)
}
func (m *FileAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileAction.Marshal(b, m, deterministic)
}
func (m *FileAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileAction.Merge(m, src)
}
func (m *FileAction) XXX_Size() int {
	return xxx_messageInfo_FileAction.Size(m)
}
func (m *FileAction) XXX_DiscardUnknown() {
	xxx_messageInfo_FileAction.DiscardUnknown(m)
}

var xxx_messageInfo_FileAction proto.InternalMessageInfo

func (m *FileAction) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

type ContentCache struct {
	// Original index of result that failed the slow cache calculation.
	Index                int64    `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContentCache) Reset()         { *m = ContentCache{} }
func (m *ContentCache) String() string { return proto.CompactTextString(m) }
func (*ContentCache) ProtoMessage()    {}
func (*ContentCache) Descriptor() ([]byte, []int) {
	return fileDescriptor_689dc58a5060aff5, []int{6}
}
func (m *ContentCache) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentCache.Unmarshal(m, b)
}
func (m *ContentCache) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentCache.Marshal(b, m, deterministic)
}
func (m *ContentCache) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentCache.Merge(m, src)
}
func (m *ContentCache) XXX_Size() int {
	return xxx_messageInfo_ContentCache.Size(m)
}
func (m *ContentCache) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentCache.DiscardUnknown(m)
}

var xxx_messageInfo_ContentCache proto.InternalMessageInfo

func (m *ContentCache) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func init() {
	proto.RegisterType((*Vertex)(nil), "errdefs.Vertex")
	proto.RegisterType((*Source)(nil), "errdefs.Source")
	proto.RegisterType((*FrontendCap)(nil), "errdefs.FrontendCap")
	proto.RegisterType((*Subrequest)(nil), "errdefs.Subrequest")
	proto.RegisterType((*Solve)(nil), "errdefs.Solve")
	proto.RegisterType((*FileAction)(nil), "errdefs.FileAction")
	proto.RegisterType((*ContentCache)(nil), "errdefs.ContentCache")
}

func init() { proto.RegisterFile("errdefs.proto", fileDescriptor_689dc58a5060aff5) }

var fileDescriptor_689dc58a5060aff5 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcd, 0x8e, 0xd3, 0x30,
	0x14, 0x85, 0x27, 0xbf, 0x43, 0x6e, 0x81, 0x85, 0x81, 0x51, 0x34, 0xab, 0x8c, 0xc5, 0xa2, 0x48,
	0x90, 0x48, 0xc3, 0x13, 0x40, 0xd1, 0x68, 0x66, 0x55, 0xc9, 0x95, 0xd8, 0xc7, 0xc9, 0x4d, 0x6b,
	0x48, 0x6c, 0xe3, 0xd8, 0xa8, 0xbc, 0x1b, 0x0f, 0x87, 0xe2, 0xa4, 0x65, 0x16, 0xdd, 0xe5, 0xe4,
	0xfb, 0x7c, 0xed, 0x63, 0xc3, 0x2b, 0x34, 0xa6, 0xc5, 0x6e, 0x2c, 0xb5, 0x51, 0x56, 0x91, 0xeb,
	0x25, 0xde, 0x7e, 0xdc, 0x0b, 0x7b, 0x70, 0xbc, 0x6c, 0xd4, 0x50, 0x0d, 0x8a, 0xff, 0xa9, 0xb8,
	0x13, 0x7d, 0xfb, 0x53, 0xd8, 0x6a, 0x54, 0xfd, 0x6f, 0x34, 0x95, 0xe6, 0x95, 0xd2, 0xcb, 0x32,
	0x5a, 0x40, 0xfa, 0x1d, 0x8d, 0xc5, 0x23, 0xb9, 0x81, 0xb4, 0x15, 0x7b, 0x1c, 0x6d, 0x1e, 0x14,
	0xc1, 0x3a, 0x63, 0x4b, 0xa2, 0x5b, 0x48, 0x77, 0xca, 0x99, 0x06, 0x09, 0x85, 0x58, 0xc8, 0x4e,
	0x79, 0xbe, 0xba, 0x7f, 0x5d, 0x6a, 0x5e, 0xce, 0xe4, 0x49, 0x76, 0x8a, 0x79, 0x46, 0xee, 0x20,
	0x35, 0xb5, 0xdc, 0xe3, 0x98, 0x87, 0x45, 0xb4, 0x5e, 0xdd, 0x67, 0x93, 0xc5, 0xa6, 0x3f, 0x6c,
	0x01, 0xf4, 0x0e, 0x56, 0x0f, 0x46, 0x49, 0x8b, 0xb2, 0xdd, 0xd4, 0x9a, 0x10, 0x88, 0x65, 0x3d,
	0xe0, 0xb2, 0xab, 0xff, 0xa6, 0x05, 0xc0, 0xce, 0x71, 0x83, 0xbf, 0x1c, 0x8e, 0xf6, 0xa2, 0xf1,
	0x37, 0x80, 0x64, 0x37, 0xf5, 0x21, 0xb7, 0xf0, 0x42, 0x48, 0xed, 0xec, 0xd3, 0xb7, 0x31, 0x0f,
	0x8a, 0x68, 0x9d, 0xb1, 0x73, 0x9e, 0xd8, 0xa0, 0x9c, 0xf4, 0x2c, 0x9c, 0xd9, 0x29, 0x93, 0x1b,
	0x08, 0x95, 0xce, 0x23, 0xdf, 0x25, 0x9d, 0x4e, 0xb9, 0xd5, 0x2c, 0x54, 0x9a, 0x7c, 0x80, 0xb8,
	0x13, 0x3d, 0xe6, 0xb1, 0x27, 0x6f, 0xca, 0xd3, 0x35, 0x3f, 0x88, 0x1e, 0xbf, 0x34, 0x56, 0x28,
	0xf9, 0x78, 0xc5, 0xbc, 0x42, 0x3e, 0x41, 0xd2, 0xd4, 0xcd, 0x01, 0xf3, 0xc4, 0xbb, 0xef, 0xce,
	0xee, 0xc6, 0xd7, 0xb3, 0x9b, 0x09, 0x3e, 0x5e, 0xb1, 0xd9, 0xfa, 0x9a, 0xc1, 0xf5, 0xe8, 0xf8,
	0x0f, 0x6c, 0x2c, 0xa5, 0x00, 0xff, 0xe7, 0x91, 0xb7, 0x90, 0x08, 0xd9, 0xe2, 0xd1, 0x37, 0x8c,
	0xd8, 0x1c, 0xe8, 0x7b, 0x78, 0xf9, 0x7c, 0xce, 0x65, 0x8b, 0xa7, 0xfe, 0x1d, 0x3f, 0xff, 0x0b,
	0x00, 0x00, 0xff, 0xff, 0x1e, 0xfa, 0x9c, 0x6f, 0x0f, 0x02, 0x00, 0x00,
}
