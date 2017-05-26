package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Printf("Hello\n")
	resp, err := http.Get("https://lkml.org/rss.php")
	if err != nil {
		panic(fmt.Sprintf("failed to get rss: %s", err))
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		panic(fmt.Sprintf("failed to read body: %s", err))
	}

	fmt.Printf(string(body))
}
