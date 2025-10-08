package xpass

import (
	"os"
)

type SymbolType string

var (
	brandPrefix = "YOUR_APP"
	signature = "prrXpass"

	SymbolAT SymbolType = "AT"
	SymbolRT SymbolType = "RT"
	SymbolAC SymbolType = "AC"
	SymbolDC SymbolType = "DC"
)

func GetSymbol(t SymbolType) string {
	prefixSign := os.Getenv("XPASS_PREFIX_SIGN")
	if prefixSign == "" {
		prefixSign = brandPrefix
	}

	return prefixSign + ":" + signature + ":" + string(t) + "_"
}