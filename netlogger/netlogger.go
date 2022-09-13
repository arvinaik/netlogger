package netlogger

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"
	"unsafe"
)

type Logger interface {
	Debug(ctx interface{}, v ...any)
	Debugln(ctx interface{}, v ...any)
	Debugf(ctx interface{}, format string, v ...any)

	Info(ctx interface{}, v ...any)
	Infoln(ctx interface{}, v ...any)
	Infof(format string, v ...any)

	Warn(ctx interface{}, v ...any)
	Warnln(ctx interface{}, v ...any)
	Warnf(ctx interface{}, format string, v ...any)

	Error(ctx interface{}, v ...any)
	Errorln(ctx interface{}, v ...any)
	Errorf(ctx interface{}, format string, v ...any)

	Fatal(ctx interface{}, v ...any)
	Fatalln(ctx interface{}, v ...any)
	Fatalf(ctx interface{}, format string, v ...any)
}

type Severity uint8

const (
	ALL   Severity = 0
	DEBUG Severity = 10
	INFO  Severity = 20
	WARN  Severity = 30
	ERROR Severity = 40
	FATAL Severity = 50
	OFF   Severity = 60
)

func (severity Severity) String() string {
	switch severity {
	case ALL:
		return "[ALL]"
	case DEBUG:
		return "[DEBUG]"
	case INFO:
		return "[INFO]"
	case WARN:
		return "[WARN]"
	case ERROR:
		return "[ERROR]"
	case FATAL:
		return "[FATAL]"
	case OFF:
		return "[OFF]"
	default:
		return fmt.Sprintf("%d", severity)
	}
}

type NetLogger struct {
	Severity Severity
}

func NewNetLogger(severity Severity) NetLogger {
	return NetLogger{
		Severity: severity,
	}
}

func (logger NetLogger) Debug(ctx context.Context, v ...any) {
	if logger.Severity <= DEBUG {
		printContextInternals(ctx)
		log.Print(createLogMessage(DEBUG, v...))
	}
}

func (logger NetLogger) Debugln(ctx context.Context, v ...any) {
	if logger.Severity <= DEBUG {
		printContextInternals(ctx)
		log.Println(createLogMessage(DEBUG, v...))
	}
}

func (logger NetLogger) Debugf(ctx context.Context, format string, v ...any) {
	if logger.Severity <= DEBUG {
		printContextInternals(ctx)
		log.Print(createLogMessagef(DEBUG, format, v...))
	}
}

func (logger NetLogger) Info(ctx context.Context, v ...any) {
	if logger.Severity <= INFO {
		printContextInternals(ctx)
		log.Print(createLogMessage(INFO, v...))
	}
}

func (logger NetLogger) Infoln(ctx context.Context, v ...any) {
	if logger.Severity <= INFO {
		printContextInternals(ctx)
		log.Println(createLogMessage(INFO, v...))
	}
}

func (logger NetLogger) Infof(ctx context.Context, format string, v ...any) {
	if logger.Severity <= INFO {
		printContextInternals(ctx)
		log.Print(createLogMessagef(INFO, format, v...))
	}
}

func (logger NetLogger) Warn(ctx context.Context, v ...any) {
	if logger.Severity <= WARN {
		printContextInternals(ctx)
		log.Print(createLogMessage(WARN, v...))
	}
}

func (logger NetLogger) Warnln(ctx context.Context, v ...any) {
	if logger.Severity <= WARN {
		printContextInternals(ctx)
		log.Println(createLogMessage(WARN, v...))
	}
}

func (logger NetLogger) Warnf(ctx context.Context, format string, v ...any) {
	if logger.Severity <= WARN {
		printContextInternals(ctx)
		log.Print(createLogMessagef(WARN, format, v...))
	}
}

func (logger NetLogger) Error(ctx context.Context, v ...any) {
	if logger.Severity <= ERROR {
		printContextInternals(ctx)
		log.Print(createLogMessage(ERROR, v...))
	}
}

func (logger NetLogger) Errorln(ctx context.Context, v ...any) {
	if logger.Severity <= ERROR {
		printContextInternals(ctx)
		log.Println(createLogMessage(ERROR, v...))
	}
}

func (logger NetLogger) Errorf(ctx context.Context, format string, v ...any) {
	if logger.Severity <= ERROR {
		printContextInternals(ctx)
		log.Print(createLogMessagef(ERROR, format, v...))
	}
}

func (logger NetLogger) Fatal(ctx context.Context, v ...any) {
	if logger.Severity <= FATAL {
		printContextInternals(ctx)
		log.Fatal(createLogMessage(FATAL, v...))
	}
}

func (logger NetLogger) Fatalln(ctx context.Context, v ...any) {
	if logger.Severity <= FATAL {
		printContextInternals(ctx)
		log.Fatalln(createLogMessage(FATAL, v...))
	}
}

func (logger NetLogger) Fatalf(ctx context.Context, format string, v ...any) {
	if logger.Severity <= FATAL {
		printContextInternals(ctx)
		log.Fatal(createLogMessagef(FATAL, format, v...))
	}
}

func createLogMessage(severity Severity, v ...any) string {
	sb := strings.Builder{}
	sb.WriteString(severity.String())
	sb.WriteString(" ")
	sb.WriteString(fmt.Sprint(v...))

	return sb.String()
}

func createLogMessagef(severity Severity, format string, v ...any) string {
	sb := strings.Builder{}
	sb.WriteString(severity.String())
	sb.WriteString(" ")
	sb.WriteString(fmt.Sprintf(format, v...))

	return sb.String()
}

func printContextInternals(ctx interface{}) {
	contextValues := reflect.ValueOf(ctx).Elem()
	contextKeys := reflect.TypeOf(ctx).Elem()

	if contextKeys.Kind() == reflect.Struct {
		for i := 0; i < contextValues.NumField(); i++ {
			reflectValue := contextValues.Field(i)
			reflectValue = reflect.NewAt(reflectValue.Type(), unsafe.Pointer(reflectValue.UnsafeAddr())).Elem()

			reflectField := contextKeys.Field(i)

			if reflectField.Name == "Context" {
				printContextInternals(reflectValue.Interface())
			} else {
				fmt.Printf("%+v", reflectField.Name)
				fmt.Printf(" : %+v\n", reflectValue.Interface())
			}
		}
	}
}
