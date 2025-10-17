package deagon

import (
	"github.com/oakroots/deagon/corpus"
)

type NameType int

const (
	NameAuto NameType = iota
	NameMale
	NameFemale
	NameFantasy

	// Bitmask for gender (0 = male, 1 = female)
	maskGender int = 0x0000001

	// Bitmask for given name index
	maskGivenName int = 0x00001FE

	// Bitmask for surname index
	maskSurname int = 0x1FFFE00

	// Total number of entries (2^25)
	totalEntriesFull int = 33554432
)

func getNamesTyped(index int, t NameType) (string, string) {
	givenIx := (index & maskGivenName) >> 1
	surIx := (index & maskSurname) >> 9

	male, female, sur, fanFirst, fanSurn := corpus.Lines()

	switch t {
	case NameFantasy:
		first := pick(fanFirst, givenIx)
		last := pick(fanSurn, surIx)

		return first, last
	case NameMale:
		first := pick(male, givenIx)
		last := pick(sur, surIx)

		return first, last
	case NameFemale:
		first := pick(female, givenIx)
		last := pick(sur, surIx)

		return first, last
	case NameAuto:
		fallthrough
	default:
		var first string
		if (index & maskGender) == 0 {
			first = pick(male, givenIx)
		} else {
			first = pick(female, givenIx)
		}

		last := pick(sur, surIx)

		return first, last
	}
}

func pick(ss []string, ix int) string {
	if len(ss) == 0 {
		return ""
	}

	return ss[ix%len(ss)]
}

// getNames returns the first name and surname for a given index.
func getNames(index int) (string, string) {
	return getNamesTyped(index, NameAuto)
}

// getName returns the formatted full name (first name + surname)
// using the provided Formatter implementation.
func getName(index int, formatter Formatter) string {
	firstname, surname := getNames(index)
	return formatter.Format(firstname, surname)
}

func GetNameWithType(index int, formatter Formatter, t NameType) string {
	first, last := getNamesTyped(index, t)

	return formatter.Format(first, last)
}
