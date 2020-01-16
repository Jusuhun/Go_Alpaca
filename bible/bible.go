package bible

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"gopkg.in/ini.v1"
)

type Bible struct {
	Name        string
	BookName    string
	IniFileName string
	OutFile     string
}

func (this Bible) GetName() string {
	return "struct Bible : " + this.Name
}

func (this Bible) Execute() {
	this.start()
}

func (self Bible) start() {
	self.openFile()
}

func (self Bible) find(key string) string {
	rkey := strings.ReplaceAll(key, ":", "_")
	cfg, err := ini.Load(self.IniFileName)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	//각종 타입으로 가져올 수 있습니다.
	return cfg.Section("소제목").Key(rkey).String()
}

//파일을 주면 로그를 남긴다.
func (self Bible) openFile() {
	file, err := os.Open(self.BookName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineCount := 0

	for scanner.Scan() {
		lineCount++

		Log := self.lineChage(scanner.Text())
		self.writeLog(Log, self.OutFile)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("reading standard input:", err)
	}
}

//파일을 주면 로그를 남긴다.
func (self Bible) lineChage(line string) string {
	result := line

	r1, _ := regexp.Compile("[가-힣]+")
	isMatch1 := r1.MatchString(result)
	r2, _ := regexp.Compile("장\\s*$")
	isMatch2 := r2.MatchString(result)
	r3, _ := regexp.Compile("([가-힣]+\\s[0-9]+:[0-9]+)")
	find3 := r3.FindString(result)
	isMatch3 := r3.MatchString(result)
	if isMatch3 {
		stitle := self.find(find3)
		if stitle != "" {
			result = "* " + stitle + "\n\n" + result
		}
	} else if isMatch2 {
		result = "### " + result
	} else if isMatch1 {
		result = "## " + result
	}

	return result + "\n"
}

func (self Bible) writeLog(log, logName string) {
	file, err := os.OpenFile(
		logName,
		os.O_CREATE|os.O_RDWR|os.O_APPEND,
		os.FileMode(0644))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Fprintln(file, log)
}
