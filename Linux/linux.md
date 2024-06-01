# Linux 操作

## 软件安装与更新

> - apt： Advanced Packaging Tool
>
> - 早期使用apt-get，ubuntu 16使用apt

```bash
# 更新可用软件包列表
$ sudo apt update

# 更新已安装的包，如果增加、更换了软件源，通常需要使用 `update` 更新一下本地缓存
$ sudo apt upgrade 

# 安装软件
$ sudo apt install 软件名

# 卸载软件
$ sudo apt remove 软件名

# 安装之后，才可以在其他系统中，通过 ssh 工具远程登陆
sudo apt install openssh-server 
```

## 配置软件源

- 下载地址文件：`/etc/apt/source.list`
- [清华镜像源](https://mirrors.tuna.tsinghua.edu.cn/) ,找到Ubuntu，并点击旁边的问号。
- 选择对应版本，并替换source.list文件(替换前最好备份)

```shell
# 拷贝一份/etc/apt/sources.list文件
sudo cp -v /etc/apt/sources.list /etc/apt/sources.list.backup

# 编辑sources.list文件
sudo gedit /etc/apt/sources.list

# 全选并删除文件的内容，粘贴清华镜像源内容，并保存

# 更新
sudo apt-get update
```

## 快捷键

```shell
# 放大终端窗口字体
ctrl + shift + =
# 缩小终端窗口字体
ctrl + -
```

## 目录说明

![003-Linux的树形示意图](.\assets\003-Linux的树形示意图.png)

| 目录                              | 存放说明                                                     |
| --------------------------------- | ------------------------------------------------------------ |
| /home/user                        | 用户工作目录或家目录  ~                                      |
| /：根目录                         | 一般只存放目录，当在终端里输入 `/home`，其实是在告诉电脑，先从 `/`（根目录）开始，再进入到 `home` 目录 |
| /bin、/usr/bin                    | 可执行二进制文件的目录，如常用的命令 ls、tar、mv、cat 等     |
| /boot：                           | 放置 linux 系统启动时用到的一些文件，如 linux 的内核文件：`/boot/vmlinuz`，系统引导管理器：`/boot/grub` |
| /dev：                            | 存放linux系统下的设备文件，访问该目录下某个文件，相当于访问某个设备，常用的是挂载光驱`mount /dev/cdrom /mnt` |
| /etc                              | 系统配置文件存放的目录，不建议在此目录下存放可执行文件，重要的配置文件有：inittab、fstab、init.d、X11、sysconfig、xinetd.d |
| /home                             | 系统默认的用户家目录，新增用户账号时，用户的家目录都存放在此目录下<br />`~` 表示当前用户的家目录<br />`~edu` 表示用户 `edu` 的家目录 |
| /lib、/usr/lib、/usr/local/lib    | 系统使用的函数库的目录，程序在执行过程中，需要调用一些额外的参数时需要函数库的协助 |
| /lost+fount                       | 系统异常产生错误时，会将一些遗失的片段放置于此目录下         |
| /mnt: /media：                    | 光盘默认挂载点，通常光盘挂载于 /mnt/cdrom 下，也不一定，可以选择任意位置进行挂载 |
| /opt：                            | 给主机额外安装软件所摆放的目录                               |
| /proc：                           | 此目录的数据都在内存中，如系统核心，外部设备，网络状态，由于数据都存放于内存中，所以不占用磁盘空间，比较重要的文件有：cpuinfo、interrupts、dma、ioports、net/ |
| /root                             | 系统管理员root的家目录                                       |
| /sbin、/usr/sbin、/usr/local/sbin | 放置系统管理员使用的可执行命令，如 fdisk、shutdown、mount 等。与 /bin 不同的是，这几个目录是给系统管理员 root 使用的命令，一般用户只能"查看"而不能设置和使用 |
| /tmp：                            | 一般用户或正在执行的程序临时存放文件的目录，任何人都可以访问，重要数据不可放置在此目录下 |
| /srv：                            | 服务启动之后需要访问的数据目录，如 www 服务需要访问的网页数据存放在 /srv/www 内 |
| /usr：                            | 应用程序存放目录<br />/usr/bin：存放应用程序<br />/usr/share：存放共享数据<br />/usr/lib：存放不能直接运行的，却是许多程序           运行所必需的一些函数库文件<br />/usr/local：存放软件升级包<br />/usr/share/doc：系统说明文件存放目录<br />/usr/share/man：程序说明文件存放目录 |
| /var                              | 放置系统执行过程中经常变化的文件<br />/var/log：随时更改的日志文件<br />/var/spool/mail：邮件存放的目录<br />/var/run：程序或服务启动后，其 PID 存放在该目录下 |

## 文件和目录操作

###  文件与目录命令

| 命令                         | 作用                     |
| :--------------------------- | :----------------------- |
| ls                           | 查看当前文件夹下的内容   |
| pwd                          | 查看当前所在文件夹       |
| cd [目录名]                  | 切换文件夹               |
| touch [文件名]               | 如果文件不存在，新建文件 |
| mkdir [目录名]               | 创建目录                 |
| rm [文件名]                  | 删除指定的文件名         |
| clear                        | 清屏                     |
| tree [文件名]                | 以树状图列出文件目录结构 |
| cp 源文件 目标文件           | 复制文件夹               |
| mv 路径/源文件 路径/目标文件 | 移动或重命名文件或者目录 |
| bash xxx                     | 执行xxx文件              |

### Linux 终端命令格式

```shell
command [-options] [parameter]

command：命令名，相应功能的英文单词或单词的缩写
[-options]：选项，可用来对命令进行控制，也可以省略
[parameter]：传给命令的参数，可以是零个、一个或者多个

# 显示命令帮助信息
command --help
```

### 显示目录 

```shell
# 显示以.开头的隐藏文件
ls -a

# 以列表方式显示文件的详细信息
ls -l

# 配合 -l 以人性化的方式显示文件大小
ls -h

# 以树状图列出文件目录结构，只显示目录
tree -d xxxx
# tree指令需要安装才能使用
sudo apt install tree
```

### 切换目录 cd命令

> Linux 所有的 **目录** 和 **文件名** 都是大小写敏感的

```shell
# 切换到当前用户的主目录(/home/用户目录)
cd
# 切换到当前用户的主目录(/home/用户目录)
cd ~
# 保持在当前目录不变
cd .
# 切换到上级目录
cd ..
# 可以在最近两次工作目录之间来回切换
cd -
```

### 创建文件或修改文件时间 touch命令

- 如果文件 **不存在**，可以创建一个空白文件
- 如果文件 **已经存在**，可以修改文件的末次修改日期

### 创建新目录 mkdir命令

```shell
# 创建新目录
mkdir xxx
# 递归创建目录，新目录不能与已有目录重名
mkdir -p xxx/xxx
```

### 删除文件或目录 rm命令

```shell
# 删除后无法恢复

# 强制删除
rm -f xxxx
#递归和删除目录下的内容，删除文件夹必须加此参数
rm -r xxxx

```

### 复制文件或目录  cp命令

```shell
# 覆盖文件前提示
cp -i
# 递归复制该目录下的所有子目录和文件，目标文件必须为一个目录名
cp -r
```

### 移动目录或文件  mv命令

 ```shell
# 覆盖文件前提示
mv -i
 ```

## 查看文件内容

### 查看文件命令

| 命令                 | 作用                                                 |
| :------------------- | :--------------------------------------------------- |
| cat 文件名           | 查看文件内容、创建文件、文件合并、追加文件内容等功能 |
| more 文件名          | 分屏显示文件内容                                     |
| grep 搜索文本 文件名 | 搜索文本文件内容                                     |
| echo                 | 将文本输出在终端上                                   |

### cat

- cat命令用来 **查看文件内容**、**创建文件**、**文件合并**、**追加文件内容** 等功能
- cat 会一次显示所有的内容，适合 **查看内容较少** 的文本文件
- Linux 中还有一个 `nl` 的命令和 `cat -b` 的效果等价

```shell
# 对非空输出行编号
cat -b
# 对输出的所有行编号
cat -n
```

### more

- `more` 命令可以用于分屏显示文件内容，每次只显示一页内容
- 适合于 **查看内容较多**的文本文件

使用 `more` 的操作键：

| 操作键   | 功能                 |
| :------- | :------------------- |
| 空格键   | 显示手册页的下一屏   |
| Enter 键 | 一次滚动手册页的一行 |
| b        | 回滚一屏             |
| f        | 前滚一屏             |
| q        | 退出                 |
| /word    | 搜索 **word** 字符串 |

### grep

- `grep`允许对文本文件进行 **模式**查找，所谓模式查找，又被称为正则表达式

```shell
# 显示匹配行及行号
grep -n
# 显示不包含匹配文本的所有行（相当于求反）
grep -v
# 忽略大小写
grep -i
# 行首查找，搜寻以a开头的行
grep ^a
# 行尾查找，搜寻以ke结束的行
grep ke
```

### echo 与重定向

- `echo` 会在终端中显示参数指定的文字，通常会和 **重定向** 联合使用

```bash
# 将hello word 输出到终端上
echo "hello word"

# 将输出的helloword输出到txt文件中，覆盖形式
echo "hello word" > hello word.txt
# 追加输出
echo "hello word" >> hello word.txt

# 将tree指令输出在终端上的内容重定向到 txt文件中
tree > tree.txt


```

### 管道 `|` 

- Linux 允许将 **一个命令的输出** 可以**通过管道** 做为 **另一个命令的输入** 
- 可以理解现实生活中的管子，管子的一头塞东西进去，另一头取出来，这里 `|` 的左右分为两端，左端塞东西（写），右端取东西（读）

常用的管道命令有：

- `more`：分屏显示内容
- `grep`：在命令执行结果的基础上查询指定的文本

## 系统管理

### 关机/重启

```bash
# 一分钟后关闭电脑
shutdown

# 重新启动操作系统，其中 now 表示现在
shutdown -r now

# 立刻关机，其中 now 表示现在
shutdown now

# 系统在今天的 20:25 会关机
shutdown 20:25

# 系统再过十分钟后自动关机
shutdown +10

# 取消之前指定的关机计划
shutdown -c
```

### 查看时间信息

```bash
# 查看系统时间
date

# calendar查看日历，-y选项可以查看一年的日历
cal -y

```

### 查看磁盘信息

```bash
# disk free 显示磁盘剩余空间
df -h

# disk usage 显示目录下的文件大小，-h以人性化的方式显示文件大小
du -h [目录名]
```

### 查看进程信息

```bash
# process status 查看进程的详细状况，ps默认显示当前用户通过终端启动的应用程序，a显示终端上的所有进程，包括其他用户的进程,u显示进程的详细状态,x显示没有控制终端的进程
ps aux

# 动态显示运行中的进程并且排序,要退出top可以直接输入q
top

# 终止指定代号的进程，`-9` 表示强行终止，最好只终止由当前用户开启的进程，而不要终止 `root` 身份开启的进程，否则可能导致系统崩溃
kill [-9] 进程代号
```



### ifconfig

- `ifconfig` 可以查看／配置计算机当前的网卡配置信息

```bash
# 查看网卡配置信息
$ ifconfig

# 查看网卡对应的 IP 地址
$ ifconfig | grep inet
```

> 提示：一台计算机中有可能会有一个 **物理网卡** 和 **多个虚拟网卡**，在 Linux 中物理网卡的名字通常以 `ensXX` 表示

- `127.0.0.1` 被称为 **本地回环/环回地址**，一般用来测试本机网卡是否正常

### ping

```bash
# 检测到目标主机是否连接正常
$ ping IP地址

# 检测本地网卡工作正常
$ ping 127.0.0.1
```

### 远程登录和复制文件

```shell
# ssh 用户名@ip 远程登录
ssh root@usernametang.com

# 远程复制文件,scp 用户名@ip:文件名或路径 用户名@ip:文件名或路径
scp /dir root@usernametang.com:/dir
```

### ssh 基础

- 利用 `SSH 协议` 可以有效**防止远程管理过程中的信息泄露** 
- 通过 `SSH 协议` 可以对所有传输的数据进行加密，也能够防止 DNS 欺骗和 IP 欺骗
- `SSH` 的另一项优点是传输的数据可以是经过压缩的，所以可以加快传输的速度

- **SSH 服务器** 的默认端口号是 `22`，默认时可以省略

- 常见服务端口号列表：

| 序号 | 服务       | 端口号 |
| :--- | :--------- | :----- |
| 01   | SSH 服务器 | 22     |
| 02   | Web 服务器 | 80     |
| 03   | HTTPS      | 443    |
| 04   | FTP 服务器 | 21     |

### SSH 客户端的简单使用

```bash
# p参数 port端口 uesr用户名 remote远程ip地址
ssh [-p port] user@remote
# 默认端口为22时，可以省略
ssh -p 22 root@usernametang.com

# 使用exit退出当前用户的登录
exit

# 如果在 `Windows` 系统中，可以安装 `PuTTY` 或者 `XShell` 客户端软件即可
```

### scp

参数P需要大写

```bash
# : 后面的路径如果不是绝对路径，则以用户的家目录作为参照路径
scp -P port 01.py user@remote:Desktop/01.py

# 把远程 家目录下的 Desktop/01.py 文件 复制到 本地当前目录下的 01.py
scp -P port user@remote:Desktop/01.py 01.py

# 加上 -r 选项可以传送文件夹
# 把当前目录下的 demo 文件夹 复制到 远程 家目录下的 Desktop
scp -r demo user@remote:Desktop

# 把远程 家目录下的 Desktop 复制到 当前目录下的 demo 文件夹
scp -r user@remote:Desktop demo

```

| 选项 | 含义                                                         |
| :--- | :----------------------------------------------------------- |
| -r   | 若给出的源文件是目录文件，则 scp 将递归复制该目录下的所有子目录和文件，目标文件必须为一个目录名 |
| -P   | 若远程 SSH 服务器的端口不是 22，需要使用大写字母 -P 选项指定端口 |

> 注意：
>
> - `scp` 这个终端命令只能在 `Linux` 或者 `UNIX` 系统下使用
> - 如果在 `Windows` 系统中，可以安装 `PuTTY`，使用 `pscp` 命令行工具或者安装 `FileZilla` 使用 `FTP` 进行文件传输

### SSH 高级（知道）

有关 SSH 配置信息都保存在用户家目录下的 `.ssh` 目录下

#### 免密登录

```shell
# 配置公钥，一直回车
ssh-keygen

# 上传公钥到服务器，让远程服务器记住公钥
ssh-copy-id -p port user@remote


```

#### 配置别名

而 **配置别名** 可以让我们进一步偷懒，譬如用：`ssh mac` 来替代上面这么一长串，那么就在 `~/.ssh/config` 里面追加以下内容：

```
Host ubuntu
    HostName ip地址
    User ubuntu
    Port 22

```

**保存之后，即可用 ssh mac 实现远程登录了，scp 同样可以使用**

## 权限与用户组

`ls -l` 可以查看文件夹下文件的详细信息，从左到右依次是：

- **权限**，第 1 个字符如果是 `d` 表示目录
- **硬链接数**，通俗地讲，就是有多少种方式，可以访问到当前目录／文件
- **拥有者**，家目录下 文件／目录 的拥有者通常都是当前用户
- **组**，在 Linux 中，很多时候，会出现组名和用户名相同的情况，后续会讲
- **大小 ** 
- **时间** 
- **名称** 

```shell
ubuntu@ubuntu:~$ ls -l
总用量 57220
drwxr-xr-x  3 ubuntu ubuntu     4096 May 16 12:37 Desktop
drwxr-xr-x  3 ubuntu ubuntu     4096 May 16 06:47 Documents
drwxr-xr-x  2 ubuntu ubuntu     4096 May 15 13:09 Downloads
-rw-r--r--  1 ubuntu ubuntu     8980 Apr 30 02:25 examples.desktop
......

```

### 修改文件权限

| 序号 | 命令  | 作用       |
| :--- | :---- | :--------- |
| 01   | chown | 修改拥有者 |
| 02   | chgrp | 修改组     |
| 03   | chmod | 修改权限   |

- 命令格式如下：

```bash
# 修改文件|目录的拥有者
chown 用户名 文件名|目录名

# 递归修改文件|目录的组
chgrp -R 组名 文件名|目录名

# 递归修改文件权限
chmod -R 755 文件名|目录名

```

- `chmod` 在设置权限时，可以简单地使用三个数字分别对应 **拥有者** ／ **组** 和 **其他** 用户的权限

```bash
# 直接修改文件|目录的 读|写|执行 权限，但是不能精确到 拥有者|组|其他
chmod +/-rwx 文件名|目录名

```

- 常见数字组合有（u表示用户／g示组／o表示其他）：
  - `777` ===> `u=rwx,g=rwx,o=rwx`
  - `755` ===> `u=rwx,g=rx,o=rx`
  - `644` ===> `u=rw,g=r,o=r`

| 序号 | 权限 |  英文  | 缩写 | 数字代号 |
| :--: | :--: | :----: | :--: | :------: |
|  01  |  读  |  read  |  r   |    4     |
|  02  |  写  | write  |  w   |    2     |
|  03  | 执行 | excute |  x   |    1     |

### sudo

```shell
`su不接用户名，可以切换到root，但是不推荐使用，因为不安全
# 切换到用户 user，执行时需要输入目标用户的密码
su <user>

# 以特权级别运行 cmd 命令，需要当前用户属于sudo组，输入当前用户的密码。
sudo <cmd>

# 切换用户，同时切换到用户家目录。
su - <user>

# 退出当前登录用户
exit

# 删除用户
sudo deluser user_name --remove-home
```

### 组管理

创建与删除组都需要通过sudo执行，组信息保存在 `/etc/group` 文件中，`/etc` 目录是专门用来保存 **系统配置信息** 的目录。在实际应用中，可以预先针对 **组** 设置好权限，然后 **将不同的用户添加到对应的组中**，从而**不用依次为每一个用户设置权限** 

```shell
# 添加组
groupadd 组名

# 删除组
groupdel 组名

# 确认组信息
cat /etc/group

# 递归修改文件/目录的所属组
chgrp -R 组名 文件/目录名
```

### 用户管理

创建、删除和修改其他用户密码的终端命令都需要通过 `sudo` 执行。

```shell
# 添加新用户 m自动建立用户家目录 g用户所在的组，否则建立一个同名的组
adduser -m -g 组 新建用户名

# 设置用户密码，如果是普通用户，直接用passwd可以修改自己的账户密码
passwd 用户名

# 删除用户，-r 选项会自动删除用户家目录
userdel -r 用户名

# 确认用户信息，新建用户后，用户信息会保存在 /etc/passwd 文件中
cat /etc/passwd | grep 用户名

# 查看用户 UID 和 GID 信息
id [用户名]

# 查看当前所有登录的用户列表
who

# 查看当前登录用户的账户名
whoami
```

### passwd 文件

`/etc/passwd` 文件存放的是用户的信息，由 6 个分号组成的 7 个信息，分别是

- 用户名
- 密码（x，表示加密的密码）
- UID（用户标识）
- GID（组标识）
- 用户全名或本地帐号
- 家目录
- 登录使用的 Shell，就是登录之后，使用的终端命令，`ubuntu` 默认是 `bash` 

### usermod

- `usermod` 可以用来设置 **用户** 的 **主组** ／ **附加组** 和 **登录 Shell**，命令格式如下：
- **主组**：通常在新建用户时指定，在 `etc/passwd` 的第 4 列 **GID 对应的组** 
- **附加组**：在 `etc/group` 中最后一列表示该组的用户列表，用于指定 **用户的附加权限**

> 提示：设置了用户的附加组之后，需要重新登录才能生效！

```bash
# 修改用户的主组（passwd 中的 GID）
usermod -g 组 用户名

# 修改用户的附加组
usermod -G 组 用户名

# 修改用户登录 Shell
usermod -s /bin/bash 用户名

```

> 注意：默认使用 `useradd` 添加的用户是没有权限使用 `sudo` 以 `root` 身份执行命令的，可以使用以下命令，将用户添加到 `sudo` 附加组中

```bash
usermod -G sudo 用户名

```

### which（重要）

> - `/etc/passwd` 是用于保存用户信息的文件
> - `/usr/bin/passwd` 是用于修改用户密码的程序

- `which` 命令可以查看执行命令所在位置，例如：

```bash
which ls

# 输出
# /bin/ls

which useradd

# 输出
# /usr/sbin/useradd

```

### bin 和 sbin

- 在 `Linux` 中，绝大多数可执行文件都是保存在 `/bin`、`/sbin`、`/usr/bin`、`/usr/sbin`
- `/bin`（`binary`）是二进制执行文件目录，主要用于具体应用
- `/sbin`（`system binary`）是系统管理员专用的二进制代码存放目录，主要用于系统管理
- `/usr/bin`（`user commands for applications`）后期安装的一些软件
- `/usr/sbin`（`super user commands for applications`）超级用户的一些管理程序

> 提示：
>
> - `cd` 这个终端命令是内置在系统内核中的，没有独立的文件，因此用 `which` 无法找到 `cd` 命令的位置





## 打包压缩

### 打包与解包

```bash
# 打包文件，tar只负责打包或解包，不负责压缩或解压缩
tar -cvf 打包文件.tar 被打包的文件／路径...

# 解包文件
tar -xvf 打包文件.tar
```

| 选项 | 含义                                                         |
| :--- | :----------------------------------------------------------- |
| c    | 生成档案文件，创建打包文件                                   |
| x    | 解开档案文件                                                 |
| v    | 列出归档解档的详细过程，显示进度                             |
| f    | 指定档案文件名称，f 后面一定是 .tar 文件，所以必须放选项最后 |

> 注意：`f` 选项必须放在最后，其他选项顺序可以随意

### 压缩与解压缩

```bash
# 压缩文件 -z调用gzip，进行压缩解压缩
tar -zcvf 打包文件.tar.gz 被压缩的文件／路径...

# 解压缩文件
tar -zxvf 打包文件.tar.gz

# 解压缩到指定路径。-C 解压缩到指定目录，注意：要解压缩的目录必须存在
tar -zxvf 打包文件.tar.gz -C 目标路径

```

### bzip2(two)

```bash
# 压缩文件
tar -jcvf 打包文件.tar.bz2 被压缩的文件／路径...

# 解压缩文件
tar -jxvf 打包文件.tar.bz2
```

## vim与vi

### 对比

- vi：

  只能是编辑文本内容，不能对字体、段落进行排版，不支持鼠标操作，没有菜单只有命令，在 系统管理、服务器管理，编辑文件时，其功能永远不是图形界面的编辑器能比拟的。

- vim(vi improved)

  vim是从vi发展出来的一个文本编辑器，支持 代码补全、编译及错误跳转等方便编程的功能特别丰富，在程序员中被广泛使用，被称为 编辑器之神.

### 打开和新建文件

- 在终端中输入 `vim` **在后面跟上文件名** 即可

```bash
# 打开或新建文件，不存在时新建，存在时打开。
vim 文件名

# 打开文件并且定位行，如果只带上+而不指定行号，会直接定位到文件末尾
vim 文件名 +行数

# 打开practice_1.txt文档
vim practice_1.txt

# 直接使用vim也可以打开vim编辑器，但是不会打开任何文件
vim

# 进入命令行模式后输入:e 文件路径同样可以打开相应文件
:e 文件路径
```

异常处理

- 如果 `vim` 异常退出，在磁盘上可能会保存有 **交换文件** 
- 下次再使用 `vim` 编辑该文件时，会看到以下屏幕信息，按下字母 `d` 可以 **删除交换文件** 即可

### 命令行模式

#### 末行模式命令

| 命令 |     英文     | 功能                           |
| :--: | :----------: | ------------------------------ |
|  w   |    write     | 保存                           |
|  q   |     quit     | 退出，如果没有保存，不允许退出 |
|  q!  |     quit     | 强行退出，不保存退出           |
|  wq  | write & quit | 保存并退出                     |
|  x   |              | 保存并退出                     |

#### 游标移动（基本）

> 学习提示

1. `vi` 的命令较多，**不要期望一下子全部记住**，个别命令忘记了，只是会影响编辑速度而已
2. 在使用 `vim` 命令时，注意 **关闭中文输入法**



要熟练使用 `vi`，首先应该学会怎么在 **命令模式** 下样快速移动光标

**编辑操作命令**，能够和 **移动命令** 结合在一起使用



1) 上、下、左、右

| 命令 | 功能 |  手指  |
| :--: | ---- | :----: |
|  h   | 向左 |  食指  |
|  j   | 向下 |  食指  |
|  k   | 向上 |  中指  |
|  l   | 向右 | 无名指 |

![005_移动光标-w551](.\assets\005_移动光标.png)

2) 行内移动（了解）

| 命令 | 英文 | 功能                           |
| :--: | :--: | ------------------------------ |
|  w   | word | 向后移动一个单词               |
|  b   | back | 向前移动一个单词               |
|  0   |      | 行首                           |
|  ^   |      | 行首，第一个不是空白字符的位置 |
|  $   |      | 行尾                           |

3) 行数移动（了解）

|  命令  | 英文 | 功能                 |
| :----: | :--: | -------------------- |
|   gg   |  go  | 文件顶部             |
|   G    |  go  | 文件末尾             |
| 数字gg |  go  | 移动到 数字 对应行数 |
| 数字G  |  go  | 移动到 数字 对应行数 |
| :数字  |      | 移动到 数字 对应行数 |

4) 屏幕移动（了解）

|   命令   |  英文   | 功能     |
| :------: | :-----: | -------- |
| Ctrl + b |  back   | 向上翻页 |
| Ctrl + f | forward | 向下翻页 |
|    H     |  Head   | 屏幕顶部 |
|    M     | Middle  | 屏幕中间 |
|    L     |   Low   | 屏幕底部 |



#### 进入插入模式（可视模式）

- 学习 `复制` 命令前，应该先学会 **怎么样选中 要复制的代码**
- 在 `vi` 中要选择文本，需要先使用 `Visual` 命令切换到 **可视模式**
- `vi` 中提供了 **三种** 可视模式，可以方便程序员选择 **选中文本的方式**
- 按 `ESC` 可以放弃选中，返回到 **命令模式**

|   命令   | 模式       | 功能                               |
| :------: | ---------- | ---------------------------------- |
|    v     | 可视模式   | 从光标位置开始按照正常模式选择文本 |
|    V     | 可视行模式 | 选中光标经过的完整行               |
| Ctrl + v | 可视块模式 | 垂直方向选中文本                   |

**可视模式**下，可以和 **移动命令** 连用，例如：`ggVG` 能够选中所有内容



#### 插入命令

在 `vim` 中除了常用的 `i` 进入 **编辑模式** 外，还提供了以下命令同样可以进入编辑模式：

| 命令 |  英文  | 功能                   |  常用  |
| :--: | :----: | ---------------------- | :----: |
|  i   | insert | 在当前字符前插入文本   |  常用  |
|  I   | insert | 在行首插入文本         | 较常用 |
|  a   | append | 在当前字符后添加文本   |        |
|  A   | append | 在行末添加文本         | 较常用 |
|  o   |        | 在当前行后面插入一空行 |  常用  |
|  O   |        | 在当前行前面插入一空行 |  常用  |

![004_插入命令-w400](.\assets\004_插入命令.png)



#### 撤销和恢复撤销

在学习编辑命令之前，先要知道怎样撤销之前一次 **错误的** 编辑动作！

|   命令   | 英文 | 功能           |
| :------: | :--: | -------------- |
|    u     | undo | 撤销上次命令   |
| CTRL + r | redo | 恢复撤销的命令 |



#### 删除文本(命令行模式)

|    命令     |  英文  | 功能                              |
| :---------: | :----: | --------------------------------- |
|      x      |  cut   | 删除光标所在字符，或者选中文字    |
| d(移动命令) | delete | 删除移动命令对应的内容            |
|     dd      | delete | 删除光标所在行，可以 ndd 复制多行 |
|      D      | delete | 删除至行尾                        |

> 提示：如果使用 **可视模式** 已经选中了一段文本，那么无论使用 `d` 还是 `x`，都可以删除选中文本

- 删除命令可以和 **移动命令** 连用，以下是常见的组合命令：

```
* dw        # 从光标位置删除到单词末尾
* d0        # 从光标位置删除到一行的起始位置
* d}        # 从光标位置删除到段落结尾
* ndd       # 从光标位置向下连续删除 n 行
* d代码行G   # 从光标所在行 删除到 指定代码行 之间的所有代码
* d'a       # 从光标所在行 删除到 标记a 之间的所有代码

