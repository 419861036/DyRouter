

--多参数0返回值
function GetBigger(a,b)
    if a >= b then
        print (a)
    else
        print (b)
    end
end

--0参数1返回值
function GetServer()
    print ("Server")
end
     local http = require("http")
function GetApplication()
    response, error = http.request("get", "http://www.baidu.com")
    print(response['body'])
    print("Application")
end


--多参数1返回值
function Compare(a,b)
    if a >= b then
        return a
    else
        return b
    end
end

--1参数多返回值
function MoreReturn(a)
    if (a == 10) then
        return "world","hello","golang"
    end
end
--0返回值0参数
function GetStr()
       dr.close(502)
    --dr.redirect("http://www.baidu.com")
    print(dr)
end
--0返回值0参数
function GetStr1()
    dr.add_res_header("bbb","111")
end