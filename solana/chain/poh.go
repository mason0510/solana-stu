package chain

import (
	"crypto/sha256"
	"encoding/binary"
	"time"
)

// ProofOfHistory 模拟 PoH 机制
func ProofOfHistory(duration time.Duration) []byte {
	start := time.Now()
	hash := sha256.Sum256([]byte("initial"))
	for time.Since(start) < duration {
		hash = sha256.Sum256(hash[:])
	}
	return hash[:]
}

type PoHState struct {
	PreviousHash [32]byte
	Count        uint64
	Timestamp    time.Time
}

func NewPoHState() *PoHState {
	return &PoHState{
		PreviousHash: sha256.Sum256([]byte("initial")),
		Count:        0,
		Timestamp:    time.Now(),
	}
}

func (s *PoHState) Tick() {
	s.Count++
	s.Timestamp = time.Now()

	data := make([]byte, 40)
	copy(data[:32], s.PreviousHash[:])
	binary.LittleEndian.PutUint64(data[32:], s.Count)

	s.PreviousHash = sha256.Sum256(data)
}

/*
*
这段时间看到的所有时钟状态.
所有节点都可以参考这个进行
*/
func SimulatePoH(duration time.Duration) []PoHState {
	state := NewPoHState()
	var history []PoHState

	start := time.Now()
	for time.Since(start) < duration {
		state.Tick()
		history = append(history, *state)
		time.Sleep(time.Millisecond)
	}

	return history
}

func VerifyPoH(history []PoHState) bool {
	for i := 1; i < len(history); i++ {
		prev := history[i-1]
		curr := history[i]

		if curr.Count != prev.Count+1 {
			return false
		}

		data := make([]byte, 40)
		copy(data[:32], prev.PreviousHash[:])
		binary.LittleEndian.PutUint64(data[32:], curr.Count)

		if sha256.Sum256(data) != curr.PreviousHash {
			return false
		}

		if !curr.Timestamp.After(prev.Timestamp) {
			return false
		}
	}
	return true
}
