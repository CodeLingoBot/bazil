// Code generated by protoc-gen-gogo.
// source: bazil.org/bazil/fs/wire/snap.proto
// DO NOT EDIT!

package wire

import proto "github.com/gogo/protobuf/proto"
import math "math"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto/gogo.pb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

// Snapshot as it is stored into database.
type SnapshotRef struct {
	Key              []byte `protobuf:"bytes,1,req,name=key" json:"key"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SnapshotRef) Reset()         { *m = SnapshotRef{} }
func (m *SnapshotRef) String() string { return proto.CompactTextString(m) }
func (*SnapshotRef) ProtoMessage()    {}

func (m *SnapshotRef) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func init() {
}
