package deagon

import (
	"testing"

	"github.com/oakroots/deagon/corpus"
)

// simple formatter for tests: "First Last"
type spaceFormatter struct{}

func (spaceFormatter) Format(first, last string) string { return first + " " + last }

// helper: build the index exactly like the implementation does
// genderBit: 0=male, 1=female (only matters for NameAuto)
// givenIx -> (index & maskGivenName) >> 1
// surIx   -> (index & maskSurname)   >> 9
func makeIndex(genderBit, givenIx, surIx int) int {
	return (surIx << 9) | (givenIx << 1) | (genderBit & 0x1)
}

func TestGetNameWithType_Male(t *testing.T) {
	male, _, surnames, _, _ := corpus.Lines()
	if len(male) == 0 || len(surnames) == 0 {
		t.Skip("empty corpus lists")
	}
	idx := makeIndex(0, 0, 0) // given=0, surname=0
	want := male[0] + " " + surnames[0]

	got := GetNameWithType(idx, spaceFormatter{}, NameMale)
	if got != want {
		t.Fatalf("want %q, got %q", want, got)
	}
}

func TestGetNameWithType_Female(t *testing.T) {
	_, female, surnames, _, _ := corpus.Lines()
	if len(female) == 0 || len(surnames) == 0 {
		t.Skip("empty corpus lists")
	}
	idx := makeIndex(0, 0, 0) // genderBit irrelevant for NameFemale
	want := female[0] + " " + surnames[0]

	got := GetNameWithType(idx, spaceFormatter{}, NameFemale)
	if got != want {
		t.Fatalf("want %q, got %q", want, got)
	}
}

func TestGetNameWithType_AutoMale(t *testing.T) {
	male, _, surnames, _, _ := corpus.Lines()
	if len(male) == 0 || len(surnames) == 0 {
		t.Skip("empty corpus lists")
	}
	idx := makeIndex(0, 0, 0) // genderBit=0 => male
	want := male[0] + " " + surnames[0]

	got := GetNameWithType(idx, spaceFormatter{}, NameAuto)
	if got != want {
		t.Fatalf("want %q, got %q", want, got)
	}
}

func TestGetNameWithType_AutoFemale(t *testing.T) {
	_, female, surnames, _, _ := corpus.Lines()
	if len(female) == 0 || len(surnames) == 0 {
		t.Skip("empty corpus lists")
	}
	idx := makeIndex(1, 0, 0) // genderBit=1 => female
	want := female[0] + " " + surnames[0]

	got := GetNameWithType(idx, spaceFormatter{}, NameAuto)
	if got != want {
		t.Fatalf("want %q, got %q", want, got)
	}
}

func TestGetNameWithType_Fantasy(t *testing.T) {
	_, _, _, fanFirst, fanSurn := corpus.Lines()
	if len(fanFirst) == 0 || len(fanSurn) == 0 {
		t.Skip("empty fantasy lists")
	}
	idx := makeIndex(0, 0, 0) // given=0, surname=0 in fantasy lists
	want := fanFirst[0] + " " + fanSurn[0]

	got := GetNameWithType(idx, spaceFormatter{}, NameFantasy)
	if got != want {
		t.Fatalf("want %q, got %q", want, got)
	}
}

func TestGetNameWithType_ModuloSafety(t *testing.T) {
	male, _, sur, _, _ := corpus.Lines()
	if len(male) == 0 || len(sur) == 0 {
		t.Skip("empty corpus lists")
	}
	hugeGiven := 1_000_000
	hugeSur := 2_000_000
	idx := makeIndex(0, hugeGiven, hugeSur)

	got := GetNameWithType(idx, spaceFormatter{}, NameMale)
	if got == "" {
		t.Fatal("expected non-empty name with large indices (modulo), got empty")
	}
}
