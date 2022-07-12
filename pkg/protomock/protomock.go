package protomock

import (
	fake "github.com/brianvoe/gofakeit/v6"
)

func init() {
	// We want deterministic content generation
	fake.Seed(42)
}
