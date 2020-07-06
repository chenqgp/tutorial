package arraySliceMap

import "fmt"

func MapWithFunValue() {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	fmt.Println(m[1](2), m[2](2), m[3](2))
}

func MapForSet() {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 3
	if mySet[n] {
		fmt.Printf("%d is existing", n)
	} else {
		fmt.Printf("%d is not existing", n)
	}
	mySet[3] = true
	fmt.Println(len(mySet))
	delete(mySet, 1)
	n = 1
	if mySet[n] {
		fmt.Printf("%d is existing", n)
	} else {
		fmt.Printf("%d is not existing", n)
	}
}

type tmap string

func (m tmap) strings() string {

	return string(m)
}

// MapKey commited by chenqgp
// It performs the features about Key in Map.
func MapKey() {
	// Does it work? the key is a Slice or something else..

	/*
	m := [[]int]string{}
	tm := tmap("map").strings
	m2 := [tm]string{} 
	*/

	// Slice, function, data structures within Slice can not be a key in Map,
	// but you can treat them as a value in map.
	m := map[string][]int{}
	m["one"] = []int{1}
	tm := map[string]tmap{"map": tmap("map"), "tmap": tmap(tmap("tmap").strings())}
	fmt.Println(m, tm)
}

// MapNil commited by chenqgp
// It performed how to create nil map, usually we don't have to do this, the best
// practice is to initialize a Map like below:
func MapNil() {
	// Does var mn could assign a value? we want store some values into mn.

	/* 
	var mn map[string]int
	mn["store"] = 1098
	*/

	// if a Map is nil, we can't store value into it. we will occur a error message 
	// if we do this:
	// Panic: `assignment to entry in nil map`.

	// let's try another situation.
	mm := make(map[string]int)
	m := map[string]int{}
	fmt.Println(mm, m)
	mm["mm"] = 1
	m["m"] = 2
	fmt.Println(mm, m)
	// It works well.
}

// MapExists commited by chenqgp
func MapExists() {
	m := make(map[string]string)
	m["color"] = "blue"
	m["shape"] = "square"
	if value, ok := m["height"]; !ok {
		fmt.Printf("value: %T, errors: %s", value, "what is height? we don't store this value.")
	}
	// We can see the type kind of the value was the "zero"(empty string) 
	// that the type was declared it before by us. We would determine the key
	// was not exists from wheter the second returned value is false or wheter 
	// the value is "zero".
	// Tip: The map must not be an nil Map.

	// Base on we talked above, these situation work well.
	if m["height"] == "" {
		fmt.Println("The height key is not exists.")
	}

	mInt := map[string]int{"one": 1}
	if mInt["second"] == 0 {
		fmt.Println("The second key is not exists.")
	}
	
	// Delete the key and value in a Map.
	delete(m, "color")

	fmt.Println(m)
}

// MapDeliver commited by chenqgp
func MapDeliver(m map[string]int) {
	m["add"] = 2

	// When you delivered a Map into a function, It won't make a copy of the Map.
	// Quite similar with Slice in this case.
}
