package detector

import (
	"math"

	"github.com/alexdesi/collisions/vectors"
	"github.com/alexdesi/collisions/wren"
)

func DetectCollision(sp1, sp2 wren.Sphere) bool {
	return (Distance(sp1, sp2) <= sp1.Radius+sp2.Radius)
}

func Distance(sp1, sp2 wren.Sphere) float64 {
	return math.Sqrt(math.Pow(sp2.X-sp1.X, 2) + math.Pow(sp2.Y-sp1.Y, 2))
}

func DetectCollisionWithBorders(sp wren.Sphere) (vectors.Vector, bool) {
	if sp.X-sp.Radius <= 0 || sp.X+sp.Radius >= 640 {
		return vectors.Vector{-1, 1}, true
	}
	if sp.Y-sp.Radius <= 0 || sp.Y+sp.Radius >= 480 {
		return vectors.Vector{1, -1}, true
	}
	return vectors.Vector{0, 0}, false
}
