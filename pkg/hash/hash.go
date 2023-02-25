package hash

import (
	"crypto/rand"
	"crypto/sha512"
)

const (
	defaultSaltSize = 16
)

type Request struct {
	Data        []byte
	Salt        []byte
	PrivateSalt []byte
	Iterations  int
}

func (r Request) getPublicSalt() []byte {
	if len(r.Salt) > 0 {
		return r.Salt
	}
	if len(r.PrivateSalt) > 0 {
		return generateSalt()
	}
	return nil
}

func (r Request) getIterations() int {
	if r.Iterations > 0 {
		return r.Iterations
	}
	return 1
}

func Compute(request Request) ([]byte, []byte) {
	publicSalt := request.getPublicSalt()
	salt := append(request.PrivateSalt[:], publicSalt[:]...)

	hasher := sha512.New()

	// First iteration, with salt if present
	if len(salt) > 0 {
		hasher.Reset()
		hasher.Write(salt)
	}
	hasher.Write(request.Data)
	hashed := hasher.Sum(nil)

	// We already ran one iteration above
	iterations := request.getIterations() - 1

	// Run remaining iterations
	for i := 0; i < iterations; i++ {
		hasher.Reset()
		hasher.Write(hashed)
		hashed = hasher.Sum(nil)
	}

	return hashed, publicSalt
}

func generateSalt() []byte {
	var salt = make([]byte, defaultSaltSize)
	_, err := rand.Read(salt[:])
	if err != nil {
		panic(err)
	}
	return salt
}
