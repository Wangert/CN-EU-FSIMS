package test

import (
	"CN-EU-FSIMS/utils/crypto"
	"fmt"
	"testing"
)

func TestSha256Hash(t *testing.T) {
	phash := "ecdf51e269847afe4f2709098dcbc7edbaafc1e9c87f53caebe47a28dd6f19c1"
	precc := "e6301f8e94258b246104a9bfa45066bf617ce0c83763a0a12fb24b363d3527fa"

	password := string(phash + precc)
	salt := "1111"

	passwordHash := crypto.CalculateSHA256(password, salt)
	fmt.Println("Password Hash: ", passwordHash)
}
