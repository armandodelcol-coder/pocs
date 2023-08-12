package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{
	color.White,
	color.RGBA{0x00, 0xff, 0xBB, 0xff},
	color.Black,
	color.RGBA{230, 230, 250, 1},
	color.RGBA{255, 140, 0, 1},
	color.RGBA{139, 0, 139, 1},
	color.RGBA{240, 230, 140, 1},
}

var cycles = 5

const (
	whiteIndex = 0
	blackIndex = 1
)

func lissajous(out io.Writer) {
	const (
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			rand_pallete_i := rand.Intn(3-1) + 1
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(rand_pallete_i))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		if k == "cycles" {
			cyclesParam, err := strconv.Atoi(v[0])
			if err != nil {
				log.Print(err)
				break
			}

			cycles = cyclesParam
			log.Print(cycles)
		}
	}
	lissajous(w)
}
