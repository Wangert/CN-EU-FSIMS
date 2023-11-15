package crypto

import (
	"crypto/sha256"
	"encoding/hex"
)

func CalculateSHA256(s, salt string) string {
	hasher := sha256.New()
	hasher.Write([]byte(salt + s))

	return hex.EncodeToString(hasher.Sum(nil))
}
