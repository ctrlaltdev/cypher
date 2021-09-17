package main

import (
	"bufio"
	"cypher/atbash"
	"cypher/polyalpha"
	"cypher/rot"
	"cypher/utils"
	"flag"
	"os"
	"strings"
)

var (
	cryptype = flag.String("t", "", "Encryption type")
	input    *string
	key      *string
	offset   = flag.Int("o", 0, "Rotation offset")
	polykey  = flag.String("k", "", "Polyalphabetic key")
	reverse  = flag.Bool("r", false, "Try to reverse encryption")
)

func main() {
	flag.Parse()

	utils.InitAlpha()

	args := strings.ToUpper(strings.Join(flag.Args(), " "))
	input = &args

	if *input == "" {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			stdin := strings.ToUpper(scanner.Text())
			input = &stdin
		}
	}

	if *polykey != "" {
		polykey := strings.ToUpper(*polykey)
		key = &polykey
	} else {
		polykey := "A"
		key = &polykey
	}

	if *reverse {

		switch *cryptype {

		case "rot":
			rot.Check(input)
			break

		case "atbash":
			atbash.Check(input)
			break

		case "polyalpha":
			polyalpha.Check(input, key)
			break

		}

	} else {

		switch *cryptype {

		case "rot":
			rot.Generate(input, offset)
			break

		case "atbash":
			atbash.Generate(input, offset)
			break

		case "polyalpha":
			polyalpha.Generate(input, key)
			break

		}

	}

}
