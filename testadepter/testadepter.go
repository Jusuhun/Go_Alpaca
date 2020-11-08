package testadepter

import (
	"encoding/json"
	"log"
	"time"
)

type ActionHandle struct {
	done   bool
	result interface{}
}

func (this ActionHandle) Join() {
	for {
		if this.IsDone() {
			break
		}
		time.Sleep(time.Duration(1) * time.Millisecond)
	}
}

func (this ActionHandle) IsDone() bool {
	return this.done
}

func (this ActionHandle) Cancle() {
}

func (this ActionHandle) GetResult() interface{} {
	return this.result
}

var MapCmd map[string]func(interface{})
var SendCmd map[string]interface{}

func SendMessage(object interface{}) interface{} {
	item := PostMessage(object)
	item.Join()
	return item.GetResult()
}

func PostMessage(object interface{}) ActionHandle {
	AppendSendList(object)
	//결과를 받으면 변수에 넣어 달라고 요청
	var action ActionHandle
	//object, action, timeOut

	return action
}

func Mapping(tree map[string]func(interface{})) {
	MapCmd = tree
}

func Marshal() ([]byte, error) {
	b, err := json.Marshal(SendCmd)
	if err != nil {
		return b, err
	}

	//SendCmd Clear

	return b, err
}

func AppendSendList(object interface{}) {
	//SendCmd["Send List"].append
	//추가하기
}

func DecodeCmd(str []byte) {
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
			for _, u := range vv {
				go RunCmd(u)
			}
		}
	}
}

func RunCmd(object interface{}) {
	m := object.(map[string]interface{})

	for k, v := range m {
		if k != "Command" {
			continue
		}

		switch vv := v.(type) {
		case string:
			MapCmd[vv](object)
		}
	}
}
