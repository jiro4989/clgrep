package section

import (
	"fmt"
	"regexp"
	"time"
)

type Section []string

func (s Section) MatchesRegexp(re *regexp.Regexp) bool {
	for _, line := range s {
		if re.MatchString(line) {
			return true
		}
	}
	return false
}

func Print(ss []Section) {
	for _, s := range ss {
		for _, v := range s {
			fmt.Println(v)
		}
		fmt.Println()
	}
}

func SearchHeader(ss []Section, re *regexp.Regexp) []Section {
	matches := make([]Section, 0)
	for _, s := range ss {
		if len(s) < 1 {
			continue
		}
		header := s[0]
		if re.MatchString(header) {
			matches = append(matches, s)
		}
	}
	return matches
}

func Reverse(ss []Section) {
	for i, j := 0, len(ss)-1; i < j; i, j = i+1, j-1 {
		ss[i], ss[j] = ss[j], ss[i]
	}
}

// TrimSpace はセクションのスペースを取り除く。
func TrimSpace(ss []Section) []Section {
	return nil
}

// MatchDate は指定の日付にマッチしたセクションのみにフィルタする。
func MatchDate(ss []Section, dt time.Time) []Section {
	return nil
}
