package alpaca

import (
	"bufio"
	"fmt"
	"os"
)

//File thrtjd
type File struct {
	Name    string
	Content []string
}

func mergeMap(markDownsMap map[int][][]string) [][]string {
	keys := make([]int, 0, len(markDownsMap))
	for key := range markDownsMap {
		keys = append(keys, key)
	}

	var ret [][]string
	ret = append(ret, make([]string, 0))
	count := len(ret) - 1

	for key := range keys {
		for i, item := range markDownsMap[key] {
			if i != 0 {
				ret = append(ret, make([]string, 0))
				count = len(ret) - 1
			}
			ret[count] = append(ret[count], item...)
		}
	}
	return ret
}

func readFileString(input string) (File, error) {
	var file File
	file.Name = input

	return readFile(file)
}

//파일을 주면 로그를 남긴다.
func readFile(input File) (File, error) {
	file, err := os.Open(input.Name)
	if err != nil {
		return input, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input.Content = append(input.Content, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return input, err
	}
	return input, nil
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
