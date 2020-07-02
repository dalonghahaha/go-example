package main

import (
	"fmt"
	"time"
)

func main() {
	//获取时间相关信息
	exampleTimeInfo()
	//时间相关操作
	exampleTimeOpration()
	//定时器
	exampleTimer()
}

func exampleTimeInfo() {
	// 返回当前本地时间
	t := time.Now()
	fmt.Println(t)
	// 返回指定时间
	t2 := time.Date(2019, time.February, 7, 0, 0, 0, 0, time.UTC)
	fmt.Println(t2)
	// 返回t的地点和时区信息
	fmt.Println(t.Location())
	// 返回采用本地和本地时区，但指向同一时间点的Time
	fmt.Println(t.Local())
	// 返回unix时间，即从时间点January 1, 1970 UTC到时间点t所经过的时间（单位秒）
	fmt.Println(t.Unix())
	// 将t表示为Unix时间，即从时间点January 1, 1970 UTC到时间点t所经过的时间（单位纳秒）
	// 如果纳秒为单位的unix时间超出了int64能表示的范围，结果是未定义的。注意这就意味着Time零值调用UnixNano方法的话，结果是未定义的
	fmt.Println(t.UnixNano())
	// 返回采用UTC和零时区，但指向同一时间点的Time
	fmt.Println(t.UTC())
	// 计算t所在的时区，返回该时区的规范名（如"CET"）和该时区相对于UTC的时间偏移量（单位秒）
	fmt.Println(t.Zone())
	// 返回时间点t对应的年、月、日
	fmt.Println(t.Date())
	// 返回t对应的那一天的时、分、秒
	fmt.Println(t.Clock())
	// 返回时间点t对应的年份
	fmt.Println(t.Year())
	// 返回时间点t对应的那一年的第几天，平年的返回值范围[1,365]，闰年[1,366]
	fmt.Println(t.YearDay())
	// 返回时间点t对应那一年的第几月
	fmt.Println(t.Month())
	// 返回时间点t对应那一月的第几日
	fmt.Println(t.Day())
	// 返回时间点t对应的那一周的周几
	fmt.Println(t.Weekday())
	// 返回t对应的那一天的第几小时，范围[0, 23]
	fmt.Println(t.Hour())
	// 返回t对应的那一小时的第几分种，范围[0, 59]
	fmt.Println(t.Minute())
	// 返回t对应的那一分钟的第几秒，范围[0, 59]
	fmt.Println(t.Second())
	// 返回t对应的那一秒内的纳秒偏移量，范围[0, 999999999]
	fmt.Println(t.Nanosecond())
	// 返回采用loc指定的地点和时区，但指向同一时间点的Time。如果loc为nil会panic
	loc, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		fmt.Println(t.In(loc))
	}
}

func exampleTimeOpration() {
	t := time.Now()
	t2 := time.Date(2019, time.February, 7, 0, 0, 0, 0, time.UTC)
	loc, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		fmt.Println(t.In(loc))
	}
	// 判断两个时间是否相同，会考虑时区的影响，因此不同时区标准的时间也可以正确比较。本方法和用t==u不同，这种方法还会比较地点和时区信息
	fmt.Println(t2.Equal(time.Now()))
	// 判断t的时间点是否在u之前
	fmt.Println(t2.Before(time.Now()))
	// 判断t的时间点是否在u之后
	fmt.Println(t2.After(time.Now()))
	// 增加2小时, 减少使用负号
	fmt.Println(t.Add(time.Hour * 2))
	// 返回增加了给出的年份、月份和天数的时间点Time。例如，时间点January 1, 2011调用AddDate(-1, 2, 3)会返回March 4, 2010
	fmt.Println(t.AddDate(1, 2, 3))
	// 获取当前时间减去指定时间的时间
	// 如果结果超出了Duration可以表示的最大值/最小值，将返回最大值/最小值。要获取时间点t-d（d为Duration），可以使用t.Add(-d)
	fmt.Println(t.Sub(time.Now()))
	// 根据layout指定的格式返回t代表的时间点的格式化文本表示
	fmt.Println(t.Format("2006-01-02 15:04:05"))
	// 返回采用如下格式字符串的格式化时间
	// "2006-01-02 15:04:05.999999999 -0700 MST"
	fmt.Println(t.String())
	// 解析一个格式化的时间字符串并返回它代表的时间
	fmt.Println(time.Parse("2006 Jan 02 15:04:05", "2019 Feb 07 12:15:30.918273645"))
	// 类似Parse但有两个重要的不同之处
	// 第一，当缺少时区信息时，Parse将时间解释为UTC时间，而ParseInLocation将返回值的Location设置为loc
	// 第二，当时间字符串提供了时区偏移量信息时，Parse会尝试去匹配本地时区，而ParseInLocation会去匹配loc
	fmt.Println(time.ParseInLocation("2006 Jan 02 15:04:05", "2019 Feb 07 12:15:30.918273645", loc))
}

func exampleTimer() {
	// 阻塞当前go程至少d代表的时间段。d<=0时，Sleep会立刻返回
	time.Sleep(time.Second * 2)
	// 创建一个Timer，它会在最少过去时间段d后到期，向其自身的C字段发送当时的时间
	timer := time.NewTimer(time.Second * 10)
	// 创建一个Ticker，它会在每次过去时间段d后到期，向其自身的C字段发送当时的时间
	ticker := time.NewTicker(time.Second * 5)
	for {
		select {
		case <-timer.C:
			fmt.Println("timer", time.Now())
		case <-ticker.C:
			fmt.Println("ticker", time.Now())
		default:

		}
	}
}
