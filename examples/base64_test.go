package examples

import (
	"encoding/base64"
	"fmt"
	"strings"
	"testing"
)

func TestBase64(t *testing.T) {
	input := "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAGQAAABkCAYA"

	b64data := input[strings.IndexByte(input, ',')+1:]
	fmt.Println(b64data)

	data, err := base64.StdEncoding.DecodeString(b64data)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(data)
}
