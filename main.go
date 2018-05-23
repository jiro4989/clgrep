package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	// 引数チェック
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	args := os.Args[1:]
	nargs := make([]string, 0)

	ignoreCaseFlag := false
	reverseSortFlag := false
	// オプション引数チェック
	for _, v := range args {
		if v == "-h" {
			printHelp()
			return
		}
		if v == "-i" {
			ignoreCaseFlag = true
			continue
		}
		if v == "-r" {
			reverseSortFlag = true
			continue
		}
		nargs = append(nargs, v)
	}

	// オプション引数しか指定されていない場合は終了
	if len(nargs) < 2 {
		printHelp()
		return
	}

	sw := nargs[0] // 検索ワード
	fn := nargs[1] // 読み込みファイル

	// 大小比較なし
	if ignoreCaseFlag {
		sw = strings.ToLower(sw)
	}

	r, err := os.Open(fn)
	if err != nil {
		log.Println(err)
		return
	}
	defer r.Close()

	re := regexp.MustCompile(".*" + sw + ".*")

	sc := bufio.NewScanner(r)
	matches := make([]string, 0) // マッチした段落
	para := make([]string, 0)    // 段落
	for sc.Scan() {
		if err := sc.Err(); err != nil {
			break
		}
		t := sc.Text()
		para = append(para, t)
		if t == "" {
			for _, v := range para {
				iv := v
				if ignoreCaseFlag {
					iv = strings.ToLower(iv)
				}
				if re.MatchString(iv) {
					matches = append(matches, strings.Join(para, "\n"))
					para = make([]string, 0)
					break
				}
			}
			para = make([]string, 0)
		}
	}

	// 逆順フラグが立ってるときは逆順ソート
	if reverseSortFlag {
		reverse(matches)
	}
	matchedText := strings.Join(matches, "\n")
	fmt.Println(matchedText)
}

func printHelp() {
	fmt.Println(`
	clgrep - grep for changelog

	usage:
		clgrep [options] search_word file_name

	options:
		-h Help
		-i Ignore case
		-r Reverse sort
	`)
}

// 配列を逆順ソートして上書きする
func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
