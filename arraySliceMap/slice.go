package arraySliceMap

import "fmt"

func SliceInit() {
	var s0 []int
	fmt.Println(len(s0), cap(s0))
	s0 = append(s0, 1)
	fmt.Println(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	fmt.Println(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	fmt.Println(len(s2), cap(s2))
	fmt.Println(s2[0], s2[1], s2[2])
	s2 = append(s2, 1)
	fmt.Println(s2[0], s2[1], s2[2], s2[3])
	fmt.Println(len(s2), cap(s2))
}

func SliceGrowing() {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		fmt.Println(len(s), cap(s))
	}
}

func SliceShareMemory() {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep",
		"Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	fmt.Println(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	fmt.Println(summer, len(summer), cap(summer))
	summer[0] = "Unknow"
	fmt.Println(Q2)
	fmt.Println(year)
}

func SliceComparing() {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}
	//if a == b { //切片只能和nil比较
	//	fmt.Println("equal")
	//}
	fmt.Println(a, b)
}

// SliceLengthCap commited by chenqgp
// Slice's length and capacity could be modifiable.
// The capacity would doubled when capcities has got full,
// Tip: The capacity would got 125% than before when
// appending more than 1000 item in Slice.
func SliceLengthCap() {
	slice := []int{1, 2, 3, 4, 5}
	a := slice[2:4]
	fmt.Println(len(a), cap(a))
	// var slice will change?
	a = append(a, 6)
	fmt.Println(len(a), cap(a), slice)
	a = append(a, 7)
	fmt.Println(len(a), cap(a))
	a = append(a, 8)
	fmt.Println(len(a), cap(a))
	a = append(a, 9)
	fmt.Println(len(a), cap(a))
	a = append(a, 10)
	fmt.Println(len(a), cap(a), len(slice), cap(slice), slice)
}

// SliceSpecifiedCap commited by chenqgp
// New Slice can specify capacity while separating the Slice,
// and the feature as same as example above.
// Tip: If the capacity that specified after `the second colon` while
// resplitting a new Slice more than the Slice, will occurr an error.
func SliceSpecifiedCap() {
	source := []int{1, 2, 3, 4, 5}
	slice := source[2:3:4]
	fmt.Println(slice)
	// Effected var source.
	slice = append(slice, 8)
	fmt.Println(source, slice, cap(slice))
	// Does it effect var source that var slice's capacity
	// more than `(4 - 2)` while appending new item?
	slice = append(slice, 9)
	fmt.Println(source, slice, cap(slice))

	// Occurred error
	//bSlice := source[2:3:6]
	//fmt.Println(bSlice)
}

// SliceNil commited by chenqgp
// If you want to create a nil or empty Slice.
// The example showing how.
func SliceNil() {
	var a []int
	b := []int{}
	c := make([]int, 0)
	fmt.Printf("%v,%v,%v", a == nil, b == nil, c == nil)
}

// SliceCap commited by chenqgp
// It will performed The advanced features about cap of Slice.
func SliceCap() {
	slice := make([]int, 2, 5)
	// the slice below created length was 2, capacity was 5; now need
	// to resplit a new one from the slice which index start from 1 and
	// give the 3 into brackets after colon that it means the length from 
    // index 0 through 2.
	// will report error under line?
	newSlice := slice[1:3]
	// Answer is no.

	// the newSlice's length and capacity: Capacities of the slice assume a `k`; So,
	// length = 3 - 1
	// cap = k - 1

	// Following below sets, the newSlice remaining capacity was 4, visit 
    // the index 2 (out of length[0,1] bound) of the newSlice that whether 
    // has the expected value 0.
	fmt.Println(len(newSlice), cap(newSlice), newSlice[2])
	// Answer is no.
	// the extra capacity is useless.
}

// SliceOverflow commited by chenqgp
// It will performed Slice created the copy when it's capacity overflowing.
func SliceOverflow() {
	slice := []int{1, 2, 3, 4}
	newS := slice[1:3]
	newS = append(newS, 5)
	// Var newS length was 2, capacitiy was 3 before it has appended number 5,
	// then length is 3, capacity 3.
	fmt.Println(slice)
	// Append agagin, the capacity has overflowed, newS will create new copy and
	// new Array for the Slice.
	newS = append(newS, 6)
	newS[0] = 10
	// Checking the variable slice has effected or not.
	fmt.Println(slice)
}

// SliceOverflow commited by chenqgp
// We can set the capacity as same as the length of the new Slice 
// while resplitting the Slice. we will see what happened example below.
func SliceAvoidSideEffect() {
	source := []int{1, 2, 3, 4, 5}
	slice := source[2:3:3]
	// once var slice has appended item, var slice capacity would overflow,
	// Go will create copy of slice and append the item(666) into the copy.
	slice = append(slice, 666)
	fmt.Println(source, slice, cap(slice))
    // A new Slice will create copy of slice and its array data after exceeding 
    // capacity through appending item, based on the length equals the capacity.
}

// SliceIterate commited by chenqgp
// It performed `range` create the copy of every item of the Slice.
// Tip: Every copy from `range` that the point has same pointer(addr).
func SliceIterate() {
    source := []int{1,2,3,4,5}
    for index, value := range source {
        if value == 3 {
            value = 999
        }
        if value == 2 {
            source[index] = 666
        }
    }
    fmt.Println(source)
}
