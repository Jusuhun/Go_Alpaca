package unicode

import (
    "log"
    "os"
    "regexp"
    "strings"
    "bufio"
    "fmt"
)


func Execute(){
    start(".")
}

func start(dirname string){
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
            folder(dirname + "/" +file.Name(), "./export/" + file.Name())
        }
    }
}

func folder(dirname, logName string){
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
            folder(dirname + "/"+file.Name(), logName+"/"+file.Name())
        } else{
            isContinue, err := regexp.MatchString("\\.(h|cpp|c)", file.Name())
            if err != nil {
                log.Fatal(err)
                continue
            }

            if isContinue {
                os.MkdirAll(logName, os.ModePerm)
                find(dirname + "/"+file.Name(), logName + "/"+file.Name())
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
func lineChage(line string) (string) {
    //#include 가 있으면넘어가야 한다.
    isContinue, err := regexp.MatchString("#\\s*include", line)
    if err != nil {
        log.Fatal(err)
        return line
    }

    if isContinue {
        return line
    }
    //extern _T("C") 도 넘겨야 한다.


    re7, _ := regexp.Compile("(\\\"[^\\\"]*\\\")")
    result := re7.ReplaceAllString(line, "_T(${1})")

	r, _ := regexp.Compile("_T\\s*\\(\\s*_T\\s*\\(\\s*\\\"[^\\\"]*\\\"\\s*\\)\\s*\\)")
    for isOk := true; isOk; {
        find6 := r.FindString(result)
        if find6 != ""{
            test6 := "_T("+re7.FindString(find6)+")"	
            result = strings.Replace(result, find6, test6, 1)
            isOk = true
        } else{
            isOk = false
        }
    }
    return result
}

func writeLog(log,logName string){
    file, err := os.OpenFile(
        logName,
        os.O_CREATE|os.O_RDWR|os.O_APPEND,
        os.FileMode(0644))
    if err != nil{
        fmt.Println(err)
        return
    }
    defer file.Close()

    fmt.Fprintln(file, log)
}