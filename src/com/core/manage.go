package core

import (
	"../config"
	"net/http"
	"net/url"
)

var SchemeMap map[string]Dispath /*创建集合 */

func init() {
	var dispath Dispath
	SchemeMap = make(map[string]Dispath)
	dispath = Proxy{}
	SchemeMap["http"] = dispath
}
func Handdler(w http.ResponseWriter, r *http.Request) {
	servers := config.GLOBAL_CONFIG.Servers[0].Proxys[0].Path
	remote, err := url.Parse(servers[0])
	if err != nil {
		panic(err)
	}
	r.URL.Scheme = remote.Scheme
	r.URL.Host = remote.Host

	d := SchemeMap[remote.Scheme]
	//处理之前的事件
	server := config.GLOBAL_CONFIG.Servers[0]
	config.HandlerLua(server.InitEvent, server.ScriptPath)
	d.Handler(r, w)

}
