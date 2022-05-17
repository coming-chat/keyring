//go:build darwin && cgo
// +build darwin,cgo

package keyring

import (
	"errors"
)

type keychain struct {
	path    string
	service string

	passwordFunc PromptFunc

	isSynchronizable         bool
	isAccessibleWhenUnlocked bool
	isTrusted                bool
}

var ErrKeyChainUnsupported = errors.New("Keychain unsupported")

func init() {
}

func (k *keychain) Get(key string) (Item, error) {
	return Item{}, ErrKeyChainUnsupported
}

func (k *keychain) GetMetadata(key string) (Metadata, error) {
	return Metadata{}, ErrKeyChainUnsupported
}

func (k *keychain) Set(item Item) error {
	return ErrKeyChainUnsupported
}

func (k *keychain) Remove(key string) error {
	return ErrKeyChainUnsupported
}

func (k *keychain) Keys() ([]string, error) {
	return []string{}, ErrKeyChainUnsupported
}
