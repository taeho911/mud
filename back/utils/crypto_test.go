package utils

import (
	"os"
	"testing"
)

func TestGenSalt(t *testing.T) {
	length := 16
	salt, err := GenSalt(length)
	if err != nil {
		t.Fatalf("genSalt failed. err = %v", err)
	}
	if len(salt) != length {
		t.Fatalf("genSalt failed. len(salt) = %v", len(salt))
	}
}

func TestHmacSHA256(t *testing.T) {
	password := []byte("testP@sswr0d")
	pepper := []byte("testPPePPer")
	peppered, err := hmacSHA256(password, pepper)
	if err != nil {
		t.Fatalf("hmacSHA256 failed. err = %v", err)
	}
	if len(peppered) != 32 {
		t.Fatalf("hmacSHA256 failed. len(peppered) = %v", len(peppered))
	}
}

func TestHashPwd(t *testing.T) {
	password := []byte("testP@sswr0d")
	os.Setenv("MUD_PEPPER", "testPPePPer")
	salt, _ := GenSalt(16)
	hashed, err := HashPwd(password, salt)
	if err != nil {
		t.Fatalf("HashPwd failed. err = %v", err)
	}
	if len(hashed) != 32 {
		t.Fatalf("HashPwd failed. len(hashed) = %v", len(hashed))
	}
}

func TestValidatePwd(t *testing.T) {
	password1 := []byte("TestValidatePwd_1")
	password2 := []byte("TestValidatePwd_2")
	os.Setenv("MUD_PEPPER", "testPPePPer")
	salt, _ := GenSalt(16)
	hashed, _ := HashPwd(password1, salt)

	// 동일한 패스워드를 비교할 때 false가 돌아오면 비교 실패
	if result, err := ValidatePwd(password1, salt, hashed); err != nil {
		t.Fatalf("ValidatePwd failed. err = %v", err)
	} else if !result {
		t.Fatalf("ValidatePwd failed. result = %v", result)
	}

	// 동일하지 않은 패스워드를 비교할 때
	if result, err := ValidatePwd(password2, salt, hashed); err == nil {
		t.Fatalf("ValidatePwd failed. err = %v", err)
	} else if result {
		t.Fatalf("ValidatePwd failed. result = %v", result)
	}
}

func TestEncodeAndDecodeBase64(t *testing.T) {
	origin := "My name is Kim"
	encoded := EncodeBase64([]byte(origin))
	decoded, err := DecodeBase64(encoded)
	if err != nil {
		t.Fatalf("DecodeBase64 failed. err = %v", err)
	}
	if origin != string(decoded) {
		t.Fatalf("origin = %v, decoded = %v", origin, string(decoded))
	}
}
