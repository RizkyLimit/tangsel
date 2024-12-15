package testing

import (
	"crypto/ecdsa"
	"encoding/hex"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestKeyWallet(t *testing.T) {

	wallet := []struct {
		PrivateHex string
		PublicHex  string
		Address    string
	}{
		{
			PrivateHex: "0xa02ee917d9a4bebecec78fa5365f34ee4174cc5c1e23c09d72e2819039e1deb8",
			PublicHex:  "0x04e1915b7e45b98d964c91ddf3e88a95881cbfe166072389e42372edca0e5af4a4f0e0b428b2b0b57dd6a2bfaf45583cebde1b1ecb9bc2315996ff0d589e962644",
			Address:    "0x76B192706d573Bf903D78065D6B8108804349c4a",
		},
		{
			PrivateHex: "0xa23e5eca301e6c7d1ef92a03113fb40ba413c69269a51a8785c09494f074e9ad",
			PublicHex:  "0x048f499665758e4bd9022646330b13e0d6982474830eacbd57fd12762557c0511dbebba881f9063f8191c63fce8f384664c9a2fa9fa6da17d4461f7fbbfa3a755d",
			Address:    "0x1Ea49C983433A3aC3d84f46DBd8981641a284B2b",
		},
	}

	var privateKeys []*ecdsa.PrivateKey
	var publicKeys []*ecdsa.PublicKey

	for i := range wallet {
		wallet[i].PrivateHex = strings.TrimPrefix(wallet[i].PrivateHex, "0x")
		wallet[i].PublicHex = strings.TrimPrefix(wallet[i].PublicHex, "0x")
	}

	t.Run("DecodePrivateKey", func(t *testing.T) {
		for i := range wallet {
			privateKeyBytes, err := hex.DecodeString(wallet[i].PrivateHex)
			require.NoError(t, err, "failed to decode private key hex: %v", err)
			privateKey, err := crypto.ToECDSA(privateKeyBytes)
			require.NoError(t, err, "failed to convert to ECDSA private key: %v", err)
			privateKeys = append(privateKeys, privateKey)
		}
	})

	t.Run("DecodePublicKey", func(t *testing.T) {
		for i := range wallet {
			publicKeyBytes, err := hex.DecodeString(wallet[i].PublicHex)
			require.NoError(t, err, "failed to decode private key hex: %v", err)
			publicKey, err := crypto.UnmarshalPubkey(publicKeyBytes)
			require.NoError(t, err, "failed to unmarshal public key: %v", err)
			publicKeys = append(publicKeys, publicKey)
		}
	})

	t.Run("MatchingKey", func(t *testing.T) {
		for i := 0; i < len(privateKeys); i++ {
			if privateKeys[i].PublicKey.X.Cmp(publicKeys[i].X) != 0 || privateKeys[i].PublicKey.Y.Cmp(publicKeys[i].Y) != 0 {
				t.Errorf("public key does not match private key")
			} else {
				derivedAddress := crypto.PubkeyToAddress(*publicKeys[i]).Hex()
				assert.Equal(t, derivedAddress, wallet[i].Address)
			}

		}
	})
}
