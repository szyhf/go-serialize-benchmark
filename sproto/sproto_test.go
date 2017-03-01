package sproto

import (
	"math"
	"testing"

	sproto "github.com/davyxu/gosproto"
)

func BenchmarkSProtoEncodeVal(b *testing.B) {
	b.StopTimer()
	msg := newValMsg()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sproto.Encode(msg)
	}
}

func BenchmarkSProtoEncodePackedVal(b *testing.B) {
	b.StopTimer()
	msg := newValMsg()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sproto.EncodePacked(msg)
	}
}

func BenchmarkSProtoDecodeVal(b *testing.B) {
	b.StopTimer()
	msgData, _ := sproto.Encode(newValMsg())
	var msg = ValMSG{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sproto.Decode(msgData, &msg)
	}
}

func BenchmarkSProtoDecodePackedVal(b *testing.B) {
	b.StopTimer()
	msgData, _ := sproto.EncodePacked(newValMsg())
	var msg = ValMSG{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sproto.DecodePacked(msgData, &msg)
	}
}

func newValMsg() ValMSG {
	return ValMSG{
		Int:         math.MaxInt64,
		IntNeg:      math.MinInt32 + 1,
		String:      "Hello",
		Bool:        true,
		ByteSlice:   []byte("World"),
		BoolSlice:   []bool{true, true, false, true, false},
		IntSlice:    []int{123, 321, 1234567},
		StringSlice: []string{"FOO", "BAR"},
		Struct: &HoldValMSG{
			Int:         1,
			String:      "Hello",
			Bool:        true,
			ByteSlice:   []byte("World"),
			BoolSlice:   []bool{true, true, false, true, false},
			IntSlice:    []int{123, 321, 1234567},
			StringSlice: []string{"FOO", "BAR"},
		},
		StructSlice: []*HoldValMSG{
			&HoldValMSG{
				Int:         2,
				String:      "Hello2",
				Bool:        true,
				ByteSlice:   []byte("World2"),
				BoolSlice:   []bool{true, true, false, true, false},
				IntSlice:    []int{123, 321, 1234567},
				StringSlice: []string{"FOO2", "BAR2"},
			},
			&HoldValMSG{
				Int:         3,
				String:      "Hello3",
				Bool:        true,
				ByteSlice:   []byte("World3"),
				BoolSlice:   []bool{true, true, false, true, false},
				IntSlice:    []int{123, 321, 1234567},
				StringSlice: []string{"FOO3", "BAR3"},
			},
		},
	}
}
