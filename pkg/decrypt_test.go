package decrypt_test

import (
	"bytes"
	"encoding/base64"
	"github.com/denysvitali/tesla-firmware-decrypt/pkg"
	"os"
	"testing"
)

func TestDecrypt_Decrypt(t *testing.T) {
	ciphertext, err := os.Open("./resources/1_0.bin")
	if err != nil {
		t.Fatalf("unable to open file: %v", err)
	}

	key := "uPaonw9KIv1tptQ3hsJsrR2Ibz3YhOcBRBaRitSi4Qg="
	keyBytes, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		t.Fatalf("unable to decode base64: %v", err)
	}
	d := decrypt.Decrypt{Key: keyBytes}

	buffer := bytes.NewBuffer(nil)
	err = d.Decrypt(ciphertext, buffer)

	if err != nil {
		t.Fatalf("unable to decrypt: %v", err)
	}

	expectedHeader := []byte("hsqs")
	header := make([]byte, len(expectedHeader))
	_, err = buffer.Read(header)
	if err != nil {
		t.Fatalf("unable to read into buffer: %v", err)
	}

	if !bytes.Equal(expectedHeader, header) {
		t.Fatalf("decryption failed: got %s but %s was expected", header, expectedHeader)
	}
}
