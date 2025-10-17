package corpus

import _ "embed"

const (
	// NameLength is width of the record (5 + 1 for newline)
	NameLength = 6
	// SurnameLength is width of the record (8 + 1 for newline)
	SurnameLength = 9
)

//go:embed gfnames.txt
var FemaleNamesBlob []byte

//go:embed gmnames.txt
var MaleNamesBlob []byte

//go:embed srnames.txt
var SurnamesBlob []byte

//go:embed fantasy_unisex_first_names.txt
var FantasyNamesBlob []byte

//go:embed fantasy_unisex_surnames.txt
var FantasySurnamesBlob []byte
