# DyRouter
DyRouter 一个基于go的动态路由服务

## 开发背景
    微服务盛行、但缺乏完善支持微服务的边缘框架、DyRouter 因此应运而生。同时他还是一款高性能，低内存占用的代理服务。
## 整体架构
    开发语言go 
    数据库：内存数据库
    支持通过接口维护路由关系，方便业务集成，同时接口支持安全验证，从底层就保证了系统安全。
    支持插件扩展，方便原生的和各种服务对接插件。

## 功能
 核心：通过接口 动态维护路由 查看路由
 
 边缘功能：支持界面集成、提供其他服务发现插件
 亮点：默认支持头路由，方便根据头进行环境切换、路由支持单条规则单独更新，提高服务更新效率，且无需重启。
 
## 路由支持 
    正则 、路径 、域名、根据头路由
## 协议支持
    http/websokect
    
## 接口：

```
{
	addRule,
	modRule,
	delRule,
	getRuleById,
	pageRule,
	启用规则,
	禁用规则,
	刷新规则
	TODO
	...
}
```


## 路由表信息新版
```
{
	 "//服务器编码":"",
	 "server1":{
		"hostName":"baidu.com",
		"port":8185,
		"dyRouters":{
			"//服务器编码":"",
			"code":"/api/bp",
			"//path 路径":"",
			"path":"/api/bp",
			"proxy_pass":[
				{
					"//type 协议":"",
					"type":"http,websokect", 
					"proxy":"http://172.16.1.84",
					"//公共扩展参数 系统内置功能需要的参数 根据情况扩展":"",
					"params":{
						"weight":1,
						"allowIp非必填":"",
						"allowIp":"127.0.0.1,127.0.0.2",
						"forbidIp非必填":"",
						"forbidIp":"192.168.0.4"
					},
					"//初始化":"",
					"eventInit":[
						{
							"//语言支持 lua js":"",
							"lang":"lua",
							"script":"内容",
							"scriptFile":"脚本文件路径",
							"params":"@see proxy_pass/params"
						}
					],
					"//请求之前":"",
					"eventReqBefore":[
						{
							"所有事件格式一致":""
						}
					],
					"//请求之后":"",
					"eventReqAfter":[
						{
							"所有事件格式一致":""
						}
					]
					
				}
			]
		}
	},
	"serverN":{
		"//说明: @see server1 多个server":""
	}
	
	
}
```
## 路由表信息
```
[
	serverN{
		code:
		hostName:baidu.com
		{
			code:便于程序 通过接口维护
			path 路径
			proxy_pass:[
				{
					proxy:127.0.0.1
					weight:1
					allowIp:127.0.0.1,127.0.0.2 非必填
					forbidIp:192.168.0.4 非必填
					routeHeader:[
						{
							env:test  比如 路由到测试环境
							match:绝对，包含，正则
						}
					]
				},
				{
					proxy:127.0.0.1
					weight:1
				}
			]
		}
	}
]
```
## 插件管理：
    TODO
## 计划提供的插件：
    Eureka服务集成（TODO）
	nacos服务集成（TODO）
	etcd服务集成（TODO）
	web管理（TODO）

## 计划表
    1、基础环境搭建（进行中）
    2、支持基本路由功能（计划中）
    3、支持路由各种配置（正则）（计划中）
    4、插件管理（计划中）
    5、插件开发（计划中）
 ### 开发邀请
 由于精力有限，欢迎有兴趣的伙伴加入
 联系方式：qq:419861036
 
