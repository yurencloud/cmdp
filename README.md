## cmdp 命令行和单文件快捷提示、下载、执行工具（远程同步）

![demo](http://cloud.yurencloud.com/index.php/s/pCYZvM1AFJGtdeW/download)



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

#### 若你已经安装go
可直接使用下面命令安装
~~~
go get github.com/yurencloud/cmdp
~~~

#### 下载安装

> 建议尽量先安装golang，然后使用`go get github.com/yurencloud/cmdp`命令安装

`待工具使用稳定后再提供下载地址，请先使用go安装`

工具可以直接使用，但建议将命令工具所在目录添加到系统路径`PATH`中

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
   2.0.0

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

---

一个人可以申请多个用户名，以方便创建各种类型的单行文本或单个文件提示，搜索，执行，方便大众搜索，下载，执行，使用。

以下用户名已经被保护，若想注册，请发邮件给我 641212003@qq.com（附上github，博客地址）

```
css,javascript,html,shell,git,java,php,python,sql,ruby,object-c,visualbasic,perl,swift,golang,scala,apache,bash,lua,npm,node,react,vue,redis,mongodb,mysql,svn,k8s,kubernetes,vagrant,linux,centos,ubuntu,blockchain,bitcoin,ethereum,android,ios,net,angular,jquery,nodejs,reactnative,hadoop,docker,spark,hive,oracle,unity3d,cocos,
```

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

#### 添加单行文本

~~~
cmdp create CONTENT KEYWORD(可选) COMMENT(可选) [create,c]
//示例
cmdp c "docker start mysql"
cmdp c "docker start mysql" dockerMysql "使用docker启动mysql容器"
~~~

> 注意你创建的所有东西，默认都是公开的，请不要上传敏感内容，例如密码，私钥等

如果你要创建非公开的，私密的内容请添加参数 --private,-p

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
cmdp s mack/docker
```

#### 执行单行文本命令

```
cmdp exec KEYWORD [exec,e]
// 示例
// 执行自己创建的命令
cmdp e dockerMysql
// 执行他人创建的命令,注意，因为执行他人命令是一件非常危险的事情，所以默认只是显示文本，若想执行，需要添加--force,-F参数
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

> 注意你创建的所有东西，默认都是公开的，请不要上传敏感内容，例如密码，私钥等

如果你要创建非公开的，私密的内容请添加参数 --private,-p

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
cmdp f mack/dockerfile
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

#### 执行单个文件命令

```
cmdp exec KEYWORD [exec,e]
// 示例
// 执行自己创建的单个文件命令（单个文件内容得是可执行命令）,注意添加--file,-f参数，默认是执行单行文本命令
cmdp e dockerMysql -f
// 执行他人创建的单个文件命令,注意，因为执行他人命令是一件非常危险的事情，所以默认只是显示单个文件的文本，若想执行，需要添加--force,-F参数
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
