package json

import (
	"math"
	"testing"

	"github.com/json-iterator/go"
)

func TestEncodeJsonIter(t *testing.T) {
	msg := newValMsg()
	var d []byte
	d, err := jsoniter.ConfigFastest.Marshal(msg)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("EncodeMsgp.Len=", len(d))

	dMsg := &ValMsg{}
	err = jsoniter.ConfigFastest.Unmarshal(d, dMsg)
	if err != nil {
		t.Error(err)
		return
	}
}

func BenchmarkEncodeJsonIter(b *testing.B) {
	b.StopTimer()
	msg := newValMsg()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		jsoniter.ConfigFastest.Marshal(msg)
	}
}

func BenchmarkDecodeJsonIter(b *testing.B) {
	b.StopTimer()
	msg := newValMsg()
	var d []byte
	d, _ = jsoniter.ConfigFastest.Marshal(msg)
	nMsg := &ValMsg{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		jsoniter.ConfigFastest.Unmarshal(d, nMsg)
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
