package main

import (
	"fmt"
	"time"
)

func main() {
	<-OrChanel(signal(2*time.Hour),
		signal(5*time.Minute),
		signal(10*time.Second),
		signal(1*time.Second),
		signal(1*time.Minute),
	)

	fmt.Println("end")
}

func signal(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func OrChanel(ch ...<-chan interface{}) <-chan interface{} {
	var receiving = make(chan interface{})
	for i, v := range ch {
		go func(chanel <-chan interface{}, index int) {
			select {
			case <-chanel:
				fmt.Println("chan send:", index)
				receiving <- chanel
			}

		}(v, i)
	}

	return receiving
}
