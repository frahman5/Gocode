//Server is a minimal "echo" and counter server
package main 

import (
    "log"
    "net/http"
    "image"
    "image/color"
    "image/gif"
    "io"
    "math"
    "math/rand"
    "strconv"
)

var palette = []color.Color{color.Black, color.RGBA{0x00, 0xff, 0x00, 0xff}}

const (
    whiteIndex = 0 // first color in palette
    blackIndex = 1 // next color in palette
)

var cycles = 20.0 // number of complete x oscillator revolutions
const (
    res     = 0.001 // angular resolution
    size    = 100   // image canvas covers [-size..+size]
    nframes = 64    // number of animation frames
    delay   = 8     // delay between frames in 10ms units
)

func main() {
    http.HandleFunc("/", handler) // each request calls handler
    log.Fatal(http.ListenAndServe("localhost:8007", nil))
}

// handler echoes the HTTP request
func handler(w http.ResponseWriter, r *http.Request) {
    u := r.URL
    q := u.Query()

    cycles_int, err := strconv.Atoi(q.Get("cycles"))
    check(err)
    cycles = float64(cycles_int)
    // fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
    // for k, v := range r.Header {
    //     fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
    // }
    // fmt.Fprintf(w, "Host = %q\n", r.Host)
    lissajous(w)
}

func lissajous(out io.Writer) {
    freq := rand.Float64() * 3.0 // relative frequency of y oscillator
    anim := gif.GIF{LoopCount: nframes}
    phase := 0.0 // phase difference
    for i := 0; i < nframes; i++ {
        rect := image.Rect(0, 0, 2*size+1, 2*size+1)
        img := image.NewPaletted(rect, palette)
        for t := 0.0; t < cycles*2*math.Pi; t += res {
            x := math.Sin(t)
            y := math.Sin(t*freq + phase)
            img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
        }
        phase += 0.1
        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }
    gif.EncodeAll(out, &anim) // Note : ignoring encoding errors
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
