package go_sproto

import (
	"math"
	"testing"

	sproto "github.com/szyhf/go-sproto"
)

func BenchmarkSZYHFEncodeSproto(b *testing.B) {
	b.StopTimer()
	msg := newValMsg()
	_, err := sproto.Encode(msg)
	if err != nil {
		b.Error("sproto.Encode failed:", err)
		return
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sproto.Encode(msg)
	}
}

func BenchmarkSZYHFDecodeSproto(b *testing.B) {
	b.StopTimer()
	msg := newValMsg()
	d, err := sproto.Encode(msg)
	if err != nil {
		b.Error("sproto.Encode failed:", err)
		return
	}
	nMsg := new(ValMsg)
	_, err = sproto.Decode(d, nMsg)
	if err != nil {
		b.Error("sproto.Decode failed:", err)
		return
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sproto.Decode(d, nMsg)
	}
}

func BenchmarkSZYHFEncodePackedSproto(b *testing.B) {
	b.StopTimer()
	msg := newValMsg()
	_, err := sproto.EncodePacked(msg)
	if err != nil {
		b.Error("sproto.EncodePacked failed:", err)
		return
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sproto.EncodePacked(msg)
	}
}

func BenchmarkSZYHFDecodePackedSproto(b *testing.B) {
	b.StopTimer()
	msg := newValMsg()
	d, err := sproto.EncodePacked(msg)
	if err != nil {
		b.Error("sproto.EncodePacked failed:", err)
		return
	}
	nMsg := new(ValMsg)
	err = sproto.DecodePacked(d, nMsg)
	if err != nil {
		b.Error("sproto.DecodePacked failed:", err)
		return
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sproto.DecodePacked(d, nMsg)
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
