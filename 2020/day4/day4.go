package day4

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Measure struct {
	Value int
	Unit  string
}

type Passport struct {
	BirthYear      int
	IssueYear      int
	ExpirationYear int
	Height         Measure
	HairColor      string
	EyeColor       string
	PassportID     string
	CountryID      string
}

func AreRequiredFieldsPresent(passport Passport) bool {
	if passport.BirthYear == 0 {
		return false
	}

	if passport.IssueYear == 0 {
		return false
	}

	if passport.ExpirationYear == 0 {
		return false
	}

	if (passport.Height == Measure{}) {
		return false
	}

	if passport.HairColor == "" {
		return false
	}

	if passport.EyeColor == "" {
		return false
	}

	if passport.PassportID == "" {
		return false
	}

	return true
}

var hexColorRegex = regexp.MustCompile(`^#[0-9a-f]{6}$`)
var passportIDRegex = regexp.MustCompile(`^\d{9}$`)

func AreFieldsValid(passport Passport) bool {
	if !isBetween(passport.BirthYear, 1920, 2002) {
		return false
	}

	if !isBetween(passport.IssueYear, 2010, 2020) {
		return false
	}

	if !isBetween(passport.ExpirationYear, 2020, 2030) {
		return false
	}

	if !validateHeight(passport) {
		return false
	}

	if !hexColorRegex.MatchString(passport.HairColor) {
		return false
	}

	if !validateEyeColor(passport) {
		return false
	}

	if !passportIDRegex.MatchString(passport.PassportID) {
		return false
	}

	return true
}

func validateHeight(passport Passport) bool {
	switch passport.Height.Unit {
	case "cm":
		return isBetween(passport.Height.Value, 150, 193)
	case "in":
		return isBetween(passport.Height.Value, 59, 76)
	default:
		return false
	}
}

func isBetween(number, min, max int) bool {
	return number >= min && number <= max
}

var eyeColors = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

func validateEyeColor(passport Passport) bool {
	for _, color := range eyeColors {
		if passport.EyeColor == color {
			return true
		}
	}

	return false
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
		passport.BirthYear = parseInt(value)
	case "iyr":
		passport.IssueYear = parseInt(value)
	case "eyr":
		passport.ExpirationYear = parseInt(value)
	case "hgt":
		passport.Height = parseMeasure(value)
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

func parseInt(value string) int {
	intVal, err := strconv.Atoi(value)

	if err != nil {
		return 0
	}

	return intVal
}

var measureRegex = regexp.MustCompile(`^(\d+)?(\D+)?$`)

func parseMeasure(value string) Measure {
	matches := measureRegex.FindStringSubmatch(value)
	measure := Measure{}

	if matches != nil {
		measure.Value = parseInt(matches[1])
		measure.Unit = matches[2]
	}

	return measure
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
		if AreRequiredFieldsPresent(ParsePassport(data)) {
			countValid++
		}
	}

	return countValid
}

func SolvePart2(inputPath string) int {
	input, err := os.Open(inputPath)

	if err != nil {
		panic(err)
	}

	defer input.Close()

	countValid := 0

	for _, data := range EachPassportData(input) {
		if AreFieldsValid(ParsePassport(data)) {
			countValid++
		}
	}

	return countValid
}
