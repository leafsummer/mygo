package main
import "fmt"

const (
    WIDTH = 1920
    HEIGHT = 1080
)

func main(){
    var ar [3]int
    f(ar) // passes a copy of ar
    fp(&ar) // passes a pointer to ar

    // var arrAge = [5]int{18, 20, 15, 22, 16}
    // var arrLazy = [...]int{5, 6, 7, 8, 22}
    // var arrLazy = []int{5, 6, 7, 8, 22}
    var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
    // var arrKeyValue = []string{3: "Chris", 4: "Ron"}

    for i:=0; i < len(arrKeyValue); i++ {
        fmt.Printf("Person at %d is %s\n", i, arrKeyValue[i])
    }

    type pixel int
    var screen [WIDTH][HEIGHT]pixel

    
    for y := 0; y < HEIGHT; y++ {
        for x := 0; x < WIDTH; x++{
            screen[x][y] = 0
        }
    }

    array := [3]float64{7.0, 8.5, 9.1}
    x := Sum(&array) // Note the explicit address-of operator
    // to pass a pointer to the array
    fmt.Printf("The sum of the array is: %f", x)
}

func f(a [3]int) {fmt.Println(a)}

func fp(a *[3]int) {fmt.Println(a)}

func Sum(a *[3]float64) (sum float64) {
    for _, v := range a { // derefencing *a to get back to the array is not necessary!
        sum += v
    }
    return
}