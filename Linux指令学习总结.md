# 常用Linux指令

## 1、文件、目录管理命令

### 基础命令：

```
pwd：输出当前工作目录

cd：切换工作目录，
	选项： /：转到根目录；
    	  ~：转到/home/user目录
    	  -：转到切换前工作目录
    	  
ls：查看目录下所有内容
	选项： -a：显示包含隐藏文件
		  -l：显示文件详细信息
		  -s：显示文件按大小排序
	
touch + file：新建文件

grep + string ：查找匹配的字符串，常用于大量文件中查找特定的字符


```

### 文件阅读命令：

```
cat + file：在屏幕显示文件全部内容
	选项：
		-n：显示行号

vi/vim + file：采用vim编辑器显示文件内容，可进行文件的修改
		i：进行文件编辑
		esc：退出编辑并进行相关保存操作
			w：保存修改，不退出vi
			w!：强制保存修改，不退出vi
			q：不保存修改，退出vi
			q!：不保存修改，强制退出vi
			wq：保存修改，退出vi
			wq!：保存修改，强制退出vi

head：查看文件开头10行内容，可修改查看数量：head -20 
tail：查看文件结尾10行内容，常用命令：tail -f file 实时查看文件的新信息

more：查看文件所有内容，以空格翻页。也可用于文件夹的查看，文件显示完自动退出或q建退出
less：查看文件所有内容，以方向键进行翻动，按q退出
```

## 2、文件修改、备份和压缩命令

```
mv：移动文件
	选项：
		-i：提示互动信息，如覆盖提示
		-f：强制操作，
		-v：显示详细信息，如移动进度

cp：拷贝命令，可用于文件及文件夹的拷贝。
	cp [source file] [target file]	拷贝文件操作
		cp a.txt b.txt	
		cp a.txt /home/data/b.txt
	cp -r [source directory] [target directory]	拷贝文件夹操作
		cp -r A/ B/
		cp -r A/ /home/data/B/
		
tar：磁盘文件备份命令，可用于大量文件打包成一个文件。
	选项：
		创建新归档tar文件
            -c：创建新归档
            -f：创建新归档时指定名称
            -v：显示文件的归档进度
            tar –cvf filename.tar directory/file /home/mine
        查看tar文件内容
        	-t：显示tar文件中的文件列表
        	tar -tvf filename.tar
        提取tar文件内容
        	-x：从归档tar文件中提取文件
        	tar -xvf filename.tar
        创建一个用tar和gzip归档并压缩的文件，
        	-z:	使用gzip压缩tar文件
        	tar –czvf filename.tgz file
        	使用gunzip命令解压filename.tgz文件，则filename.tgz会被删除，以filename.tar代替      
zip/gzip：把文件以gzip来压缩
	gzip filename	压缩文件，并保存为filename.gz，源文件会被删除
	gunzip filename.gz	解压文件，源文件会被删除
	gzip -r filename.gz file1 file2 file3 /usr/work/school 	压缩多个文件和目录


```



## 3、系统命令

```
ps：查看目前程序执行的情况
	选项：
		-u：列出使用者的名称和使用时间
		-x：列出所有程序，包括那些没有终端机的程序
		-r：只列出正在执行的前台程序，不列出其他信息
		-m：列出内存分布的情况
		-a：显示所有终端机下执行的程序
		-l：查看此次pid所有信息
		ps -aux | grep named  查看named进程详细信息
		ps axo pid,comm,pcpu  查看进程的PID、名称以及CPU占用率
		ps -ef | grep filename/port	根据字段、端口查看进程id
		
top：动态查看目前程序执行的情景和内存使用的情况
	选项：
		-p：指定进程
		
kill：终止正在运行的进程
	kill -9 进程id	杀进程

```

## 4、权限命令

