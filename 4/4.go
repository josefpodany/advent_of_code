package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Passport is a type alias to map[string]string
type Passport map[string]string

// Passports is a type alias to a Passport slice
type Passports []Passport

var (
	eyeColorCodes  = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	hairColorRegex = regexp.MustCompile(`^#[\w]{6}$`)
	pidRegex       = regexp.MustCompile(`^[0-9]{9}$`)
	validators     = map[string]func(string) bool{
		// validateYear is a curried function. This way, we have one generic
		// factory that builds different specializations of the validator.
		"byr": validateYear(1920, 2002),
		"iyr": validateYear(2010, 2020),
		"eyr": validateYear(2020, 2030),
		"hgt": isHeightValid,
		"ecl": isEyeColorValid,
		"hcl": func(hc string) bool { return hairColorRegex.MatchString(hc) },
		"pid": func(pid string) bool { return pidRegex.MatchString(pid) },
	}
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Printf("couldn't read file: %s\n", err.Error())
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ps := Passports{}
	p := Passport{}

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			// This is the newline that separates the passports.
			// Add the passport to the rest and create a new one.
			ps = append(ps, p)
			p = Passport{}
			continue
		}
		pairs := strings.Split(line, " ")
		for _, pair := range pairs {
			items := strings.Split(pair, ":")
			if len(items) != 2 {
				continue
			}
			p[items[0]] = items[1]
		}
	}

	// Add the last passport.
	ps = append(ps, p)

	mandatoryFields := []string{}
	// Get all mandatory fields from the validations.
	for field := range validators {
		mandatoryFields = append(mandatoryFields, field)
	}
	withMandatory := ps.FilterMandatory(mandatoryFields)
	withMandatoryValid := ps.FilterMandatory(mandatoryFields).FilterValid()
	fmt.Printf("passports w/ mandatory fields: %d\n", len(withMandatory))
	fmt.Printf("valid passports: %d\n", len(withMandatoryValid))
}

// HasMandatoryFields described by the mandatory slice passed in
func (P Passport) HasMandatoryFields(mandatory []string) bool {
	for _, field := range mandatory {
		if _, ok := P[field]; !ok {
			return false
		}
	}
	return true
}

// Valid returns if the passport is valid according to the validators declared above.
// Function relies on the passport having only the mandatory and optional keys.
func (P Passport) Valid() bool {
	for k, v := range P {
		if k != "cid" && !validators[k](v) {
			return false
		}
	}
	return true
}

// FilterMandatory keys in the passports
func (Ps Passports) FilterMandatory(mandatory []string) Passports {
	valid := Passports{}
	for _, p := range Ps {
		if p.HasMandatoryFields(mandatory) {
			valid = append(valid, p)
		}
	}
	return valid
}

// FilterValid passports with the validators declared above
func (Ps Passports) FilterValid() Passports {
	valid := Passports{}
	for _, p := range Ps {
		if p.Valid() {
			valid = append(valid, p)
		}
	}
	return valid
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

func isEyeColorValid(ecl string) bool {
	for _, validEcl := range eyeColorCodes {
		if ecl == validEcl {
			return true
		}
	}
	return false
}
