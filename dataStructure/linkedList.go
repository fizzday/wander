package main

type linkList interface {
	Length() int
	Lpush(data int) *node
	Rpush(data int) *node
}
type node struct {
	data int
	next *node
}
