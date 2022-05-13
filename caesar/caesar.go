package caesar

import (
	"fmt"
)

// Key is the number of character shifting
type Key byte

const alphanum byte = 'z' - 'a' + 1

// New Key is generated
func New(n byte) Key { return Key(n) }

// Encrypt plain message by using Caesar Crypto
func (key Key) Encrypt(plain []byte) (ret []byte, err error) {
	for _, b := range plain {
		base, err := choiseBaseChar(b)
		if err != nil {
			return nil, err
		}
		shift := (b-base+byte(key))%alphanum + base
		ret = append(ret, shift)
	}
	return
}

// Decrypt plain message by using Caesar Crypto
func (key Key) Decrypt(encrypted []byte) (ret []byte, err error) {
	return New(alphanum - byte(key)).Encrypt(encrypted)
}

func isLowerAlphabet(b byte) bool { return ('a' <= b && b <= 'z') }
func isUpperAlphabet(b byte) bool { return ('A' <= b && b <= 'Z') }
func isAlphabet(b byte) bool      { return isLowerAlphabet(b) || isUpperAlphabet(b) }

func choiseBaseChar(b byte) (byte, error) {
	if !isAlphabet(b) {
		return 0, fmt.Errorf("invalid character: %x", b)
	}
	if isLowerAlphabet(b) {
		return 'a', nil
	}
	return 'A', nil
}
