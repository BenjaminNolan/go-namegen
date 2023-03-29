package namegen

import (
	"math/rand"
	"strings"

	"github.com/samber/lo"
)

type Generator struct {
	Capitalize bool
	Rhyme      bool
	Separator  string
}

func New() *Generator {
	return &Generator{
		Capitalize: true,
		Rhyme:      false,
		Separator:  " ",
	}
}

func (g *Generator) Generate() string {
	noun := Data.Nouns[rand.Intn(len(Data.Nouns))]

	var adjectives []string
	if g.Rhyme {
		adjectives = lo.Filter(Data.Adjectives, func(adjective string, idx int) bool {
			return strings.HasPrefix(adjective, noun[0:1])
		})
		if len(adjectives) == 0 {
			adjectives = Data.Adjectives
		}
	} else {
		adjectives = Data.Adjectives
	}

	res := strings.ReplaceAll(
		strings.Join([]string{
			adjectives[rand.Intn(len(adjectives))],
			noun,
		}, " "),
		" ",
		g.Separator,
	)

	if g.Capitalize {
		res = strings.Title(res)
	}

	return res
}

// Generate returns a random name, equivalent to calling namegen.New().Generate()
func Generate() string {
	return New().Generate()
}
