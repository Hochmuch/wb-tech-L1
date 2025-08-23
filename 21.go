package main

import "fmt"

type Bad struct {
	val string
}

func (bad Bad) output() string {
	return bad.val
}

type Good interface {
	printVal()
}

type Adapter struct {
	bad Bad
}

func (a Adapter) printVal() {
	fmt.Println(a.bad.val)
}

func getGood(good Good) {
	good.printVal()
}

func main() {
	var bad = Bad{val: "value"}
	var adapter = Adapter{bad: bad}

	getGood(adapter)
}
