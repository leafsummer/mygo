package main

import (
    "fmt"
    "time"
)

const LIM = 50

var fibs [LIM]uint64

func main() {
    var result uint64 = 0
    start := time.Now()
    for i:=0; i < LIM; i++ {
        result = fibonacci(i)
        fmt.Printf("fibonacci(%d) is: %d ", i, result)
    }
    end := time.Now()
    delta := end.Sub(start)
    fmt.Printf("\nlong Calculation took this amount of time: %s\n", delta)

    result = 0
    start = time.Now()
    for i:=0; i < LIM; i++ {
        result = fibonacciMem(i)
        fmt.Printf("fibonacci(%d) is: %d ", i, result)
    }
    end = time.Now()
    delta = end.Sub(start)
    fmt.Printf("\nlong Calculation took this amount of time: %s\n", delta)
}

func fibonacci(n int) (res uint64) {
    if n <= 1 {
        res = 1
    } else {
        res = fibonacci(n-1) + fibonacci(n-2)
    }
    return
}

func fibonacciMem(n int) (res uint64) {
    if fibs[n] != 0 {
        res = fibs[n]
        return
    }
    if n <= 1 {
        res = 1
    } else {
        res = fibonacciMem(n-1) + fibonacciMem(n-2)
    }
    fibs[n] = res
    return
}
