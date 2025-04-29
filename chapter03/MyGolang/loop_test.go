package MyGolang

import "testing"

func CreateSource(len int) []int {
	nums := make([]int, 0, len)

	for i := 0; i < len; i++ {
		nums = append(nums, i)
	}

	return nums
}

func BenchmarkLoopStep1(b *testing.B) {
	// 制作源数据，长度为 10000
	src := CreateSource(10000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Loop(src, 1)
	}
}

func BenchmarkLoopStep2(b *testing.B) {
	// 制作源数据，长度为 10000
	src := CreateSource(10000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Loop(src, 2)
	}
}

func BenchmarkLoopStep3(b *testing.B) {
	// 制作源数据，长度为 10000
	src := CreateSource(10000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Loop(src, 3)
	}
}

func BenchmarkLoopStep4(b *testing.B) {
	// 制作源数据，长度为 10000
	src := CreateSource(10000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Loop(src, 4)
	}
}

func BenchmarkLoopStep5(b *testing.B) {
	// 制作源数据，长度为 10000
	src := CreateSource(10000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Loop(src, 5)
	}
}

func BenchmarkLoopStep6(b *testing.B) {
	// 制作源数据，长度为 10000
	src := CreateSource(10000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Loop(src, 6)
	}
}

func BenchmarkLoopStep12(b *testing.B) {
	// 制作源数据，长度为 10000
	src := CreateSource(10000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Loop(src, 12)
	}
}

func BenchmarkLoopStep16(b *testing.B) {
	// 制作源数据，长度为 10000
	src := CreateSource(10000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Loop(src, 16)
	}
}
