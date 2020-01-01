package service 

import (
	"github.com/cjf/smog/gen-go/rpc"
)

func GetComment(bid int32) (r []*rpc.Comment, err error)  {
	//test
	return []*rpc.Comment {
		&rpc.Comment{Bid: 1, Uname: "che12", Content: "简洁明了", Date: "2019-12-30", Zancnt: 1},	
		&rpc.Comment{Bid: 1, Uname: "毛毛三", Content: "最后的图挂了鸭", Date: "2019-12-30", Zancnt: 0},
	}, nil
}