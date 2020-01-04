package array_and_slice

import (
	"fmt"
	sc "strconv"
)

func Example_slice_test1() {
	n := 3
	fruits := make([]string, n)

	fruits[0] = "사과"
	fruits[1] = "바나나"
	fruits[2] = "토마토"

	fmt.Println(fruits)

	// Output:
	// [사과 바나나 토마토]
}

func Example_slicing() {
	nums := []int{11, 12, 13, 14, 15}
	fmt.Println(nums)
	fmt.Println(nums[1:3])
	fmt.Println(nums[2:])
	fmt.Println(nums[:3])
	fmt.Println(nums[1 : len(nums)-1])
	// Output:
	//[11 12 13 14 15]
	//[12 13]
	//[13 14 15]
	//[11 12 13]
	//[12 13 14]
}

func Example_append() {
	nums := []int{11, 12, 13, 14, 15}

	nums = append(nums, 16)
	fmt.Println(nums)

	i, _ := sc.Atoi("17")
	nums = append(nums, i)
	fmt.Println(nums)

	nums = append(nums, 18, 19, 20)
	fmt.Println(nums)

	// Output:
	//[11 12 13 14 15 16]
	//[11 12 13 14 15 16 17]
	//[11 12 13 14 15 16 17 18 19 20]
}

func Example_attach_slice() {
	f1 := []string{"사과", "바나나", "토마토"}
	f2 := []string{"포도", "딸기"}

	f3 := append(f1, f2...)
	fmt.Println(f3)

	f4 := append(f1[:2], f2...)
	fmt.Println(f4)

	// compile error, because f2 is not string but string slice
	//f5 := append(f1,f2)
	//fmt.Println(f5)

	f5 := append(f1, "감", "키위")
	fmt.Println(f5)

	// Output:
	//[사과 바나나 토마토 포도 딸기]
	//[사과 바나나 포도 딸기]
	//[사과 바나나 토마토 감 키위]
}

func Example_sliceCap() {
	nums := []int{1, 2, 3, 4, 5}

	fmt.Println(nums)
	fmt.Println("len:", len(nums))
	fmt.Println("cap:", cap(nums))
	fmt.Println()

	sliced1 := nums[:3]
	fmt.Println(sliced1)
	fmt.Println("sliced1 len:", len(sliced1))
	fmt.Println("sliced1 cap:", cap(sliced1))
	fmt.Println()

	sliced2 := nums[2:]
	fmt.Println(sliced2)
	fmt.Println("sliced2 len:", len(sliced2))
	fmt.Println("sliced2 cap:", cap(sliced2))
	fmt.Println()

	sliced3 := sliced1[:4]
	fmt.Println(sliced3)
	fmt.Println("sliced3 len:", len(sliced3))
	fmt.Println("sliced3 cap:", cap(sliced3))
	fmt.Println()

	nums[2] = 100
	fmt.Println(nums, sliced1, sliced2, sliced3)

	// Output:
	//[1 2 3 4 5]
	//len: 5
	//cap: 5
	//
	//[1 2 3]
	//sliced1 len: 3
	//sliced1 cap: 5
	//
	//[3 4 5]
	//sliced2 len: 3
	//sliced2 cap: 3
	//
	//[1 2 3 4]
	//sliced3 len: 4
	//sliced3 cap: 5
	//
	//[1 2 100 4 5] [1 2 100] [100 4 5] [1 2 100 4]
}

func Example_sliceCopy() {
	src := []int{30, 20, 50, 10, 40}
	dest := make([]int, len(src))
	dest2 := make([]int, len(src))
	dest3 := make([]int, len(src)-1)

	for i := range src {
		dest[i] = src[i]
	}
	fmt.Println(dest)

	copy(dest2, src)
	fmt.Println(dest2)

	if n := copy(dest3, src); n != len(src) {
		fmt.Printf("복사가 덜됐네... : %d\n", n)
	}

	dest4 := append([]int(nil), src...)
	fmt.Println(dest4)

	// Output:
	//[30 20 50 10 40]
	//[30 20 50 10 40]
	//복사가 덜됐네... : 4
	//[30 20 50 10 40]
}
