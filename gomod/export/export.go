package export

import (
	"fmt"
	"rsc.io/quote"
)

type Export struct {
}

func (c Export) ExampleExport() {
	fmt.Println(quote.Go())
}

func (c Export) NewFunction() {
	fmt.Println("Hello World!")
}

func (c Export) FooBar() {
	fmt.Println("Hello Git insteadOf!")
}

func (c Export) FooBarV2() {
	fmt.Println("Hello Git insteadOf in a different version!")
}
