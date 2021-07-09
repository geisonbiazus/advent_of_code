package day4_test

import (
	"strings"
	"testing"

	"github.com/geisonbiazus/advent_of_code/2020/day4"
	"github.com/stretchr/testify/assert"
)

func AreRequiredFieldsPresent(t *testing.T) {
	t.Run("It returns true when all fields are set", func(t *testing.T) {
		passport := buildPassport()
		assert.True(t, day4.AreRequiredFieldsPresent(passport))
	})

	t.Run("It returns true when CountryId is not present", func(t *testing.T) {
		passport := buildPassport()
		passport.CountryID = ""
		assert.True(t, day4.AreRequiredFieldsPresent(passport))
	})

	t.Run("It validates presence of BirthYear", func(t *testing.T) {
		passport := buildPassport()
		passport.BirthYear = 0
		assert.False(t, day4.AreRequiredFieldsPresent(passport))
	})

	t.Run("It validates presence of IssueYear", func(t *testing.T) {
		passport := buildPassport()
		passport.IssueYear = 0
		assert.False(t, day4.AreRequiredFieldsPresent(passport))
	})

	t.Run("It validates presence of ExpirationYear", func(t *testing.T) {
		passport := buildPassport()
		passport.ExpirationYear = 0
		assert.False(t, day4.AreRequiredFieldsPresent(passport))
	})

	t.Run("It validates presence of Height", func(t *testing.T) {
		passport := buildPassport()
		passport.Height = day4.Measure{}
		assert.False(t, day4.AreRequiredFieldsPresent(passport))
	})

	t.Run("It validates presence of HairColor", func(t *testing.T) {
		passport := buildPassport()
		passport.HairColor = ""
		assert.False(t, day4.AreRequiredFieldsPresent(passport))
	})

	t.Run("It validates presence of EyeColor", func(t *testing.T) {
		passport := buildPassport()
		passport.EyeColor = ""
		assert.False(t, day4.AreRequiredFieldsPresent(passport))
	})

	t.Run("It validates presence of PassportID", func(t *testing.T) {
		passport := buildPassport()
		passport.PassportID = ""
		assert.False(t, day4.AreRequiredFieldsPresent(passport))
	})
}

