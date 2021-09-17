package polyalpha

import (
	"cypher/rot"
	"cypher/utils"
	"fmt"
	"strings"
)

func polyalpha(cipher *string, key *string, reverse bool) *string {
	var substitution string

	keylist := strings.Split(*key, "")

	for i, l := range strings.Split(*cipher, "") {

		var offset int
		if reverse {
			offset = 0 - (utils.AlphaMap[keylist[i%len(keylist)]] - 1)
		} else {
			offset = utils.AlphaMap[keylist[i%len(keylist)]] - 1
		}

		res := rot.Sub(&l, offset, false)

		substitution += res
	}

	return &substitution
}

func Generate(plaintext *string, key *string) {
	sub := polyalpha(plaintext, key, false)
	fmt.Printf("%s\n", *sub)
}

func Check(cipher *string, key *string) {
	sub := polyalpha(cipher, key, true)
	fmt.Printf("%s\n", *sub)
}
