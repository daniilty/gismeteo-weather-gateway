package requests

import (
	"bytes"
	"io"
)

func GetEmptyBody() io.Reader {
	emptyBytes := []byte{}

	return bytes.NewBuffer(emptyBytes)
}
