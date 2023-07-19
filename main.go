package main

// Fact はnの階乗を生成する関数である
func Fact(num int) int {
	f := 1
	for i := num; i >= 1; i-- {
		f *= i
	}
	return f
}

func Fact2(num int) int {
	if num <= 1 {
		return 1
	} else {
		return num * Fact2(num-1)
	}
}