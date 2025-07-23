package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
✅Goroutine
题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
考察点 ： go 关键字的使用、协程的并发执行。
题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
考察点 ：协程原理、并发任务调度。
*/

func PrintNums() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func(arrs []int) {
		defer wg.Done()
		for _, num := range arrs {
			if num%2 == 0 {
				fmt.Println("偶数：", num)
			}
			time.Sleep(10 * time.Millisecond)
		}
	}(nums)

	go func(arrs []int) {
		defer wg.Done()
		for _, num := range arrs {
			if num%2 != 0 {
				fmt.Println("奇数：", num)
			}
			time.Sleep(100 * time.Millisecond)
		}
	}(nums)
	wg.Wait()
}
func InitFuncs(a int, b int) int {
	return a + b
}

func TriggerTask(funs []func(int, int) int) {
	wg := sync.WaitGroup{}
	wg.Add(len(funs))
	for _, f := range funs {
		num1 := rand.Intn(100)
		num2 := rand.Intn(100)
		go func(a int, b int) {
			defer wg.Done()
			start := time.Now()
			result := f(a, b)
			end := time.Now()
			elapsed := end.Sub(start)
			fmt.Println("结果：", result, "耗时：", elapsed)
		}(num1, num2)
	}
}

func TriggerTask2(funs []func(int, int) int) {
	wg := sync.WaitGroup{}
	wg.Add(len(funs))
	for _, f := range funs {
		num1 := rand.Intn(100)
		num2 := rand.Intn(100)
		go fmt.Println(f(num1, num2))

		defer wg.Done()
	}
}
