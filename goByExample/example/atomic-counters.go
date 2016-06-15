package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

//为了模拟并发更新，我们启动 50 个 Go 协程，
//对计数器每隔 1ms （译者注：应为非准确时间）进行一次加一操作。
func main() {
	var ops uint64 = 0

	for i := 0; i < 50; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1)

				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
