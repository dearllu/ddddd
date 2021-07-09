package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)
var (
	Cfg *ini.File
	RunMode string
	HTTPPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
	PageSize int
	JwtSecret string
)
func init() {
	var err error
<<<<<<< Updated upstream
	Cfg, err = ini.Load("conf/app.ini")
=======
	Cfg, err = ini.Load("conf/app.ini") //将配置文件中的信息读取出来
>>>>>>> Stashed changes
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	LoadBase()
	LoadServer()
	LoadApp()
}
<<<<<<< HEAD

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v",err)
=======
func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}
func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
>>>>>>> dev
	}
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
<<<<<<< HEAD
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

=======
	WriteTimeout =  time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}
>>>>>>> dev
func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}
<<<<<<< HEAD

<<<<<<< Updated upstream
=======
>>>>>>> dev
	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
=======
func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug") //解析出来RunMode
}

func LoadServer() {
	sec, err := Cfg.GetSection("server") //SEC是SERVER类型的结构体
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v",err)
	}
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug") //解析出来RunMode

	//获取对应参数
	//HTTPPort = sec.Key("HTTP_PORT").MustString("de")//.MustString("debug")
	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp(){
	sec, err := Cfg.GetSection("app") //读取配置文件的APP部分
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v",err)
	}
	//获取对应参数
	JwtSecret = sec.Key("JWT.SECRET").MustString("!@)*#)!@U#@*!@!)")
>>>>>>> Stashed changes
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}