```



#### 复制、粘贴

`vi` 中提供有一个 **被复制文本的缓冲区**

- **复制** 命令会将选中的文字保存在缓冲区 
- **删除** 命令删除的文字会被保存在缓冲区
- 在需要的位置，使用 **粘贴** 命令可以将缓冲区的文字插入到光标所在位置

|    命令     | 英文  | 功能                        |
| :---------: | :---: | --------------------------- |
| y(移动命令) | copy  | 复制                        |
|     yy      | copy  | 复制一行，可以 nyy 复制多行 |
|      p      | paste | 粘贴                        |

**提示**

- 命令 `d`、`x` 类似于图形界面的 **剪切操作** —— `CTRL + X`
- 命令 `y` 类似于图形界面的 **复制操作** —— `CTRL + C`
- 命令 `p` 类似于图形界面的 **粘贴操作** —— `CTRL + V`
- `vi` 中的 **文本缓冲区同样只有一个**，如果后续做过 **复制、剪切** 操作，之前缓冲区中的内容会被替换

**注意**

- `vi` 中的 **文本缓冲区** 和系统的 **剪贴板** 不是同一个
- 所以在其他软件中使用 `CTRL + C` 复制的内容，不能在 `vi` 中通过 `P` 命令粘贴
- 可以在 **编辑模式** 下使用 **鼠标右键粘贴**



#### 退出vim

从普通模式输入`:`进入命令行模式，输入`wq`回车，保存并退出编辑

以下为其它几种退出方式：

| 命令               | 说明             |
| ------------------ | ---------------- |
| `:q!`              | 强制退出，不保存 |
| `:q`               | 退出             |
| `:wq!`             | 强制保存并退出   |
| `:w <文件路径>`    | 另存为           |
| `:saveas 文件路径` | 另存为           |
| `:x`               | 保存并退出       |
| `:wq`              | 保存并退出       |

#### 替换

| 命令 |  英文   | 功能                   | 工作模式 |
| :--: | :-----: | ---------------------- | -------- |
|  r   | replace | 替换当前字符           | 命令模式 |
|  R   | replace | 替换当前行光标后的字符 | 替换模式 |

- `R` 命令可以进入 **替换模式**，替换完成后，按下 `ESC` 可以回到 **命令模式**
- **替换命令** 的作用就是不用进入 **编辑模式**，对文件进行 **轻量级的修改**

#### 缩排和重复执行

| 命令 | 功能         |
| :--: | ------------ |
|  >>  | 向右增加缩进 |
|  <<  | 向左减少缩进 |
|  .   | 重复上次命令 |

- **缩排命令** 在开发程序时，**统一增加代码的缩进** 比较有用！
  - 一次性 **在选中代码前增加 4 个空格**，就叫做 **增加缩进**
  - 一次性 **在选中代码前删除 4 个空格**，就叫做 **减少缩进**
- 在 **可视模式** 下，缩排命令只需要使用 **一个** `>` 或者 `<` 

> 在程序中，**缩进** 通常用来表示代码的归属关系
>
> 前面空格越少，代码的级别越高
>
> 前面空格越多，代码的级别越低



### 查找

#### 常规查找 

| 命令 | 功能     |
| :--: | -------- |
| /str | 查找 str |

- 查找到指定内容之后，使用 `Next` 查找下一个出现的位置：
  - `n`: 查找下一个
  - `N`: 查找上一个
- 如果不想看到高亮显示，可以随便查找一个文件中不存在的内容即可

#### 单词快速匹配

| 命令 | 功能                     |
| :--: | ------------------------ |
|  *   | 向后查找当前光标所在单词 |
|  #   | 向前查找当前光标所在单词 |

- 在开发中，通过单词快速匹配，可以快速看到这个单词在其他什么位置使用过

#### 查找并替换

- 在 `vim` 中查找和替换命令需要在 **末行模式** 下执行
- 记忆命令格式：

```
:%s///g

