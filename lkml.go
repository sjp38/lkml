package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	resp, err := http.Get("https://lkml.org/rss.php")
	if err != nil {
		panic(fmt.Sprintf("failed to get rss: %s", err))
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		panic(fmt.Sprintf("failed to read body: %s", err))
	}

	rsslines := strings.Split(string(body), "\n")

	var titles []string
	for _, line := range rsslines {
		txt := strings.TrimSpace(line)
		if strings.HasPrefix(txt, "<title>") && strings.HasSuffix(txt,
			"</title>") {
			sidx := len("<title>")
			eidx := len(txt) - len("</title>")
			titles = append(titles, txt[sidx:eidx])
		}
	}

	// 0th index is rss channel title. So, skip it.
	for i := len(titles) - 1; i > 0; i-- {
		fmt.Printf("%s\n", titles[i])
	}
}
