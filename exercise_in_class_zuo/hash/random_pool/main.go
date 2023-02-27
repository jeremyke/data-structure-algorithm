package main

import (
	"math/rand"
	"time"
)

/**
【题目】设计一种结构，在该结构中有如下三个功能:
	insert(key):将某个key加入到该结构，做到不重复加入
	delete(key):将原本在结构中的某个key移除
	getRandom(): 等概率随机返回结构中的任何一个key。
【要求】Insert、delete和getRandom方法的时间复杂度都是O(1)
*/

type randomPool struct {
	keyIndexMap map[string]int
	indexKeyMap map[int]string
	size        int
}

func (rp *randomPool) insert(insertKey string) {
	if _, ok := rp.keyIndexMap[insertKey]; !ok {
		rp.size++
		rp.keyIndexMap[insertKey] = rp.size
		rp.indexKeyMap[rp.size] = insertKey
	}
}

func (rp *randomPool) delete(delKey string) {
	if index, ok := rp.keyIndexMap[delKey]; ok {
		var (
			lasKey   = rp.size
			lasValue = rp.indexKeyMap[rp.size]
		)
		rp.indexKeyMap[index] = lasValue
		delete(rp.indexKeyMap, lasKey)
		rp.keyIndexMap[lasValue] = index
		delete(rp.keyIndexMap, delKey)
	}
}

func (rp *randomPool) getRandom() string {
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(rp.size + 1)
	return rp.indexKeyMap[randomNum]
}
