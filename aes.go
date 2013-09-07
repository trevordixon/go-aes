package aes

import ()

type aes struct {
	key         []byte
	expandedKey []byte
}

func New(key []byte) (*aes, error) {
	a := &aes{key: key}

	return a, nil
}
