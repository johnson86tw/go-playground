package main

import "testing"

// go test -v -bench=. -benchmem

// 印出來的東西：
// 1. 每秒跑幾次
// 2. 每一次需要多久
// 3. 每次一個記憶體空間為多少 Byte
// 4. 每次執行要搭配多少記憶體空間

func BenchmarkAddByShareMemory(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addByShareMemory(100)
	}
}

func BenchmarkAddByShareCommunicate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addByShareCommunicate(100)
	}
}
