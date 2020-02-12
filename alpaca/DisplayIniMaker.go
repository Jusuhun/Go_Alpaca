package alpaca

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"

	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
	"gopkg.in/ini.v1"
)

//DisplayIniMaker dim
type DisplayIniMaker struct {
	Name        string
	Path        string
	IniFileName string
	Match       string //"\\.(h|cpp|c)"
}

//GetName 인스턴스 이름을 받아온다.
func (dim DisplayIniMaker) GetName() string {
	return "struct DisplayIniMaker : " + dim.Name
}

//Execute 실행한다
func (dim DisplayIniMaker) Execute() {
	files, _ := folderSerch(dim.Path+"/", dim.Match)

	os.MkdirAll(dim.Path+"_Export", os.ModePerm)

	// 출력파일 생성
	os.Create(dim.Path + "_Export/" + dim.IniFileName)
	// fo, err := os.Create(dim.Path + "_Export" + dim.IniFileName)
	// if err != nil {
	//     panic(err)
	// }
	// fo.Close()
	cfg, err := ini.Load(dim.Path + "_Export/" + dim.IniFileName)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		orgin, _ := readFile(file)
		_ = orgin
		result := File{}
		result.Path = strings.Replace(file.Path, dim.Path, dim.Path+"_Export", 1)
		result.Name = orgin.Name

		count := 0
		keyBase := orgin.Name
		for _, line := range orgin.Content {
			key := fmt.Sprintf("%s_%03d", keyBase, count+1)

			line = eucKRDecoder(line)

			ischaged, chage, text := lineChage(line, key)
			addComment := false
			for ischaged {
				cfg.Section("KOREAN").Key(key).SetValue(text)
				cfg.SaveTo(dim.Path + "_Export/" + dim.IniFileName)
				addComment = true
				count++

				ischaged, chage, text = lineChage(chage, key)
			}
			if addComment {
				chage = chage + "//200212 auto.JSH"
			}
			chage = eucKREncoder(chage)
			result.Content = append(result.Content, chage)
		}

		os.MkdirAll(result.Path, os.ModePerm)
		writeFile(result.Content, result.Path+result.Name, os.O_CREATE|os.O_RDWR)
	}
}

//기준 문자열에서 바꿀 문자를 찾고, Key값으로 변경한다.
//기준 문자열과, Key값을 준다
//변경여부와, 변경이 적용된 문자열과, key값으로 변경된 문자열을 반환한다.
func lineChage(line, key string) (bool, string, string) {
	re, _ := regexp.Compile("_T\\s*\\(\\s*\"[^\"]*[ㄱ-ㅎ|ㅏ-ㅣ|가-힣]+[^\"]*\"\\s*\\)")
	re2nd, _ := regexp.Compile("[가-힣]+")
	find := re.FindString(line)
	find2 := re2nd.FindString(find)
	if find2 == "" {
		re2, _ := regexp.Compile("\"[^\"]*[ㄱ-ㅎ|ㅏ-ㅣ|가-힣]+[^\"]*\"")
		find = re2.FindString(line)
		find2 = re2nd.FindString(find)
	}

	if find2 == "굴림" {
		find2 = ""
	}

	isOk := (find2 != "")

	if isOk {
		chageMsg := fmt.Sprintf("ReadDisplayMessage(_T(\"%s\"))", key)
		result := strings.Replace(line, find, chageMsg, 1)

		re3, _ := regexp.Compile("[^\"]*[ㄱ-ㅎ|ㅏ-ㅣ|가-힣]+[^\"]*")
		value := re3.FindString(find)

		return isOk, result, value
	}

	return isOk, line, ""
}

func eucKRDecoder(input string) string {
	var bufs bytes.Buffer
	wr := transform.NewWriter(&bufs, korean.EUCKR.NewDecoder())
	wr.Write([]byte(input))
	wr.Close()
	return bufs.String()
}

func eucKREncoder(input string) string {
	var bufs bytes.Buffer
	wr := transform.NewWriter(&bufs, korean.EUCKR.NewEncoder())
	wr.Write([]byte(input))
	wr.Close()
	return bufs.String()
}
