package cryptopals

import (
	"bytes"
)

func padPKCS7(msg []byte, blocksize int) []byte {
	pad_len := blocksize - len(msg) % blocksize
	return append(msg, bytes.Repeat([]byte{byte(pad_len)}, pad_len) ...)
}
