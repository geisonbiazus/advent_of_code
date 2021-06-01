package day2

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strconv"
)

type Policy struct {
	Min    int
	Max    int
	Letter string
}

type PasswordPolicy struct {
	Policy   Policy
	Password string
}

func ValidatePasswordMinMaxRule(policy Policy, password string) bool {
	count := 0

	for i := 0; i < len(password); i++ {
		if string(password[i]) == policy.Letter {
			count++
		}
	}

	return count >= policy.Min && count <= policy.Max
}

func ValidatePasswordPositionRule(policy Policy, password string) bool {
	validMin := string(password[policy.Min-1]) == policy.Letter
	validMax := string(password[policy.Max-1]) == policy.Letter

	return (validMin || validMax) && !(validMin && validMax)
}

type ValidatorFunc func(policy Policy, password string) bool

func CountValid(passwordPolicies []PasswordPolicy, validatorFunc ValidatorFunc) int {
	count := 0

	for _, p := range passwordPolicies {
		if validatorFunc(p.Policy, p.Password) {
			count++
		}
	}

	return count
}

func ParseInput(input io.Reader) []PasswordPolicy {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	passwordPolicies := []PasswordPolicy{}

	for scanner.Scan() {
		passwordPolicies = append(passwordPolicies, parsePasswordPolicy(scanner.Text()))
	}

	return passwordPolicies
}

var re = regexp.MustCompile(`^(\d+)-(\d+) (\S): (\S+)$`)

func parsePasswordPolicy(line string) PasswordPolicy {
	matches := re.FindStringSubmatch(line)

	min := toInt(matches[1])
	max := toInt(matches[2])
	letter := matches[3]
	password := matches[4]

	return PasswordPolicy{
		Policy:   Policy{Min: min, Max: max, Letter: letter},
		Password: password,
	}
}

func toInt(str string) int {
	n, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return n
}

func SolvePart1(inputPath string) int {
	input, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer input.Close()

	return CountValid(ParseInput(input), ValidatePasswordMinMaxRule)
}

func SolvePart2(inputPath string) int {
	input, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer input.Close()

	return CountValid(ParseInput(input), ValidatePasswordPositionRule)
}
