// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math"
)

//find the second biggest number
func find_biggest_and_second_biggest() {
	array := []int32{10, 2, 200, 100, 342, 0, 220, 500}

	max := int32(0)
	min := int32(0)
	fmt.Printf("len %v \n", len(array))
	for i := 0; i < len(array); i++ {
		if array[i] > max {
			min = max
			max = array[i]
		} else if array[i] > min {
			min = array[i]
		}

	}
	fmt.Printf("Hello, 2hd biggest %v biggest %v \n", min, max)
}

//find the second biggest number
func find_lowest_and_second_lowest() {
	array := []uint32{10, 2, 200, 100, 342, 0, 220, 500}

	var max uint32 = math.MaxUint32
	var min uint32 = math.MaxUint32
	fmt.Printf("len %v \n", len(array))
	for i := 0; i < len(array); i++ {
		if array[i] < max {
			min = max
			max = array[i]
		} else if array[i] < min {
			min = array[i]
		}

	}
	fmt.Printf("Hello, 2hd smallest %v smallest %v \n", max, min)
}

func main() {
	find_biggest_and_second_biggest()
	find_lowest_and_second_lowest()
}
