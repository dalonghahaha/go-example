package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// 获取系统基本信息
	example()
	// 文件夹操作
	exampleDir()
	// 文件操作
	exampleFile()
	// 进程操作
	exampleProcess()
	// 执行命令
	exampleCommand()
}

const (
	DirPath  = "testdata/"
	FilePath = "testdata/test_os.txt"
)

func example() {
	// 返回内核提供的主机名
	fmt.Println(os.Hostname())
	// 返回底层的系统内存页的尺寸
	fmt.Println("agesize:", os.Getpagesize())
	// 返回调用者的用户ID
	fmt.Println("uid:", os.Getuid())
	// 返回调用者的有效用户ID
	fmt.Println("euid:", os.Geteuid())
	// 返回调用者的组ID
	fmt.Println("gid:", os.Getgid())
	// 返回调用者的有效组ID
	fmt.Println("egid:", os.Getegid())
	// 返回调用者所属的所有用户组的组ID
	fmt.Println(os.Getgroups())
	// 返回调用者所在进程的进程ID
	fmt.Println("pid:", os.Getpid())
	// 返回调用者所在进程的父进程的进程ID
	fmt.Println("ppid:", os.Getppid())
	// 返回表示环境变量的格式为"key=value"的字符串的切片拷贝
	fmt.Println(os.Environ())
	// 返回对应当前工作目录的根路径。如果当前目录可以经过多条路径抵达（因为硬链接），Getwd会返回其中一个
	fmt.Println(os.Getwd())
	// 返回用于保管临时文件的默认目录
	fmt.Println(os.TempDir())
	// 返回用于用户的缓存数据的默认根目录。用户应该在这个目录中创建自己的特定于应用程序的子目录，并使用它
	fmt.Println(os.UserCacheDir())
	// 用户文件夹路径
	fmt.Println(os.UserHomeDir())
}

func exampleDir() {
	// 使用指定的权限和名称创建一个目录。如果出错，会返回*PathError底层类型的错误
	if err := os.Mkdir(DirPath, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	// 使用指定的权限和名称创建一个目录，包括任何必要的上级目录，并返回nil，否则返回错误
	// 权限位perm会应用在每一个被本函数创建的目录上
	// 如果path指定了一个已经存在的目录，MkdirAll不做任何操作并返回nil
	if err := os.MkdirAll(DirPath, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

func exampleFile() {
	// 以读写方式打开文件，并且内容写入方式为添加，如果文件不存在以0755操作模式创建文件
	f, err := os.OpenFile(FilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		log.Fatal(err)
	}
	// 关闭文件
	defer f.Close()
	// 写入内容（向文件内添加内容）
	if _, err := f.Write([]byte("appended some data\n")); err != nil {
		log.Fatal(err)
	}
	// 判断文件是否存在
	if _, err := os.Stat(FilePath); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("file does not exist")
		} else {
			log.Fatal(err)
		}
	}
	// 修改指定文件的mode操作模式
	// 如果name指定的文件是一个符号链接，它会修改该链接的目的地文件的mode。如果出错，会返回*PathError底层类型的错误
	if err := os.Chmod(FilePath, 0644); err != nil {
		log.Fatal(err)
	}
	// 修改指定文件的用户id和组id
	// 如果name指定的文件是一个符号链接，它会修改该链接的目的地文件的用户id和组id。如果出错，会返回*PathError底层类型的错误
	if err := os.Chown(FilePath, 501, 20); err != nil {
		log.Fatal(err)
	}
	// 返回一个描述指定文件的FileInfo
	// 如果指定的文件对象是一个符号链接，返回的FileInfo描述该符号链接的信息，本函数不会试图跳转该链接。如果出错，返回的错误值为*PathError类型
	fi, err := os.Lstat(FilePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("permissions: %#o\n", fi.Mode().Perm()) // 0400, 0777, etc.
	switch mode := fi.Mode(); {
	case mode.IsRegular():
		fmt.Println("regular file")
	case mode.IsDir():
		fmt.Println("directory")
	case mode&os.ModeSymlink != 0:
		fmt.Println("symbolic link")
	case mode&os.ModeNamedPipe != 0:
		fmt.Println("named pipe")
	}
}

func exampleProcess() {
	// 初始化一个 保管将被StartProcess函数用于一个新进程的属性
	attr := &os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}, //其他变量如果不清楚可以不设定
	}
	// 使用提供的属性、程序名、命令行参数开始一个新进程
	// StartProcess函数是一个低水平的接口
	// os/exec包提供了高水平的接口，应该尽量使用该包。如果出错，错误的底层类型会是*PathError
	p, err := os.StartProcess("/usr/bin/vim", []string{"/usr/bin/vim", "temp.txt"}, attr)
	if err != nil {
		log.Fatal(err)
	}
	// 根据进程id查找一个运行中的进程
	p, err = os.FindProcess(p.Pid)
	if err != nil {
		log.Fatal(err)
	}
	// 立刻退出进程
	if err := p.Kill(); err != nil {
		log.Fatal(err)
	}
}

func exampleCommand() {
	cmd := exec.Command("go", "doc", "exec")
	// 设置输入内容
	cmd.Stdin = strings.NewReader("")
	// 声明buffer
	var out bytes.Buffer
	// 设置输出内容填充地址
	cmd.Stdout = &out
	// 执行c包含的命令，并阻塞直到完成
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out.String())
}
