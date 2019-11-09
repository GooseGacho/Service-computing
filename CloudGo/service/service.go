package service
import (
    "github.com/codegangsta/martini" 
)
func NewServer(port string) {   
    m := martini.Classic()
    //提交请求的处理
    m.Get("/", func(params martini.Params) string {
        return "Welcome to use this service!"
    })

    m.RunOnAddr(":"+port)   
}