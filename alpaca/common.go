package alpaca

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

//File thrtjd
type File struct {
	Path    string
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
	file, err := os.Open(input.Path + input.Name)
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

func findFolder(dirname string) ([]string, error) {
	var folderNames []string
	f, err := os.Open(dirname)
	if err != nil {
		return folderNames, err
	}
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return folderNames, err
	}

	for _, file := range files {
		if file.IsDir() {
			folderNames = append(folderNames, file.Name())
		}
	}
	return folderNames, nil
}
func folderSerch(dirname, matchString string) ([]File, error) {
	var findFiles []File
	f, err := os.Open(dirname)
	if err != nil {
		return findFiles, err
	}
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return findFiles, err
	}

	for _, file := range files {
		if file.IsDir() {
			temp, _ := folderSerch(dirname+"/"+file.Name(), matchString)
			findFiles = append(findFiles, temp...)
		} else {
			isMatch, err := regexp.MatchString(matchString, file.Name())
			if err != nil {
				continue
			}
			if isMatch {
				findFiles = append(findFiles, File{Path: dirname + "/", Name: file.Name()})
			}
		}
	}
	return findFiles, nil
}
