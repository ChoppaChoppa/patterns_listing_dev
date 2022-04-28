package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Wgetter struct {
	Url     string
	Body    string
	Headers http.Header
	Code    int
}

func main() {
	flag.Parse()

	wg := wget(flag.Args()...)
	fmt.Println(wg)
}

func wget(urls ...string) []Wgetter {
	wgetters := make([]Wgetter, len(urls))

	for i, v := range urls {
		resp, errGet := http.Get(v)
		if errGet != nil {
			fmt.Println(errGet)
			return nil
		}

		bodyB, errRead := ioutil.ReadAll(resp.Body)
		if errRead != nil {
			fmt.Println(errRead)
			return nil
		}

		wgetter := Wgetter{
			Url:     v,
			Body:    string(bodyB),
			Headers: resp.Header,
			Code:    resp.StatusCode,
		}

		wgetters[i] = wgetter
		resp.Body.Close()
	}

	return wgetters
}
