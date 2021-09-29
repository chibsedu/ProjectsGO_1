// Simple function call

package main

import ( 
    "fmt"
)

func main() {
    result := sum(7, 6)
    fmt.Println(result)
    
}

func sum(x int, y int)int{
    return x + y
}