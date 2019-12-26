package errorwork

import (
    "log"
    "os"
    "regexp"
    "strconv"
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
            folder(dirname + "/"+file.Name(), file.Name())
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
            folder(dirname + "/"+file.Name(), logName)
        } else{
            find(dirname + "/"+file.Name(), logName)
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

        isOk, number, test := lineCheck(scanner.Text())
        if isOk {
            //Log 남기기
            writeLog(fmt.Sprint(name, "\t",lineCount, "\t", number, "\t", test), logName)
        }  
    }
    if err := scanner.Err(); err != nil {
        fmt.Println("reading standard input:", err)
    }
}

//파일을 주면 로그를 남긴다.
func lineCheck(line string) (bool,int,string) {
    test := strings.Split(line, "//")[0]
    test = strings.TrimSpace(test)

    matched, _ := regexp.MatchString("return\\s+[0-9]+", test)
    if(matched){
        r := regexp.MustCompile("[^0-9]+")
        numberStr := r.ReplaceAllString(test, "")
        number,_ := strconv.Atoi(numberStr)
        if number != 0{
            return true,number,test
        }
    }

    matched, _ = regexp.MatchString("generateErrorCode\\s*\\(\\s*[0-9]+", test)
    if(matched){
        r := regexp.MustCompile("[^0-9]+")
        numberStr := r.ReplaceAllString(test, "")
        number,_ := strconv.Atoi(numberStr)
        if number != 0{
            return true,number,test
        }
    }

    matched, _ = regexp.MatchString("ret\\s*\\=\\s*[0-9]+", test)
    if(matched){
        r := regexp.MustCompile("[^0-9]+")
        numberStr := r.ReplaceAllString(test, "")
        number,_ := strconv.Atoi(numberStr)
        if number != 0{
            return true,number,test
        }
    }

    return false, 0, ""
}

func writeLog(log,logName string){
    file, err := os.OpenFile(
        logName+".tsv",
        os.O_CREATE|os.O_RDWR|os.O_APPEND,
        os.FileMode(0644))
    if err != nil{
        fmt.Println(err)
        return
    }
    defer file.Close()

    fmt.Fprintln(file, log)
}