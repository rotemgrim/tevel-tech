package main

import (
	"math/rand"
	"sync"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}


var db = map[int]string{}
func saveIntoPersistentDB(packet dataPacket) {
	db[packet.index] = packet.data
}

var indexes = map[string]int{}
var mutex = new(sync.Mutex)
func setIndex(listenerName string, i int) {
	mutex.Lock()
	indexes[listenerName] = i
	mutex.Unlock()
}
func getIndex(listenerName string) int {
	return indexes[listenerName]
}