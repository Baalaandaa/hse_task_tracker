package hash

import (
	"crypto/sha512"
	"fmt"
)

func Hash(in string) string {
	h := sha512.New()
	h.Write([]byte("task_salt_" + in + "_task_salt"))
	return fmt.Sprintf("%x", h.Sum(nil))
}
