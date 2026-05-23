package main

import (
	"fmt"
	"sync"
	"time"
)

var wait sync.WaitGroup

func Shopping(name string) {
	fmt.Printf("%s 开始购物\n", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s 结束购物\n", name)
	wait.Done()
}

func main() {
	StartTime := time.Now()
	//现在的模式是购物接力
	//Shopping("张三")
	//Shopping("李四")
	//Shopping("王五")
	wait.Add(3)
	//主线程结束，协程函数跟着结束
	go Shopping("张三")
	go Shopping("李四")
	go Shopping("王五")
	wait.Wait()
	fmt.Println("购物完成", time.Since(StartTime))

}
