// channel_test project channel_test.go
package channelLearn

import (
	"fmt"
)

func test(c chan int) {

	for i := 0; i < 10; i++ {
		c <- i
	}

	close(c)
}

func Channel_One() {

	c := make(chan int)

	go test(c)

	<-c

	for v := range c {
		fmt.Println(v)
	}

}
