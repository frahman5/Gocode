// Fetch prints the content found at each specified URL
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "strings"
)

func main() {
    for _, url := range os.Args[1:] {
        if !strings.HasPrefix(url, "http://") {
            url = "http://" + url
        }
        response, err := http.Get(url)
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
            os.Exit(1)
        }
        body, err := ioutil.ReadAll(response.Body)
        response.Body.Close()
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: reading %s : %v\n", url, err)
            os.Exit(1)
        }
        fmt.Printf("%s", body)
    }
}