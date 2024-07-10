package main

// https://leetcode.cn/problems/two-sum/description/
func main() {

}

func twoSum(nums []int, target int) (result []int) {
	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			if v+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSum2(nums []int, target int) (result []int) {
	hashTable := make(map[int]int)
	for i, x := range nums {
		v, ok := hashTable[target-x]
		if ok {
			// If v exists, the index of v must be smaller than i, so it must be placed in front of result
			return []int{v, i}
		}
		hashTable[x] = i
	}
	return nil
}
