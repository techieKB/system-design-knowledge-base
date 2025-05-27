package main

import (
	"crypto/sha256"
	"fmt"
	"math/big"
)

var TOTAL_KEYS = big.NewInt(117);

func computeSlot(val string) uint64 {
	h := sha256.New()
	h.Write([]byte(val))
	hashDigest := new(big.Int).SetBytes(h.Sum(nil))
	return new(big.Int).Mod(hashDigest,TOTAL_KEYS).Uint64()
}

type NodeInfo struct {
	index uint64
	keys []string
}

type HashSpace struct{
	nodes []uint64
}

func (i *HashSpace) AppendNode(node string){
	slot := computeSlot(node)
	left := 0
	right := len(i.nodes)
	for left < right {
		mid := left + (right - left) / 2
		if(slot < i.nodes[mid]){
			right = mid;
		} else if(slot > i.nodes[mid]){
			left = mid + 1;
		} else {
			panic("Slot collision !!!")
		}
	}
	targetIdx := left
	i.nodes = append(i.nodes, 0)
	copy(i.nodes[targetIdx+1:],i.nodes[targetIdx:])
	i.nodes[targetIdx] = slot
}

func main() {
	info := &HashSpace{}
	info.AppendNode("A")
	info.AppendNode("E")
	fmt.Println(info.nodes)
}
