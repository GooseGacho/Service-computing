package service 

import (
	"github.com/cjf/smog/gen-go/rpc"
)

func GetTags(uid int32) (r []*rpc.Tag, err error) {
	//test
	return []*rpc.Tag {
		&rpc.Tag{Tid: 1, UID: 1, Tname: "Dijkstra"},
		&rpc.Tag{Tid: 1, UID: 1, Tname: "最短路径"},
	}, nil
}