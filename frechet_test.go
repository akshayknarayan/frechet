package frechet

import (
	"testing"
)

func TestFrechet(t *testing.T) {
	curve1 := []Point{Point{x: 0, y: 0}, Point{x: 1, y: 1}, Point{x: 2, y: 2}}
	curve2 := []Point{Point{x: 0, y: 1}, Point{x: 1, y: 2}, Point{x: 2, y: 3}}

	dist := Frechet(curve1, curve2)
	if dist != 1.0 {
		t.Fatalf("%v != 1.0", dist)
	}
}

func BenchmarkFrechet(b *testing.B) {
	curve1 := []Point{}
	curve2 := []Point{}

	for i := 0; i < 1000; i++ {
		curve1 = append(curve1, Point{x: float64(i), y: float64(i)})
		curve2 = append(curve2, Point{x: float64(i), y: float64(i + 1)})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dist := Frechet(curve1, curve2)
		if dist != 1.0 {
			b.Fatalf("%v != 1.0", dist)
		}
	}
}
