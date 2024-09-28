# FastNoise Lite

[Web Preview App](https://auburn.github.io/FastNoiseLite)

FastNoise Lite is an noise generation package with a large selection of noise algorithms.

## Features

- 2D & 3D sampling
- OpenSimplex2 noise
- OpenSimplex2S noise
- Cellular (Voronoi) noise
- Perlin noise
- Value noise
- Value Cubic noise
- OpenSimplex2-based domain warp
- Basic Grid Gradient domain warp
- Multiple fractal options for all of the above
- Supports floats and/or doubles

## Getting Started

Here's an example for creating a 128x128 array of OpenSimplex2 noise

```go
// Create and configure noise state (either float32 or float64)
var noise = fastnoise.New[float32]()
noise.NoiseType(fastnoise.OpenSimplex2)

// Gather noise data
var noiseData [128][128]float32

for x := 0; x < 128; x++ {
	for y := 0; y < 128; y++ {
		noiseData[x][y] = noise.Noise2D(x, y)
	}
}

// Do something with this data...
```
