package main

import (
	"fmt"
	"sync"
)

type listener struct {
	lock      sync.Mutex
	id        string
	lastIndex int
	name      string
	ch        chan dataPacket
}

func createListener(name string) listener {
	lis := listener{
		id:        randSeq(10),
		name:      name,
		ch:        make(chan dataPacket, 1),
	}
	setIndex(name, -1)
	go lis.startListening()
	return lis
}

func (l *listener) getId() string {
	return l.id
}

func (l *listener) getLastIndex() int {
	return getIndex(l.name)
}

func (l *listener) setLastIndex(i int) {
	setIndex(l.name, i)
}

func (l *listener) startListening() {
	for {
		select {
		case packet := <-l.ch:
			fmt.Println(l.name + " Received data: " + packet.data)
			l.lock.Lock()
			if getIndex(l.name) < packet.index {
				setIndex(l.name, packet.index)
			}
			l.lock.Unlock()
		}
	}
}
