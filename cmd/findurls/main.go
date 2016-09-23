// findurls prints the unique URLs found in Duelyst's source code.
// This program expects to be executed from the root of this repo.
package main

import (
	"fmt"
	"io/ioutil"

	"github.com/mvdan/xurls"
)

func main() {
	bytes, err := ioutil.ReadFile("src/duelyst.js")
	if err != nil {
		panic(err)
	}
	urls := xurls.Strict.FindAllString(string(bytes), -1)
	uniqueURLs := make(map[string]bool)
	for _, url := range urls {
		_, ok := uniqueURLs[url]
		if !ok {
			fmt.Println(url)
			uniqueURLs[url] = true
		}
	}
}