```

1) 全局替换

- **一次性**替换文件中的 **所有出现的旧文本**
- 命令格式如下：

```
:%s/旧文本/新文本/g

```

2) 可视区域替换

- **先选中** 要替换文字的 **范围**
- 命令格式如下：

```
:s/旧文本/新文本/g

```

3) 确认替换

- 如果把末尾的 `g` 改成 `gc` 在替换的时候，会有提示！**推荐使用！**

```
:%s/旧文本/新文本/gc

```

1. `y` - `yes` 替换
2. `n` - `no` 不替换
3. `a` - `all` 替换所有
4. `q` - `quit` 退出替换
5. `l` - `last` 最后一个，并把光标移动到行首
6. `^E` 向下滚屏
7. `^Y` 向上滚屏



### 演练

#### 演练 1 —— 编辑命令和数字连用

在开发中，可能会遇到连续输入 `N` 个同样的字符

> 在 `Python` 中有简单的方法，但是其他语言中通常需要自己输入

- 例如：`**********` 连续 10 个星号

要实现这个效果可以在 **命令模式** 下

1. 输入 `10`，表示要重复 10 次
2. 输入 `i` 进入 **编辑模式**
3. 输入 `*` 也就是重复的文字
4. 按下 `ESC` 返回到 **命令模式**，返回之后 `vim` 就会把第 `2、3` 两步的操作重复 `10` 次

> 提示：正常开发时，在 **进入编辑模式之前，不要按数字**

#### 演练 2 —— 利用 可视块 给多行代码增加注释

- 在开发中，可能会遇到一次性给多行代码 **增加注释** 的情况

> 在 `Python` 中，要给代码增加注释，可以在代码前增加一个 `# `

