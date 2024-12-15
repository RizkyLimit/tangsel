package wallet

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

func GenerateWallet() string {
	privateKey, err := crypto.GenerateKey() // Generate a new private key using ECDSA SECP256k1
	if err != nil {
		log.Fatalf("Failed to generate private key: %v", err)
	}

	privateKeyHex := fmt.Sprintf("0x%x", crypto.FromECDSA(privateKey)) // Convert private key to bytes (hex format to string)
	publicKeyECDSA, ok := privateKey.Public().(*ecdsa.PublicKey)       // Get the public key
	if !ok {
		log.Fatalf("Failed to cast public key to ECDSA")
	}

	publicKeyHex := fmt.Sprintf("0x%x", crypto.FromECDSAPub(publicKeyECDSA))           // Convert public key to bytes (hex format to string)
	compressPublicKeyHex := fmt.Sprintf("0x%x", crypto.CompressPubkey(publicKeyECDSA)) // Convert public key to Compress Public Key (hex format to string)
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()                           // Generate address from public key
	return "===== ECDSA SECP256k1 Key Pair =====\nPrivate Key: " + privateKeyHex + "\nPublic Key: " + publicKeyHex + "\nCompress Public Key: " + compressPublicKeyHex + "\nEthereum Address: " + address
}
