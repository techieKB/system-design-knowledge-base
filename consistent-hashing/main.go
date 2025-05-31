package main

import (
	"crypto/sha256"
	"fmt"
	"math/big"
)

var TOTAL_KEYS = big.NewInt(1117);

func computeSlot(val string) uint64 {
	h := sha256.New()
	h.Write([]byte(val))
	hashDigest := new(big.Int).SetBytes(h.Sum(nil))
	return new(big.Int).Mod(hashDigest,TOTAL_KEYS).Uint64()
}

type NodeInfo struct {
	name string
	index uint64
	keys []string
}

type HashSpace struct{
	nodes []*NodeInfo
}

func (i *HashSpace) AppendNode(node string){
	slot := computeSlot(node)
	fmt.Printf("Node %v, appending to %v \n",node,slot)
	targetIdx := i.findSlotForNode(slot)
	i.nodes = append(i.nodes, nil)
	copy(i.nodes[targetIdx+1:],i.nodes[targetIdx:])
	i.nodes[targetIdx] = &NodeInfo{
		index: slot,
		name:  node,
		keys: []string{},
	}
}

func (i *HashSpace) AssignKey(key string){
	slot := computeSlot(key)
	fmt.Printf("Requested key %v, assigned slot %v \n",key,slot)
	idx := i.findClosestStorageNode(slot)
	i.nodes[idx].keys = append(i.nodes[idx].keys, key)
}

func(i *HashSpace) GetStorageNode(key string) *NodeInfo {
	slot := computeSlot(key)
	return i.nodes[i.findClosestStorageNode(slot)]
}

func (i *HashSpace) findSlotForNode(slot uint64) uint64 {
	left := 0
	right := len(i.nodes)
	for left < right {
		mid := left + (right - left) / 2
		if slot < i.nodes[mid].index {
			right = mid;
		} else if slot > i.nodes[mid].index {
			left = mid + 1;
		} else {
			panic("Slot collision !!!")
		}
	}
	return uint64(left)
}

func (i *HashSpace) findClosestStorageNode(slot uint64) uint64 {
	left := 0
	right := len(i.nodes)
	for left < right {
		mid := left + (right - left) / 2
		if slot > i.nodes[mid].index {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if right == len(i.nodes) {
		right = 0
	}
	return uint64(right)
}

func main() {
	info := &HashSpace{}
	info.AppendNode("node-1")
	info.AppendNode("node-2")
	info.AppendNode("node-3")
	info.AssignKey("file1.txt")
	info.AssignKey("file2.txt")
	for _,node := range info.nodes {
		fmt.Printf("Node %v, Index %v, Keys: %v \n",node.name,node.index,node.keys)
	}
}