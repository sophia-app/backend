package utils

import (
	"os"

	myhash "github.com/sophia-app/backend/pkg/hash"
)

// GetHashEncrypter returns a new Argon2HashEncrypter with the salt length and memory parameters.
func GetHashEncrypter() *myhash.Argon2HashEncrypter {
	saltLen := uint32(len([]byte(GetHashSalt())))

	return myhash.New(1, saltLen, 64*1024, 4, 32)
}

// GetHashSalt returns the hash salt from the environment variable.
func GetHashSalt() string {
	return os.Getenv("HASH_SALT")
}
