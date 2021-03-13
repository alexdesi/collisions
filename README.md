# Collisions
Elastic collisions of Sphere using Ebiten game library	Elastic collisions of Sphere using Ebiten game library

*Collisions* is a simple elastic collisions simulator of two spheres in a 2D space.

![Elastic collisions of two spheres in Go]("images/elastic collision of two spheres.png")

It calculates the final velocities of the spheres using simple linear algebra and avoiding the use of complex trigonometry.

The vector operations needed to calculate the elastic collision are simply outlined in the doc [2-Dimensional Elastic Collisions without Trigonometry](https://www.vobarian.com/collisions/2dcollisions2.pdf) written by [@vobarian](https://github.com/vobarian).

I have not imported any external library for the vector operations, I wrote *vectors.go* with few basic functions.

*Collisions* is made of four packages:

- wren: it contains the core function to compute the speed of the sphere after the impact.
- detector: it contains the function to detect the collision of the sphere, with each other, or with the edges.
- shapes2D: it contains the function to draw the circles (the spheres).
- vectors: it contains the basic functions to work with vectors.

## How did you animated it?
I used [Ebiten](https://github.com/hajimehoshi/ebiten) - a simple graphics library to build games in Go.


## How to run *Collisions*?

```
go run main.go
```

## Why did you write *Collisions*?

A while ago I wrote something similar in Java, but I wasn't happy with the result and I did not share it.
Now, I am learning Go, and I decided to resume that old idea to write my Hello Word in Go.

## What's next?
Some ideas are:

- Use any number of spheres
- Add friction
- Add gravity
- Add inelastic collision
- ...

Feel free to clone this project and expand it at your whish.

---
[Alessandro De Simone](https://alessandro.desi)
