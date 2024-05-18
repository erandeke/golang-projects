package main

import "fmt"

/*
Two sums
Input: nums = [2,7,11,15], target = 9
 * Output: [0,1]
 * Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
*/

//Aproach in n*n and space in O(n)

func computeTwoSum(arr [4]int, target int) []int {

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				return []int{i, j}
			}
		}

	}
	return []int{}

}

//Appraoch two : using map TC : O(n) and space O(n)

func computeTwoSumUsingHash(arr [4]int, target int) []int {

	m := make(map[int]int, 10)

	//get the map filled using arr
	for idx, elements := range arr {
		m[elements] = idx
	}

	//iterate over array to compute the sum

	for index, elem := range arr {
		if indexY, ok := m[target-elem]; ok && index != indexY {
			return []int{index, indexY}
		}
	}

	return []int{}
}

func main() {
	var balance = [4]int{2, 7, 11, 15}
	target := 9
	a := computeTwoSumUsingHash(balance, target)
	fmt.Println(a)

	fmt.Println(IsNumberPalindrome(121))

}
