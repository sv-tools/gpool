package gpool

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
)

func TestPool(t *testing.T) {
	p := Pool[int64]{}
	x := p.Get()
	if x != 0 {
		t.Fatalf("expected x = 0, but got %d", x)
	}
	p.New = func() int64 {
		// should return only a number which is greater than zero
		v := rand.Int63()
		for v <= 0 {
			v = rand.Int63()
		}
		return v
	}
	x = p.Get()
	if x == 0 {
		t.Fatal("expected x != 0, but got 0")
	}
	p.Put(-1)
	// getting until the pool returns -1 or fails by timeout
	for x != -1 {
		x = p.Get()
	}
}

var benchmarkSyncPoolResult any

func BenchmarkSyncPool(b *testing.B) {
	p := sync.Pool{New: func() any { return 42 }}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		var res any
		for pb.Next() {
			res = p.Get()
			p.Put(res)
		}
		benchmarkSyncPoolResult = res
	})
}

var benchmarkPoolResult int64

func BenchmarkPool(b *testing.B) {
	p := Pool[int64]{New: func() int64 { return 42 }}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		var res int64
		for pb.Next() {
			res = p.Get()
			p.Put(res)
		}
		benchmarkPoolResult = res
	})
}

func ExamplePool_int64() {
	p := Pool[int64]{New: func() int64 { return 42 }}
	x := p.Get()
	defer p.Put(x)
	fmt.Printf("x = (%T) %d", x, x)
	// Output: x = (int64) 42
}

func ExamplePool_string() {
	p := Pool[string]{New: func() string { return "foo" }}
	x := p.Get()
	defer p.Put(x)
	fmt.Printf("x = (%T) %s", x, x)
	// Output: x = (string) foo
}
