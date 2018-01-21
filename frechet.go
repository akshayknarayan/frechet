package frechet

import "math"

// Point is used to represent curves
type Point struct {
	x float64
	y float64
}

func euclideanDistance(p1 Point, p2 Point) float64 {
	dx := p2.x - p1.x
	dy := p2.y - p1.y
	return math.Sqrt(dx*dx + dy*dy)
}

func min(x float64, y float64, z float64) float64 {
	if x < y {
		return math.Min(x, z)
	}

	return math.Min(y, z)
}

// Frechet is a dynamic programming implementation calculating the frechet distance between the two curves c1 and c2.
func Frechet(c1 []Point, c2 []Point) float64 {
	I := len(c1)
	J := len(c2)
	res := make([][]float64, I)
	for i := range res {
		res[i] = make([]float64, J)
		for j := range res[i] {
			res[i][j] = -1.0
		}
	}

	for i := 0; i < I; i++ {
		for j := 0; j < J; j++ {
			if res[i][j] > -1 {
				return res[i][j]
			}

			if i == 0 && j == 0 {
				res[i][j] = euclideanDistance(c1[0], c2[0])
			} else if i > 0 && j == 0 {
				res[i][j] = math.Max(res[i-1][0], euclideanDistance(c1[i], c2[0]))
			} else if i == 0 && j > 0 {
				res[i][j] = math.Max(res[0][j-1], euclideanDistance(c1[0], c2[j]))
			} else if i > 0 && j > 0 {
				res[i][j] = math.Max(
					min(
						res[i-1][j-1],
						res[i][j-1],
						res[i-1][j],
					),
					euclideanDistance(c1[i], c2[j]),
				)
			} else {
				res[i][j] = math.Inf(1)
			}
		}
	}

	return res[I-1][J-1]
}
