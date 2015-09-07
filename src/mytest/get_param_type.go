package main

import (
    "fmt"
    "reflect"
)

func main() {
    var x float64 = 3.4
    var a = 'c'
    var t = reflect.TypeOf(x)
    var ta = reflect.TypeOf(a)
    fmt.Printf("%v\n", t)
    fmt.Println("type:", t)
    fmt.Printf("%v\n", t.Kind())
    fmt.Println("a type:", ta)
    var b = int(a)
    fmt.Println(b)
    var c = string(b)
    // var c = rune(b)
    var tc = reflect.TypeOf(c)
    fmt.Println(tc)

    for _, char := range []rune("世界你好") {
        fmt.Println(char)
        fmt.Println(string(char))
    }
}