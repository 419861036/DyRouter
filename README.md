# DyRouter
DyRouter 一个基于go的动态路由服务

-- 开发背景
通过接口 动态维护路由 查看路由 支持界面集成
接口：{
	addRule,
	modRule,
	delRule,
	getRuleById,
	pageRule,
	启用规则,
	禁用规则,
	刷新规则
}
路由支持 正则 、路径 、域名、根据头路由
协议支持：http/websokect
路由表信息
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
插件：
	nacos服务集成
	etcd服务集成
	web管理
