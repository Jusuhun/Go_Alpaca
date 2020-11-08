package testcontrl

import "Go_Alpaca/testobject"

func Mapping() map[string]func(interface{}) {
	var mapCmd map[string]func(interface{})

	//명령어 함수 맵을 만듬
	mapCmd["run"] = func(object interface{}) {
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
		axisY := testobject.Motion(Name:="LoadY") //objectName
		cylTool := testobject.Cylinder(Name:="ToolUD")              //objectName
		action := PostMessage(cylTool.Up()) //adepter
		result := SendMessage(axisY.Move("Load")) //adepter
		action.IsDone()
		action.Join()
		action.Cancel()
		result2 := action.GetResult()
	}

	mapCmd["StepLoading"] = func(object interface{}) {
		//커맨드 리스트에 접근할 권한이 필요 그리고 함수화
	}

	return mapCmd
}

func doRunStep() {
	step := 4
	switch step {
	case 0:
		executeLoadReady() //커맨드 리스트에 접근할 권한이 필요 그리고 함수화
		step = 1
		//next wait jump //option continue
	case 1:
	case 2:
	case 3:
	case 4:
	}
}

func executeLoadReady() {
	axisY := testobject.Motion(Name:="LoadY") //objectName
	cylTool := testobject.Cylinder(Name:="ToolUD")              //objectName
	action := PostMessage(cylTool.Up()) //adepter
	result := SendMessage(axisY.Move("Load")) //adepter
	action.IsDone()
	action.Join()
	action.Cancel()
	result2 := action.GetResult()
}
