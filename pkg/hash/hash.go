package hash

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"errors"

	"golang.org/x/crypto/argon2"
)

// HashSalt struct used to store generated hash and salt used to generate the hash.
type HashSalt struct {
	Hash, Salt string
}

type Argon2HashEncrypter struct {
	// time represents the number of passed over the specified memory.
	time    uint32
	// cpu memory to be used.
	memory  uint32
	// threads for parallelism aspect of the algorithm.
	threads uint8
	// keyLen of the generate hash key.
	keyLen  uint32
	// saltLen the length of the salt used.
	saltLen uint32
}

// New constructor function for Argon2HashEncrypter.
func New(time, saltLen uint32, memory uint32, threads uint8, keyLen uint32) *Argon2HashEncrypter {
	return &Argon2HashEncrypter{
		time:    time,
		saltLen: saltLen,
		memory:  memory,
		threads: threads,
		keyLen:  keyLen,
	}
}

// GenerateHash using the password and provided salt.
// If no salt value is provided fallback to random value generated of a given length.
func (a *Argon2HashEncrypter) GenerateHash(password, salt string) (*HashSalt, error) {
	var err error

	// Convert password and salt to byte slices.
	passwordBytes := []byte(password)
	saltBytes := []byte(salt)

	// If salt is not provided generate a salt of the configured salt length.
	if len(saltBytes) == 0 {
		saltBytes, err = RandomSecret(a.saltLen)
	}
	if err != nil {
		return nil, err
	}

	// Generate hash
	hash := argon2.IDKey(passwordBytes, saltBytes, a.time, a.memory, a.threads, a.keyLen)

	// Encode the hash and salt using base64 for readable output
	encodedHash := base64.StdEncoding.EncodeToString(hash)
	encodedSalt := string(saltBytes)

	// Return the generated hash and salt used for storage.
	return &HashSalt{Hash: encodedHash, Salt: encodedSalt}, nil
}

// Compare generated hash with store hash.
func (a *Argon2HashEncrypter) Compare(hash, salt, password string) error {
	// Generate hash for comparison.
	hashSalt, err := a.GenerateHash(password, salt)
	if err != nil {
		return err
	}

	// Compare the generated hash with the stored hash.
	// If they don't match return error.
	if !bytes.Equal([]byte(hash), []byte(hashSalt.Hash)) {
		return errors.New("hash doesn't match")
	}

	return nil
}

// RandomSecret generates a random secret of a given length.
func RandomSecret(length uint32) ([]byte, error) {
	secret := make([]byte, length)

	_, err := rand.Read(secret)
	if err != nil {
		return nil, err
	}

	return secret, nil
}
