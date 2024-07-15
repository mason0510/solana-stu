package address

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"

	"github.com/mr-tron/base58"
)

// SolanaKeyPair 结构体用于存储 Solana 密钥对和地址信息
type SolanaKeyPair struct {
	PublicKey        ed25519.PublicKey
	PrivateKey       ed25519.PrivateKey
	PublicKeyHex     string
	PrivateKeyHex    string
	PublicKeyBase58  string
	PrivateKeyBase58 string
	SolanaAddress    string
}

// GenerateSolanaKeyPair 生成 Solana 密钥对和地址
func GenerateSolanaKeyPair() (*SolanaKeyPair, error) {
	publicKey, privateKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to generate ed25519 key pair: %v", err)
	}

	publicKeyHex := hex.EncodeToString(publicKey)
	privateKeyHex := hex.EncodeToString(privateKey)
	publicKeyBase58 := base58.Encode(publicKey)
	privateKeyBase58 := base58.Encode(privateKey)
	solanaAddress := publicKeyBase58 // Solana 地址就是公钥的 Base58 编码

	return &SolanaKeyPair{
		PublicKey:        publicKey,
		PrivateKey:       privateKey,
		PublicKeyHex:     publicKeyHex,
		PrivateKeyHex:    privateKeyHex,
		PublicKeyBase58:  publicKeyBase58,
		PrivateKeyBase58: privateKeyBase58,
		SolanaAddress:    solanaAddress,
	}, nil
}
