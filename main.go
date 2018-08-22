package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	flags "github.com/jessevdk/go-flags"
	options "github.com/jiro4989/clgrep/internal/options"
	"github.com/jiro4989/clgrep/section"
)

// opts はコマンドライン引数です
var (
	Version string
)

// オプションでない引数の数によって挙動が変化する
// 引数1つの場合 -> 標準入力に対してgrep
// 引数2つの場合 -> 第二引数は検索対象のファイル
func main() {
	var opts options.Options
	opts.Version = func() {
		fmt.Println(Version)
		os.Exit(0)
	}

	args, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(0)
	}

	if len(args) < 1 {
		fmt.Println("Need arguments. args=", args)
		os.Exit(1)
	}

	// 引数と検索オプションにマッチしたセクションをgrep
	ss, err := clgrep(args, opts)
	if err != nil {
		panic(err)
	}

	ss.Print()
}

func clgrep(args []string, opts options.Options) (section.Sections, error) {
	l := len(args)
	if l < 1 {
		return nil, errors.New("引数が不足しています。")
	}

	// 引数が一つの場合は標準入力からデータ読み取り
	// 引数が２つ以上のときは、ファイル読み取り
	var r *os.File
	sw := args[0]
	if l < 2 {
		r = os.Stdin
	} else {
		var err error
		r, err = os.Open(args[1])
		if err != nil {
			return nil, err
		}
		defer r.Close()
	}

	// オプション指定がある場合はIgnoreCaseを有効化
	var re *regexp.Regexp
	if opts.IgnoreCase {
		re = regexp.MustCompile(`(?i)` + sw)
	} else {
		re = regexp.MustCompile(sw)
	}

	// セクション内でマッチしたものにフィルタ1
	ss, err := findMatchedSections(r, re)
	if err != nil {
		return nil, err
	}

	if len(ss) < 1 {
		return ss, nil
	}

	// ヘッダ行から検索フィルタ
	if opts.Tag {
		ss = ss.MatchTag(re)
	}

	if len(ss) < 1 {
		return ss, nil
	}

	// 逆順ソート
	if opts.Reverse {
		ss.Reverse()
	}

	if !opts.ShowIndent {
		ss = ss.TrimSpace()
	}

	return ss, nil
}

func findMatchedSections(r *os.File, re *regexp.Regexp) (section.Sections, error) {
	var sect section.Section = make([]string, 0)
	matchedSections := make([]section.Section, 0)

	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := sc.Text()
		if isBlankLine(line) {
			if sect.MatchesRegexp(re) {
				matchedSections = append(matchedSections, sect)
			}
			sect = make([]string, 0)
			continue
		}
		sect = append(sect, line)
	}
	if 0 < len(sect) {
		if sect.MatchesRegexp(re) {
			matchedSections = append(matchedSections, sect)
		}
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}
	return matchedSections, nil
}

func isBlankLine(s string) bool {
	return strings.TrimSpace(s) == ""
}
