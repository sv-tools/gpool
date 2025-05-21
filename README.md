# gpool

Generic wrapper for sync.Pool in Go.

The repository is archived because the feature is fully implemented and will be replaced by sync v2 package in the future.

## Usage

```shell
go get github.com/sv-tools/gpool
```

`int64` type
```go
p := Pool[int64]{New: func() int64 { return 42 }}
x := p.Get()
defer p.Put(x)
fmt.Printf("x = (%T) %d", x, x)
// Output: x = (int64) 42
```

`string` type
```go
p := Pool[string]{New: func() string { return "foo" }}
x := p.Get()
defer p.Put(x)
fmt.Printf("x = (%T) %s", x, x)
// Output: x = (string) foo
```

## Benchmarks

```shell
% go test -bench=. -benchmem ./...
goos: darwin
goarch: arm64
pkg: github.com/sv-tools/gpool
BenchmarkSyncPool-8     699275571                1.614 ns/op           0 B/op          0 allocs/op
BenchmarkPool-8         647708158                1.732 ns/op           0 B/op          0 allocs/op
```
