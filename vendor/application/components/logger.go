package components

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var (
	defaultLogger *Logger
)

//Logger just wrap logrus some method
//more detail setting please use Logrus() method get logrus.Logger,
//then use logrus setting
type Logger struct {
	logrus        *logrus.Logger
	writerManager *loggerWriterManager
}

//NewLogger return new Logger. Please EnablePrint or add Other Writer by yourself
func NewLogger() *Logger {
	logger := &Logger{
		logrus:        logrus.New(),
		writerManager: newloggerWriterManager(),
	}

	logger.logrus.Out = logger.writerManager

	return logger
}

//Logrus return logrus.Logger for more complex cases
func (l *Logger) Logrus() *logrus.Logger {
	return l.logrus
}

//SetLevel set loglevel
func (l *Logger) SetLevel(level logrus.Level) {
	l.logrus.Level = level
}

//Debug same as logrus Debug
func (l *Logger) Debug(args ...interface{}) {
	l.logrus.Debug(args...)
}

//Info same as logrus Info
func (l *Logger) Info(args ...interface{}) {
	l.logrus.Info(args...)
}

//Print same as logrus Print
func (l *Logger) Print(args ...interface{}) {
	l.logrus.Print(args...)
}

//Warning same as logrus Warning
func (l *Logger) Warning(args ...interface{}) {
	l.logrus.Warning(args...)
}

//Error same as logrus Error
func (l *Logger) Error(args ...interface{}) {
	l.logrus.Error(args...)
}

//Fatal same as logrus Fatal
func (l *Logger) Fatal(args ...interface{}) {
	l.logrus.Fatal(args...)
}

//Panic same as logrus Panic
func (l *Logger) Panic(args ...interface{}) {
	l.logrus.Panic(args...)
}

//EnablePrint add os.Stdout to logger writer
func (l *Logger) EnablePrint() {
	l.writerManager.AddWriter("print", os.Stdout)
}

//DisablePrint remove os.Stdout from logger writer
func (l *Logger) DisablePrint() {
	l.writerManager.RemoveWriter("print")
}

type loggerWriterManager struct {
	writers map[string]io.Writer
}

func newloggerWriterManager() *loggerWriterManager {
	return &loggerWriterManager{
		writers: make(map[string]io.Writer),
	}
}

func (lwm *loggerWriterManager) RemoveWriter(name string) {
	delete(lwm.writers, name)
}

func (lwm *loggerWriterManager) AddWriter(name string, w io.Writer) {
	lwm.writers[name] = w
}

func (lwm *loggerWriterManager) Writers() map[string]io.Writer {
	return lwm.writers
}

func (lwm *loggerWriterManager) Write(p []byte) (n int, err error) {
	size := len(p)
	if size == 0 {
		return 0, nil
	}
	for _, w := range lwm.writers {
		n, err := w.Write(p)
		if err != nil {
			return n, err
		}
	}

	return size, nil
}

//we can use default logger to do something

func Logrus() *logrus.Logger {
	return defaultLogger.Logrus()
}

func SetLogLevel(l logrus.Level) {
	defaultLogger.SetLevel(l)
}

func Debug(args ...interface{}) {
	defaultLogger.Debug(args...)
}

func Info(args ...interface{}) {
	defaultLogger.Info(args...)
}

func Print(args ...interface{}) {
	defaultLogger.Print(args...)
}

func Warning(args ...interface{}) {
	defaultLogger.Warning(args...)
}

func Error(args ...interface{}) {
	defaultLogger.Error(args...)
}

func Fatal(args ...interface{}) {
	defaultLogger.Fatal(args...)
}

func Panic(args ...interface{}) {
	defaultLogger.Panic(args...)
}

func EnablePrint() {
	defaultLogger.EnablePrint()
}

func DisablePrint() {
	defaultLogger.DisablePrint()
}

func GetLogger() *Logger {
	defaultLogger = NewLogger()
	defaultLogger.EnablePrint()
	return defaultLogger
}
