# 简介（Info）

一个简单的go语言下的序列化压测，非常简单的，仅供参考。

Just a simple bench for some serialization in golang.

## 支持（Support）

现在只实现了“gosproto”和“protobuf.v3"的压测，更多的信息可以尝试找找相关目录中的readme.md（我不一定写了=。=）

Now just finished "gosproto" and "google/proto.v3", more info try look at inner readme in directory.

## 使用（Usage）

``` shell
sh ./bench.sh
```

## 我自己机子的结果（My Result）

2017.03.01:

```txt
--- PASS: TestEncodeProto3 (0.00s)
        proto3_test.go:17: EncodeProto3.Len= 217
BenchmarkEncodeProto3-8          1000000              1848 ns/op            1576 B/op         15 allocs/op
BenchmarkDecodeProto3-8           500000              2872 ns/op            1160 B/op         46 allocs/op
PASS
ok      go.szyhf.org/serialize-benchmark/proto3 3.343s
=== RUN   TestEncodeSproto
--- PASS: TestEncodeSproto (0.00s)
        xjdrew_sproto_test.go:17: SprotoEncode.Len= 362
        xjdrew_sproto_test.go:24: SprotoEncodePacked.Len= 213
BenchmarkSZYHFEncodeSproto-8              300000              4099 ns/op            1376 B/op         36 allocs/op
BenchmarkSZYHFDecodeSproto-8              300000              5446 ns/op            3136 B/op         71 allocs/op
BenchmarkSZYHFEncodePackedSproto-8        300000              4995 ns/op            1856 B/op         37 allocs/op
BenchmarkSZYHFDecodePackedSproto-8        200000              6299 ns/op            3808 B/op         73 allocs/op
BenchmarkEncodeSproto-8                   300000              3774 ns/op            1376 B/op         36 allocs/op
BenchmarkDecodeSproto-8                   200000              6252 ns/op            3168 B/op         76 allocs/op
BenchmarkEncodePackedSproto-8             300000              4711 ns/op            1856 B/op         37 allocs/op
BenchmarkDecodePackedSproto-8             200000              7114 ns/op            3840 B/op         78 allocs/op
```

## 计划（Plan）

json

fastjson

msgp