要实现这个效果可以在 **命令模式** 下

1. 移动到要添加注释的 **第 1 行代码**，按 `^` 或者 `gg`来到行首
2. 按 `CTRL + v` 进入 **可视块** 模式
3. 使用 `j` 向下连续选中要添加的代码行
4. 输入 `I` 进入 **编辑模式**，并在 **行首插入**，注意：一定要使用 **I**
5. 输入 `# ` 也就是注释符号
6. 按下 `ESC` 返回到 **命令模式**，返回之后 `vi` 会在之前选中的每一行代码 **前** 插入 `# `



### 末行命令扩展

**末行命令** 主要是针对文件进行操作的：**保存**、**退出**、**保存&退出**、**搜索&替换**、**另存**、**新建**、**浏览文件**

|   命令    | 英文  | 功能                                           |
| :-------: | :---: | ---------------------------------------------- |
|   :e .    | edit  | 会打开内置的文件浏览器，浏览要当前目录下的文件 |
| :n 文件名 |  new  | 新建文件                                       |
| :w 文件名 | write | 另存为，但是仍然编辑当前文件，并不会切换文件   |

> 提示：切换文件之前，必须保证当前这个文件已经被保存！

- 已经学习过的 **末行命令**：

|   命令   |     英文     | 功能                           |
| :------: | :----------: | ------------------------------ |
|    :w    |    write     | 保存                           |
|    :q    |     quit     | 退出，如果没有保存，不允许退出 |
|   :q!    |     quit     | 强行退出，不保存退出           |
|   :wq    | write & quit | 保存并退出                     |
|    :x    |              | 保存并退出                     |
| :%s///gc |              | 确认搜索并替换                 |

