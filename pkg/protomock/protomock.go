package protomock

import (
	fake "github.com/brianvoe/gofakeit/v6"
	"google.golang.org/protobuf/proto"
)

var faker *fake.Faker

func init() {
	// We want deterministic content generation
	Seed(42)
}

func Seed(seed int64) {
	faker = fake.New(seed)
}

func Mock(msg proto.Message) {
	mockMessage(msg, 0)
}

func Faker() *fake.Faker {
	return faker
}
