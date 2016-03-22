// Code generated by protoc-gen-gogo.
// source: broker.proto
// DO NOT EDIT!

package core

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DataBrokerReq struct {
	Payload  *LoRaWANData `protobuf:"bytes,1,opt,name=Payload,json=payload" json:"Payload,omitempty"`
	Metadata *Metadata    `protobuf:"bytes,2,opt,name=Metadata,json=metadata" json:"Metadata,omitempty"`
}

func (m *DataBrokerReq) Reset()                    { *m = DataBrokerReq{} }
func (m *DataBrokerReq) String() string            { return proto.CompactTextString(m) }
func (*DataBrokerReq) ProtoMessage()               {}
func (*DataBrokerReq) Descriptor() ([]byte, []int) { return fileDescriptorBroker, []int{0} }

func (m *DataBrokerReq) GetPayload() *LoRaWANData {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *DataBrokerReq) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type DataBrokerRes struct {
	Payload  *LoRaWANData `protobuf:"bytes,1,opt,name=Payload,json=payload" json:"Payload,omitempty"`
	Metadata *Metadata    `protobuf:"bytes,2,opt,name=Metadata,json=metadata" json:"Metadata,omitempty"`
}

func (m *DataBrokerRes) Reset()                    { *m = DataBrokerRes{} }
func (m *DataBrokerRes) String() string            { return proto.CompactTextString(m) }
func (*DataBrokerRes) ProtoMessage()               {}
func (*DataBrokerRes) Descriptor() ([]byte, []int) { return fileDescriptorBroker, []int{1} }

func (m *DataBrokerRes) GetPayload() *LoRaWANData {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *DataBrokerRes) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type JoinBrokerReq struct {
	AppEUI   []byte    `protobuf:"bytes,1,opt,name=AppEUI,json=appEUI,proto3" json:"AppEUI,omitempty"`
	DevEUI   []byte    `protobuf:"bytes,2,opt,name=DevEUI,json=devEUI,proto3" json:"DevEUI,omitempty"`
	DevNonce []byte    `protobuf:"bytes,3,opt,name=DevNonce,json=devNonce,proto3" json:"DevNonce,omitempty"`
	Metadata *Metadata `protobuf:"bytes,4,opt,name=Metadata,json=metadata" json:"Metadata,omitempty"`
}

func (m *JoinBrokerReq) Reset()                    { *m = JoinBrokerReq{} }
func (m *JoinBrokerReq) String() string            { return proto.CompactTextString(m) }
func (*JoinBrokerReq) ProtoMessage()               {}
func (*JoinBrokerReq) Descriptor() ([]byte, []int) { return fileDescriptorBroker, []int{2} }

func (m *JoinBrokerReq) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type JoinBrokerRes struct {
	Payload  *LoRaWANJoinAccept `protobuf:"bytes,1,opt,name=Payload,json=payload" json:"Payload,omitempty"`
	DevAddr  []byte             `protobuf:"bytes,2,opt,name=DevAddr,json=devAddr,proto3" json:"DevAddr,omitempty"`
	Metadata *Metadata          `protobuf:"bytes,3,opt,name=Metadata,json=metadata" json:"Metadata,omitempty"`
}

func (m *JoinBrokerRes) Reset()                    { *m = JoinBrokerRes{} }
func (m *JoinBrokerRes) String() string            { return proto.CompactTextString(m) }
func (*JoinBrokerRes) ProtoMessage()               {}
func (*JoinBrokerRes) Descriptor() ([]byte, []int) { return fileDescriptorBroker, []int{3} }

func (m *JoinBrokerRes) GetPayload() *LoRaWANJoinAccept {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *JoinBrokerRes) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*DataBrokerReq)(nil), "core.DataBrokerReq")
	proto.RegisterType((*DataBrokerRes)(nil), "core.DataBrokerRes")
	proto.RegisterType((*JoinBrokerReq)(nil), "core.JoinBrokerReq")
	proto.RegisterType((*JoinBrokerRes)(nil), "core.JoinBrokerRes")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Broker service

