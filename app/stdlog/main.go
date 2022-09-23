package main

import (
	"log"

	"github.com/natefinch/lumberjack"
)

// 使用lumberjack日志滚动记录器
func main() {
	w := &lumberjack.Logger{
		Filename:   "./rolling.log",
		MaxSize:    500, // megabytes
		MaxAge:     28,  // days
		MaxBackups: 3,
		LocalTime:  false, // disabled by default
		Compress:   false, // disabled by default
	}
	log.SetOutput(w)

	log.Printf("lalala")
}

// 打印日志到多个目标
//func main() {
//	file, err := os.OpenFile("./file.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
//	if err != nil {
//		log.Fatalf("os.OpenFile error:%+v", err)
//	}
//	defer file.Close()
//
//	multi := io.MultiWriter(os.Stdout, file)
//
//	log.SetOutput(multi)
//	log.SetFlags(log.LstdFlags | log.Lmsgprefix)
//	log.SetPrefix("[multi]")
//
//	log.Println("log print")
//}

// 打印日志到文件
//func main_to_file() {
//	// todo:这里的0666可以用标准库自带的啥常量代替嘛?
//	file, err := os.OpenFile("./file.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
//	if err != nil {
//		log.Fatalf("os.OpenFile error:%+v", err)
//	}
//	defer file.Close()
//
//	log.SetOutput(file)
//	log.SetFlags(log.LstdFlags | log.Lmsgprefix)
//	log.SetPrefix("[file]")
//
//	log.Println("log print")
//}
