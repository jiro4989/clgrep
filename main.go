package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	flags "github.com/jessevdk/go-flags"
)

// options オプション引数
type options struct {
	IgnoreCase  bool   `short:"i" long:"ignore-case" description:"Ignore case"`
	Reverse     bool   `short:"r" long:"reverse" description:"Reverse sort"`
	Tag         bool   `short:"t" long:"tag" description:"Search tags"`
	TodayFormat string `long:"today-format" description:"Only today"`
	ShowIndent  bool   `long:"show-indent" description:"Show indent"`
}

// opts はコマンドライン引数です
var (
	opts  options
	today string
)

func main() {
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

	today = time.Now().Format(opts.TodayFormat)

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

	swRe := regexp.MustCompile(".*" + sw + ".*")

	sc := bufio.NewScanner(r)
	matches := make([]string, 0) // マッチした段落
	para := make([]string, 0)    // 段落
	appendFlag := true
	for sc.Scan() {
		if err := sc.Err(); err != nil {
			break
		}
		t := sc.Text()

		// 今日の日付のみ設定があるときだけフラグ操作
		if opts.TodayFormat != "" && containsDateString(t) {
			if isToday(t) {
				appendFlag = true
			} else {
				appendFlag = false
			}
		}
		if !appendFlag {
			continue
		}

		// デフォルトではインデントは削除
		// 指定があるときだけインデントを残す
		if !opts.ShowIndent {
			t = strings.Replace(t, "\t", "", 1)
		}
		para = append(para, t)

		// 空文字を段落の区切れ目と判定
		if t == "" {
			// 段落単位で検索ワードの出現を検査
			// 段落内にワードが見つかったらその段落はまるごとmatchに追加してbreak
			for _, v := range para {
				iv := v
				if opts.IgnoreCase {
					iv = strings.ToLower(iv)
				}
				if swRe.MatchString(iv) {
					matches = append(matches, strings.Join(para, "\n"))
					para = make([]string, 0)
					break
				}
			}
			para = make([]string, 0)
		}
	}

	// 残ってる可能性があるので
	if 0 < len(para) {
		matches = append(matches, strings.Join(para, "\n"))
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

// isToday は文字列が今日の日付かどうかを判定します。
func isToday(s string) bool {
	return strings.HasPrefix(s, today)
}

// isDateString は文字列が日付文字列を含むかを判定します。
func containsDateString(s string) bool {
	l := len(opts.TodayFormat)
	if len(s) < l {
		return false
	}
	s2 := s[:l]
	_, err := time.Parse(opts.TodayFormat, s2)
	return err == nil
}
