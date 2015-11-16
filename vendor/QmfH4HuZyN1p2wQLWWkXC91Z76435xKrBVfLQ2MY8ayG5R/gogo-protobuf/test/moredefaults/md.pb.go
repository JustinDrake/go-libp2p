// Code generated by protoc-gen-gogo.
// source: md.proto
// DO NOT EDIT!

/*
	Package moredefaults is a generated protocol buffer package.

	It is generated from these files:
		md.proto

	It has these top-level messages:
		MoreDefaultsB
		MoreDefaultsA
*/
package moredefaults

import proto "QmfH4HuZyN1p2wQLWWkXC91Z76435xKrBVfLQ2MY8ayG5R/gogo-protobuf/proto"
import fmt "fmt"
import math "math"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto"
import test "QmfH4HuZyN1p2wQLWWkXC91Z76435xKrBVfLQ2MY8ayG5R/gogo-protobuf/test/example"

import bytes "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MoreDefaultsB struct {
	Field1           *string `protobuf:"bytes,1,opt,name=Field1" json:"Field1,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MoreDefaultsB) Reset()         { *m = MoreDefaultsB{} }
func (m *MoreDefaultsB) String() string { return proto.CompactTextString(m) }
func (*MoreDefaultsB) ProtoMessage()    {}

func (m *MoreDefaultsB) GetField1() string {
	if m != nil && m.Field1 != nil {
		return *m.Field1
	}
	return ""
}

type MoreDefaultsA struct {
	Field1           *int64         `protobuf:"varint,1,opt,name=Field1,def=1234" json:"Field1,omitempty"`
	Field2           int64          `protobuf:"varint,2,opt,name=Field2" json:"Field2"`
	B1               *MoreDefaultsB `protobuf:"bytes,3,opt,name=B1" json:"B1,omitempty"`
	B2               MoreDefaultsB  `protobuf:"bytes,4,opt,name=B2" json:"B2"`
	A1               *test.A        `protobuf:"bytes,5,opt,name=A1" json:"A1,omitempty"`
	A2               test.A         `protobuf:"bytes,6,opt,name=A2" json:"A2"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *MoreDefaultsA) Reset()         { *m = MoreDefaultsA{} }
func (m *MoreDefaultsA) String() string { return proto.CompactTextString(m) }
func (*MoreDefaultsA) ProtoMessage()    {}

const Default_MoreDefaultsA_Field1 int64 = 1234

func (m *MoreDefaultsA) GetField1() int64 {
	if m != nil && m.Field1 != nil {
		return *m.Field1
	}
	return Default_MoreDefaultsA_Field1
}

func (m *MoreDefaultsA) GetField2() int64 {
	if m != nil {
		return m.Field2
	}
	return 0
}

func (m *MoreDefaultsA) GetB1() *MoreDefaultsB {
	if m != nil {
		return m.B1
	}
	return nil
}

func (m *MoreDefaultsA) GetB2() MoreDefaultsB {
	if m != nil {
		return m.B2
	}
	return MoreDefaultsB{}
}

func (m *MoreDefaultsA) GetA1() *test.A {
	if m != nil {
		return m.A1
	}
	return nil
}

func (m *MoreDefaultsA) GetA2() test.A {
	if m != nil {
		return m.A2
	}
	return test.A{}
}

func (this *MoreDefaultsB) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*MoreDefaultsB)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Field1 != nil && that1.Field1 != nil {
		if *this.Field1 != *that1.Field1 {
			return false
		}
	} else if this.Field1 != nil {
		return false
	} else if that1.Field1 != nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MoreDefaultsA) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*MoreDefaultsA)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Field1 != nil && that1.Field1 != nil {
		if *this.Field1 != *that1.Field1 {
			return false
		}
	} else if this.Field1 != nil {
		return false
	} else if that1.Field1 != nil {
		return false
	}
	if this.Field2 != that1.Field2 {
		return false
	}
	if !this.B1.Equal(that1.B1) {
		return false
	}
	if !this.B2.Equal(&that1.B2) {
		return false
	}
	if !this.A1.Equal(that1.A1) {
		return false
	}
	if !this.A2.Equal(&that1.A2) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func NewPopulatedMoreDefaultsB(r randyMd, easy bool) *MoreDefaultsB {
	this := &MoreDefaultsB{}
	if r.Intn(10) != 0 {
		v1 := randStringMd(r)
		this.Field1 = &v1
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedMd(r, 2)
	}
	return this
}

func NewPopulatedMoreDefaultsA(r randyMd, easy bool) *MoreDefaultsA {
	this := &MoreDefaultsA{}
	if r.Intn(10) != 0 {
		v2 := int64(r.Int63())
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		this.Field1 = &v2
	}
	this.Field2 = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Field2 *= -1
	}
	if r.Intn(10) != 0 {
		this.B1 = NewPopulatedMoreDefaultsB(r, easy)
	}
	v3 := NewPopulatedMoreDefaultsB(r, easy)
	this.B2 = *v3
	if r.Intn(10) != 0 {
		this.A1 = test.NewPopulatedA(r, easy)
	}
	v4 := test.NewPopulatedA(r, easy)
	this.A2 = *v4
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedMd(r, 7)
	}
	return this
}

type randyMd interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneMd(r randyMd) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringMd(r randyMd) string {
	v5 := r.Intn(100)
	tmps := make([]rune, v5)
	for i := 0; i < v5; i++ {
		tmps[i] = randUTF8RuneMd(r)
	}
	return string(tmps)
}
func randUnrecognizedMd(r randyMd, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldMd(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldMd(data []byte, r randyMd, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateMd(data, uint64(key))
		v6 := r.Int63()
		if r.Intn(2) == 0 {
			v6 *= -1
		}
		data = encodeVarintPopulateMd(data, uint64(v6))
	case 1:
		data = encodeVarintPopulateMd(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateMd(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateMd(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateMd(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateMd(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
