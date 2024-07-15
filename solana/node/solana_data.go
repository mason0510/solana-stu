package node

//
//import (
//	"fmt"
//	"math"
//	"math/rand"
//	"time"
//)
//
//// Node 代表网络中的一个节点
//type Node struct {
//	ID       int
//	Children []*Node
//	Data     []byte
//}
//
//// Network 代表整个网络
//type Network struct {
//	Nodes []*Node
//}
//
//// 创建一个新的网络
//func NewNetwork(nodeCount int) *Network {
//	network := &Network{
//		Nodes: make([]*Node, nodeCount),
//	}
//	for i := 0; i < nodeCount; i++ {
//		network.Nodes[i] = &Node{ID: i}
//	}
//	return network
//}
//
//// 构建树状结构
//func (n *Network) BuildTree(fanout int) {
//	for i, node := range n.Nodes {
//		for j := 1; j <= fanout; j++ {
//			childIndex := i*fanout + j
//			if childIndex < len(n.Nodes) {
//				node.Children = append(node.Children, n.Nodes[childIndex])
//			}
//		}
//	}
//}
//
//// 模拟数据传播
//func (n *Network) PropagateData(data []byte, shardSize int) {
//	// 将数据分片
//	shards := splitData(data, shardSize)
//
//	// 从根节点开始传播
//	n.propagateShards(n.Nodes[0], shards)
//}
//
//// 递归传播分片
//func (n *Network) propagateShards(node *Node, shards [][]byte) {
//	if len(shards) == 0 {
//		return
//	}
//
//	// 模拟网络延迟
//	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
//
//	// 给当前节点分配一个分片
//	node.Data = shards[0]
//	fmt.Printf("Node %d received shard: %v\n", node.ID, node.Data)
//
//	// 将剩余分片分配给子节点
//	remainingShards := shards[1:]
//	if len(node.Children) > 0 && len(remainingShards) > 0 {
//		shardsPerChild := len(remainingShards) / len(node.Children)
//		if shardsPerChild == 0 {
//			shardsPerChild = 1
//		}
//		for i, child := range node.Children {
//			start := i * shardsPerChild
//			if start >= len(remainingShards) {
//				break
//			}
//			end := (i + 1) * shardsPerChild
//			if end > len(remainingShards) {
//				end = len(remainingShards)
//			}
//			n.propagateShards(child, remainingShards[start:end])
//		}
//	}
//}
//
//// 将数据分片
//func splitData(data []byte, shardSize int) [][]byte {
//	shardCount := int(math.Ceil(float64(len(data)) / float64(shardSize)))
//	shards := make([][]byte, shardCount)
//
//	for i := 0; i < shardCount; i++ {
//		start := i * shardSize
//		end := (i + 1) * shardSize
//		if end > len(data) {
//			end = len(data)
//		}
//		shards[i] = data[start:end]
//	}
//
//	return shards
//}
