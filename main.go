package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	flags "github.com/jessevdk/go-flags"
)

// options オプション引数
type options struct {
	IgnoreCase bool `short:"i" long:"ignore-case" description:"Ignore case"`
	Reverse    bool `short:"r" long:"reverse" description:"Reverse sort"`
	Tag        bool `short:"t" long:"tag" description:"Search tags"`
	Today      bool `long:"today" description:"Only today"`
	NoIndent   bool `long:"no-indent" description:"Remove indent"`
}

func main() {
	var opts options
	args, err := flags.Parse(&opts)
	if err != nil {
		log.Println(err)
		return
	}

	if len(args) < 1 {
		log.Println("Need arguments. Use -h options.")
		return
	}

	sw := args[0] // 検索ワード
	fn := args[1] // 読み込みファイル

	// 大小比較なし
	if opts.IgnoreCase {
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
		if opts.NoIndent {
			t = strings.Replace(t, "\t", "", 1)
		}
		para = append(para, t)
		if t == "" {
			for _, v := range para {
				iv := v
				if opts.IgnoreCase {
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
	if opts.Reverse {
		reverse(matches)
	}
	matchedText := strings.Join(matches, "\n")
	fmt.Println(matchedText)
}

// 配列を逆順ソートして上書きする
func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
