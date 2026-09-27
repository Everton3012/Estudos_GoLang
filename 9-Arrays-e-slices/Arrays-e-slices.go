package main

import "fmt"

func main() {

	var array1 = [5]int{1, 2, 3, 4, 5}

	fmt.Println("Array 1:", array1)

	array2 := [6]int{0, 1, 2, 3, 4, 5}
	fmt.Println("Array 2:", array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println("Array 3:", array3)

	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice 1:", slice1)

	slice1 = append(slice1, 6)
	fmt.Println("Slice 1 after append:", slice1)

	slice2 := make([]int, 5)
	fmt.Println("Slice 2:", slice2)

	slice3 := make([]int, 5, 10)
	fmt.Println("Slice 3:", slice3)
	fmt.Println("Slice 3 length:", len(slice3))
	fmt.Println("Slice 3 capacity:", cap(slice3))

	slice4 := array1[1:4]
	fmt.Println("Slice 4:", slice4)

	slice4[0] = 10
	fmt.Println("Slice 4 after modification:", slice4)
	fmt.Println("Array 1 after slice modification:", array1)
}
