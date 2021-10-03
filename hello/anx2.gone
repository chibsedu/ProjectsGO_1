// Function with multiple return values

package main

import ( 
    "fmt"
	"math"
	"errors"
)

func main() {
    result, err := sqrt(-9)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func sqrt(x float64) (float64, error){
	if x < 0 {
		return 0, errors.New("Undefined for negative values")
	}
	return math.Sqrt(x), nil
}