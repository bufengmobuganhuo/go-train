package main

import "testing"

func TestTriangle(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{30000, 40000, 50000},
	}

	for _, tt := range tests {
		if actual := calcTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("calcTriangle(%d, %d); got %d; expected %d", tt.a, tt.b, actual, tt.c)
		}
	}
}

func BenchmarkTriangle(b *testing.B) {
	num1, num2 := 300, 400
	ans := 500
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if actual := calcTriangle(num1, num2); actual != ans {
			b.Errorf("got %d for input (%d, %d); expected %d", actual, num1, num2, ans)
		}
	}
}
