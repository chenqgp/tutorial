package sample

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
	// if a == b { //切片只能和nil比较
	// 	fmt.Println("equal")
	// }
	fmt.Println(a, b)
}
