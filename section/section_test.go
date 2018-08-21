package section

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sect = func(s ...string) Section { return s }
var sects = func(s ...Section) Sections { return s }

func TestMatchesRegexp(t *testing.T) {
	re := regexp.MustCompile(`abc`)

	assert.Equal(t, true, sect("*test: foo", "abc").MatchesRegexp(re))
	assert.Equal(t, true, sect("*test: foo", "	abc").MatchesRegexp(re))
	assert.Equal(t, true, sect("*test: foo", "xabcx").MatchesRegexp(re))
	assert.Equal(t, true, sect("	*test: foo", "	xabcx").MatchesRegexp(re))
	assert.Equal(t, false, sect("foobar").MatchesRegexp(re))
	assert.Equal(t, false, sect().MatchesRegexp(re))

	re = regexp.MustCompile(``)
	assert.Equal(t, true, sect("*test: foo", "abc").MatchesRegexp(re))

	re = regexp.MustCompile(`f..`)
	assert.Equal(t, true, sect("*test: foo", "abc").MatchesRegexp(re))
}

func TestMatchesTag(t *testing.T) {
	re := regexp.MustCompile(`main`)

	assert.Equal(t, true, sect("*main.go: foobar", "abc").MatchesTag(re))
	assert.Equal(t, true, sect("*main.go: foobar").MatchesTag(re))
	assert.Equal(t, true, sect("*main.go: ").MatchesTag(re))
	assert.Equal(t, true, sect("*main.go:").MatchesTag(re))
	assert.Equal(t, true, sect("	*main.go:").MatchesTag(re))
	assert.Equal(t, false, sect("*hoge.go: main", "abc").MatchesTag(re))
	assert.Equal(t, false, sect("*hoge.go: main", "main").MatchesTag(re))
	assert.Equal(t, false, sect("*hoge.go: fuga", "main").MatchesTag(re))
	assert.Equal(t, false, sect("*main.go fuga", "main").MatchesTag(re))
	assert.Equal(t, false, sect("main.go: fuga", "main").MatchesTag(re))
	assert.Equal(t, false, sect("main.go fuga", "main").MatchesTag(re))
	assert.Equal(t, false, sect("main.go fuga", "*main.go: foobar", "main").MatchesTag(re))
	assert.Equal(t, false, sect().MatchesTag(re))

	re = regexp.MustCompile(`m...`)
	assert.Equal(t, true, sect("*main.go:").MatchesTag(re))
}

type TestMatchTagData struct {
	expect Sections
	in     Sections
}

func TestMatchTag(t *testing.T) {
	tds := []TestMatchTagData{
		TestMatchTagData{
			expect: sects(
				sect("*main.go: foobar", "abc"),
				sect("*main.go: hogebar", "abc"),
				sect("*main.go: foobar", "abc"),
			),
			in: sects(
				sect("*main.go: foobar", "abc"),
				sect("*main.go: hogebar", "abc"),
				sect("*main.go: foobar", "abc"),
			),
		},
		TestMatchTagData{
			expect: Sections{},
			in: sects(
				sect("*foobar.go: hogebar", "abc"),
				sect("*MAIN.go: hogebar", "abc"),
				sect("*foobar.go: hogebar", "abc"),
			),
		},
		TestMatchTagData{
			expect: Sections{},
			in: sects(
				sect("*foobar.go: main", "abc"),
				sect("*MAIN.go: main", "abc"),
				sect("*foobar.go: main", "abc"),
			),
		},
		TestMatchTagData{
			expect: Sections{},
			in:     Sections{},
		},
	}
	re := regexp.MustCompile(`main`)
	for _, v := range tds {
		assert.Equal(t, v.expect, v.in.MatchTag(re))
	}
}

func TestReverse(t *testing.T) {
	tds := []TestMatchTagData{
		TestMatchTagData{
			expect: sects(
				sect("*main.go: 3", "6"),
				sect("*main.go: 2", "5"),
				sect("*main.go: 1", "4"),
			),
			in: sects(
				sect("*main.go: 1", "4"),
				sect("*main.go: 2", "5"),
				sect("*main.go: 3", "6"),
			),
		},
		TestMatchTagData{
			expect: sects(
				sect("*main.go: 2", "5"),
				sect("*main.go: 1", "4"),
			),
			in: sects(
				sect("*main.go: 1", "4"),
				sect("*main.go: 2", "5"),
			),
		},
		TestMatchTagData{
			expect: sects(sect("*main.go: 1", "4")),
			in:     sects(sect("*main.go: 1", "4")),
		},
		TestMatchTagData{
			expect: Sections{},
			in:     Sections{},
		},
	}
	for _, v := range tds {
		v.in.Reverse()
		assert.Equal(t, v.expect, v.in)
	}
}

func TestTrimSpace(t *testing.T) {
	// TODO
}

func TestMatchDate(t *testing.T) {
	// TODO
}
