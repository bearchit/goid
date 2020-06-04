# goid

## Get Started
```
go get github.com/bearchit/goid 
```

## Usage
```go
package main

import (
	"fmt"
	"github.com/bearchit/goid"
)

func main() {
	// Create a new ID generator
	// Using dash character with UUIDv4 string
	g := goid.NewUuidV4Generator(true)

	// Generate ID
	id := g.Generate()

	fmt.Println(id.Nil())    // Returns boolean whether id is nil or not
	fmt.Println(id.String()) // Converts ID to string
	fmt.Println(id)

	target := g.Generate()
	ids := goid.NewIDs(
		target,
		g.Generate(),
		g.Generate(),
	)
	fmt.Println(ids)
	fmt.Println(ids.Contains(target))

	// Create a UUIDv4 generator without dash
	g2 := goid.NewUuidV4Generator(false)
	fmt.Println(g2.Generate())
}
```

### Namespace Generator

#### Default delimiter
```go
g := goid.NewNamespaceGenerator("com")
id := g.Generate("github", "bearchit", "goid")
// id = com.github.bearchit.goid
```

#### Custom delimiter
```go
g := goid.NewNamespaceGenerator("com", goid.WithDelimiter("-"))
id := g.Generate("github", "bearchit", "goid")
// id = com-github.bearchit-goid
```
