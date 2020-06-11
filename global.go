package goid

import (
	"encoding/base64"
	"fmt"
)

type globalGenerator struct {
	prefix    string
	id        ID
	delimiter string
}

func NewGlobalGenerator(
	opts ...func(*globalGenerator),
) *globalGenerator {
	opt := globalGenerator{
		delimiter: ":",
	}

	for _, x := range opts {
		x(&opt)
	}

	return &opt
}

func (g globalGenerator) Generate(args ...string) ID {
	id := fmt.Sprintf("%s%s%s", g.prefix, g.delimiter, g.id)
	return FromString(base64.URLEncoding.EncodeToString([]byte(id)))
}

func (g globalGenerator) Prefix(prefix string) globalGenerator {
	g.prefix = prefix
	return g
}

func (g globalGenerator) ID(id ID) globalGenerator {
	g.id = id
	return g
}
