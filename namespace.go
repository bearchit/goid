package goid

import (
	"fmt"
	"strings"
)

type namespaceGenerator struct {
	namespace string
	delimiter string
	names     []string
}

func NewNamespaceGenerator(
	namespace string,
	opts ...func(*namespaceGenerator),
) *namespaceGenerator {
	opt := namespaceGenerator{
		namespace: namespace,
		delimiter: ".",
	}

	for _, x := range opts {
		x(&opt)
	}

	return &opt
}

func (n namespaceGenerator) Names(names ...string) namespaceGenerator {
	n.names = names
	return n
}

func (n namespaceGenerator) Generate() ID {
	return FromString(
		fmt.Sprintf("%s%s%s", n.namespace, n.delimiter, strings.Join(n.names, n.delimiter)),
	)
}

func WithDelimiter(delimiter string) func(opt *namespaceGenerator) {
	return func(g *namespaceGenerator) {
		g.delimiter = delimiter
	}
}
