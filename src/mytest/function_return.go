package main

import (
    "fmt"
    "strings"
)

func main() {
    // make an Add2 function, give it a name p2, and call it:
    p2 := Add2()
    fmt.Printf("Call Add2 for 3 gives: %v\n", p2(3))
    // make a special Adder function, a gets value 3:
    TwoAdder := Adder(3)
    fmt.Printf("The result is: %v\n", TwoAdder(3))

    addBmp := MakeAddSuffix(".bmp")
    addJpeg := MakeAddSuffix(".jpeg")

    fmt.Printf("the addBmp function result is: %v\n", addBmp("file"))
    fmt.Printf("the addJpeg function result is; %v\n", addJpeg("file"))
}

func Add2() func(b int) int {
    return func(b int) int {
        return b + 2
    }
}

func Adder(a int) func(b int) int {
    return func(b int) int {
        return a + b
    }
}

func MakeAddSuffix(suffix string) func(string) string {
    return func(name string) string {
        if !strings.HasSuffix(name, suffix) {
            return name + suffix
        }
        return name
    }
}