package main

import (
	"fmt"
)

///////| Testing how embedding works, the go substitute for direct inheritance of methods. |/////////////

type Alpha struct {
	contents string
}

func (a *Alpha) getContents() string {
	return a.contents
}

type Beta struct { //With embedding, the attribute name is identical to the struct name.
	Alpha
}

type Greek interface {
	getContents() string
}

func experiment_embedding() {
	var greek Greek
	a := Alpha{"test-this\n"}
	fmt.Print(a.getContents())

	b := Beta{Alpha{"test-that\n"}}
	fmt.Print(b.getContents())

	fmt.Print(b.Alpha)

	greek = &a
	greek = &b
	fmt.Print(greek) //Embedding does indeed lead to both structs fulfilling the interface conditions. Neat.
}
