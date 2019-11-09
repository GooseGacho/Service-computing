一、概述
开发简单 web 服务程序 cloudgo，了解 web 服务器工作原理。

**任务目标**

 - 熟悉 go 服务器工作原理
 - 基于现有 web 库，编写一个简单 web 应用类似 cloudgo。
 - 使用 curl 工具访问 web 程序
 - 对 web 执行压力测试

二、任务要求

 - 编程 web 服务程序 类似 cloudgo 应用。
 - 要求有详细的注释
 - 是否使用框架、选哪个框架自己决定。请在 README.md 说明你决策的依据
 - 使用 curl 测试，将测试结果写入 README.md
 - 使用 ab 测试，将测试结果写入 README.md。并解释重要参数。


【框架选择】
我使用的web开发框架是Martini。Martini 是一个非常新的 Go 语言的 Web 框架，使用 Go 的 net/http 接口开发，类似 Sinatra 或者 Flask 之类的框架，也可使用自己的 DB 层、会话管理和模板。

选择的理由

 1. 使用非常简单 
 2. 无侵入设计 
 3. 可与其他 Go 的包配合工作 
 4. 超棒的路径匹配和路由 
 5. 模块化设计，可轻松添加工具 
 6. 大量很好的处理器和中间件
 7. 很棒的开箱即用特性 
 8. 完全兼容 http.HandlerFunc 接口

**安装与测试**
1.安装
go get github.com/codegangsta/martini
2.测试
编写以下代码server.go：
```
package main

import "github.com/codegangsta/martini"

func main() {
  m := martini.Classic()　　//创建一个典型的martini实例
  m.Get("/", func() string {     //接收对\的GET方法请求，第二个参数是对一请求的处理方法
    return "Hello world!"
  })
  m.Run()  //运行服务器
}
```
2.测试
输入命令go run main.go运行代码后，打开网页http://localhost:3000，可以看到“Hello world!”显示在网页上。
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191106212724939.png)

【应用编写】
简单地写一个在屏幕上显示文字的应用。在cloudgo文件夹中创建main.go和service文件夹（内有service.go）。

main.go
使用了老师博客中给出的代码，完成绑定端口为8080、解析端口、启动server完成操作的操作。
```
package main

import (
    "os"
    "github.com/shanzhulizhi/cloudgo/service"
    flag "github.com/spf13/pflag"
)

//设置默认端口为8080
const (
    PORT string = "8080"
)

func main() {
    port := os.Getenv("PORT")
    if len(port) == 0 {
        port = PORT  //如果没有监听端口，则设为默认端口
    }
 
    //用户可以自己加上-p参数设置端口
    pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
    flag.Parse()
    if len(*pPort) != 0 {
        port = *pPort
    }
 
    //启动server
	service.NewServer(port)
}
```
**service.go**
使用martini框架中的函数格式具体定义main.go文件中启动server后要具体进行的操作：显示文字“Welcome to use web service!”。其中martini安装在codegangsta文件夹中。
```
package service
import (
   "github.com/codegangsta/martini" 
)
func NewServer(port string) {   
    m := martini.Classic()
    //提交请求的处理
    m.Get("/", func(params martini.Params) string {
        return "hello world"
    })

    m.RunOnAddr(":"+port)   
}
```
【运行测试】
命令行测试
输入命令go run main.go -p 8090，监听端口为8090。在l浏览器中输入http://localhost:8090，可以看到文本信息。
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191106232135332.png)
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191107001146953.png)

**curl测试**
运行main.go后，在另外一个终端输入curl -v http://localhost:8090，可以看到连接成功并显示文本信息。
![在这里插入图片描述](https://img-blog.csdnimg.cn/2019110700102856.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3h1YW5fdGluZw==,size_16,color_FFFFFF,t_70)

**ab测试**
1.安装Apache web压力测试程序：
yum -y install httpd-tools
2.运行main.go后，输入以下命令执行压力测试：
ab -n 1000 -c 100 http://localhost:8090/
可以看到以下信息：
![在这里插入图片描述](https://img-blog.csdnimg.cn/2019110700180892.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3h1YW5fdGluZw==,size_16,color_FFFFFF,t_70)

**参数解释**

 - ab的参数-n 1000表示执行的请求数量为1000个，-c 100表示并发请求的个数为100个。
 - 服务器的主机名是localhost，监听端口是8090.。 
 - Document Path: 请求的资源 
 - Document Length:文档返回的长度，此处是28bytes 
 - Concurrency Level: 并发个数 
 - Time taken for tests:总请求时间 
 - Complete requests: 总成功请求数 
 - Failed requests: 失败的请求数 
 - Write errors:错误数 
 - Total transferred: 传输总字节数 
 - HTML transferred: HTML传输字节数 
 - Requests per second: 平均每秒的请求数 
 - Time per request: 平均每个请求消耗的时间 
 - Time per request:平均请求消耗时间除以并发数 
 - Transfer rate: 传输速率 
 - Connection Times：说明了连接、处理、等待和总时间的最小值、最大值、中间值和均值。 
 - Percentage of the requests served within a certain time：说明了请求完成的百分比和所用时间。

**重点指标**
对压力测试的结果重点关注吞吐率（Requests per second）、用户平均请求等待时间（Time per request）指标：

吞吐率（Requests per second）：
服务器并发处理能力的量化描述，单位是reqs/s，指的是在某个并发用户数 下单位时间内处理的请求数。某个并发用户数下单位时间内能处理的最大请求数，称之为最大吞吐率。
记住：吞吐率是基于并发用户数的。这句话代表了两个含义：
a.吞吐率和并发用户数相关
b.不同的并发用户数下，吞吐率一般是不同的
计算公式：总请求数/处理完成这些请求数所花费的时间，即
Request per second=Complete requests/Time taken for tests
必须要说明的是，这个数值表示当前机器的整体性能，值越大越好。
用户平均请求等待时间（Time per request）：
计算公式：处理完成所有请求数所花费的时间/（总请求数/并发用户数），即：
Time per request=Time taken for tests/（Complete requests/Concurrency Level）
服务器平均请求等待时间（Time per request:across all concurrent requests）：
计算公式：处理完成所有请求数所花费的时间/总请求数，即：
Time taken for/testsComplete requests
可以看到，它是吞吐率的倒数。
同时，它也等于用户平均请求等待时间/并发用户数，即
Time per request/Concurrency Level。

