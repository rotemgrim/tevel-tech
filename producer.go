package main

import (
	"strconv"
	"sync"
	"time"
)

type producer struct {
	lock         sync.Mutex
	index        int
	listenerList []listener
}

type dataPacket struct {
	index int
	data  string
}

func (p *producer) produce(wg *sync.WaitGroup, howMany int) {
	wg.Add(1)
	defer wg.Done()

	for i := 0; i < howMany; i++ {
		// just for fake data creation
		data := strconv.Itoa(p.index) + "_" + randSeq(7)
		println("produced: " + data)

		// generate payload for sending
		packet := dataPacket{
			index: p.index,
			data: data,
		}

		saveIntoPersistentDB(packet)
		p.notifyAll(packet)

		p.index++
		time.Sleep(time.Second)
	}
}

func (p *producer) register(lis listener) {
	p.lock.Lock()
	p.listenerList = append(p.listenerList, lis)
	println(lis.name + " registered! | " + strconv.Itoa(lis.getLastIndex()))

	// this will send all the data prior to registration
	// that came after the lastIndex
	lastIndex := lis.getLastIndex()
	if p.index > lastIndex {
		for i, data := range db {
			if i > lastIndex {
				packet := dataPacket{index:i, data:data}
				p.notify(packet, lis)
			}
		}
	}
	p.lock.Unlock()
}

func (p *producer) unregister(lis listener) {
	p.lock.Lock()
	p.listenerList = removeFromSlice(p.listenerList, lis)
	println(lis.name + " removed! | " + strconv.Itoa(lis.getLastIndex()))
	p.lock.Unlock()
}

func (p *producer) notifyAll(packet dataPacket) {
	p.lock.Lock()
	for _, lis := range p.listenerList {
		p.notify(packet, lis)
	}
	p.lock.Unlock()
}

func (p *producer) notify(packet dataPacket, lis listener) {
	lis.ch <- packet
}

func removeFromSlice(listenerList []listener, listenerToRemove listener) []listener {
	listenerListLength := len(listenerList)
	for i, listener := range listenerList {
		if listenerToRemove.getId() == listener.getId() {
			listenerList[listenerListLength-1], listenerList[i] = listenerList[i], listenerList[listenerListLength-1]
			return listenerList[:listenerListLength-1]
		}
	}
	return listenerList
}
