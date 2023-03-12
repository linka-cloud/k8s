/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: k8s.io/api/scheduling/v1alpha1/generated.proto

package v1alpha1

import (
	fmt "fmt"

	io "io"

	proto "github.com/gogo/protobuf/proto"

	k8s_io_api_core_v1 "go.linka.cloud/k8s/core/v1"

	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

func (m *PriorityClass) Reset()      { *m = PriorityClass{} }
func (*PriorityClass) ProtoMessage() {}
func (*PriorityClass) Descriptor() ([]byte, []int) {
	return fileDescriptor_260442fbb28d876a, []int{0}
}
func (m *PriorityClass) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PriorityClass) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *PriorityClass) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PriorityClass.Merge(m, src)
}
func (m *PriorityClass) XXX_Size() int {
	return m.Size()
}
func (m *PriorityClass) XXX_DiscardUnknown() {
	xxx_messageInfo_PriorityClass.DiscardUnknown(m)
}

var xxx_messageInfo_PriorityClass proto.InternalMessageInfo

func (m *PriorityClassList) Reset()      { *m = PriorityClassList{} }
func (*PriorityClassList) ProtoMessage() {}
func (*PriorityClassList) Descriptor() ([]byte, []int) {
	return fileDescriptor_260442fbb28d876a, []int{1}
}
func (m *PriorityClassList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PriorityClassList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *PriorityClassList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PriorityClassList.Merge(m, src)
}
func (m *PriorityClassList) XXX_Size() int {
	return m.Size()
}
func (m *PriorityClassList) XXX_DiscardUnknown() {
	xxx_messageInfo_PriorityClassList.DiscardUnknown(m)
}

var xxx_messageInfo_PriorityClassList proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PriorityClass)(nil), "go.linka.cloud.k8s.scheduling.v1alpha1.PriorityClass")
	proto.RegisterType((*PriorityClassList)(nil), "go.linka.cloud.k8s.scheduling.v1alpha1.PriorityClassList")
}

func init() {
	proto.RegisterFile("go.linka.cloud/k8s/scheduling/v1alpha1/generated.proto", fileDescriptor_260442fbb28d876a)
}

