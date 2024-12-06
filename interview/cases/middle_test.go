package cases

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIslands(t *testing.T) {
	t.Run("3", func(t *testing.T) {
		input := [][]int{
			{1, 0, 0, 1},
			{0, 0, 1, 1},
			{1, 0, 0, 1},
			{1, 0, 1, 1},
		}
		assert.Equal(t, 3, countIslands(input))
	})
	t.Run("1", func(t *testing.T) {
		input := [][]int{
			{1, 1, 1, 1},
			{0, 0, 0, 1},
			{1, 0, 0, 1},
			{1, 1, 1, 1},
		}
		assert.Equal(t, 1, countIslands(input))
	})
	t.Run("diagonal", func(t *testing.T) {
		input := [][]int{
			{1, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 1, 0},
			{0, 0, 0, 1},
		}
		assert.Equal(t, 4, countIslands(input))
	})
	t.Run("ring", func(t *testing.T) {
		input := [][]int{
			{1, 1, 1, 1, 1},
			{1, 0, 0, 0, 1},
			{1, 0, 1, 0, 1},
			{1, 0, 0, 0, 1},
			{1, 1, 1, 1, 1},
		}
		assert.Equal(t, 2, countIslands(input))
	})
}

func countIslands(islandMap [][]int) int {
	var result int
	for i, line := range islandMap {
		for j, point := range line {
			if point == 1 {
				result++
				processIsland(islandMap, i, j)
			}
		}
	}

	return result
}

func processIsland(islandMap [][]int, i, j int) {
	if i < 0 || j < 0 || i >= len(islandMap) || j >= len(islandMap[0]) {
		return
	}
	if islandMap[i][j] == 0 {
		return
	}
	islandMap[i][j] = 0
	processIsland(islandMap, i, j+1)
	processIsland(islandMap, i+1, j)
	processIsland(islandMap, i, j-1)
	processIsland(islandMap, i-1, j)
}

///////////
