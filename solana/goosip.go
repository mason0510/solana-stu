//go:build node
// +build node

package main

import (
	"fmt"
	"solana/node"
	"time"
)

// 提供最终会一致性 牺牲即时的强一致性
func main() {
	gossip := node.NewGossipProtocol()

	// 添加节点到 Gossip 网络
	for i := 0; i < 10; i++ {
		newNode := &node.Node{ID: i}
		gossip.AddNode(newNode)
	}

	// 启动 Gossip 过程
	gossip.StartGossiping()

	// 运行一段时间后
	time.Sleep(10 * time.Second)
	fmt.Printf("Total rumors spread: %d\n", gossip.GetRumorsCount())
}
