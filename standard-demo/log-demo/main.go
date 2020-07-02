package main

import (
	"log"
	"os"
)

type user struct {
	name string
}

func main() {
	exampleLog()

	exampleCustomLog()
}

func exampleLog() {
	u := user{"tang"}
	//Printf 格式化输出
	log.Printf("%v\n", u)        //格式化输出结构
	log.Printf("%+v\n", u)       //格式化输出结构
	log.Printf("%#v\n", u)       //输出值的Go语言表示方法
	log.Printf("%T\n", u)        //输出值的类型的Go语言表示
	log.Printf("%t\n", true)     //输出值的 true 或 false
	log.Printf("%b\n", 1024)     //二进制表示
	log.Printf("%c\n", 11111111) //数值对应的 Unicode 编码字符
	log.Printf("%d\n", 10)       //十进制表示
	log.Printf("%o\n", 8)        //八进制表示
	log.Printf("%q\n", 22)       //转化为十六进制并附上单引号
	log.Printf("%x\n", 1223)     //十六进制表示，用a-f表示
	log.Printf("%X\n", 1223)     //十六进制表示，用A-F表示
	log.Printf("%U\n", 1233)     //Unicode表示
	log.Printf("%b\n", 12.34)    //无小数部分，两位指数的科学计数法6946802425218990p-49
	log.Printf("%e\n", 12.345)   //科学计数法，e表示
	log.Printf("%E\n", 12.34455) //科学计数法，E表示
	log.Printf("%f\n", 12.3456)  //有小数部分，无指数部分
	log.Printf("%g\n", 12.3456)  //根据实际情况采用%e或%f输出
	log.Printf("%G\n", 12.3456)  //根据实际情况采用%E或%f输出
	log.Printf("%s\n", "wqdew")  //直接输出字符串或者[]byte
	log.Printf("%q\n", "dedede") //双引号括起来的字符串
	log.Printf("%x\n", "abczxc") //每个字节用两字节十六进制表示，a-f表示
	log.Printf("%X\n", "asdzxc") //每个字节用两字节十六进制表示，A-F表示
}

func exampleCustomLog() {
	logFile := "./log"
	// 创建日志文件
	file, err := os.Create(logFile)
	if err != nil {
		log.Fatal(err)
	}
	// 创建一个Logger
	// 参数out设置日志信息写入的目的地
	// 参数prefix会添加到生成的每一条日志前面
	// 参数flag定义日志的属性（时间、文件等等）
	logger := log.New(file, "[INFO]", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("111")
	logger.Println("222")
	logger.Println("333")
	// 设置logger的输出选项(日志属性)
	logger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	// 设置logger的输出前缀
	logger.SetPrefix("[ERROR]")
	// 设置标准logger的输出目的地，默认是标准错误输出
	logger.SetOutput(file)
	logger.Println("444")
	logger.Println("555")
	logger.Println("666")
}
