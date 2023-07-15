package logger

import (
	"fmt"
	"os"
	"time"
)

// 日志写入器
var Writer = NewFlieLogger("/root/liweiran/project/webServer/log/log_data", "web.log")

// 定义日志等级
type LogLevel int8

const (
	UNKOWN LogLevel = iota
	INFO
	WARN
	ERROR
)

// 获取日志的字符串格式
func getLogStr(level LogLevel) string {

	switch level {
	case INFO:
		return "info"
	case WARN:
		return "warnig"
	case ERROR:
		return "error"
	default:
		return "unknow"
	}
}

// 日志写入器结构
type FileLogger struct {
	filePath string
	fileName string
	fileObj  *os.File
}

// 构造函数
func NewFlieLogger(fp, fn string) *FileLogger {

	fl := &FileLogger{
		filePath: fp,
		fileName: fn,
	}
	err := fl.initLogger()
	if err != nil {
		panic(err)
	}
	return fl
}

// 初始化日志写入器
func (f *FileLogger) initLogger() error {
	now := time.Now()
	year, month, day := now.Year(), now.Month(), now.Day()
	path := fmt.Sprintf("%s/%s_%d-%d-%d", f.filePath, f.fileName, year, month, day)
	fileObj, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file fail, err: %v\n", err)
		return err
	}
	f.fileObj = fileObj
	return nil
}

// 检查是否切割日志文件
func (f *FileLogger) checkSplitLog() error {
	now := time.Now()
	year, month, day := now.Year(), now.Month(), now.Day()
	path := fmt.Sprintf("%s/%s_%d-%d-%d", f.filePath, f.fileName, year, month, day)
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		fileObj, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("open log file fail, err: %v\n", err)
			return err
		}
		f.fileObj = fileObj
	}
	return nil
}

// 打印日志操作
func (f *FileLogger) Log(level LogLevel, msg string) {
	now := time.Now()
	fmt.Fprintf(f.fileObj, "[%s] [%s] %s\n", now.Format("2006-01-02 15:04:05"), getLogStr(level), msg)
}

// info类型日志
func (f FileLogger) Info(msg string, a ...interface{}) {
	msg = fmt.Sprintf(msg, a...)
	f.checkSplitLog()
	f.Log(INFO, msg)
}

// warning类型日志
func (f FileLogger) Warning(msg string, a ...interface{}) {
	msg = fmt.Sprintf(msg, a...)
	f.checkSplitLog()
	f.Log(WARN, msg)
}

// error类型日志
func (f FileLogger) Error(msg string, a ...interface{}) {
	msg = fmt.Sprintf(msg, a...)
	f.checkSplitLog()
	f.Log(ERROR, msg)
}

func (f *FileLogger) close() {
	f.fileObj.Close()
}
