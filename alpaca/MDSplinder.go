package alpaca

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

//MDSplinder mds
type MDSplinder struct {
	Name   string
	InFile string
	Path   string
}

//GetName 인스턴스 이름을 받아온다.
func (mds MDSplinder) GetName() string {
	return "struct MDSplinder : " + mds.Name
}

//Execute 실행한다
func (mds MDSplinder) Execute() {
	orgin, _ := readFileString(mds.Path + mds.InFile)

	items := spritMarkDown(orgin)

	markDownSave(items, mds.Path)
}

func spritMarkDown(orgin File) [][]string {
	markDownsMap := make(map[int][][]string)

	for i := 0; ; i++ {
		if len(orgin.Content) < 50 {
			markDownsMap[i] = spLines(orgin.Content)
			break
		} else {
			markDownsMap[i] = spLines(orgin.Content[:50])
			orgin.Content = orgin.Content[50:]
		}
	}

	return mergeMap(markDownsMap)
}

func spLines(lines []string) [][]string {
	items := make([][]string, 1)

	count := 0

	for _, line := range lines {
		isT2, err := regexp.MatchString("^## ", line)
		if err != nil {
			log.Fatal(err)
		}

		if isT2 {
			count = count + 1
			items = append(items, make([]string, 0))
		}

		items[count] = append(items[count], strings.Replace(line, "## ", "# ", 1))
	}

	return items
}

func markDownSave(mdFiles [][]string, path string) {
	for i, item := range mdFiles {
		if i == 0 {
			continue
		}

		outName := item[0]
		index := fmt.Sprintf("%03d_", i)
		outName = strings.Replace(outName, "# ", index, 1)
		writeFile(item, path+outName+".md",
			os.O_CREATE|os.O_RDWR)
	}
}
