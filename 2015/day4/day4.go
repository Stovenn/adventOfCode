package day4

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

const secretKey = "bgvyzdsv"

var lowest int = 0

func Run() {
	// secret + number = hash
	h := md5.New()

	for {
		h.Reset()
		h.Write([]byte(fmt.Sprintf("%s%d", secretKey, lowest)))
		b := h.Sum(nil)
		str := hex.EncodeToString(b)
		if strings.HasPrefix(str, "000000") {
			break
		}
		lowest++
	}

	fmt.Println("Lowest decimal:", lowest)
}
