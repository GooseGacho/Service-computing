package read

import (
	"bufio"
	"fmt"
	"github.com/Eayne/selpgGO/arg"
	"io"
	"os"
	"os/exec"
)

func Process_input(args *arg.Selpg_args) {
	var input_scanner *bufio.Reader //io.Reader其实也可以
	var output_writer io.Writer     //Fprint中，第一个参数是io.Writer接口 (大部分读写都是实现这个接口)，且不用指针，因为接口类型不是指针
	var pr *io.PipeReader
	var pw *io.PipeWriter
	var cmd *exec.Cmd
	var file *os.File

	if (*args).In_filename != "" {
		file, err := os.Open((*args).In_filename) // 其中一个声明一个没有声明
		if err != nil {
			fmt.Fprint(os.Stderr,"selpg: input file error: ", err.Error())
			os.Exit(2)
		}
		input_scanner = bufio.NewReader(file)
	} else {
		input_scanner = bufio.NewReader(os.Stdin)
	}

	if (*args).Print_dest != "" {
		arg := fmt.Sprintf("-d %s", (*args).Print_dest)
		cmd = exec.Command("lp", arg, (*args).In_filename)
		//cmd = exec.Command("cat")
		pr, pw = io.Pipe()
		cmd.Stdin = pr
		//cmd.Stdout = os.Stdout
		output_writer = pw
		if output_writer == nil {
			fmt.Fprint(os.Stderr,"selpg: could not open pipe to ", arg)
			os.Exit(8)
		}
	} else {
		output_writer = bufio.NewWriter(os.Stdout) //指针类型分清楚
	}

	wordByte := make([]byte, 0)
	line_ctr, page_ctr := 1, 1
	for {
		c, err := input_scanner.ReadByte()
		if err == nil {
			if page_ctr >= args.Start_page && page_ctr <= args.End_page {
				wordByte = append(wordByte, c)
			}
			if !(*args).Page_type {
				if c == '\n' {
					line_ctr++
				}
				if line_ctr > args.Page_len {
					line_ctr = 1
					page_ctr++
				}
			} else {
				if c == '\f' {
					page_ctr++
				}
			}
		} else {
			break
		}
	}

	_, ok := output_writer.(*bufio.Writer)
	s := string(wordByte)
	if ok {
		fmt.Fprintf(output_writer, "%s", s)
		(output_writer.(*bufio.Writer)).Flush() // 必须要flush，否则不能输出出来。已经类型转换要注意
	} else {
		go func() {
			fmt.Fprintf(output_writer, "%s", s)
			output_writer.(*io.PipeWriter).Close()
		}()
	}

	if page_ctr < args.Start_page {
		fmt.Fprint(os.Stderr,"selpg: start_page (", args.Start_page, ") greater than total pages (", page_ctr, "),", " no output written")
	} else if page_ctr < args.End_page {
		fmt.Fprint(os.Stderr,"selpg: end_page (", args.End_page, ") greater than total pages (", page_ctr, "),", " less output than expected")
	}
	if cmd != nil {
		err := cmd.Run()
		if err != nil {
			fmt.Fprint(os.Stderr,"selpg: lp error")
			os.Exit(10)
		}
	}
	file.Close()
	if pr != nil {
		pr.Close()
	} //Close()？
}
