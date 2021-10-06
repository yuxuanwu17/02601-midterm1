package main

import (
	"fmt"
	"testing"
)

func TestCountSwitches(t *testing.T) {
	fmt.Println("=======================第一题===================================")
	booleanTest1 := []bool{}
	fmt.Println(CountSwitches(booleanTest1))
	booleanTest2 := []bool{true, true, false, false, true, false, true, false}
	fmt.Println(CountSwitches(booleanTest2))
}

func TestLeftmostNegative(t *testing.T) {
	fmt.Println("=======================第二题===================================")
	test1 := [][]int{
		{3, 4, -5, 1, 7},
		{-10, 10, 10},
		{2, 1, 0, -1},
		{3, -3, -3, 0},
		{4, -4, 5, 0},
	}
	fmt.Println(LeftmostNegative(test1))

	test2 := [][]int{
		{3, 4, 5, 1, 7},
		{10, 10, 10},
		{2, 1, 0, 1},
		{3, 3, 3, 0},
		{4, 4, 5, 0},
	}
	fmt.Println(LeftmostNegative(test2))
}

func TestFirstUnique(t *testing.T) {
	fmt.Println("=======================第三题===================================")
	fmt.Println(FirstUnique([]int{2, 3, 4, 5, 2, 4, 5}))
	fmt.Println(FirstUnique([]int{7, 8, 2, 8, 2, 7, 8}))
	fmt.Println(FirstUnique([]int{8, 8, 8, 9, 8, 6, 8, 8}))
}

func TestIsCircularPermutation(t *testing.T) {
	fmt.Println("=======================第四题===================================")
	// true
	//a := []int{2, 4, 5, 6, 7, 2}
	//b := []int{4, 5, 6, 7, 2, 2}
	//b := []int{7, 8, 1, 10, 14, 31} // false case

	//true
	//c := []int{2, 3, 3, 4}
	//d := []int{3, 3, 4, 2}
	//d := []int{3, 4, 4, 2}

	//fmt.Println(IsCircularPermutation(a, b))
	//fmt.Println(IsCircularPermutation(c, d))
	fmt.Println(IsCircularPermutation([]int{5, 6, 7, 8, 5, 12}, []int{8, 7, 6, 5, 12, 5}))
	//fmt.Println(IsCircularPermutation([]int{2, 3, 3, 3}, []int{3, 2, 3, 3}))
	//fmt.Println(IsCircularPermutation([]int{1,1,1,1,1,1 }, []int{1,2,1,1,1,1}))
}

func TestContains(t *testing.T) {
	fmt.Println("=======================第五题===================================")
	// l2's 10 occurs 3 times, while l1 10 occurs 2 times => false
	l1 := []int{1, 7, 8, 10, 10, 31, 14}
	l2 := []int{10, 10, 10, 31, 14, 1, 7, 8}
	fmt.Println(Contains(l1, l2))

	// true
	l3 := []int{1, 7, 8, 10, 31, 14}
	l4 := []int{10, 31, 14, 1, 7, 8}
	fmt.Println(Contains(l3, l4))

	// true
	l5 := []int{1, 7, 8, 10, 10, 31, 14}
	l6 := []int{10, 31, 14, 1, 7, 8}
	fmt.Println(Contains(l5, l6))

}

func TestSquareNumbers(t *testing.T) {
	nums := []int{1, 5, 9, 9, 20}
	fmt.Println(SquareNumbers(nums))
	nums2 := []int{1, 5, 9, 9, 20, 21, 23, 36}
	fmt.Println(SquareNumbers(nums2))
	nums3 := []int{1, 5, 9, 9, 20}
	fmt.Println(SquareNumbers(nums3))

	nums4 := make([]int, 0)
	for i := 0; i < 1000; i++ {
		nums4 = append(nums4, i)
	}
	fmt.Println(SquareNumbers(nums4))
}

func TestHasCycle(t *testing.T) {

}
