package section

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchesRegexp(t *testing.T) {
	f := func(s ...string) Section { return s }
	re := regexp.MustCompile(`abc`)

	assert.Equal(t, true, f("*test: foo", "abc").MatchesRegexp(re))
	assert.Equal(t, true, f("*test: foo", "xabcx").MatchesRegexp(re))
	assert.Equal(t, false, f("foobar").MatchesRegexp(re))
	assert.Equal(t, false, f().MatchesRegexp(re))

	re = regexp.MustCompile(``)
	assert.Equal(t, true, f("*test: foo", "abc").MatchesRegexp(re))

	re = regexp.MustCompile(`f..`)
	assert.Equal(t, true, f("*test: foo", "abc").MatchesRegexp(re))
}

func TestMatchesTag(t *testing.T) {
	f := func(s ...string) Section { return s }
	re := regexp.MustCompile(`main`)

	assert.Equal(t, true, f("*main.go: foobar", "abc").MatchesTag(re))
	assert.Equal(t, true, f("*main.go: foobar").MatchesTag(re))
	assert.Equal(t, true, f("*main.go: ").MatchesTag(re))
	assert.Equal(t, true, f("*main.go:").MatchesTag(re))
	assert.Equal(t, false, f("*hoge.go: main", "abc").MatchesTag(re))
	assert.Equal(t, false, f("*hoge.go: main", "main").MatchesTag(re))
	assert.Equal(t, false, f("*hoge.go: fuga", "main").MatchesTag(re))
	assert.Equal(t, false, f("*main.go fuga", "main").MatchesTag(re))
	assert.Equal(t, false, f("main.go: fuga", "main").MatchesTag(re))
	assert.Equal(t, false, f("main.go fuga", "main").MatchesTag(re))
	assert.Equal(t, false, f("main.go fuga", "*main.go: foobar", "main").MatchesTag(re))
	assert.Equal(t, false, f().MatchesTag(re))

	re = regexp.MustCompile(`m...`)
	assert.Equal(t, true, f("*main.go:").MatchesTag(re))
}

func TestMatchTag(t *testing.T) {
	// TODO
}

func TestReverse(t *testing.T) {
	// TODO
}

func TestMatchDate(t *testing.T) {
	// TODO
}
