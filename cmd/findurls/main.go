package main

import (
	"fmt"
	"io/ioutil"

	"github.com/mvdan/xurls"
)

func main() {
	b, err := ioutil.ReadFile("src/duelyst.js")
	if err != nil {
		panic(err)
	}
	s := string(b)
	urls := xurls.Strict.FindAllString(s, -1)
	for _, u := range urls {
		fmt.Println(u)
	}
}
