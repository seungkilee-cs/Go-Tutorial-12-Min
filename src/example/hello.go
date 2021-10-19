package main

import (
	"errors"
	"fmt"
	"math"
)

func sum(x int, y int) int {
	return x + y
}

// specify the return value(s)
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Square Root Undefined for Negative Number")
	} else {
		return math.Sqrt(x), nil
	}
}

// Class Struct
type person struct {
	name string
	age  int
}

func main() {
	p1 := person{name: "Seung Ki", age: 28}

	// make for map, delete for value of the map
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	m["sum"] = sum(m["a"], m["b"])
	float := -4.0
	sqrt_float, err := sqrt(float)

	// loop over maps!
	for key, value := range m {
		fmt.Println("key: ", key, "value: ", value)
	}

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sqrt_float)
	}

	fmt.Print(p1)

	// PTR for pass by ref
	ptr := 7
	inc(&ptr)
	fmt.Println(ptr)

}

func inc(x *int) {
	*x++
}