> 在实际开发中，可以使用 `w` 命令 **阶段性的备份代码**



## 常用命令速查图

![vi](.\assets\vim.png)

### vimrc

`vimrc` 是 `vim` 的配置文件，可以设置 vim 的配置，包括：**热键**、**配色**、**语法高亮**、**插件** 等

`Linux` 中 `vimrc` 有两个位置，**家目录下的配置文件优先级更高** 

```
/etc/vim/vimrc
~/.vimrc

```

常用的插件有：

- 代码补全
- 代码折叠
- 搜索
- Git 集成
- ……

网上有很多高手已经配置好的针对 `python` 开发的 `vimrc` 文件，可以下载过来直接使用，或者等大家多 `Linux` 比较熟悉后，再行学习！



## 附录：HOW

> 官网地址：<https://github.com/chenjiandongx/how>

### 💡 IDEA

Linux 是每位开发者必备的技能，如何高效地掌握 Linux 命令就成为一件很重要的事了。[jaywcjlove/linux-command](https://github.com/jaywcjlove/linux-command) 项目收集和整理了 500+ 的 Linux 命令使用文档，不过缺少了一个命令行版本，`how` 决定来填补这个空缺。

### 🔰 安装

**pip 安装**

```
$ pip install how

```

**源码安装**

```
$ git clone https://github.com/chenjiandongx/how.git
$ cd how
$ pip install -r requirements.txt
$ python setup.py install

```

### 📏 使用

```
$ how
usage: how [-h] [-i] [-v] [COMMAND [COMMAND ...]]

Lovely Linux commands cheat sheet.

positional arguments:
  COMMAND        the puzzling command

optional arguments:
  -h, --help     show this help message and exit
  -i, --init     initialize all commands
  -v, --version  displays the current version of `how`

```

> Note: 建议第一次使用 `how` 时先初始化所有的命令文档，`how -i`，该命令会将 <https://github.com/jaywcjlove/linux-command> 的 .md 文档下载到 `~/.command` 本地路径下。不过这个操作不是必须的，因为如果 `how some-command` 在本地路径中查询不到的话，会尝试先向远程地址下载。

### 🔖 示例

初始化所有文档，同时也是更新所有文档的命令

```
$ how -i
Initializing commands: 96/562 

```

查询如何使用 `man` 命令

```
$ how man
# man

查看 Linux 中的指令帮助

##  补充说明

man 命令 是 Linux 下的帮助指令，通过 man 指令可以查看
Linux 中的指令帮助、配置文件帮助和编程帮助等信息。

###  语法

man(选项)(参数)

###  选项

-a：在所有的 man 帮助手册中搜索；
-f：等价于 whatis 指令，显示给定关键字的简短描述信息；
-P：指定内容时使用分页程序；
-M：指定 man 手册搜索的路径。

###  参数

- 数字：指定从哪本 man 手册中搜索帮助；
- 关键字：指定要搜索帮助的关键字。

###  数字代表内容

1：用户在 shell 环境可操作的命令或执行文件；
2：系统内核可调用的函数与工具等
3：一些常用的函数(function) 与函数库(library)，大部分为 C
的函数库(libc)
4：设备文件说明，通常在/dev 下的文件
5：配置文件或某些文件格式
6：游戏(games)
7：惯例与协议等，如 Linux 文件系统，网络协议，ASCII code
等说明
8：系统管理员可用的管理命令
9：跟 kernel 有关的文件

###  实例

我们输入 man ls，它会在最左上角显示“LS（1）”，在这里，“LS”表示手
册名称，而“（1）”表示该手册位于第一节章，同样，我们输 man
ifconfig
它会在最左上角显示“IFCONFIG（8）”。也可以这样输入命令：“man
[章节号] 手册名称”。

man 是按照手册的章节号的顺序进行搜索的，比如：

man sleep

只会显示 sleep 命令的手册,如果想查看库函数 sleep，就要输入:

man 3 sleep

```

## 在虚拟机中安装 Ubuntu

### 安装前的准备

- 访问 `https://www.ubuntu.com/download/desktop` 下载 `Ubuntu` 18.04 版本
- 在操作系统上安装 `VMWare` 虚拟机软件

> 为什么要使用虚拟机？
>
> 1. 不需要准备 额外的电脑
> 2. 在虚拟机中可以 **大胆尝试任何操作**，不用担心造成不可逆转的破坏

### 基本安装

- 只需要按照提示一步一步操作即可安装 `Ubuntu` 了
- 安装结束后，登录并且进入系统，可以看到以下画面：

![assets](.\assets\002_默认安装完成.png)

## 设置语言环境

按照以下步骤操作可以修改 ubuntu 的语言环境

- 1) 进入 **系统设置 System Settings**
- 2) 打开 **语言支持 Language Support**
- 3) 通过 **安装/删除语言 Install/Remove Languages** 添加 **简体中文 Chinese(simplified)**
- 4) 将 **汉语** 拖拽到 **顶部**，如下图所示：

![assets](.\assets\003_将汉语拖拽到顶部.png)

- 5) **重新启动系统**
- 6) **不** 修改文件夹的名称，如下图所示：

![assets](.\assets\004_不修改文件夹名称.png)

## 安装常用软件

1 设置服务器镜像源

2 在启动栏添加 **终端** 图标

3 **apt** 终端命令

4 谷歌浏览器 chrome

5 搜狗输入法

### 设置服务器镜像源

