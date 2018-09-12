//Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "time"
    "bufio"
)
    
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    start := time.Now()
    ch := make(chan string)
    for _, url := range os.Args[1:] {
        go fetch(url, ch) // start a goroutine
    }
    f, err := os.Create("output.txt")
    check(err)
    w := bufio.NewWriter(f)
    for range os.Args[1:] {
        _, err := w.WriteString(<-ch + "\n")
        check(err)
    }
    s := fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds())
    _, err = w.WriteString(s)
    check(err)
    w.Flush()
    f.Close()
}

func fetch(url string, ch chan<- string) {
    start := time.Now()
    resp, err := http.Get(url)
    if err != nil {
        ch <- fmt.Sprint(err) // send to channel ch
        return
    }

    nbytes, err := io.Copy(ioutil.Discard, resp.Body)
    resp.Body.Close() // don't leak resources
    if err != nil {
        ch <- fmt.Sprintf("while reading %s: %v", url, err)
        return
    }
    secs := time.Since(start).Seconds()
    ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}