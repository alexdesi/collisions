package vectors

import "math"

type Vector [2]float64

func (u Vector) Dot(v Vector) float64 {
	return u[0]*v[0] + u[1]*v[1]
}

func (v Vector) Multiply(n float64) Vector {
	return Vector{v[0] * n, v[1] * n}
}

func (u Vector) Minus(v Vector) Vector {
	return Vector{u[0] - v[0], u[1] - v[1]}
}

func (u Vector) Add(v Vector) Vector {
	return Vector{u[0] + v[0], u[1] + v[1]}
}

func Normalize(v Vector) Vector {
	m := Magnitude(v)

	return Vector{v[0] / m, v[1] / m}
}

func Magnitude(v Vector) float64 {
	return math.Sqrt(math.Pow(v[0], 2) + math.Pow(v[1], 2))
}

func UnitNormal(v Vector) Vector {
	return Normalize(v)
}

func UnitTangent(v Vector) Vector {
	Un := UnitNormal(v)

	return Vector{-Un[1], Un[0]}
}