func TestAreFieldsValid(t *testing.T) {
	t.Run("It returns true when all fields are valid", func(t *testing.T) {
		passport := buildPassport()
		assert.True(t, day4.AreFieldsValid(passport))
	})

	t.Run("It returns true when CountryId is not present", func(t *testing.T) {
		passport := buildPassport()
		passport.CountryID = ""
		assert.True(t, day4.AreFieldsValid(passport))
	})

	t.Run("It validates BirthYear between 1920 and 2002", func(t *testing.T) {
		passport := buildPassport()
		passport.BirthYear = 1919
		assert.False(t, day4.AreFieldsValid(passport))

		passport.BirthYear = 2003
		assert.False(t, day4.AreFieldsValid(passport))

		passport.BirthYear = 1920
		assert.True(t, day4.AreFieldsValid(passport))

		passport.BirthYear = 2002
		assert.True(t, day4.AreFieldsValid(passport))
	})

	t.Run("It validates IssueYear between 2010 and 2020", func(t *testing.T) {
		passport := buildPassport()
		passport.IssueYear = 2009
		assert.False(t, day4.AreFieldsValid(passport))

		passport.IssueYear = 2021
		assert.False(t, day4.AreFieldsValid(passport))

		passport.IssueYear = 2010
		assert.True(t, day4.AreFieldsValid(passport))

		passport.IssueYear = 2020
		assert.True(t, day4.AreFieldsValid(passport))
	})

	t.Run("It validates ExpirationYear between 2020 and 2030", func(t *testing.T) {
		passport := buildPassport()
		passport.ExpirationYear = 2019
		assert.False(t, day4.AreFieldsValid(passport))

		passport.ExpirationYear = 2031
		assert.False(t, day4.AreFieldsValid(passport))

		passport.ExpirationYear = 2020
		assert.True(t, day4.AreFieldsValid(passport))

		passport.ExpirationYear = 2030
		assert.True(t, day4.AreFieldsValid(passport))
	})

	t.Run("It validates Height", func(t *testing.T) {
		t.Run("In cm between 150 adn 193", func(t *testing.T) {
			passport := buildPassport()
			passport.Height = day4.Measure{149, "cm"}
			assert.False(t, day4.AreFieldsValid(passport))

			passport.Height = day4.Measure{194, "cm"}
			assert.False(t, day4.AreFieldsValid(passport))

			passport.Height = day4.Measure{150, "cm"}
			assert.True(t, day4.AreFieldsValid(passport))

			passport.Height = day4.Measure{193, "cm"}
			assert.True(t, day4.AreFieldsValid(passport))
		})

		t.Run("In in between 59 adn 76", func(t *testing.T) {
			passport := buildPassport()
			passport.Height = day4.Measure{58, "in"}
			assert.False(t, day4.AreFieldsValid(passport))

			passport.Height = day4.Measure{77, "in"}
			assert.False(t, day4.AreFieldsValid(passport))

			passport.Height = day4.Measure{59, "in"}
			assert.True(t, day4.AreFieldsValid(passport))

			passport.Height = day4.Measure{76, "in"}
			assert.True(t, day4.AreFieldsValid(passport))
		})

		t.Run("It fails if unit is not cm or in", func(t *testing.T) {
			passport := buildPassport()
			passport.Height = day4.Measure{59, "aa"}
			assert.False(t, day4.AreFieldsValid(passport))

			passport.Height = day4.Measure{150, "aa"}
			assert.False(t, day4.AreFieldsValid(passport))
		})
	})

	t.Run("It validates HairColor", func(t *testing.T) {
		passport := buildPassport()
		passport.HairColor = ""
		assert.False(t, day4.AreFieldsValid(passport))

		passport.HairColor = "#ffffff"
		assert.True(t, day4.AreFieldsValid(passport))

		passport.HairColor = "#f9f9f9"
		assert.True(t, day4.AreFieldsValid(passport))

		passport.HairColor = "#fffff"
		assert.False(t, day4.AreFieldsValid(passport))

		passport.HairColor = "#fffffff"
		assert.False(t, day4.AreFieldsValid(passport))

		passport.HairColor = "ffffff"
		assert.False(t, day4.AreFieldsValid(passport))
	})

	t.Run("It validates EyeColor", func(t *testing.T) {
		passport := buildPassport()
		passport.EyeColor = "amb"
		assert.True(t, day4.AreFieldsValid(passport))

		passport.EyeColor = "blu"
		assert.True(t, day4.AreFieldsValid(passport))

		passport.EyeColor = "gry"
		assert.True(t, day4.AreFieldsValid(passport))

		passport.EyeColor = "grn"
		assert.True(t, day4.AreFieldsValid(passport))

		passport.EyeColor = "hzl"
		assert.True(t, day4.AreFieldsValid(passport))

		passport.EyeColor = "oth"
		assert.True(t, day4.AreFieldsValid(passport))

		passport.EyeColor = "aaa"
		assert.False(t, day4.AreFieldsValid(passport))

		passport.EyeColor = ""
		assert.False(t, day4.AreFieldsValid(passport))
	})

	t.Run("It validates PassportID", func(t *testing.T) {
		passport := buildPassport()
		passport.PassportID = ""
		assert.False(t, day4.AreFieldsValid(passport))

		passport.PassportID = "000000000"
		assert.True(t, day4.AreFieldsValid(passport))

		passport.PassportID = "123456789"
		assert.True(t, day4.AreFieldsValid(passport))

		passport.PassportID = "1234567899"
		assert.False(t, day4.AreFieldsValid(passport))

		passport.PassportID = "12345678"
		assert.False(t, day4.AreFieldsValid(passport))
	})
}

func TestParsePassport(t *testing.T) {
	t.Run("It returns empty passport with empty string", func(t *testing.T) {
		input := ""
		assert.Equal(t, day4.Passport{}, day4.ParsePassport(input))
	})

	t.Run("It parses passport fields", func(t *testing.T) {
		assert.Equal(t, day4.Passport{BirthYear: 1937}, day4.ParsePassport("byr:1937"))
		assert.Equal(t, day4.Passport{IssueYear: 2017}, day4.ParsePassport("iyr:2017"))
		assert.Equal(t, day4.Passport{ExpirationYear: 2020}, day4.ParsePassport("eyr:2020"))
		assert.Equal(t, day4.Passport{Height: day4.Measure{183, "cm"}}, day4.ParsePassport("hgt:183cm"))
		assert.Equal(t, day4.Passport{HairColor: "#fffffd"}, day4.ParsePassport("hcl:#fffffd"))
		assert.Equal(t, day4.Passport{EyeColor: "gry"}, day4.ParsePassport("ecl:gry"))
		assert.Equal(t, day4.Passport{PassportID: "860033327"}, day4.ParsePassport("pid:860033327"))
		assert.Equal(t, day4.Passport{CountryID: "147"}, day4.ParsePassport("cid:147"))
	})

	t.Run("It parses Height field correctly", func(t *testing.T) {
		assert.Equal(t, day4.Passport{Height: day4.Measure{183, ""}}, day4.ParsePassport("hgt:183"))
		assert.Equal(t, day4.Passport{Height: day4.Measure{0, "cm"}}, day4.ParsePassport("hgt:cm"))
		assert.Equal(t, day4.Passport{Height: day4.Measure{90, "in"}}, day4.ParsePassport("hgt:90in"))
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
	assert.Equal(t, 158, day4.SolvePart2("input.txt"))
}

func buildPassport() day4.Passport {
	return day4.Passport{
		BirthYear:      1937,
		IssueYear:      2017,
		ExpirationYear: 2020,
		Height:         day4.Measure{183, "cm"},
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
