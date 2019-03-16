# 要解决的问题

	如何解决外部服务异常场景测试困难的问题。
	详细描述：外部服务出现非预期错误码，大量改动服务代码来兼容异常，但是线下又难以复现，如何在验证代码改动的有效性。

# 解决方案
	mock server，实现第三方服务的各类接口，并支持自定义异常错误码、状态接口。
	

## 达成目标
	以server形式存在，具有mock外部服务名、协议、状态码、返回值等默认配置，可以随意构造异常错误，以提供回归测试。

## 解决方案构成

| 构成 | 解释 |
| --- | --- |
| Server | 服务端 |
| Conf | 配置文件 |
| UI |mockserver前端|

## 解决方案示例
####配置示例
<<< @/../mock-server/conf.json

```
{
    "/testUrl":{
        "protocol":"http",
        "return_http_code":200,
        "return_content_type":"application/json",
        "return_body":{
            "err_no":0,
            "err_msg":"success"
        }
    }
}
```
### response.go
<<< @/../mock-server/response.go

### main.go
<<< @/../mock-server/main.go

运行后，server每个1秒刷一次配置，可随时手动修改，到达mock效果



