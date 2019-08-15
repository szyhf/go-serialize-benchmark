package gob

import (
	"bytes"
	"encoding/gob"
	"math"
	"testing"
)

var buf = bytes.NewBuffer(make([]byte, 1024))
var encoder *gob.Encoder
var decoder *gob.Decoder

func init() {
	gob.Register(new(ValMsg))
	gob.Register(new(HoldValMsg))

	encoder = gob.NewEncoder(buf)
	decoder = gob.NewDecoder(buf)
}

func TestEncodeGob(t *testing.T) {
	buf.Reset()
	msg := newValMsg()
	err := encoder.Encode(msg)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("EncodeGob.Len=", buf.Len())

	dMsg := &ValMsg{}
	err = decoder.Decode(dMsg)
	if err != nil {
		t.Error(err)
		return
	}
}

func BenchmarkEncodeGob(b *testing.B) {
	b.StopTimer()
	msg := newValMsg()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		encoder.Encode(msg)
		buf.Reset()
	}
}

func BenchmarkDecodeGob(b *testing.B) {
	b.StopTimer()
	msg := newValMsg()
	for i := 0; i < b.N; i++ {
		encoder.Encode(msg)
	}
	nMsg := &ValMsg{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		decoder.Decode(nMsg)
	}
}

func newValMsg() *ValMsg {
	return &ValMsg{
		Int:         math.MaxInt64,
		IntNeg:      math.MinInt32 + 1,
		Str:         "Hello",
		Bool:        true,
		ByteSlice:   []byte("World"),
		BoolSlice:   []bool{true, true, false, true, false},
		IntSlice:    []int{123, 321, 1234567},
		StringSlice: []string{"FOO", "BAR"},
		Struct: &HoldValMsg{
			Int:         1,
			Str:         "Hello",
			Bool:        true,
			ByteSlice:   []byte("World"),
			BoolSlice:   []bool{true, true, false, true, false},
			IntSlice:    []int{123, 321, 1234567},
			StringSlice: []string{"FOO", "BAR"},
		},
		StructSlice: []*HoldValMsg{
			&HoldValMsg{
				Int:         2,
				Str:         "Hello2",
				Bool:        true,
				ByteSlice:   []byte("World2"),
				BoolSlice:   []bool{true, true, false, true, false},
				IntSlice:    []int{123, 321, 1234567},
				StringSlice: []string{"FOO2", "BAR2"},
			},
			&HoldValMsg{
				Int:         3,
				Str:         "Hello3",
				Bool:        true,
				ByteSlice:   []byte("World3"),
				BoolSlice:   []bool{true, true, false, true, false},
				IntSlice:    []int{123, 321, 1234567},
				StringSlice: []string{"FOO3", "BAR3"},
			},
		},
	}
}
