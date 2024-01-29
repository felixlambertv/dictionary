package dictionary

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

type Point struct {
	x, y, z float64
}

func (p *Point) ClosestPoints(points []*Point, limit int) []*Point {
	var closestPoints []*Point
	var maxDist float64

	for _, point := range points {
		dist := p.DistanceTo(point)

		if len(closestPoints) < limit {
			closestPoints = append(closestPoints, point)
			if dist > maxDist {
				maxDist = dist
			}
		} else if dist < maxDist {
			for i, cp := range closestPoints {
				if p.DistanceTo(cp) == maxDist {
					closestPoints[i] = point
					break
				}
			}

			maxDist = 0
			for _, cp := range closestPoints {
				if d := p.DistanceTo(cp); d > maxDist {
					maxDist = d
				}
			}
		}
	}

	return closestPoints
}

func NewPoint(x float64, y float64, z float64) *Point {
	return &Point{x: clamp(x), y: clamp(y), z: clamp(z)}
}

func (p *Point) DistanceTo(other *Point) float64 {
	return math.Sqrt(math.Pow(p.x-other.x, 2) + math.Pow(p.y-other.y, 2) + math.Pow(p.z-other.z, 2))
}

func clamp(coordinate float64) float64 {
	if coordinate < 0 {
		return 0
	}
	if coordinate > 1 {
		return 1
	}
	return coordinate
}

func RandomPoint() *Point {
	return NewPoint(rand.Float64(), rand.Float64(), rand.Float64())
}

func TestShortest(t *testing.T) {
	points := make([]*Point, 10_000_000)
	for i := 0; i < 10_000_000; i++ {
		points[i] = RandomPoint()
	}
	point := NewPoint(0, 0, 0)
	fmt.Println(point.ClosestPoints(points, 10))
}

func BenchmarkName(b *testing.B) {
	b.ResetTimer()
	points := make([]*Point, 10_000_000)
	for i := 0; i < 10_000_000; i++ {
		points[i] = RandomPoint()
	}
	for i := 0; i < b.N; i++ {
		point := NewPoint(0, 0, 0)
		point.ClosestPoints(points, 10)
	}
}
