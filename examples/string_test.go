package examples

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestStrings(t *testing.T) {
	var buf []byte
	in := []byte(`p�ߺ�}@j���ɏv2`)
	encodeHex(buf, in)
	fmt.Println(buf)
}
func encodeHex(dst, uuid []byte) {
	hex.Encode(dst, uuid[:4])
	dst[8] = '-'
	hex.Encode(dst[9:13], uuid[4:6])
	dst[13] = '-'
	hex.Encode(dst[14:18], uuid[6:8])
	dst[18] = '-'
	hex.Encode(dst[19:23], uuid[8:10])
	dst[23] = '-'
	hex.Encode(dst[24:], uuid[10:])
}
