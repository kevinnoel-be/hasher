package hash

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeSimple(t *testing.T) {
	request := Request{
		Data:        []byte("password"),
		Salt:        []byte("public"),
		PrivateSalt: []byte("private"),
	}
	hash, salt := Compute(request)
	assert.Equal(t, "3A7L1NrrEtrJ6Ivs6YxMq/cxMCvZDAo4A+E11zVxMtYMUGEgc8a1Ecte9gE+ty0FKTJsjUq1vO321yg+bu8lDw==", b64(hash))
	assert.Equal(t, b64(request.Salt), b64(salt))
}

func TestComputeSimpleWithIterations(t *testing.T) {
	request := Request{
		Data:        []byte("password"),
		Salt:        []byte("public"),
		PrivateSalt: []byte("private"),
		Iterations:  10000,
	}
	hash, salt := Compute(request)
	assert.Equal(t, "kxUoqUqScwZvJSReRRSHvLNMOt7YUZs+eOeTMlWnaeQ5kEEIsX1W28CYx/PCqarGCxfdFpHo2O4fTnojD1W4IA==", b64(hash))
	assert.Equal(t, b64(request.Salt), b64(salt))
}

func TestComputePublicSalt(t *testing.T) {
	request := Request{
		Data: []byte("password"),
		Salt: []byte("public"),
	}
	hash, salt := Compute(request)
	assert.Equal(t, "njSCr4cxHp7tfIJjD7RfqCEESxAs50tb/CgaKGOdgxjvAymux3m17GY9k29bIbJLAaAIKRrn52Vq4c98JSed4g==", b64(hash))
	assert.Equal(t, b64(salt), b64(request.Salt))
}

func TestComputeNoSalts(t *testing.T) {
	request := Request{
		Data: []byte("password"),
	}
	hash, salt := Compute(request)
	assert.Equal(t, "sQnzu7wkTrgkQZF+0G1hi5AI3Qmzvv0bXgc5THBqi7mAsdd4Xll27ASbRt9fEyavWi6m0QP9B8lThf+rDKy8hg==", b64(hash))
	assert.Nil(t, salt)
}

func b64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}
