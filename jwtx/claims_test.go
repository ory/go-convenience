package jwtx

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestParseMapStringInterfaceClaims(t *testing.T) {
	assert.EqualValues(t, &Claims{
		JTI:       "jti",
		Subject:   "sub",
		Issuer:    "iss",
		Audience:  []string{"aud"},
		ExpiresAt: time.Unix(1234, 0),
		IssuedAt:  time.Unix(1234, 0),
		NotBefore: time.Unix(1234, 0),
	}, ParseMapStringInterfaceClaims(map[string]interface{}{
		"jti": "jti",
		"aud": "aud",
		"iss": "iss",
		"sub": "sub",
		"exp": 1234,
		"iat": 1234,
		"nbf": 1234,
	}))

	assert.EqualValues(t, &Claims{
		Audience:  []string{"aud", "dua"},
		ExpiresAt: time.Unix(1234, 0),
		IssuedAt:  time.Unix(1234, 0),
		NotBefore: time.Unix(1234, 0),
	}, ParseMapStringInterfaceClaims(map[string]interface{}{
		"aud": []string{"aud", "dua"},
		"exp": 1234,
		"iat": 1234,
		"nbf": 1234,
	}))
}
