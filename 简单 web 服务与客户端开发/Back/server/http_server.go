package server

import (
	"fmt"
	"net/http"
	"github.com/cjf/smog/middle"
)

func Run(port int) {
	http.HandleFunc("/usersv/getUserInfo", wrapHttpHandlerFunc(middle.GetUserHttpHandler()))
	http.HandleFunc("/blogsv/getBlogListByUID", wrapHttpHandlerFunc(middle.GetBlogsByUidHttpHandler()))
	http.HandleFunc("/blogsv/getBlogByBID", wrapHttpHandlerFunc(middle.GetBlogByIdHttpHandler()))
	http.HandleFunc("/blogsv/getBlogListByTagAndUID", wrapHttpHandlerFunc(middle.GetBlogsByTIDAndUIDHttpHandler()))
	http.HandleFunc("/commentsv/getCommentListByBID", wrapHttpHandlerFunc(middle.GetCommentListByBIDHttpHandler()))
	http.HandleFunc("/tagsv/getTagsByUID", wrapHttpHandlerFunc(middle.GetTagsByUIDHttpHandler()))
	fmt.Printf("Listening ... %v\n", port)
    http.ListenAndServe(":8989", nil)
}

func wrapHttpHandlerFunc(handle func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request)  {
	return func (w http.ResponseWriter, r *http.Request) {
		http_preprocess(w, r)
		handle(w, r)
	}
}

//solve CORS
func http_preprocess(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Add("Access-Control-Allow-Headers",  "content-type")    
    w.Header().Set("Access-Control-Max-Age", "86400")
    w.Header().Set("Access-Control-Allow-Methods", "POST")
    w.Header().Set("Access-Control-Allow-Credentials", "true");
}