`Ubuntu` 中 **大部分** 的软件 **安装/更新** 都是利用 `apt` 命令，从 **ubuntu 的服务器** 直接安装的

`Ubuntu` **官方的服务器在国外**，为了提高软件 **安装/更新速度**，**ubuntu** 提供了 **选择最佳服务器** 的功能，可以帮助我们方便的找到一个速度最快的 **镜像服务器**！

> 所谓 **镜像服务器**，就是 **所有服务器的内容是相同的（镜像）**，但是根据所在位置不同，速度不同，通常国内服务器速度会更快一些！

按照以下步骤操作可以设置 `ubuntu` 的服务器

1) 进入 **系统设置** 

2) 打开 **软件和更新** 

3) 设置 **下载自...** **其他站点** 

4) 通过 **选择最佳服务器** 选择速度最快的 **镜像源**，如下图所示：

![assets](.\assets\001_ubuntu设置软件源.png)

**提示**：更换服务器之后，需要一个相对比较长时间的更新过程，需要耐心等待。更新完成后，再安装软件都会从新设置的服务器下载软件了

### 在启动栏添加 **终端** 图标

`Ubuntu` 的 **启动栏** 类似于 Windows 的 **任务栏**，可以显示 **常用软件图标** 和 **正在运行的软件图标**

通常 Linux 的用户使用 **终端命令** 的频率非常高，所以会把 **终端** 添加到 **启动栏**

通过最上方的 **搜索** 按钮可以 **搜索并启动** 需要使用的程序

**点击右键** 可以 **保留**、**删除** 或者 **移动** 启动栏中图标

可以把最常用的软件图标，**保留** 在启动栏，可以方便快速启动程序



# shell

## shell简介

### shell

shell是一个命令行解释器，它接收应用程序/用户命令，然后调用操作系统内核。shell还是一个功能相当强大的变成语言，易编写，易调试，灵活性强。

**查看shell的命令**

```bash
# 查看linux提供的shell
cat /etc/shells

# 查看bash与sh的关系
ll | grep bash

# 查看默认解析器
echo $SHELL
```

### 脚本格式

脚本以`\#!/bin/bash`开头（指定解析器）

脚本执行方式：

```bash
# 1. bash 或 sh+脚本的相对路径或绝对路径（不用赋予脚本+x 权限）
 sh ./helloworld.sh # sh +相对路径
 bash ./helloworld.sh # 相对路径
 
# 2. 输入脚本的绝对路径或相对路径执行脚本（必须具有可执行权限+x）
chmod +x helloworld.sh # 赋予脚本执行权限
./helloworld.sh 
/helloworld.sh 

# 3. 在脚本的路径前加上“.”或者 source
. helloworld.sh
source helloworld.sh
```

**区别** ：

`sh`、`bash`、`./`、`/`方式是在当前shell中打开一个子shell来执行脚本内容，当前脚本内容结束，则关闭子shell，回到父shell中。

`.`,`source`的方式，可以使脚本在当前shell里执行，而无需打开子shell！这也是为什么我们每次要修改完/etc/profile 文件以后，需要 source 一下的原因。

开子 shell 与不开子 shell 的区别就在于，环境变量的继承关系，如在子 shell 中设置的当前变量，父 shell 是不可见的。

## 变量

### 系统预定义变量

| 变量   | 说明         |
| ------ | ------------ |
| $HOME  | 家目录       |
| $PWD   | 当前工作目录 |
| $SHELL | 当前的shell  |
| $USER  | 当前的用户   |

```bash
# 显示当前 Shell 中所有变量
set
```

### 自定义变量

**语法：**

- 定义变量：变量名=变量值，=号前后不能用空格
- 撤销变量：unset 变量名
- 声明静态变量：readonly 变量，不能 unset

**定义规则：**

- 变量名称可以由字母、数字和下划线组成，但是不能以数字开头，**环境变量名建议大写**
- 等号两侧不能有空格
- 在 bash 中，变量默认类型都是字符串类型，无法直接进行数值运算。
- 变量的值如果有空格，需要使用双引号或单引号括起来。

```bash
# 定义变量A
[ubuntu@VM-12-11 shells]$ A=5
[ubuntu@VM-12-11 shells]$ echo $A
5
# 给变量A重新赋值
[ubuntu@VM-12-11 shells]$ A=8
[ubuntu@VM-12-11 shells]$ echo $A
8
# 撤销变量A
[ubuntu@VM-12-11 shells]$ unset A
[ubuntu@VM-12-11 shells]$ echo $A
# 声明静态的变量 B=2，不能 unset
[ubuntu@VM-12-11 shells]$ readonly B=2
[ubuntu@VM-12-11 shells]$ echo $B
2
[ubuntu@VM-12-11 shells]$ B=9
-bash: B: readonly variable

# 在 bash 中，变量默认类型都是字符串类型，无法直接进行数值运算
[atguigu@hadoop102 ~]$ C=1+2
[atguigu@hadoop102 ~]$ echo $C
1+2

# 变量的值如果有空格，需要使用双引号或单引号括起来
[atguigu@hadoop102 ~]$ D=I love xianbaobao
-bash: world: command not found
[atguigu@hadoop102 ~]$ D="I love xianbaobao"
[atguigu@hadoop102 ~]$ echo $D
I love xianbaobao
# 可把变量提升为全局环境变量，可供其他 Shell 程序使用
export 变量名
[ubuntu@VM-12-11 shells]$ vim helloworld.sh

# 在 helloworld.sh 文件中增加 echo $B
#!/bin/bash
echo "helloworld"
echo $B
[ubuntu@VM-12-11 shells]$ export B
[ubuntu@VM-12-11 shells]$ ./helloworld.sh
helloworld
2
```

## 特殊变量

| 变量 | 说明                                                         |
| ---- | ------------------------------------------------------------ |
| $n   | n 为数字，`$0` 代表该脚本名称，$1-$9 代表第一到第九个参数，十以上的参数，十以上的参数需要用大括号包含，如${10} |
| $#   | 获取所有输入参数个数，常用于循环,判断参数的个数是否正确以及加强脚本的健壮性 |
| $*   | 代表命令行中所有的参数，$*把所有的参数看成一个整体(类似于解包) |
| $@   | 代表命令行中所有的参数，不过$@把每个参数区分对待             |
| $？  | 最后一次执行的命令的返回状态。如果这个变量的值为 0，证明上一个命令正确执行；如果这个变量的值为非 0（具体是哪个数，由命令自己来决定），则证明上一个命令执行不正确了。） |

```bash
# $n
[root@vultr]$ touch parameter.sh
[root@vultr]$ vim parameter.sh
#!/bin/bash
echo '==========$n=========='
echo $0
echo $1
echo $2
[root@vultr]$ chmod 777 parameter.sh
[root@vultr]$ ./parameter.sh cls xz
==========$n==========
./parameter.sh
cls
xz

# $#
[root@vultr]$ vim parameter.sh
#!/bin/bash
echo '==========$n=========='
echo $0
echo $1
echo $2
echo '==========$#=========='
echo $#
[root@vultr]$ chmod 777 parameter.sh
[root@vultr]$ ./parameter.sh cls xz
==========$n==========
./parameter.sh
cls
xz
==========$#==========
2

# $*、$@
[root@vultr]$ vim parameter.sh
#!/bin/bash
echo '==========$n=========='
echo $0
echo $1
echo $2
echo '==========$#=========='
echo $#
echo '==========$*=========='
echo $*
echo '==========$@=========='
echo $@
[root@vultr]$ ./parameter.sh a b c d e f g
==========$n==========
./parameter.sh
a
b
==========$#==========
7
==========$*==========
a b c d e f g
==========$@==========
a b c d e f g

# $?
[root@vultr]$ ./helloworld.sh
hello world
[root@vultr]$ echo $?
0
```

## 运算符

语法：`“$((运算式))” 或 “$[运算式]”`

```bash
# 计算2+3*4的值
[root@vultr shells]# S=$[(2+3)*4]
[root@vultr shells]# echo $S
```

## 条件判断

语法：

- test condition
- [ condition ]（注意 condition 前后要有空格）

注意：条件非空即为 true，[ atguigu ]返回 true，[ ] 返回 false。

**常用判断条件：**

**两个数之间的比较**