type BrokerClient interface {
	HandleData(ctx context.Context, in *DataBrokerReq, opts ...grpc.CallOption) (*DataBrokerRes, error)
	HandleJoin(ctx context.Context, in *JoinBrokerReq, opts ...grpc.CallOption) (*JoinBrokerRes, error)
}

type brokerClient struct {
	cc *grpc.ClientConn
}

func NewBrokerClient(cc *grpc.ClientConn) BrokerClient {
	return &brokerClient{cc}
}

func (c *brokerClient) HandleData(ctx context.Context, in *DataBrokerReq, opts ...grpc.CallOption) (*DataBrokerRes, error) {
	out := new(DataBrokerRes)
	err := grpc.Invoke(ctx, "/core.Broker/HandleData", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brokerClient) HandleJoin(ctx context.Context, in *JoinBrokerReq, opts ...grpc.CallOption) (*JoinBrokerRes, error) {
	out := new(JoinBrokerRes)
	err := grpc.Invoke(ctx, "/core.Broker/HandleJoin", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Broker service

type BrokerServer interface {
	HandleData(context.Context, *DataBrokerReq) (*DataBrokerRes, error)
	HandleJoin(context.Context, *JoinBrokerReq) (*JoinBrokerRes, error)
}

func RegisterBrokerServer(s *grpc.Server, srv BrokerServer) {
	s.RegisterService(&_Broker_serviceDesc, srv)
}

func _Broker_HandleData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(DataBrokerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BrokerServer).HandleData(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Broker_HandleJoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(JoinBrokerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BrokerServer).HandleJoin(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Broker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "core.Broker",
	HandlerType: (*BrokerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleData",
			Handler:    _Broker_HandleData_Handler,
		},
		{
			MethodName: "HandleJoin",
			Handler:    _Broker_HandleJoin_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

func (m *DataBrokerReq) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *DataBrokerReq) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Payload != nil {
		data[i] = 0xa
		i++
		i = encodeVarintBroker(data, i, uint64(m.Payload.Size()))
		n1, err := m.Payload.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Metadata != nil {
		data[i] = 0x12
		i++
		i = encodeVarintBroker(data, i, uint64(m.Metadata.Size()))
		n2, err := m.Metadata.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *DataBrokerRes) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *DataBrokerRes) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Payload != nil {
		data[i] = 0xa
		i++
		i = encodeVarintBroker(data, i, uint64(m.Payload.Size()))
		n3, err := m.Payload.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.Metadata != nil {
		data[i] = 0x12
		i++
		i = encodeVarintBroker(data, i, uint64(m.Metadata.Size()))
		n4, err := m.Metadata.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}

func (m *JoinBrokerReq) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *JoinBrokerReq) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.AppEUI != nil {
		if len(m.AppEUI) > 0 {
			data[i] = 0xa
			i++
			i = encodeVarintBroker(data, i, uint64(len(m.AppEUI)))
			i += copy(data[i:], m.AppEUI)
		}
	}
	if m.DevEUI != nil {
		if len(m.DevEUI) > 0 {
			data[i] = 0x12
			i++
			i = encodeVarintBroker(data, i, uint64(len(m.DevEUI)))
			i += copy(data[i:], m.DevEUI)
		}
	}
	if m.DevNonce != nil {
		if len(m.DevNonce) > 0 {
			data[i] = 0x1a
			i++
			i = encodeVarintBroker(data, i, uint64(len(m.DevNonce)))
			i += copy(data[i:], m.DevNonce)
		}
	}
	if m.Metadata != nil {
		data[i] = 0x22
		i++
		i = encodeVarintBroker(data, i, uint64(m.Metadata.Size()))
		n5, err := m.Metadata.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	return i, nil
}

func (m *JoinBrokerRes) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *JoinBrokerRes) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Payload != nil {
		data[i] = 0xa
		i++
		i = encodeVarintBroker(data, i, uint64(m.Payload.Size()))
		n6, err := m.Payload.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if m.DevAddr != nil {
		if len(m.DevAddr) > 0 {
			data[i] = 0x12
			i++
			i = encodeVarintBroker(data, i, uint64(len(m.DevAddr)))
			i += copy(data[i:], m.DevAddr)
		}
	}
	if m.Metadata != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintBroker(data, i, uint64(m.Metadata.Size()))
		n7, err := m.Metadata.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	return i, nil
}

func encodeFixed64Broker(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Broker(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintBroker(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *DataBrokerReq) Size() (n int) {
	var l int
	_ = l
	if m.Payload != nil {
		l = m.Payload.Size()
		n += 1 + l + sovBroker(uint64(l))
	}
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovBroker(uint64(l))
	}
	return n
}

func (m *DataBrokerRes) Size() (n int) {
	var l int
	_ = l
	if m.Payload != nil {
		l = m.Payload.Size()
		n += 1 + l + sovBroker(uint64(l))
	}
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovBroker(uint64(l))
	}
	return n
}

func (m *JoinBrokerReq) Size() (n int) {
	var l int
	_ = l
	if m.AppEUI != nil {
		l = len(m.AppEUI)
		if l > 0 {
			n += 1 + l + sovBroker(uint64(l))
		}
	}
	if m.DevEUI != nil {
		l = len(m.DevEUI)
		if l > 0 {
			n += 1 + l + sovBroker(uint64(l))
		}
	}
	if m.DevNonce != nil {
		l = len(m.DevNonce)
		if l > 0 {
			n += 1 + l + sovBroker(uint64(l))
		}
	}
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovBroker(uint64(l))
	}
	return n
}

func (m *JoinBrokerRes) Size() (n int) {
	var l int
	_ = l
	if m.Payload != nil {
		l = m.Payload.Size()
		n += 1 + l + sovBroker(uint64(l))
	}
	if m.DevAddr != nil {
		l = len(m.DevAddr)
		if l > 0 {
			n += 1 + l + sovBroker(uint64(l))
		}
	}
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovBroker(uint64(l))
	}
	return n
}

func sovBroker(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozBroker(x uint64) (n int) {
	return sovBroker(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DataBrokerReq) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBroker
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DataBrokerReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataBrokerReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBroker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBroker
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Payload == nil {
				m.Payload = &LoRaWANData{}
			}
			if err := m.Payload.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBroker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBroker
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &Metadata{}
			}
			if err := m.Metadata.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBroker(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBroker
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
func (m *DataBrokerRes) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBroker
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DataBrokerRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataBrokerRes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBroker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBroker
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Payload == nil {
				m.Payload = &LoRaWANData{}
			}
			if err := m.Payload.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBroker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBroker
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &Metadata{}
			}
			if err := m.Metadata.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBroker(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBroker
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
func (m *JoinBrokerReq) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBroker
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: JoinBrokerReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: JoinBrokerReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppEUI", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBroker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBroker
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppEUI = append(m.AppEUI[:0], data[iNdEx:postIndex]...)
			if m.AppEUI == nil {
				m.AppEUI = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DevEUI", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBroker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBroker
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DevEUI = append(m.DevEUI[:0], data[iNdEx:postIndex]...)
			if m.DevEUI == nil {
				m.DevEUI = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DevNonce", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBroker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBroker
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DevNonce = append(m.DevNonce[:0], data[iNdEx:postIndex]...)
			if m.DevNonce == nil {
				m.DevNonce = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBroker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBroker
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &Metadata{}
			}
			if err := m.Metadata.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBroker(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBroker
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
func (m *JoinBrokerRes) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBroker
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: JoinBrokerRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: JoinBrokerRes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBroker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBroker
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Payload == nil {
				m.Payload = &LoRaWANJoinAccept{}
			}
			if err := m.Payload.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DevAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBroker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBroker
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DevAddr = append(m.DevAddr[:0], data[iNdEx:postIndex]...)
			if m.DevAddr == nil {
				m.DevAddr = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBroker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBroker
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &Metadata{}
			}
			if err := m.Metadata.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBroker(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBroker
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
func skipBroker(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBroker
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowBroker
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowBroker
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthBroker
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowBroker
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipBroker(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthBroker = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBroker   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorBroker = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x2a, 0xca, 0xcf,
	0x4e, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x49, 0xce, 0x2f, 0x4a, 0x95, 0xe2,
	0xcd, 0xc9, 0x2f, 0x4a, 0x2c, 0x4f, 0xcc, 0x83, 0x08, 0x4a, 0x71, 0x81, 0x04, 0x21, 0x6c, 0xa5,
	0x0c, 0x2e, 0x5e, 0x97, 0xc4, 0x92, 0x44, 0x27, 0xb0, 0xa6, 0xa0, 0xd4, 0x42, 0x21, 0x6d, 0x2e,
	0xf6, 0x80, 0xc4, 0xca, 0x9c, 0xfc, 0xc4, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x41,
	0x3d, 0xb0, 0x72, 0x9f, 0xfc, 0xa0, 0xc4, 0x70, 0x47, 0x3f, 0x90, 0xe2, 0x20, 0xf6, 0x02, 0x88,
	0x0a, 0x21, 0x2d, 0x2e, 0x0e, 0xdf, 0xd4, 0x92, 0xc4, 0x14, 0xa0, 0xa0, 0x04, 0x13, 0x58, 0x35,
	0x1f, 0x44, 0x35, 0x4c, 0x34, 0x88, 0x23, 0x17, 0xca, 0x42, 0xb7, 0xa9, 0x98, 0x76, 0x36, 0xb5,
	0x33, 0x72, 0xf1, 0x7a, 0xe5, 0x67, 0xe6, 0x21, 0x3c, 0x25, 0xc6, 0xc5, 0xe6, 0x58, 0x50, 0xe0,
	0x1a, 0xea, 0x09, 0xb6, 0x89, 0x27, 0x88, 0x2d, 0x11, 0xcc, 0x03, 0x89, 0xbb, 0xa4, 0x96, 0x81,
	0xc4, 0x99, 0x20, 0xe2, 0x29, 0x60, 0x9e, 0x90, 0x14, 0x17, 0x07, 0x50, 0xdc, 0x2f, 0x3f, 0x2f,
	0x39, 0x55, 0x82, 0x19, 0x2c, 0xc3, 0x91, 0x02, 0xe5, 0xa3, 0xb8, 0x84, 0x85, 0x80, 0x4b, 0x3a,
	0xd0, 0x5c, 0x52, 0x2c, 0x64, 0x88, 0xee, 0x69, 0x71, 0x14, 0x4f, 0x83, 0x14, 0x3b, 0x26, 0x27,
	0xa7, 0x16, 0x94, 0x20, 0xbc, 0x2e, 0xc1, 0xc5, 0x0e, 0x74, 0x8c, 0x63, 0x4a, 0x4a, 0x11, 0xd4,
	0x95, 0xec, 0x29, 0x10, 0x2e, 0x8a, 0x53, 0x98, 0xf1, 0x3b, 0xc5, 0xa8, 0x82, 0x8b, 0x0d, 0xe2,
	0x0a, 0x21, 0x33, 0x2e, 0x2e, 0x8f, 0xc4, 0xbc, 0x94, 0x9c, 0x54, 0x50, 0x08, 0x0b, 0x09, 0x43,
	0x74, 0xa0, 0x24, 0x02, 0x29, 0x2c, 0x82, 0xc5, 0x08, 0x7d, 0x20, 0x47, 0xc2, 0xf4, 0xa1, 0x84,
	0xb3, 0x14, 0x16, 0xc1, 0x62, 0x27, 0x81, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x00, 0xf1, 0x03, 0x20,
	0x9e, 0xf1, 0x58, 0x8e, 0x21, 0x89, 0x0d, 0x9c, 0xf6, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x57, 0x4c, 0xa7, 0xbb, 0xac, 0x02, 0x00, 0x00,
}
