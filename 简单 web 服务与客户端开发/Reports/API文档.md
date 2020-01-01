# API说明
使用rpc风格（详细请求/响应内容在表格以下）：

| api                            | method | description                        |
| ------------------------------ | ------ | ----------------------------------|
| /api                           | GET    | 获得所有支持API                    |
| /usersv/getUserInfo            | POST   | 通过用户名获取相应用户信息         |
| /blogsv/getBlogListByUID       | POST   | 通过用户id获取相应博客列表         |
| /blogsv/getBlogByBID           | POST   | 通过博客id获取相应博客内容         |
| /blogsv/getBlogListByTagAndUID | POST   | 通过标签id和用户id获取相应博客列表 |
| /commentsv/getCommentListByBID | POST   | 通过标签id和用户id获取相应博客列表 |
| /tagsv/getTagsByUID            | POST   | 通过用户id获取标签列表             |


具体请求/响应内容：

```
struct User {
  1: required i32 id,
  2: required string name,
  10: required i32 creacnt, 
  11: required i32 fancnt,
  12: required i32 zancnt,
  13: required i32 commentcnt,
  14: required i32 visitcnt
}

struct Blog {
  1: required i32 bid,
  2: required string uname,
  3: required bool iscopy,
  4: required string tname,
  10: required string title,
  11: required string content,
  12: required i32 readcnt, 
  13: required string cdate,
  14: required i32 ccnt
}

struct Tag {
  1: required i32 uid,
  2: required i32 tid,
  10: required string tname,
}

struct Comment {
  1: required i32 bid,
  2: required string uname,
  3: required string content,
  10: required string date,
  12: required i32 zancnt
}

struct Intreq {
  1: required i32 i
}

struct Int2req {
  1: required i32 i1,
  2: required i32 i2
}

struct Stringreq {
  1: required string s
}

service UserService {
  User GetUserInfo(1: Stringreq s)
}

service TagService {
  list<Tag> GetTagsByUID(1: Intreq i)
}

service BlogService {
  list<Blog> GetBlogsByTIDAndUID(1: Int2req i),
  list<Blog> GetBlogsByUid(1: Intreq i),
  Blog GetBlogById(1: Intreq i)
}

service CommentService {
  list<Comment> GetCommentListByBID(1: Intreq i)
}

```

