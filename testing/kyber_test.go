package testing

import (
	"fmt"
	"testing"

	"go.dedis.ch/kyber/v4/suites"
)

func TestKyber(t *testing.T) {
	s := suites.MustFind("Ed25519")
	x := s.Scalar().Zero()
	fmt.Println(x)
}
