package testcontrl

//유닛별로 이 쓰래드 하나씩이라고 봐야함

import (
	"GO_Alpaca/testadepter"
	"Go_Alpaca/testobject"
)

var life bool

func Mapping() map[string]func(interface{}) {
	var mapCmd map[string]func(interface{})

	//명령어 함수 맵을 만듬
	mapCmd["run"] = func(object interface{}) {
		//make Result
		//appendSendList(Result)
	}

	mapCmd["stop"] = func(object interface{}) {
		//make Result
		//appendSendList(Result)
	}

	mapCmd["standby"] = func(object interface{}) {
		//make Result
		//appendSendList(Result)
	}

	mapCmd["GetCommandList"] = func(object interface{}) {
		//make Result
		//appendSendList(Result)
	}

	mapCmd["ExecuteLoadReady"] = func(object interface{}) {
		axisY := testobject.Motion("motion", "LoadY", nil)   //objectName
		cylTool := testobject.Cylinder("cylinder", "ToolUD") //objectName
		action := testadepter.PostMessage(cylTool.Up())
		result := testadepter.SendMessage(axisY.Move("Load"))
		action.Join()
		result2 := action.GetResult()
	}

	mapCmd["StepLoading"] = func(object interface{}) {
	}

	return mapCmd
}

func run() {
	life = true
	go doRunStep()
}

func stop() {
	life = false
}

func standby() {
	//필요에 따라 안전위치로
}

func doRunStep() {
	conti := false
	step := 4
	for life || conti {
		conti = false
		switch step {
		case 0:
			testadepter.RunCmd("ExecuteLoadReady")
			step = 1
			break
			//next wait jump //option continue
		case 1:
		case 2:
		case 3:
		case 4:
		}
	}
}
