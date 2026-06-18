// Benchmark untuk string concatenation, slice pre-allocation, dan map vs slice lookup.
// Jalankan: go test -bench=. -benchmem ./performance/

package main

import (
	"fmt"
	"strings"
	"testing"
)

// ---------------------------------------------------------------
// Benchmark 1: String Concatenation
// + operator vs strings.Builder vs fmt.Sprintf
// ---------------------------------------------------------------

// concatPlus menggabungkan n string menggunakan operator +
func concatPlus(n int, s string) string {
	r := ""
	for range n {
		r += s
	}
	return r
}

// concatBuilder menggabungkan n string menggunakan strings.Builder
func concatBuilder(n int, s string) string {
	var b strings.Builder
	b.Grow(n * len(s)) // alokasi cukup sejak awal
	for range n {
		b.WriteString(s)
	}
	return b.String()
}

// concatSprintf menggabungkan n string menggunakan fmt.Sprintf
func concatSprintf(n int, s string) string {
	r := ""
	for range n {
		r = fmt.Sprintf("%s%s", r, s)
	}
	return r
}

// Parameter kecil: 10 iterasi, string pendek
func BenchmarkConcatPlus_Small(b *testing.B) {
	for range b.N {
		concatPlus(10, "x")
	}
}

func BenchmarkConcatBuilder_Small(b *testing.B) {
	for range b.N {
		concatBuilder(10, "x")
	}
}

func BenchmarkConcatSprintf_Small(b *testing.B) {
	for range b.N {
		concatSprintf(10, "x")
	}
}

// Parameter besar: 1000 iterasi, string panjang
func BenchmarkConcatPlus_Large(b *testing.B) {
	for range b.N {
		concatPlus(1000, "hello-")
	}
}

func BenchmarkConcatBuilder_Large(b *testing.B) {
	for range b.N {
		concatBuilder(1000, "hello-")
	}
}

func BenchmarkConcatSprintf_Large(b *testing.B) {
	for range b.N {
		concatSprintf(1000, "hello-")
	}
}

// ---------------------------------------------------------------
// Benchmark 2: Slice Pre-allocation vs Append Dinamis
// ---------------------------------------------------------------

// sliceNoCap membuat slice tanpa kapasitas awal, append n elemen
func sliceNoCap(n int) []int {
	s := []int{}
	for range n {
		s = append(s, 1)
	}
	return s
}

// slicePreAlloc membuat slice dengan kapasitas awal n, append n elemen
func slicePreAlloc(n int) []int {
	s := make([]int, 0, n)
	for range n {
		s = append(s, 1)
	}
	return s
}

func BenchmarkSliceNoCap_100(b *testing.B) {
	for range b.N {
		sliceNoCap(100)
	}
}

func BenchmarkSlicePreAlloc_100(b *testing.B) {
	for range b.N {
		slicePreAlloc(100)
	}
}

func BenchmarkSliceNoCap_10000(b *testing.B) {
	for range b.N {
		sliceNoCap(10000)
	}
}

func BenchmarkSlicePreAlloc_10000(b *testing.B) {
	for range b.N {
		slicePreAlloc(10000)
	}
}

// ---------------------------------------------------------------
// Benchmark 3: Map vs Slice untuk Lookup Kecil
// ---------------------------------------------------------------

var map5 = map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true}
var slice5 = []int{1, 2, 3, 4, 5}
var map100 = func() map[int]bool {
	m := make(map[int]bool, 100)
	for i := range 100 {
		m[i] = true
	}
	return m
}()
var slice100 = func() []int {
	s := make([]int, 100)
	for i := range 100 {
		s[i] = i
	}
	return s
}()

// lookupSlice memeriksa keberadaan nilai secara linear
func lookupSlice(s []int, v int) bool {
	for _, x := range s {
		if x == v {
			return true
		}
	}
	return false
}

// lookupMap memeriksa keberadaan nilai menggunakan map
func lookupMap(m map[int]bool, v int) bool {
	return m[v]
}

func BenchmarkSliceLookup5_Hit(b *testing.B) {
	for range b.N {
		lookupSlice(slice5, 5)
	}
}

func BenchmarkMapLookup5_Hit(b *testing.B) {
	for range b.N {
		lookupMap(map5, 5)
	}
}

func BenchmarkSliceLookup5_Miss(b *testing.B) {
	for range b.N {
		lookupSlice(slice5, 99)
	}
}

func BenchmarkMapLookup5_Miss(b *testing.B) {
	for range b.N {
		lookupMap(map5, 99)
	}
}

func BenchmarkSliceLookup100_Hit(b *testing.B) {
	for range b.N {
		lookupSlice(slice100, 99)
	}
}

func BenchmarkMapLookup100_Hit(b *testing.B) {
	for range b.N {
		lookupMap(map100, 99)
	}
}

func BenchmarkSliceLookup100_Miss(b *testing.B) {
	for range b.N {
		lookupSlice(slice100, 999)
	}
}

func BenchmarkMapLookup100_Miss(b *testing.B) {
	for range b.N {
		lookupMap(map100, 999)
	}
}
