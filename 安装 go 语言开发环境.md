## 安装 go 语言开发环境

# 安装vscode
首先在我使用的centOS中，已经在yum存储库中发布稳定的64位VS代码，以下脚本将安装密​​钥和存储库

```
sudo rpm --import https://packages.microsoft.com/keys/microsoft.asc
sudo sh -c ‘echo -e “[code]\nname=Visual Studio 
 Code\nbaseurl=https://packages.microsoft.com/yumrepos/vscode\nenabled=1\ngpgcheck=1\ngp
 gkey=https://packages.microsoft.com/keys/microsoft.asc” > /etc/yum.repos.d/vscode.repo’
```
之后使用yum来更新包缓存并安装包：

```
yum check-update
sudo yum install code
```
这样就安装好了vscode，使用时可以再菜单启动，也可以在命令行输入code
## 安装 golang
在安装golang的过程中，如果centOS配置的是google源并可以连接到服务器直接使用

```
$ sudo yum install golang
```
安装即可，若已经将源换为国内源或是不能连接到google服务器，则可以在https://golang.org/dl/官网下载最新的安装包
![在这里插入图片描述](https://img-blog.csdnimg.cn/20190914154141154.png)
下载后进入下载目录，使用命令行来解压并安装

```
sudo tar -C /usr/local -xzf go1.13.linux-amd64.tar.gz
```
可以使用

```
$ go version
```
来测试安装是否成功。
安装完成之后进行环境变量的配置，在/etc/profile中添加如下语句：

```
export PATH=$PATH:/usr/local/go/bin
```
之后重新登录当前用户来应用这些配置。

还需要为工作区设置一下环境变量，首先创建工作空间

```
$ mkdir $HOME/gowork
```
然后配置环境变量

```
export GOPATH=HOME/goworkexportPATH=HOME/goworkexport PATH=HOME/goworkexportPATH=PATH:$GOPATH/bin
```
然后执行这个语句来执行配置

```
$ source $HOME/.profile
```
## 安装必要的工具和插件
创建一个.go文件，然后用vscode打开，这里直接点击出现的提示中的install all即可。
若连接不到golang.org则需先下载源码：

```
mkdir $GOPATH/src/golang.org/x/
go get -d github.com/golang/tools
cp $GOPATH/src/github.com/golang/tools $GOPATH/src/golang.org/x/ -rf
```
然后安装工具包：

```
$ go install golang.org/x/tools/go/buildutil
```
退出 vscode，再进入，按提示安装。
## 开始写第一个包并进行测试
根据官方文档的指示，建立工作空间，具体过程和上述安装golang时类似，并为新建的工作空间配置GOPATH环境变量。
设置工作期与环境变量

```
$ mkdir $HOME/work
$ export GOPATH=$HOME/work
$ export PATH=PATH:PATH:PATH:GOPATH/bin
```
为包添加路径

```
$ mkdir -p $GOPATH/src/github.com/user
```
创建第一个程序

 - 程序路径
 - `$ mkdir $GOPATH/src/github.com/user/hello`
创建完成之后就可以创建并编写第一个程序了

```
package main
import “fmt”
func main() {
fmt.Printf(“Hello, world.\n”)
}
```
用 go 工具构建并安装此程序

```
$ go install github.com/user/hello
```
之后运行此程序

```
$ $GOPATH/bin/hello
Hello, world.
```
下面是创建第一个库

 - 为包添加路径
 - `$ mkdir $GOPATH/src/github.com/user/stringutil`
接着，在该目录中创建名为 reverse.go 的文件，内容如下
![在这里插入图片描述](https://img-blog.csdnimg.cn/20190914154656809.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2FkZ2hqZ2Y=,size_16,color_FFFFFF,t_70)
现在用 go build 命令来测试该包的编译

```
$ go build github.com/user/stringutil

```
之后将刚才的hello.go文件修改为
![在这里插入图片描述](https://img-blog.csdnimg.cn/20190914154734559.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2FkZ2hqZ2Y=,size_16,color_FFFFFF,t_70)
最后通过

```
$ go install github.com/user/hello
```
来安装 hello 程序时，stringutil 包也会被自动安装。
此时工作空间是这样的：
![在这里插入图片描述](https://img-blog.csdnimg.cn/20190914154804426.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2FkZ2hqZ2Y=,size_16,color_FFFFFF,t_70)
之后可以创建一个测试文件

```
$GOPATH/src/github.com/user/stringutil/reverse_test.go
```
![在这里插入图片描述](https://img-blog.csdnimg.cn/2019091415483876.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2FkZ2hqZ2Y=,size_16,color_FFFFFF,t_70)
接着测试结果为：
![在这里插入图片描述](https://img-blog.csdnimg.cn/20190914154850722.png)
最后要学习的就是安装远程包的步骤：

```
$ go get github.com/golang/example/hello
$ $GOPATH/bin/hello
Hello, Go examples!
```
完成之后可以看到新出现的golang部分。
![在这里插入图片描述](https://img-blog.csdnimg.cn/20190914154924607.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2FkZ2hqZ2Y=,size_16,color_FFFFFF,t_70)
