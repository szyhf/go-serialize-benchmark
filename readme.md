# 简介（Info）

一个简单的go语言下的序列化压测，非常简单的，仅供参考。

Just a simple bench for some serialization in golang.

## 已实现（Exist）

+ google/golang/encoding/gob
+ google/golang/encoding/json
+ xjdrew/gosproto
+ szyhf/go-sproto
+ google/proto.v3
+ thinylib/msgp
+ vmihailenco/msgpack
+ json-iterator/go
+ mailru/easyjson

更多的信息可以尝试找找相关目录中的readme.md（我不一定写了=。=）

More info try look at inner readme in directory.

## 使用（Usage）

``` shell
sh ./bench.sh
```

## 使用的结构体（Struct）

``` golang
type ValMsg struct {
	Int         int
	Str         string
	Bool        bool
	Struct      *HoldValMsg
	ByteSlice   []byte
	BoolSlice   []bool
	IntSlice    []int
	StringSlice []string
	StructSlice []*HoldValMsg
	IntNeg      int
}

type HoldValMsg struct {
	Int         int
	Str         string
	Bool        bool
	ByteSlice   []byte
	BoolSlice   []bool
	IntSlice    []int
	StringSlice []string
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
```

## 我自己机子的结果（My Result）

2019.08.13:

go version go1.12.6 darwin/amd64

```txt
=== RUN   TestEncodeEasyJSON
--- PASS: TestEncodeEasyJSON (0.00s)
    easyjson_test.go:16: EncodeEasyJSON.Len= 697
goos: darwin
goarch: amd64
pkg: github.com/szyhf/go-serialize-benchmark/internal/easyjson
BenchmarkEncodeEasyJSON-8   	 1000000	      1873 ns/op	    1312 B/op	       5 allocs/op
BenchmarkDecodeEasyJSON-8   	  300000	      4447 ns/op	     720 B/op	      24 allocs/op
PASS
ok  	github.com/szyhf/go-serialize-benchmark/internal/easyjson	3.291s
goos: darwin
goarch: amd64
pkg: github.com/szyhf/go-serialize-benchmark/internal/go_sproto
BenchmarkSZYHFEncodeSproto-8         	  500000	      3648 ns/op	    1376 B/op	      36 allocs/op
BenchmarkSZYHFDecodeSproto-8         	  300000	      5423 ns/op	    3136 B/op	      71 allocs/op
BenchmarkSZYHFEncodePackedSproto-8   	  300000	      4523 ns/op	    1856 B/op	      37 allocs/op
BenchmarkSZYHFDecodePackedSproto-8   	  200000	      6227 ns/op	    3808 B/op	      73 allocs/op
PASS
ok  	github.com/szyhf/go-serialize-benchmark/internal/go_sproto	6.269s
=== RUN   TestEncodeGob
--- PASS: TestEncodeGob (0.00s)
    gob_test.go:30: EncodeGob.Len= 571
goos: darwin
goarch: amd64
pkg: github.com/szyhf/go-serialize-benchmark/internal/gob
BenchmarkEncodeGob-8   	  500000	      3587 ns/op	     384 B/op	      12 allocs/op
BenchmarkDecodeGob-8   	  300000	      4666 ns/op	    1048 B/op	      42 allocs/op
PASS
ok  	github.com/szyhf/go-serialize-benchmark/internal/gob	4.513s
=== RUN   TestEncodeSproto
--- PASS: TestEncodeSproto (0.00s)
    xjdrew_sproto_test.go:17: SprotoEncode.Len= 362
    xjdrew_sproto_test.go:24: SprotoEncodePacked.Len= 213
goos: darwin
goarch: amd64
pkg: github.com/szyhf/go-serialize-benchmark/internal/gosproto
BenchmarkEncodeSproto-8         	  500000	      3408 ns/op	    1376 B/op	      36 allocs/op
BenchmarkDecodeSproto-8         	  200000	      6198 ns/op	    3168 B/op	      76 allocs/op
BenchmarkEncodePackedSproto-8   	  300000	      4270 ns/op	    1856 B/op	      37 allocs/op
BenchmarkDecodePackedSproto-8   	  200000	      7055 ns/op	    3840 B/op	      78 allocs/op
PASS
ok  	github.com/szyhf/go-serialize-benchmark/internal/gosproto	5.861s
=== RUN   TestEncodeJSON
--- PASS: TestEncodeJSON (0.00s)
    json_test.go:17: EncodeJSON.Len= 697
goos: darwin
goarch: amd64
pkg: github.com/szyhf/go-serialize-benchmark/internal/json
BenchmarkEncodeJSON-8   	  500000	      3209 ns/op	     704 B/op	       1 allocs/op
BenchmarkDecodeJSON-8   	  100000	     16197 ns/op	     400 B/op	      34 allocs/op
PASS
ok  	github.com/szyhf/go-serialize-benchmark/internal/json	3.430s
=== RUN   TestEncodeJSONIter
--- PASS: TestEncodeJSONIter (0.00s)
    json_iterator_test.go:18: EncodeJSONIter.Len= 697
goos: darwin
goarch: amd64
pkg: github.com/szyhf/go-serialize-benchmark/internal/json_iterator
BenchmarkEncodeJSONIter-8   	  500000	      2964 ns/op	     744 B/op	       6 allocs/op
BenchmarkDecodeJSONIter-8   	  300000	      4245 ns/op	     128 B/op	      20 allocs/op
PASS
ok  	github.com/szyhf/go-serialize-benchmark/internal/json_iterator	2.840s
=== RUN   TestEncodeProto3
--- PASS: TestEncodeProto3 (0.00s)
    proto3_test.go:17: EncodeProto3.Len= 217
goos: darwin
goarch: amd64
pkg: github.com/szyhf/go-serialize-benchmark/internal/proto3
BenchmarkEncodeProto3-8   	 1000000	      1137 ns/op	     256 B/op	       2 allocs/op
BenchmarkDecodeProto3-8   	  500000	      2552 ns/op	     984 B/op	      46 allocs/op
PASS
ok  	github.com/szyhf/go-serialize-benchmark/internal/proto3	2.461s
=== RUN   TestEncodeTinylibMsgp
--- PASS: TestEncodeTinylibMsgp (0.00s)
    tinylib_msgp_test.go:16: EncodeTinylibMsgp.Len= 428
goos: darwin
goarch: amd64
pkg: github.com/szyhf/go-serialize-benchmark/internal/tinylib_msgp
BenchmarkEncodeTinylibMsgp-8   	 3000000	       490 ns/op	     640 B/op	       1 allocs/op
BenchmarkDecodeTinylibMsgp-8   	 1000000	      1111 ns/op	      56 B/op	      12 allocs/op
PASS
ok  	github.com/szyhf/go-serialize-benchmark/internal/tinylib_msgp	3.097s
=== RUN   TestEncodeVmihailencoMsgp
--- PASS: TestEncodeVmihailencoMsgp (0.00s)
    vmihailenco_msgp_test.go:18: EncodeVmihailencoMsgp.Len= 528
goos: darwin
goarch: amd64
pkg: github.com/szyhf/go-serialize-benchmark/internal/vmihailenco_msgp
BenchmarkEncodeVmihailencoMsgp-8   	  300000	      5346 ns/op	    1344 B/op	      11 allocs/op
BenchmarkDecodeVmihailencoMsgp-8   	  200000	      8754 ns/op	     880 B/op	      55 allocs/op
PASS
ok  	github.com/szyhf/go-serialize-benchmark/internal/vmihailenco_msgp	3.510
```