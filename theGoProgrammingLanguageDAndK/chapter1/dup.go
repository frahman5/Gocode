//Dup prints the count and text of lines that appear more than once 
// in the input. It reads from stdin or from a list of named files.
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    counts := make(map[string]int)
    filenametracker := make(map[string]string)
    files := os.Args[1:]
    if len(files) == 0 {
        countLines(os.Stdin, counts, filenametracker)
    } else {
        for _, arg := range files {
            f, err := os.Open(arg)
            if err != nil {
                fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
                continue
            }
            countLines(f, counts, filenametracker)
            f.Close()
        }
    }
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\t%s\n", n, line, filenametracker[line])
        }
    }
}

func countLines(f *os.File, counts map[string]int, filenametracker map[string]string) {
    input := bufio.NewScanner(f)
    for input.Scan() {
        counts[input.Text()]++
        if counts[input.Text()] > 1 {
            filenametracker[input.Text()] = f.Name()
        }
    }
    // NOTE: ignoring potential errors from input.Err()`
}