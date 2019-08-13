package msgp

import (
	"math"
	"testing"
)

func TestEncodeMsgp(t *testing.T) {
	msg := newValMsg()
	var d []byte
	d, err := msg.MarshalMsg(d)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("EncodeMsgp.Len=", len(d))

	dMsg := &ValMsg{}
	d, err = dMsg.UnmarshalMsg(d)
	if err != nil {
		t.Error(err)
		return
	}
}

func BenchmarkEncodeMsgp(b *testing.B) {
	b.StopTimer()
	msg := newValMsg()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var d []byte
		msg.MarshalMsg(d)
	}
}

func BenchmarkDecodeMsgp(b *testing.B) {
	b.StopTimer()
	msg := newValMsg()
	var d []byte
	d, _ = msg.MarshalMsg(d)
	nMsg := &ValMsg{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		nMsg.UnmarshalMsg(d)
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
