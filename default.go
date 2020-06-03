package goid

var (
	DefaultGenerator     = NewUuidV4Generator(true)
	MockedID             = ID("mocked-id")
	DefaultMockGenerator = MockGenerator{id: MockedID}
)

func Generate() ID {
	return DefaultGenerator.Generate()
}
