package day3_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/geisonbiazus/advent_of_code/2020/day3"
	"github.com/stretchr/testify/assert"
)

func TestMove(t *testing.T) {
	t.Run("It stays on place when right and down are zeroes", func(t *testing.T) {
		nextX, nextY, value := day3.Move([][]int{{1}}, 0, 0, 0, 0)

		assert.Equal(t, 0, nextX)
		assert.Equal(t, 0, nextY)
		assert.Equal(t, 1, value)
	})

	t.Run("It moves to the right", func(t *testing.T) {
		nextX, nextY, value := day3.Move([][]int{{1, 1, 0, 1}}, 0, 0, 2, 0)

		assert.Equal(t, 2, nextX)
		assert.Equal(t, 0, nextY)
		assert.Equal(t, 0, value)
	})

	t.Run("It moves down", func(t *testing.T) {
		nextX, nextY, value := day3.Move([][]int{
			{1},
			{1},
			{0},
			{1},
		}, 0, 0, 0, 2)

		assert.Equal(t, 0, nextX)
		assert.Equal(t, 2, nextY)
		assert.Equal(t, 0, value)
	})

	t.Run("It moves right and down", func(t *testing.T) {
		nextX, nextY, value := day3.Move([][]int{
			{1, 1, 1, 1},
			{1, 1, 1, 1},
			{1, 1, 0, 1},
			{1, 1, 1, 1},
		}, 0, 0, 2, 2)

		assert.Equal(t, 2, nextX)
		assert.Equal(t, 2, nextY)
		assert.Equal(t, 0, value)
	})

	t.Run("It starts from different positions", func(t *testing.T) {
		nextX, nextY, value := day3.Move([][]int{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 1},
		}, 1, 1, 2, 2)

		assert.Equal(t, nextX, 3)
		assert.Equal(t, nextY, 3)
		assert.Equal(t, value, 1)
	})

	t.Run("It loops horizontally", func(t *testing.T) {
		nextX, nextY, value := day3.Move([][]int{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 3, 0},
			{0, 0, 0, 0},
		}, 0, 0, 6, 2)

		assert.Equal(t, nextX, 2)
		assert.Equal(t, nextY, 2)
		assert.Equal(t, value, 3)
	})

	t.Run("It returns -1 when the passing trough the bottom", func(t *testing.T) {
		nextX, nextY, value := day3.Move([][]int{
			{0, 0, 0, 0},
		}, 0, 0, 0, 2)

		assert.Equal(t, 0, nextX)
		assert.Equal(t, 0, nextY)
		assert.Equal(t, -1, value)
	})
}

func TestCountTrees(t *testing.T) {
	assert.Equal(t, 0, day3.CountTrees([][]int{{0}}, 2, 2))

	assert.Equal(t, 1, day3.CountTrees([][]int{
		{0, 0},
		{0, 1},
	}, 1, 1))

	fmt.Println("start here")
	assert.Equal(t, 4, day3.CountTrees([][]int{
		{0, 1},
		{0, 1},
		{0, 1},
		{0, 1},
		{0, 1},
		{0, 1},
		{0, 1},
		{0, 1},
	}, 1, 1))
}

func TestParseInput(t *testing.T) {
	assert.Equal(t, [][]int{{0}}, day3.ParseInput(strings.NewReader(".\n")))
	assert.Equal(t, [][]int{{1}}, day3.ParseInput(strings.NewReader("#\n")))
	assert.Equal(t, [][]int{{1, 0, 0, 1}}, day3.ParseInput(strings.NewReader("#..#\n")))
	assert.Equal(t, [][]int{
		{1, 0, 0, 1},
		{0, 0, 0, 0},
		{0, 1, 0, 1},
	}, day3.ParseInput(strings.NewReader(""+
		"#..#\n"+
		"....\n"+
		".#.#\n",
	)))
}

func TestSolvePuzzle(t *testing.T) {
	assert.Equal(t, 294, day3.SolvePart1("./input.txt"))
	assert.Equal(t, 5774564250, day3.SolvePart2("./input.txt"))
}
