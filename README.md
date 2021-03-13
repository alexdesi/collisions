# Collisions

*Collisions* is a simple elastic collisions simulator of two spheres in a 2D space.

It calculates the velocities and directions of the spheres after collisions, using linear algebra and avoiding the use of complex trigonometry.

![Elastic collisions of two spheres in Go](https://github.com/alexdesi/collisions/blob/dfae43c9d7b510f166a5fbddae4a3687a576485d/images/elastic%20collision%20of%20two%20spheres.png)

The vector operations needed to calculate the elastic collision are simply outlined in the doc [2-Dimensional Elastic Collisions without Trigonometry](https://www.vobarian.com/collisions/2dcollisions2.pdf) written by [@vobarian](https://github.com/vobarian).

I have not imported any external library for the vector operations, I wrote *vectors.go* with few basic functions.

*Collisions* is made of four packages:

- wren: it contains the core function to calculate the speed of the spheres after the impact.
- detector: it contains the functions to detect the collision of the sphere, with each other, or with the edges.
- shapes2D: it contains the function to draw the circles (the spheres).
- vectors: it contains the basic functions to work with vectors.

## How did you animated the spheres?
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
