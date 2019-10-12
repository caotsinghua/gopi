package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

var palette = []color.Color{
	color.White,
	// color.Black,
	color.RGBA{0, 0, 255, 1},
	color.RGBA{0, 255, 0, 1},
	color.RGBA{255, 0, 0, 1},
}

const (
	whiteIndex = 0
	blackIndex = 1
)

func lissajous(out io.Writer, cycles float64) {
	const (
		res     = 0.001
		size    = 100
		nframes = 128
		delay   = 8
	)
	freq := rand.Float64() * 3
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	// 循环次数 nframes 生成64张图
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2.0*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(math.Floor((rand.Float64()*3)+1)))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
