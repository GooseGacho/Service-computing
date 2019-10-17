# go_agenda


## Agenda 业务需求

### 用户注册

* 注册新用户时，用户需设置一个唯一的用户名和一个密码。另外，还需登记邮箱及电话信息。
* 如果注册时提供的用户名已由其他用户使用，应反馈一个适当的出错信息；成功注册后，亦应反馈一个成功注册的信息。
* 使用说明：register -u username -P password -e email -p phone
* 运行结果：
```
//成功
[wyy@centos7 agenda]$ go run main.go register -u wuyy -P 2333 -e 121212@163.com -p 17766727
2018/10/31 19:54:39 register called by wuyy 2333 121212@163.com 17766727
2018/10/31 19:54:39 READUSER success
2018/10/31 19:54:39 WRITEUSER success

//已注册，失败
[wyy@centos7 agenda]$ go run main.go register -u wuyy -P 2333 -e 121212@163.com -p 17766727
2018/10/31 20:04:25 register called by wuyy 2333 121212@163.com 17766727
2018/10/31 20:04:25 READUSER success
2018/10/31 20:04:25 username already registered

```

### 用户登录

* 用户使用用户名和密码登录 Agenda 系统。
* 用户名和密码同时正确则登录成功并反馈一个成功登录的信息。否则，登录失败并反馈一个失败登录的信息。
* 使用说明：login -u username -p password
* 运行结果：
```
// 成功
[centos-manager@centos-manager agenda]$ go run main.go login -u wxl -p 123
2018/11/01 21:15:45 READUSER success
2018/11/01 21:15:45 Log in Success!
2018/11/01 21:15:45 Welcome! wxl

// 用户不存在
[centos-manager@centos-manager agenda]$ go run main.go login -u wrongUser -p 111
2018/11/01 21:13:59 READUSER success
2018/11/01 21:13:59 Warning! Wrong UserName

// 密码错误
[centos-manager@centos-manager agenda]$ go run main.go login -u wxl -p 111
2018/11/01 21:14:31 READUSER success
2018/11/01 21:14:31 Warning! Wrong Password

```

### 用户登出

* 已登录的用户登出系统后，只能使用用户注册和用户登录功能。
* 使用说明：logout
* 运行结果：
```
[wyy@centos7 agenda]$ go run main.go logout
2018/10/31 19:56:22 logout success

```

### 用户查询

* 已登录的用户可以查看已注册的所有用户的用户名、邮箱及电话信息。
* 使用说明：searchUser -u username
* 运行结果：
```
// 查询所有用户
[centos-manager@centos-manager agenda]$ go run main.go searchUser -u _ALL_
2018/11/01 21:23:53 READUSER success
2018/11/01 21:23:53 NAME: TestUser   EMAIL: 123@qq.com   TEL: 123
2018/11/01 21:23:53 NAME: wuyy   EMAIL: 121212@163.com   TEL: 17766727
2018/11/01 21:23:53 NAME: wxl   EMAIL: 123@163.com   TEL: 12121

// 查询特定用户
[centos-manager@centos-manager agenda]$ go run main.go searchUser -u wxl
2018/11/01 21:25:18 READUSER success
2018/11/01 21:25:18 NAME: wxl   EMAIL: 123@163.com   TEL: 12121

// 查询用户不存在
[centos-manager@centos-manager agenda]$ go run main.go searchUser -u noUser
2018/11/01 21:27:25 READUSER success
2018/11/01 21:27:25 User does not exist

```

### 用户删除

* 已登录的用户可以删除本用户账户（即销号）。
* 操作成功，需反馈一个成功注销的信息；否则，反馈一个失败注销的信息。
* 删除成功则退出系统登录状态。删除后，该用户账户不再存在。
* 用户账户删除以后：
   * 以该用户为 发起者 的会议将被删除
   * 以该用户为 参与者 的会议将从 参与者 列表中移除该用户。若因此造成会议 参与者 人数为0，则会议也将被删除。
* 使用说明：deleteUser -p password
* 运行结果：
```
// 注销成功
[centos-manager@centos-manager agenda]$ go run main.go deleteUser -p 123456
2018/11/01 23:10:34 READUSER success
2018/11/01 23:10:34 READUSER success
2018/11/01 23:10:34 Delete user successfully.

// 注销失败，密码错误
[centos-manager@centos-manager agenda]$ go run main.go deleteUser -p 111111
2018/11/01 23:10:28 READUSER success
2018/11/01 23:10:28 Wrong password!

```

### 创建会议

* 已登录的用户可以添加一个新会议到其议程安排中。会议可以在多个已注册 用户间举行，不允许包含未注册用户。添加会议时提供的信息应包括：
   * 会议主题(title)（在会议列表中具有唯一性）
   * 会议参与者(participator)
   * 会议起始时间(start time)
   * 会议结束时间(end time)
   * 注意，任何用户都无法分身参加多个会议。如果用户已有的会议安排（作为发起者或参与者）与将要创建的会议在时间上重叠 （允许仅有端点重叠的情况），则无法创建该会议。
   * 用户应获得适当的反馈信息，以便得知是成功地创建了新会议，还是在创建过程中出现了某些错误。
