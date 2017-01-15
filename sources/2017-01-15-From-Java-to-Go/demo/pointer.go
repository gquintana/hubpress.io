package main

import "fmt"

type Article struct {
	Name  string
	Price int
}

// tag::main[]
func RaisePriceByVal(a Article) {
	a.Price = (a.Price * 110) / 100
}

func RaisePriceByRef(a *Article) {
	a.Price = (a.Price * 110) / 100                           // <1>
}

func main() {
	a := Article{"Go in Action", 25}
	fmt.Printf("Article %s %d\n", a.Name, a.Price)
	RaisePriceByVal(a)
	fmt.Printf("By val: Article %s %d\n", a.Name, a.Price)    // <2>
	RaisePriceByRef(&a)                                       // <3>
	fmt.Printf("By ref: Article %s %d\n", a.Name, a.Price)    // <4>
}
// end::main[]
