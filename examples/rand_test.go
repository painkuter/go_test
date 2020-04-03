package examples

import (
	"crypto/rand"
	"fmt"
	"testing"
)

func TestGetRand(t *testing.T) {
	r, _ := rand.Prime(rand.Reader, 40)
	fmt.Println(r.Int64() % 10)
}
