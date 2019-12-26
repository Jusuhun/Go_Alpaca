package example

import "bytes"
import "fmt"
import "regexp"

//설명
func ExRegexp() {
	//다음은 패턴이 문자열과 일치하는지를 검사합니다.
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)
	
	//위에서는 문자열 패턴을 바로 사용했지만, 다른 정규식 작업을 위해선 최적화된 Regexp 구조체를 Compile` 해야합니다.
	r, _ := regexp.Compile("p([a-z]+)ch")
	
	//이 구조체에선 많은 메서드를 사용할 수 있습니다. 다음은 이전에 봤던것과 같은 일치 검사입니다.
	fmt.Println(r.MatchString("peach"))
	
	//다음은 정규식과 일치하는 문자열을 찾습니다.
	fmt.Println(r.FindString("peach punch"))
	
	//다음 메서드 또한 첫번째로 매칭되는 문자열을 찾지만 일치하는 텍스트 대신 일치하는 텍스트의 첫 인덱스와 마지막 인덱스를 반환합니다.
	fmt.Println(r.FindStringIndex("peach punch"))

	//Submatch 변형은 전체 패턴 일치와 해당 일치의 부분 일치에 대한 정보를 모두 포함합니다. 예를 들어 다음은 p([a-z]+)ch와 ([a-z]+)에 대한 정보를 모두 반환합니다.
 	fmt.Println(r.FindStringSubmatch("peach punch"))

	 //유사하게 다음은 전체 일치와 부분 일치의 인덱스에 대한 정보를 반환합니다.
 	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	 //All 변형은 입력값의 첫번째만이 아닌 모든 일치 작업에 적용됩니다. 예를 들면 정규식에 대해 모든 일치 항목들을 찾는 경우가 있습니다.
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	//다음의 All 변형은 위에서 본것과 마찬가지로 다른 함수에 대해 사용할 수 있습니다.
  	fmt.Println(r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	//다음 함수의 두번째 인자에 음이 아닌 정수를 전달하면 일치 항목의 갯수를 제한합니다.
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	//위의 예시들은 문자열 인자를 사용하며 MatchString과 같은 이름을 사용했습니다. 함수명에서 String을 없애고 인자로 []byte를 전달할 수도 있습니다.
	fmt.Println(r.Match([]byte("peach")))

	//정규표현식으로 상수를 만들때 Compile의 변형인 MustCompile을 사용할 수 있습니다. 일반 Compile은 반환값이 2개라 상수로 사용할 수 없습니다.
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	//regexp 패키지는 부분문자열을 다른값으로 바꾸는데 사용할 수도 있습니다.
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	//Func 변형을 사용하여 주어진 함수로 일치된 텍스트를 변환시킬 수 있습니다.
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}


func ExRegexpReplace() {
	re1, _ := regexp.Compile("\\.[a-zA-Z]+$")
	s1 := re1.ReplaceAllString("hello example.com", ".net")
                        // 정규표현식에 해당하는 문자열을 지정된 문자열로 바꿈
	fmt.Println(s1) // hello example.net: 맨 마지막에 오는 ".영문"을 .net으로 바꿈

	re2, _ := regexp.Compile("([a-zA-Z,]+) ([a-zA-Z!]+)")
	s2 := re2.ReplaceAllString("Hello, world!", "${2}")
                        // 정규표현식에 해당하는 문자열을 지정된 문자열로 바꿈
	fmt.Println(s2) // world!: ${2}만 있으므로 두 번째로 찾은 문자열 world!만 남김


	re2_1, _ := regexp.Compile("([a-zA-Z]+,)")
	s2_1 := re2_1.ReplaceAllString("Hello, Happy, world!", "_T(${1})")
                        // 정규표현식에 해당하는 문자열을 지정된 문자열로 바꿈
	fmt.Println(s2_1) // world!: ${2}만 있으므로 두 번째로 찾은 문자열 world!만 남김

	re3, _ := regexp.Compile("([a-zA-Z,]+) ([a-zA-Z!]+)")
	s3 := re3.ReplaceAllString("Hello, world!", "${2} ${1}")
                        // 정규표현식에 해당하는 문자열을 지정된 문자열로 바꿈
	fmt.Println(s3) // world! Hello,: ${2} ${1}로 설정했으므로 두 문자열의 위치를 바꿈

	re4, _ := regexp.Compile("(?P<first>[a-zA-Z,]+) (?P<second>[a-zA-Z!]+)")
                        // 찾은 문자열에 ${first}, ${second}로 이름을 정함
	s4 := re4.ReplaceAllString("Hello, world!", "${second} ${first}")
                        // 정규표현식에 해당하는 문자열을 지정된 문자열로 바꿈
	fmt.Println(s4) // world! Hello,: ${second} ${first}로 설정했으므로
                        // 두 문자열의 위치를 바꿈

	re5, _ := regexp.Compile("(?P<first>[a-zA-Z,]+) (?P<second>[a-zA-Z!]+)")
	s5 := re5.ReplaceAllLiteralString("Hello, world!", "${second} ${first}")
                        // 문자열을 바꿀 때 정규표현식 기호를 무시함
	fmt.Println(s5) // ${second} ${first}: 정규표현식 기호를 무시했으므로
                        // ${second} ${first}가 그대로 표시됨
}