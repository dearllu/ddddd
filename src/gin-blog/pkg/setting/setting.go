package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg *ini.File //配置文件类型的数据类型
	//从配置文件中解析出来的数据
	RunMode        string
	HTTPPort       int
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration //time.duration是变量的类型

    PageSize       int
	JwtSecret      string
)

func init() {
	var err error
	Cfg, err := ini.Load("conf/app.ini") //将配置文件中的信息读取出来
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v",err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
}

func LoadBase()