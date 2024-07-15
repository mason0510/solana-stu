package node

import (
	"fmt"
	"strings"
)

type Node struct {
	ID         int
	Children   []*Node
	Data       []byte
	ShardIndex int // 新增：分片索引
}

func NewNode(id int) *Node {
	return &Node{
		ID:       id,
		Children: []*Node{},
	}
}

func (n *Node) AddChild(child *Node) {
	n.Children = append(n.Children, child)
}

func (n *Node) PropagateShards(shards [][]byte, startIndex int) {
	if len(shards) == 0 {
		return
	}

	n.Data = shards[0]
	n.ShardIndex = startIndex
	fmt.Printf("Node %d received shard %d: %s\n", n.ID, startIndex, shards[0])

	if len(shards) > 1 && len(n.Children) > 0 {
		remainingShards := shards[1:]
		childCount := len(n.Children)
		for i, child := range n.Children {
			start := i * len(remainingShards) / childCount
			end := (i + 1) * len(remainingShards) / childCount
			if start < len(remainingShards) {
				child.PropagateShards(remainingShards[start:end], startIndex+1+start)
			}
		}
	}
}

func (n *Node) PrintTree(depth int) {
	fmt.Printf("%s- Node %d\n", strings.Repeat("  ", depth), n.ID)
	for _, child := range n.Children {
		child.PrintTree(depth + 1)
	}
}