var fileDescriptor_260442fbb28d876a = []byte{
	// 480 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x31, 0x8f, 0xd3, 0x30,
	0x1c, 0xc5, 0xe3, 0x1e, 0x91, 0x7a, 0xae, 0x2a, 0x7a, 0x99, 0xa2, 0x0e, 0x4e, 0x54, 0x96, 0x2c,
	0x67, 0xd3, 0x13, 0x02, 0x24, 0xb6, 0x50, 0x09, 0x21, 0x81, 0xa8, 0x32, 0x30, 0x20, 0x06, 0xdc,
	0xd4, 0x97, 0x9a, 0x26, 0x71, 0x14, 0x3b, 0x95, 0xba, 0xf1, 0x11, 0xf8, 0x52, 0x48, 0x1d, 0x6f,
	0xbc, 0x29, 0xa2, 0x66, 0xe0, 0x3b, 0x30, 0xa1, 0xa4, 0x77, 0x97, 0xb6, 0x81, 0x83, 0x2d, 0xff,
	0xe7, 0xf7, 0x7b, 0x76, 0x5e, 0x62, 0x88, 0x97, 0xcf, 0x25, 0xe6, 0x82, 0xd0, 0x8c, 0x13, 0x19,
	0x2e, 0xd8, 0xbc, 0x88, 0x79, 0x1a, 0x91, 0xd5, 0x98, 0xc6, 0xd9, 0x82, 0x8e, 0x49, 0xc4, 0x52,
	0x96, 0x53, 0xc5, 0xe6, 0x38, 0xcb, 0x85, 0x12, 0x16, 0xda, 0xf9, 0x31, 0xcd, 0x38, 0x6e, 0xfc,
	0xf8, 0xd6, 0x3f, 0x3c, 0x8f, 0xb8, 0x5a, 0x14, 0x33, 0x1c, 0x8a, 0x84, 0x44, 0x22, 0x12, 0xa4,
	0xc6, 0x66, 0xc5, 0x65, 0x3d, 0xd5, 0x43, 0xfd, 0xb4, 0x8b, 0x1b, 0x8e, 0xf6, 0xb6, 0x0f, 0x45,
	0xce, 0xc8, 0xaa, 0xb5, 0xe5, 0xf0, 0x49, 0xe3, 0x49, 0x68, 0xb8, 0xe0, 0x29, 0xcb, 0xd7, 0x24,
	0x5b, 0x46, 0x95, 0x20, 0x49, 0xc2, 0x14, 0xfd, 0x13, 0x45, 0xfe, 0x46, 0xe5, 0x45, 0xaa, 0x78,
	0xc2, 0x5a, 0xc0, 0xd3, 0x7f, 0x01, 0xd5, 0xeb, 0x26, 0xf4, 0x98, 0x1b, 0xfd, 0xec, 0xc0, 0xfe,
	0x34, 0xe7, 0x22, 0xe7, 0x6a, 0xfd, 0x32, 0xa6, 0x52, 0x5a, 0x9f, 0x60, 0xb7, 0x3a, 0xd5, 0x9c,
	0x2a, 0x6a, 0x03, 0x17, 0x78, 0xbd, 0x8b, 0xc7, 0xb8, 0xa9, 0xed, 0x2e, 0x1c, 0x67, 0xcb, 0xa8,
	0x12, 0x24, 0xae, 0xdc, 0x78, 0x35, 0xc6, 0xef, 0x66, 0x9f, 0x59, 0xa8, 0xde, 0x32, 0x45, 0x7d,
	0x6b, 0x53, 0x3a, 0x86, 0x2e, 0x1d, 0xd8, 0x68, 0xc1, 0x5d, 0xaa, 0xe5, 0x40, 0x73, 0x45, 0xe3,
	0x82, 0xd9, 0x1d, 0x17, 0x78, 0xa6, 0x7f, 0xaa, 0x4b, 0xc7, 0x7c, 0x5f, 0x09, 0xc1, 0x4e, 0xb7,
	0x9e, 0xc1, 0x7e, 0x14, 0x8b, 0x19, 0x8d, 0x27, 0xec, 0x92, 0x16, 0xb1, 0xb2, 0x4f, 0x5c, 0xe0,
	0x75, 0xfd, 0x33, 0x5d, 0x3a, 0xfd, 0x57, 0xfb, 0x0b, 0xc1, 0xa1, 0xcf, 0x1a, 0xc3, 0xde, 0x9c,
	0xc9, 0x30, 0xe7, 0x99, 0xe2, 0x22, 0xb5, 0x1f, 0xb8, 0xc0, 0x3b, 0xf5, 0x1f, 0xea, 0xd2, 0xe9,
	0x4d, 0x1a, 0x39, 0xd8, 0xf7, 0x58, 0x11, 0x1c, 0x64, 0x39, 0x63, 0x49, 0x3d, 0x4d, 0x45, 0xcc,
	0xc3, 0xb5, 0x6d, 0xd6, 0xdc, 0x0b, 0x5d, 0x3a, 0x83, 0xe9, 0xd1, 0xda, 0xaf, 0xd2, 0x79, 0xd4,
	0xfe, 0xea, 0xf8, 0xd8, 0x16, 0xb4, 0x42, 0x47, 0xdf, 0x00, 0x3c, 0x3b, 0x68, 0xfa, 0x0d, 0x97,
	0xca, 0xfa, 0xd8, 0x6a, 0x1b, 0xff, 0x5f, 0xdb, 0x15, 0x5d, 0x77, 0x3d, 0xb8, 0xe9, 0xba, 0x7b,
	0xab, 0xec, 0x35, 0x1d, 0x40, 0x93, 0x2b, 0x96, 0x48, 0xbb, 0xe3, 0x9e, 0x78, 0xbd, 0x8b, 0x73,
	0x7c, 0xff, 0xff, 0x8f, 0x0f, 0xce, 0xe7, 0xf7, 0x6f, 0x92, 0xcd, 0xd7, 0x55, 0x46, 0xb0, 0x8b,
	0xf2, 0x27, 0x9b, 0x2d, 0x32, 0xae, 0xb6, 0xc8, 0xb8, 0xde, 0x22, 0xe3, 0x8b, 0x46, 0x60, 0xa3,
	0x11, 0xb8, 0xd2, 0x08, 0x5c, 0x6b, 0x04, 0xbe, 0x6b, 0x04, 0xbe, 0xfe, 0x40, 0xc6, 0x07, 0x74,
	0xff, 0xcd, 0xfc, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x5c, 0xcf, 0xda, 0x02, 0xba, 0x03, 0x00, 0x00,
}

