package logger

import (
	"os"

	"github.com/Sirupsen/logrus"
	{% if cookiecutter.use_viper_config == "y" %}"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/config"{% else %}"github.com/spf13/pflag"{% endif %}
	"github.com/natefinch/lumberjack"
)

// Logger defines a set of methods for writing application logs. Derived from and
// inspired by logrus.Entry.
type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Debugln(args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Errorln(args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatalln(args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Infoln(args ...interface{})
	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
	Panicln(args ...interface{})
	Print(args ...interface{})
	Printf(format string, args ...interface{})
	Println(args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Warning(args ...interface{})
	Warningf(format string, args ...interface{})
	Warningln(args ...interface{})
	Warnln(args ...interface{})
}

var defaultLogger *logrus.Logger

func Log() Logger {
	return defaultLogger
}

func LogPtr() *logrus.Logger {
	return defaultLogger
}
{% if cookiecutter.use_viper_config == "y" %}
func init() {
	defaultLogger = newLogrusLogger(config.Config())
}

func NewLogger(cfg config.Provider) *logrus.Logger {
	return newLogrusLogger(cfg)
}

func newLogrusLogger(cfg config.Provider) *logrus.Logger {
	l := logrus.New()
	if cfg.GetBool("json_logs") {
		l.Formatter = new(logrus.JSONFormatter)
	}

	if cfg.GetString("logfile") != "" {
		l.Out = &lumberjack.Logger{
			Filename:   cfg.GetString("logfile"),
			MaxSize:    1,
			MaxBackups: 3,
			MaxAge:     1,
			Compress:   true,
			LocalTime:  true,
		}
	} else {
		l.Out = os.Stdout
	}

	switch cfg.GetString("loglevel") {
	case "debug":
		l.Level = logrus.DebugLevel
		l.SetReportCaller(true)
	case "info":
		l.Level = logrus.InfoLevel
	case "warn":
		l.Level = logrus.WarnLevel
	case "error":
		l.Level = logrus.ErrorLevel
	case "fatal":
		l.Level = logrus.FatalLevel
	case "panic":
		l.Level = logrus.PanicLevel
	default:
		l.Level = logrus.InfoLevel
	}
	return l
}

// ReloadLogrusLogger reloads config of a logger
func ReloadLogrusLoggerFromConfig(l *logrus.Logger, cfg config.Provider) {
	if cfg.GetBool("json_logs") {
		l.Formatter = new(logrus.JSONFormatter)
	}

	if cfg.GetString("logfile") != "" {
		l.Out = &lumberjack.Logger{
			Filename:   cfg.GetString("logfile"),
			MaxSize:    1,
			MaxBackups: 3,
			MaxAge:     1,
			Compress:   true,
			LocalTime:  true,
		}
	} else {
		l.Out = os.Stdout
	}

	switch cfg.GetString("loglevel") {
	case "debug":
		l.Level = logrus.DebugLevel
		l.SetReportCaller(true)
	case "info":
		l.Level = logrus.InfoLevel
	case "warn":
		l.Level = logrus.WarnLevel
	case "error":
		l.Level = logrus.ErrorLevel
	case "fatal":
		l.Level = logrus.FatalLevel
	case "panic":
		l.Level = logrus.PanicLevel
	default:
		l.Level = logrus.InfoLevel
	}
}
{% else %}
func init() {
	defaultLogger = newLogrusLogger()
}

func NewLogger() *logrus.Logger {
	return newLogrusLogger()
}

func newLogrusLogger() *logrus.Logger {
	l := logrus.New()
	return l
}

// ReloadLogrusLogger reloads logger config
func ReloadLogrusLoggerFromFlagSet(l *logrus.Logger, flagSet *pflag.FlagSet) {
	if jsonLogs, err := flagSet.GetBool("json_logs"); (err == nil) && jsonLogs {
		l.Formatter = new(logrus.JSONFormatter)
	}

	if logFile, err := flagSet.GetString("logfile"); (err == nil) && (logFile != "") {
		l.Formatter = new(logrus.JSONFormatter)
		l.Out = &lumberjack.Logger{
			Filename:   logFile,
			MaxSize:    1,
			MaxBackups: 3,
			MaxAge:     1,
			Compress:   true,
			LocalTime:  true,
		}
	} else {
		l.Out = os.Stdout
	}

	if logLevel, err := flagSet.GetString("loglevel"); (err == nil) && (logLevel != "") {
		switch logLevel {
		case "debug":
			l.Level = logrus.DebugLevel
			l.SetReportCaller(true)
		case "info":
			l.Level = logrus.InfoLevel
		case "warn":
			l.Level = logrus.WarnLevel
		case "error":
			l.Level = logrus.ErrorLevel
		case "fatal":
			l.Level = logrus.FatalLevel
		case "panic":
			l.Level = logrus.PanicLevel
		default:
			l.Level = logrus.InfoLevel
		}
	}
}
{% endif %}

type Fields map[string]interface{}

func (f Fields) With(k string, v interface{}) Fields {
	f[k] = v
	return f
}

func (f Fields) WithFields(f2 Fields) Fields {
	for k, v := range f2 {
		f[k] = v
	}
	return f
}

func WithFields(fields Fields) Logger {
	return defaultLogger.WithFields(logrus.Fields(fields))
}

// Debug package-level convenience method.
func Debug(args ...interface{}) {
	defaultLogger.Debug(args...)
}

// Debugf package-level convenience method.
func Debugf(format string, args ...interface{}) {
	defaultLogger.Debugf(format, args...)
}

// Debugln package-level convenience method.
func Debugln(args ...interface{}) {
	defaultLogger.Debugln(args...)
}

// Error package-level convenience method.
func Error(args ...interface{}) {
	defaultLogger.Error(args...)
}

// Errorf package-level convenience method.
func Errorf(format string, args ...interface{}) {
	defaultLogger.Errorf(format, args...)
}

// Errorln package-level convenience method.
func Errorln(args ...interface{}) {
	defaultLogger.Errorln(args...)
}

// Fatal package-level convenience method.
func Fatal(args ...interface{}) {
	defaultLogger.Fatal(args...)
}

// Fatalf package-level convenience method.
func Fatalf(format string, args ...interface{}) {
	defaultLogger.Fatalf(format, args...)
}

// Fatalln package-level convenience method.
func Fatalln(args ...interface{}) {
	defaultLogger.Fatalln(args...)
}

// Info package-level convenience method.
func Info(args ...interface{}) {
	defaultLogger.Info(args...)
}

// Infof package-level convenience method.
func Infof(format string, args ...interface{}) {
	defaultLogger.Infof(format, args...)
}

// Infoln package-level convenience method.
func Infoln(args ...interface{}) {
	defaultLogger.Infoln(args...)
}

// Panic package-level convenience method.
func Panic(args ...interface{}) {
	defaultLogger.Panic(args...)
}

// Panicf package-level convenience method.
func Panicf(format string, args ...interface{}) {
	defaultLogger.Panicf(format, args...)
}

// Panicln package-level convenience method.
func Panicln(args ...interface{}) {
	defaultLogger.Panicln(args...)
}

// Print package-level convenience method.
func Print(args ...interface{}) {
	defaultLogger.Print(args...)
}

// Printf package-level convenience method.
func Printf(format string, args ...interface{}) {
	defaultLogger.Printf(format, args...)
}

// Println package-level convenience method.
func Println(args ...interface{}) {
	defaultLogger.Println(args...)
}

// Warn package-level convenience method.
func Warn(args ...interface{}) {
	defaultLogger.Warn(args...)
}

// Warnf package-level convenience method.
func Warnf(format string, args ...interface{}) {
	defaultLogger.Warnf(format, args...)
}

// Warning package-level convenience method.
func Warning(args ...interface{}) {
	defaultLogger.Warning(args...)
}

// Warningf package-level convenience method.
func Warningf(format string, args ...interface{}) {
	defaultLogger.Warningf(format, args...)
}

// Warningln package-level convenience method.
func Warningln(args ...interface{}) {
	defaultLogger.Warningln(args...)
}

// Warnln package-level convenience method.
func Warnln(args ...interface{}) {
	defaultLogger.Warnln(args...)
}
