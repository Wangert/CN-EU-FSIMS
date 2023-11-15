package test

import (
	"CN-EU-FSIMS/utils/crypto"
	"fmt"
	"testing"
)

func TestSha256Hash(t *testing.T) {
	password := "password765"
	salt := "SALT"

	passwordHash := crypto.CalculateSHA256(password, salt)
	fmt.Println("Password Hash: ", passwordHash)
}
