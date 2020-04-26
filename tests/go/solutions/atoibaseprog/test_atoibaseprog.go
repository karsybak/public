package main

import (
	"./base"

	"github.com/01-edu/z01"
)

// this is the function that creates the TESTS
func main() {
	type node struct {
		s    string
		base string
	}

	table := []node{}

	// 5 random pairs of string numbers with valid bases
	for i := 0; i < 5; i++ {
		validBaseToInput := base.Valid()
		val := node{
			s:    base.StringFrom(validBaseToInput),
			base: validBaseToInput,
		}
		table = append(table, val)
	}
	// 5 random pairs of string numbers with invalid bases
	for i := 0; i < 5; i++ {
		invalidBaseToInput := base.Invalid()
		val := node{
			s:    "thisinputshouldnotmatter",
			base: invalidBaseToInput,
		}
		table = append(table, val)
	}
	table = append(table,
		node{s: "125", base: "0123456789"},
		node{s: "1111101", base: "01"},
		node{s: "7D", base: "0123456789ABCDEF"},
		node{s: "uoi", base: "choumi"},
		node{s: "bbbbbab", base: "-ab"},
	)
	for _, arg := range table {
		z01.ChallengeMain("atoibaseprog", arg.s, arg.base)
	}
	z01.ChallengeMain("atoibaseprog")
	z01.ChallengeMain("atoibaseprog", "125", "0123456789", "something")
}

// TODO: fix base exercises