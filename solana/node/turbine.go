package node

import (
	"fmt"
)

func RunTurbineSimulation() {
	// 创建一个有 50 个节点的网络
	network := NewNetwork(50)

	// 构建一个扇出度为 3 的树
	network.BuildTree(3)

	// 模拟大块数据
	data := []byte("This is a large block of data that needs to be propagated through the network quickly and efficiently. " +
		"We are adding more text to make the data larger and test the system more thoroughly.")

	// 传播数据，每个分片大小为 20 字节
	network.PropagateData(data, 20)

	// 打印网络树结构
	fmt.Println("\nNetwork Tree Structure:")
	network.PrintTree()

	// 验证数据完整性
	// 验证数据完整性
	reconstructedData := make([]byte, 0, len(data))
	for i := 0; i < len(network.Nodes); i++ {
		for _, node := range network.Nodes {
			if node.ShardIndex == i {
				reconstructedData = append(reconstructedData, node.Data...)
				break
			}
		}
	}

	fmt.Println("\nOriginal data:", string(data))
	fmt.Println("Reconstructed data:", string(reconstructedData))

	if string(data) == string(reconstructedData) {
		fmt.Println("Data successfully reconstructed!")
	} else {
		fmt.Println("Data reconstruction failed.")
		// 打印每个节点的详细信息以进行调试
		for _, node := range network.Nodes {
			if len(node.Data) > 0 {
				fmt.Printf("Node %d: Index %d, Data: %s\n", node.ID, node.ShardIndex, string(node.Data))
			}
		}
	}

}
