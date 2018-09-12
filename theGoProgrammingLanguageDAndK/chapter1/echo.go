//Echo prints its command-line arguments
package main 

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    for index, elem := range(os.Args[1:]) {
        fmt.Println(strconv.Itoa(index) + " " + elem)
    }
}