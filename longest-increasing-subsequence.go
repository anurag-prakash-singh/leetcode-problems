package main

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func lengthOfLIS(nums []int) int {
	n := len(nums)
	mem := make([]int, n)
	lis := 0

	for i := 0; i < n; i++ {
		lisI := 1

		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] {
				lisI = max(lisI, mem[j]+1)
			}
		}

		mem[i] = lisI

		lis = max(mem[i], lis)
	}

	return lis
}

func main() {

}
