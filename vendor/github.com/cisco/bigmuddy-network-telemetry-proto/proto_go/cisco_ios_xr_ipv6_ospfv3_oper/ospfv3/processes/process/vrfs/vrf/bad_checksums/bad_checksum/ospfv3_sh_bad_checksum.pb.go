// Code generated by protoc-gen-go.
// source: ospfv3_sh_bad_checksum.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_bad_checksums_bad_checksum is a generated protocol buffer package.

It is generated from these files:
	ospfv3_sh_bad_checksum.proto

It has these top-level messages:
	Ospfv3ShBadChecksum_KEYS
	Ospfv3ShBadChecksum
	Ospfv3EdmTime
*/
package cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_bad_checksums_bad_checksum

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

// OSPFv3 Bad Checksum
type Ospfv3ShBadChecksum_KEYS struct {
	ProcessName  string `protobuf:"bytes,1,opt,name=process_name,json=processName" json:"process_name,omitempty"`
	VrfName      string `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	PacketNumber uint32 `protobuf:"varint,3,opt,name=packet_number,json=packetNumber" json:"packet_number,omitempty"`
}

func (m *Ospfv3ShBadChecksum_KEYS) Reset()                    { *m = Ospfv3ShBadChecksum_KEYS{} }
func (m *Ospfv3ShBadChecksum_KEYS) String() string            { return proto.CompactTextString(m) }
func (*Ospfv3ShBadChecksum_KEYS) ProtoMessage()               {}
func (*Ospfv3ShBadChecksum_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Ospfv3ShBadChecksum_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *Ospfv3ShBadChecksum_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ospfv3ShBadChecksum_KEYS) GetPacketNumber() uint32 {
	if m != nil {
		return m.PacketNumber
	}
	return 0
}

type Ospfv3ShBadChecksum struct {
	// Received Checksum
	ReceivedChecksum uint32 `protobuf:"varint,50,opt,name=received_checksum,json=receivedChecksum" json:"received_checksum,omitempty"`
	// Computed Checksum
	ComputedChecksum uint32 `protobuf:"varint,51,opt,name=computed_checksum,json=computedChecksum" json:"computed_checksum,omitempty"`
	// Packet Timestamp (relative to 1970/1/1 00:00)
	Timestamp *Ospfv3EdmTime `protobuf:"bytes,52,opt,name=timestamp" json:"timestamp,omitempty"`
	// Received Hexadecimal Data
	ReceivedData []byte `protobuf:"bytes,53,opt,name=received_data,json=receivedData,proto3" json:"received_data,omitempty"`
}

func (m *Ospfv3ShBadChecksum) Reset()                    { *m = Ospfv3ShBadChecksum{} }
func (m *Ospfv3ShBadChecksum) String() string            { return proto.CompactTextString(m) }
func (*Ospfv3ShBadChecksum) ProtoMessage()               {}
func (*Ospfv3ShBadChecksum) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Ospfv3ShBadChecksum) GetReceivedChecksum() uint32 {
	if m != nil {
		return m.ReceivedChecksum
	}
	return 0
}

func (m *Ospfv3ShBadChecksum) GetComputedChecksum() uint32 {
	if m != nil {
		return m.ComputedChecksum
	}
	return 0
}

func (m *Ospfv3ShBadChecksum) GetTimestamp() *Ospfv3EdmTime {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Ospfv3ShBadChecksum) GetReceivedData() []byte {
	if m != nil {
		return m.ReceivedData
	}
	return nil
}

// OSPFv3 time stamp
type Ospfv3EdmTime struct {
	// Seconds
	Second uint32 `protobuf:"varint,1,opt,name=second" json:"second,omitempty"`
	// Nano seconds
	Nanosecond uint32 `protobuf:"varint,2,opt,name=nanosecond" json:"nanosecond,omitempty"`
}

func (m *Ospfv3EdmTime) Reset()                    { *m = Ospfv3EdmTime{} }
func (m *Ospfv3EdmTime) String() string            { return proto.CompactTextString(m) }
func (*Ospfv3EdmTime) ProtoMessage()               {}
func (*Ospfv3EdmTime) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Ospfv3EdmTime) GetSecond() uint32 {
	if m != nil {
		return m.Second
	}
	return 0
}

func (m *Ospfv3EdmTime) GetNanosecond() uint32 {
	if m != nil {
		return m.Nanosecond
	}
	return 0
}

func init() {
	proto.RegisterType((*Ospfv3ShBadChecksum_KEYS)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.bad_checksums.bad_checksum.ospfv3_sh_bad_checksum_KEYS")
	proto.RegisterType((*Ospfv3ShBadChecksum)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.bad_checksums.bad_checksum.ospfv3_sh_bad_checksum")
	proto.RegisterType((*Ospfv3EdmTime)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.bad_checksums.bad_checksum.ospfv3_edm_time")
}

func init() { proto.RegisterFile("ospfv3_sh_bad_checksum.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xc9, 0x84, 0xe9, 0xb2, 0x0e, 0xb5, 0x87, 0x51, 0x51, 0xa4, 0xce, 0x4b, 0x41, 0xe8,
	0x61, 0x53, 0xbf, 0x80, 0x7a, 0x10, 0x61, 0x87, 0x7a, 0xd2, 0xcb, 0x23, 0x4b, 0x5f, 0x59, 0x19,
	0x69, 0x42, 0x92, 0x15, 0x8f, 0x5e, 0xfd, 0x1c, 0x7e, 0x51, 0x59, 0x9a, 0x6e, 0x55, 0x76, 0xf5,
	0x52, 0xde, 0xfb, 0xff, 0xfe, 0xfd, 0xbf, 0xf6, 0x25, 0xf4, 0x42, 0x1a, 0x55, 0xd4, 0x33, 0x30,
	0x4b, 0x58, 0xb0, 0x1c, 0xf8, 0x12, 0xf9, 0xca, 0xac, 0x45, 0xaa, 0xb4, 0xb4, 0x32, 0x7c, 0xe7,
	0xa5, 0xe1, 0x12, 0x4a, 0x69, 0xe0, 0x43, 0x43, 0xa9, 0xea, 0x7b, 0xf0, 0x7e, 0xa9, 0x50, 0xa7,
	0x4d, 0xbd, 0xf1, 0x72, 0x34, 0x06, 0x4d, 0x5b, 0xa5, 0xb5, 0x2e, 0xdc, 0x23, 0xed, 0x66, 0x9a,
	0x5f, 0xdd, 0xe4, 0x93, 0xd0, 0xf3, 0xfd, 0xc3, 0xe1, 0xe5, 0xe9, 0xed, 0x35, 0xbc, 0xa2, 0x81,
	0x8f, 0x83, 0x8a, 0x09, 0x8c, 0x48, 0x4c, 0x92, 0x41, 0x36, 0xf4, 0xda, 0x9c, 0x09, 0x0c, 0xcf,
	0xe8, 0x51, 0xad, 0x8b, 0x06, 0xf7, 0x1c, 0x3e, 0xac, 0x75, 0xe1, 0xd0, 0x35, 0x1d, 0x29, 0xc6,
	0x57, 0x68, 0xa1, 0x5a, 0x8b, 0x05, 0xea, 0xe8, 0x20, 0x26, 0xc9, 0x28, 0x0b, 0x1a, 0x71, 0xee,
	0xb4, 0xc9, 0x77, 0x8f, 0x8e, 0xf7, 0x7f, 0x42, 0x78, 0x43, 0x4f, 0x35, 0x72, 0x2c, 0x6b, 0xdc,
	0x89, 0xd1, 0xd4, 0x65, 0x9c, 0xb4, 0xe0, 0xa1, 0x63, 0xe6, 0x52, 0xa8, 0xb5, 0xed, 0x9a, 0x67,
	0x8d, 0xb9, 0x05, 0x5b, 0xf3, 0x17, 0xa1, 0x03, 0x5b, 0x0a, 0x34, 0x96, 0x09, 0x15, 0xdd, 0xc6,
	0x24, 0x19, 0x4e, 0x57, 0xe9, 0xff, 0x2d, 0xda, 0xbf, 0x0c, 0x98, 0x0b, 0xd8, 0xcc, 0xcd, 0x76,
	0xd3, 0x37, 0x5b, 0xda, 0xfe, 0x65, 0xce, 0x2c, 0x8b, 0xee, 0x62, 0x92, 0x04, 0x59, 0xd0, 0x8a,
	0x8f, 0xcc, 0xb2, 0xc9, 0x33, 0x3d, 0xfe, 0x13, 0x11, 0x8e, 0x69, 0xdf, 0x20, 0x97, 0x55, 0xee,
	0x4e, 0x65, 0x94, 0xf9, 0x2e, 0xbc, 0xa4, 0xb4, 0x62, 0x95, 0xf4, 0xac, 0xe7, 0x58, 0x47, 0x59,
	0xf4, 0xdd, 0xb5, 0x9a, 0xfd, 0x04, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x18, 0x59, 0xeb, 0x76, 0x02,
	0x00, 0x00,
}
