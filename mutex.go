package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var ticket = 10 //10张票

var mutex sync.Mutex // 创建锁头

var wg sync.WaitGroup

func main() {
	/**
	4个goroutine，模拟4个售票口

	在使用互斥锁的时候，对资源操作完，一定要解锁，否则会出现程序异常，死锁等问题
	defer语句
	*/
	wg.Add(4)
	go saleTickets("售票口1")
	go saleTickets("售票口2")
	go saleTickets("售票口3")
	go saleTickets("售票口4")

	wg.Wait()

	//time.Sleep(5 * time.Second)
}

func saleTickets(name string) {
	rand.Seed(time.Now().UnixNano())
	defer wg.Done()
	for {
		//上锁
		mutex.Lock()
		if ticket > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(name, "售出", ticket)
			ticket--
		} else {
			mutex.Unlock() // 条件不满足，也要解锁
			fmt.Println(name, "售罄，没有票了。。。")
			break
		}
		mutex.Unlock() //解锁
	}
}