func (m *PriorityClass) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PriorityClass) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PriorityClass) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PreemptionPolicy != nil {
		i -= len(*m.PreemptionPolicy)
		copy(dAtA[i:], *m.PreemptionPolicy)
		i = encodeVarintGenerated(dAtA, i, uint64(len(*m.PreemptionPolicy)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Description != nil {
		i -= len(*m.Description)
		copy(dAtA[i:], *m.Description)
		i = encodeVarintGenerated(dAtA, i, uint64(len(*m.Description)))
		i--
		dAtA[i] = 0x22
	}
	if m.GlobalDefault != nil {
		i--
		if *m.GlobalDefault {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.Value != nil {
		i = encodeVarintGenerated(dAtA, i, uint64(*m.Value))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.ObjectMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *PriorityClassList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PriorityClassList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PriorityClassList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Items) > 0 {
		for iNdEx := len(m.Items) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Items[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenerated(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.ListMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenerated(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PriorityClass) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if m.Value != nil {
		n += 1 + sovGenerated(uint64(*m.Value))
	}
	if m.GlobalDefault != nil {
		n += 2
	}
	if m.Description != nil {
		l = len(*m.Description)
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.PreemptionPolicy != nil {
		l = len(*m.PreemptionPolicy)
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func (m *PriorityClassList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ListMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *PriorityClass) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PriorityClass{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ObjectMeta), "ObjectMeta", "v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Value:` + valueToStringGenerated(this.Value) + `,`,
		`GlobalDefault:` + valueToStringGenerated(this.GlobalDefault) + `,`,
		`Description:` + valueToStringGenerated(this.Description) + `,`,
		`PreemptionPolicy:` + valueToStringGenerated(this.PreemptionPolicy) + `,`,
		`}`,
	}, "")
	return s
}
func (this *PriorityClassList) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForItems := "[]PriorityClass{"
	for _, f := range this.Items {
		repeatedStringForItems += strings.Replace(strings.Replace(f.String(), "PriorityClass", "PriorityClass", 1), `&`, ``, 1) + ","
	}
	repeatedStringForItems += "}"
	s := strings.Join([]string{`&PriorityClassList{`,
		`ListMeta:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ListMeta), "ListMeta", "v1.ListMeta", 1), `&`, ``, 1) + `,`,
		`Items:` + repeatedStringForItems + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *PriorityClass) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: PriorityClass: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PriorityClass: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Value = &v
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GlobalDefault", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
			b := bool(v != 0)
			m.GlobalDefault = &b
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Description = &s
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreemptionPolicy", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := k8s_io_api_core_v1.PreemptionPolicy(dAtA[iNdEx:postIndex])
			m.PreemptionPolicy = &s
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func (m *PriorityClassList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: PriorityClassList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PriorityClassList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ListMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, PriorityClass{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
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
				return 0, ErrInvalidLengthGenerated
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenerated
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenerated
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenerated        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenerated = fmt.Errorf("proto: unexpected end of group")
)
