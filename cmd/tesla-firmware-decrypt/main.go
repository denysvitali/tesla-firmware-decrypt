package main

import (
	"encoding/base64"
	"github.com/alexflint/go-arg"
	"github.com/denysvitali/tesla-firmware-decrypt/pkg"
	"github.com/sirupsen/logrus"
	"os"
)

var args struct {
	InputFile string `arg:"positional,required"`
	Key       string `arg:"-k,required"`
}

var logger = logrus.New()

func main() {
	arg.MustParse(&args)
	f, err := os.Open(args.InputFile)
	if err != nil {
		logger.Fatalf("unable to open file: %v", err)
	}

	keyBytes, err := base64.StdEncoding.DecodeString(args.Key)
	if err != nil {
		logger.Fatalf("unable to decode base64 key: %v", err)
	}
	if len(keyBytes) != 32 {
		logger.Fatalf("invalid key size: expected 32 but received %d", len(keyBytes))
	}

	dec := decrypt.Decrypt{
		Key: keyBytes,
	}
	err = dec.Decrypt(f, os.Stdout)
	if err != nil {
		logger.Fatalf("unable to decrypt: %v", err)
	}
}
