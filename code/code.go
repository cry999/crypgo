package code

import (
	"strings"
)

// Key of code is a code book which has pairs of plain and encrypted word
type Key struct {
	encBook map[string]string
	decBook map[string]string
}

// New Key of code is generated
func New(encBook map[string]string) *Key {
	decBook := map[string]string{}
	for word, code := range encBook {
		decBook[code] = word
	}
	return &Key{
		encBook: encBook,
		decBook: decBook,
	}
}

// Encrypt plain contents by using code book
func (key *Key) Encrypt(plain []byte) ([]byte, error) {
	words := strings.Split(string(plain), " ")
	ret := make([]string, 0)
	for _, word := range words {
		var tail string
		switch word[len(word)-1:] {
		case ",", ".":
			tail = word[len(word)-1:]
			word = word[:len(word)-1]
		}
		if code, ok := key.encBook[word]; ok {
			ret = append(ret, code+tail)
		} else {
			ret = append(ret, word+tail)
		}
	}
	return []byte(strings.Join(ret, " ")), nil
}

// Decrypt plain contents by using code book
func (key *Key) Decrypt(encrypted []byte) ([]byte, error) {
	codes := strings.Split(string(encrypted), " ")
	ret := make([]string, 0)
	for _, code := range codes {
		var tail string
		switch code[len(code)-1:] {
		case ",", ".":
			tail = code[len(code)-1:]
			code = code[:len(code)-1]
		}
		if word, ok := key.decBook[code]; ok {
			ret = append(ret, word+tail)
		} else {
			ret = append(ret, code+tail)
		}
	}
	return []byte(strings.Join(ret, " ")), nil
}
