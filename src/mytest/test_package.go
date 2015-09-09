package main

import (
    "fmt"
    "unsafe"
    "container/list"
)

func main() {
    var i int = 10
    size := unsafe.Sizeof(i)
    fmt.Println("The size of an int is: ", size)

    lst := list.New()
    lst.PushBack(100)
    lst.PushBack(101)
    lst.PushBack(102)

    fmt.Println("Here is the double linked list:\n", lst)

    for e := lst.Front(); e != nil; e = e.Next() {
        fmt.Println(e)
        fmt.Println(e.Value)
    }
}