package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func getHosting1(ts, top []int) int {
	stime := time.Now()
	sort.Ints(ts)
	sort.Ints(top)

	total_cnt := 0

	for _, v := range top {
		cnt := 0
		for i := 0; i < len(ts); i++ {
			if ts[i] > v || cnt >= 5 {
				break
			}
			cnt++
		}
		ts = ts[cnt:]
		total_cnt += cnt
	}
	elapse := time.Since(stime)
	fmt.Printf("elapse time : %s\n", elapse)
	return total_cnt
}

func getHosting2(ts, top []int) int {
	stime := time.Now()
	ts1 := make([]int, 60)

	for _, v := range ts {
		ts1[v] += 1
	}

	total_cnt := 0

	for _, v := range top {
		start := v
		remain := 5
		for start >= 0 {
			if ts1[start] > remain {
				ts1[start] -= remain
				total_cnt += remain
				break
			} else {
				remain -= ts1[start]
				total_cnt += ts1[start]
				ts1[start] = 0
				start--
			}
		}
	}

	elapse := time.Since(stime)
	fmt.Printf("elapse time : %s\n", elapse)
	return total_cnt
}

func getHosting(ts, top []int) int {
	return getHosting2(ts, top)
}

// timestamp 에서 top 의 각 시간 안의 것을 삭제. 그리고 5 개 까지만.
func main() {
	ts := []int{0, 1, 1, 2, 4, 5}
	top := []int{5}
	fmt.Println(getHosting(ts, top))

	ts = []int{1, 2, 2, 3, 4, 5, 6, 6, 13, 16}
	top = []int{10, 15}
	fmt.Println(getHosting(ts, top))

	ts = []int{2, 2, 4, 8, 11, 28, 30}
	top = []int{0, 5, 5, 15}
	fmt.Println(getHosting(ts, top))

	ts = []int{0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	top = []int{6, 6, 6, 6}
	fmt.Println(getHosting(ts, top))

	ts = make([]int, 100000000)
	for i := 0; i < 10000000; i++ {
		ts[i] = rand.Intn(60)
	}
	top = make([]int, 3000)
	for i := 0; i < 3000; i++ {
		top[i] = rand.Intn(60)
	}
	fmt.Println(getHosting(ts, top))
}
