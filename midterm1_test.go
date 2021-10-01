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
	test1 := [][]int{
		{3, 4, -5, 1, 7},
		{10, 10, 10},
		{2, 1, 0, -1},
		{3, -3, -3, 0},
		{4, -4, 5, 0},
	}
	fmt.Println(LeftmostNegative(test1))

	fmt.Println("======================================")
	test2 := [][]int{
		{3, 4, 5, 1, 7},
		{10, 10, 10},
		{2, 1, 0, 1},
		{3, 3, 3, 0},
		{4, 4, 5, 0},
	}
	fmt.Println(LeftmostNegative(test2))
}
