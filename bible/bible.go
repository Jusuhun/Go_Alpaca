package bible

import (
	"gopkg.in/ini.v1"
)

type Bible struct {
	Name        string
	BookName    string
	IniFileName string
}

func (this Bible) GetName() string {
	return "struct Bible : " + this.Name
}

func (this Bible) Execute() {
	this.start()
}

func (self Bible) start() {

	self.find("요 1:1")
	self.find("요 1:2")
	self.find("요 1:3")
}

func (self Bible) find(key string) string {
	cfg, err := ini.Load(self.IniFileName)
	if err != nil {
		return ""
	}

	//각종 타입으로 가져올 수 있습니다.
	return cfg.Section("소제목").Key(key).String()
}
