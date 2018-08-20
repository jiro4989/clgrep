package section

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

type Section []string

// MatchesRegexp はセクション内のいずれかの箇所で正規表現マッチするかを判定する
func (s Section) MatchesRegexp(re *regexp.Regexp) bool {
	for _, line := range s {
		if re.MatchString(line) {
			return true
		}
	}
	return false
}

// MatchesTag はヘッダのタグで正規表現マッチしたかを判定する
func (s Section) MatchesTag(re *regexp.Regexp) bool {
	if len(s) < 1 {
		return false
	}

	head := s[0]
	head = strings.TrimSpace(head)
	if strings.Index(head, ":") == -1 {
		return false
	}
	if strings.Index(head, "*") == -1 {
		return false
	}
	tag := strings.Split(head, ":")[0]

	return re.MatchString(tag)
}

type Sections []Section

func (ss Sections) MatchTag(re *regexp.Regexp) Sections {
	if len(ss) < 1 {
		return ss
	}

	matches := make([]Section, 0)
	for _, s := range ss {
		if len(s) < 1 {
			continue
		}
		if s.MatchesTag(re) {
			matches = append(matches, s)
		}
	}
	return matches
}

func (ss Sections) Reverse() {
	for i, j := 0, len(ss)-1; i < j; i, j = i+1, j-1 {
		ss[i], ss[j] = ss[j], ss[i]
	}
}

// TrimSpace はセクションのスペースを取り除く。
func (ss Sections) TrimSpace() Sections {
	return nil
}

// MatchDate は指定の日付にマッチしたセクションのみにフィルタする。
func (ss Sections) MatchDate(dt time.Time) Sections {
	return nil
}

func (ss Sections) Print() {
	for _, s := range ss {
		for _, v := range s {
			fmt.Println(v)
		}
		fmt.Println()
	}
}
