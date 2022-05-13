package scytale

import (
	"fmt"
	"math/rand"
	"time"
)

// Key of scytale crypto is circumference of scytale
type Key int

// New Key is generated
func New(c int) Key { return Key(c) }

// Encrypt plain words by using scytale crypto
func (key Key) Encrypt(plain []byte) ([]byte, error) {
	circum := int(key)
	length := len(plain) / circum
	if len(plain)%circum != 0 {
		length++
	}

	plain = append(plain, []byte(randString(length*circum-len(plain), alpha))...)

	encrypted := make([]byte, length*circum)
	for i, b := range plain {
		row := i % length
		col := i / length
		encrypted[row*circum+col] = b
	}
	return encrypted, nil
}

// Decrypt plain words by using scytale crypto
func (key Key) Decrypt(encrypted []byte) ([]byte, error) {
	circum := int(key)
	if len(encrypted)%circum != 0 {
		return nil, fmt.Errorf("encrypted is invalid length")
	}
	length := len(encrypted) / circum

	plain := make([]byte, length*circum)
	for i, b := range encrypted {
		row := i / circum
		col := i % circum
		plain[row+col*length] = b
	}
	return plain, nil
}

const (
	lower = "abcdefghijklmnopqrstuvwxyz"
	upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alpha = lower + upper
)

var rander = rand.New(rand.NewSource(time.Now().UnixMicro()))

func randString(n int, allowed string) string {
	var s string
	for i := 0; i < n; i++ {
		s += string(allowed[rander.Intn(len(allowed))])
	}
	return s
}

/*
0  6 12 18
1  7 13 19
2  8 14 20
3  9 15 21
4 10 16 22
5 11 17 23
*/
