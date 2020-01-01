
GetUserInfo = function() {
    var transport = new Thrift.Transport('http://127.0.0.1:8989/usersv/getUserInfo')
    var protocol = new Thrift.Protocol(transport)
    var client = new UserServiceClient(protocol)
    this.client = client
    var username = new Stringreq()
    username.s = ""
    var res = client.GetUserInfo(username)
    return res
}

GetT = function() {
    var transport = new Thrift.Transport('http://127.0.0.1:8989/tagsv/getTagsByUID')
    var protocol = new Thrift.Protocol(transport)
    var client = new TagServiceClient(protocol)
    this.client = client
    var id = new Intreq()
    id.i = 0
    var res = client.GetTagsByUID(id)
    return res
}

GetBTL = function(a) {
    var transport = new Thrift.Transport('http://127.0.0.1:8989/blogsv/getBlogListByTagAndUID')
    var protocol = new Thrift.Protocol(transport)
    var client = new BlogServiceClient(protocol)
    this.client = client
    var id = new Int2req()
    id.i1 = 0
    id.i2 = a
    var res = client.GetBlogsByTIDAndUID (id)
    return res
}

GetBL= function() {
    var transport = new Thrift.Transport('http://127.0.0.1:8989/blogsv/getBlogListByUID')
    var protocol = new Thrift.Protocol(transport)
    var client = new BlogServiceClient(protocol)
    this.client = client
    var id = new Intreq()
    id.i = 0
    var res = client.GetBlogsByUid(id)
    return res
}

GetB= function(a) {
    var transport = new Thrift.Transport('http://127.0.0.1:8989/blogsv/getBlogByBID')
    var protocol = new Thrift.Protocol(transport)
    var client = new BlogServiceClient(protocol)
    this.client = client
    var id = new Intreq()
    id.i = a
    var res = client.GetBlogById(id)
    return res
}

GetC = function(a) {
    var transport = new Thrift.Transport('http://127.0.0.1:8989/commentsv/getCommentListByBID')
    var protocol = new Thrift.Protocol(transport)
    var client = new CommentServiceClient(protocol)
    this.client = client
    var id = new Intreq()
    id.i = a
    var res = client.GetCommentListByBID(id)
    return res
}