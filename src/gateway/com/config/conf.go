package config

import (
	"encoding/json"
	"github.com/edolphin-ydf/gopherlua-debugger"
    "github.com/cjoudrey/gluahttp"
	"github.com/yuin/gopher-lua"
	luajson  "gateway/com/utils/json"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

var GLOBAL_CONFIG *Config
var conf_file = "../../config.json"
var Host_port map[string]Server
var LUA *LuaHandler

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
type PathMap struct {
	Path []string
}
type LuaHandler struct {
	L *lua.LState
}

func (p *PathMap) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	size := len(p.Path)
	n := rand.Intn(size)
	path := p.Path[n]
	remote, _ := url.Parse(path)
	r.URL.Scheme = remote.Scheme
	r.URL.Host = remote.Host
}

func init() {
	filePtr, err := os.Open(conf_file)
	if err != nil {
		log.Fatal("Open file failed [Err:%s]", err.Error())
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
	//遍历主机
	Host_port = map[string]Server{}
	for _, server := range config.Servers {
		Host_port[strconv.Itoa(server.Port)] = server
	}
	GLOBAL_CONFIG = &config
	LUA = &LuaHandler{lua.NewState()}
	//debug
	lua_debugger.Preload(LUA.L)
	//http
	LUA.L.PreloadModule("http", gluahttp.NewHttpModule(&http.Client{}).Loader)
	//json
	luajson.Preload(LUA.L)

	LUA.HandlerLua(config.InitEvent, config.ScriptPath)
	log.Println("当前进程：" + strconv.Itoa(os.Getpid()))
}

func (l *LuaHandler) HandlerLua(function string, path string) {
	if !(len(function) > 0 && len(path) > 0) {
		return
	}
	//加载lua引擎
	l.L.DoFile(path)
	//L.DoString("GetStr()")
	//事件回调
	fn := l.L.GetGlobal(function)
	if err := l.L.CallByParam(lua.P{Fn: fn, NRet: 1, Protect: true}, nil); err != nil {
		panic(err)
	}
}
