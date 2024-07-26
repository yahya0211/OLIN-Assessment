package models

import (
	"fmt"
	"sort"
	"test/entities"
)

func TwoSum(nums []int, target int) entities.TwoSums {
	numsMap := make(map[int]int)
	for i, num := range nums {
		pair := target - num
		if idx, found := numsMap[pair]; found {
			return entities.TwoSums{Output: []int{idx, i}}
		}
		numsMap[num] = i
	}
	return entities.TwoSums{Output: []int{}}
}

func ThreeSum(nums []int) entities.ThreeSums {
	var result [][]int

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return entities.ThreeSums{Output: result}
}

func FindSubstring(s string, words []string) entities.ExpectedOutput {
	if len(words) == 0 || len(words[0]) == 0 {
		return entities.ExpectedOutput{Output: []int{}}
	}

	wordLen := len(words[0])
	fmt.Println("Word length:", wordLen)
	numWords := len(words)
	fmt.Println("Number of words:", numWords)
	totalLen := wordLen * numWords
	fmt.Println("Total length:", totalLen)
	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}

	var result []int

	for i := 0; i <= len(s)-totalLen; i++ {
		seen := make(map[string]int)
		j := 0
		for ; j < numWords; j++ {
			word := s[i+j*wordLen : i+(j+1)*wordLen]
			if count, exists := wordCount[word]; exists {
				seen[word]++
				if seen[word] > count {
					break
				}
			} else {
				break
			}
		}
		if j == numWords {
			result = append(result, i)
		}
	}


	return entities.ExpectedOutput{Output: result}

	// Jika menginginkan input yang diharapkan, maka untuk "S":"goodgoodgoodbestwordword" 
	// dan "words":["word","good","best","good"] atau 
	// 
}
