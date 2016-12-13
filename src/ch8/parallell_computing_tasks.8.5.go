package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"net/http"
	"runtime"
	"sync"
	"time"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

// 4 workers initialized.
// Paralell verison rendered in: 845.243275ms
// No paralell version rendered in: 2.313840059s
func main() {
	paralell()
	noParalell()
}

func paralell() {
	workers := runtime.NumCPU()
	fmt.Printf("%d workers initialized.\n", workers)
	var wg sync.WaitGroup
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	start := time.Now()
	rows := make(chan int, height) // Channel with buffer the capacity of height
	for row := 0; row < height; row++ {
		rows <- row
	}
	// Close channel, can't use this channel to send data anymore,
	// but other goroutine can revceive from that channel.
	close(rows)
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			for py := range rows {
				y := float64(py)/width*(ymax-ymin) + ymin
				for px := 0; px < width; px++ {
					x := float64(px)/width*(xmax-xmin) + xmin
					z := complex(x, y)
					img.Set(px, py, newton(z))
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("Paralell verison rendered in:", time.Since(start))
	http.HandleFunc("/paralell", func(w http.ResponseWriter, r *http.Request) {
		png.Encode(w, img) // NOTE: ignoring errors
	})
	// 这句话是阻塞的，一直listen在8080端口。所以得注释掉才能继续往下运行下一个函数。
	// log.Fatal(http.ListenAndServe(":8080", nil))
}

func noParalell() {
	start := time.Now()
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/width*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, newton(z))
		}
	}

	fmt.Println("No paralell version rendered in:", time.Since(start))
	http.HandleFunc("/no_paralell", func(w http.ResponseWriter, r *http.Request) {
		png.Encode(w, img) // NOTE: ignoring errors
	})
	// log.Fatal(http.ListenAndServe(":8081", nil))
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	iterations := 37
	for n := uint8(0); int(n) < iterations; n++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(cmplx.Pow(z, 4)-1) < 1e-6 {
			return color.Gray{255 - uint8(math.Log(float64(n))/math.Log(float64(iterations+0))*255)}
		}
	}
	return color.Black
}
