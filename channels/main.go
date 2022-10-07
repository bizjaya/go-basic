package main

import (
	"fmt"
	"time"
)


func main() {

    channel()
//	buffChannel()
	//for multi-channnel queue
//	rangeSelect()

	//fmt.Print("Hlello ", test, "w orld")
}

func channel(){

	ch := make(chan string)

	go func() {
		fmt.Println(time.Now(), "taking a nap")

		time.Sleep(2 * time.Second)

		ch <- "hello"
	}()

	fmt.Println(time.Now(), "waiting for message")

	v := <- ch
	
	fmt.Println(time.Now(), "received", v)

}

func buffChannel(){
	ch := make(chan int, 2)

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(time.Now(), i, "sending")
			ch <- i
			fmt.Println(time.Now(), i, "sent")
		}

		// XXX: There could be cases where this message is not completed,
		// this is solved in futured examples
		fmt.Println(time.Now(), "all completed")
	}()

	time.Sleep(2 * time.Second)

	fmt.Println(time.Now(), "waiting for messages")

	fmt.Println(time.Now(), "received", <-ch)
	fmt.Println(time.Now(), "received", <-ch)
	fmt.Println(time.Now(), "received", <-ch)

	fmt.Println(time.Now(), "exiting")
}

func rangeSelect(){
	ch := make(chan int, 2)
	exit := make(chan struct{})

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(time.Now(), i, "sending")
			ch <- i
			fmt.Println(time.Now(), i, "sent")

			time.Sleep(1 * time.Second)
		}

		fmt.Println(time.Now(), "all completed, leaving")

		close(ch)
	}()

	go func() {
		// XXX: This is overcomplicated because is only channel only, "select"
		// shines when using multiple channels.
		for {
			select {
			case v, open := <-ch:
				if !open {
					close(exit)
					return
				}

				fmt.Println(time.Now(), "received", v)
			}
		}

		// XXX: In cases where only one channel is used
		// for v := range ch {
		// 	fmt.Println(time.Now(), "received", v)
		// }

		// close(exit)
	}()

	fmt.Println(time.Now(), "waiting for everything to complete")

	<-exit

	fmt.Println(time.Now(), "exiting")
}