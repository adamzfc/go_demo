package main

import (
    "fmt"
    "unsafe"
    // "html"
    // "log"
    // "net/http"
)

func main() {
    x := 0x12345678

    p := unsafe.Pointer(&x)
    n := (*[4]byte)(p)

    for i := 0; i < len(n); i++ {
        fmt.Printf("%X ", n[i])
    }
}
