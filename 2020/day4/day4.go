package day4

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strings"
)

type Passport struct {
	BirthYear      string
	IssueYear      string
	ExpirationYear string
	Height         string
	HairColor      string
	EyeColor       string
	PassportID     string
	CountryID      string
}

func (p Passport) IsValid() bool {
	requiredFields := []string{
		p.BirthYear,
		p.IssueYear,
		p.ExpirationYear,
		p.Height,
		p.HairColor,
		p.EyeColor,
		p.PassportID,
	}

	for _, v := range requiredFields {
		if v == "" {
			return false
		}
	}

	return true
}

var whiteSpaceRe = regexp.MustCompile(`\s`)

func ParsePassport(input string) Passport {
	fields := whiteSpaceRe.Split(input, -1)
	passport := Passport{}

	for _, field := range fields {
		passport = parseField(passport, field)
	}

	return passport
}

func parseField(passport Passport, field string) Passport {
	fieldValue := strings.Split(field, ":")

	if len(fieldValue) == 2 {
		passport = setField(passport, fieldValue[0], fieldValue[1])
	}

	return passport
}

func setField(passport Passport, field, value string) Passport {
	switch field {
	case "byr":
		passport.BirthYear = value
	case "iyr":
		passport.IssueYear = value
	case "eyr":
		passport.ExpirationYear = value
	case "hgt":
		passport.Height = value
	case "hcl":
		passport.HairColor = value
	case "ecl":
		passport.EyeColor = value
	case "pid":
		passport.PassportID = value
	case "cid":
		passport.CountryID = value
	}
	return passport
}

func EachPassportData(input io.Reader) []string {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	passports := []string{}
	passport := ""

	for scanner.Scan() {
		text := strings.Trim(scanner.Text(), " ")

		if text == "" && passport != "" {
			passports = append(passports, strings.Trim(passport, " \n"))
			passport = ""
		} else {
			passport += "\n" + text
		}
	}

	if passport != "" {
		passports = append(passports, strings.Trim(passport, " \n"))
	}

	return passports
}

func SolvePart1(inputPath string) int {
	input, err := os.Open(inputPath)

	if err != nil {
		panic(err)
	}

	defer input.Close()

	countValid := 0

	for _, data := range EachPassportData(input) {
		if ParsePassport(data).IsValid() {
			countValid++
		}
	}

	return countValid
}
