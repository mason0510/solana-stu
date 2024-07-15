package chain

import "sync"

// GulfStream 模拟 Gulf Stream 机制
type GulfStream struct {
	transactions []string
	mu           sync.Mutex
}

func NewGulfStream() *GulfStream {
	return &GulfStream{
		transactions: make([]string, 0),
	}
}

func (gs *GulfStream) AddTransaction(tx string) {
	gs.mu.Lock()
	defer gs.mu.Unlock()
	gs.transactions = append(gs.transactions, tx)
}

func (gs *GulfStream) GetTransactions() []string {
	gs.mu.Lock()
	defer gs.mu.Unlock()
	return gs.transactions
}
