// Code generated by protoc-gen-gogo.
// source: bazil.org/bazil/cas/wire/manifest.proto
// DO NOT EDIT!

/*
Package wire is a generated protocol buffer package.

It is generated from these files:
	bazil.org/bazil/cas/wire/manifest.proto

It has these top-level messages:
	Manifest
*/
package wire

import proto "github.com/gogo/protobuf/proto"
import math "math"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto/gogo.pb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Manifest struct {
	Root             []byte `protobuf:"bytes,1,req,name=root" json:"root"`
	Size_            uint64 `protobuf:"varint,2,req,name=size" json:"size"`
	ChunkSize        uint32 `protobuf:"varint,3,req,name=chunkSize" json:"chunkSize"`
	Fanout           uint32 `protobuf:"varint,4,req,name=fanout" json:"fanout"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Manifest) Reset()         { *m = Manifest{} }
func (m *Manifest) String() string { return proto.CompactTextString(m) }
func (*Manifest) ProtoMessage()    {}

func (m *Manifest) GetRoot() []byte {
	if m != nil {
		return m.Root
	}
	return nil
}

func (m *Manifest) GetSize_() uint64 {
	if m != nil {
		return m.Size_
	}
	return 0
}

func (m *Manifest) GetChunkSize() uint32 {
	if m != nil {
		return m.ChunkSize
	}
	return 0
}

func (m *Manifest) GetFanout() uint32 {
	if m != nil {
		return m.Fanout
	}
	return 0
}

func init() {
}
