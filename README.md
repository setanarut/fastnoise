# FastNoise Lite

[Web Preview App](https://auburn.github.io/FastNoiseLite)

FastNoise Lite is a noise generation package with a large selection of noise algorithms.

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

Here's an example of OpenSimplex2 noise with DomainWarp

```go
	chars := " .:;+=xX$"
	noiseState := fastnoise.NewNoiseState[float32]()
	noiseState.Frequency = 0.05
	noiseState.DomainWarpType = fastnoise.DomainWarpOpenSimplex2Reduced
	noiseState.DomainWarpAmp = 30
	for y := range 10 {
		for x := range 80 {
			wx, wy := fastnoise.DomainWarp2D(x, y, noiseState)
			v := fastnoise.Value2D(wx, wy, noiseState)
			i := int((v + 1) / 2 * float32(len(chars)))
			fmt.Print(string(chars[i]))
		}
		fmt.Println()
	}
	// Output:
	// ++=xXX=+;;;+++==+;. :x$x;..    .:;;++;::....:;;. .;++==xXX$XX=+xX+. ;X$Xx+;;;;;:
	// =xxX$x+;;+++;;;;:..;xXx;;;+++;:....::..      .:. .;+====xxXxx==XX; .=X$Xx++;;;;:
	// XX$$X=+;+++++;;:::;xXx++xX$$XXx+;:...         .. .;==========xX$x:.;X$Xx=++;;;;;
	// X$Xx=+++====++;;;+=x=++x$XXxxxxx=+;:.    ...  .. .+xXXxx===xxX$X+.;x$Xx+;;;;;;;:
	// +=++;++=xx===++;+=x=;+xx=+;;;;+++=+;:............:=X$$$XXXXX$$X=;;x$X=;:..::::::
	// .::;+=xxxxx==++++==;;==+::;++;;:;+=+;::::::::::..+X$$$$$$$$$$X=++xXx;:.    ..:::
	// :;+=xxxxxxxx==+++=+:;++::;===++::;+=++;;;;;;;;:::x$$XXXxxxxxx=+=xX=;::::........
	// =xXXXXXXx======++=;:;+:.:=xxx=+;::;+=++++++++;::+X$XXxx====++=xXX=++=xxx=+;::...
	// X$$$$$Xx=++++==++=;.:;..;xXXxx=+:.:;++++++++;;:+xXXXXx=======xXXx++xX$$XXx=+;::.
	// $$$$$$Xx=;;;;=====;.::..+xXXxx=+;..::;;;;;;;::;=XxX$XxxxXXXXXXXx==xX$$$XXx===+;;
```
