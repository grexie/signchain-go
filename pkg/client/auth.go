package client

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"encoding/binary"
	"strings"
	"time"
)

type AuthSecretKey string
type AuthSignature string

func (s AuthSignature) String() string {
	return string(s)
}

func (k AuthSecretKey) Sign(timestamp time.Time, data []byte) (AuthSignature, error) {
	var nonce [32]byte
	var tb [8]byte
	binary.BigEndian.PutUint64(tb[:], uint64(timestamp.UnixMilli()))

	if _, err := rand.Read(nonce[:]); err != nil {
		return "", err
	} else {
		hash := sha256.New()
		hash.Write(data)
		hash.Write(nonce[:])
		hash.Write(tb[:])
		hash.Write([]byte(k))
		
		signature := strings.ToLower(base32.StdEncoding.EncodeToString(nonce[:])) + "." + strings.ToLower(base32.StdEncoding.EncodeToString(tb[:])) + "." + strings.ToLower(base32.StdEncoding.EncodeToString(hash.Sum([]byte{})))
		
		return AuthSignature(signature), nil
	}
}