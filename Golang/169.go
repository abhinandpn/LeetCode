package golang

// func MajorityElement(nums []int) int {
// 	var elem, count int
// 	for i := range nums {
// 		if count == 0 {
// 			count++
// 			elem = nums[i]
// 		} else {
// 			if nums[i] != elem {
// 				count--
// 			} else {
// 				count++
// 			}
// 		}

//		}
//		return elem
//	}
func MajorityElement(nums []int) int {
	key := make(map[int]int)
	for _, n := range nums {
		key[n]++
	}

	res := 0
	max := 0
	for i, n := range key {
		if n > max {
			max = n
			res = i
		}
	}

	return res
}
