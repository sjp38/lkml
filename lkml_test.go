package main

import (
	"reflect"
	"testing"
)

func TestItemsAfter(t *testing.T) {
	items := []rssItem{
		rssItem{
			title:  "abc",
			author: "foo",
			link:   "1",
		},
		rssItem{
			title:  "def",
			author: "foo2",
			link:   "2",
		},
		rssItem{
			title:  "ghi",
			author: "foo3",
			link:   "3",
		},
		rssItem{
			title:  "jkl",
			author: "foo4",
			link:   "4",
		},
	}

	after := itemsAfter(items, items[1])
	if !reflect.DeepEqual(after, items[2:]) {
		t.Error("expected %v but %v\n", items[2:], after)
	}

	after = itemsAfter(items, items[3])
	if !reflect.DeepEqual(after, []rssItem{}) {
		t.Error("expected %v but %v\n", []rssItem{}, after)
	}
}
