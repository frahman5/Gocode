// just a playfile to test different ways of declaring variables
package main 

import "fmt"

func main() {
    s := "Hello world"
    var t string
    var m, n = "m", "n"
    var l string = "Hello I am n l with an unnecessary long function declaration"
    t = "Hello World I'm a t"
    fmt.Println(s)
    fmt.Println(t)
    fmt.Println(m)
    fmt.Println(n)
    fmt.Println(l)
}