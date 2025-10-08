package xpass

import (
	"os"
)

type SymbolType string

var (
	defaultPrefixBrand = "YOUR-APP"

	SymbolAT SymbolType = "AT"
	SymbolRT SymbolType = "RT"
	SymbolAC SymbolType = "AC"
	SymbolDC SymbolType = "DC"
)

func GetSymbol(t SymbolType) string {
	prefixSign := os.Getenv("XPASS_PREFIX_SIGN")
	if prefixSign == "" {
		prefixSign = defaultPrefixBrand
	}

	return prefixSign + ":PrrXpass:" + string(t) + "."
}