package array

func TwoSum(nums []int, target int) []int {
	var resultIndices []int
	n := len(nums)

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				resultIndices = append(resultIndices, i)
				resultIndices = append(resultIndices, j)
			}
		}
	}
	return resultIndices
}
