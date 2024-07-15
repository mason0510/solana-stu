// File: node/gossip.go

package node

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// GossipProtocol 实现了简化版的 Gossip 协议
type GossipProtocol struct {
	nodes       []*Node
	mutex       sync.RWMutex
	rumorsCache sync.Map
	logMutex    sync.Mutex
}

// NewGossipProtocol 创建一个新的 GossipProtocol 实例
func NewGossipProtocol() *GossipProtocol {
	return &GossipProtocol{
		nodes: make([]*Node, 0),
	}
}

// AddNode 添加一个节点到 Gossip 网络
func (gp *GossipProtocol) AddNode(node *Node) {
	gp.mutex.Lock()
	defer gp.mutex.Unlock()
	gp.nodes = append(gp.nodes, node)
}

// StartGossiping 开始 Gossip 过程
// 异步执行 不影响其他的查询过程
func (gp *GossipProtocol) StartGossiping() {
	go func() {
		for {
			gp.gossipRound()
			time.Sleep(1 * time.Second)
		}
	}()
}

// gossipRound 执行一轮 Gossip
func (gp *GossipProtocol) gossipRound() {
	gp.mutex.RLock()
	nodes := make([]*Node, len(gp.nodes))
	copy(nodes, gp.nodes)
	gp.mutex.RUnlock()

	for _, node := range nodes {
		go gp.gossipNode(node)
	}
}

// gossipNode 对特定节点执行 Gossip
func (gp *GossipProtocol) gossipNode(node *Node) {
	if len(gp.nodes) < 2 {
		return
	}

	// 随机选择一个其他节点
	target := gp.selectRandomNode(node)
	if target == nil {
		return
	}

	// 交换信息
	gp.exchangeInfo(node, target)
}

// selectRandomNode 随机选择一个不同的节点
func (gp *GossipProtocol) selectRandomNode(exclude *Node) *Node {
	gp.mutex.RLock()
	defer gp.mutex.RUnlock()

	candidates := make([]*Node, 0, len(gp.nodes)-1)
	for _, node := range gp.nodes {
		if node != exclude {
			candidates = append(candidates, node)
		}
	}

	if len(candidates) == 0 {
		return nil
	}

	return candidates[rand.Intn(len(candidates))]
}

// exchangeInfo 在两个节点之间交换信息
func (gp *GossipProtocol) exchangeInfo(node1, node2 *Node) {
	rumor := fmt.Sprintf("Rumor from %d to %d", node1.ID, node2.ID)

	_, loaded := gp.rumorsCache.LoadOrStore(rumor, true)
	if !loaded {
		gp.log("Gossip: %s\n", rumor)
	}
}

// GetRumorsCount 返回当前已传播的谣言数量
func (gp *GossipProtocol) GetRumorsCount() int {
	count := 0
	gp.rumorsCache.Range(func(key, value interface{}) bool {
		count++
		return true
	})
	return count
}

// log 用于安全地输出日志
func (gp *GossipProtocol) log(format string, args ...interface{}) {
	gp.logMutex.Lock()
	defer gp.logMutex.Unlock()
	fmt.Printf(format, args...)
}
