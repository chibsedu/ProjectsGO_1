package main

import "fmt"

func main() {
    // Intro
    fmt.Println("Hello, World!")

//// A R R A Y S
    // Array 5 int
    var a [5]int
    a[2] = 7
    fmt.Println(a)

    // Array Other Initialization
    b := [5]int{5, 4, 3, 2, 1}
    fmt.Println(b)

    // The length of an Array is part of the Array's type and can't be changed.


//// S L I C E S
    // To overcome this, we use Slices by simply removing the element count (ici: 5).
    c := []int{5, 4, 3, 2, 1}
    c = append(c, 13)  // append doesn't modifie the original slice and returns a new one
    fmt.Println(c)

//// M A P
    vertices := make(map[string]int) // ie: map[keys]values

    vertices["triangle"] = 2
    vertices["square"] = 5
    vertices["dodecagon"] = 32

    // fmt.Println(vertices["square"])  --> Gives 5
    // delete(vertices, "square")       --> Deletes "square" element from the map
    fmt.Println(vertices)           

//// L O O P S

    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
    /* while loop
    i := 0
    for i < 5 {
        fmt.Println(i)
        i++
    } 
    */

    // Iterate over each element of an array or slice
    arr := []string{"a", "b", "c"}
    for index, value := range arr {
        fmt.Println("index:", index, "value:", value)
    }
    
    // Iterate over each element of a map
    for key, value := range vertices {
        fmt.Println("key:", key, "value:", value)
    }
    
}

