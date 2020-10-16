package inter

import (
	"github.com/lanvard/syslog"
)

type Severity = syslog.Priority
type Facility = syslog.Priority

type Logger interface {
	SetApp(app Maker) Logger
	Clear()
	Log(severity Severity, message string, arguments ...interface{})
	LogWith(severity Severity, message string, data interface{})
	Emergency(message string, arguments ...interface{})
	EmergencyWith(message string, data interface{})
	Alert(message string, arguments ...interface{})
	AlertWith(message string, data interface{})
	Critical(message string, arguments ...interface{})
	CriticalWith(message string, data interface{})
	Error(message string, arguments ...interface{})
	ErrorWith(message string, data interface{})
	Warning(message string, arguments ...interface{})
	WarningWith(message string, data interface{})
	Notice(message string, arguments ...interface{})
	NoticeWith(message string, data interface{})
	Info(message string, arguments ...interface{})
	InfoWith(message string, data interface{})
	Debug(message string, arguments ...interface{})
	DebugWith(message string, data interface{})
}