package arg

import (
	"fmt"
	"github.com/spf13/pflag"
	"os"
)

const INT_MAX = (int)(^uint(0) >> 1)

type Selpg_args struct {
	Start_page, End_page, Page_len int
	Page_type                      bool //false:l  true:f
	In_filename, Print_dest        string
}

func Usage() {
	fmt.Println("USAGE: selpg -s start_page -e end_page [ -f | -l lines_per_page ] [ -ddest ] [ in_filename ]")
}

func Bind(sa *Selpg_args) {
	pflag.IntVarP(&(*sa).Start_page, "startPage", "s", -1, "input start page")
	pflag.IntVarP(&(*sa).End_page, "endPage", "e", -1, "input end page")
	pflag.IntVarP(&(*sa).Page_len, "pageLen", "l", -1, "input page length")
	pflag.BoolVarP(&(*sa).Page_type, "pageType", "f", false, "input page type")
	pflag.StringVarP(&(*sa).Print_dest, "printDestination", "d", "", "input print Destination")
}

func Parse(sa *Selpg_args) {
	pflag.Parse()
	if pflag.NArg() > 1 {
		fmt.Fprint(os.Stderr,"selpg: unknown option")
		Usage()
		os.Exit(1)
	} else if pflag.NArg() == 1 {
		(*sa).In_filename = pflag.Arg(0)
	}
	if !(*sa).Page_type && (*sa).Page_len == -1 {
		(*sa).Page_len = 72
	}
}

func Process_args(args *Selpg_args) {
	//check the command-line arguments for validity
	if (*args).Start_page == -1 || (*args).End_page == -1 {
		fmt.Fprint(os.Stderr,"selpg: not enough arguments")
		Usage()
		os.Exit(3)
	}
	//handle 1st arg - start page
	if ((*args).Start_page < 1) || ((*args).Start_page > (INT_MAX - 1)) {
		fmt.Fprint(os.Stderr,"selpg: invalid start page ", ((*args).Start_page))
		Usage()
		os.Exit(4)
	}
	//handle 2nd arg - end page
	if ((*args).End_page < 1) || ((*args).End_page > (INT_MAX - 1)) || ((*args).End_page < (*args).Start_page) {
		fmt.Fprint(os.Stderr,"selpg: invalid end page ", ((*args).End_page))
		Usage()
		os.Exit(5)
	}
	// while there more args and they start with a '-'
	if ((*args).Page_type) && ((*args).Page_len != -1) {
		fmt.Fprint(os.Stderr,"selpg: option should be \"-f\" or -l pageLength")
		Usage()
		os.Exit(6)
	} else if (!(*args).Page_type) && (((*args).Page_len < 1) || ((*args).Page_len > (INT_MAX - 1))) {
		fmt.Fprint(os.Stderr,"selpg: invalid page length ", ((*args).Page_len))
		Usage()
		os.Exit(7)
	}
}
