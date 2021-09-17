package atbash

import (
	"cypher/rot"
	"cypher/utils"
	"fmt"
)

func Generate(plaintext *string, offset *int) {
	sub := rot.Rot(plaintext, *offset, true)
	fmt.Printf("%s\n", *sub)
}

func Check(cipher *string) {
	for i := 0; i < len(utils.AlphaList); i++ {
		sub := rot.Rot(cipher, i, true)
		fmt.Printf("%d\t%s\n", i, *sub)
	}
}
