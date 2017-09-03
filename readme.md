# 简介（Info）

一个简单的go语言下的序列化压测，非常简单的，仅供参考。

Just a simple bench for some serialization in golang.

## 已实现（Exist）

+ google/golang/encoding/json
+ xjdrew/gosproto
+ szyhf/go-sproto
+ google/proto.v3
+ thinylib/msgp
+ json-iterator/go

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

2017.09.03:

At golang v1.9.

```txt
pkg: github.com/szyhf/go-serialize-benchmark/go_sproto
BenchmarkSZYHFEncodeSproto-6         	  200000	      5843 ns/op	    1376 B/op	      36 allocs/op
BenchmarkSZYHFDecodeSproto-6         	  200000	      6961 ns/op	    3136 B/op	      71 allocs/op
BenchmarkSZYHFEncodePackedSproto-6   	  200000	      6918 ns/op	    1856 B/op	      37 allocs/op
BenchmarkSZYHFDecodePackedSproto-6   	  200000	      8326 ns/op	    3808 B/op	      73 allocs/op
PASS
ok  	github.com/szyhf/go-serialize-benchmark/go_sproto	5.949s
=== RUN   TestEncodeSproto
--- PASS: TestEncodeSproto (0.00s)
	xjdrew_sproto_test.go:17: SprotoEncode.Len= 362
	xjdrew_sproto_test.go:24: SprotoEncodePacked.Len= 213
pkg: github.com/szyhf/go-serialize-benchmark/gosproto
BenchmarkEncodeSproto-6         	  300000	      5245 ns/op	    1376 B/op	      36 allocs/op
BenchmarkDecodeSproto-6         	  200000	      8082 ns/op	    3168 B/op	      76 allocs/op
BenchmarkEncodePackedSproto-6   	  300000	      6423 ns/op	    1856 B/op	      37 allocs/op
BenchmarkDecodePackedSproto-6   	  200000	      9271 ns/op	    3840 B/op	      78 allocs/op
PASS
ok  	github.com/szyhf/go-serialize-benchmark/gosproto	8.334s
=== RUN   TestEncodeJSON
--- PASS: TestEncodeJSON (0.00s)
	json_test.go:17: EncodeMsgp.Len= 697
pkg: github.com/szyhf/go-serialize-benchmark/json
BenchmarkEncodeJSON-6   	  200000	      7109 ns/op	    2368 B/op	       9 allocs/op
BenchmarkDecodeJSON-6   	   50000	     23929 ns/op	     544 B/op	      37 allocs/op
PASS
ok  	github.com/szyhf/go-serialize-benchmark/json	2.950s
=== RUN   TestEncodeJsonIter
--- PASS: TestEncodeJsonIter (0.00s)
	json_iterator_test.go:18: EncodeMsgp.Len= 697
pkg: github.com/szyhf/go-serialize-benchmark/json_iterator
BenchmarkEncodeJsonIter-6   	  300000	      3670 ns/op	     720 B/op	       2 allocs/op
BenchmarkDecodeJsonIter-6   	  300000	      4889 ns/op	     208 B/op	      20 allocs/op
PASS
ok  	github.com/szyhf/go-serialize-benchmark/json_iterator	2.686s
=== RUN   TestEncodeMsgp
--- PASS: TestEncodeMsgp (0.00s)
	msgp_test.go:16: EncodeMsgp.Len= 428
pkg: github.com/szyhf/go-serialize-benchmark/msgp
BenchmarkEncodeMsgp-6            2000000               750 ns/op             640 B/op          1 allocs/op
BenchmarkDecodeMsgp-6            1000000              1585 ns/op              56 B/op         12 allocs/op
PASS
ok  	github.com/szyhf/go-serialize-benchmark/msgp	4.700s
=== RUN   TestEncodeProto3
--- PASS: TestEncodeProto3 (0.00s)
	proto3_test.go:17: EncodeProto3.Len= 217
pkg: github.com/szyhf/go-serialize-benchmark/proto3
BenchmarkEncodeProto3-6   	  500000	      2494 ns/op	    1576 B/op	      15 allocs/op
BenchmarkDecodeProto3-6   	  500000	      3992 ns/op	    1160 B/op	      46 allocs/op
PASS
ok  	github.com/szyhf/go-serialize-benchmark/proto3	3.323s
```
