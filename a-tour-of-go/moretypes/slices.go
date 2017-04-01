package main

import "fmt"

func main() {
    s := []int{2, 3, 5, 7, 11,13}
    fmt.Println("s ==", s)

    for i := 0; i < len(s); i++ {
        fmt.Println("s[%d] == %d\n", i, s[i])
    }
}
