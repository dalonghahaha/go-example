package main

import (
	"fmt"
	"strings"
)

func main() {
	// 判断两个utf-8编码字符串，大小写不敏感
	s, t := "hello go", "hello Go"
	is_equal := strings.EqualFold(s, t)
	fmt.Println("EqualFold: ", is_equal) // EqualFold:  true

	// 判断s是否有前缀字符串prefix
	prefix := "hello"
	has_prefix := strings.HasPrefix(s, prefix)
	fmt.Println("HasPrefix: ", has_prefix) // true

	// 判断s是否有后缀字符串suffix
	suffix := "go"
	has_suffix := strings.HasSuffix(s, suffix)
	fmt.Println("HasSuffix: ", has_suffix) // true

	// 判断字符串s是否包含子串substr
	substr := "lo"
	con := strings.Contains(s, substr)
	fmt.Println("Contains: ", con) // true

	// 判断字符串s是否包含utf-8码值r
	r := rune(101)
	ru := 'e'
	con_run := strings.ContainsRune(s, r)
	fmt.Println("ContainsRune: ", con_run, r, ru) // true

	// 子串sep在字符串s中第一次出现的位置，不存在则返回-1
	sep := "o"
	sep_idnex := strings.Index(s, sep)
	fmt.Println(sep_idnex) // 4

	// 子串sep在字符串s中最后一次出现的位置，不存在则返回-1
	sep_lastindex := strings.LastIndex(s, sep)
	fmt.Println(sep_lastindex) // 7

	// 返回s中每个单词的首字母都改为标题格式的字符串拷贝
	title := strings.Title(s)
	fmt.Println(title) // Hello Go

	// 返回将所有字母都转为对应的标题版本的拷贝
	to_title := strings.ToTitle(s)
	fmt.Println(to_title) // HELLO GO

	// 返回将所有字母都转为对应的小写版本的拷贝
	s_lower := strings.ToLower(s)
	fmt.Println(s_lower) // hello go

	// 返回count个s串联的字符串
	s_repeat := strings.Repeat(s, 3)
	fmt.Println(s_repeat) // hello gohello gohello go

	// 返回将s中前n个不重叠old子串都替换为new的新字符串，如果n<0会替换所有old子串
	s_old, s_new := "go", "world"
	s_replace := strings.Replace(s, s_old, s_new, -1)
	fmt.Println(s_replace) // hello world

	// 返回将s前后端所有cutset包含的utf-8码值都去掉的字符串
	s, cutset := "#abc!!!", "#!"
	s_new = strings.Trim(s, cutset)
	fmt.Println(s, s_new) // #abc!!! abc

	// 返回将字符串按照空白（unicode.IsSpace确定，可以是一到多个连续的空白字符）分割的多个字符串
	s = "hello world! go language"
	s_fields := strings.Fields(s)
	for k, v := range s_fields {
		fmt.Println(k, v)
	}
	// 0 hello
	// 1 world!
	// 2 go
	// 3 language

	// 用去掉s中出现的sep的方式进行分割，会分割到结尾，并返回生成的所有片段组成的切片
	s_split := strings.Split(s, " ")
	fmt.Println(s_split) // [hello world! go language]

	// 将一系列字符串连接为一个字符串，之间用sep来分隔
	s_join := strings.Join([]string{"a", "b", "c"}, "/")
	fmt.Println(s_join) // a/b/c

	// 将s的每一个unicode码值r都替换为mapping(r)，返回这些新码值组成的字符串拷贝。如果mapping返回一个负值，将会丢弃该码值而不会被替换
	map_func := func(r rune) rune {
		switch {
		case r > 'A' && r < 'Z':
			return r + 32
		case r > 'a' && r < 'z':
			return r - 32
		}
		return r
	}
	s = "Hello World!"
	s_map := strings.Map(map_func, s)
	fmt.Println(s_map) // hELLO wORLD!
}
