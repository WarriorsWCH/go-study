### mac上安装go环境
#### 1.Homebrew的使用：

- （1）安装软件：brew install 软件名，例如：brew install wget

- （2）搜索软件：brew search 软件名

- （3）卸载软件：brew uninstall 软件名

- （4）更新软件：brew upgrade 软件名，例如：brew upgrade git

- （5）查看使用brew已安装的软件列表：brew list

- （6）查看软件信息：brew info /home 软件名，例如：brew info git/brew home git

- （7）查看哪些已安装的程序需要更新：brew outdated

brew安装好之后可以使用上述的一些命令来测试一下是否安装成功，例如使用brew list来查看一下当前brew安装的软件列表

#### 2.使用brew安装go

- 1、安装命令：brew install go

- 2、检查：输入brew info go或者go env即可查看当前安装的golang版本信息

#### 3.配置go路径环境

主要是GOROOT和GOPATH

GOROOT：就是go的安装环境
GOPATH：作为编译后二进制的存放目的地和import包时的搜索路径。其实说通俗点就是你的go项目工作目录。通常情况下GOPATH包含三个目录：bin、pkg、src。
src目录下主要存放go的源文件

pkg目录存放编译好的库文件，主要是*.a文件;

bin目录主要存放可执行文件
注意：千万不要把GOPATH设置成go的安装路径，可以自己在用户目录下创建一个目录，例如mygo

一般安装好go之后，使用go env查看一下当前环境。此时显示出来的GOROOT就是你使用brew安装go的安装目录，这个路径要记下来。接下来要在bash_profile文件中进行配置。

使用vim ~/.bash_profile

然后在这个文件中进行编辑，下面以我的电脑为例，路径这种要根据不同人的情况而定
```
export GOROOT=/usr/local/go
export GOPATH=/Users/wangcc/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
```

- GOROOT： go安装目录
- GOPATH：go工作目录
- GOBIN：go可执行文件目录
- PATH：将go可执行文件加入PATH中，使GO命令与我们编写的GO应用可以全局调用

编辑完之后退出保存文件，然后使用

source ~/.bash_profile

使之生效，然后再使用go env查看当前环境，可以发现已经是你配置文件中设置的路径环境了

- go run 运行当个.go文件
- go install 在编译源代码之后还安装到指定的目录
- go build 加上可编译的go源文件可以得到一个可执行文件
- go get = git clone + go install 从指定源上面下载或者更新指定的代码和依赖，并对他们进行编译和安装