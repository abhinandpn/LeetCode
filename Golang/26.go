package golang

func RemoveDuplicates(nums []int) int {

	count := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] != nums[j] {
				count++
			}
		}
	}
	return count
}
