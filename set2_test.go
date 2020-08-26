package cryptopals

import (
	"bytes"
	"testing"
)

func TestProblem9(t *testing.T) {
	res := padPKCS7([]byte("YELLOW SUBMARINE"), 20)
	if !bytes.Equal(res, []byte("YELLOW SUBMARINE\x04\x04\x04\x04")) {
		t.Error("Wrong answer", res)
	}
}
