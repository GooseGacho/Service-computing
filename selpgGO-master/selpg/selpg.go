package main

import (
	"github.com/Eayne/selpgGO/arg"
	"github.com/Eayne/selpgGO/read"
)

func main() {
	var sa arg.Selpg_args
	arg.Bind(&sa)
	arg.Parse(&sa)
	arg.Process_args(&sa)
	read.Process_input(&sa)
}
