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