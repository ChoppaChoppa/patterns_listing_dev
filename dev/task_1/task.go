package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

func main() {
	GetTime("0.beevik-ntp.pool.ntp.org")
}

func GetTime(host string){
	time, err := ntp.Time(host)
	if err != nil {
		code, _ := fmt.Fprintln(os.Stderr, err)
		os.Exit(code)
	}

	fmt.Println(time)
}