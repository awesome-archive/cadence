// The MIT License (MIT)

// Copyright (c) 2017-2020 Uber Technologies Inc.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: uber/cadence/shared/v1/error.proto

package sharedv1

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/gogo/protobuf/proto"

	v11 "github.com/uber/cadence-idl/go/proto/admin/v1"
	v1 "github.com/uber/cadence-idl/go/proto/api/v1"
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

type CurrentBranchChangedError struct {
	CurrentBranchToken   []byte   `protobuf:"bytes,1,opt,name=current_branch_token,json=currentBranchToken,proto3" json:"current_branch_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CurrentBranchChangedError) Reset()         { *m = CurrentBranchChangedError{} }
func (m *CurrentBranchChangedError) String() string { return proto.CompactTextString(m) }
func (*CurrentBranchChangedError) ProtoMessage()    {}
func (*CurrentBranchChangedError) Descriptor() ([]byte, []int) {
	return fileDescriptor_3688ca0fd170c8f9, []int{0}
}
func (m *CurrentBranchChangedError) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CurrentBranchChangedError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CurrentBranchChangedError.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CurrentBranchChangedError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CurrentBranchChangedError.Merge(m, src)
}
func (m *CurrentBranchChangedError) XXX_Size() int {
	return m.Size()
}
func (m *CurrentBranchChangedError) XXX_DiscardUnknown() {
	xxx_messageInfo_CurrentBranchChangedError.DiscardUnknown(m)
}

var xxx_messageInfo_CurrentBranchChangedError proto.InternalMessageInfo

func (m *CurrentBranchChangedError) GetCurrentBranchToken() []byte {
	if m != nil {
		return m.CurrentBranchToken
	}
	return nil
}

type InternalDataInconsistencyError struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InternalDataInconsistencyError) Reset()         { *m = InternalDataInconsistencyError{} }
func (m *InternalDataInconsistencyError) String() string { return proto.CompactTextString(m) }
func (*InternalDataInconsistencyError) ProtoMessage()    {}
func (*InternalDataInconsistencyError) Descriptor() ([]byte, []int) {
	return fileDescriptor_3688ca0fd170c8f9, []int{1}
}
func (m *InternalDataInconsistencyError) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InternalDataInconsistencyError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InternalDataInconsistencyError.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InternalDataInconsistencyError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InternalDataInconsistencyError.Merge(m, src)
}
func (m *InternalDataInconsistencyError) XXX_Size() int {
	return m.Size()
}
func (m *InternalDataInconsistencyError) XXX_DiscardUnknown() {
	xxx_messageInfo_InternalDataInconsistencyError.DiscardUnknown(m)
}

var xxx_messageInfo_InternalDataInconsistencyError proto.InternalMessageInfo

type EventAlreadyStartedError struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventAlreadyStartedError) Reset()         { *m = EventAlreadyStartedError{} }
func (m *EventAlreadyStartedError) String() string { return proto.CompactTextString(m) }
func (*EventAlreadyStartedError) ProtoMessage()    {}
func (*EventAlreadyStartedError) Descriptor() ([]byte, []int) {
	return fileDescriptor_3688ca0fd170c8f9, []int{2}
}
func (m *EventAlreadyStartedError) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventAlreadyStartedError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventAlreadyStartedError.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventAlreadyStartedError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventAlreadyStartedError.Merge(m, src)
}
func (m *EventAlreadyStartedError) XXX_Size() int {
	return m.Size()
}
func (m *EventAlreadyStartedError) XXX_DiscardUnknown() {
	xxx_messageInfo_EventAlreadyStartedError.DiscardUnknown(m)
}

var xxx_messageInfo_EventAlreadyStartedError proto.InternalMessageInfo

type RemoteSyncMatchedError struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoteSyncMatchedError) Reset()         { *m = RemoteSyncMatchedError{} }
func (m *RemoteSyncMatchedError) String() string { return proto.CompactTextString(m) }
func (*RemoteSyncMatchedError) ProtoMessage()    {}
func (*RemoteSyncMatchedError) Descriptor() ([]byte, []int) {
	return fileDescriptor_3688ca0fd170c8f9, []int{3}
}
func (m *RemoteSyncMatchedError) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RemoteSyncMatchedError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RemoteSyncMatchedError.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RemoteSyncMatchedError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteSyncMatchedError.Merge(m, src)
}
func (m *RemoteSyncMatchedError) XXX_Size() int {
	return m.Size()
}
func (m *RemoteSyncMatchedError) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteSyncMatchedError.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteSyncMatchedError proto.InternalMessageInfo

type RetryTaskV2Error struct {
	DomainId             string                  `protobuf:"bytes,1,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	WorkflowExecution    *v1.WorkflowExecution   `protobuf:"bytes,2,opt,name=workflow_execution,json=workflowExecution,proto3" json:"workflow_execution,omitempty"`
	StartEvent           *v11.VersionHistoryItem `protobuf:"bytes,3,opt,name=start_event,json=startEvent,proto3" json:"start_event,omitempty"`
	EndEvent             *v11.VersionHistoryItem `protobuf:"bytes,4,opt,name=end_event,json=endEvent,proto3" json:"end_event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *RetryTaskV2Error) Reset()         { *m = RetryTaskV2Error{} }
func (m *RetryTaskV2Error) String() string { return proto.CompactTextString(m) }
func (*RetryTaskV2Error) ProtoMessage()    {}
func (*RetryTaskV2Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_3688ca0fd170c8f9, []int{4}
}
func (m *RetryTaskV2Error) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RetryTaskV2Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RetryTaskV2Error.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RetryTaskV2Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetryTaskV2Error.Merge(m, src)
}
func (m *RetryTaskV2Error) XXX_Size() int {
	return m.Size()
}
func (m *RetryTaskV2Error) XXX_DiscardUnknown() {
	xxx_messageInfo_RetryTaskV2Error.DiscardUnknown(m)
}

var xxx_messageInfo_RetryTaskV2Error proto.InternalMessageInfo

func (m *RetryTaskV2Error) GetDomainId() string {
	if m != nil {
		return m.DomainId
	}
	return ""
}

func (m *RetryTaskV2Error) GetWorkflowExecution() *v1.WorkflowExecution {
	if m != nil {
		return m.WorkflowExecution
	}
	return nil
}

func (m *RetryTaskV2Error) GetStartEvent() *v11.VersionHistoryItem {
	if m != nil {
		return m.StartEvent
	}
	return nil
}

func (m *RetryTaskV2Error) GetEndEvent() *v11.VersionHistoryItem {
	if m != nil {
		return m.EndEvent
	}
	return nil
}

type ShardOwnershipLostError struct {
	Owner                string   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShardOwnershipLostError) Reset()         { *m = ShardOwnershipLostError{} }
func (m *ShardOwnershipLostError) String() string { return proto.CompactTextString(m) }
func (*ShardOwnershipLostError) ProtoMessage()    {}
func (*ShardOwnershipLostError) Descriptor() ([]byte, []int) {
	return fileDescriptor_3688ca0fd170c8f9, []int{5}
}
func (m *ShardOwnershipLostError) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ShardOwnershipLostError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ShardOwnershipLostError.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ShardOwnershipLostError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShardOwnershipLostError.Merge(m, src)
}
func (m *ShardOwnershipLostError) XXX_Size() int {
	return m.Size()
}
func (m *ShardOwnershipLostError) XXX_DiscardUnknown() {
	xxx_messageInfo_ShardOwnershipLostError.DiscardUnknown(m)
}

var xxx_messageInfo_ShardOwnershipLostError proto.InternalMessageInfo

func (m *ShardOwnershipLostError) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func init() {
	proto.RegisterType((*CurrentBranchChangedError)(nil), "uber.cadence.shared.v1.CurrentBranchChangedError")
	proto.RegisterType((*InternalDataInconsistencyError)(nil), "uber.cadence.shared.v1.InternalDataInconsistencyError")
	proto.RegisterType((*EventAlreadyStartedError)(nil), "uber.cadence.shared.v1.EventAlreadyStartedError")
	proto.RegisterType((*RemoteSyncMatchedError)(nil), "uber.cadence.shared.v1.RemoteSyncMatchedError")
	proto.RegisterType((*RetryTaskV2Error)(nil), "uber.cadence.shared.v1.RetryTaskV2Error")
	proto.RegisterType((*ShardOwnershipLostError)(nil), "uber.cadence.shared.v1.ShardOwnershipLostError")
}

func init() {
	proto.RegisterFile("uber/cadence/shared/v1/error.proto", fileDescriptor_3688ca0fd170c8f9)
}

var fileDescriptor_3688ca0fd170c8f9 = []byte{
	// 440 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xe5, 0xf2, 0x47, 0xcd, 0x96, 0x03, 0x58, 0x55, 0x31, 0x41, 0x8a, 0x22, 0x23, 0xa1,
	0x72, 0xb1, 0x49, 0x11, 0x27, 0x4e, 0x34, 0x04, 0x11, 0x44, 0x85, 0xe4, 0x94, 0x22, 0x71, 0xb1,
	0x36, 0xbb, 0x43, 0xbc, 0x4a, 0x3c, 0x1b, 0xcd, 0x6e, 0x1c, 0xfc, 0x2a, 0x3c, 0x11, 0x47, 0x1e,
	0x01, 0xe5, 0x49, 0xd0, 0xee, 0xba, 0x82, 0xf4, 0xd6, 0x9b, 0x3d, 0xdf, 0xef, 0xfb, 0x3c, 0xfe,
	0x34, 0x2c, 0xdd, 0xcc, 0x81, 0x72, 0xc1, 0x25, 0xa0, 0x80, 0xdc, 0x54, 0x9c, 0x40, 0xe6, 0xcd,
	0x28, 0x07, 0x22, 0x4d, 0xd9, 0x9a, 0xb4, 0xd5, 0xf1, 0x89, 0x63, 0xb2, 0x8e, 0xc9, 0x02, 0x93,
	0x35, 0xa3, 0xfe, 0x70, 0xcf, 0xcb, 0xd7, 0xca, 0x19, 0x85, 0xae, 0x6b, 0x8d, 0xc1, 0xd9, 0x7f,
	0xb6, 0x4f, 0xc8, 0x5a, 0xa1, 0x63, 0x2a, 0x65, 0xac, 0xa6, 0x36, 0x40, 0xe9, 0x05, 0x7b, 0x32,
	0xde, 0x10, 0x01, 0xda, 0x73, 0xe2, 0x28, 0xaa, 0x71, 0xc5, 0x71, 0x01, 0x72, 0xe2, 0x36, 0x88,
	0x5f, 0xb2, 0x63, 0x11, 0xc4, 0x72, 0xee, 0xd5, 0xd2, 0xea, 0x25, 0x60, 0x12, 0x0d, 0xa3, 0xd3,
	0x07, 0x45, 0x2c, 0xfe, 0x37, 0x5e, 0x3a, 0x25, 0x1d, 0xb2, 0xc1, 0x14, 0x2d, 0x10, 0xf2, 0xd5,
	0x3b, 0x6e, 0xf9, 0x14, 0x85, 0x46, 0xa3, 0x8c, 0x05, 0x14, 0xad, 0xcf, 0x4c, 0xfb, 0x2c, 0x99,
	0x34, 0x80, 0xf6, 0xed, 0x8a, 0x80, 0xcb, 0x76, 0x66, 0x39, 0xd9, 0xee, 0x7b, 0x69, 0xc2, 0x4e,
	0x0a, 0xa8, 0xb5, 0x85, 0x59, 0x8b, 0xe2, 0x82, 0x5b, 0x51, 0x5d, 0x2b, 0x3f, 0x0f, 0xd8, 0xc3,
	0x02, 0x2c, 0xb5, 0x97, 0xdc, 0x2c, 0xaf, 0xce, 0xc2, 0x7a, 0x4f, 0x59, 0x4f, 0xea, 0x9a, 0x2b,
	0x2c, 0x95, 0xf4, 0x3b, 0xf5, 0x8a, 0xc3, 0x30, 0x98, 0xca, 0xf8, 0x0b, 0x8b, 0xb7, 0x9a, 0x96,
	0xdf, 0x57, 0x7a, 0x5b, 0xc2, 0x0f, 0x10, 0x1b, 0xab, 0x34, 0x26, 0x07, 0xc3, 0xe8, 0xf4, 0xe8,
	0xec, 0x79, 0xb6, 0x57, 0x2a, 0x5f, 0xab, 0xac, 0x19, 0x65, 0x5f, 0x3b, 0x7c, 0x72, 0x4d, 0x17,
	0x8f, 0xb6, 0x37, 0x47, 0xf1, 0x47, 0x76, 0x64, 0xdc, 0xca, 0x25, 0xb8, 0x9f, 0x48, 0xee, 0xf8,
	0xbc, 0x17, 0x37, 0xf2, 0x5c, 0xd5, 0x2e, 0xf1, 0x0a, 0xc8, 0x28, 0x8d, 0x1f, 0x42, 0xe3, 0x53,
	0x0b, 0x75, 0xc1, 0xbc, 0xdb, 0x37, 0x10, 0xbf, 0x67, 0x3d, 0x40, 0xd9, 0x25, 0xdd, 0xbd, 0x6d,
	0xd2, 0x21, 0xa0, 0xf4, 0x39, 0x69, 0xce, 0x1e, 0xcf, 0x2a, 0x4e, 0xf2, 0xf3, 0x16, 0x81, 0x4c,
	0xa5, 0xd6, 0x9f, 0xb4, 0xb1, 0xa1, 0xa2, 0x63, 0x76, 0x4f, 0xbb, 0x69, 0x57, 0x4f, 0x78, 0x39,
	0x1f, 0xff, 0xda, 0x0d, 0xa2, 0xdf, 0xbb, 0x41, 0xf4, 0x67, 0x37, 0x88, 0xbe, 0xbd, 0x5e, 0x28,
	0x5b, 0x6d, 0xe6, 0x99, 0xd0, 0x75, 0xbe, 0x77, 0x32, 0xd9, 0x02, 0x30, 0xf7, 0x67, 0xf2, 0xef,
	0x36, 0xdf, 0x84, 0xa7, 0x66, 0x34, 0xbf, 0xef, 0x95, 0x57, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x2d, 0xfb, 0xbd, 0x38, 0xc5, 0x02, 0x00, 0x00,
}

func (m *CurrentBranchChangedError) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CurrentBranchChangedError) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CurrentBranchChangedError) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.CurrentBranchToken) > 0 {
		i -= len(m.CurrentBranchToken)
		copy(dAtA[i:], m.CurrentBranchToken)
		i = encodeVarintError(dAtA, i, uint64(len(m.CurrentBranchToken)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *InternalDataInconsistencyError) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InternalDataInconsistencyError) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InternalDataInconsistencyError) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func (m *EventAlreadyStartedError) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventAlreadyStartedError) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventAlreadyStartedError) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func (m *RemoteSyncMatchedError) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RemoteSyncMatchedError) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RemoteSyncMatchedError) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func (m *RetryTaskV2Error) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RetryTaskV2Error) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RetryTaskV2Error) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.EndEvent != nil {
		{
			size, err := m.EndEvent.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintError(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.StartEvent != nil {
		{
			size, err := m.StartEvent.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintError(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.WorkflowExecution != nil {
		{
			size, err := m.WorkflowExecution.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintError(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.DomainId) > 0 {
		i -= len(m.DomainId)
		copy(dAtA[i:], m.DomainId)
		i = encodeVarintError(dAtA, i, uint64(len(m.DomainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ShardOwnershipLostError) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ShardOwnershipLostError) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ShardOwnershipLostError) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintError(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintError(dAtA []byte, offset int, v uint64) int {
	offset -= sovError(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CurrentBranchChangedError) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CurrentBranchToken)
	if l > 0 {
		n += 1 + l + sovError(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *InternalDataInconsistencyError) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *EventAlreadyStartedError) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RemoteSyncMatchedError) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RetryTaskV2Error) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DomainId)
	if l > 0 {
		n += 1 + l + sovError(uint64(l))
	}
	if m.WorkflowExecution != nil {
		l = m.WorkflowExecution.Size()
		n += 1 + l + sovError(uint64(l))
	}
	if m.StartEvent != nil {
		l = m.StartEvent.Size()
		n += 1 + l + sovError(uint64(l))
	}
	if m.EndEvent != nil {
		l = m.EndEvent.Size()
		n += 1 + l + sovError(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ShardOwnershipLostError) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovError(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovError(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozError(x uint64) (n int) {
	return sovError(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CurrentBranchChangedError) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowError
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
			return fmt.Errorf("proto: CurrentBranchChangedError: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CurrentBranchChangedError: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentBranchToken", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowError
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
				return ErrInvalidLengthError
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthError
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CurrentBranchToken = append(m.CurrentBranchToken[:0], dAtA[iNdEx:postIndex]...)
			if m.CurrentBranchToken == nil {
				m.CurrentBranchToken = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipError(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthError
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *InternalDataInconsistencyError) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowError
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
			return fmt.Errorf("proto: InternalDataInconsistencyError: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InternalDataInconsistencyError: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipError(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthError
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EventAlreadyStartedError) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowError
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
			return fmt.Errorf("proto: EventAlreadyStartedError: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventAlreadyStartedError: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipError(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthError
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RemoteSyncMatchedError) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowError
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
			return fmt.Errorf("proto: RemoteSyncMatchedError: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RemoteSyncMatchedError: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipError(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthError
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RetryTaskV2Error) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowError
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
			return fmt.Errorf("proto: RetryTaskV2Error: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RetryTaskV2Error: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DomainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowError
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
				return ErrInvalidLengthError
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthError
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DomainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WorkflowExecution", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowError
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
				return ErrInvalidLengthError
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthError
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.WorkflowExecution == nil {
				m.WorkflowExecution = &v1.WorkflowExecution{}
			}
			if err := m.WorkflowExecution.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartEvent", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowError
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
				return ErrInvalidLengthError
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthError
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.StartEvent == nil {
				m.StartEvent = &v11.VersionHistoryItem{}
			}
			if err := m.StartEvent.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndEvent", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowError
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
				return ErrInvalidLengthError
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthError
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.EndEvent == nil {
				m.EndEvent = &v11.VersionHistoryItem{}
			}
			if err := m.EndEvent.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipError(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthError
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ShardOwnershipLostError) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowError
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
			return fmt.Errorf("proto: ShardOwnershipLostError: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ShardOwnershipLostError: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowError
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
				return ErrInvalidLengthError
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthError
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipError(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthError
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipError(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowError
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
					return 0, ErrIntOverflowError
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
					return 0, ErrIntOverflowError
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
				return 0, ErrInvalidLengthError
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupError
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthError
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthError        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowError          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupError = fmt.Errorf("proto: unexpected end of group")
)
