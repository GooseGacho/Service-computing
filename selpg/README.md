
# selpg
具体要求与说明：[ 开发 Linux 命令行实用程序 ](https://www.ibm.com/developerworks/cn/linux/shell/clutil/index.html)
# 说明
## 1. 命令格式  
```java
selpg [-s startPage] [-e endPage] [-l linePerPage | -f] [-d dest] input_file >output_file 2>error_file
   ```

## 2. 参数
   `-s`：startPage，后面接开始读取的页号   
    `-e`：endPage，后面接结束读取的页号  
    `-l`：后面跟行数，代表多少行分为一页  
    `-f`：该标志无参数，代表按照分页符’\f’ 分页，一般默认72行为一页  
    `-d`：“-dDestination”选项将选定的页直接发送至打印机，“Destination”应该是 lp 命令“-d”选项可接受的打印目的地名称  
   ` input_file`，`output_file 2`，`error_file`：输入文件、输出文件、错误信息文件的名字  
   
(-s和-e是必须的参数，其它为可选参数，-l和-f参数不可能同时出现）

# 设计

 1. 导入所需包
 ```java
 package main

import (

	"fmt"
	"flag"
	"os"
	"io"
	"bufio"
	"os/exec"

)
```

 2. 定义结构体
 ```java
 type selpg_args struct
{
	start_page int
	end_page int
	in_filename string
	page_len int
	page_type bool

	print_dest string
}
 ```
3. main函数
```java
func main(){
	var args selpg_args
	get(&args)
	process_args(&args)
	process_input(&args)
}
```
4. 实现各个函数的功能
```java
get(&args)	//初始化，定义结构体各个变量的值
process_args(&args)	//检查输入的命令是否合理有效
process_input(&args)	//处理输入的命令
```
# 运行使用
使用以下命令
```java
go install github.com/github-user/selpg
selpg [-s startPage] [-e endPage] [-l linePerPage | -f] [-d dest] input_file >output_file 2>error_file
```
# 测试结果
按文档 **使用 selpg** 章节要求测试你的程序   
input.txt里面是1到100行的内容分别为1到100的输入文件

1. `$ selpg -s1 -e1 input_file`  
![](https://github.com/SHIKD-pu/Service-computing/blob/master/selpg/picture/1-1.PNG)  
一直到72  
![](https://github.com/SHIKD-pu/Service-computing/blob/master/selpg/picture/1-3.PNG)  
2. `$ selpg -s1 -e1 < input_file`  
![](https://github.com/SHIKD-pu/Service-computing/blob/master/selpg/picture/2.PNG)  
也是输出直到72，**后面的图片结果省略22到72的图片**  
3. `$ other_command | selpg -s10 -e20`  
![](https://github.com/SHIKD-pu/Service-computing/blob/master/selpg/picture/3.PNG)  
4. `$ selpg -s10 -e20 input_file >output_file`  
![](https://github.com/SHIKD-pu/Service-computing/blob/master/selpg/picture/4-1.PNG)  
![](https://github.com/SHIKD-pu/Service-computing/blob/master/selpg/picture/4-2.PNG)  
5. `$ selpg -s10 -e20 input_file 2>error_file`  
![](https://github.com/SHIKD-pu/Service-computing/blob/master/selpg/picture/5-1.PNG)  
![](https://github.com/SHIKD-pu/Service-computing/blob/master/selpg/picture/5-2.PNG)  
6. `$ selpg -s10 -e20 input_file >output_file` 2>error_file  
![](https://github.com/SHIKD-pu/Service-computing/blob/master/selpg/picture/6.PNG)  
7. `$ selpg -s10 -e20 input_file >output_file` 2>/dev/null  
![](https://github.com/SHIKD-pu/Service-computing/blob/master/selpg/picture/7.PNG)  
8. `$ selpg -s10 -e20 input_file >/dev/null`  
![](https://github.com/SHIKD-pu/Service-computing/blob/master/selpg/picture/8.PNG)  
9. `$ selpg -s10 -e20 input_file | other_command`  
![](https://github.com/SHIKD-pu/Service-computing/blob/master/selpg/picture/9.PNG)  
10. `$ selpg -s10 -e20 input_file 2>error_file | other_command`  
![](https://github.com/SHIKD-pu/Service-computing/blob/master/selpg/picture/10.PNG)  
---  
1. `$ selpg -s10 -e20 -l66 input_file`  
![](https://github.com/SHIKD-pu/Service-computing/blob/master/selpg/picture/21.PNG)  
2. `$ selpg -s10 -e20 -f input_file`  
![](https://github.com/SHIKD-pu/Service-computing/blob/master/selpg(/picture/22.PNG)  
3. `$ selpg -s10 -e20 -dlp1 input_file`  
![](https://github.com/SHIKD-pu/Service-computing/blob/master/selpg/picture/23.PNG)  
4. `$ selpg -s10 -e20 input_file > output_file 2>error_file &`  
![](https://github.com/SHIKD-pu/Service-computing/blob/master/selpg/picture/24.PNG)   
