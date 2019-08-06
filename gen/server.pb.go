// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/ernestrc/dispatchd/gen/server.proto

package gen

import (
	fmt "fmt"
	amqp "github.com/ernestrc/dispatchd/amqp"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ExchangeState struct {
	Name                 string      `protobuf:"bytes,1,opt,name=name" json:"name"`
	ExType               uint8       `protobuf:"varint,2,opt,name=ex_type,json=exType,casttype=uint8" json:"ex_type"`
	Passive              bool        `protobuf:"varint,3,opt,name=passive" json:"passive"`
	Durable              bool        `protobuf:"varint,4,opt,name=durable" json:"durable"`
	AutoDelete           bool        `protobuf:"varint,5,opt,name=auto_delete,json=autoDelete" json:"auto_delete"`
	Internal             bool        `protobuf:"varint,6,opt,name=internal" json:"internal"`
	System               bool        `protobuf:"varint,7,opt,name=system" json:"system"`
	Arguments            *amqp.Table `protobuf:"bytes,8,opt,name=arguments" json:"arguments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ExchangeState) Reset()         { *m = ExchangeState{} }
func (m *ExchangeState) String() string { return proto.CompactTextString(m) }
func (*ExchangeState) ProtoMessage()    {}
func (*ExchangeState) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d24e92367ef8f05, []int{0}
}
func (m *ExchangeState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExchangeState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExchangeState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExchangeState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExchangeState.Merge(m, src)
}
func (m *ExchangeState) XXX_Size() int {
	return m.Size()
}
func (m *ExchangeState) XXX_DiscardUnknown() {
	xxx_messageInfo_ExchangeState.DiscardUnknown(m)
}

var xxx_messageInfo_ExchangeState proto.InternalMessageInfo

type BindingState struct {
	Id                   []byte      `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	QueueName            string      `protobuf:"bytes,2,opt,name=queue_name,json=queueName" json:"queue_name"`
	ExchangeName         string      `protobuf:"bytes,3,opt,name=exchange_name,json=exchangeName" json:"exchange_name"`
	Key                  string      `protobuf:"bytes,4,opt,name=key" json:"key"`
	Arguments            *amqp.Table `protobuf:"bytes,5,opt,name=arguments" json:"arguments,omitempty"`
	Topic                bool        `protobuf:"varint,6,opt,name=topic" json:"topic"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *BindingState) Reset()         { *m = BindingState{} }
func (m *BindingState) String() string { return proto.CompactTextString(m) }
func (*BindingState) ProtoMessage()    {}
func (*BindingState) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d24e92367ef8f05, []int{1}
}
func (m *BindingState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BindingState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BindingState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BindingState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindingState.Merge(m, src)
}
func (m *BindingState) XXX_Size() int {
	return m.Size()
}
func (m *BindingState) XXX_DiscardUnknown() {
	xxx_messageInfo_BindingState.DiscardUnknown(m)
}

var xxx_messageInfo_BindingState proto.InternalMessageInfo

type QueueState struct {
	Name                 string      `protobuf:"bytes,1,opt,name=name" json:"name"`
	Durable              bool        `protobuf:"varint,2,opt,name=durable" json:"durable"`
	Arguments            *amqp.Table `protobuf:"bytes,3,opt,name=arguments" json:"arguments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *QueueState) Reset()         { *m = QueueState{} }
func (m *QueueState) String() string { return proto.CompactTextString(m) }
func (*QueueState) ProtoMessage()    {}
func (*QueueState) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d24e92367ef8f05, []int{2}
}
func (m *QueueState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueueState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueueState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueueState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueueState.Merge(m, src)
}
func (m *QueueState) XXX_Size() int {
	return m.Size()
}
func (m *QueueState) XXX_DiscardUnknown() {
	xxx_messageInfo_QueueState.DiscardUnknown(m)
}

var xxx_messageInfo_QueueState proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ExchangeState)(nil), "gen.ExchangeState")
	proto.RegisterType((*BindingState)(nil), "gen.BindingState")
	proto.RegisterType((*QueueState)(nil), "gen.QueueState")
}

func init() {
	proto.RegisterFile("github.com/ernestrc/dispatchd/gen/server.proto", fileDescriptor_8d24e92367ef8f05)
}

