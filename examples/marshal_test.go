package examples

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Warehouse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func TestMarshal(t *testing.T) {
	w := &Warehouse{
		ID:   1,
		Name: "Имя",
	}
	bytes, err := json.Marshal(w)
	assert.Nil(t, err)
	fmt.Println(string(bytes))

	result := &Warehouse{}
	err = json.Unmarshal(bytes, result)
	assert.Nil(t, err)

	assert.Equal(t, 1, result.ID)
}
