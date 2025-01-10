package conf

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const (
	desc = `CodingSoul系列课程包括:
手写操作系统、手写数据库、手写网络协议栈等课程, 目前还在更新中
详情请前往 [https://s.c1ns.cn/Jz2bS] 查看.`
)

var courses = map[string]string{
	"重铸编程之魂":  "https://s.c1ns.cn/273dW",
	"手写操作系统":  "https://s.c1ns.cn/av4xB",
	"手写网络协议栈": "https://s.c1ns.cn/P1lzo",
	"手写数据库系统": "https://s.c1ns.cn/xuunr",
}

var (
	Name        = "cstool"
	Version     = "1.0.0"
	Intro       = "CodingSoul系列课程专用工具集"
	Description string
)

func init() {
	// banner
	var banner string
	file, err := os.OpenFile("conf/banner.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
		return
	}
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
		return
	}
	banner = string(data)
	// courses
	var builder strings.Builder
	for k, v := range courses {
		builder.WriteString(fmt.Sprintf("- %30s - %s\n", v, k))
	}
	Description = fmt.Sprintf("%s%8s\n%s\n\n%s", banner, Version, desc, builder.String())
}
