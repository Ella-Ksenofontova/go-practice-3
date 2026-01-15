package main

import "fmt"

func work(x int) int {
	if (x%5 == 0 || x%4 == 0) && x > 0 {
		return x + 1
	}
	return x - 1
}

type CacheMap struct {
	cache map[int]int
}

func (cm CacheMap) getValueOfWork(key int) int {
	val, ok := cm.cache[key]
	if !ok {
		val = work(key)
		cm.cache[key] = val
	}
	return val
}

func main() {
	cm := CacheMap{make(map[int]int)}
	var a int

	fmt.Println("Введите 10 целых чисел")
	for range 10 {
		fmt.Scan(&a)
		val := cm.getValueOfWork(a)

		fmt.Print(val, " ")
	}

	print("time limit ok")
}
