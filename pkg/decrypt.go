package decrypt

import (
	"encoding/binary"
	"fmt"
	"golang.org/x/crypto/salsa20"
	"io"
)

type Decrypt struct {
	Key []byte
}

func (d *Decrypt) Decrypt(reader io.Reader, writer io.Writer) error {
	if len(d.Key) != 32 {
		return fmt.Errorf("key must be 32 bytes long")
	}
	blockNumber := 0

	nonceBytes := make([]byte, 24)
	block := make([]byte, 256)
	outBlock := make([]byte, 256)
	for {
		binary.LittleEndian.PutUint64(nonceBytes, uint64(blockNumber))
		blockBytes, err := reader.Read(block)
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return fmt.Errorf("unable to read file: %v", err)
		}
		salsa20.XORKeyStream(outBlock, block, nonceBytes, (*[32]byte)(d.Key))
		blockNumber++
		_, err = writer.Write(outBlock[0:blockBytes])
		if err != nil {
			return fmt.Errorf("unable to write: %v", err)
		}
	}
}
