package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/crypto/argon2"
	"strings"
)

// Parameter untuk algoritma Argon2
type hashParams struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

// GeneratePasswordHash melakukan hashing password dengan Argon2id
func GeneratePasswordHash(password string) (string, error) {
	// Parameter yang direkomendasikan untuk Argon2id
	p := &hashParams{
		memory:      64 * 1024, // 64MB
		iterations:  3,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}

	// Generate salt acak
	salt := make([]byte, p.saltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	// Hashing password dengan Argon2id
	hash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	// Format hasil hash: $argon2id$v=19$m=65536,t=3,p=2$salt$hash
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// Encode parameter dan hasil hash dalam format string
	encodedHash := fmt.Sprintf("$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s",
		p.memory, p.iterations, p.parallelism, b64Salt, b64Hash)

	return encodedHash, nil
}

// VerifyPassword membandingkan password dengan hash yang tersimpan
func VerifyPassword(password, encodedHash string) (bool, error) {
	// Parsing hash string
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return false, errors.New("format hash tidak valid")
	}

	// Validasi format hash
	if vals[1] != "argon2id" {
		return false, errors.New("algoritma hashing tidak didukung")
	}

	var version int
	_, err := fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return false, err
	}
	if version != 19 {
		return false, errors.New("versi argon2id tidak didukung")
	}

	// Parsing parameter
	p := &hashParams{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.memory, &p.iterations, &p.parallelism)
	if err != nil {
		return false, err
	}

	// Decode salt
	salt, err := base64.RawStdEncoding.DecodeString(vals[4])
	if err != nil {
		return false, err
	}
	p.saltLength = uint32(len(salt))

	// Decode hash
	hash, err := base64.RawStdEncoding.DecodeString(vals[5])
	if err != nil {
		return false, err
	}
	p.keyLength = uint32(len(hash))

	// Hitung hash ulang dari password yang diberikan
	compareHash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	// Bandingkan hash (menggunakan perbandingan waktu konstan untuk mencegah timing attack)
	return subtle.ConstantTimeCompare(hash, compareHash) == 1, nil
}
