package chain

import (
	"sync"
)

// Sealevel 模拟 Sealevel 并行执行机制
func Sealevel(transactions []string) {
	var wg sync.WaitGroup
	for _, tx := range transactions {
		wg.Add(1)
		go func(t string) {
			defer wg.Done()
			// 模拟交易执行
			_ = t
		}(tx)
	}
	wg.Wait()
}
