package option

// options オプション引数
type Options struct {
	IgnoreCase  bool   `short:"i" long:"ignore-case" description:"Ignore case"`
	Reverse     bool   `short:"r" long:"reverse" description:"Reverse sort"`
	Tag         bool   `short:"t" long:"tag" description:"Search tags"`
	Version     func() `short:"v" long:"version" description:"Version"`
	TodayFormat string `long:"today-format" description:"Only today" default:"2006/01/02"`
	ShowIndent  bool   `long:"show-indent" description:"Show indent"`
}