var fileDescriptor_8d24e92367ef8f05 = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4d, 0x8e, 0xd3, 0x30,
	0x18, 0x86, 0xeb, 0xa4, 0xbf, 0x5f, 0x5b, 0x16, 0x5e, 0x20, 0x6b, 0x84, 0xda, 0x68, 0x10, 0xa8,
	0xb3, 0xa0, 0x91, 0x58, 0x21, 0x96, 0x15, 0x6c, 0x91, 0x28, 0xb3, 0xaf, 0xdc, 0xe4, 0x23, 0xb5,
	0x68, 0x9c, 0x8c, 0x7f, 0x46, 0xad, 0xb8, 0x00, 0xc7, 0x80, 0x03, 0x70, 0x8f, 0x59, 0xce, 0x8e,
	0xdd, 0x08, 0xf5, 0x18, 0xac, 0x50, 0x9c, 0x16, 0x1c, 0x09, 0x0d, 0xb3, 0x89, 0xe2, 0xe7, 0x7d,
	0x1d, 0x7d, 0x79, 0x6c, 0x98, 0x67, 0xc2, 0x6c, 0xec, 0x7a, 0x9e, 0x14, 0x79, 0x8c, 0x4a, 0xa2,
	0x36, 0x2a, 0x89, 0x53, 0xa1, 0x4b, 0x6e, 0x92, 0x4d, 0x1a, 0x67, 0x28, 0x63, 0x8d, 0xea, 0x1a,
	0xd5, 0xbc, 0x54, 0x85, 0x29, 0x68, 0x98, 0xa1, 0x3c, 0x7b, 0x71, 0xff, 0x26, 0x9e, 0x5f, 0x95,
	0xee, 0x51, 0xef, 0x69, 0xd4, 0xb3, 0x22, 0x2b, 0x62, 0x87, 0xd7, 0xf6, 0xa3, 0x5b, 0xb9, 0x85,
	0x7b, 0xab, 0xeb, 0xe7, 0xdf, 0x03, 0x18, 0xbf, 0xdd, 0x25, 0x1b, 0x2e, 0x33, 0xfc, 0x60, 0xb8,
	0x41, 0xca, 0xa0, 0x2d, 0x79, 0x8e, 0x8c, 0x44, 0x64, 0x36, 0x58, 0xb4, 0x6f, 0xee, 0xa6, 0xad,
	0xa5, 0x23, 0xf4, 0x39, 0xf4, 0x70, 0xb7, 0x32, 0xfb, 0x12, 0x59, 0x10, 0x91, 0xd9, 0x78, 0x31,
	0xae, 0xc2, 0x5f, 0x77, 0xd3, 0x8e, 0x15, 0xd2, 0xbc, 0x5a, 0x76, 0x71, 0x77, 0xb9, 0x2f, 0x91,
	0x4e, 0xa0, 0x57, 0x72, 0xad, 0xc5, 0x35, 0xb2, 0x30, 0x22, 0xb3, 0xfe, 0xf1, 0x23, 0x27, 0x58,
	0xe5, 0xa9, 0x55, 0x7c, 0xbd, 0x45, 0xd6, 0xf6, 0xf3, 0x23, 0xa4, 0xcf, 0x60, 0xc8, 0xad, 0x29,
	0x56, 0x29, 0x6e, 0xd1, 0x20, 0xeb, 0x78, 0x1d, 0xa8, 0x82, 0x37, 0x8e, 0xd3, 0x08, 0xfa, 0x42,
	0x1a, 0x54, 0x92, 0x6f, 0x59, 0xd7, 0xeb, 0xfc, 0xa1, 0xf4, 0x09, 0x74, 0xf5, 0x5e, 0x1b, 0xcc,
	0x59, 0xcf, 0xcb, 0x8f, 0x8c, 0x5e, 0xc0, 0x80, 0xab, 0xcc, 0xe6, 0x28, 0x8d, 0x66, 0xfd, 0x88,
	0xcc, 0x86, 0x2f, 0x87, 0x73, 0x67, 0xf2, 0xb2, 0x1a, 0x63, 0xf9, 0x37, 0x7d, 0xdd, 0xff, 0xf2,
	0x75, 0xda, 0xba, 0xfd, 0x36, 0x6d, 0x9d, 0xff, 0x20, 0x30, 0x5a, 0x08, 0x99, 0x0a, 0x99, 0xd5,
	0xba, 0x1e, 0x41, 0x20, 0x52, 0x27, 0x6b, 0xb4, 0x0c, 0x44, 0x4a, 0x9f, 0x02, 0x5c, 0x59, 0xb4,
	0xb8, 0x72, 0x12, 0x03, 0x4f, 0xe2, 0xc0, 0xf1, 0x77, 0x95, 0xc9, 0x0b, 0x18, 0xe3, 0x51, 0x7a,
	0xdd, 0x0b, 0xbd, 0xde, 0xe8, 0x14, 0xb9, 0xea, 0x63, 0x08, 0x3f, 0xe1, 0xde, 0x89, 0x3a, 0x15,
	0x2a, 0xd0, 0x9c, 0xbe, 0x73, 0xdf, 0xf4, 0xf4, 0x0c, 0x3a, 0xa6, 0x28, 0x45, 0xd2, 0xb0, 0x54,
	0x23, 0xef, 0xcf, 0x3e, 0x03, 0xbc, 0xaf, 0x06, 0xfc, 0xdf, 0x2d, 0xf0, 0x4e, 0x2f, 0xf8, 0xd7,
	0xe9, 0x35, 0x06, 0x0b, 0x1f, 0xa6, 0x75, 0x31, 0xba, 0x39, 0x4c, 0xc8, 0xed, 0x61, 0x42, 0x7e,
	0x1e, 0x26, 0xe4, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5c, 0x3e, 0xa3, 0x26, 0x28, 0x03, 0x00,
	0x00,
}

