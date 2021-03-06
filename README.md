


<h2 align="center">cmdp</h2>
<p align="center">Shell脚本命令文件跨平台收藏、搜索和同步</p>
<p align="center">
    <img src="https://img.shields.io/github/stars/yurencloud/cmdp.svg" alt="">
    <img src="https://img.shields.io/github/issues/yurencloud/cmdp.svg" alt="">
    <img src="https://img.shields.io/github/forks/yurencloud/cmdp.svg" alt="">
    <img src="https://img.shields.io/github/license/yurencloud/cmdp.svg" alt="">
</p>

<p align="center">
  <a href="https://cmdp.yurencloud.com/document" target="_blank">使用文档</a>
  |
  <a href="https://cmdp.yurencloud.com/user"  target="_blank">用户排名</a>
  |
  <a href="http://cmdp.yurencloud.com"  target="_blank">官方网站</a>
</p>

欢迎到官方网站上查看`官方推荐`、`用户排名`、`个人中心` 可以快速查看他人和自己的命令或文件，以及可视化创建命令及文件。

#### 一、安装

方式1：支持各平台直接下载使用

[mac](https://github.com/yurencloud/cmdp/releases/download/v3.0.0/cmdp.mac.v3.0.0.tar.gz) | [windows](https://github.com/yurencloud/cmdp/releases/download/v3.0.0/cmdp.windows.v3.0.0.tar.gz) |[linux64](https://github.com/yurencloud/cmdp/releases/download/v3.0.0/cmdp.linux.v3.0.0.tar.gz) | [linux32](https://github.com/yurencloud/cmdp/releases/download/v3.0.0/cmdp.linux32.v3.0.0.tar.gz)

方式2：支持go安装

```
go get github.com/yurencloud/cmdp
```

方式3：支持shell脚本一键安装（仅限linux平台）

```
wget https://github.com/yurencloud/cmdp/releases/download/v3.0.0/cmdp.linux.v3.0.0.tar.gz && tar -zxvf cmdp.linux.v3.0.0.tar.gz && rm -rf cmdp.linux.v3.0.0.tar.gz && chmod +x cmdp && mv cmdp /bin/cmdp && cmdp version && echo "install success"
```



`NEW`已经公开cmdp服务端程序，允许自行搭建cmdp服务端 

包含cmdp服务端的接口文档和搭建教程

[cmdp服务端搭建教程](https://github.com/yurencloud/cmdp-server-public)

#### 二、快速使用

搜索自己收藏的所有和mysql相关的shell命令或脚本文件`全文模糊搜索`

```bash
cmdp s mysql
```

```bash
*** 输出 ***
docker exec -it mysql bash |  docker进入mysql容器 private id:9
docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD={{password}} -d mysql | docker-mysql docker快速创建mysql容器 public id:120
...
```

快速执行创建docker mysql 容器命令

```bash
cmdp e docker-mysql
# 相当于执行 docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD={{password}} -d mysql
```

支持参数替换

```bash
cmdp e docker-mysql 123456
# 或 cmdp e docker-mysql password=123456
# 相当于执行 docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -d mysql
```

支持执行他人公开的命令

```
cmdp e -F someone/docker-mysql
```

支持上传和下载脚本文件

```bash
cmdp p updateYumToAliyun.sh yum "快速切换yum源为阿里云源脚本"
cmdp l yum # 下载
```

支持执行远程脚本文件

```
cmdp e -f yum
cmdp e -F -f someone/yum # 执行他人的公开的脚本文件
# 相当于直接执行远程的 updateYumToAliyun.sh 脚本文件
```



其他特性

- 支持关注他人
- 支持fork他人命令或文件
- 支持多平台登录
- 支持远程同步命令
- 支持模糊搜索
- 支持循环嵌套支持命令
- 更多特性见下文...



![demo](http://cloud.yurencloud.com/index.php/s/pCYZvM1AFJGtdeW/download)



### 二、详细文档

```
$ cmdp help
NAME:
   cmdp - A new cli application

USAGE:
   cmdp [global options] command [command options] [arguments...]

VERSION:
   3.0.0

COMMANDS:
     search, s       search command, code, account, text, etc.
     version, v      show version
     create, c       create command to remote
     delete, d       delete command by id
     exec, e         exec command
     forkcmd, fc     fork command, code, account, text, etc.
     register        user register
     login           login by username and password
     logout          logout
     reset           reset password
     info            update user introduction
     push, p         push your file to remote
     pull, l, pl, P  pull your file from remote
     find, f         find your files from remote by keyword
     remove, r       remove your remote file by id
     forkfile, ff    fork file.
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



#### 添加命令

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



#### 搜索命令

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

搜索他人的命令

```
// CONTENT中如果包含"/"符号的，左边为他人用户名，右边为搜索内容
// 只会显示他人公开的内容
cmdp s mackwang/docker
// 搜索他人的全部内容, 右边为空即可
cmdp s mackwang/
```



#### 执行命令

```
cmdp exec KEYWORD [exec,e]
// 示例
// 执行自己创建的命令
cmdp e dockerMysql
// 执行他人创建的命令,注意，因为执行他人命令是一件非常危险的事情，所以默认只是显示文本，若想执行，需要添加--force,-F参数
// 只有他人公开的内容，你才可见 [--force,-F]
cmdp e --force tom/dockerMysql
```



#### 添加占位参数

```bash
{{PARAM_NAME}}
# docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD={{password}} -d mysql
# mysql -u{{username}} -p{{password}}
```



#### 执行命令时使用占位参数

```
// 如果mysql命令是mysql -u{{username}} -p{{password}}

cmdp e mysql root 123456
cmdp e mysql root password=123456
cmdp e mysql username=root password=123456
# 具名参数优先替换，然后匿名参数依次替换
```



#### 删除命令

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
cmdp e -f --force tom/dockerMysql
```



#### 单个文件添加占位参数（添加方式和命令一样）

```bash
{{PARAM_NAME}}
# docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD={{password}} -d mysql
# mysql -u{{username}} -p{{password}}
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



#### fork他人的命令

fork是复制拷贝，所以即使他人删除此命令，你fork的命令仍然在。

```
cmdp forkcmd USERNAME/KEYWORD [forkcmd,fc]
// 示例
cmdp forkcmd mackwang/mysql
cmdp fc mackwang/mysql
```



#### fork他人的文件

fork是复制拷贝，所以即使他人删除此文件，你fork的文件仍然在。

```
cmdp forkfile USERNAME/KEYWORD [forkfile,ff]
// 示例
cmdp forkfile mackwang/mysql.sh
cmdp ff mackwang/mysql.sh
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



#### 三、简介

cmdp相当于git中的一个小功能，可以可以搜索，记录，提交，上传，下载，执行自己的或他人的1行文字，或者1个文件

cmdp相当于wget，可以下载1个文件，可以通过关键词，下载1个文件

cmdp相当于alias，可以记录大量的命令别名，远程同步，只要登录cmdp账号，就可以使用和执行自己或他人的命令别名

cmdp相当于man，可以制作和查看各种语言的help

cmdp相当于小云盘，可以储存1行文字或1个文件



#### 四、功能

##### 单行文本 (command,code,path,content...)

- 上传记录命令，代码，路径，文本等，并可添加关键词，注释。
- 搜索自己或他人添加的命令，代码，路径，文本等。（支持全局模糊搜索，包括内容，关键词，注释）
- 快速执行自己或他人添加的命令。

##### 单个文件(.yaml.md,.txt,.sh,.js,.java,.py,.php,Vagrantfile,Dockerfile...)

- 上传单个文件，并可添加关键词，注释。
- 搜索自己或他人上传的单个文件。（支持全局模糊搜索，包括关键词，注释）
- 下载自己或他人上传的单个文件。（根据关键词）
- 快速执行自己或他人文件命令。（文件内容是命令时）

---



#### 五、使用场景

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






