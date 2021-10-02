package main

import (
	"fmt"
	"testing"
)

func TestCountSwitches(t *testing.T) {
	booleanTest1 := []bool{true, true, false, false, true, false}
	fmt.Println(CountSwitches(booleanTest1))
	booleanTest2 := []bool{true, true, false, false, true, false, true, false}
	fmt.Println(CountSwitches(booleanTest2))
}

func TestLeftmostNegative(t *testing.T) {
	fmt.Println("=============老师给的case========================")
	test1 := [][]int{
		{3, 4, -5, 1, 7},
		{-10, 10, 10},
		{2, 1, 0, -1},
		{3, -3, -3, 0},
		{4, -4, 5, 0},
	}
	fmt.Println(LeftmostNegative(test1))

	fmt.Println("=============全是正数的情况=========================")
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
	fmt.Println(FirstUnique([]int{2, 3, 4, 5, 2, 4, 5}))
	fmt.Println(FirstUnique([]int{7, 8, 2, 8, 2, 7, 8}))
	fmt.Println(FirstUnique([]int{8, 8, 8, 9, 8, 6, 8, 8}))
}

func TestIsCircularPermutation(t *testing.T) {
	a := []int{1, 7, 8, 10, 31, 14}
	b := []int{10, 31, 14, 1, 7, 8}
	//b := []int{7, 8, 1, 10, 14, 31}

	c := []int{2, 3, 3, 4}
	//d := []int{3,3,4,2}
	d := []int{3, 4, 4, 2}

	fmt.Println(IsCircularPermutation(a, b))
	fmt.Println(IsCircularPermutation(c, d))
}
