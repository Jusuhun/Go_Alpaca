package main

import (
	"GO_Alpaca/testadepter"
	"GO_Alpaca/testcontrl"
	"log"
	"net"
	"time"
)

//커맨드맵에서 명령을 찾는 것과 보낸 명령의 응답을 기다리는 맵에서 데이터를 분석하는것의 처리 시간이 주요 포인트 일듯
//TCP 통신의 RTT 시간도 중요 포인트
func main() {
	testadepter.Mapping(testcontrl.Mapping())

	conn, err := net.Dial("tcp", "127.0.0.1:9200")
	if nil != err {
		log.Println(err)
	}

	go func() {
		data := make([]byte, 4096)

		for {
			n, err := conn.Read(data)
			if err != nil {
				log.Println(err)
				return
			}

			testadepter.decodeCmd(data[:n])
			//time.Sleep(time.Duration(1) * time.Millisecond) //Block이면 필요 없지 않나
		}
	}()

	for {
		b, err := testadepter.Marshal()
		if err != nil {
			continue
		}

		n, err := conn.Write(b)
		if err != nil {
			log.Println(err)
			break
		}
		time.Sleep(time.Duration(1) * time.Millisecond)
	}
}
