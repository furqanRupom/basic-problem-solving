package main

import (
	"fmt"

	"github.com/furqanrupom/basic-problem-solving/problems"
)

func main() {
	/* 1. Reverse word */
	problem1 := problems.ReverseWords("Hello World")
	fmt.Println(problem1)

	/*2. Find all pairs with a given sum in an array */
	numList := []int{2, 6, 10, 12}
	targetSum := 16
	problem2 := problems.SumInArray(numList, targetSum)
	fmt.Println(problem2)

	/*2. check the array is palindrome or not */
	palindromeList := []int{1, 2, 2, 1, 1, 19}
	problem3 := problems.CheckPalindrome(palindromeList)
	fmt.Println(problem3)

	fmt.Println("BASIC PROBLEMS SOLVING IN GO")
}
