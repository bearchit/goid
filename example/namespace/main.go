package main

import (
	"fmt"

	"github.com/bearchit/goid"
)

func main() {
	g1 := goid.NewNamespaceGenerator("com")
	id1 := g1.Generate("github", "bearchit", "goid")
	fmt.Println(id1.String())

	g2 := goid.NewNamespaceGenerator("com", goid.WithDelimiter("-"))
	id2 := g2.Generate("github", "bearchit", "goid")
	fmt.Println(id2.String())

	// Output
	// com.github.bearchit.goid
	// com-github-bearchit-goid
}
