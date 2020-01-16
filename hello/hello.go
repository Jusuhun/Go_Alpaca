package main

import (
	_ "GO_Alpaca/errorwork"
	_ "GO_Alpaca/fizzbuzz"
	"fmt"

	"GO_Alpaca/bible"
	_ "GO_Alpaca/example"
	"GO_Alpaca/unicode"
	"strconv"
)

type selecter interface {
	GetName() string
	Execute()
}

func main() {
	u := new(unicode.Unicode)
	u.Name = "134"

	b := new(bible.Bible)
	b.Name = "Instance1"
	b.BookName = "../바른성경.txt"
	b.IniFileName = "../소제목.txt"
	b.OutFile = "../바른성경_Out.md"

	//start([]selecter{u})
	b.Execute()
}

func start(items []selecter) {
	for {
		for index, item := range items {
			fmt.Println(index, " _ ", item.GetName())
		}
		fmt.Println("exit _ program exit")

		var s1 string

		_, err := fmt.Scan(&s1)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("input _ ", s1)

		if s1 == "exit" {
			break
		}

		count, err2 := strconv.Atoi(s1)
		if err2 != nil {
			fmt.Println(err)
			return
		}

		if len(items) > count {
			items[count].Execute()
		}
	}
}
