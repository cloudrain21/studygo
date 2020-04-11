package calc

import (
	"testing"
)

func TestSum(t *testing.T) {
	r := Sum1(1, 2)
	if r != 3 {
		t.Errorf("want : 3, act : %d", r)
	}
}

func BenchmarkSum1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum1(i, i+1)
	}
}

func BenchmarkSum2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum2(i, i+1)
	}
}
