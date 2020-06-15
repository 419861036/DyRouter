package config

import (
	"encoding/json"
	"fmt"
	"github.com/yuin/gopher-lua"
	"log"
	"os"
	"strconv"
)

var GLOBAL_CONFIG *Config
var conf_file = "config.json"

type Proxy struct {
	Location    string
	Path        []string
	BeforeEvent string
	AfterEvent  string
}
type Server struct {
	Port       int
	Proxys     []Proxy
	InitEvent  string
	Name       string
	ScriptPath string
}
type Config struct {
	Hostname   string
	Logfile    string
	Servers    []Server
	ScriptPath string
	InitEvent  string
}

func init() {
	filePtr, err := os.Open(conf_file)
	if err != nil {
		fmt.Println("Open file failed [Err:%s]", err.Error())
		return
	}
	defer filePtr.Close()
	var config Config
	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&config)
	if err != nil {
		log.Println("Decoder failed " + err.Error())
	} else {
		log.Println(config)
	}
	GLOBAL_CONFIG = &config
	HandlerLua(config.InitEvent, config.ScriptPath)
	log.Println("当前进程：" + strconv.Itoa(os.Getpid()))
}

func HandlerLua(function string, path string) {
	if !(len(function) > 0 && len(path) > 0) {
		return
	}
	//加载lua引擎
	L := lua.NewState()
	L.OpenLibs()
	defer L.Close()
	L.DoFile(path)
	//L.DoString("GetStr()")
	fn := L.GetGlobal(function)
	if err := L.CallByParam(lua.P{Fn: fn, NRet: 1, Protect: true}, nil); err != nil {
		panic(err)
	}
}
