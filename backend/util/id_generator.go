package util

import (
	"crypto/rand"
	"io"
	"time"

	"github.com/oklog/ulid/v2"
)


type Identifier struct {
	identifier string
}

type IdentifierGenerator interface {
	Generate() Identifier
}

func GenerateIdentifier() Identifier {
	return IDGenerator.Generate()
}

func (i Identifier) Value() string {
	return i.identifier
}

func (i Identifier) Equal(other Identifier) bool {
	return i.identifier == other.identifier
}

var IDGenerator IdentifierGenerator

func init() {
	IDGenerator = newULIDGenerator(rand.Reader)
}

type ULIDGenerator struct {
	entropy *ulid.MonotonicEntropy
}

func newULIDGenerator(reader io.Reader) *ULIDGenerator {
	return &ULIDGenerator{
		entropy: ulid.Monotonic(reader, 0),
	}
}

func (g *ULIDGenerator) Generate() Identifier {
	id := ulid.MustNew(ulid.Timestamp(time.Now()), g.entropy)

	return Identifier{
		identifier: id.String(),
	}
}