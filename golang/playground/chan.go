package playground

import (
	"fmt"
	"time"
)

func testChannelSend() {
	fmt.Println("=============== test channel send ==================")

	c := make(chan int)
	defer close(c)
	go func() {
		c <- 3 + 4
	}()
	i := <-c
	fmt.Println(i)

	fmt.Println("====================================================")
}

func sendSumToChannel(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func testChannelBlock() {
	fmt.Println("=============== test channel block ==================")

	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sendSumToChannel(s[:len(s)/2], c)
	go sendSumToChannel(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c
	fmt.Println(x, y, x+y)

	fmt.Println("====================================================")
}


func testChannelRange() {
	fmt.Println("=============== test channel range ==================")

	go func() {
		time.Sleep(1 * time.Hour)
	}()
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i = i + 1 {
			c <- i
		}
		close(c)
	}()
	for i := range c {  // iter until c is closed
		fmt.Println(i)
	}
	fmt.Println("Finished")

	fmt.Println("====================================================")
}

//select语句选择一组可能的send操作和receive操作去处理。
//它类似switch,但是只是用来处理通讯(communication)操作。
//它的case可以是send语句，也可以是receive语句，亦或者default。
//receive语句可以将值赋值给一个或者两个变量。它必须是一个receive操作。
//最多允许有一个default case,它可以放在case列表的任何位置，尽管我们大部分会将它放在最后。
func testChannelSelect() {
	fmt.Println("=============== test channel select ==================")

	c := make(chan int)
	quit := make(chan int)

	fib := func() {
		x, y := 0, 1
		//如果有同时多个case去处理,比如同时有多个channel可以接收数据，
		//那么Go会伪随机的选择一个case处理(pseudo-random)。
		//如果没有case需要处理，则会选择default去处理， 如果default case存在的情况下。
		//如果没有default case，则select语句会阻塞，直到某个case需要处理。
		for {
			select {
			case c <- x:
				x, y = y, x+y
			case <-quit:
				fmt.Println("quit")
				return
			}
		}
	}
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fib()

	fmt.Println("====================================================")
}



func testChannelTimeout() {
	fmt.Println("=============== test channel timeout ==================")

	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	// blocked neither no value pushed to c1 nor timeout less than 1 second
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	fmt.Println("====================================================")
}


func testChannelTimer() {
	fmt.Println("=============== test channel timer ==================")

	fmt.Println("Timer 1 start")
	timer1 := time.NewTimer(time.Second * 2)
	<-timer1.C
	fmt.Println("Timer 1 expired")

	fmt.Println("Timer 2 start")
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C  // no panic if stop is called
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()  // stop immediately
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	fmt.Println("====================================================")
}


func testChannelTicker() {
	fmt.Println("=============== test channel ticker ==================")

	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		fmt.Println("Ticker 1 start")
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()
	fmt.Println("Timer 1 start")
	timer := time.NewTimer(time.Second * 3)
	<-timer.C  // yield to ticker goroutine
	fmt.Println("Timer 1 stop")
	ticker.Stop()
	fmt.Println("Ticker 1 stop")

	fmt.Println("====================================================")
}


func testChannelClose() {
	fmt.Println("=============== test channel close ==================")

	go func() {
		time.Sleep(time.Hour)
	}()
	c := make(chan int, 10)
	c <- 1
	c <- 2
	close(c)  // cannot send to channel anymore
	// panic
	//	c <- 3

	// output

	i, ok := <-c
	fmt.Printf("%d, %v\n", i, ok)  // 1, true
	fmt.Println(<-c)  // 2
	i, ok = <-c
	fmt.Printf("%d, %v\n", i, ok)  // 0, false
	fmt.Println(<-c)  // 0

	//for i := range c {  // 1, 2 only
	//	fmt.Println(i)
	//}

	fmt.Println("====================================================")
}


func testChannelSync() {
	fmt.Println("=============== test channel close ==================")

	worker := func(done chan bool) {
		time.Sleep(time.Second)
		// 通知任务已完成
		done <- true
	}
	done := make(chan bool, 1)
	go worker(done)
	// 等待任务完成
	fmt.Println("worker start")
	<-done
	fmt.Println("worker stop")

	fmt.Println("====================================================")
}


func TestChannel() {
	// https://www.runoob.com/w3cnote/go-channel-intro.html
	//testChannelSend()
	//testChannelBlock()
	//testChannelRange()
	//testChannelSelect()
	//testChannelTimeout()
	//testChannelTimer()
	//testChannelTicker()
	testChannelClose()
	testChannelSync()

}