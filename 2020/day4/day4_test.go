package day4_test

import (
	"strings"
	"testing"

	"github.com/geisonbiazus/advent_of_code/2020/day4"
	"github.com/stretchr/testify/assert"
)

func TestPassport(t *testing.T) {
	t.Run("IsValid", func(t *testing.T) {
		t.Run("It returns true when all fields are set", func(t *testing.T) {
			passport := buildPassport()
			assert.True(t, passport.IsValid())
		})

		t.Run("It returns true when CountryId is not present", func(t *testing.T) {
			passport := buildPassport()
			passport.CountryID = ""
			assert.True(t, passport.IsValid())
		})

		t.Run("It validates presence of BirthYear", func(t *testing.T) {
			passport := buildPassport()
			passport.BirthYear = ""
			assert.False(t, passport.IsValid())
		})

		t.Run("It validates presence of IssueYear", func(t *testing.T) {
			passport := buildPassport()
			passport.IssueYear = ""
			assert.False(t, passport.IsValid())
		})

		t.Run("It validates presence of ExpirationYear", func(t *testing.T) {
			passport := buildPassport()
			passport.ExpirationYear = ""
			assert.False(t, passport.IsValid())
		})

		t.Run("It validates presence of Height", func(t *testing.T) {
			passport := buildPassport()
			passport.Height = ""
			assert.False(t, passport.IsValid())
		})

		t.Run("It validates presence of HairColor", func(t *testing.T) {
			passport := buildPassport()
			passport.HairColor = ""
			assert.False(t, passport.IsValid())
		})

		t.Run("It validates presence of EyeColor", func(t *testing.T) {
			passport := buildPassport()
			passport.EyeColor = ""
			assert.False(t, passport.IsValid())
		})

		t.Run("It validates presence of PassportID", func(t *testing.T) {
			passport := buildPassport()
			passport.PassportID = ""
			assert.False(t, passport.IsValid())
		})
	})
}

func TestParsePassport(t *testing.T) {
	t.Run("It returns empty passport with empty string", func(t *testing.T) {
		input := ""
		assert.Equal(t, day4.Passport{}, day4.ParsePassport(input))
	})

	t.Run("It parses passport fields", func(t *testing.T) {
		assert.Equal(t, day4.Passport{BirthYear: "1937"}, day4.ParsePassport("byr:1937"))
		assert.Equal(t, day4.Passport{IssueYear: "2017"}, day4.ParsePassport("iyr:2017"))
		assert.Equal(t, day4.Passport{ExpirationYear: "2020"}, day4.ParsePassport("eyr:2020"))
		assert.Equal(t, day4.Passport{Height: "183cm"}, day4.ParsePassport("hgt:183cm"))
		assert.Equal(t, day4.Passport{HairColor: "#fffffd"}, day4.ParsePassport("hcl:#fffffd"))
		assert.Equal(t, day4.Passport{EyeColor: "gry"}, day4.ParsePassport("ecl:gry"))
		assert.Equal(t, day4.Passport{PassportID: "860033327"}, day4.ParsePassport("pid:860033327"))
		assert.Equal(t, day4.Passport{CountryID: "147"}, day4.ParsePassport("cid:147"))
	})

	t.Run("It parses multiple fields at once", func(t *testing.T) {
		assert.Equal(t, buildPassport(), day4.ParsePassport("ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm"))
	})

	t.Run("It parses multiline passports", func(t *testing.T) {
		assert.Equal(t, buildPassport(), day4.ParsePassport("ecl:gry pid:860033327\neyr:2020 hcl:#fffffd byr:1937\niyr:2017 cid:147 hgt:183cm"))
		assert.Equal(t, buildPassport(), day4.ParsePassport("ecl:gry pid:860033327\neyr:2020 hcl:#fffffd byr:1937\niyr:2017 cid:147 hgt:183cm"))
	})
}

func TestEachPassportData(t *testing.T) {
	t.Run("It returns a slice of each passport data split", func(t *testing.T) {
		assert.Equal(t, splitPassports, day4.EachPassportData(strings.NewReader(passports)))
	})
}

func TestSolvePuzzle(t *testing.T) {
	assert.Equal(t, 250, day4.SolvePart1("input.txt"))
}

func buildPassport() day4.Passport {
	return day4.Passport{
		BirthYear:      "1937",
		IssueYear:      "2017",
		ExpirationYear: "2020",
		Height:         "183cm",
		HairColor:      "#fffffd",
		EyeColor:       "gry",
		PassportID:     "860033327",
		CountryID:      "147",
	}
}

var passports = `
ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in
`

var splitPassports = []string{
	`ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm`,

	`iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929`,

	`hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm`,

	`hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`,
}
