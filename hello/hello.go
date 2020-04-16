package main

import (
	_ "GO_Alpaca/errorwork"
	_ "GO_Alpaca/fizzbuzz"
	"fmt"

	"GO_Alpaca/alpaca"
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
	basePath := ""
	if true {
		basePath = "../../../test/"
	}

	u := new(unicode.Unicode)
	u.Name = "134"

	b := new(bible.Bible)
	b.Name = "Instance1"
	b.BookName = basePath + "바른성경.txt"
	b.IniFileName = basePath + "소제목.txt"
	b.OutFile = basePath + "바른성경_Out.md"

	mds := new(alpaca.MDSplinder)
	mds.Name = "Instance1"
	mds.InFile = "001_모세오경.md"
	mds.Path = basePath

	dim := new(alpaca.DisplayIniMaker)
	dim.Name = "Instance1"
	dim.IniFileName = "DispMessage.ini"
	dim.Path = basePath + "DIM_Work/"
	dim.Match = "\\.(cpp|c)"

	btm := new(alpaca.BibleToMarkdown)
	btm.Name = "Instance1"
	btm.IniFileName = "BibleInfo.ini"
	btm.BookName = "./개역개정.txt"
	btm.OutPath = basePath + "Output/"

	btm.Execute()

	//start([]selecter{u, b, mds})
	//mdSpinder.InFile = "009_공동서신.md"
	//mdSpinder.Execute()
}

func start(items []selecter) {
	for {
		for index, item := range items {
			fmt.Println(index, " _ ", item.GetName())
		}
		fmt.Println("exit _ program exit")

		var s1 string

		_, err := fmt.Scanln(&s1)

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
