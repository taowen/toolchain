package main

import (
    "example.com/a/example/lib"
    "fmt"
    "github.com/thoas/go-funk"
)

func main() {
    fmt.Println(funk.Uniq(lib.Data))
}
