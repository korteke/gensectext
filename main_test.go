package main

import (
	"strings"
	"testing"
	"time"

	"github.com/ProtonMail/gopenpgp/v2/helper"
)

func TestExpiresTime(t *testing.T) {
	tim, _ := time.Parse("2006-01-02", "2099-12-12")

	if !tim.After(time.Now()) {
		t.Errorf("Expires date is not in the future.")
	}

	if tim.String() != "2099-12-12 00:00:00 +0000 UTC" {
		t.Error("Expected 2099-12-12 got: ", tim.String())
	}

}

func TestGenerateKey(t *testing.T) {
	pass := []byte{42, 42, 42, 42, 42}
	name := "Security Example CO"
	email := "security@test.test"

	rsaKey, _ := helper.GenerateKey(name, email, pass, keyType, rsaBits)
	if !strings.HasPrefix(rsaKey, "-----BEGIN PGP PRIVATE KEY BLOCK-----") {
		t.Errorf("The RSA key generator did not produce a correctly formatted key.")
	}

	if !strings.HasSuffix(rsaKey, "-----END PGP PRIVATE KEY BLOCK-----") {
		t.Errorf("The RSA key generator did not produce a correctly formatted key.")
	}
}
