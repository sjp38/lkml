package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type rssItem struct {
	title	string
	author	string
	link	string
}

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

	var items []rssItem
	var item rssItem
	for _, line := range rsslines {
		txt := strings.TrimSpace(line)
		txt = strings.Replace(txt, "&lt;", "<", -1)
		txt = strings.Replace(txt, "&gt;", ">", -1)
		if txt == "<item>" {
			item = rssItem{}
			continue
		}
		if strings.HasPrefix(txt, "<title>") && strings.HasSuffix(txt,
			"</title>") {
			sidx := len("<title>")
			eidx := len(txt) - len("</title>")
			item.title = txt[sidx:eidx]
			continue
		}
		if strings.HasPrefix(txt, "<author>") && strings.HasSuffix(txt, "</author>") {
			sidx := len("<author>")
			eidx := len(txt) - len("</author>")
			item.author = txt[sidx:eidx]
			continue
		}
		if strings.HasPrefix(txt, "<link>") && strings.HasSuffix(txt, "</link>") {
			sidx := len("<link>")
			eidx := len(txt) - len("</link>")
			item.link = txt[sidx:eidx]
			continue
		}
		if txt == "</item>" {
			items = append(items, item)
		}
	}

	// 0th index is rss channel title. So, skip it.
	for i := len(items) - 1; i > 0; i-- {
		fmt.Printf("%s\n\t%s\n\t%s\n", items[i].title, items[i].author, items[i].link)
	}
}
