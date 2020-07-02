# Go相关demo代码

## quick-start-1：快速入门（1）

- number: [数值类型demo](quick-start-1/number/main.go)
- string: [字符串类型demo](quick-start-1/string/main.go)
- array: [数组类型demo](quick-start-1/array/main.go)
- slice: [切片类型demo](quick-start-1/slice/main.go)
- map: [字典类型demo](quick-start-1/map/main.go)
- condition: [条件语句demo](quick-start-1/condition/main.go)
- loop: [遍历语句demo](quick-start-1/loop/main.go)

## quick-start-2：快速入门（2）

- function: [函数使用demo](quick-start-2/function/main.go)
- stuct: [结构体使用demo](quick-start-2/stuct/main.go)
- pointer: [指针使用demo](quick-start-2/pointer/main.go)
- interface: [接口使用demo](quick-start-2/interface/main.go)
- exception: [错误处理使用demo](quick-start-2/exception/main.go)

## gin-demo：[gin使用](gin-demo/)

``` bash
├── assets          //静态资源文件
├── conf            //配置文件
├── controllers     //控制器代码
├── go.mod          //mod文件（声明依赖）
├── go.sum
├── helpers         //工具函数包
├── main.go         //入口文件
├── models          //模型代码
├── server          //服务声明
├── services        //逻辑代码
└── views           //视图模板文件
```

## cobra-demo：[cobra使用](cobra-demo/)

``` bash
├── cmd             //命令代码
├── conf            //配置文件
├── go.mod          //mod文件（声明依赖）
├── go.sum
├── helpers         //工具函数包
├── main.go         //入口文件
├── models          //模型代码
└── services        //逻辑代码
```

## concurrency-demo：并发编程

- goroutine: [goroutine-协程demo](concurrency-demo/goroutine/main.go)
- channel: [channel-通道demo](concurrency-demo/channel/main.go)
- lock: [锁机制demo](concurrency-demo/lock/main.go)
- sync-with-global: [基于全局变量的并发控制demo](concurrency-demo/sync-with-global/main.go)
- sync-with-channel: [基于channel的并发控制demo](concurrency-demo/sync-with-channel/main.go)
- sync-with-waitGroup: [基于WaitGroup的并发控制demo](concurrency-demo/sync-with-waitGroup/main.go)
- sync-with-context: [基于Context的并发控制demo](concurrency-demo/sync-with-context/main.go)
- demo-wordCount: [综合实例——文件夹文件大小统计demo](concurrency-demo/demo-wordCount/main.go)
- demo-request: [综合实例——并发http请求demo](concurrency-demo/demo-request/main.go)

## standard-demo：常用标准库

- fmt: [fmt使用demo](standard-demo/fmt-demo/main.go)
- log: [log使用demo](standard-demo/log-demo/main.go)
- strings: [strings使用demo](standard-demo/strings-demo/main.go)
- strconv: [strconv使用demo](standard-demo/strconv-demo/main.go)
- time: [time使用demo](standard-demo/time-demo/main.go)
- os: [os使用demo](standard-demo/os-demo/main.go)
- http: [http使用demo](standard-demo/http-demo/main.go)
- sql: [sql使用demo](standard-demo/sql-demo/main.go)
- template: [sql使用demo](standard-demo/template-demo/main.go)

## oop-demo：面向对象编程

- define: [类定义和实例化demo](oop-demo/define/main.go)
- encapsulation: [封装demo](oop-demo/encapsulation/)
- inheritance: [继承demo](oop-demo/inheritance/main.go)
- polymorphism: [多态demo](oop-demo/polymorphism/main.go)