package sproto

import (
	"math"
	"testing"

	sproto "github.com/szyhf/gosproto"
)

func TestEncodeSproto(t *testing.T) {
	msg := newMsg()
	d, err := sproto.Encode(msg)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("SprotoEncode.Len=", len(d))

	d, err = sproto.EncodePacked(msg)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("SprotoEncodePacked.Len=", len(d))
}

func BenchmarkEncodeSproto(b *testing.B) {
	b.StopTimer()
	msg := newMsg()
	_, err := sproto.Encode(msg)
	if err != nil {
		b.Errorf("sproto.Encode failed:", err)
		return
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sproto.Encode(msg)
	}
}

func BenchmarkDecodeSproto(b *testing.B) {
	b.StopTimer()
	msg := newMsg()
	d, err := sproto.Encode(msg)
	if err != nil {
		b.Errorf("sproto.Encode failed:", err)
		return
	}
	nMsg := new(PtrMsg)
	_, err = sproto.Decode(d, nMsg)
	if err != nil {
		b.Errorf("sproto.Decode failed:", err)
		return
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sproto.Decode(d, nMsg)
	}
}

func BenchmarkEncodePackedSproto(b *testing.B) {
	b.StopTimer()
	msg := newMsg()
	_, err := sproto.EncodePacked(msg)
	if err != nil {
		b.Errorf("sproto.EncodePacked failed:", err)
		return
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sproto.EncodePacked(msg)
	}
}

func BenchmarkDecodePackedSproto(b *testing.B) {
	b.StopTimer()
	msg := newMsg()
	d, err := sproto.EncodePacked(msg)
	if err != nil {
		b.Errorf("sproto.EncodePacked failed:", err)
		return
	}
	nMsg := new(PtrMsg)
	err = sproto.DecodePacked(d, nMsg)
	if err != nil {
		b.Errorf("sproto.DecodePacked failed:", err)
		return
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sproto.DecodePacked(d, nMsg)
	}
}

func newMsg() *PtrMsg {
	return &PtrMsg{
		Int:         sproto.Int(math.MaxInt64),
		IntNeg:      sproto.Int(math.MinInt32 + 1),
		String:      sproto.String("Hello"),
		Bool:        sproto.Bool(true),
		ByteSlice:   []byte("World"),
		BoolSlice:   []bool{true, true, false, true, false},
		IntSlice:    []int{123, 321, 1234567},
		StringSlice: []string{"FOO", "BAR"},
		Struct: &HoldPtrMsg{
			Int:         sproto.Int(1),
			String:      sproto.String("Hello"),
			Bool:        sproto.Bool(true),
			ByteSlice:   []byte("World"),
			BoolSlice:   []bool{true, true, false, true, false},
			IntSlice:    []int{123, 321, 1234567},
			StringSlice: []string{"FOO", "BAR"},
		},
		StructSlice: []*HoldPtrMsg{
			&HoldPtrMsg{
				Int:         sproto.Int(2),
				String:      sproto.String("Hello2"),
				Bool:        sproto.Bool(true),
				ByteSlice:   []byte("World2"),
				BoolSlice:   []bool{true, true, false, true, false},
				IntSlice:    []int{123, 321, 1234567},
				StringSlice: []string{"FOO2", "BAR2"},
			},
			&HoldPtrMsg{
				Int:         sproto.Int(3),
				String:      sproto.String("Hello3"),
				Bool:        sproto.Bool(true),
				ByteSlice:   []byte("World3"),
				BoolSlice:   []bool{true, true, false, true, false},
				IntSlice:    []int{123, 321, 1234567},
				StringSlice: []string{"FOO3", "BAR3"},
			},
		},
	}
}
