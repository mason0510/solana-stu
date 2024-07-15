package node

import (
	_ "fmt"
)

type Network struct {
	Nodes []*Node
	Root  *Node
}

func NewNetwork(size int) *Network {
	network := &Network{
		Nodes: make([]*Node, size),
	}
	for i := 0; i < size; i++ {
		network.Nodes[i] = NewNode(i)
	}
	return network
}

func (n *Network) BuildTree(fanout int) {
	n.Root = n.Nodes[0]
	queue := []*Node{n.Root}
	currentIndex := 1

	for len(queue) > 0 && currentIndex < len(n.Nodes) {
		parent := queue[0]
		queue = queue[1:]

		for i := 0; i < fanout && currentIndex < len(n.Nodes); i++ {
			child := n.Nodes[currentIndex]
			parent.AddChild(child)
			queue = append(queue, child)
			currentIndex++
		}
	}
}

func (n *Network) PropagateData(data []byte, shardSize int) {
	shards := splitData(data, shardSize)
	n.Root.PropagateShards(shards, 0)
}

func splitData(data []byte, shardSize int) [][]byte {
	shardCount := (len(data) + shardSize - 1) / shardSize
	shards := make([][]byte, shardCount)

	for i := 0; i < shardCount; i++ {
		start := i * shardSize
		end := start + shardSize
		if end > len(data) {
			end = len(data)
		}
		shards[i] = data[start:end]
	}

	return shards
}

func (n *Network) PrintTree() {
	n.Root.PrintTree(0)
}
