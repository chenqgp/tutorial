package sample

import "fmt"

func ArrayInit() {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr3 := [...]int{1, 3, 4, 5}
	arr1[1] = 5
	fmt.Printf("%+v, %+v", arr[1], arr[2])
	fmt.Printf("%+v, %+v", arr1, arr3)
}

func ArrayTravel() {
	arr3 := [...]int{1, 3, 4, 5}
	for i := 0; i < len(arr3); i++ {
		fmt.Printf("%+v", arr3[i])
	}
	for _, e := range arr3 {
		fmt.Printf("%+v", e)
	}
}

func ArraySection() {
	arr3 := []int{1, 2, 3, 4, 5}
	arr3_sec := arr3[:]
	fmt.Printf("%+v", arr3_sec)
}

// ArrayTypeOf commited by chenqgp
// It performed the how array is an `value type` in GO.
func ArrayTypeOf() {
	var arr1 [2]int
	arr1[0] = 1
	arr1[1] = 2
	newA := arr1

	newA[0] = 3

	fmt.Println(arr1, newA)
}

// ArrayCopy commited by chenqgp
// It performed assignment operation of the multidimensional array and
// the subset of the multidimensinal array assigned to one dimensonal array.
func ArrayCopy() {
	var arr1 [2]int
	var arr2 [2]int
	arr1[0] = 1
	arr1[1] = 2
	arr2 = arr1
	fmt.Println(arr1, arr2)
	arr1 = [...]int{3, 4}
	arr2 = arr1
	fmt.Println(arr1, arr2)

	var arr3 = [3][2]int{[2]int{1, 2}, [2]int{5, 6}}
	arr2 = arr3[1]
	fmt.Println(arr3[1], arr3[2], arr2)
}

// ArrayPointer commited by chenqgp
// It performed assignment operation of pointer value from two array
// and uncertainly length of an array assign to certainly length of an array.
func ArrayPointer() [2]*int {
	var arrV [2]*int
	arrP := [...]*int{new(int), new(int)}
	arrV = arrP

	*arrP[0] = 10
	*arrP[1] = 20

	return arrV
}

// ReceiveLargeArray commited by chenqgp
// It performed large array passed through into function's problem
// and how to change its value.
func ReceiveLargeArray(array *[1e6]int) {
	// ...
	val := array[0]
	fmt.Println()
	fmt.Printf("%T, %v", val, val)

	array[0] = 7

	val = 8
}
