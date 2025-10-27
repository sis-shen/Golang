package main

import "testing"

func TestFib(t *testing.T) {
	fMap := map[int]int{}
	fMap[0] = 0
	fMap[1] = 1
	fMap[2] = 1
	fMap[3] = 2
	fMap[4] = 3
	fMap[5] = 5
	fMap[6] = 8
	fMap[7] = 13
	fMap[8] = 21
	fMap[9] = 34

	for k, v := range fMap {
		fib := Fib(k)
		if fib == v {
			t.Logf("结果正确:n为%d,值为%d", k, fib)
		} else {
			t.Errorf("<UNK>:n<UNK>%d,<UNK>%d", k, fib)
		}
	}
}

func BenchmarkFib(b *testing.B) {
	n := 15
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; i < b.N; i++ {
			Fib(n)
		}
	})
}
