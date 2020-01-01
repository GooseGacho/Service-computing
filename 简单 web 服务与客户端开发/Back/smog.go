package main

import (
	"fmt"
	"github.com/cjf/smog/server"
	"github.com/cjf/smog/dao"
)


func main() {
	fmt.Println("Smog start!")
	dao.InitDB()
    server.Run(8989)
}

