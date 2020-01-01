﻿**本次项目我完成了一个极简博客的客户端和服务端，主要工作有：设计（API、架构、数据库）、前后端代码开发、测试、文档完成，以下对工作的总体流程和一些想法做一下简要review。 （实验环境为centos7）**

## **总体设计**

 - 由于之前RESTful服务使用得比较多，这次想尝试一下别的，所以在服务器的web API的选择上，采用RPC风格。
   

总的来说，可以把RESTful和RPC理解为不同的交流协议，前者通过资源符和状态去与服务端交互，核心是资源，后者则通过方法地址（名）和参数，核心是方法，当然这只是表面的理解，它们的内在区别还有很多，各有优劣。基于RPC的API更加适用行为(也就是命令和过程)，基于REST的API更加适用于构建模型(也就是资源和实体)，处理CRUD。
 - RPC框架我选择的是Apache thrift，另外可供选择的还有JSON-RPC。
   由于本次作业为web开发，相当于只能把http当成一个传输层协议（而非应用层协议），不能很好的利用http的很多feature。但是无妨，重点是学习RPC和thrift。注：thrift常用于后端（包含中间件）之间而非后端直接与前端，从而能很好地利用thrift基于tcp/udp、利用二进制作为消息传递方式的特点。
   前端使用vue，后端使用Go，使用go提供的http服务作为基本框架。

## Swagger API

我使用Swagger Editor来编写Swagger API文档，官网本地安装即可。 设计时注意的：

 1. 由于是以RPC为基础，所以在设计API时局限比较大，RPC只能以GET或POST的方式传输，幸好本次作业也基本是请求数据，这个限制也没什么区别。
 2. 同样不同于RESTful，路径往往是一个动作，而非资源。 

## Thrift

下载之后： 首先需要一点上网技巧，然后还要装好boost 1.53.0，c++11等等很多需要的库（按官网），然后：
```
$ tar -xf thrift-0.13.0.tar.gz && cd cd thrift-0.13.0/
$ ./configure --prefix=/usr/local/ --with-boost=/usr/local --without-python --without-ruby --without-java --without-haskell --without-erlang --without-perl --without-php --without-cpp --without-json --without-as3 --without-csharp --without-erl --without-cocoa --without-ocaml --without-hs --without-xsd --without-html --without-delphi --without-gv --without-lua --without-qt
$ make
$ make install
```
其中configure的参数自己选定，不需要什么语言最好without修饰去掉。接着使用idl编写的接口定义文件： 类似于：
```c
struct User {
  1: required i32 id,
  2: required string name,
  10: required i32 creacnt, 
  11: required i32 fancnt,
  12: required i32 zancnt,
  13: required i32 commentcnt,
  14: required i32 visitcnt
}

service UserService {
  User GetUserInfo(1: Stringreq s)
}
```
使用thrift生成代码：thrift -r --gen js xxx.thrift && thrift -r --gen go xxx.thrift 接着就可以使用生成的代码了，注意还需要导入thrift文件夹下/lib/go或/lib/js的thrift文件。

## 客户端

客户端使用的是vue。 简述：

 1. 安装node.js和npm（代理需要改一下npm的config文件，否则使用cnpm） 
 2. 安装vue和vue-cli
 3. 使用webpack初始化一个vue项目 
 4. 编写页面并且引入生成代码和Thrift库。网上js通过http使用thrift的资料比较少，具体用法可以看项目源码。参考官网：http://thrift.apache.org/tutorial/js 注意：js只支持AJAX通信，且只能使用JsonProtocol。 

## 服务端

使用go及其基本包。 简述：

 1. 创建HTTP服务器，设定路径与处理函数 
 2. 编写处理函数，也就是将HTTP请求使用thrift处理
 3. 编写thrift处理函数（使用到了flag和thrift的NewThriftHandlerFunc函数）编写业务函数（数据处理、数据库等）

