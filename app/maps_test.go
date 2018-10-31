package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveBlueCars(t *testing.T) {
	m := map[string]string{"vaz": "blue", "volvo": "yellow", "bmw": "blue"}
	for key, value := range m {
		if value == "blue" {
			delete(m, key)
		}
	}

	assert.Equal(t, map[string]string{"volvo": "yellow"}, m)
}

func TestNotFoundByKeyReturnsDefaultValue(t *testing.T) {
	mp := make(map[int]string)
	mp[1] = "A"
	test, _ := mp[2]
//	assert.True(t, ok)
	assert.Equal(t, "", test)
}
