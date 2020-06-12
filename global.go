package goid

import (
	"encoding/base64"
	"fmt"
)

type globalGenerator struct {
	prefix    string
	generator Generator
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

func WithGenerator(generator Generator) func(*globalGenerator) {
	return func(g *globalGenerator) {
		g.generator = generator
	}
}

func (g globalGenerator) Generate(args ...string) ID {
	id := fmt.Sprintf("%s%s%s", g.prefix, g.delimiter, g.generator.Generate(args...))
	return FromString(base64.URLEncoding.EncodeToString([]byte(id)))
}

func (g globalGenerator) Prefix(prefix string) globalGenerator {
	g.prefix = prefix
	return g
}
