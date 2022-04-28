package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	var timeout int
	var host string
	var port string

	flag.IntVar(&timeout, "timeout", 10, "timout")
	flag.StringVar(&host, "host", "", "host")
	flag.StringVar(&port, "port", "", "port")

	flag.Parse()

	addr := host + ":" + port

	fmt.Println(timeout)
	connect(addr, time.Duration(timeout))
}

func connect(addr string, timeout time.Duration) {
	conn, errConn := net.DialTimeout("tcp", addr, timeout*time.Second)
	if errConn != nil {
		fmt.Println("conn")
		fmt.Println(errConn)
		os.Exit(0)
	}
	defer func() {
		fmt.Println("defer")
		conn.Close()
	}()

	go func() {
		for {
			buf := make([]byte, 1024)
			_, errRead := conn.Read(buf)
			if errRead != nil {
				fmt.Println(errRead)
				os.Exit(0)
			}
			fmt.Println(string(buf))
		}
	}()

	for {
		var source string
		fmt.Scanln(&source)

		_, errWrite := conn.Write([]byte(source))
		if errWrite != nil {
			fmt.Println(errWrite)
			os.Exit(0)
		}
	}
}
