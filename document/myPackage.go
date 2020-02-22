// 이 패키지는 go 언어의 문서 작성 규칙이 얼마나 간단한지 보여주기 위한 것이다.
// 아주아주아주 단순하다.
package document

// PI 는 상수이자 전역변수이다.
// 이런 부분은 코드 자체로 이해할 수 있다면 굳이 주석을 달지 않다도 된다.
const PI = 3.14159265

// GetStrLen() 함수는 string 의 길이를 구하는 함수이다.
func GetStrLen(s string) int {
	if s == "" {
		return 0
	}
	n := 0
	for range s {
		n++
	}
	return n
}

// GetDouble() 함수는 주어진 정수의 제곱을 리턴한다.
func GetDouble(n int) int {
	return n * n
}
