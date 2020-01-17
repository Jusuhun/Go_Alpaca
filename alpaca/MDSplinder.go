package alpaca

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type MDSplinder struct {
	Name   string
	InFile string
	Path   string
}

func (this MDSplinder) GetName() string {
	return "struct MDSplinder : " + this.Name
}

func (this MDSplinder) Execute() {
	this.Splin()
}

func (this MDSplinder) Splin() {
	lines := openToLines(this.Path + this.InFile)

	var items [][]string

	items = append(items, make([]string, 0))
	count := len(items) - 1

	for {
		tempItems := spLines(lines[:50])
		for i, item := range tempItems {
			if i != 0 {
				items = append(items, make([]string, 0))
				count = len(items) - 1
			}
			items[count] = append(items[count], item...)
		}
		if len(lines) < 50 {
			break
		}
		lines = lines[50:]
	}

	for i, item := range items {
		if i == 0 {
			continue
		}

		outName := item[0]
		index := fmt.Sprintf("%03d_", i)
		outName = strings.Replace(outName, "# ", index, 1)
		writeFile(item, this.Path+outName+".md",
			os.O_CREATE|os.O_RDWR)
	}

}

func spLines(lines []string) [][]string {
	//var items [][]string
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

//파일을 주면 로그를 남긴다.
func openToLines(name string) []string {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("reading standard input:", err)
	}
	return lines
}

func writeFile(lines []string, logName string, flg int) {
	file, err := os.OpenFile(
		logName,
		flg,
		os.FileMode(0644))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	for _, line := range lines {
		fmt.Fprintln(file, line)
	}
}
