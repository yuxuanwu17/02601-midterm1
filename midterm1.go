// 02-601 Midterm #1, Fall 2021
//
// You have until Friday, October 8th at 11:59pm to complete this exam.  You
// may not discuss the exam with anyone except the course staff.
//
// You may look at any material on golang.org or on the course's Diderot site.
// No other material may be consulted.
//
// There are 7 programming problems. Do not change the function signatures.
// Please leave the problem descriptions in your submitted file.
//
// Please submit a gzipped tar file that contains only one file (midterm1.go)
// that contains the solutions to all the problems.
//
// Each function has a dummy return statement that you should replace with your
// code. Your code should compile, so if you choose to not do a problem, keep
// the an appropriate `return` statement in the function.
//
// You may want to create a `main` function in a separate main.go file that
// tests your functions. Do not turn in such a main.go file, and do not add a
// main() function to this midterm1.go file. You can add helper functions to
// this file if you believe that is the best organization of your code.
//
// You can use any standard packages included with Go.
//
// Each problem is worth 14 points. You get 2 points for correctly submitting
// the file.
//
// NOTE: You will only be able to submit this file ONCE. So be sure you are
// finished before submitting.

// Replace the following with your information:
// Name: Yuxuan Wu
// Andrew ID: yuxuanwu

package main

/*******************************************************************************************
Problem 1. Write a function that takes in a list []bool and returns the number
of times the list switches from true to  false or from false to true when read
left-to-right.

For example, if the list is []bool{true, true, false, false, true, false}, your
function should return 3.
*******************************************************************************************/

func CountSwitches(b []bool) int {
	count := 0
	for i := 0; i < len(b)-1; i++ {
		if b[i] == b[i+1] {
			continue
		} else {
			count++
		}
	}

	return count
}

/*******************************************************************************************
Problem 2. Write a function that takes in a [][]int list and return the
smallest integer i such that f[x][i] < 0 for some integer x. If no such i
exists, return -1.

For example, if the input is:
    []int{
        []int{3,4,-5,1,7},
        []int{10,10,10},
        []int{2, 1, 0, -1},
        []int{3,-3,-3,0},
        []int{4,-4,5,0},
    }

Your function should return 1.
*******************************************************************************************/

func LeftmostNegative(f [][]int) int {
	// create a list to store the index of each subarray's first negative value
	storeFirstNegVal := make([]int, 0)
	for _, ints := range f {
		for i, i2 := range ints {
			if i2 < 0 {
				storeFirstNegVal = append(storeFirstNegVal, i)
				break
			}
		}
	}
	1if len(storeFirstNegVal) == 0 {
		return -1
	}

	minIndex := min(storeFirstNegVal)

	return minIndex
}

func min(l []int) (min int) {
	min = l[0]
	for _, v := range l {
		if v < min {
			min = v
		}
	}
	return
}

/*******************************************************************************************
Problem 3. Write a function called FirstUnique that takes a slice of integers
and returns the first integer that only occurs once in the slice. If no
integers occur only once, return 0.

Examples:
    FirstUnique([]int{2, 3, 4, 5, 2, 4, 5}) should return 3.
    FirstUnique([]int{7, 8, 2, 8, 2, 7, 8}) should return 0.
    FirstUnique([]int{8, 8, 8, 9, 8, 6, 8, 8}) should return 9.
*******************************************************************************************/

func FirstUnique(list []int) int {
	return 0
}

/*******************************************************************************************
Problem 4. A list of numbers A is a circular permutation of another list B if
you can write both lists around a circle and rotate the circles so that the
positions of the numbers are identical.

For example:
    A = 1,7,8,10,31,14
    B = 10,31,14,1,7,8
    C = 7,8,1,10,14,31
A and B are circular permutations of each other, while C is not a circular
permutation of A or B.

Write a function IsCircularPermutation that takes two int lists as parameters
and returns true if one is a circular permutation of the other (and false
otherwise).
*******************************************************************************************/

func IsCircularPermutation(a, b []int) bool {
	return true
}

/*******************************************************************************************
Problem 5. Write a function Contains(l1, l2 []int) bool that returns true if
the integers in l2 are a subset  of the integers in l1 and false otherwise.
That is, the function returns true if every integer that appears in l2 also
appears someplace in l1. If an integer occurs in l2 more than once, it needs to
occur at least that many times in l1.
*******************************************************************************************/

func Contains(l1, l2 []int) bool {
	return true
}

/*******************************************************************************************
Problem 6. Write a function SquareNumbers that takes a list of integers and
returns the numbers in that list that are square numbers (can be written as x*x
for some integer x). The output list should be in the same order as the input list.

Example: If your input is []int{1,5,9,9,20} the returned value should be []int{1,9,9}.
*******************************************************************************************/

func SquareNumbers(a []int) []int {
	return []int{}
}

/*******************************************************************************************
Problem 7. Consider the following Node type. It contains a single field which
is a pointer to another Node. That other Node of course will itself contain a
field pointing to a Node.  By following these pointers we can walk through a
series of Nodes. If `next` is nil, the series of Nodes ends.

Write a function HasCycle that walks through the series of Nodes staarting with
the parameter n and returns true if you can ever return back to a node you have
already visited, and false otherwise.

Be sure to consider the following case:

n -> n1 -> n2 -> n3 -> n4 -> n5 -+
                 ^               |
                 |               |
                 +---------------+

which should return true.

You cannot modify the Node type or the pointers in the Node.
*******************************************************************************************/

type Node struct {
	next *Node
}

func HasCycle(n *Node) bool {
	return true
}

// END OF MIDTERM