注意：

 1. js只能通过AJAX通信，也就是只可以通过HTTP的方式，所以我们需要在服务器增加HTTP解析这一步骤。具体所用API参考GoDoc：https://godoc.org/github.com/apache/thrift/lib/go/thrift。
 2. TServer在thrift框架中的主要任务是接收client请求，并转发到某个processor上进行请求处理。针对不同的访问规模，thrift提供了不同TServer模型。thrift目前支持的server模型包括：
```c
     - TSimpleServer: 单线程服务器端使用标准的阻塞式I/O
    - TTHreaadPoolServer:   多线程服务器端使用标准的阻塞式I/O
    - TNonblockingServer:多线程服务器端使用非阻塞式I/O   
    - TThreadedServer:多线程网络模型，使用阻塞式I/O,为每个请求创建一个线程 对于Go，只有TSimpleServer模型。
```
 3. 注意跨域问题

## 数据库

数据库为boltDB，它是一个非关系型数据库，数据是以简单的值键对形式来存储的，相当于一个map。 大致使用流程：

 1. open函数。打开/创建数据库 
 2. CreateBucket函数。创建表。 
 3. Update函数。增/删/改数据。 
 4. View函数。查找数据。

## 遇到的困难

其实在安装的时候遇到了很多问题:依赖、很多环境等，但是这里还是主要讨论开发的问题吧：

 1. vue-cli的代理问题。vue-cli代理是看系统配置，而非npm的配置。并且在profile文件中需要大写的HTTPS-PROXY并且需要http://前缀！
 2. thrift二进制为何可以使用HTTP的疑问:https://stackoverflow.com/questions/38088324/thrift-can-use-http-but-it-is-a-binary-communication-protocol
 3. thrift要用json protocol和HTTPprocessor。所以要使用NewTJSONProtocolFactory，而且需要使用http作为服务器不可以使用thrift的，在这里卡了很久，网上没有相关资料，只能自己在GoDoc中看API，最后使用NewThriftHandlerFunc函数得以解决。参考:https://stackoverflow.com/questions/16350995/thrift-transport-in-javascript-client
 4. 按照js教程在客户端与服务端之间一直连不通，经过以下步骤 
```
	首先是查阅发现js只支持http，进行了上面第3步的改进
    怀疑vue中没能成功导入js包，于是尝试多种方法，甚至直接将js代码插入，但是虽然可以检测到js的对象，但是会一直报语法错误
    怀疑是vue的虚拟dom问题，于是换成传统js实验，解决了连接问题，能检测到，但是报network error
    通过在服务端thrift文件以及生成文件中打断点发现post请求体为空，报eof错误。
    修改前端thrift生成的代码，但是未发现错误，只能将firefox换为chrome（切换为沙盒模式），在这里不得不说chrome的错误信息详细得多得多。
    ```
    发现是跨域问题（CORS），参考https://www.ruanyifeng.com/blog/2016/04/cors.html，由于为复杂请求，所以需要处理预检请求：
    ```c w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Add("Access-Control-Allow-Headers",  "content-type")    
    w.Header().Set("Access-Control-Max-Age", "86400")
    w.Header().Set("Access-Control-Allow-Methods", "POST")
    w.Header().Set("Access-Control-Allow-Credentials", "true");
    if(r.Method == "OPTIONS") {  return } ```
```
## 小结

本次项目业务要求不高，老师只要求了资源（用户、博客等）的获取，只有CRUD中的R，所以业务还是比较简单的，很明显老师是希望让我们去学习API设计规范、Go构建web服务、RESTful/RPC实践、使用web客户端调用远端服务，这些更加有用的知识，而不是业务本身。 但是在thrift的安装和使用过程中还是遇到了挺多困难的，一度想要放弃去使用熟悉的Restful，最终能实现自己最初的想法挺不容易。 通过这次作业，我学到了很多东西，有RPC，有VUE，有thrift等等，从设计到测试，从前端一个小控件，到后端数据库的CRUD，这些过程对我来说都是崭新的，十分有益的。Go语言方面我深入学习了module等方面的知识，在服务计算这门课的核心技能上也有了很大的长进，感谢老师和TA。
