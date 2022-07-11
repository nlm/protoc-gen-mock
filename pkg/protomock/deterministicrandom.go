package protomock

import (
	"fmt"
	"math/rand"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type DeterministicRandom struct {
	Seed   int64
	rand   *rand.Rand
	source rand.Source
}

func (p *DeterministicRandom) Source() rand.Source {
	if p.source == nil {
		p.source = rand.NewSource(p.Seed)
	}
	return p.source
}

func (p *DeterministicRandom) Rand() *rand.Rand {
	if p.rand == nil {
		p.rand = rand.New(p.Source())
	}
	return p.rand
}

func (p *DeterministicRandom) RandName() string {
	return cases.Title(language.Und).String(left[p.Rand().Intn(len(left))] + " " + right[p.Rand().Intn(len(right))])
}

func (p *DeterministicRandom) RandEmail() string {
	return cases.Lower(language.Und).String(left[p.Rand().Intn(len(left))] + "." + right[p.Rand().Intn(len(right))] + "@example.com")
}

func (p *DeterministicRandom) RandUUID() string {
	hexString := fmt.Sprintf("%0x%0x", p.Rand().Uint64(), p.Rand().Uint64())
	return hexString[0:8] + "-" + hexString[8:12] + "-" + hexString[12:16] + "-" + hexString[16:20] + "-" + hexString[20:32]
}

var deterministicRandom *DeterministicRandom

func init() {
	deterministicRandom = &DeterministicRandom{Seed: 42}
}
