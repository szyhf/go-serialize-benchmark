# 简介（Info）

一个简单的go语言下的序列化压测，非常简单的，仅供参考。

Just a simple bench for some serialization in golang.

## 已实现（Exist）

+ xjdrew/gosproto
+ szyhf/go-sproto
+ google/proto.v3
+ thinylib/msgp

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

2017.03.05:

```txt
?   	github.com/szyhf/go-serialize-benchmark/fastjson	[no test files]
=== RUN   TestEncodeJSON
--- PASS: TestEncodeJSON (0.00s)
	json_test.go:17: EncodeMsgp.Len= 697
BenchmarkEncodeJSON-6   	  200000	      8874 ns/op	    2376 B/op	      10 allocs/op
BenchmarkDecodeJSON-6   	10000000	       173 ns/op	     320 B/op	       2 allocs/op
PASS
ok  	github.com/szyhf/go-serialize-benchmark/json	3.828s
=== RUN   TestEncodeMsgp
--- PASS: TestEncodeMsgp (0.00s)
	msgp_test.go:16: EncodeMsgp.Len= 428
BenchmarkEncodeMsgp-6   	 2000000	       735 ns/op	     640 B/op	       1 allocs/op
BenchmarkDecodeMsgp-6   	200000000	         8.16 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/szyhf/go-serialize-benchmark/msgp	4.671s
=== RUN   TestEncodeProto3
--- PASS: TestEncodeProto3 (0.00s)
	proto3_test.go:17: EncodeProto3.Len= 217
BenchmarkEncodeProto3-6   	  500000	      2441 ns/op	    1576 B/op	      15 allocs/op
BenchmarkDecodeProto3-6   	  300000	      3670 ns/op	    1160 B/op	      46 allocs/op
PASS
ok  	github.com/szyhf/go-serialize-benchmark/proto3	2.405s
=== RUN   TestEncodeSproto
--- PASS: TestEncodeSproto (0.00s)
	xjdrew_sproto_test.go:17: SprotoEncode.Len= 362
	xjdrew_sproto_test.go:24: SprotoEncodePacked.Len= 213
BenchmarkSZYHFEncodeSproto-6         	  300000	      5302 ns/op	    1376 B/op	      36 allocs/op
BenchmarkSZYHFDecodeSproto-6         	  200000	      7471 ns/op	    3136 B/op	      71 allocs/op
BenchmarkSZYHFEncodePackedSproto-6   	  300000	      6474 ns/op	    1856 B/op	      37 allocs/op
BenchmarkSZYHFDecodePackedSproto-6   	  200000	      8206 ns/op	    3808 B/op	      73 allocs/op
BenchmarkEncodeSproto-6              	  300000	      4857 ns/op	    1376 B/op	      36 allocs/op
BenchmarkDecodeSproto-6              	  200000	      8444 ns/op	    3168 B/op	      76 allocs/op
BenchmarkEncodePackedSproto-6        	  300000	      6110 ns/op	    1856 B/op	      37 allocs/op
BenchmarkDecodePackedSproto-6        	  200000	      9960 ns/op	    3840 B/op	      78 allocs/op
PASS
ok  	github.com/szyhf/go-serialize-benchmark/sproto	14.238s
```

## 计划（Plan）

json

fastjson
