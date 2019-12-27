package main

import (
	_ "GO_Alpaca/errorwork"
	_ "GO_Alpaca/fizzbuzz"

	_ "GO_Alpaca/example"
	"GO_Alpaca/unicode"
)

type selecter interface {
	GetName() string
}

func main() {
	unicode.Execute()
}
