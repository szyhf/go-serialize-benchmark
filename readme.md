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

2018.01.19:

At golang v1.9.2.

```txt
goos: darwin
goarch: amd64
pkg: github.com/szyhf/go-serialize-benchmark/go_sproto
BenchmarkSZYHFEncodeSproto-8              300000              4104 ns/op            1376 B/op         36 allocs/op
BenchmarkSZYHFDecodeSproto-8              300000              5018 ns/op            3136 B/op         71 allocs/op
BenchmarkSZYHFEncodePackedSproto-8        300000              4865 ns/op            1856 B/op         37 allocs/op
BenchmarkSZYHFDecodePackedSproto-8        300000              5889 ns/op            3808 B/op         73 allocs/op
PASS
ok      github.com/szyhf/go-serialize-benchmark/go_sproto       6.183s
=== RUN   TestEncodeGob
--- PASS: TestEncodeGob (0.00s)
        gob_test.go:30: EncodeMsgp.Len= 571
goos: darwin
goarch: amd64
pkg: github.com/szyhf/go-serialize-benchmark/gob
BenchmarkEncodeGob-8      500000              3684 ns/op             384 B/op         12 allocs/op
BenchmarkDecodeGob-8      300000              4809 ns/op            1048 B/op         42 allocs/op
PASS
ok      github.com/szyhf/go-serialize-benchmark/gob     4.651s
=== RUN   TestEncodeSproto
--- PASS: TestEncodeSproto (0.00s)
        xjdrew_sproto_test.go:17: SprotoEncode.Len= 362
        xjdrew_sproto_test.go:24: SprotoEncodePacked.Len= 213
goos: darwin
goarch: amd64
pkg: github.com/szyhf/go-serialize-benchmark/gosproto
BenchmarkEncodeSproto-8                   300000              3707 ns/op            1376 B/op         36 allocs/op
BenchmarkDecodeSproto-8                   300000              6112 ns/op            3168 B/op         76 allocs/op
BenchmarkEncodePackedSproto-8             300000              4641 ns/op            1856 B/op         37 allocs/op
BenchmarkDecodePackedSproto-8             200000              6845 ns/op            3840 B/op         78 allocs/op
PASS
ok      github.com/szyhf/go-serialize-benchmark/gosproto        5.940s
=== RUN   TestEncodeJSON
--- PASS: TestEncodeJSON (0.00s)
        json_test.go:17: EncodeMsgp.Len= 697
goos: darwin
goarch: amd64
pkg: github.com/szyhf/go-serialize-benchmark/json
BenchmarkEncodeJSON-8             300000              4977 ns/op            2368 B/op          9 allocs/op
BenchmarkDecodeJSON-8             100000             20016 ns/op             544 B/op         37 allocs/op
PASS
ok      github.com/szyhf/go-serialize-benchmark/json    3.758s
=== RUN   TestEncodeJsonIter
--- PASS: TestEncodeJsonIter (0.00s)
        json_iterator_test.go:18: EncodeMsgp.Len= 697
goos: darwin
goarch: amd64
pkg: github.com/szyhf/go-serialize-benchmark/json_iterator
BenchmarkEncodeJsonIter-8         500000              2700 ns/op             720 B/op          2 allocs/op
BenchmarkDecodeJsonIter-8         300000              4344 ns/op             256 B/op         24 allocs/op
PASS
ok      github.com/szyhf/go-serialize-benchmark/json_iterator   2.739s
=== RUN   TestEncodeMsgp
--- PASS: TestEncodeMsgp (0.00s)
        msgp_test.go:16: EncodeMsgp.Len= 428
goos: darwin
goarch: amd64
pkg: github.com/szyhf/go-serialize-benchmark/msgp
BenchmarkEncodeMsgp-8            3000000               506 ns/op             640 B/op          1 allocs/op
BenchmarkDecodeMsgp-8            1000000              1290 ns/op              56 B/op         12 allocs/op
PASS
ok      github.com/szyhf/go-serialize-benchmark/msgp    3.357s
=== RUN   TestEncodeProto3
--- PASS: TestEncodeProto3 (0.00s)
        proto3_test.go:17: EncodeProto3.Len= 217
goos: darwin
goarch: amd64
pkg: github.com/szyhf/go-serialize-benchmark/proto3
BenchmarkEncodeProto3-8          1000000              1726 ns/op            1576 B/op         15 allocs/op
BenchmarkDecodeProto3-8           500000              2646 ns/op            1160 B/op         46 allocs/op
PASS
ok      github.com/szyhf/go-serialize-benchmark/proto3  3.108s
```