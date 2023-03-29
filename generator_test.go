package namegen

import (
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCapitalize(t *testing.T) {
	generator := New()
	generator.Capitalize = true

	testRegex := regexp.MustCompile(`^[A-Z][a-z]+( [A-Z][a-z]+)+$`)

	for i := 0; i < 100; i++ {
		assert.Regexp(t, testRegex, generator.Generate())
	}
}

func TestRhyme(t *testing.T) {
	generator := New()
	generator.Rhyme = true

	for i := 0; i < 100; i++ {
		name := generator.Generate()
		assert.True(t, strings.Contains(name, " "+name[0:1]))
	}
}

func TestSeparatorDash(t *testing.T) {
	generator := New()
	generator.Separator = "-"

	testRegex := regexp.MustCompile(`^[A-Z][a-z]+(-[A-Z][a-z]+)+$`)

	for i := 0; i < 100; i++ {
		assert.Regexp(t, testRegex, generator.Generate())
	}
}

func TestSeparatorDot(t *testing.T) {
	generator := New()
	generator.Separator = "."

	testRegex := regexp.MustCompile(`^[A-Z][a-z]+(\.[A-Z][a-z]+)+$`)

	for i := 0; i < 100; i++ {
		assert.Regexp(t, testRegex, generator.Generate())
	}
}

func TestLowerCase(t *testing.T) {
	generator := New()
	generator.Capitalize = false

	testRegex := regexp.MustCompile(`^[a-z]+( [a-z]+)+$`)

	for i := 0; i < 100; i++ {
		assert.Regexp(t, testRegex, generator.Generate())
	}
}

func TestGenerate(t *testing.T) {
	testRegex := regexp.MustCompile(`^[A-Z][a-z]+( [A-Z][a-z]+)+$`)

	for i := 0; i < 100; i++ {
		assert.Regexp(t, testRegex, Generate())
	}
}
