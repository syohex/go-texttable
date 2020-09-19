# go-texttable ![](https://github.com/syohex/go-texttable/workflows/CI/badge.svg)

## Sample Code

```go
package main

import "github.com/syohex/go-texttable"
import "fmt"

func main () {
	tbl := &texttable.TextTable{}
	tbl.SetHeader("Country", "Capital")

	tbl.AddRow("United States of America", "Washington D.C")
	tbl.AddRow("France", "Paris")
	tbl.AddRow("United Kingdom", "London")
	tbl.AddRow("Japan", "Tokyo")
	tbl.AddRow("Taiwan", "Taipei")

	fmt.Println(tbl.Draw())
}

```

Output of above code is

![go-texttable1](image/go-texttable1.png)


`go-texttable` also supports multibyte characters such as Japanese.

```go
package main

import "github.com/syohex/go-texttable"
import "fmt"

func main () {
	tbl := &texttable.TextTable{}
	tbl.SetHeader("プログラミング言語", "よみがな", "作者")

	tbl.AddRow("Python", "ぱいそん", "グイド・ヴァンロッサム")
	tbl.AddRow("Perl", "ぱーる", "ラリーウォール")
	tbl.AddRow("Ruby", "るびぃ", "まつもとゆきひろ")
	tbl.AddRow("Erlang", "あーらん", "ジョーアームストロング")
	tbl.AddRow("D言語", "でぃーげんご", "ウォルター・ブライト")

	fmt.Println(tbl.Draw())
}
```

![go-texttable2](image/go-texttable2.png)
