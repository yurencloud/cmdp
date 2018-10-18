## cmdp 命令行和单文件快速提示、上传、下载、执行工具（远程同步）

![demo](http://cloud.yurencloud.com/index.php/s/pCYZvM1AFJGtdeW/download)



#### 简介

cmdp相当于git中的一个小功能，可以可以搜索，记录，提交，上传，下载，执行自己的或他人的1行文字，或者1个文件

cmdp相当于wget，可以下载1个文件，可以通过关键词，下载1个文件

cmdp相当于alias，可以记录大量的命令别名，远程同步，只要登录cmdp账号，就可以使用和执行自己或他人的命令别名

cmdp相当于man，可以制作和查看各种语言的help

cmdp相当于小云盘，可以储存1行文字或1个文件


#### 功能

##### 单行文本 (command,code,path,content...)

- 上传记录命令，代码，路径，文本等，并可添加关键词，注释。
- 搜索自己或他人添加的命令，代码，路径，文本等。（支持全局模糊搜索，包括内容，关键词，注释）
- 快速执行自己或他人添加的命令。

##### 单个文件(.yaml.md,.txt,.sh,.js,.java,.py,.php,Vagrantfile,Dockerfile...)

- 上传单个文件，并可添加关键词，注释。
- 搜索自己或他人上传的单个文件。（支持全局模糊搜索，包括关键词，注释）
- 下载自己或他人上传的单个文件。（根据关键词）
- 快速执行自己或他人文件命令。（文件内容是命令时）



#### 使用场景

> 如果你也经常遇到下述场景，那么cmdp，就可以帮助你快速解决此类问题，并且会成为你经常使用的一个命令行工具！

命令行参数太多，太难记，平常可能会为了1条命令，新建一个笔记，来记录，查询麻烦。

```
// 快速创建笔记，并自己上传此命令到你的账号中
cmdp c "docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=secret -d mysql" mysql

// 搜索和mysql相关的所有命令
cmdp s mysql

// 快速执行
// 相当于执行了docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=secret -d mysql
cmdp e mysql

// 如果自己不想创建笔记，可以直接搜索或者立即执行官方或其他大神创建的笔记
cmdp s docker/mysql
cmdp s dashen/mysql
cmdp e docker/mysql 
cmdp e dashen/mysql
```

经常要在云盘或笔记中收藏各种shell脚本，bash命令，没有地方快速收藏，快速执行。

```
// 上传shell脚本文件
cmdp p my.sh
// 上传markdown笔记文件
cmdp p my.md
// 上传单个代码文件，比如全国城市村镇的json
cmdp p country.json

// 搜索所有和mysql相关的我的文件
cmdp f mysql

// 执行shell脚本文件
cmdp e my.sh
// 下载markdown笔记文件
cmdp l my.md
// 下载全国城市村镇的json
cmdp l country.json

// 如果自己不想上传文件，可以直接搜索或者下载或者立即执行官方或其他dashen上传的文件
cmdp s office/a.sh
```

虽然linux有ctrl+R查询，但对于经常创建linux系统，使用不同的服务器，不同的电脑的人来说，命令提示不能同步到所有电脑。

```
// 搜索如何解压tar文件命令
cmdp s tar 

// 搜索如何快速显示所有目录的总大小命令
cmdp s du

// 搜索自己创建的docker命令
cmdp s docker
```

docker创建容器的命令太长，规则太多。经常需要在笔记中记录如何创建，或者记录下命令。

```
// 搜索自己创建并记录的所有和docker相关的命令
cmdp s docker
// 直接执行命令，创建mysql容器
cmdp e dockerMysql
```

Dockerfile，docker-compose.yaml等容器配置文件，如果想多台电脑通用，总是得收藏，或放到云盘，或多台电脑同步文件

```
// 下载java spring服务器容器的Dockerfile
cmdp l java
// 下载官方或dashen上传的java spring服务器容器的Dockerfile
cmdp l office/java
cmdp l dashen/java
```

Vagrantfile容器文件经常要放到自己的云盘

```
// 下载适用windows的hyperV专用的centos7的Vagrantfile
// 并启动容器
cmdp l centos7 && vagrant up
```

.gitignore.npmignore,package.json,各种config文件，经常要复制，粘贴

```
// 下载自己的.gitgnore文件
cmdp l .git

// 下载官方或dashen提供的.gitgnore文件
cmdp l dashen/.gitgnore
```

经常要分享同一个文件，配置文件给不同的同事或合作伙伴。

```
cmdp l xiaoming/config.json
```

等等



### 安装

#### 方式一：go安装

先安装golang

然后，可直接使用下面命令安装
~~~
go get github.com/yurencloud/cmdp
~~~



#### 方式二: wget 直接安装 （仅限linux）

以管理员身份执行下述命令

若正常输出版本号，即为安装成功`cmdp version 2.0.0`
64位
```
wget https://github.com/yurencloud/cmdp/releases/download/v2.1.0/cmdp.linux.tar.gz && tar -zxvf cmdp.linux.tar.gz && rm -rf cmdp.linux.tar.gz && chmod +x cmdp && mv cmdp /bin/cmdp && cmdp version
```
32位
```
wget https://github.com/yurencloud/cmdp/releases/download/v2.1.0/cmdp.linux32.tar.gz && tar -zxvf cmdp.linux32.tar.gz && rm -rf cmdp.linux.tar.gz && chmod +x cmdp && mv cmdp /bin/cmdp && cmdp version
```



#### 方式二：直接下载
> 建议尽量先安装golang，然后使用`go get github.com/yurencloud/cmdp`命令安装，这样无需设置环境变量或执行权限等问题

点击[下载地址](https://github.com/yurencloud/cmdp/releases)

解压后，工具可以直接使用，但建议将命令工具所在目录添加到系统路径`PATH`中

```
//windows
设置环境变量
点击Path，编辑
假设cmdp在C:\tool\cmdp，只要在Path中添加一条C:\tool

//linux,mac
假设cmdp在/home/tool/cmdp
vi ~/.bash_profile 或.bashrc
添加export PATH=$PATH:/home/tool
source ~/.bash_profile
```

在mac或linux若出现`Permission denied`问题，请用下面方法添加可执行权限
~~~
chmod +x cmdp
~~~



### 快速使用

```
$ cmdp help
NAME:
   cmdp - A new cli application

USAGE:
   cmdp [global options] command [command options] [arguments...]

VERSION:
   2.1.0

COMMANDS:
     search, s       search command, code, account, text, etc.
     version, v      show version
     create, c       create command to remote
     delete, d       delete command by id
     exec, e         exec command
     register        user register
     login           login by username and password
     logout          logout
     reset           reset password
     info            update user introduction
     push, p         push your file to remote
     pull, l, pl, P  pull your file from remote
     find, f         find your files from remote by keyword
     remove, r       remove your remote file by id
     star            star other user
     update          update cmdp version
     user, u         search users, order by stars count desc, cmds count desc, files count desc
     help, h         Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version

```



~~~
以下均以linux下操作为例，windows下也差不多
~~~


#### 注册

~~~
cmdp register USERNAME PASSWORD
~~~
注册后自动在HOME目录下生成.cmdp/.cmdp_token文件，*请不要删除*，token有效期1年

#### 登录

若已经注册，或切换登录，或token过期，或token删除，请使用登录功能，会生成新的token
~~~
cmdp login USERNAME PASSWORD
~~~

#### 修改用户简介

他人关注你时，简介会显示在他的关注列表

```
cmdp info CONTENT
cmdp info "我是tom,专门添加shell命令提示和bash脚本收藏"
```

#### 修改用户登录密码

```
cmdp reset NEW_PASSWORD
cmdp reset 654321
```

#### 搜索已经注册的所有用户！！

搜索结果按star,cmds,files数量依次排序

其中cmds,files仅统计公开的数量

```
cmdp user KEYWORD(可选) [user,u]

// 搜索所有用户
cmdp u

// 搜索用户简介info中，包含shell的用户，他可能精通shell
cmdp u shell

// 搜索指定用户名的用户，比如cmdp作者mackwang
cmdp u mackwang
```

![image](http://cloud.yurencloud.com/index.php/s/0MBRPyLG0zM3Cet/download)

> 搜索这些用户干什么用呢？

比如你想获取大量和docker有关的命令或文件

你就可以先搜索

```
cmdp u docker
```



然后根据排名，关注第一个用户，例如叫mackwang

```
cmdp star mackwang
```



然后查看他的所有的cmds和files，这样你就可以直接使用他的所有公开的命令或文件啦

```
// 注意，/符号不能少，符号左边是用户名，符号右边是关键词，留空时，默认搜索全部
cmdp s mackwang/
cmdp f mackwang/
```



其他参数

```
--page,-p 设置页数

--size,-s 设置每页显示条数，默认20

--all,-a 显示全部
```



#### 添加单行文本

~~~
cmdp create CONTENT KEYWORD(可选) COMMENT(可选) [create,c]
//示例
cmdp c "docker start mysql"
cmdp c "docker start mysql" dockerMysql "使用docker启动mysql容器"
~~~

> 安全：你创建的所有东西，默认都是隐私的，保密的，只有你可见，若想公开，可以添加--public,-p参数

```
cmdp c "docker start mysql" -p
```

#### 搜索单行文本

~~~
cmdp search CONTENT [search,s] 
//示例
cmdp s docker
~~~
显示结果，彩色
~~~
docker start mysql |  使用docker启动mysql容器 id:2
mysql -uroot -p |    docker登入mysql id:31
total:2, size:20, page:1/1
~~~

其他参数

```
--page,-p 设置页数
--size,-s 设置每页显示条数，默认20
--all,-a 显示全部
```

搜索他人的单行文本

```
// CONTENT中如果包含"/"符号的，左边为他人用户名，右边为搜索内容
// 只会显示他人公开的内容
cmdp s mackwang/docker
// 搜索他人的全部内容, 右边为空即可
cmdp s mackwang/
```

#### 执行单行文本命令

```
cmdp exec KEYWORD [exec,e]
// 示例
// 执行自己创建的命令
cmdp e dockerMysql
// 执行他人创建的命令,注意，因为执行他人命令是一件非常危险的事情，所以默认只是显示文本，若想执行，需要添加--force,-F参数
// 只有他人公开的内容，你才可见
cmdp e tom/dockerMysql --force
```

#### 删除单选文本

先查询，后根据结尾显示的id进行删除
~~~
cmdp delete ID 
cmdp d 14
~~~

#### 上传单个文件

```
cmdp push PATH KEYWORD(可选) COMMENT(可选) [push,p]
//示例
// 若不写keyword,默认keyword为文件名
cmdp p test.sh
cmdp p test.sh test "快速单元测试"
```

> 安全：你创建的所有东西，默认都是隐私的，保密的，只有你可见，若想公开，可以添加--public,-p参数

```
cmdp p test.sh -p
```

#### 查找单个文件

```
cmdp find CONTENT [find,f] 
//示例
cmdp f dockerfile
```

显示结果，彩色

```
dockerfile start mysql |  使用docker启动mysql容器 id:2
mysql -uroot -p |    docker登入mysql id:31
total:2, size:20, page:1/1
```

其他参数

```
--page,-p 设置页数
--size,-s 设置每页显示条数，默认20
--all,-a 显示全部
```

查找他人的单个文件

```
// CONTENT中如果包含"/"符号的，左边为他人用户名，右边为搜索内容
// 只会显示他人公开的内容
cmdp f mackwang/dockerfile
// 搜索他人的全部内容, 右边为空即可
cmdp s mackwang/
```

#### 下载单个文件

```
cmdp pull KEYWORD [pull,l,pl,P]
// 示例
cmdp l dockerfile
// 下载他人的单个文件
cmdp l tom/dockerfile
默认会下载到当前目录
```

仅打印文件内容，不下载

```
cmdp l dockerfile -p
```



#### 执行单个文件命令

```
cmdp exec KEYWORD [exec,e]
// 示例
// 执行自己创建的单个文件命令（单个文件内容得是可执行命令）,注意添加--file,-f参数，默认是执行单行文本命令
cmdp e dockerMysql -f
// 执行他人创建的单个文件命令,注意，因为执行他人命令是一件非常危险的事情，所以默认只是显示单个文件的文本，若想执行，需要添加--force,-F参数
// 只有他人公开的文件才可下载或执行
cmdp e tom/dockerMysql -f --force
```

#### 删除单个文件

先查找，后根据结尾显示的id进行删除

```
cmdp remvoe ID [remove,r]
cmdp r 14
```

#### 关注或收藏他人（其他用户，组织，官方账号）

```
cmdp star USERNAME
// 示例
cmdp star mackwang
```

#### 显示所有的关注或收藏列表

```
// 不添加任务参数，就会显示所有已关注或收藏的列表
cmdp star
```

#### 取消关注

```
先找到该关注的id,再取消
cmdp star -d ID
// 示例
cmdp start -d 12
```

####  快速升级cmdp到最新版本

```
cmdp update
```