```
chmod：改变文件许可权限。操作r/4(读)、w/2(写)、x/1(可执行)三种权限，ls -l可查看文件的权限。
		具体为文件使用组添加不同的权限，u(文件拥有者)、g(群组其他人)、o(其他人)、a(所有人)
		chmod 777 file	为文件开通最高权限
		
su：切换用户状态。
	
```

## 5、线上查询命令

```
man：查询和解释命令的使用方法
	man 命令
		选项：
		-a：显示所有
		-f：显示给定关键字的简短描述信息
		man ll		//查看ll命令相关用法

locate：查询定位文件和目录
	locate filename
		locate test		//查询test的文件或目录
		locate /etc/sh	//查询etc目录下sh开头的文件
		
注：locate与find区别，locate会为硬盘中的所有档案和目录资料先建立一个索引数据库，在执行loacte时直接找该索引，索引数据库一般是由操作系统管理，查询速度会较快
[root@ac-test03 cs-engine-message-server]# cd /var/lib/mlocate
[root@ac-test03 mlocate]# ls
mlocate.db


whatis：查询某个命令的含义，相当于man -f 命令
	whatis 命令
		whatis ll	//查询ll命令用法

whereis：查找二进制程序、代码等相关文件的路径。whereis命令只能用于程序名的搜索，也是从数据库中查找数据
	选项：
		-b：只查找二进制文件
		-B<目录>：在设置的目录下查找二进制文件
		-f：不显示文件名前的路径
		-s：只查找原始代码文件


find：在指定目录下查找文件
	选项：
		find / -name "file"	//在根目录下查找file

```

## 6、网络操作命令

```
ftp：用来设置文件系统相关功能
	选项：
		-d：详细显示指令执行过程，便于排错或分析程序执行的情况；
        -i：关闭互动模式，不询问任何问题；
        -g：关闭本地主机文件名称支持特殊字符的扩充特性；
        -n：不使用自动登录；
        -v：显示指令执行过程
        
        登录ftp.dark.com主机：ftp ftp.dark.com	//需要有权限
        

scp：加密的方式在本地主机和远程主机之间复制文件，与cp命令不同的是，cp只用于本机进行文件拷贝。
	scp 选项 源文件 目标服务地址
	选项：
		-r：以递归方式复制
		-P：指定远程主机的端口号
		-C：使用压缩
		scp filename user@host:some/directory	//传输文件
		scp -r directory user@host:some/directory	//传输文件夹

netstat：查看网络状态信息，与ps命令搭配使用，查询某个应用的进程信息
	选项：
		-a	显示所有连线中的Socket
		-n  接使用ip地址，而不通过域名服务器
		-p  示正在使用Socket的程序识别码和程序名称
		-t  示TCP传输协议的连线状况
        -u  示UDP传输协议的连线状况
        -v  示指令执行过程
		netstat -nap |grep port	  根据端口查网络状态
		netstat    -s             显示所有端口的统计信息		
		netstat    -l             只显示监听端口
		netstat    -a             列出所有端口

```

## 7、磁盘命令

```
df：df命令用来检查硬盘分区和已挂在的文件系统的磁盘空间
	选项：
		-a  包含全部的文件系统
		-h  以可读性较高的方式来显示信息
		-i  显示inode的信息
		-t<文件系统类型>或--type=<文件系统类型>：仅显示指定文件系统类型的磁盘信息
		-l  仅显示本地端的文件系统
		
		
du：显示每个文件和目录的磁盘使用空间，也是查看使用空间的，但是与df命令不同的是du命令是对文件和目录磁盘使用的空间的查看
	选项：
		-a：显示目录中个别文件的大小
		-b:显示文件目录大小，以B单位显示
		-c, 除了显示个别目录或文件的大小外，同时也显示所有目录或文件的总和
		du -sh * |sort -rh	//文件从大到小排序
		du -sh ./*/		//只显示当前目录下子目录的大小。


```

## 8、其他命令

```
mkdir：创建目录

rmdir：删除目录

echo：打印一行文件

```