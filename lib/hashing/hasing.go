package hashing

import (
	"crypto/md5"
	"encoding/hex"
	"math/big"
)

func GenMd5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func GenMd5HashInt(text string) (*big.Int, error) {
	hasher := md5.New()
	hasher.Write([]byte(text))
	hrv := GenMd5Hash(text)
	i := new(big.Int)
	i.SetString(hrv, 16)
	return i, nil
}

func ModByString(s string, divisor int) int {
	m := big.NewInt(int64(divisor))
	i, _ := GenMd5HashInt(s)
	return int(new(big.Int).Mod(i, m).Int64())
}