func (m *ExchangeState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExchangeState) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintServer(dAtA, i, uint64(len(m.Name)))
	i += copy(dAtA[i:], m.Name)
	dAtA[i] = 0x10
	i++
	i = encodeVarintServer(dAtA, i, uint64(m.ExType))
	dAtA[i] = 0x18
	i++
	if m.Passive {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	dAtA[i] = 0x20
	i++
	if m.Durable {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	dAtA[i] = 0x28
	i++
	if m.AutoDelete {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	dAtA[i] = 0x30
	i++
	if m.Internal {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	dAtA[i] = 0x38
	i++
	if m.System {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	if m.Arguments != nil {
		dAtA[i] = 0x42
		i++
		i = encodeVarintServer(dAtA, i, uint64(m.Arguments.Size()))
		n1, err1 := m.Arguments.MarshalTo(dAtA[i:])
		if err1 != nil {
			return 0, err1
		}
		i += n1
	}
	return i, nil
}

func (m *BindingState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BindingState) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintServer(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintServer(dAtA, i, uint64(len(m.QueueName)))
	i += copy(dAtA[i:], m.QueueName)
	dAtA[i] = 0x1a
	i++
	i = encodeVarintServer(dAtA, i, uint64(len(m.ExchangeName)))
	i += copy(dAtA[i:], m.ExchangeName)
	dAtA[i] = 0x22
	i++
	i = encodeVarintServer(dAtA, i, uint64(len(m.Key)))
	i += copy(dAtA[i:], m.Key)
	if m.Arguments != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintServer(dAtA, i, uint64(m.Arguments.Size()))
		n2, err2 := m.Arguments.MarshalTo(dAtA[i:])
		if err2 != nil {
			return 0, err2
		}
		i += n2
	}
	dAtA[i] = 0x30
	i++
	if m.Topic {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	return i, nil
}

func (m *QueueState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueueState) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintServer(dAtA, i, uint64(len(m.Name)))
	i += copy(dAtA[i:], m.Name)
	dAtA[i] = 0x10
	i++
	if m.Durable {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	if m.Arguments != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintServer(dAtA, i, uint64(m.Arguments.Size()))
		n3, err3 := m.Arguments.MarshalTo(dAtA[i:])
		if err3 != nil {
			return 0, err3
		}
		i += n3
	}
	return i, nil
}

func encodeVarintServer(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ExchangeState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	n += 1 + l + sovServer(uint64(l))
	n += 1 + sovServer(uint64(m.ExType))
	n += 2
	n += 2
	n += 2
	n += 2
	n += 2
	if m.Arguments != nil {
		l = m.Arguments.Size()
		n += 1 + l + sovServer(uint64(l))
	}
	return n
}

func (m *BindingState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != nil {
		l = len(m.Id)
		n += 1 + l + sovServer(uint64(l))
	}
	l = len(m.QueueName)
	n += 1 + l + sovServer(uint64(l))
	l = len(m.ExchangeName)
	n += 1 + l + sovServer(uint64(l))
	l = len(m.Key)
	n += 1 + l + sovServer(uint64(l))
	if m.Arguments != nil {
		l = m.Arguments.Size()
		n += 1 + l + sovServer(uint64(l))
	}
	n += 2
	return n
}

func (m *QueueState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	n += 1 + l + sovServer(uint64(l))
	n += 2
	if m.Arguments != nil {
		l = m.Arguments.Size()
		n += 1 + l + sovServer(uint64(l))
	}
	return n
}

func sovServer(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozServer(x uint64) (n int) {
	return sovServer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ExchangeState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowServer
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ExchangeState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExchangeState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthServer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthServer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExType", wireType)
			}
			m.ExType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExType |= uint8(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Passive", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Passive = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Durable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Durable = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AutoDelete", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AutoDelete = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Internal", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Internal = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field System", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.System = bool(v != 0)
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Arguments", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthServer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthServer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Arguments == nil {
				m.Arguments = &amqp.Table{}
			}
			if err := m.Arguments.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipServer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthServer
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthServer
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BindingState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowServer
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BindingState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BindingState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthServer
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthServer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = append(m.Id[:0], dAtA[iNdEx:postIndex]...)
			if m.Id == nil {
				m.Id = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueueName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthServer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthServer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QueueName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExchangeName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthServer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthServer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExchangeName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthServer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthServer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Arguments", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthServer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthServer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Arguments == nil {
				m.Arguments = &amqp.Table{}
			}
			if err := m.Arguments.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Topic", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Topic = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipServer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthServer
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthServer
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueueState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowServer
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueueState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueueState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthServer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthServer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Durable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Durable = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Arguments", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthServer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthServer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Arguments == nil {
				m.Arguments = &amqp.Table{}
			}
			if err := m.Arguments.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipServer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthServer
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthServer
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipServer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowServer
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowServer
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowServer
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthServer
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthServer
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowServer
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipServer(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthServer
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthServer = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowServer   = fmt.Errorf("proto: integer overflow")
)
