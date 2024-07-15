package chain

// Turbine 模拟 Turbine 区块传播机制
func Turbine(block []byte, validators int) {
	chunkSize := len(block) / validators
	for i := 0; i < validators; i++ {
		start := i * chunkSize
		end := (i + 1) * chunkSize
		if i == validators-1 {
			end = len(block)
		}
		go func(chunk []byte) {
			// 模拟发送块的一部分到验证者
			_ = chunk
		}(block[start:end])
	}
}
