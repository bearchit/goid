package goid

import (
	"encoding/base64"
	"fmt"
	"regexp"
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

type GlobalID struct {
	Prefix    string
	Delimiter string
	ID        ID
}

func ParseGlobalID(id ID) (*GlobalID, error) {
	decoded, err := base64.URLEncoding.DecodeString(id.String())
	if err != nil {
		return nil, err
	}

	reg := regexp.MustCompile(`(.+)(:)(.+)`)
	tokens := reg.FindStringSubmatch(string(decoded))

	if len(tokens) != 4 {
		return nil, fmt.Errorf("invalid GlobalID format: %v", string(decoded))
	}

	return &GlobalID{
		Prefix:    tokens[1],
		Delimiter: tokens[2],
		ID:        FromString(tokens[3]),
	}, nil
}
