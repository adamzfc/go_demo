package main

import (
    "log"
    "errors"
    "time"
)

type GoDoer func()

func Go(todo GoDoer) {
    go func() {
        defer func() {
            if err := recover(); err != nil {
                log.Printf("Panic in safely.Go: %s", err)
            }
        }()
        todo()
    }()
}

func message() {
    println("Inside goroutine")
    panic(errors.New("Oops!"))
}

func main() {
    Go(message)
    println("Outside goroutine")
    time.Sleep(2000)
}
