package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	flag "github.com/spf13/pflag"
)

func parse(str string) ([]int, int) {
	a := strings.Split(str, "/")
	mask, _ := strconv.Atoi(a[1])
	stra := strings.Split(a[0], ".")

	ipa := make([]int, 4)

	for p := range stra {
		ipa[p], _ = strconv.Atoi(stra[p])
	}
	return ipa, mask
}

// IsOverlap returns whether two IPv4 CIDR blocks overlap or not.
// It assumes that the CIDR blocks are valid and doesn't perform any validation
// of its own. The un-use of go's "net" http libary is intentional :).
func IsOverlap(cidra, cidrb string) bool {
	a, maska := parse(cidra)
	b, maskb := parse(cidrb)
	mask := min(maska, maskb)

	bitseta, bitsetb, c := 0, 0, 0

	// run for each block in IP block
	for nn := 0; nn < 4; nn++ {
		n := 7
		for c < mask && n >= 0 {
			bitseta |= a[nn] & (1 << n)
			bitsetb |= b[nn] & (1 << n)
			c++
			n--
		}
		if c >= mask {
			break
		}
	}

	return bitseta == bitsetb
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var cidra, cidrb string
	flag.StringVar(&cidra, "a", "", "CIDR block a")
	flag.StringVar(&cidrb, "b", "", "CIDR block b")
	flag.Parse()

	if cidra == "" || cidrb == "" {
		fmt.Println("cannot pass empty cidr")
		os.Exit(1)
	}

	ans := IsOverlap(cidra, cidrb)

	fmt.Println(ans)
}