* 使用说明：createMeeting -m meetingtitle -M month -d day -t [time]1/2/3/4 -c [command]update/new 增加新的时间/创建新的会议
* 运行结果：
```
// 成功
[centos-manager@centos-manager agenda]$ go run main.go createMeeting  -m TestMeeting -M 1 -d 1 -t 2 -c new
2018/11/01 21:41:20 READUSER success
2018/11/01 21:41:20 READMEETINGS success
2018/11/01 21:41:20 Apply Success! Please add Paticipators!
2018/11/01 21:41:20 WRITEUSER success
2018/11/01 21:41:20 WRITEMEETINGS success

// 更新会议时间成功，会议分两期两天进行
[centos-manager@centos-manager agenda]$ go run main.go createMeeting  -m TestMeeting -M 1 -d 1 -t 3 -c update
2018/11/01 22:22:38 READUSER success
2018/11/01 22:22:38 READMEETINGS success
2018/11/01 22:22:38 Apply Success!
2018/11/01 22:22:38 WRITEUSER success
2018/11/01 22:22:38 WRITEMEETINGS success
[centos-manager@centos-manager agenda]$ go run main.go createMeeting  -m TestMeeting -M 1 -d 3 -t 1 -c update
2018/11/01 22:22:53 READUSER success
2018/11/01 22:22:53 READMEETINGS success
2018/11/01 22:22:53 Apply Success!
2018/11/01 22:22:53 WRITEUSER success
2018/11/01 22:22:53 WRITEMEETINGS success

// 创建题目冲突
[centos-manager@centos-manager agenda]$ go run main.go createMeeting  -m cheat -M 1 -d 1 -t 1 -c new
2018/11/01 21:40:53 READUSER success
2018/11/01 21:40:53 READMEETINGS success
2018/11/01 21:40:53 Wrong! You should use -command update

// 创建时间冲突
[centos-manager@centos-manager agenda]$ go run main.go createMeeting  -m TestMeeting -M 1 -d 1 -t 1 -c new
2018/11/01 21:36:17 READUSER success
2018/11/01 21:36:17 READMEETINGS success
2018/11/01 21:36:17 Wrong! Time had applied!

```


### 增删会议参与者

* 已登录的用户可以向 自己发起的某一会议增加/删除 参与者 。
* 增加参与者时需要做 时间重叠 判断（允许仅有端点重叠的情况）。
* 删除会议参与者后，若因此造成会议 参与者 人数为0，则会议也将被删除。
* 使用说明：changeMeetingPar -m meetingtitle -c a/d 增加/删除 -p username
* 运行结果：
```
// 成功增加
[centos-manager@centos-manager agenda]$ go run main.go changeMeetingPar -m TestMeeting -c a -p wuyy
2018/11/01 22:02:20 READUSER success
2018/11/01 22:02:20 READMEETINGS success
2018/11/01 22:02:20 Add success!
2018/11/01 22:02:20 WRITEUSER success
2018/11/01 22:02:20 WRITEMEETINGS success
[centos-manager@centos-manager agenda]$ go run main.go changeMeetingPar -m TestMeeting -c a -p TestUser
2018/11/01 22:10:53 READUSER success
2018/11/01 22:10:53 READMEETINGS success
2018/11/01 22:10:53 Add success!
2018/11/01 22:10:53 WRITEUSER success
2018/11/01 22:10:53 WRITEMEETINGS success


// 成功删除用户
[centos-manager@centos-manager agenda]$ go run main.go changeMeetingPar -m TestMeeting -c d -p TestUser
2018/11/01 22:11:27 READUSER success
2018/11/01 22:11:27 READMEETINGS success
2018/11/01 22:11:27 Delete success!
2018/11/01 22:11:27 WRITEUSER success
2018/11/01 22:11:27 WRITEMEETINGS success

// 成功删除，会议无用户，自动清除
[centos-manager@centos-manager agenda]$ go run main.go changeMeetingPar -m TestMeeting -c d -p wuyy
2018/11/01 22:06:54 READUSER success
2018/11/01 22:06:54 READMEETINGS success
2018/11/01 22:06:54 Empty meeting! Delete Automitaic!
2018/11/01 22:06:54 Delete success!
2018/11/01 22:06:54 WRITEUSER success
2018/11/01 22:06:54 WRITEMEETINGS success

// 增删失败，会议不存在
[centos-manager@centos-manager agenda]$ go run main.go changeMeetingPar -m NoMeeting -c a -p wuyy
2018/11/01 21:53:49 READUSER success
2018/11/01 21:53:49 READMEETINGS success
2018/11/01 21:53:49 Dont has this Meeting

// 增加失败，被邀请用户是会议发起人
[centos-manager@centos-manager agenda]$ go run main.go changeMeetingPar -m TestMeeting -c a -p wxl
2018/11/01 22:01:13 READUSER success
2018/11/01 22:01:13 READMEETINGS success
2018/11/01 22:01:13 Cann't add sponsor to meeting as participate

// 增加失败，用户不存在
[centos-manager@centos-manager agenda]$ go run main.go changeMeetingPar -m TestMeeting -c a -p NoUser
2018/11/01 21:57:09 READUSER success
2018/11/01 21:57:09 READMEETINGS success
2018/11/01 21:57:09 Dont have user named NoUser

// 增加失败，重复增加
[centos-manager@centos-manager agenda]$ go run main.go changeMeetingPar -m TestMeeting -c a -p wuyy
2018/11/01 22:04:26 READUSER success
2018/11/01 22:04:26 READMEETINGS success
2018/11/01 22:04:26 wuyy was participator! Add failed

```

