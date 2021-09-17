package rot

import (
	"cypher/utils"
	"fmt"
	"strings"
)

func Sub(letter *string, offset int, atbash bool) string {

	if utils.AlphaMap[*letter] == 0 {
		return *letter
	}

	pos := utils.AlphaMap[*letter] - 1

	if atbash {
		offset = len(utils.AlphaList) - pos + offset
	} else {
		offset = pos + offset
	}

	if offset < 0 {
		offset = len(utils.AlphaList) + offset
	}

	return utils.AlphaList[offset%len(utils.AlphaList)]
}

func Rot(cipher *string, offset int, atbash bool) *string {

	var substitution string

	for _, l := range strings.Split(*cipher, "") {

		res := Sub(&l, offset, atbash)

		substitution += res
	}

	return &substitution
}

func Generate(plaintext *string, offset *int) {
	sub := Rot(plaintext, *offset, false)
	fmt.Printf("%s\n", *sub)
}

func Check(cipher *string) {
	for i := 0; i < len(utils.AlphaList); i++ {
		sub := Rot(cipher, i, false)
		fmt.Printf("%d\t%s\n", i, *sub)
	}
}
