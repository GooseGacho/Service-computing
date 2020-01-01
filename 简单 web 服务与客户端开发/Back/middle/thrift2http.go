package middle 

import (	
	"context"
	"flag"
	"fmt"
    "os"
    "net/http"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/cjf/smog/gen-go/rpc"
	"github.com/cjf/smog/service"
)

type UserService struct {
}

type BlogService struct {
}

type TagService struct {
}

type CommentService struct {
}

//thrift service func implement 
func (this *UserService) GetUserInfo(ctx context.Context, s *rpc.Stringreq) (r *rpc.User, err error) {
    fmt.Println("OHH!!")
	return service.GetUser(s.S)
}

func (this *TagService) GetTagsByUID(ctx context.Context, i *rpc.Intreq) (r []*rpc.Tag, err error) {
	return service.GetTags(i.I)
}

func (this *BlogService) GetBlogsByTIDAndUID(ctx context.Context, i *rpc.Int2req) (r []*rpc.Blog, err error) {
	return service.GetBlogsByTUID(i.I1, i.I2)
}
func (this *BlogService) GetBlogsByUid(ctx context.Context, i *rpc.Intreq) (r []*rpc.Blog, err error) {
	return service.GetBlogsByUID(i.I)
}
func (this *BlogService) GetBlogById(ctx context.Context, i *rpc.Intreq) (r *rpc.Blog, err error) {
	return service.GetBlog(i.I)
}

func (this *CommentService) GetCommentListByBID(ctx context.Context, i *rpc.Intreq) (r []*rpc.Comment, err error) {
	return service.GetComment(i.I)
}

// Handler transform
func GetTagsByUIDHttpHandler() func(w http.ResponseWriter, r *http.Request) {
	protocolFactory := thrift.NewTJSONProtocolFactory()
	handler := &TagService{}
	processor := rpc.NewTagServiceProcessor(handler)
	return thrift.NewThriftHandlerFunc(processor, protocolFactory, protocolFactory)
}

func GetBlogsByTIDAndUIDHttpHandler() func(w http.ResponseWriter, r *http.Request) {
	return GetBlogsByUidHttpHandler()
}

func GetBlogsByUidHttpHandler() func(w http.ResponseWriter, r *http.Request) {
	protocolFactory := thrift.NewTJSONProtocolFactory()
	handler := &BlogService{}
	processor := rpc.NewBlogServiceProcessor(handler)
	return thrift.NewThriftHandlerFunc(processor, protocolFactory, protocolFactory)
}

func GetBlogByIdHttpHandler() func(w http.ResponseWriter, r *http.Request) {
	return GetBlogsByUidHttpHandler() 
}

func GetCommentListByBIDHttpHandler() func(w http.ResponseWriter, r *http.Request) {
	protocolFactory := thrift.NewTJSONProtocolFactory()
	handler := &CommentService{}
	processor := rpc.NewCommentServiceProcessor(handler)
	return thrift.NewThriftHandlerFunc(processor, protocolFactory, protocolFactory)
}


func GetUserHttpHandler() func(w http.ResponseWriter, r *http.Request) {
	//命令行参数
    protocol := flag.String("P", "json", "Specify the protocol (binary, compact, json, simplejson)")
    //framed := flag.Bool("framed", false, "Use framed transport")
    //buffered := flag.Bool("buffered", false, "Use buffered transport")
    //addr := flag.String("addr", "localhost:9090", "Address to listen to")

    flag.Parse()

    //protocol
    var protocolFactory thrift.TProtocolFactory
    switch *protocol {
    case "compact":
        protocolFactory = thrift.NewTCompactProtocolFactory()
    case "simplejson":
        protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
    case "json":
        protocolFactory = thrift.NewTJSONProtocolFactory()
    case "binary", "":
        protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
    default:
        fmt.Fprint(os.Stderr, "Invalid protocol specified", protocol, "\n")
        os.Exit(1)
    }
    //handler
	handler := &UserService{}
	//process
	processor := rpc.NewUserServiceProcessor(handler)
	return thrift.NewThriftHandlerFunc(processor, protocolFactory, protocolFactory)


    //buffered
    // var transportFactory thrift.TTransportFactory
    // if *buffered {
    //     transportFactory = thrift.NewTBufferedTransportFactory(8192)
    // } else {
    //     transportFactory = thrift.NewTTransportFactory()
    // }

    // //framed
    // if *framed {
    //     transportFactory = thrift.NewTFramedTransportFactory(transportFactory)
    // }

    // //transport,no secure
    // var err error
    // var transport thrift.TServerTransport
    // transport, err = thrift.NewTServerSocket(*addr)
    // if err != nil {
    //     fmt.Println("error running server:", err)
    // }

    //processor

    //fmt.Println("Starting the simple server... on ", *addr)
    
    //start tcp server
    // server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
    // err = server.Serve()

    //if err != nil {
    //    fmt.Println("error running server:", err)
    //}
}
