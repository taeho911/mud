package utils

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"os"

	"golang.org/x/crypto/scrypt"
)

const (
	SCRYPT_N       int = 32768
	SCRYPT_R       int = 8
	SCRYPT_P       int = 1
	SCRYPT_KEY_LEN int = 32
)

func MakeRandom(length int) ([]byte, error) {
	random := make([]byte, length)
	_, err := rand.Read(random)
	return random, err
}

func HashPwd(pwd, salt []byte) ([]byte, error) {
	peppered, err := hmacSHA256(pwd, []byte(os.Getenv("MUD_PEPPER")))
	if err != nil {
		return nil, err
	}
	return scrypt.Key(peppered, salt, SCRYPT_N, SCRYPT_R, SCRYPT_P, SCRYPT_KEY_LEN)
}

func hmacSHA256(in, key []byte) ([]byte, error) {
	mac := hmac.New(sha256.New, key)
	_, err := mac.Write(in)
	return mac.Sum(nil), err
}

func ValidatePwd(pwd, salt, hashed []byte) (bool, error) {
	hashedPwd, err := HashPwd(pwd, salt)
	if err != nil {
		return false, err
	}
	if subtle.ConstantTimeCompare(hashedPwd, hashed) != 1 {
		return false, fmt.Errorf("invalid username or password")
	}
	return true, nil
}

func EncodeBase64(origin []byte) string {
	return base64.StdEncoding.EncodeToString(origin)
}

func DecodeBase64(str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(str)
}