| 符号 | 说明                      |
| ---- | ------------------------- |
| -eq  | 等于（equal）             |
| -ne  | 不等于（not equal）       |
| -lt  | 小于（less than）         |
| -le  | 小于等于（less equal）    |
| -gt  | 大于（greater than）      |
| -ge  | 大于等于（greater equal） |

**注：如果是字符串之间的比较 ，用等号“=”判断相等；用“!=”判断不等。**

**按照文件权限进行判断**

| 符号 | 说明                    |
| ---- | ----------------------- |
| -r   | 有读的权限（read）      |
| -w   | 有写的权限（write）     |
| -x   | 有执行的权限（execute） |

**按照文件类型进行判断**

| 符号 | 说明                                 |
| ---- | ------------------------------------ |
| -e   | e 文件存在（existence）              |
| -f   | 文件存在并且是一个常规的文件（file） |
| -d   | 文件存在并且是一个目录（directory）  |

```bash
# 23 是否大于等于 22
[root@vultr shells]$ [ 23 -ge 22 ]
[root@vultr shells]$ echo $?
0

# helloworld.sh 是否具有写权限
[root@vultr shells]$ [ -w helloworld.sh ]
[root@vultr shells]$ echo $?
0

#/home/cls.txt 目录中的文件是否存在
[root@vultr shells]$ [ -e /home/cls.txt ]
[root@vultr shells]$ echo $?
1

# 多条件判断（&& 表示前一条命令执行成功时，才执行后一条命令，|| 表示上一条命令执行失败后，才执行下一条命令）
[root@vultr ~]$ [ atguigu ] && echo OK || echo notOK
OK
[root@vultr shells]$ [ ] && echo OK || echo notOK
notOK
```

## 流程控制

### if

```bash
# 单分支
if [ 条件判断式 ];then
	程序
fi

# 单分支
if [ 条件判断式 ]
then
	程序
fi

# 多分枝
if [ 条件判断式 ]
then
	程序
elif [ 条件判断式 ]
then
	程序
else
	程序
fi

```

注意事项：

- [ 条件判断式 ]，中括号和条件判断式之间必须有空格
- if 后要有空格

```bash
# 输入一个数字，如果是 1，则输出 xianbaobao zhen haokan，如果是 2，则输出 xianbaobao zhen mei，如果是其它，什么也不输出。
[root@vultr shells]$ touch if.sh
[root@vultr shells]$ vim if.sh
#!/bin/bash
if [ $1 -eq 1 ]
then
	echo "banzhang zhen shuai"
elif [ $1 -eq 2 ]
then
	echo "cls zhen mei"
fi
[root@vultr shells]$ chmod 777 if.sh
[root@vultr shells]$ ./if.sh 1
xianbaobao zhen haokan
```

### case语句

```bash
case $变量名 in
"值 1"）
	如果变量的值等于值 1，则执行程序 1
;;
"值 2"）
	如果变量的值等于值 2，则执行程序 2
;;
	…省略其他分支…
*）
	如果变量的值都不是以上的值，则执行此程序
;;
esac
```

注意事项：

- case 行尾必须为单词“in”，每一个模式匹配必须以右括号“）”结束。
- 双分号“**;;**”表示命令序列结束，相当于 break。
- 最后的“*）”表示默认模式，相当于  default。

```bash
# 输入一个数字，如果是 1，则输出 xianbaobao，如果是 2，则输出 cls，如果是其它，输出yaya。
[root@vultr shells]$ touch case.sh
[root@vultr shells]$ vim case.sh

!/bin/bash
	case $1 in
"1")
	echo "xianbaobao"
;;
"2")
	echo "cls"
;;
*)
	echo "yaya"
;;
esac
[root@vultr shells]$ chmod 777 case.sh
[root@vultr shells]$ ./case.sh 1
1
```

### for循环

```bash
# 方式一
for (( 初始值;循环控制条件;变量变化 ))
do
	程序
done

# 方式二
for 变量 in 值 1 值 2 值 3…
do
	程序
done
```

```shell
# 从 1 加到 100
[root@vultr shells]$ touch for1.sh
[root@vultr shells]$ vim for1.sh
#!/bin/bash
sum=0
for((i=0;i<=100;i++))
do
	sum=$[$sum+$i]
done
echo $sum
[root@vultr shells]$ chmod 777 for1.sh
[root@vultr shells]$ ./for1.sh
5050

# 打印所有输入参数
[root@vultr shells]$ touch for2.sh
[root@vultr shells]$ vim for2.sh
#!/bin/bash
#打印数字
for i in cls mly wls
do
	echo "xxxx love $i"
done
[root@vultr shells]$ chmod 777 for2.sh
[root@vultr shells]$ ./for2.sh
xxxx love cls
xxxx love mly
xxxx love wls

# 比较$*和$@的区别，$*和$@都表示传递给函数或脚本的所有参数，不被双引号“”包含时，都以$1 $2 …$n的形式输出所有参数。
[root@vultr shells]$ touch for3.sh
[root@vultr shells]$ vim for3.sh
#!/bin/bash
echo '=============$*============='
for i in $*
do
echo "ban zhang love $i"
done
echo '=============$@============='
for j in $@
do
echo "ban zhang love $j"
done
[root@vultr shells]$ chmod 777 for3.sh
[root@vultr shells]$ ./for3.sh cls mly wls
=============$*=============
banzhang love cls
banzhang love mly
banzhang love wls
=============$@=============
banzhang love cls
banzhang love mly
banzhang love wls

# 当它们被双引号“”包含时，$*会将所有的参数作为一个整体，以“$1 $2 …$n”的形式输出所有参数；$@会将各个参数分开，以“$1” “$2”…“$n”的形式输出所有参数。
[root@vultr shells]$ vim for4.sh
#!/bin/bash
echo '=============$*============='
for i in "$*"
#$*中的所有参数看成是一个整体，所以这个 for 循环只会循环一次
do
echo "ban zhang love $i"
done
echo '=============$@============='
for j in "$@"
#$@中的每个参数都看成是独立的，所以“$@”中有几个参数，就会循环几次
do
echo "ban zhang love $j"
done
[root@vultr shells]$ chmod 777 for4.sh
[root@vultr shells]$ ./for4.sh cls mly wls
=============$*=============
banzhang love cls mly wls # 类似于解包
=============$@=============
banzhang love cls
banzhang love mly
banzhang love wls
```

### while循环

```shell
while [ 条件判断式 ]
do
	程序
done
```

从 1 加到 100

```shell
[root@vultr shells]$ touch while.sh
[root@vultr shells]$ vim while.sh
#!/bin/bash
sum=0
i=1
while [ $i -le 100 ]
do
sum=$[$sum+$i]
i=$[$i+1]
done
echo $sum
[root@vultr shells]$ chmod 777 while.sh
[root@vultr shells]$ ./while.sh
5050
```

## read读取控制台输入

基本语法：read (选项) （参数）

- 选项：
  - -p：指定读取值时的提示符
  - -t：指定读取值时等待的时间(秒)，如果-t不加则表示一直等待
- 参数
  - 变量：指定读取值得变量名

```shell
# 提示 7 秒内，读取控制台输入的名称
[root@vultr shells]$ touch read.sh
[root@vultr shells]$ vim read.sh
#!/bin/bash
read -t 7 -p "Enter your name in 7 seconds :" NN
echo $NN
[root@vultr shells]$ ./read.sh
Enter your name in 7 seconds : hsiangya
hsiangya
```

## 函数

### dirname

语法：`dirname 文件绝对路径`

说明：从给定的包含绝对路径的文件名中去除文件名（非目录的部分），然后返回剩下的路径（目录的部分））dirname 可以理解为取文件路径的绝对路径名称

```shell
# 获取 hsiangya.txt 文件的路径。
[root@vultr ~]$ dirname /home/hsiangya.txt
/home
```

### basename

语法：`basename [string / pathname] [suffix]`

说明：basename 命令会删掉所有的前缀包括最后一个（‘/’）字符，然后将字符串显示出来。basename 可以理解为取路径里的文件名称

选项：suffix 为后缀，如果 suffix 被指定了，basename 会将 pathname 或 string 中的 suffix 去掉。

```shell
# 截取该/home/atguigu/banzhang.txt 路径的文件名称。
[root@vultr shells]$ basename /home/atguigu/banzhang.txt
banzhang.txt
[root@vultr shells]$ basename /home/atguigu/banzhang.txt .txt
banzhang
```

