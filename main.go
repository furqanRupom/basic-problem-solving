package main

import (
	"fmt"

	"github.com/furqanrupom/basic-problem-solving/problems"
)

func main() {

	/* 1. Reverse word */
	problem1 := problems.ReverseWords("Hello World")
	fmt.Println(problem1)

	/* 2. Find all pairs with a given sum */
	numList := []int{2, 6, 10, 12}
	targetSum := 16
	problem2 := problems.SumInArray(numList, targetSum)
	fmt.Println(problem2)

	/* 3. Check palindrome */
	palindromeList := []int{1, 2, 2, 1}
	problem3 := problems.CheckPalindrome(palindromeList)
	fmt.Println(problem3)

	/* 4. Rotate matrix */
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	problems.RotateMatrix(matrix)
	fmt.Println(matrix)

	/* 5. Remove duplicates */
	duplicatesWord := "favourr"
	problem5 := problems.RemoveDuplications(duplicatesWord)
	fmt.Println(problem5)

	/* 6. LIS */
	longList := []int{1, 2, 6, 8, 19, 111}
	problem6 := problems.LenOfLongest(longList)
	fmt.Println(problem6)

	/* 7. Count vowels (your version) */
	vowelInput := "hello world"
	problem7 := problems.CountVowels(vowelInput)
	fmt.Println(problem7)

	/* 8. Majority element */
	majorityArr := []int{2, 2, 1, 2, 3, 2, 2}
	problem8 := problems.MajorityElements(majorityArr)
	fmt.Println(problem8)

	/* 9. Rotate string */
	problem9 := problems.RotateString("hello", 2)
	fmt.Println(problem9)

	/* 10. Missing elements in sequence */
	missingArr := []int{1, 2, 4, 7}
	problem10 := problems.FindMissing(missingArr)
	fmt.Println(problem10)

	/* 11. Rotate array */
	rotateArr := []int{1, 2, 3, 4, 5}
	problems.RotateArray(rotateArr, 2)
	fmt.Println(rotateArr)

	/* 12. Missing number */
	missingNumArr := []int{1, 2, 4, 5}
	problem12 := problems.MissingNumber(missingNumArr)
	fmt.Println(problem12)

	/* 13. Merge arrays */
	a := []int{1, 3, 5}
	b := []int{2, 4, 6}
	problem13 := problems.MegetSortedArray(a, b)
	fmt.Println(problem13)

	/* 14. Longest substring without repeating */
	problem14 := problems.LengthOfLongestSubstring("abcabcbb")
	fmt.Println(problem14)

	/* 15. Check arrays equal */
	arr1 := []int{1, 2, 3}
	arr2 := []int{3, 2, 1}
	problem15 := problems.ArraysEqual(arr1, arr2)
	fmt.Println(problem15)

	/* 16. Move zeros */
	zeroArr := []int{0, 1, 0, 3, 12}
	problems.MoveZeroes(zeroArr)
	fmt.Println(zeroArr)

	/* 17. Custom map */
	mapped := problems.Map([]int{1, 2, 3}, func(x int) int {
		return x * 2
	})
	fmt.Println(mapped)

	/* 18. Intersection */
	int1 := []int{1, 2, 3}
	int2 := []int{2, 3, 4}
	problem18 := problems.Intersection(int1, int2)
	fmt.Println(problem18)

	/* 19. First non-repeating char */
	problem19 := problems.FirstUnique("swiss")
	fmt.Println(string(problem19))

	/* 20. Perfect square */
	problem20 := problems.IsPerfectSquare(16)
	fmt.Println(problem20)

	/* 23. Count Special Char */

	problem23 := problems.CountTotalSpecialChar("AaBb")
	fmt.Println(problem23)

	fmt.Println("BASIC PROBLEMS SOLVING IN GO")
}
