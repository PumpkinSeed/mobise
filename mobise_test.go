package mobise

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestSearch(t *testing.T) {
	var arr = []int{7, 8, 9, 10, 11, 13, 24, 1, 2, 3, 4, 5, 6}

	value := Search(arr, 0, len(arr)-1)
	if arr[value] != 1 {
		t.Errorf("Value should be 1, instead of %d", arr[value])
	}
}

func BenchmarkSearch(b *testing.B) {
	var arr = generateTestCase()
	fmt.Println(len(arr))

	for i := 0; i < b.N; i++ {
		Search(arr, 0, len(arr)-1)
	}
}

func generateTestCase() []int {
	max := random(10000, 99999)
	var first, second []int
	for i := max / 2; i < max; i += random(2, 4) {
		second = append(second, i)
	}

	for i := 0; i < max/2; i += random(2, 4) {
		first = append(first, i)
	}

	return append(second, first...)
}

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
