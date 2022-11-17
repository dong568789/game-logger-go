package utils

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"time"
)

var Logger *logger

//Logger 日志，实现了Write接口
type logger struct {
	t    time.Time
	fp   *os.File
	path string
	m    sync.RWMutex
}

//new 初始化
func InitLogger(path string) {
	if path == "" {
		path = "logs"
	}
	if !isDir(path) {
		err := os.Mkdir(path, 0755)
		if err != nil {
			panic("无法创建runtime目录")
		}
	}
	Logger = &logger{
		t:    time.Now(),
		path: path,
	}
	Logger.setLogfile()
}

func (l *logger) Info(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	l.Println("Info", msg)
}

func (l *logger) Warning(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	l.Println("Warning", msg)
}

func (l *logger) Error(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	l.Println("Error", msg)
}

func (l *logger) Println(prefix string, msg string) {
	msg = fmt.Sprintf(
		"%s %s %s\n",
		"["+prefix+"]",
		time.Now().Format("2006-01-02 15:04:05"),
		msg,
	)
	l.Write([]byte(msg))
}

//Write 实现Write接口，用于写入
func (l *logger) Write(msg []byte) (n int, err error) {
	today := dateToStr(time.Now())
	loggerDate := dateToStr(l.t)

	//如果当前日期与logger日期不一致，表示是新的一天，需要关闭原日志文件，并更新日期与日志文件
	if today != loggerDate && l.fp != nil {
		l.fp.Close()
		l.fp = nil
	}

	if l.fp == nil {
		l.setLogfile()
	}

	//写入
	if l.fp != nil {
		return l.fp.Write(msg)
	}

	return 0, errors.New("无法写入日志")
}

//setLogfile 更新日志文件
func (l *logger) setLogfile() error {
	year, month, day := time.Now().Date()
	//dir := fmt.Sprintf("logs/%d/%02d", year, month)
	//
	////锁住，防止并发时，多次执行创建。os.MkdirAll在目录存在时，也不会返回错误，锁不锁都行
	l.m.Lock()
	defer l.m.Unlock()
	logfile := fmt.Sprintf("%s/%d-%02d-%d.log", l.path, year, month, day)
	//打开新的日志文件，用于写入
	fp, err := os.OpenFile(logfile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	l.fp = fp
	return nil
}

//isDir 是否是目录
func isDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

//dateToStr 时间转换为日期字符串
func dateToStr(t time.Time) string {
	return t.Format("2006-01-02")
}
