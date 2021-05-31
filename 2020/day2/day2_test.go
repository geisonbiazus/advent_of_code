package day2_test

import (
	"strings"
	"testing"

	"github.com/geisonbiazus/advent_of_code/2020/day2"
	"github.com/stretchr/testify/assert"
)

func TestValidatePassword(t *testing.T) {
	assert.True(t, day2.ValidatePassword(policy(1, 3, "a"), "abcde"))
	assert.True(t, day2.ValidatePassword(policy(1, 3, "a"), "aaabcde"))
	assert.False(t, day2.ValidatePassword(policy(1, 3, "a"), "bcde"))
	assert.False(t, day2.ValidatePassword(policy(1, 3, "a"), "aaaabcde"))
	assert.True(t, day2.ValidatePassword(policy(3, 6, "b"), "abbbcde"))
	assert.True(t, day2.ValidatePassword(policy(3, 6, "b"), "aaabbbbbbcde"))
	assert.False(t, day2.ValidatePassword(policy(3, 6, "b"), "abbcde"))
	assert.False(t, day2.ValidatePassword(policy(3, 6, "b"), "abbbbbbbbcde"))
}

func TestCountValidPasswords(t *testing.T) {
	assert.Equal(t, day2.CountValid([]day2.PasswordPolicy{}), 0)
	assert.Equal(t, day2.CountValid([]day2.PasswordPolicy{
		pwPolicy(1, 3, "a", "asd"),
	}), 1)
	assert.Equal(t, day2.CountValid([]day2.PasswordPolicy{
		pwPolicy(2, 3, "a", "asd"),
	}), 0)

	assert.Equal(t, day2.CountValid([]day2.PasswordPolicy{
		pwPolicy(1, 3, "a", "asd"),
		pwPolicy(1, 3, "a", "aaaasd"),
		pwPolicy(1, 3, "a", "aasd"),
	}), 2)
}

func TestParseInput(t *testing.T) {
	assert.Equal(t, []day2.PasswordPolicy{
		pwPolicy(1, 3, "a", "abcde"),
		pwPolicy(1, 3, "b", "cdefg"),
		pwPolicy(2, 9, "c", "ccccccccc"),
	}, day2.ParseInput(strings.NewReader(sampleInput)))
}

func TestSolvePuzzle(t *testing.T) {
	assert.Equal(t, 396, day2.SolvePart1("./input.txt"))
}

func policy(min, max int, letter string) day2.Policy {
	return day2.Policy{Min: min, Max: max, Letter: letter}
}

func pwPolicy(min, max int, letter, password string) day2.PasswordPolicy {
	return day2.PasswordPolicy{
		Policy:   policy(min, max, letter),
		Password: password,
	}
}

const sampleInput = `1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc
`
