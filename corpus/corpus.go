package corpus

import (
	"bytes"
	"strings"
	"sync"
	"unicode"
)

var (
	once            sync.Once
	maleNames       []string
	femaleNames     []string
	surnames        []string
	fantasyNames    []string
	fantasySurnames []string
)

func Lines() (male, female, sur, fanFirst, fanSurn []string) {
	once.Do(func() {
		maleNames = parseBlobToLines(MaleNamesBlob, NameLength)
		femaleNames = parseBlobToLines(FemaleNamesBlob, NameLength)
		surnames = parseBlobToLines(SurnamesBlob, SurnameLength)
		fantasyNames = parseBlobToLines(FantasyNamesBlob, 0)
		fantasySurnames = parseBlobToLines(FantasySurnamesBlob, 0)
	})

	return maleNames, femaleNames, surnames, fantasyNames, fantasySurnames
}

func parseBlobToLines(blob []byte, fixedLen int) []string {
	if len(blob) == 0 {
		return nil
	}
	if bytes.ContainsRune(blob, '\n') || bytes.ContainsRune(blob, '\r') || fixedLen <= 0 {
		lines := bytes.Split(blob, []byte("\n"))

		return compactLines(lines)
	}

	out := make([]string, 0, len(blob)/fixedLen)
	for i := 0; i+fixedLen <= len(blob); i += fixedLen {
		chunk := blob[i : i+fixedLen]
		s := strings.TrimSpace(string(chunk))
		if s != "" {
			out = append(out, stripNonPrint(s))
		}
	}

	return out
}

func compactLines(lines [][]byte) []string {
	out := make([]string, 0, len(lines))
	for _, ln := range lines {
		s := strings.TrimSpace(string(ln))
		if s != "" {
			out = append(out, stripNonPrint(s))
		}
	}

	return out
}

func stripNonPrint(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for _, r := range s {
		if r == '\r' {
			continue
		}
		if unicode.IsPrint(r) {
			b.WriteRune(r)
		}
	}

	return b.String()
}
