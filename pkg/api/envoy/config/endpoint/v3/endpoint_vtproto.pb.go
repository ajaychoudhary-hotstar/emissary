//go:build vtprotobuf
// +build vtprotobuf

// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// source: envoy/config/endpoint/v3/endpoint.proto

package endpointv3

import (
	protohelpers "github.com/planetscale/vtprotobuf/protohelpers"
	durationpb "github.com/planetscale/vtprotobuf/types/known/durationpb"
	wrapperspb "github.com/planetscale/vtprotobuf/types/known/wrapperspb"
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *ClusterLoadAssignment_Policy_DropOverload) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClusterLoadAssignment_Policy_DropOverload) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *ClusterLoadAssignment_Policy_DropOverload) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.DropPercentage != nil {
		if vtmsg, ok := interface{}(m.DropPercentage).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.DropPercentage)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Category) > 0 {
		i -= len(m.Category)
		copy(dAtA[i:], m.Category)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Category)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ClusterLoadAssignment_Policy) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClusterLoadAssignment_Policy) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *ClusterLoadAssignment_Policy) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.WeightedPriorityHealth {
		i--
		if m.WeightedPriorityHealth {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.EndpointStaleAfter != nil {
		size, err := (*durationpb.Duration)(m.EndpointStaleAfter).MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x22
	}
	if m.OverprovisioningFactor != nil {
		size, err := (*wrapperspb.UInt32Value)(m.OverprovisioningFactor).MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.DropOverloads) > 0 {
		for iNdEx := len(m.DropOverloads) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.DropOverloads[iNdEx].MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x12
		}
	}
	return len(dAtA) - i, nil
}

func (m *ClusterLoadAssignment) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClusterLoadAssignment) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *ClusterLoadAssignment) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.NamedEndpoints) > 0 {
		for k := range m.NamedEndpoints {
			v := m.NamedEndpoints[k]
			baseI := i
			size, err := v.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = protohelpers.EncodeVarint(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.Policy != nil {
		size, err := m.Policy.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Endpoints) > 0 {
		for iNdEx := len(m.Endpoints) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.Endpoints[iNdEx].MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.ClusterName) > 0 {
		i -= len(m.ClusterName)
		copy(dAtA[i:], m.ClusterName)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.ClusterName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ClusterLoadAssignment_Policy_DropOverload) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Category)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.DropPercentage != nil {
		if size, ok := interface{}(m.DropPercentage).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.DropPercentage)
		}
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *ClusterLoadAssignment_Policy) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.DropOverloads) > 0 {
		for _, e := range m.DropOverloads {
			l = e.SizeVT()
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	if m.OverprovisioningFactor != nil {
		l = (*wrapperspb.UInt32Value)(m.OverprovisioningFactor).SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.EndpointStaleAfter != nil {
		l = (*durationpb.Duration)(m.EndpointStaleAfter).SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.WeightedPriorityHealth {
		n += 2
	}
	n += len(m.unknownFields)
	return n
}

func (m *ClusterLoadAssignment) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClusterName)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if len(m.Endpoints) > 0 {
		for _, e := range m.Endpoints {
			l = e.SizeVT()
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	if m.Policy != nil {
		l = m.Policy.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if len(m.NamedEndpoints) > 0 {
		for k, v := range m.NamedEndpoints {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.SizeVT()
			}
			l += 1 + protohelpers.SizeOfVarint(uint64(l))
			mapEntrySize := 1 + len(k) + protohelpers.SizeOfVarint(uint64(len(k))) + l
			n += mapEntrySize + 1 + protohelpers.SizeOfVarint(uint64(mapEntrySize))
		}
	}
	n += len(m.unknownFields)
	return n
}