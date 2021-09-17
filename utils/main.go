package utils

import (
	"log"
)

var (
	AlphaMap  = make(map[string]int)
	AlphaList = [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
)

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func InitAlpha() {
	for i, l := range AlphaList {
		AlphaMap[l] = i + 1
	}
}