### 查询会议

* 已登录的用户可以查询自己的议程在某一时间段(time interval)内的所有会议安排。
* 用户给出所关注时间段的起始时间和终止时间，返回该用户议程中在指定时间范围内找到的所有会议安排的列表。
* 在列表中给出每一会议的起始时间、终止时间、主题、以及发起者和参与者。
* 注意，查询会议的结果应包括用户作为 发起者或参与者 的会议。
* 使用说明：searchMeeting -S startMonth -s startDay -E endMonth -e endDay
* 运行结果：
```
[centos-manager@centos-manager agenda]$ go run main.go searchMeeting -S 1 -s 1 -E 1 -e 3
2018/11/01 22:25:13 READMEETINGS success
2018/11/01 22:25:13 You Sponsor: 
2018/11/01 22:25:13 Title: TestMeeting
2018/11/01 22:25:13 Sponsor: wxl
2018/11/01 22:25:13 Date: 1.1
2018/11/01 22:25:13 11:00~12:00
2018/11/01 22:25:13 16:00~17:00
2018/11/01 22:25:13 Participate: 
2018/11/01 22:25:13 [wuyy]
2018/11/01 22:25:13 
2018/11/01 22:25:13 You Sponsor: 
2018/11/01 22:25:13 Title: TestMeeting
2018/11/01 22:25:13 Sponsor: wxl
2018/11/01 22:25:13 Date: 1.3
2018/11/01 22:25:13 10:00~11:00
2018/11/01 22:25:13 Participate: 
2018/11/01 22:25:13 [wuyy]

```

### 取消会议

* 已登录的用户可以取消 自己发起 的某一会议安排。
* 取消会议时，需提供唯一标识：会议主题（title）。
* 使用说明：deleteMeeting -m meeting
* 运行结果：
//成功取消
```
// 取消成功
[centos-manager@centos-manager agenda]$ go run main.go deleteMeeting -m TestMeeting1
2018/11/01 22:35:42 READUSER success
2018/11/01 22:35:42 READMEETINGS success
2018/11/01 22:35:42 Delete Meeting Success!
2018/11/01 22:35:42 WRITEUSER success
2018/11/01 22:35:42 WRITEMEETINGS success

//取消失败
[centos-manager@centos-manager agenda]$ go run main.go deleteMeeting -m NoMeeting
2018/11/01 22:36:35 READUSER success
2018/11/01 22:36:35 READMEETINGS success
2018/11/01 22:36:35 Dont have this Meeting

```

### 退出会议

* 已登录的用户可以退出 自己参与 的某一会议安排。
* 退出会议时，需提供一个唯一标识：会议主题（title）。若因此造成会议 参与者 人数为0，则会议也将被删除。
* 使用说明：exitMeeting -m meeting
* 运行结果：
```
// 退出成功
[centos-manager@centos-manager agenda]$ go run main.go exitMeeting -m TestMeeting
2018/11/01 22:32:02 READUSER success
2018/11/01 22:32:02 READMEETINGS success
2018/11/01 22:32:02 Exit meeting successfully
2018/11/01 22:32:02 WRITEUSER success
2018/11/01 22:32:02 WRITEMEETINGS success

// 未参加会议
[centos-manager@centos-manager agenda]$ go run main.go exitMeeting -m NoMeeting
2018/11/01 22:33:24 READUSER success
2018/11/01 22:33:24 READMEETINGS success
2018/11/01 22:33:24 Not Participate Meeting!

```

### 清空会议

* 已登录的用户可以清空 自己发起 的所有会议安排。
* 使用说明：clearMeeting -p password
* 运行结果：
```
// 清空成功
[centos-manager@centos-manager agenda]$ go run main.go clearMeeting -p 123456
2018/11/01 23:04:09 Delete TestMeeting2 Success!
2018/11/01 23:04:09 Delete TestMeeting3 Success!
2018/11/01 23:04:09 Exit TestMeeting1 successfully

// 清空失败，密码错误
[centos-manager@centos-manager agenda]$ go run main.go clearMeeting -p 111111
2018/11/01 22:40:13 READUSER success
2018/11/01 22:40:13 READMEETINGS success
2018/11/01 22:40:13 Wrong password!

```
