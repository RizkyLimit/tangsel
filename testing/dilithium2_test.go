package testing

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/cloudflare/circl/sign/dilithium/mode2"
	"github.com/stretchr/testify/require"
)

func TestDil(t *testing.T) {
	pk, sk, err := mode2.GenerateKey(rand.Reader)
	require.NoError(t, err, "Gagal menghasilkan key pair: %v", err)
	pkBytes, err := pk.MarshalBinary()
	require.NoError(t, err, "Gagal serialisasi public key: %v", err)
	skBytes, err := sk.MarshalBinary()
	require.NoError(t, err, "Gagal serialisasi secret key: %v", err)

	fmt.Printf("Public Key (%d bytes): %s\n\n\n", len(pkBytes), hex.EncodeToString(pkBytes))
	fmt.Printf("Secret Key (%d bytes): %s\n", len(skBytes), hex.EncodeToString(skBytes))
}
