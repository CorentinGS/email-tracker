package crypto

import (
	"crypto/rand"
	"encoding/gob"
	"io"
	"log/slog"
	"os"

	"github.com/pkg/errors"
	"golang.org/x/crypto/ed25519"
)

type KeyManager struct {
	seed       io.Reader
	privateKey ed25519.PrivateKey
	publicKey  ed25519.PublicKey
}

type KeyManagerOption func(*KeyManager)

func WithSeed(seed io.Reader) KeyManagerOption {
	return func(km *KeyManager) {
		km.seed = seed
	}
}

func WithKeys(publicKey ed25519.PublicKey, privateKey ed25519.PrivateKey) KeyManagerOption {
	return func(km *KeyManager) {
		km.publicKey = publicKey
		km.privateKey = privateKey
	}
}

func NewKeyManager(options ...KeyManagerOption) (*KeyManager, error) {
	publicKey, privateKey, err := retrieveKeys()
	if err != nil {
		slog.Debug("Error retrieving keys", slog.Any("error", err))
		publicKey, privateKey, err = GenerateKeyPair()
		if err != nil {
			slog.Error("Error generating keys", slog.Any("error", err))
			return nil, err
		}
	}

	km := &KeyManager{
		seed:       rand.Reader,
		privateKey: privateKey,
		publicKey:  publicKey,
	}

	for _, option := range options {
		option(km)
	}

	return km, nil
}

func (k *KeyManager) GetPrivateKey() ed25519.PrivateKey {
	return k.privateKey
}

func (k *KeyManager) GetPublicKey() ed25519.PublicKey {
	return k.publicKey
}

func (k *KeyManager) ParseEd25519Key() error {
	publicKey, privateKey, err := GenerateKeyPair()
	if err != nil {
		return err
	}

	k.privateKey = privateKey
	k.publicKey = publicKey

	return nil
}

func GenerateKeyPair() (ed25519.PublicKey, ed25519.PrivateKey, error) {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, nil, errors.Wrap(err, "Error generating keys")
	}

	// Store keys in a .gob file
	err = storeKeys(publicKey, privateKey)
	if err != nil {
		return nil, nil, err
	}

	return publicKey, privateKey, nil
}

func storeKeys(publicKey ed25519.PublicKey, privateKey ed25519.PrivateKey) error {
	// Open the .gob file for writing
	file, err := os.Create("./keys.gob")
	if err != nil {
		return err
	}
	defer file.Close()

	// Create an encoder
	encoder := gob.NewEncoder(file)

	// Encode and write the keys to the file
	err = encoder.Encode(publicKey)
	if err != nil {
		return err
	}

	err = encoder.Encode(privateKey)
	if err != nil {
		return err
	}

	return nil
}

func retrieveKeys() (ed25519.PublicKey, ed25519.PrivateKey, error) {
	// Open the .gob file for reading
	file, err := os.Open("./keys.gob")
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	// Create a decoder
	decoder := gob.NewDecoder(file)

	// Decode and retrieve the keys from the file
	var publicKey ed25519.PublicKey
	err = decoder.Decode(&publicKey)
	if err != nil {
		return nil, nil, err
	}

	var privateKey ed25519.PrivateKey
	err = decoder.Decode(&privateKey)
	if err != nil {
		return nil, nil, err
	}

	slog.Debug("Keys retrieved successfully")
	return publicKey, privateKey, nil
}
