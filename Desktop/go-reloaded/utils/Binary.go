package utils

import (
	"log"
	"strings"
	"strconv"
)

func Bin(b string) string {
	BinPos:= strings.Index(b, "(bin)")
	if BinPos == -1 {
		return b
	}

	beforeBin := strings.TrimSpace(b[:BinPos])
	if beforeBin == " " {
		return b
	}

	lastSpace:= strings.LastIndex(beforeBin, " ")
	var binNumber string
	if lastSpace == -1 {
		binNumber += beforeBin
	} else {
		binNumber += beforeBin[lastSpace+1:]
	}

	for _, c := range binNumber {
		if c != '1' && c != '0' {
			log.Fatalf("Error: Invalid binary digit '%c' in '%s'", c, binNumber)
		}
	}

	bin, err := strconv.ParseInt(binNumber, 2, 64)
	if err != nil {
		log.Fatal("Error to coverting to int (Binary)", err)
		return b
	}

	var result string
	if lastSpace == -1 {
		result += strconv.Itoa(int(bin)) + " "
	} else{
		result += beforeBin[:lastSpace + 1 ] + strconv.Itoa(int(bin))
	}

	if BinPos+5 < len(b) {
        result += b[BinPos+5:]
    }
	
	return result
}