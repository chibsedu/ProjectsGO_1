// P O I N T E R S

package main

import ( 
    "fmt"
)

func main() {
	i:= 7
	inc(&i)		// Passage par reference (&i <--> Memory adress of the variable i)
				// L'appel inc(i) passera une copie de i seulement.
	fmt.Println(i)    
}

func inc(x *int) {
	*x++		// Dereference, THEN modifie the original value
				// This will modifie the original value   
				// Without * (dereferencing), this will increment the memory adress
}
