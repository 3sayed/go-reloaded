package utils

import(
	"strconv"
	"strings"
	"log"
)

func Hex(h string) string {

	//Find The Position Of (Hex)
	hexPos := strings.Index(h, "(hex)")
	if hexPos == -1 {
		return h
	}

	// Retreive strings before (Hex)
	beforeHex := strings.TrimSpace(h[:hexPos])
	if beforeHex == " " {
		return h
	}

	lastSapce := strings.LastIndex(beforeHex, " ")
	var hexNumber string
	if lastSapce == -1 {
		hexNumber += beforeHex
	} else{
		hexNumber += beforeHex[lastSapce+1:] // only the word before hex
	}

	//Convert to from string to int and from int to decimal 
	dec, err := strconv.ParseInt(hexNumber, 16, 64)
	if err!= nil {
		log.Printf("Error converting hex '%s' to decimal: %v", hexNumber, err)
		return h
	}

	var result string 
	if lastSapce == -1 {
		result += strconv.Itoa(int(dec))
	} else {
		result += beforeHex[:lastSapce + 1] + strconv.Itoa(int(dec))
	}

	if hexPos+5 < len(h) {
		result += h[hexPos+5:]
	}
	return result
}