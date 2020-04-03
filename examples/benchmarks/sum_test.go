package benchmarks

import "testing"

func BenchmarkSumForward(b *testing.B) {
	nums := []int{}
	for i := 0; i < 5; i++ {
		nums = append(nums, i)
	}
	for n := 0; n < b.N; n++ {
		sum := nums[0] + nums[1] + nums[2] + nums[3] + nums[4]
		_ = sum
	}
}
func BenchmarkSumBackward(b *testing.B) {
	nums := []int{}
	for i := 0; i < 5; i++ {
		nums = append(nums, i)
	}
	for n := 0; n < b.N; n++ {
		sum := nums[4] + nums[3] + nums[2] + nums[1] + nums[0]
		_ = sum
	}
}
