package goid

import (
	"fmt"
	"strings"
)

type namespaceGenerator struct {
	namespace string
	delimiter string
}

func NewNamespaceGenerator(
	namespace string,
	opts ...func(*namespaceGenerator),
) Generator {
	opt := namespaceGenerator{
		namespace: namespace,
		delimiter: ".",
	}

	for _, x := range opts {
		x(&opt)
	}

	return &opt
}

func (n namespaceGenerator) Generate(args ...string) ID {
	return FromString(
		fmt.Sprintf("%s%s%s", n.namespace, n.delimiter, strings.Join(args, n.delimiter)),
	)
}

func WithDelimiter(delimiter string) func(opt *namespaceGenerator) {
	return func(g *namespaceGenerator) {
		g.delimiter = delimiter
	}
}
