package examples

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransform(t *testing.T) {
	transformer := newTransformer()

	testCases := []struct {
		a int64
		b int64
	}{
		{a: 0, b: 0},
		{a: 100, b: 1608260665},
		{a: 1_000_000_000, b: 685885591},
		{a: 1_000_000_000 + 1, b: 809342380},
	}

	for _, pair := range testCases {
		y := transformer.modularTransform(pair.a)
		assert.EqualValues(t, pair.b, y)

		z := transformer.modularInverse(y)
		assert.Equal(t, pair.a, z)
	}
}

func TestTransformRandom(t *testing.T) {
	transformer := newTransformer()

	for i := 0; i < 1000; i++ {
		a := rand.Int63n(1<<31 - 1)
		b := transformer.modularTransform(a)
		c := transformer.modularInverse(b)
		assert.EqualValues(t, a, c)
	}
}

func TestTransformUniq(t *testing.T) {
	transformer := newTransformer()
	set := make(map[int64]int64)
	for i := 0; i < 1000; i++ {
		a := rand.Int63n(1<<31 - 1)
		if oldValue, ok := set[a]; !ok {
			set[a] = transformer.modularTransform(a)
		} else {
			assert.Equal(t, oldValue, transformer.modularTransform(a))
		}
	}
}