### 自定义函数

- 必须在调用函数地方之前，先声明函数，shell 脚本是逐行运行。不会像其它语言一样先编译。
- 函数返回值，只能通过$?系统变量获得，可以显示加：return 返回，如果不加，将以最后一条命令运行结果，作为返回值。return 后跟数值 n(0-255)

语法：

```shell
[ function ] funname[()]
{
	Action;
	[return int;]
}
```

```bash
# 计算两个输入参数的和。
[root@vultr shells]$ touch fun.sh
[root@vultr shells]$ vim fun.sh
#!/bin/bash
function sum()
{
	s=0
	s=$[$1+$2]
	echo "$s"
}
read -p "Please input the number1: " n1;
read -p "Please input the number2: " n2;
sum $n1 $n2;
[root@vultr shells]$ chmod 777 fun.sh
[root@vultr shells]$ ./fun.sh
Please input the number1: 2
Please input the number2: 5
7
```

## 文本处理工具

### cut

cut 的工作就是“剪”，具体的说就是在文件中负责剪切数据用的。cut 命令从文件的每一行剪切字节、字符和字段并将这些字节、字符和字段输出。

语法：`cut [选项参数] filename`

说明：默认分隔符是制表符

| 选项参数 | 说明                                           |
| -------- | ---------------------------------------------- |
| -f       | 列号，提取第几列                               |
| -d       | 分隔符，按照指定分隔符分割列，默认是制表符“\t” |
| -c       | 按字符进行切割 后加加 n 表示取第几列 比如 -c 1 |

```shell
# 数据准备
[root@vultr shells]$ touch cut.txt
[root@vultr shells]$ vim cut.txt
dong shen
guan zhen
wo wo
lai lai
le le
# 切割 cut.txt 第一列
[root@vultr shells]$ cut -d " " -f 1 cut.txt
dong
guan
wo
lai
le
# 切割 cut.txt 第二、三列
[root@vultr shells]$ cut -d " " -f 2,3 cut.txt
shen
zhen
wo
lai
le
# 在 cut.txt 文件中切割出 guan
[root@vultr shells]$ cat cut.txt |grep guan | cut -d " " -f 1
guan
# 选取系统 PATH 变量值，第 2 个“：”开始后的所有路径：
[root@vultr shells]$ echo $PATH
/usr/local/bin:/usr/bin:/usr/local/sbin:/usr/sbin:/home/atguigu/.local/bin:/
home/atguigu/bin
[root@vultr shells]$ echo $PATH | cut -d ":" -f 3-
/usr/local/sbin:/usr/sbin:/home/atguigu/.local/bin:/home/atguigu/bin
# 切割 ifconfig 后打印的 IP 地址
[root@vultr shells]$ ifconfig ens33 | grep netmask | cut -d " " -f 10
192.168.111.101
```

### awk

一个强大的文本分析工具，把文件逐行的读入，以空格为默认分隔符将每行切片，切开的部分再进行分析处理。

语法：`awk [选项参数] ‘/pattern1/{action1} /pattern2/{action2}...’ filename`

- pattern：表示 awk 在数据中查找的内容，就是匹配模式
- action：在找到匹配内容时所执行的一系列命令

| 选项参数 | 说明                 |
| -------- | -------------------- |
| -F       | 指定输入文件分隔符   |
| -v       | 赋值一个用户定义变量 |

```shell
# 数据准备
[root@vultr shells]$ sudo cp /etc/passwd ./
passwd 数据的含义
用户名:密码(加密过后的):用户 id:组 id:注释:用户家目录:shell 解析器

# 搜索 passwd 文件以 root 关键字开头的所有行，并输出该行的第 7 列。
[root@vultr shells]$ awk -F : '/^root/{print $7}' passwd 
/bin/bash

# 搜索 passwd 文件以 root 关键字开头的所有行，并输出该行的第 1 列和第 7 列，中间以“，”号分割。（只有匹配了 pattern 的行才会执行 action。）
[root@vultr shells]$ awk -F : '/^root/{print $1","$7}' passwd 
root,/bin/bash

# 只显示/etc/passwd 的第一列和第七列，以逗号分割，且在所有行前面添加列名 user，shell 在最后一行添加"dahaige，/bin/zuishuai"。
[root@vultr shells]$ awk -F : 'BEGIN{print "user, shell"} {print $1","$7}
END{print "dahaige,/bin/zuishuai"}' passwd
user, shell
root,/bin/bash
bin,/sbin/nologin 。。。
atguigu,/bin/bash
dahaige,/bin/zuishuai
# 注意：BEGIN 在所有数据读取行之前执行；END 在所有数据执行之后执行。

# 将 passwd 文件中的用户 id 增加数值 1 并输出
[root@vultr shells]$ awk -v i=1 -F : '{print $3+i}' passwd
1
2
3
4
```

awk内置变量

| 变量     | 说明                                   |
| -------- | -------------------------------------- |
| FILENAME | 文件名                                 |
| NR       | 已读的记录数（行号）                   |
| NF       | 浏览记录的域的个数（切割后，列的个数） |

```shell
# 统计 passwd 文件名，每行的行号，每行的列数
[root@vultr shells]$ awk -F : '{print "filename:" FILENAME ",linenum:"
NR ",col:"NF}' passwd
filename:passwd,linenum:1,col:7
filename:passwd,linenum:2,col:7
filename:passwd,linenum:3,col:7
...

# 查询 ifconfig 命令输出结果中的空行所在的行号
[root@vultr shells]$ ifconfig | awk '/^$/{print NR}'
9
18
26

# 切割 IP
[root@vultr shells]$ ifconfig ens33 | awk '/netmask/ {print $2}'
192.168.6.101
```

## 综合应用案例

### 归档文件

实际生产应用中，往往需要对重要数据进行归档备份。

需求：实现一个每天对指定目录归档备份的脚本，输入一个目录名称（末尾不带/），将目录下所有文件按天归档保存，并将归档日期附加在归档文件名上，放在/root/archive 下。这里用到了归档命令：tar后面可以加上-c 选项表示归档，加上-z 选项表示同时进行压缩，得到的文件后缀名为.tar.gz。

脚本实现如下：

```shell
#!/bin/bash
# 首先判断输入参数个数是否为 1
if [ $# -ne 1 ]
then
	echo "参数个数错误！应该输入一个参数，作为归档目录名"
exit
fi
# 从参数中获取目录名称
if [ -d $1 ]
then
	echo
else
	echo
	echo "目录不存在！"
	echo
exit
fi
DIR_NAME=$(basename $1)
DIR_PATH=$(cd $(dirname $1); pwd)

# 获取当前日期
DATE=$(date +%y%m%d)

# 定义生成的归档文件名称
FILE=archive_${DIR_NAME}_$DATE.tar.gz
DEST=/root/archive/$FILE

# 开始归档目录文件
echo "开始归档..."
echo
tar -czf $DEST $DIR_PATH/$DIR_NAME
if [ $? -eq 0 ]
then
	echo
	echo "归档成功！"
	echo "归档文件为：$DEST"
	echo
else
	echo "归档出现问题！"
	echo
fi
exit
```

### 发送消息

我们可以利用 Linux 自带的 mesg 和 write 工具，向其它用户发送消息。

需求：实现一个向某个用户快速发送消息的脚本，输入用户名作为第一个参数，后面直接跟要发送的消息。脚本需要检测用户是否登录在系统中、是否打开消息功能，以及当前发送消息是否为空。脚本实现如下：

```shell
#!/bin/bash
login_user=$(who | grep -i -m 1 $1 | awk '{print $1}')
if [ -z $login_user ]
then
	echo "$1 不在线！"
	echo "脚本退出.."
exit
fi
is_allowed=$(who -T | grep -i -m 1 $1 | awk '{print $2}')
if [ $is_allowed != "+" ]
then
	echo "$1 没有开启消息功能"
	echo "脚本退出.."
exit
fi
if [ -z $2 ]
then
	echo "没有消息发出"
	echo "脚本退出.."
exit
fi
whole_msg=$(echo $* | cut -d " " -f 2- )
user_terminal=$(who | grep -i -m 1 $1 | awk '{print $2}')
echo $whole_msg | write $login_user $user_terminal
if [ $? != 0 ]
then
	echo "发送失败！"
else
	echo "发送成功！"
fi
exit
```

