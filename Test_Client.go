package main

import (
	"GO_Alpaca/testcontrl"
	"encoding/json"
	"log"
	"net"
	"time"
)

var mapCmd map[string]func(interface{})

func main() {
	mapCmd := testcontrl.Mapping()

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

			decodeCmd(data[:n])
			//time.Sleep(time.Duration(1) * time.Millisecond) //Block이면 필요 없지 않나
		}
	}()

	for {
		b, err := json.Marshal()
		//Clear
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

func appendSendList(object interface{}) {
}

func decodeCmd(str []byte) {
	var f interface{}

	if err := json.Unmarshal(str, &f); err != nil {
		log.Println(err)
		return
	}

	m := f.(map[string]interface{})
	for k, v := range m {
		if k != "Send List" {
			continue
		}

		switch vv := v.(type) {
		case []interface{}:
			for i, u := range vv {
				go runCmd(u)
			}
		}
	}
}

func runCmd(object interface{}) {
	m := object.(map[string]interface{})

	for k, v := range m {
		if k != "Command" {
			continue
		}

		switch vv := v.(type) {
		case string:
			mapCmd[vv](object)
		}
	}
}
