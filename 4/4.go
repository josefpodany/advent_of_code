package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var validEyeColors = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
var validations = map[string]func(string) bool{
	"byr": validateYear(1920, 2002),
	"iyr": validateYear(2010, 2020),
	"eyr": validateYear(2020, 2030),
	"hgt": isHeightValid,
	"hcl": isHairColorValid,
	"ecl": isEyeColorValid,
	"pid": isPIDValid,
}

func hasAllMandatory(pass map[string]string, mandatory []string) bool {
	for _, field := range mandatory {
		if _, ok := pass[field]; !ok {
			return false
		}
	}
	return true
}
func filterMandatory(passports []map[string]string, mandatory []string) []map[string]string {
	withMandatory := []map[string]string{}
	for _, pass := range passports {
		if hasAllMandatory(pass, mandatory) {
			withMandatory = append(withMandatory, pass)
		}
	}
	return withMandatory
}

func validateYear(atLeast, atMost int) func(string) bool {
	return func(year string) bool {
		if len(year) != 4 {
			return false
		}

		birth, err := strconv.Atoi(year)
		if err != nil {
			return false
		}
		return birth >= atLeast && birth <= atMost
	}
}

func isHeightValid(height string) bool {
	if len(height) < 4 {
		return false
	}
	units := height[len(height)-2:]
	h, err := strconv.Atoi(height[:len(height)-2])
	if err != nil {
		fmt.Printf("malformed height '%s': invalid number\n", height)
		return false
	}

	switch units {
	case "in":
		return h >= 59 && h <= 76
	case "cm":
		return h >= 150 && h <= 193
	default:
		fmt.Printf("malformed height '%s': invalid units\n", height)
	}
	return false
}

func isHairColorValid(hc string) bool {
	ok, _ := regexp.MatchString(`^#[\w]{6}$`, hc)
	return ok
}

func isEyeColorValid(ecl string) bool {
	for _, validEcl := range validEyeColors {
		if ecl == validEcl {
			return true
		}
	}
	return false
}

func isPIDValid(pid string) bool {
	ok, _ := regexp.MatchString(`^[0-9]{9}$`, pid)
	return ok
}

func isPassportValid(passport map[string]string) bool {
	for k, v := range passport {
		if k == "cid" {
			continue
		}
		if !validations[k](v) {
			return false
		}
	}
	return true
}

func filterValid(passports []map[string]string) []map[string]string {
	valid := []map[string]string{}
	for _, pass := range passports {
		if isPassportValid(pass) {
			valid = append(valid, pass)
		}
	}
	return valid
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Printf("couldn't read file: %s\n", err.Error())
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	passports := []map[string]string{}
	passport := map[string]string{}

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			// This is the newline that separates the passports.
			// Add the passport to the rest and create a new one.
			passports = append(passports, passport)
			passport = map[string]string{}
			continue
		}
		pairs := strings.Split(line, " ")
		for _, pair := range pairs {
			items := strings.Split(pair, ":")
			if len(items) != 2 {
				continue
			}
			passport[items[0]] = items[1]
		}
	}

	// Add the last passport.
	passports = append(passports, passport)

	mandatoryFields := []string{}
	// Get all mandatory fields from the validations.
	for field := range validations {
		mandatoryFields = append(mandatoryFields, field)
	}

	passports = filterMandatory(passports, mandatoryFields)
	fmt.Printf("passports w/ mandatory fields: %d\n", len(passports))
	passports = filterValid(passports)
	fmt.Printf("valid passports: %d\n", len(passports))
}
