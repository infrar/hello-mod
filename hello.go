package hello

import "rsc.io/quote/v3"

// Hello - a demo function
func Hello() string {
	return quote.HelloV3()
}

// Proverb - return a proverb
func Proverb() string {
	return quote.Concurrency()
}