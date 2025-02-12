//basic random number generator

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10)) //rand.Intn(n) returns a random integer between 0 and n-1	
}
