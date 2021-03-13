package wren

import (
	"github.com/alexdesi/collisions/vectors"
)

type Sphere struct {
	X, Y     float64
	Velocity vectors.Vector
	Radius   float64
	Mass     float64
}

// u1 = (v1 * ( m1 - m2) + 2 * m2 * v2 ) /( m1 + m2)
// u2 = (v2 * ( m2 - m1) + 2 * m1 * v1 ) /( m1 + m2)

func Impact(sp1, sp2 *Sphere) {
	vectorNormal := vectors.Vector{sp1.X - sp2.X, sp1.Y - sp2.Y} // Vector whose direction is normal (perpendicular) to the surfaces of the objects at the point of collision.

	Un := vectors.UnitNormal(vectorNormal)
	Ut := vectors.UnitTangent(Un)

	v1n := Un.Dot(sp1.Velocity)
	v1t := Ut.Dot(sp1.Velocity)

	v2n := Un.Dot(sp2.Velocity)
	v2t := Ut.Dot(sp2.Velocity)

	v1TangAfter := v1t
	v2TangAfter := v2t

	// v1n_after = ( v1n * (m1 + m2) + 2 (m2 * v2n)  ) / m1 + m2
	v1NorAfter := (v1n*(sp1.Mass-sp2.Mass) + 2*sp2.Mass*v2n) / (sp1.Mass + sp2.Mass)
	v2NorAfter := (v2n*(sp2.Mass-sp1.Mass) + 2*sp1.Mass*v1n) / (sp1.Mass + sp2.Mass)

	// v1NorAfter := v2n // in case of m1 == m2
	// v2NotAfter := v1n // in case of m1 == m2
	V1After := Un.Multiply(v1NorAfter).Add(Ut.Multiply(v1TangAfter))
	V2After := Un.Multiply(v2NorAfter).Add(Ut.Multiply(v2TangAfter))

	sp1.Velocity = V1After
	sp2.Velocity = V2After
}
