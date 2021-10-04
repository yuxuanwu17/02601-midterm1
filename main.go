package main

import "fmt"

func main() {

	// bool values
	b1 := []bool{true, true, false, false, true, false}
	b2 := []bool{true, false, true, false}
	b3 := []bool{true}
	b4 := []bool{true, true, true, true, true}

	t1 := CountSwitches(b1)
	t2 := CountSwitches(b2)
	t3 := CountSwitches(b3)
	t4 := CountSwitches(b4)
	fmt.Println("Testing Q1...")
	fmt.Println(t1, t2, t3, t4)

	// int values

	i1 := []int{3, 4, -5, 1, 7}
	i2 := []int{10, 10, 10}
	i3 := []int{2, 1, 0, -1}
	i4 := []int{3, -3, -3, 0}
	i5 := []int{4, -4, 5, 0}

	// list values
	l1 := [][]int{i1, i2, i3, i4, i5}
	l2 := [][]int{i2, i3}
	l3 := [][]int{i1, i2}
	l4 := [][]int{i2}
	f1 := LeftmostNegative(l1)
	f2 := LeftmostNegative(l2)
	f3 := LeftmostNegative(l3)
	f4 := LeftmostNegative(l4)
	fmt.Println("Testing Q2...")
	fmt.Println(f1, f2, f3, f4)

	i6 := []int{2, 3, 4, 5, 2, 4, 5}
	i7 := []int{7, 8, 2, 8, 2, 7, 8}
	i8 := []int{8, 8, 8, 9, 8, 6, 8, 8}

	r1 := FirstUnique(i4)
	r2 := FirstUnique(i6)
	r3 := FirstUnique(i7)
	r4 := FirstUnique(i8)
	r5 := FirstUnique([]int{4})
	r6 := FirstUnique([]int{0, 0})
	fmt.Println("Testing Q3...")
	fmt.Println(r1, r2, r3, r4, r5, r6) // 3 3 0 9 4 0

	i9 := []int{1, 7, 8, 10, 31, 14}
	i10 := []int{10, 31, 14, 1, 7, 8}
	i11 := []int{7, 8, 1, 10, 14, 31}
	i12 := []int{2, 4, 5, 6, 7, 2}
	i13 := []int{4, 5, 6, 7, 2, 2}
	q1 := IsCircularPermutation(i9, i10)
	q2 := IsCircularPermutation(i9, i11)
	q3 := IsCircularPermutation(i10, i11)
	q4 := IsCircularPermutation(i12, i13)
	q5 := IsCircularPermutation([]int{2, 2, 2, 3}, []int{3, 2, 2, 2})
	q6 := IsCircularPermutation([]int{2, 2, 2, 3}, []int{4, 2, 2, 2})
	q7 := IsCircularPermutation([]int{2}, []int{2})
	fmt.Println("Testing Q4...")
	fmt.Println(q1, q2, q3, q4, q5, q6, q7) // true false false true true false true

	i14 := []int{1, 2, 3, 4, 5, 5, 6}
	i15 := []int{1, 2, 3, 6, 6}
	i16 := []int{1, 2, 3, 4, 5}
	i17 := []int{2}
	a1 := Contains(i14, i15)
	a2 := Contains(i14, i17)
	a3 := Contains(i14, i16)
	a4 := Contains(i15, i14)
	fmt.Println("Testing Q5...")
	fmt.Println(a1, a2, a3, a4) // false true true false

	i18 := []int{121, 1, 5, 9, 9, 25, 81}
	i19 := []int{1, 9}
	i20 := []int{1}
	i21 := []int{2, 5, 8, 10}
	i22 := []int{0}
	c1 := SquareNumbers(i18)
	c2 := SquareNumbers(i19)
	c3 := SquareNumbers(i20)
	c4 := SquareNumbers(i21)
	c5 := SquareNumbers(i22)
	fmt.Println("Testing Q6...")
	fmt.Println(c1, c2, c3, c4, c5)
}
