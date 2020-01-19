package _map

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func count(s string, codeCount map[rune]int) {
	for _, r := range s {
		codeCount[r]++
	}
}

func TestCount(t *testing.T) {
	codeCount := map[rune]int{}
	count("가나다나", codeCount)
	if !reflect.DeepEqual(
		map[rune]int{'가': 1, '나': 2, '다': 1},
		codeCount,
	) {
		t.Error("codeCount mismatch:", codeCount)
	}
}

func ExampleCount() {
	codeCount := map[rune]int{}
	count("가나다나", codeCount)
	for _, key := range []rune{'가', '나', '다'} {
		fmt.Println(string(key), codeCount[key])
	}

	// Output:
	//가 1
	//나 2
	//다 1
}

func ExampleCountSort() {
	codeCount := map[rune]int{}
	count("가나다나", codeCount)

	// codeCount map 으로부터 key 들만 뽑아서 sort 를 위한 slice 에 담는다.
	var keys sort.IntSlice
	for key := range codeCount {
		keys = append(keys, int(key))
	}

	// 담아둔 slice 를 sorting 한다.
	sort.Sort(keys)

	// sorting 된 slice 에서 값(key)을 하나씩 가져와서 map 에서 찾는다.
	for _, key := range keys {
		fmt.Println(string(key), codeCount[rune(key)])
	}

	// map 의 value 만 가져와서 sorting 후 찍고 싶을 때
	var values sort.IntSlice
	for _, val := range codeCount {
		values = append(values, int(val))
	}
	sort.Sort(values)
	for _, val := range values {
		fmt.Println(val)
	}

	// Output:
	// 가 1
	// 나 2
	// 다 1
	// 1
	// 1
	// 2
}
