//go:build base58
// +build base58

package main

import (
	"fmt"
	"solana/address"
)

func main() {
	pair, err := address.GenerateSolanaKeyPair()
	if err != nil {
		fmt.Printf("生成密钥对时出错: %v\n", err)
		return
	}

	fmt.Println("成功生成 Solana 密钥对:")
	fmt.Printf("公钥 (原始字节): %v\n", pair.PublicKey)
	fmt.Printf("私钥 (原始字节): %v\n", pair.PrivateKey)
	fmt.Printf("公钥 (Hex): %s\n", pair.PublicKeyHex)
	fmt.Printf("私钥 (Hex): %s\n", pair.PrivateKeyHex)
	fmt.Printf("公钥 (Base58): %s\n", pair.PublicKeyBase58)
	fmt.Printf("私钥 (Base58): %s\n", pair.PrivateKeyBase58)
	fmt.Printf("Solana 地址: %s\n", pair.SolanaAddress)
}
