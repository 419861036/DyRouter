package core

import (
	"../config"
	lua "github.com/yuin/gopher-lua"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"time"
)

var SchemeMap map[string]Dispath /*创建集合 */

func init() {
	var dispath Dispath
	SchemeMap = make(map[string]Dispath)
	dispath = Proxy{}
	SchemeMap["http"] = dispath
	dispath = HttpProxy{}
	SchemeMap["h"] = dispath
	dispath = StreamProxy{}
	SchemeMap["default"] = dispath

}

type match struct {
	path        []string
	beforeEvent string
	afterEvent  string
}

var req *http.Request
var res http.ResponseWriter

func Handdler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	addrs := ctx.Value("local-addr").(*net.TCPAddr)
	server := config.Host_port[strconv.Itoa(addrs.Port)]
	//路径匹配
	var bMatch match
	var lMatch match
	for _, ser := range server.Proxys {
		match, _ := regexp.MatchString(ser.Location, r.URL.Path)
		if match {
			bMatch.path = ser.Path
			bMatch.afterEvent = ser.AfterEvent
			bMatch.beforeEvent = ser.BeforeEvent
			break
		}
		match1, _ := regexp.MatchString(r.URL.Path, ser.Location)
		if match1 {
			lMatch.path = ser.Path
			lMatch.beforeEvent = ser.BeforeEvent
			lMatch.afterEvent = ser.AfterEvent
		}
	}
	bestMatch := bMatch
	if reflect.DeepEqual(bMatch, match{}) {
		bestMatch = lMatch
	}
	if len(bestMatch.path) <= 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	//随机负载均衡
	rand.Seed(time.Now().Unix())
	remote, err := url.Parse(bestMatch.path[rand.Intn(len(bestMatch.path))])

	if err != nil {
		panic(err)
	}
	r.URL.Scheme = remote.Scheme
	r.URL.Host = remote.Host
	d := SchemeMap[r.URL.Scheme]
	//处理请求之前的事件
	config.LUA.L.SetGlobal("m", lua.LString("1"))
	//这里就可以加载mymodule模块
	req = r
	res = w
	config.LUA.L.RegisterModule("dr", exports)
	config.LUA.HandlerLua(bestMatch.beforeEvent, server.ScriptPath)
	d.Handler(r, w, func(request *http.Request, writer http.ResponseWriter) {
		writer.Header().Add("goof", "111")
		config.LUA.HandlerLua(bestMatch.afterEvent, server.ScriptPath)
	})

}

// 导出对象
var exports = map[string]lua.LGFunction{
	"get_req_header": get_req_header,
	"add_req_header": add_req_header,
	"del_req_header": del_req_header,
	"get_res_header": get_res_header,
	"add_res_header": add_res_header,
	"del_res_header": del_res_header,
	"redirect":       redirect,
}

func get_req_header(L *lua.LState) int {
	key := L.ToString(1)
	val := req.Header.Get(key)
	L.Push(lua.LString(val))
	return 1
}
func add_req_header(L *lua.LState) int {
	key := L.ToString(1)
	val := L.ToString(2)
	req.Header.Add(key, val)
	return 0
}
func del_req_header(L *lua.LState) int {
	key := L.ToString(1)
	req.Header.Del(key)
	return 0
}
func get_res_header(L *lua.LState) int {
	key := L.ToString(1)
	val := res.Header().Get(key)
	L.Push(lua.LString(val))
	return 1
}
func add_res_header(L *lua.LState) int {
	key := L.ToString(1)
	val := L.ToString(2)
	res.Header().Add(key, val)
	return 0
}
func del_res_header(L *lua.LState) int {
	key := L.ToString(1)
	res.Header().Del(key)
	return 0
}
func redirect(L *lua.LState) int {
	path := L.ToString(1)
	http.Redirect(res, req, path, http.StatusTemporaryRedirect)
	return 0
}
