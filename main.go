package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/alexdesi/collisions/detector"
	"github.com/alexdesi/collisions/shapes2D"
	"github.com/alexdesi/collisions/vectors"
	"github.com/alexdesi/collisions/wren"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

func (g *Game) Update() error {
	// Actual vel for sp1

	sp1.X = sp1.X + sp1.Velocity[0]
	sp2.X = sp2.X + sp2.Velocity[0]

	sp1.Y = sp1.Y + sp1.Velocity[1]
	sp2.Y = sp2.Y + sp2.Velocity[1]

	if detector.DetectCollision(sp1, sp2) {
		wren.Impact(&sp1, &sp2)
	}

	borderVector, borderCollision := detector.DetectCollisionWithBorders(sp1)
	if borderCollision {
		sp1.Velocity = vectors.Vector{sp1.Velocity[0] * borderVector[0], sp1.Velocity[1] * borderVector[1]}
	}

	borderVector, borderCollision = detector.DetectCollisionWithBorders(sp2)
	if borderCollision {
		sp2.Velocity = vectors.Vector{sp2.Velocity[0] * borderVector[0], sp2.Velocity[1] * borderVector[1]}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	red := color.RGBA{0xff, 0, 0, 0xff}

	shapes2D.DrawCircle(screen, int(sp1.X), int(sp1.Y), int(sp1.Radius), red)
	// drawCircle(screen, int(sp1.X), int(sp1.Y), 2, red)
	// ebitenutil.DrawLine(screen, sp1.X, sp1.Y, sp1.X+sp1.Velocity[0]*100, sp1.Y+sp1.Velocity[1]*100, red)

	shapes2D.DrawCircle(screen, int(sp2.X), int(sp2.Y), int(sp2.Radius), red)
	// drawCircle(screen, int(sp2.X), int(sp2.Y), 2, red)
	// ebitenutil.DrawLine(screen, sp2.X, sp2.Y, sp2.X+sp2.Velocity[0]*100, sp2.Y+sp2.Velocity[1]*100, red)

	// drawVector(screen, Un)
	// drawVector(screen, Ut)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

var sp1 = wren.Sphere{X: 100, Y: 340, Velocity: vectors.Vector{3, 1}, Mass: 1, Radius: 20}
var sp2 = wren.Sphere{X: 300, Y: 240, Velocity: vectors.Vector{-1, 0}, Mass: 2, Radius: 100}

func main() {
	fmt.Println("Elastic collision")
	fmt.Println(sp1)

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
