////go:build solana
//// +build solana
//
//package main
//
///*
//*
//模拟并行处理 然后广播
//*/
//func main() {
//	// 模拟 PoH
//	poh := chain.ProofOfHistory(time.Millisecond * 400)
//	fmt.Printf("PoH hash: %x\n", poh)
//
//	// 模拟 Gulf Stream
//	gs := chain.NewGulfStream()
//	gs.AddTransaction("tx1")
//	gs.AddTransaction("tx2")
//	gs.AddTransaction("tx3")
//
//	// 模拟 Sealevel
//	transactions := gs.GetTransactions()
//	chain.Sealevel(transactions)
//
//	// 模拟 Turbine
//	block := []byte("simulated block data")
//	chain.Turbine(block, 4)
//
//	fmt.Println("Solana simulation completed")
//}
//go:build solana
// +build solana

package main

import (
	"fmt"
	"solana/chain"
	"time"
)

func main() {
	fmt.Println("Starting Solana simulation...")

	// 模拟 PoH
	duration := time.Millisecond * 400
	pohHistory := chain.SimulatePoH(duration)
	fmt.Printf("Generated %d PoH entries\n", len(pohHistory))
	fmt.Printf("First PoH entry: %+v\n", pohHistory[0])
	fmt.Printf("Last PoH entry: %+v\n", pohHistory[len(pohHistory)-1])

	if chain.VerifyPoH(pohHistory) {
		fmt.Println("PoH history verified successfully")
	} else {
		fmt.Println("PoH history verification failed")
	}

	// 模拟 Gulf Stream
	gs := chain.NewGulfStream()
	gs.AddTransaction("tx1")
	gs.AddTransaction("tx2")
	gs.AddTransaction("tx3")

	// 模拟 Sealevel
	transactions := gs.GetTransactions()
	chain.Sealevel(transactions)

	// 模拟 Turbine
	block := []byte("simulated block data")
	chain.Turbine(block, 4)

	fmt.Println("Solana simulation completed")
}
