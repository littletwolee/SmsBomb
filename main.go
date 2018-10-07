package main

import (
	"application"
)

const VERSION = "1.0.1"

var app *application.App

type queue []int

func main() {

	app = application.NewApp("test", "/Users/wuqimeng/go/go_project/src/SmsBomb/config.toml")
	var q queue
	q.push(1)
	q.push(2)
	q.push(3)

	print(q.pop())

}

func test() int {
	lastOccurred := make(map[byte]int)
	s := "abcbaf"
	b := []byte(s)
	maxLen := 0
	start := 0
	for i, ch := range b {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLen
}

func mapKeyToString(data map[string]int) string {
	s := ""
	for k, _ := range data {
		s += k
	}
	return s

}

func (q *queue) push(v int) {
	*q = append(*q, v)
}

func (q *queue) pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}
