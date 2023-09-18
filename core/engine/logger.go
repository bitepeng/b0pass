package engine

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/logrusorgru/aurora"
	"github.com/mattn/go-colorable"
)

// LogWriter 多端写日志类
type LogWriter struct {
	*MultiLogWriter
}
type MultiLogWriter struct {
	sync.Map
}

var logWriter = &LogWriter{new(MultiLogWriter)}
var multiLogger = log.New(logWriter.MultiLogWriter, "", log.LstdFlags)
var colorLogger = log.New(colorable.NewColorableStdout(), "", log.LstdFlags)

func init() {
	log.SetOutput(logWriter)
}
func (w *LogWriter) Write(data []byte) (n int, err error) {
	os.Stdout.Write(data)
	return w.MultiLogWriter.Write(data)
}
func (w *MultiLogWriter) Write(data []byte) (n int, err error) {
	w.Range(func(k, v interface{}) bool {
		n, err = k.(io.Writer).Write(data)
		if err != nil {
			w.Delete(k)
		}
		return true
	})
	return
}

// AddWriter 添加日志输出端
func AddWriter(wn io.Writer) {
	logWriter.Store(wn, wn)
}

// MayBeError 优雅错误判断加日志辅助函数
func MayBeError(info error) (hasError bool) {
	if hasError = info != nil; hasError {
		Print(aurora.Red(info))
	}
	return
}
func getNoColor(v ...interface{}) (noColor []interface{}) {
	noColor = append(noColor, v...)
	for i, value := range v {
		if vv, ok := value.(aurora.Value); ok {
			noColor[i] = vv.Value()
		}
	}
	return
}

// Print 带颜色识别
func Print(v ...interface{}) {
	noColor := getNoColor(v...)
	colorLogger.Output(2, fmt.Sprint(v...))
	multiLogger.Output(2, fmt.Sprint(noColor...))
}

// Printf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	noColor := getNoColor(v...)
	colorLogger.Output(2, fmt.Sprintf(format, v...))
	multiLogger.Output(2, fmt.Sprintf(format, noColor...))
}

// Println calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Println.
func Println(v ...interface{}) {
	noColor := getNoColor(v...)
	colorLogger.Output(2, fmt.Sprintln(v...))
	multiLogger.Output(2, fmt.Sprintln(noColor...))
}
