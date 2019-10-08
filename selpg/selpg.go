package main

import (

	"fmt"
	"flag"
	"os"
	"io"
	"bufio"
	"os/exec"

)

type selpg_args struct
{
	start_page int
	end_page int
	in_filename string
	page_len int
	page_type bool

	print_dest string
}

var progname string

func get(parg *selpg_args) {
	flag.IntVar(&parg.start_page, "s", -1, "Start page.")
	flag.IntVar(&parg.end_page, "e", -1, "End page.")
	flag.IntVar(&parg.page_len, "l", 72, "Line number per page.")
	flag.BoolVar(&parg.page_type, "f", false, "Determine form-feed-delimited")
	flag.StringVar(&parg.print_dest, "d", "", "specify the printer")
	flag.Parse()

	args_left := flag.Args()
	if(len(args_left) > 0){
		parg.in_filename = string(args_left[0])
	} else {
		parg.in_filename = ""
	}
}

func process_args(parg *selpg_args) {

	if parg == nil{
		fmt.Fprintf(os.Stderr, "\n[Error]The args is nil!Please check your program!\n\n")
		os.Exit(1)
	}else if(parg.start_page == -1) || (parg.end_page == -1){
		fmt.Fprintf(os.Stderr, "\n[Error]The startPage and endPage is not allowed empty!Please check your command!\n\n")
		os.Exit(2)
	}else if (parg.start_page < 0) || (parg.end_page < 0){
		fmt.Fprintf(os.Stderr, "\n[Error]The startPage and endPage is not negative!Please check your command!\n\n")
		os.Exit(3)
	}else if parg.start_page > parg.end_page{
		fmt.Fprintf(os.Stderr, "\n[Error]The startPage can not be bigger than the endPage!Please check your command!\n\n")
		os.Exit(4)
	}

}

func process_input(parg *selpg_args) {
	var fin *os.File
	var fout *os.File
	var fout_d io.WriteCloser

	if parg.in_filename == ""{
		fin = os.Stdin
	} else {
		var err_fin error
		fin, err_fin = os.Open(parg.in_filename)

		if err_fin != nil {
			fmt.Fprintf(os.Stderr, "\n[Error]inputFile:");
			panic(err_fin)
		}
	}

	if len(parg.print_dest) == 0 {
		fout = os.Stdout;
		fout_d = nil
	} else {
		fout = nil
		var err_dest error
		cmd := exec.Command("./" + parg.print_dest)
		fout_d, err_dest = cmd.StdinPipe()
		if err_dest != nil {
			fmt.Fprintf(os.Stderr, "\n[Error]fout_dest:");
			panic(err_dest)
		}

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err_start := cmd.Start()
		if err_start != nil {
			fmt.Fprintf(os.Stderr, "\n[Error]command-start:");
			panic(err_start)
		}
	}

	if(fout != nil){
		//output_to_file
		line_ctr := 0
		page_ctr := 1
		buf := bufio.NewReader(fin)

		for true {
			line,err := buf.ReadString('\n')
			if err == io.EOF{
				break
			}
			if err != nil {
				fmt.Fprintf(os.Stderr, "\n[Error]file_in_out:");
				panic(err)
			}

			line_ctr++
			if line_ctr > parg.page_len{
				page_ctr++
				line_ctr = 1
			}
			if (page_ctr >= parg.start_page) && (page_ctr <= parg.end_page){
				fmt.Fprintf(fout, "%s", line)
			}
		}
	} else {
		//output_to_exc
		line_ctr := 0
		page_ctr := 1
		buf := bufio.NewReader(fin)

		for true {
			bytes, err := buf.ReadByte()
			if err == io.EOF{
				break
			}
			if line_ctr > parg.page_len{
				page_ctr++
				line_ctr = 1
			}
			if (page_ctr >= parg.start_page) && (page_ctr <= parg.end_page){
				fout_d.Write([]byte{bytes})
			}
		}
	}

}

func main(){
	var args selpg_args
	get(&args)
	process_args(&args)
	process_input(&args)
}
