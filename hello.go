package hello

import "rsc.io/quote/v3"

func Hello() string {
	return "v0.2.0" + quote.HelloV3()
}

func Proverb() string {
	return quote.Concurrency()
}
