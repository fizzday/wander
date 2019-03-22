package main

import (
	"fmt"
	"sync"
)

type Manager struct{
	Note string
}

var m *Manager
var once sync.Once

func GetInstance() *Manager {
	once.Do(func() {
		m = &Manager{
			"单例实例化",
		}
	})
	return m
}

func (p *Manager) Manage() {
	p.Note = "1234"
}

func main() {
	ins := GetInstance()
	ins.Manage()
	fmt.Println(ins.Note)
}
