package unicode

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func Execute() {
	start(".")
}

func start(dirname string) {
	f, err := os.Open(dirname)
	if err != nil {
		log.Fatal(err)
	}
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			folder(dirname+"/"+file.Name(), "./export/"+file.Name())
		}
	}
}

func folder(dirname, logName string) {
	f, err := os.Open(dirname)
	if err != nil {
		log.Fatal(err)
	}
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			folder(dirname+"/"+file.Name(), logName+"/"+file.Name())
		} else {
			isContinue, err := regexp.MatchString("\\.(h|cpp|c)", file.Name())
			if err != nil {
				log.Fatal(err)
				continue
			}

			if isContinue {
				os.MkdirAll(logName, os.ModePerm)
				find(dirname+"/"+file.Name(), logName+"/"+file.Name())
			}
		}
	}
}

//파일을 주면 로그를 남긴다.
func find(name, logName string) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineCount := 0

	for scanner.Scan() {
		lineCount++

		line := lineChage(scanner.Text())
		writeLog(line, logName)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("reading standard input:", err)
	}
}

//파일을 주면 로그를 남긴다.
func lineChage(line string) string {
	//제외 경우 판단.
	if true {

		//#include 가 있으면넘어가야 한다.
		if true {
			isContinue, err := regexp.MatchString("#\\s*include", line)
			if err != nil {
				log.Fatal(err)
				return line
			}

			if isContinue {
				return line
			}
		}
		//CLimitSingleInstance SingleInstanceObject(TEXT(_T("WinOlb")));
		//TRACE0(_T("Failed to create status bar\n"));
		//extern _T("C") 도 넘겨야 한다.
		if true {
			isContinue, err := regexp.MatchString("extern", line)
			if err != nil {
				log.Fatal(err)
				return line
			}

			if isContinue {
				return line
			}
		}
		// GetProcAddress 는 예외
		if true {
			isContinue, err := regexp.MatchString("GetProcAddress", line)
			if err != nil {
				log.Fatal(err)
				return line
			}

			if isContinue {
				return line
			}
		}
		// acsc_ 함수들도 제외.
		if true {
			isContinue, err := regexp.MatchString("acsc_", line)
			if err != nil {
				log.Fatal(err)
				return line
			}

			if isContinue {
				return line
			}
		}
		// MIL_TEXT 함수들도 제외.
		if true {
			isContinue, err := regexp.MatchString("MIL_TEXT", line)
			if err != nil {
				log.Fatal(err)
				return line
			}

			if isContinue {
				return line
			}
		}
		// GetBoxTrack 함수들도 제외.
		if true {
			isContinue, err := regexp.MatchString("GetBoxTrack", line)
			if err != nil {
				log.Fatal(err)
				return line
			}

			if isContinue {
				return line
			}
		}
		// GetCrossTrack 함수들도 제외.
		if true {
			isContinue, err := regexp.MatchString("GetCrossTrack", line)
			if err != nil {
				log.Fatal(err)
				return line
			}

			if isContinue {
				return line
			}
		}
	}

	result := line

	//문자열 함수 바꾸기
	if true {
		// atoi -> _wtoi
		if true {
			r, _ := regexp.Compile("atoi")
			result = r.ReplaceAllString(result, "_wtoi")
		}
		// Atof/atof ->_wtol
		if true {
			r, _ := regexp.Compile("[aA]tof")
			result = r.ReplaceAllString(result, "_wtol")
		}
		// strncpy -> wcsncpy
		if true {
			r, _ := regexp.Compile("strncpy")
			result = r.ReplaceAllString(result, "wcsncpy")
		}
		// itoa -> _itow
		if true {
			r, _ := regexp.Compile("itoa")
			result = r.ReplaceAllString(result, "_itow")
		}
		// fopen -> _wfopen
		if true {
			r, _ := regexp.Compile("fopen")
			result = r.ReplaceAllString(result, "_wfopen")
		}
		// _strtime -> _wstrtime
		if true {
			r, _ := regexp.Compile("_strtime")
			result = r.ReplaceAllString(result, "_wstrtime")
		}
		// _strdate -> _wstrdate
		if true {
			r, _ := regexp.Compile("_strdate")
			result = r.ReplaceAllString(result, "_wstrdate")
		}
		// sprintf -> swprintf
		if true {
			r, _ := regexp.Compile("sprintf")
			result = r.ReplaceAllString(result, "swprintf")
		}
		// fprintf ->  fwprintf
		if true {
			r, _ := regexp.Compile("fprintf")
			result = r.ReplaceAllString(result, "fwprintf")
		}
		// fputs -> fputws
		if true {
			r, _ := regexp.Compile("fputs")
			result = r.ReplaceAllString(result, "fputws")
		}
		// fgets -> fgetws
		if true {
			r, _ := regexp.Compile("fgets")
			result = r.ReplaceAllString(result, "fgetws")
		}
		// strlen -> lstrlen 문자열 길이
		if true {
			r, _ := regexp.Compile("strlen")
			result = r.ReplaceAllString(result, "lstrlen")
		}
		// strcmp -> lstrcmp 문자열 비교
		if true {
			r, _ := regexp.Compile("strcmp")
			result = r.ReplaceAllString(result, "lstrcmp")
		}
		// strcpy -> lstrcpy  문자열 복사
		if true {
			r, _ := regexp.Compile("strcpy")
			result = r.ReplaceAllString(result, "lstrcpy")
		}
		// strtok -> wcstok 문자열 가르기
		if true {
			r, _ := regexp.Compile("strtok")
			result = r.ReplaceAllString(result, "wcstok")
		}
		// sscanf -> swscanf 문자열 입력
		if true {
			r, _ := regexp.Compile("sscanf")
			result = r.ReplaceAllString(result, "swscanf")
		}
		// strtol -> wcstol 문자열 -> 숫자
		if true {
			r, _ := regexp.Compile("strtol")
			result = r.ReplaceAllString(result, "wcstol")
		}
		// ltoa -> _ltow  숫자 -> 문자열
		if true {
			r, _ := regexp.Compile("ltoa")
			result = r.ReplaceAllString(result, "_ltow")
		}
	}

	re7, _ := regexp.Compile("(\\\"[^\\\"]*\\\")")
	result = re7.ReplaceAllString(result, "_T(${1})")

	r, _ := regexp.Compile("_T\\s*\\(\\s*_T\\s*\\(\\s*\\\"[^\\\"]*\\\"\\s*\\)\\s*\\)")
	for isOk := true; isOk; {
		find6 := r.FindString(result)
		if find6 != "" {
			test6 := "_T(" + re7.FindString(find6) + ")"
			result = strings.Replace(result, find6, test6, 1)
			isOk = true
		} else {
			isOk = false
		}
	}
	return result
}

func writeLog(log, logName string) {
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
