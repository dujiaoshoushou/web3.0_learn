package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/**
✅锁机制
题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ： sync.Mutex 的使用、并发数据安全。
题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ：原子操作、并发数据安全。
*/

func GetMutexNum() {
	num := 0
	lock := sync.Mutex{}
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				lock.Lock()
				num++
				lock.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Println(num)
}

func GetAtomicNum() {
	num := atomic.Int32{}
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				num.Add(1)
			}
		}()
	}

	wg.Wait()
	fmt.Println(num.Load())
}
