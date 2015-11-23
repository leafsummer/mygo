package main

import "fmt"

func main() {
    // fmt.Println("Starting the program")
    // panic("A severe error occurred: stopping the program!")
    // fmt.Println("Ending the program")
    fmt.Printf("Calling test\r\n")
    test()
    fmt.Print("Test completed\r\n")
}

func badCall() {
    panic("bad end")
}

func test() {
    defer func() {
        if e := recover(); e != nil {
            fmt.Printf("Panicing %s\r\n", e)
        }
    }()
    badCall()
    fmt.Prinf("After bad call\r\n")
}