package alpaca

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"gopkg.in/ini.v1"
)

type BibleToMarkdown struct {
	Name        string
	BookName    string
	IniFileName string
	OutPath     string
}

func (btm BibleToMarkdown) GetName() string {
	return "struct Bible : " + btm.Name
}

func (btm BibleToMarkdown) Execute() {
	origin, err := readFileString(btm.BookName)
	if err != nil {
		fmt.Println(err)
		return
	}

	cfg, err := ini.Load(btm.IniFileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	// if true {
	// 	for _, line := range origin.Content {
	// 		re, _ := regexp.Compile("^[0-9|a-z|A-Z]+\\s[0-9]+:[0-9]+")
	// 		find := re.FindString(line)
	// 		if find == "" {
	// 			fmt.Println(line)
	// 		}
	// 	}
	// }

	// return

	var origin2 File
	if true {
		for _, line := range origin.Content {
			re, _ := regexp.Compile("^[0-9|a-z|A-Z]+\\s[0-9]+:[0-9]+")
			re2nd, _ := regexp.Compile("^[0-9|a-z|A-Z]+")
			find := re.FindString(line)
			find2 := re2nd.FindString(find)
			exchage := cfg.Section("English").Key(find2).String()
			ccc := strings.Replace(find, find2, exchage, 1)
			line2 := strings.Replace(line, find, ccc, 1)
			origin2.Content = append(origin2.Content, line2)
		}
	}

	// if true {
	// 	for _, line := range origin2.Content {
	// 		re, _ := regexp.Compile("^[가-힣]+\\s[0-9]+:[0-9]+")
	// 		find := re.FindString(line)
	// 		re2nd, _ := regexp.Compile("^[가-힣]+")
	// 		find2 := re2nd.FindString(find)
	// 		if find == "" || find2 == "" {
	// 			fmt.Println(line)
	// 		}
	// 	}
	// }

	// return

	var origin3 File
	if true {
		oldTitle := ""
		oldSubTitle := ""
		for _, line := range origin2.Content {
			re, _ := regexp.Compile("^[가-힣]+\\s[0-9]+:[0-9]+")
			find := re.FindString(line)
			re2nd, _ := regexp.Compile("^[가-힣]+")
			find2 := re2nd.FindString(find)
			exchage := cfg.Section("KOR").Key(find2).String()
			title := "## " + exchage

			re3nd, _ := regexp.Compile("[0-9]+:")
			find3 := re3nd.FindString(find)
			find3 = strings.Replace(find3, ":", "장", 1)
			subTitle := "### " + exchage + " " + find3

			key := strings.Replace(find, ":", "_", 1)
			ttt := cfg.Section("소제목").Key(key).String()
			subSubTitle := ""
			if ttt != "" {
				subSubTitle = "* " + ttt
			}

			//Old Title 비교
			if oldTitle != title {
				oldTitle = title
				origin3.Content = append(origin3.Content, title, "")
			}

			//Old Sub Title 비교
			if oldSubTitle != subTitle {
				oldSubTitle = subTitle
				origin3.Content = append(origin3.Content, subTitle, "")
			}

			//Old 소제목 찾기
			if subSubTitle != "" {
				origin3.Content = append(origin3.Content, subSubTitle, "")
			}

			// 본문
			origin3.Content = append(origin3.Content, line, "")
		}
	}

	path := btm.OutPath
	var output []File
	if true {
		items := spritMarkDown(origin3)
		for i, item := range items {
			if len(item) < 1 {
				continue
			}

			var file File
			file.Content = item

			index := fmt.Sprintf("%03d_", i)
			outName := strings.Replace(item[0], "# ", index, 1)
			file.Name = outName + ".md"

			title := strings.Replace(item[0], "# ", "", 1)
			outPath := cfg.Section("분류").Key(title).String()
			file.Path = path + outPath + "/"

			output = append(output, file)
		}
	}

	for _, file := range output {
		os.MkdirAll(file.Path, os.ModePerm)
		writeFile(file.Content, file.Path+file.Name,
			os.O_CREATE|os.O_RDWR)
	}
}
