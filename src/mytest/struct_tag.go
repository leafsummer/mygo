package main

import (
    "fmt"
    "reflect"
)

type TagType struct { //tags
    field1 bool "An important answer"
    field2 string "The name of the thing"
    field3 int "How much there are"
}

func main() {
    tt := TagType{true, "Barak obama", 1}
    for i := 0; i < 3; i++ {
        refTag(tt, i)
    }
}

func refTag(tt TagType, ix int){
    ixType := reflect.TypeOf(ix)
    ttType := reflect.TypeOf(tt)
    ixField := ttType.Field(ix)
    fmt.Printf("%v %v %v\n", ixType, ttType, ixField.Tag)
}